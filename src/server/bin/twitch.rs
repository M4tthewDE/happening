use std::env;

use dotenvy::dotenv;
use twitch_api::{helix::users::GetUsersRequest, types::UserIdRef, TwitchClient};
use twitch_oauth2::AppAccessToken;

pub struct TwitchApi<'a> {
    token: AppAccessToken,
    client: TwitchClient<'a, reqwest::Client>,
}

impl TwitchApi<'_> {
    pub async fn new() -> TwitchApi<'static> {
        let client: TwitchClient<reqwest::Client> = TwitchClient::default();
        TwitchApi {
            token: generate_token().await.unwrap(),
            client,
        }
    }

    pub async fn is_valid_user_id(&self, id: &str) -> bool {
        let ids: &[_] = &[UserIdRef::from_str(id)];
        let req = GetUsersRequest::ids(ids);
        let response = &self.client.helix.req_get(req, &self.token).await.unwrap();

        !response.data.is_empty()
    }
}

pub async fn generate_token() -> anyhow::Result<AppAccessToken> {
    dotenv().ok();

    let client_id = env::var("TWITCH_CLIENT_ID")
        .ok()
        .map(twitch_oauth2::ClientId::new)
        .expect("TWITCH_CLIENT_ID not set");

    let client_secret = env::var("TWITCH_CLIENT_SECRET")
        .ok()
        .map(twitch_oauth2::ClientSecret::new)
        .expect("TWITCH_CLIENT_SECRET not set");

    let reqwest = reqwest::Client::builder().build()?;
    let token = twitch_oauth2::AppAccessToken::get_app_access_token(
        &reqwest,
        client_id,
        client_secret,
        vec![],
    )
    .await?;

    Ok(token)
}
