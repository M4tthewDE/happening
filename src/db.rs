use std::{env, time::Duration};

use anyhow::{Context, Result};
use diesel::{
    r2d2::{self, ConnectionManager, Pool},
    PgConnection, RunQueryDsl,
};
use dotenvy::dotenv;
use redis::Commands;

use crate::{models::NewSubscription, schema::subscription, twitch::AppAccessToken};

pub type DbPool = Pool<ConnectionManager<PgConnection>>;

pub struct Db {
    pool: DbPool,
}

impl Db {
    pub fn new() -> Result<Db> {
        dotenv().with_context(|| "Error loading .env")?;

        let db_url =
            env::var("DATABASE_URL").with_context(|| "DATABASE_URL must be set".to_string())?;

        let manager = ConnectionManager::<PgConnection>::new(db_url);
        let pool = r2d2::Pool::builder()
            .build(manager)
            .with_context(|| "DATABASE_URL must be set")?;

        Ok(Db { pool })
    }

    pub fn save_subscription(&self, target_id: &str, subscription_type: &str) -> Result<()> {
        let new_subscription = NewSubscription {
            target_id,
            subscription_type,
        };

        diesel::insert_into(subscription::table)
            .values(&new_subscription)
            .execute(&mut self.pool.get()?)
            .with_context(|| "Error saving subscription")?;

        Ok(())
    }
}

pub struct RedisClient {
    con: redis::Connection,
}

impl RedisClient {
    pub fn new() -> Result<RedisClient> {
        let client = redis::Client::open("redis://127.0.0.1/")
            .with_context(|| "failed to connect to redis")?;

        let con = client.get_connection()?;

        Ok(RedisClient { con })
    }

    pub fn save_token(&mut self, token: AppAccessToken) -> Result<()> {
        self.con
            .set("token", token.access_token)
            .with_context(|| "failed to set token")?;

        self.con
            .set("expires_in", token.expires_in.as_secs())
            .with_context(|| "failed to set token")?;

        Ok(())
    }

    pub fn get_token(&mut self) -> Result<String> {
        let token: String = self
            .con
            .get("token")
            .with_context(|| "failed to get token")?;
        Ok(token)
    }

    pub fn get_expires_in(&mut self) -> Result<Duration> {
        let expires_in: u64 = self
            .con
            .get("expires_in")
            .with_context(|| "failed to get token")?;

        Ok(Duration::from_secs(expires_in))
    }
}
