package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, event events.CloudWatchEvent) error {
	log.Println(event)
	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
