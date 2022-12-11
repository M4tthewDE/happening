use std::env;

use anyhow::{Context, Result};
use dotenvy::dotenv;
use twitch_api::{helix::users::GetUsersRequest, types::UserIdRef, TwitchClient};
use twitch_oauth2::{AccessToken, AppAccessToken, ClientId, ClientSecret};

#[derive(Clone)]
pub struct TwitchApi<'a> {
    secret: ClientSecret,
    client_id: ClientId,
    client: TwitchClient<'a, reqwest::Client>,
}

impl TwitchApi<'_> {
    pub async fn new() -> Result<TwitchApi<'static>> {
        dotenv().with_context(|| "Error loading .env")?;

        let secret = env::var("TWITCH_CLIENT_SECRET")
            .ok()
            .map(twitch_oauth2::ClientSecret::new)
            .with_context(|| "TWITCH_CLIENT_SECRET not set")?;

        let client_id = env::var("TWITCH_CLIENT_ID")
            .ok()
            .map(twitch_oauth2::ClientId::new)
            .with_context(|| "TWITCH_CLIENT_ID not set")?;

        let client: TwitchClient<reqwest::Client> = TwitchClient::default();

        Ok(TwitchApi {
            secret,
            client,
            client_id,
        })
    }

    pub async fn generate_token(&self) -> Result<AppAccessToken> {
        let reqwest = reqwest::Client::builder().build()?;
        let token = twitch_oauth2::AppAccessToken::get_app_access_token(
            &reqwest,
            self.client_id.clone(),
            self.secret.clone(),
            vec![],
        )
        .await?;

        Ok(token)
    }

    pub async fn is_valid_user_id(&self, token: String, id: &str) -> Result<bool> {
        let token = AppAccessToken::from_existing_unchecked(
            AccessToken::new(token),
            None,
            self.client_id.clone(),
            self.secret.clone(),
            None,
            None,
        );

        let ids: &[_] = &[UserIdRef::from_str(id)];
        let req = GetUsersRequest::ids(ids);
        let response = &self
            .client
            .helix
            .req_get(req, &token)
            .await
            .with_context(|| "Error fetching user")?;

        Ok(!response.data.is_empty())
    }
}
