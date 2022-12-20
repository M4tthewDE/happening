package internal

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func HandleNewSubscription(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	var newSubscriptionBody NewSubscriptionBody
	err := json.Unmarshal([]byte(request.Body), &newSubscriptionBody)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "",
			StatusCode: 400,
		}
	}

	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx, func(o *config.LoadOptions) error {
		o.Region = "us-east-1"
		return nil
	})
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "",
			StatusCode: 500,
		}
	}

	client := dynamodb.NewFromConfig(cfg)
	d := NewDao(client)

	_, err = d.GetAuth(ctx)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "",
			StatusCode: 500,
		}
	}

	resp, err := json.Marshal(newSubscriptionBody)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "",
			StatusCode: 400,
		}
	}

	return events.APIGatewayProxyResponse{
		Body:       string(resp),
		StatusCode: 200,
	}
}

type NewSubscriptionBody struct {
	TargetUserID string `json:"target_id"`
	SubType      string `json:"subscription_type"`
}
