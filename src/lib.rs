use std::env;

use diesel::{
    r2d2::{self, ConnectionManager, Pool},
    PgConnection, RunQueryDsl,
};
use dotenvy::dotenv;

use crate::models::NewSubscription;

pub mod models;
pub mod schema;

pub fn establish_connection_pool() -> Pool<ConnectionManager<PgConnection>> {
    dotenv().ok();

    let db_url = env::var("DATABASE_URL").expect("DATABASE_URL must be set");
    let manager = ConnectionManager::<PgConnection>::new(&db_url);
    let pool = r2d2::Pool::builder()
        .build(manager)
        .expect("Failed to create pool.");
    pool
}

pub fn create_subscription(conn: &mut PgConnection, target_id: &str, subscription_type: &str) {
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
