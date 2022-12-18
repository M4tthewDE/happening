package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/m4tthewde/happening/backend/internal"
)

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	response, status := internal.HandleNewSubscription(request.Body)

	return events.APIGatewayProxyResponse{
		Body:       response,
		StatusCode: status,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}

type NewSubscriptionBody struct {
	TargetUserID string `json:"target_id"`
	SubType      string `json:"subscription_type"`
}
