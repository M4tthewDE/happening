package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Processing Lambda request %t\n", request.RequestContext)
	return events.APIGatewayProxyResponse{
		Body:       "Hello world",
		StatusCode: 200,
	}, nil
}

func main() {
	log.Printf("Start lambda")
	lambda.Start(HandleRequest)
}
