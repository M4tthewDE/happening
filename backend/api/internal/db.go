package internal

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
)

type DB struct {
	ddb *dynamodb.Client
}

func NewDao(ddb *dynamodb.Client) *DB {
	return &DB{ddb: ddb}
}

func (d DB) GetAuth(ctx context.Context) (string, error) {
	out, err := d.ddb.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
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

func (d DB) GetPermissions(ctx context.Context, user_id string) ([]string, error) {
	out, err := d.ddb.Scan(ctx, &dynamodb.ScanInput{
		TableName:        aws.String(os.Getenv("PERMISSIONS_TABLE_NAME")),
		FilterExpression: aws.String("user_id = :uid"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":uid": &types.AttributeValueMemberS{Value: user_id},
		},
	})
	if err != nil {
		return nil, err
	}

	if len(out.Items) == 0 {
		return make([]string, 0), nil
	}

	permissions := out.Items[0]["permissions"].(*types.AttributeValueMemberSS).Value
	return permissions, nil
}
