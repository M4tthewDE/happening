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

func GetUser(request events.APIGatewayProxyRequest) (string, int) {
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

	token, err := d.GetAuth(ctx)
	if err != nil {
		log.Println(err)
		return "", 500
	}

	client, err := helix.NewClient(&helix.Options{
		ClientID:       os.Getenv("TWITCH_CLIENT_ID"),
		AppAccessToken: token,
	})
	if err != nil {
		log.Println(err)
		return "", 500
	}

	name := request.QueryStringParameters["name"]

	resp, err := client.GetUsers(&helix.UsersParams{Logins: []string{name}})
	if err != nil {
		log.Println(err)
		return "", 500
	}

	if len(resp.Data.Users) == 0 {
		return "", 404
	}

	bodyBytes, err := json.Marshal(resp.Data.Users[0])
	if err != nil {
		log.Println(err)
		return "", 500
	}

	return string(bodyBytes), 200
}
