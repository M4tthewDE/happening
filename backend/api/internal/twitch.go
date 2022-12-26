package internal

import (
	"context"
	"os"

	"github.com/nicklaw5/helix"
)

func getClient() (*helix.Client, error) {
	token, err := GetAuth(context.TODO())
	if err != nil {
		return nil, err
	}

	return helix.NewClient(&helix.Options{
		ClientID:       os.Getenv("TWITCH_CLIENT_ID"),
		AppAccessToken: token,
	})
}

func ValidateToken(token string) (bool, *helix.ValidateTokenResponse, error) {
	client, err := helix.NewClient(&helix.Options{
		ClientID:       os.Getenv("TWITCH_CLIENT_ID"),
		AppAccessToken: token,
	})
	if err != nil {
		return false, nil, err
	}

	return client.ValidateToken(token)
}

func GetUsers(params *helix.UsersParams) (*helix.UsersResponse, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetUsers(params)
}

func RemoveEventSubSubscription(id string) (*helix.RemoveEventSubSubscriptionParamsResponse, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.RemoveEventSubSubscription(id)
}

func GetEventSubSubscriptions() (*helix.EventSubSubscriptionsResponse, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetEventSubSubscriptions(&helix.EventSubSubscriptionsParams{})
}

func CreateEventSubSubscription(sub *helix.EventSubSubscription) (*helix.EventSubSubscriptionsResponse, error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.CreateEventSubSubscription(sub)
}
