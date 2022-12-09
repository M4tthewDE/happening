use crate::schema::subscription;
use diesel::{Insertable, Queryable};

#[derive(Queryable)]
pub struct Subscription {
    pub id: i32,
    pub target_id: String,
    pub subscription_type: String,
}

#[derive(Insertable)]
#[diesel(table_name = subscription)]
pub struct NewSubscription<'a> {
    pub target_id: &'a str,
    pub subscription_type: &'a str,
}
