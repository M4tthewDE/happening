package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/nicklaw5/helix/v2"
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

func (d DB) DeleteAuth(ctx context.Context, id string) error {
	_, err := d.ddb.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(table_name),
		Key: map[string]types.AttributeValue{
			hash_key: &types.AttributeValueMemberS{Value: id},
		},
	})

	return err
}

func (d DB) GetAuth(ctx context.Context) (string, string, bool, error) {
	out, err := d.ddb.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String(table_name),
	})
	if err != nil {
		return "", "", false, err
	}

	if len(out.Items) == 0 {
		return "", "", false, nil
	}

	token := out.Items[0]["access_token"].(*types.AttributeValueMemberS).Value
	id := out.Items[0]["id"].(*types.AttributeValueMemberS).Value
	return token, id, true, nil
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

func GenerateToken() (*helix.AccessCredentials, error) {
	client, err := helix.NewClient(&helix.Options{
		ClientID:     os.Getenv("TWITCH_CLIENT_ID"),
		ClientSecret: os.Getenv("TWITCH_SECRET"),
	})
	if err != nil {
		return nil, nil
	}

	resp, err := client.RequestAppAccessToken([]string{""})
	if err != nil {
		return nil, nil
	}

	return &resp.Data, nil
}

func ShouldRefresh(token string) (bool, error) {
	client, err := helix.NewClient(&helix.Options{
		ClientID: os.Getenv("TWITCH_CLIENT_ID"),
	})
	if err != nil {
		return false, err
	}

	_, resp, err := client.ValidateToken(token)
	if err != nil {
		return false, err
	}

	return resp.Data.ExpiresIn < 660, nil
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

	token, id, exists, err := d.GetAuth(ctx)
	if err != nil {
		return err
	}

	if !exists {
		token, err := GenerateToken()
		if err != nil {
			return err
		}

		err = d.SaveAuth(ctx, token.AccessToken)
		if err != nil {
			return err
		}

		return nil
	}

	shouldRefresh, err := ShouldRefresh(token)
	if err != nil {
		return err
	}

	if shouldRefresh {
		token, err := GenerateToken()
		if err != nil {
			return err
		}

		log.Println("Refreshing token...")

		err = d.DeleteAuth(ctx, id)
		if err != nil {
			return err
		}

		err = d.SaveAuth(ctx, token.AccessToken)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
