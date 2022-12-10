use std::env;

use dotenvy::dotenv;
use twitch_oauth2::AppAccessToken;

pub fn validate_user_id(id: &str) {
    todo!("implement");
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
