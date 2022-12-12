use std::{env, time::Duration};

use anyhow::{bail, Context, Result};
use dotenvy::dotenv;
use serde::Deserialize;

#[derive(Clone)]
pub struct TwitchApi {
    secret: String,
    client_id: String,
    reqwest: reqwest::Client,
}

impl TwitchApi {
    pub async fn new() -> Result<TwitchApi> {
        dotenv().with_context(|| "Error loading .env")?;

        let secret = env::var("TWITCH_CLIENT_SECRET")
            .ok()
            .with_context(|| "TWITCH_CLIENT_SECRET not set")?;

        let client_id = env::var("TWITCH_CLIENT_ID")
            .ok()
            .with_context(|| "TWITCH_CLIENT_ID not set")?;

        let reqwest = reqwest::Client::builder().build()?;

        Ok(TwitchApi {
            secret,
            client_id,
            reqwest,
        })
    }

    pub async fn generate_token(&self) -> Result<AppAccessToken> {
        let res = self
            .reqwest
            .post("https://id.twitch.tv/oauth2/token")
            .query(&[
                ("client_id", &self.client_id),
                ("client_secret", &self.secret),
                ("grant_type", &"client_credentials".to_string()),
            ])
            .send()
            .await
            .with_context(|| "token request failed")?;

        if !res.status().is_success() {
            bail!("error requesting token: {}", res.status());
        }

        let body: TokenBody = res.json().await?;
        Ok(AppAccessToken {
            access_token: body.access_token,
            expires_in: Duration::from_secs(body.expires_in),
        })
    }

    pub async fn is_valid_user_id(&self, token: String, id: &str) -> Result<bool> {
        let res = self
            .reqwest
            .get("https://api.twitch.tv/helix/users")
            .query(&[("id", id)])
            .header("Authorization", &format!("Bearer {token}"))
            .header("Client-Id", &self.client_id)
            .send()
            .await
            .with_context(|| "user request failed")?;

        if !res.status().is_success() {
            bail!("error fetching user: {}", res.status());
        }

        let body: UserBody = res.json().await?;
        Ok(!body.data.is_empty())
    }
}

#[derive(Deserialize)]
struct UserBody {
    data: Vec<UserData>,
}

#[derive(Deserialize)]
struct UserData {}

#[derive(Deserialize)]
struct TokenBody {
    access_token: String,
    expires_in: u64,
}

pub struct AppAccessToken {
    pub access_token: String,
    pub expires_in: Duration,
}
