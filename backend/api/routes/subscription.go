package routes

import (
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/m4tthewde/happening/backend/api/internal"
	"github.com/nicklaw5/helix"
)

func DeleteSubscription(request events.APIGatewayProxyRequest) (string, int) {
	id := request.QueryStringParameters["id"]

	_, err := internal.RemoveEventSubSubscription(id)
	if err != nil {
		log.Println(err)
		return "", 500
	}
	return "", 200
}

func GetSubscriptions(request events.APIGatewayProxyRequest) (string, int) {
	subscriptions := make([]GetSubscription, 0)

	resp, err := internal.GetEventSubSubscriptions()
	if err != nil {
		log.Println(err)
		return "", 500
	}

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

func PostSubscription(request events.APIGatewayProxyRequest) (string, int) {
	var newSubscriptionBody NewSubscriptionBody
	err := json.Unmarshal([]byte(request.Body), &newSubscriptionBody)
	if err != nil {
		log.Println(err)
		return "", 400

	}

	_, err = internal.CreateEventSubSubscription(&helix.EventSubSubscription{
		Type:    newSubscriptionBody.SubType,
		Version: "1",
		Condition: helix.EventSubCondition{
			BroadcasterUserID: newSubscriptionBody.TargetUserID,
		},
		Transport: helix.EventSubTransport{
			Method:   "webhook",
			Callback: os.Getenv("API_URL") + "/twitch",
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

type GetSubscriptionBody struct {
	Subscriptions []GetSubscription `json:"subscriptions"`
}

type GetSubscription struct {
	ID           string `json:"id"`
	TargetUserID string `json:"target_id"`
	SubType      string `json:"subscription_type"`
	Status       string `json:"status"`
}
