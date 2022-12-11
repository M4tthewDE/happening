#![feature(proc_macro_hygiene, decl_macro)]

use happening::Db;
use rocket::{http::Status, serde::json::Json, Build, Rocket, State};
use rocket_cors::{AllowedOrigins, CorsOptions};
use twitch::TwitchApi;
use types::Subscription;

#[macro_use]
extern crate rocket;

mod twitch;
mod types;

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

    let twitch_api = twitch::TwitchApi::new().await;
    let db = Db::new();

    rocket
        .mount("/", routes![new_subscription])
        .manage(twitch_api)
        .manage(db)
}

#[post("/api/subscription", format = "json", data = "<subscription>")]
async fn new_subscription(
    subscription: Json<Subscription<'_>>,
    twitch_api: &State<TwitchApi<'_>>,
    db: &State<Db>,
) -> (Status, &'static str) {
    let target_id = &subscription.target_id;

    if !twitch_api.is_valid_user_id(target_id).await {
        return (Status::BadRequest, "Target user does not exist");
    }

    let subscription_type = subscription.subscription_type.to_string();
    db.save_subscription(target_id, &subscription_type);

    (Status::Ok, "subscription created")
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
