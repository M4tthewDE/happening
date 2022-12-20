package internal

import (
	"context"

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
