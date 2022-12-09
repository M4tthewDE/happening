// @generated automatically by Diesel CLI.

diesel::table! {
    subscription (subscription_id) {
        subscription_id -> Integer,
        target_id -> Text,
        subscription_type -> Text,
    }
}
