package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var newSubscriptionBody NewSubscriptionBody
	err := json.Unmarshal([]byte(request.Body), &newSubscriptionBody)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Hello world",
			StatusCode: 400,
		}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       "Hello world",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}

type NewSubscriptionBody struct {
	TargetUserID string `json:"target_user_id"`
	SubType      string `json:"sub_type"`
}
