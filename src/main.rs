#![feature(proc_macro_hygiene, decl_macro)]

use rocket::http::RawStr;
use rocket_contrib::json::{Json, JsonValue};

#[macro_use]
extern crate rocket;
#[macro_use]
extern crate rocket_contrib;
#[macro_use]
extern crate serde_derive;

fn main() {
    rocket().launch();
}

fn rocket() -> rocket::Rocket {
    rocket::ignite().mount("/", routes![new_subscription])
}

#[derive(Serialize, Deserialize, Debug)]
struct Subscription {
    target_id: String,
    subscription_type: SubscriptionType,
}

#[derive(Serialize, Deserialize, Debug)]
enum SubscriptionType {
    Follow,
}

#[post("/api/subscription", format = "json", data = "<subscription>")]
fn new_subscription(subscription: Json<Subscription>) {}

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
