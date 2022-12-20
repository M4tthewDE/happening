package internal

import (
	"bytes"
	"encoding/json"
	"log"
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

	var vals eventSubNotification
	err := json.NewDecoder(bytes.NewReader([]byte(request.Body))).Decode(&vals)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "",
			StatusCode: 400,
		}
	}

	if vals.Challenge != "" {
		return events.APIGatewayProxyResponse{
			Body:       vals.Challenge,
			StatusCode: 200,
		}
	}
	switch request.Path {
	case "/api/twitch/follow":
		return handleFollowEvent(request)
	default:
		return events.APIGatewayProxyResponse{
			Body:       "",
			StatusCode: 404,
		}
	}
}

type eventSubNotification struct {
	Subscription helix.EventSubSubscription `json:"subscription"`
	Challenge    string                     `json:"challenge"`
	Event        json.RawMessage            `json:"event"`
}

func handleFollowEvent(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	var followEvent helix.EventSubChannelFollowEvent
	err := json.NewDecoder(bytes.NewReader([]byte(request.Body))).Decode(&followEvent)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "",
			StatusCode: 400,
		}
	}

	log.Println(followEvent)

	return events.APIGatewayProxyResponse{
		Body:       "",
		StatusCode: 200,
	}
}
