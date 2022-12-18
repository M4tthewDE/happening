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
			Body:       "",
			StatusCode: 400,
		}, err
	}

	resp, err := json.Marshal(newSubscriptionBody)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
		}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(resp),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}

type NewSubscriptionBody struct {
	TargetUserID string `json:"target_id"`
	SubType      string `json:"subscription_type"`
}
