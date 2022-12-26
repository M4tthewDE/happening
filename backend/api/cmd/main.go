package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/m4tthewde/happening/backend/api/routes"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	lambda.Start(distributeRequest)
}

func distributeRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var body string
	var status int

	switch request.PathParameters["proxy"] {
	case "api/subscription":
		body, status = subscriptionRequest(request)
	case "api/user":
		body, status = userRequest(request)
	case "api/twitch":
		body, status = twitchRequest(request)
	case "api/permissions":
		body, status = permissionsRequest(request)
	default:
		body = ""
		status = 404
	}

	headers := make(map[string]string, 0)
	headers["Access-Control-Allow-Origin"] = "*"
	headers["Access-Control-Allow-Methods"] = "GET, POST, DELETE"
	headers["Access-Control-Allow-Headers"] = "Origin, X-Requested-With, Content-Type, Accept"

	return events.APIGatewayProxyResponse{
		Body:       body,
		StatusCode: status,
		Headers:    headers,
	}, nil
}

func subscriptionRequest(request events.APIGatewayProxyRequest) (string, int) {
	switch request.HTTPMethod {
	case "POST":
		return routes.PostSubscription(request)
	case "GET":
		return routes.GetSubscriptions(request)
	case "DELETE":
		return routes.DeleteSubscription(request)
	case "OPTIONS":
		return "", 200
	default:
		return "", 405
	}
}

func userRequest(request events.APIGatewayProxyRequest) (string, int) {
	switch request.HTTPMethod {
	case "GET":
		return routes.GetUser(request)
	case "OPTIONS":
		return "", 200
	default:
		return "", 405
	}
}

func twitchRequest(request events.APIGatewayProxyRequest) (string, int) {
	switch request.HTTPMethod {
	case "POST":
		return routes.HandleNewEventsubEvent(request)
	default:
		return "", 405
	}
}

func permissionsRequest(request events.APIGatewayProxyRequest) (string, int) {
	switch request.HTTPMethod {
	case "GET":
		return routes.GetPermissions(request)
	case "OPTIONS":
		return "", 200
	default:
		return "", 405
	}
}
