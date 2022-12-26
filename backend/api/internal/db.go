package internal

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
)

func getDdb(ctx context.Context) (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx, func(o *config.LoadOptions) error {
		o.Region = "us-east-1"
		return nil
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return dynamodb.NewFromConfig(cfg), nil

}

func GetAuth(ctx context.Context) (string, error) {
	ddb, err := getDdb(ctx)
	if err != nil {
		return "", err
	}

	out, err := ddb.Scan(ctx, &dynamodb.ScanInput{
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

func GetPermissions(ctx context.Context, user_id string) (bool, error) {
	ddb, err := getDdb(ctx)
	if err != nil {
		return false, err
	}

	out, err := ddb.Scan(ctx, &dynamodb.ScanInput{
		TableName:        aws.String(os.Getenv("PERMISSIONS_TABLE_NAME")),
		FilterExpression: aws.String("user_id = :uid"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":uid": &types.AttributeValueMemberS{Value: user_id},
		},
	})
	if err != nil {
		return false, err
	}

	if len(out.Items) == 0 {
		return false, nil
	}

	return true, nil
}
