use std::env;

use diesel::{
    r2d2::{self, ConnectionManager, Pool},
    PgConnection, RunQueryDsl,
};
use dotenvy::dotenv;

use crate::models::NewSubscription;
use crate::schema::subscription;

pub mod models;
pub mod schema;

pub type DbPool = Pool<ConnectionManager<PgConnection>>;

pub struct Db {
    pool: DbPool,
}

impl Db {
    pub fn new() -> Db {
        dotenv().ok();

        let db_url = env::var("DATABASE_URL").expect("DATABASE_URL must be set");
        let manager = ConnectionManager::<PgConnection>::new(&db_url);
        let pool = r2d2::Pool::builder()
            .build(manager)
            .expect("Failed to create pool.");

        Db { pool }
    }

    pub fn save_subscription(&self, target_id: &str, subscription_type: &str) {
        let new_subscription = NewSubscription {
            target_id,
            subscription_type,
        };

        diesel::insert_into(subscription::table)
            .values(&new_subscription)
            .execute(&mut self.pool.get().unwrap())
            .expect("Error saving subscription");
    }
}
