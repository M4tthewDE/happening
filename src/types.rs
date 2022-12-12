use std::fmt::Display;

use rocket::serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug)]
#[serde(crate = "rocket::serde")]
pub struct Subscription<'r> {
    pub target_id: &'r str,
    pub subscription_type: SubscriptionType,
}

#[derive(Serialize, Deserialize, Debug)]
#[serde(crate = "rocket::serde")]
pub enum SubscriptionType {
    Follow,
    Sub,
}

impl Display for SubscriptionType {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            SubscriptionType::Follow => write!(f, "Follow"),
            SubscriptionType::Sub => write!(f, "Sub"),
        }
    }
}

impl SubscriptionType {
    pub fn get_twitch_type(&self) -> String {
        match self {
            SubscriptionType::Follow => "channel.follow".to_string(),
            SubscriptionType::Sub => "channel.subscribe".to_string(),
        }
    }
}
