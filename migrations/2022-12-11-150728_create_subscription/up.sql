-- Your SQL goes here

CREATE TABLE subscription (
    subscription_id SERIAL PRIMARY KEY,
    target_id VARCHAR NOT NULL,
    subscription_type VARCHAR NOT NULL
)