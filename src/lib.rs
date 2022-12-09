use std::env;

use diesel::{Connection, RunQueryDsl, SqliteConnection};
use dotenvy::dotenv;

use crate::models::NewSubscription;

pub mod models;
pub mod schema;

pub fn establish_connection() -> SqliteConnection {
    dotenv().ok();

    let db_url = env::var("DATABASE_URL").expect("DATABASE_URL must be set");
    SqliteConnection::establish(&db_url).unwrap_or_else(|_| panic!("Error connecting to {db_url}"))
}

pub fn create_subscription(conn: &mut SqliteConnection, target_id: &str, subscription_type: &str) {
    use crate::schema::subscription;

    let new_subscription = NewSubscription {
        target_id,
        subscription_type,
    };

    diesel::insert_into(subscription::table)
        .values(&new_subscription)
        .execute(conn)
        .expect("Error saving subscription");
}
