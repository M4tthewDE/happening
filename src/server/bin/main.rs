#![feature(proc_macro_hygiene, decl_macro)]

use happening::{create_subscription, establish_connection};
use rocket::{serde::json::Json, Build, Rocket, State};
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

    rocket
        .mount("/", routes![new_subscription])
        .manage(twitch_api)
}

#[post("/api/subscription", format = "json", data = "<subscription>")]
fn new_subscription(subscription: Json<Subscription>, twitch_api: &State<TwitchApi>) {
    let mut conn = establish_connection();

    let target_id = &subscription.target_id;
    twitch_api.validate_user_id(&target_id);

    let subscription_type = subscription.subscription_type.to_string();
    create_subscription(&mut conn, target_id, &subscription_type);
}

#[cfg(test)]
mod test {
    use super::rocket;
    use rocket::{
        http::{ContentType, Status},
        local::asynchronous::Client,
    };

    #[rocket::async_test]
    async fn new_subscription() {
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
