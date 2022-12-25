package internal

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
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

	id, ok := request.QueryStringParameters["id"]
	if !ok {
		return "", 400
	}

	permissions, err := d.GetPermissions(ctx, id)
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
