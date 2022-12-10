#![feature(proc_macro_hygiene, decl_macro)]

use happening::{create_subscription, establish_connection};
use rocket::fairing::AdHoc;
use rocket_contrib::json::Json;
use rocket_cors::{AllowedOrigins, CorsOptions};
use types::Subscription;

#[macro_use]
extern crate rocket;
#[macro_use]
extern crate serde_derive;

mod twitch;
mod types;

#[tokio::main]
async fn main() {
    let token = twitch::generate_token().await.unwrap();
    rocket().launch();
}

fn rocket() -> rocket::Rocket {
    rocket::ignite()
        .mount("/", routes![new_subscription])
        .attach(AdHoc::on_attach("CORS", |rocket| {
            match rocket.config().get_bool("cors_allow_all").unwrap_or(false) {
                true => {
                    let cors = CorsOptions::default().allowed_origins(AllowedOrigins::all());

                    Ok(rocket.attach(cors.to_cors().unwrap()))
                }
                false => Ok(rocket),
            }
        }))
}

#[post("/api/subscription", format = "json", data = "<subscription>")]
fn new_subscription(subscription: Json<Subscription>) {
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
        local::Client,
    };

    #[test]
    fn new_subscription() {
        let client = Client::new(rocket()).expect("valid rocket instance");
        let response = client
            .post("/api/subscription")
            .header(ContentType::JSON)
            .body(r#"{"target_id": "1234", "subscription_type": "Follow"}"#)
            .dispatch();

        assert_eq!(response.status(), Status::Ok);
    }
}
