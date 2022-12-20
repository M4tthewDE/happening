package internal

import (
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/nicklaw5/helix/v2"
)

func HandleNewEventsubEvent(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	headers := make(map[string][]string)

	for key, value := range request.Headers {
		v := make([]string, 1)
		v = append(v, value)
		headers[key] = v
	}

	if !helix.VerifyEventSubNotification(os.Getenv("EVENTSUB_SECRET"), headers, request.Body) {
		return events.APIGatewayProxyResponse{
			Body:       "",
			StatusCode: 403,
		}
	}

	return events.APIGatewayProxyResponse{
		Body:       "",
		StatusCode: 200,
	}
}
