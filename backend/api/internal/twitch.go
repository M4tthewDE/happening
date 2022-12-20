package internal

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/nicklaw5/helix/v2"
)

func HandleNewEventsubEvent(request events.APIGatewayProxyRequest) (string, int) {
	if !VerifyEventSubNotification(os.Getenv("EVENTSUB_SECRET"), request.Headers, request.Body) {
		log.Println("Failed to verify signature")
		return "", 403
	}

	var vals eventSubNotification
	err := json.NewDecoder(bytes.NewReader([]byte(request.Body))).Decode(&vals)
	if err != nil {
		log.Println(err)
		return "", 400
	}

	if vals.Challenge != "" {
		return vals.Challenge, 400
	}

	// TODO: hit helix and check if we want to handle this event
	// compare IDs

	switch vals.Subscription.Type {
	case "channel.follow":
		return handleFollowEvent(request)
	default:
		return "", 404
	}
}

func VerifyEventSubNotification(secret string, headers map[string]string, message string) bool {
	hmacMessage := []byte(fmt.Sprintf("%s%s%s", headers["twitch-eventsub-message-id"], headers["twitch-eventsub-message-timestamp"], message))
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(hmacMessage)
	hmacsha256 := fmt.Sprintf("sha256=%s", hex.EncodeToString(mac.Sum(nil)))
	return hmacsha256 == headers["twitch-eventsub-message-signature"]
}

type eventSubNotification struct {
	Subscription helix.EventSubSubscription `json:"subscription"`
	Challenge    string                     `json:"challenge"`
	Event        json.RawMessage            `json:"event"`
}

func handleFollowEvent(request events.APIGatewayProxyRequest) (string, int) {
	var followEvent helix.EventSubChannelFollowEvent
	err := json.NewDecoder(bytes.NewReader([]byte(request.Body))).Decode(&followEvent)
	if err != nil {
		return "", 400
	}

	// TODO: empty struct
	log.Println(followEvent)

	return "", 200
}
