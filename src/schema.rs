// @generated automatically by Diesel CLI.

diesel::table! {
    subscription (subscription_id) {
        subscription_id -> Int4,
        target_id -> Varchar,
        subscription_type -> Varchar,
        eventsub_id -> Varchar,
    }
}
