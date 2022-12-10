use std::fmt::Display;

#[derive(Serialize, Deserialize, Debug)]
pub struct Subscription {
    pub target_id: String,
    pub subscription_type: SubscriptionType,
}

#[derive(Serialize, Deserialize, Debug)]
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
