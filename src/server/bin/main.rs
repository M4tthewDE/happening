#![feature(proc_macro_hygiene, decl_macro)]

use happening::{create_subscription, establish_connection};
use rocket::fairing::AdHoc;
use rocket_contrib::json::Json;
use rocket_cors::{AllowedOrigins, CorsOptions};

#[macro_use]
extern crate rocket;
#[macro_use]
extern crate serde_derive;

fn main() {
    rocket().launch();
}

fn rocket() -> rocket::Rocket {
    rocket::ignite()
        .mount("/", routes![new_subscription])
        .attach(AdHoc::on_attach("Cors config", |rocket| {
            match rocket.config().get_bool("cors_allow_all").unwrap_or(false) {
                true => {
                    let cors = CorsOptions::default().allowed_origins(AllowedOrigins::all());

                    Ok(rocket.attach(cors.to_cors().unwrap()))
                }
                false => Ok(rocket),
            }
        }))
}

#[derive(Serialize, Deserialize, Debug)]
pub struct Subscription {
    target_id: String,
    subscription_type: String,
}

#[post("/api/subscription", format = "json", data = "<subscription>")]
fn new_subscription(subscription: Json<Subscription>) {
    let mut conn = establish_connection();
    let target_id = &subscription.target_id;
    let subscription_type = &subscription.subscription_type;

    create_subscription(&mut conn, target_id.as_str(), subscription_type.as_str())
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
