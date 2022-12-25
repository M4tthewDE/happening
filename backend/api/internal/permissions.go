package internal

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/nicklaw5/helix"
)

func GetPermissions(request events.APIGatewayProxyRequest) (string, int) {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx, func(o *config.LoadOptions) error {
		o.Region = "us-east-1"
		return nil
	})
	if err != nil {
		log.Println(err)
		return "", 500
	}

	dbClient := dynamodb.NewFromConfig(cfg)
	d := NewDao(dbClient)

	token, ok := request.QueryStringParameters["token"]
	if !ok {
		return "", 400
	}

	client, err := helix.NewClient(&helix.Options{
		ClientID:       os.Getenv("TWITCH_CLIENT_ID"),
		AppAccessToken: token,
	})
	if err != nil {
		log.Println(err)
		return "", 500
	}

	isValid, resp, err := client.ValidateToken(token)
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

	permissions, err := d.GetPermissions(ctx, resp.Data.UserID)
	if err != nil {
		log.Println(err)
		return "", 500
	}

	bodyBytes, err := json.Marshal(permissions)
	if err != nil {
		log.Println(err)
		return "", 500
	}

	return string(bodyBytes), 200
}
