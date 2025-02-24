package main

import (
	"context"

	"url-shortener/internal/db"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler function for AWS Lambda
func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	shortCode := request.PathParameters["short_code"]

	// Fetch original URL
	longURL, err := db.GetURL(shortCode)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 404, Body: "Short URL not found"}, nil
	}

	// Redirect user
	headers := map[string]string{"Location": longURL}
	return events.APIGatewayProxyResponse{StatusCode: 302, Headers: headers}, nil
}

// Lambda main entry point
func main() {
	db.InitDB() // Initialize DynamoDB connection
	lambda.Start(handler)
}
