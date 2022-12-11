#![feature(proc_macro_hygiene, decl_macro)]

use happening::{create_subscription, establish_connection};
use rocket::serde::json::Json;
use rocket_cors::{AllowedOrigins, CorsOptions};
use types::Subscription;

#[macro_use]
extern crate rocket;

mod twitch;
mod types;

#[launch]
fn rocket() -> _ {
    let twitch_api = twitch::TwitchApi::new();

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

    rocket.mount("/", routes![new_subscription])
}

#[post("/api/subscription", format = "json", data = "<subscription>")]
fn new_subscription(subscription: Json<Subscription<'_>>) {
    let mut conn = establish_connection();

    let target_id = &subscription.target_id;
    twitch::validate_user_id(target_id);
    let subscription_type = subscription.subscription_type.to_string();

    create_subscription(&mut conn, target_id, &subscription_type);
}

#[cfg(test)]
mod test {
    use super::rocket;
    use rocket::{
        http::{ContentType, Status},
        local::blocking::Client,
    };

    #[test]
    fn new_subscription() {
        let client = Client::tracked(rocket()).expect("valid rocket instance");
        let response = client
            .post("/api/subscription")
            .header(ContentType::JSON)
            .body(r#"{"target_id": "1234", "subscription_type": "Follow"}"#)
            .dispatch();

        assert_eq!(response.status(), Status::Ok);
    }
}
