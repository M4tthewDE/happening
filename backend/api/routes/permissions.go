package routes

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/m4tthewde/happening/backend/api/internal"
)

func GetPermissions(request events.APIGatewayProxyRequest) (string, int) {
	token, ok := request.QueryStringParameters["token"]
	if !ok {
		return "", 400
	}

	isValid, resp, err := internal.ValidateToken(token)
	if err != nil {
		log.Println(err)
		return "", 500
	}

	if !isValid {
		return "", 403
	}

	if resp.Data.ClientID != os.Getenv("TWITCH_CLIENT_ID") {
		return "", 403
	}

	isAllowed, err := internal.GetPermissions(context.TODO(), resp.Data.UserID)
	if err != nil {
		log.Println(err)
		return "", 500
	}

	if isAllowed {
		return "", 200
	}

	return "", 403
}
