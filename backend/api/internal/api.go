package internal

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
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

func (d DB) GetAuth(ctx context.Context) (string, error) {
	out, err := d.ddb.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String(table_name),
	})
	if err != nil {
		return "", err
	}

	if len(out.Items) == 0 {
		return "", nil
	}

	token := out.Items[0]["access_token"].(*types.AttributeValueMemberS).Value
	return token, nil
}

func HandleNewSubscription(body string) (response string, status int) {
	var newSubscriptionBody NewSubscriptionBody
	err := json.Unmarshal([]byte(body), &newSubscriptionBody)
	if err != nil {
		return "", 400
	}

	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx, func(o *config.LoadOptions) error {
		o.Region = "us-east-1"
		return nil
	})
	if err != nil {
		return "", 500
	}

	client := dynamodb.NewFromConfig(cfg)
	d := NewDao(client)

	token, err := d.GetAuth(ctx)
	if err != nil {
		return "", 500
	}

	resp, err := json.Marshal(newSubscriptionBody)
	if err != nil {
		return "", 400
	}

	return string(resp), 200
}

type NewSubscriptionBody struct {
	TargetUserID string `json:"target_id"`
	SubType      string `json:"subscription_type"`
}
