package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/m4tthewde/happening/backend/api/internal"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	lambda.Start(distributeRequest)
}

func distributeRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println(request.Path)

	switch request.Path {
	case "/api/subscription":
		return subscriptionRequest(request), nil
	case "/api/twitch":
		return twitchRequest(request), nil
	default:
		return events.APIGatewayProxyResponse{
			Body:       "",
			StatusCode: 404,
		}, nil
	}
}

func subscriptionRequest(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	switch request.HTTPMethod {
	case "POST":
		return internal.HandleNewSubscription(request)
	default:
		return events.APIGatewayProxyResponse{
			Body:       "",
			StatusCode: 405,
		}
	}
}

func twitchRequest(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	switch request.HTTPMethod {
	case "POST":
		return internal.HandleNewEventsubEvent(request)
	default:
		return events.APIGatewayProxyResponse{
			Body:       "",
			StatusCode: 405,
		}
	}
}
