-- Your SQL goes here

CREATE TABLE subscription (
    subscription_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    target_id text NOT NULL,
    subscription_type text NOT NULL
)