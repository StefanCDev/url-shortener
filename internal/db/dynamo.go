package db

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// URL struct - represents a short URL entry in DynamoDB
type URL struct {
	ShortCode string `json:"short_code"` // Shortened URL ID
	LongURL   string `json:"long_url"`   // Original full URL
}

// DynamoDB client instance
var dbClient *dynamodb.DynamoDB

// Initialize DynamoDB client
func InitDB() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")), // Use environment variables
	}))
	dbClient = dynamodb.New(sess)
}

// Save a new short URL to DynamoDB
func SaveURL(shortCode, longURL string) error {
	item := URL{
		ShortCode: shortCode,
		LongURL:   longURL,
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return fmt.Errorf("failed to marshal item: %v", err)
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(os.Getenv("DYNAMODB_TABLE")), // Table name from env
		Item:      av,
	}

	_, err = dbClient.PutItem(input)
	if err != nil {
		return fmt.Errorf("failed to save URL: %v", err)
	}

	return nil
}

// Retrieve original URL from short code
func GetURL(shortCode string) (string, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("DYNAMODB_TABLE")),
		Key: map[string]*dynamodb.AttributeValue{
			"short_code": {S: aws.String(shortCode)},
		},
	}

	result, err := dbClient.GetItem(input)
	if err != nil {
		return "", fmt.Errorf("failed to get URL: %v", err)
	}

	if result.Item == nil {
		return "", fmt.Errorf("short URL not found")
	}

	var url URL
	err = dynamodbattribute.UnmarshalMap(result.Item, &url)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal result: %v", err)
	}

	return url.LongURL, nil
}
