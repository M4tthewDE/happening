package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/nicklaw5/helix"
)

const (
	table_name = "auth"
	hash_key   = "id"
)

type DB struct {
	ddb *dynamodb.Client
}

func NewDao(ddb *dynamodb.Client) *DB {
	return &DB{ddb: ddb}
}

func (d DB) GetAuth(ctx context.Context) (string, bool, error) {
	out, err := d.ddb.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String(table_name),
	})
	if err != nil {
		return "", false, err
	}

	if len(out.Items) == 0 {
		return "", false, nil
	}

	return "", len(out.Items) != 0, nil
}

func (d DB) SaveAuth(ctx context.Context, token string) error {
	d.ddb.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(table_name),
		Item: map[string]types.AttributeValue{
			hash_key:       &types.AttributeValueMemberS{Value: uuid.New().String()},
			"access_token": &types.AttributeValueMemberS{Value: token},
		},
	})

	return nil
}

func HandleRequest(ctx context.Context, event events.CloudWatchEvent) error {
	cfg, err := config.LoadDefaultConfig(ctx, func(o *config.LoadOptions) error {
		o.Region = "us-east-1"
		return nil
	})
	if err != nil {
		return err
	}

	client := dynamodb.NewFromConfig(cfg)
	d := NewDao(client)

	_, exists, err := d.GetAuth(ctx)
	if err != nil {
		return err
	}

	if !exists {
		token, err := GenerateToken()
		if err != nil {
			return err
		}

		err = d.SaveAuth(ctx, token)
		if err != nil {
			return err
		}
	}

	return nil
}

func GenerateToken() (string, error) {
	client, err := helix.NewClient(&helix.Options{
		ClientID:     os.Getenv("TWITCH_CLIENT_ID"),
		ClientSecret: os.Getenv("TWITCH_SECRET"),
	})
	if err != nil {
		return "", nil
	}

	resp, err := client.RequestAppAccessToken([]string{""})
	if err != nil {
		return "", nil
	}

	return resp.Data.AccessToken, nil
}

func main() {
	lambda.Start(HandleRequest)
}
