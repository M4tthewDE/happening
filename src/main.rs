#![feature(proc_macro_hygiene, decl_macro)]

use std::{sync::Mutex, time::Duration};

use happening::{Db, RedisClient};
use rocket::{http::Status, serde::json::Json, Build, Rocket, State};
use rocket_cors::{AllowedOrigins, CorsOptions};
use tokio::{task, time};
use twitch::TwitchApi;
use twitch_oauth2::TwitchToken;
use types::Subscription;

#[macro_use]
extern crate rocket;

mod twitch;
mod types;

struct ApiState<'a> {
    twitch_api: TwitchApi<'a>,
    redis_client: Mutex<RedisClient>,
    db: Db,
}

#[rocket::main]
async fn main() -> Result<(), rocket::Error> {
    let _rocket = rocket().await.launch().await?;
    Ok(())
}

async fn rocket() -> Rocket<Build> {
    let rocket = if rocket::Config::figment()
        .extract_inner("cors_allow_all")
        .unwrap_or(false)
    {
        rocket::build().attach(
            CorsOptions::default()
                .allowed_origins(AllowedOrigins::all())
                .to_cors()
                .unwrap(),
        )
    } else {
        rocket::build()
    };

    // save access token in redis
    let twitch_api = twitch::TwitchApi::new().await.unwrap();
    let mut redis_client = RedisClient::new().unwrap();
    let token = twitch_api.generate_token().await.unwrap();
    redis_client.set_token(token.access_token.as_str()).unwrap();
    redis_client.set_expires_in(token.expires_in()).unwrap();

    // start token refresh loop
    let new_twitch_api = twitch_api.clone();
    task::spawn(async move {
        let mut redis_client = RedisClient::new().unwrap();

        loop {
            time::sleep(Duration::from_secs(60)).await;

            let expires_in = redis_client.get_expires_in().unwrap();
            if expires_in.as_secs() < 300 {
                let new_token = new_twitch_api.generate_token().await.unwrap();
                redis_client
                    .set_token(new_token.access_token.as_str())
                    .unwrap();
                redis_client.set_expires_in(new_token.expires_in()).unwrap();
            }
        }
    });

    let db = Db::new().unwrap();

    let api_state = ApiState {
        twitch_api,
        redis_client: Mutex::new(redis_client),
        db,
    };

    rocket
        .mount("/", routes![new_subscription])
        .manage(api_state)
}

#[post("/api/subscription", format = "json", data = "<subscription>")]
async fn new_subscription(
    subscription: Json<Subscription<'_>>,
    api_state: &State<ApiState<'_>>,
) -> (Status, &'static str) {
    let target_id = &subscription.target_id;
    let token = match api_state.redis_client.lock().unwrap().get_token() {
        Ok(token) => token,
        Err(_) => {
            return (Status::InternalServerError, "error fetching token");
        }
    };

    match api_state
        .twitch_api
        .is_valid_user_id(token, target_id)
        .await
    {
        Ok(is_valid) => {
            if !is_valid {
                return (Status::BadRequest, "Target user does not exist");
            }
        }
        Err(_) => return (Status::InternalServerError, "error accesing twitch api"),
    };

    let subscription_type = subscription.subscription_type.to_string();

    match api_state
        .db
        .save_subscription(target_id, &subscription_type)
    {
        Err(_) => (Status::InternalServerError, "error accesing database"),
        Ok(_) => (Status::Ok, "subscription created"),
    }
}

#[cfg(test)]
mod test {
    use super::rocket;
    use rocket::{
        http::{ContentType, Status},
        local::asynchronous::Client,
    };

    #[rocket::async_test]
    async fn new_subscription_invalid_target() {
        let rocket = rocket().await;
        let client = Client::tracked(rocket)
            .await
            .expect("valid rocket instance");

        let response = client
            .post("/api/subscription")
            .header(ContentType::JSON)
            .body(r#"{"target_id": "1234556", "subscription_type": "Follow"}"#)
            .dispatch()
            .await;

        assert_eq!(response.status(), Status::BadRequest);
    }

    #[rocket::async_test]
    async fn new_subscription_valid_target() {
        let rocket = rocket().await;
        let client = Client::tracked(rocket)
            .await
            .expect("valid rocket instance");

        let response = client
            .post("/api/subscription")
            .header(ContentType::JSON)
            .body(r#"{"target_id": "1234", "subscription_type": "Follow"}"#)
            .dispatch()
            .await;

        assert_eq!(response.status(), Status::Ok);
    }
}
