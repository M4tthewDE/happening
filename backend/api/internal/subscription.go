package internal

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/nicklaw5/helix"
)

type GetSubscriptionBody struct {
	Subscriptions []GetSubscription `json:"subscriptions"`
}

type GetSubscription struct {
	ID           string `json:"id"`
	TargetUserID string `json:"target_id"`
	SubType      string `json:"subscription_type"`
	Status       string `json:"status"`
}

func GetSubscriptions(request events.APIGatewayProxyRequest) (string, int) {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx, func(o *config.LoadOptions) error {
		o.Region = "us-east-1"
		return nil
	})
	if err != nil {
		log.Println(err)
		return "", 500
	}

	dbClient := dynamodb.NewFromConfig(cfg)
	d := NewDao(dbClient)

	token, err := d.GetAuth(ctx)
	if err != nil {
		log.Println(err)
		return "", 500
	}

	client, err := helix.NewClient(&helix.Options{
		ClientID:       os.Getenv("TWITCH_CLIENT_ID"),
		AppAccessToken: token,
	})
	if err != nil {
		log.Println(err)
		return "", 500
	}

	resp, err := client.GetEventSubSubscriptions(&helix.EventSubSubscriptionsParams{})
	if err != nil {
		log.Println(err)
		return "", 500
	}

	subscriptions := make([]GetSubscription, 0)
	for _, sub := range resp.Data.EventSubSubscriptions {
		subscriptions = append(subscriptions, GetSubscription{
			ID:           sub.ID,
			TargetUserID: sub.Condition.BroadcasterUserID,
			SubType:      sub.Type,
			Status:       sub.Status,
		})
	}

	body := GetSubscriptionBody{
		Subscriptions: subscriptions,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Println(err)
		return "", 500
	}

	return string(bodyBytes), 200
}

func HandleNewSubscription(request events.APIGatewayProxyRequest) (string, int) {
	var newSubscriptionBody NewSubscriptionBody
	err := json.Unmarshal([]byte(request.Body), &newSubscriptionBody)
	if err != nil {
		log.Println(err)
		return "", 400

	}

	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx, func(o *config.LoadOptions) error {
		o.Region = "us-east-1"
		return nil
	})
	if err != nil {
		log.Println(err)
		return "", 500
	}

	dbClient := dynamodb.NewFromConfig(cfg)
	d := NewDao(dbClient)

	token, err := d.GetAuth(ctx)
	if err != nil {
		log.Println(err)
		return "", 500
	}

	client, err := helix.NewClient(&helix.Options{
		ClientID:       os.Getenv("TWITCH_CLIENT_ID"),
		AppAccessToken: token,
	})
	if err != nil {
		log.Println(err)
		return "", 500
	}

	sub_type, err := newSubscriptionBody.getTypeString()
	if err != nil {
		log.Println(err)
		return "", 500
	}

	_, err = client.CreateEventSubSubscription(&helix.EventSubSubscription{
		Type:    sub_type,
		Version: "1",
		Condition: helix.EventSubCondition{
			BroadcasterUserID: newSubscriptionBody.TargetUserID,
		},
		Transport: helix.EventSubTransport{
			Method:   "webhook",
			Callback: "https://happening.fdm.com.de/api/twitch",
			Secret:   os.Getenv("EVENTSUB_SECRET"),
		},
	})
	if err != nil {
		log.Println(err)
		return "", 500
	}

	return "", 200
}

type NewSubscriptionBody struct {
	TargetUserID string `json:"target_id"`
	SubType      string `json:"subscription_type"`
}

var ErrSubTypeNotFound = errors.New("sub type not found")

func (s NewSubscriptionBody) getTypeString() (string, error) {
	switch s.SubType {
	case "FOLLOW":
		return "channel.follow", nil
	case "SUB":
		return "channel.subscribe", nil
	default:
		return "", ErrSubTypeNotFound
	}
}
