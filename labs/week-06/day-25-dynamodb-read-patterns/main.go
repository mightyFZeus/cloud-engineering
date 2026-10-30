package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go"
)

const (
	region         = "eu-west-1"
	requestTimeout = 15 * time.Second
	userPK         = "USER#demo-a"
)

func main() {
	os.Exit(run())
}

func run() int {
	tableName := flag.String("table", "cloud-eng-day25-jobs", "DynamoDB table to read")
	flag.Parse()
	if *tableName == "" {
		return fail("TableNameRequired")
	}

	// AWS_PROFILE selects the existing temporary-credential profile externally.
	configCtx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	cfg, err := config.LoadDefaultConfig(configCtx, config.WithRegion(region))
	cancel()
	if err != nil {
		return fail(errorCategory(err))
	}
	client := dynamodb.NewFromConfig(cfg)

	itemCtx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	item, err := client.GetItem(itemCtx, &dynamodb.GetItemInput{
		TableName: aws.String(*tableName),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: userPK},
			"SK": &types.AttributeValueMemberS{Value: "JOB#001"},
		},
		ConsistentRead: aws.Bool(true),
	})
	cancel()
	if err != nil {
		return fail(errorCategory(err))
	}
	if item == nil || !matchesJob(item.Item, "JOB#001", "queued") {
		return fail("UnexpectedItem")
	}
	fmt.Println("item check: PASS")

	queryCtx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	jobs, err := client.Query(queryCtx, &dynamodb.QueryInput{
		TableName:              aws.String(*tableName),
		KeyConditionExpression: aws.String("#pk = :user AND begins_with(#sk, :prefix)"),
		ExpressionAttributeNames: map[string]string{
			"#pk": "PK",
			"#sk": "SK",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":user":   &types.AttributeValueMemberS{Value: userPK},
			":prefix": &types.AttributeValueMemberS{Value: "JOB#"},
		},
		ConsistentRead:   aws.Bool(true),
		ScanIndexForward: aws.Bool(true),
	})
	cancel()
	if err != nil {
		return fail(errorCategory(err))
	}
	if jobs == nil {
		return fail("UnexpectedQuery")
	}
	// This lab expects one tiny page; it does not implement general pagination.
	if len(jobs.LastEvaluatedKey) != 0 {
		return fail("IncompletePage")
	}
	if len(jobs.Items) != 2 ||
		!matchesJob(jobs.Items[0], "JOB#001", "queued") ||
		!matchesJob(jobs.Items[1], "JOB#002", "complete") {
		return fail("UnexpectedQuery")
	}
	fmt.Println("query check: PASS")
	return 0
}

func matchesJob(item map[string]types.AttributeValue, sortKey, jobStatus string) bool {
	return stringEquals(item["PK"], userPK) &&
		stringEquals(item["SK"], sortKey) &&
		stringEquals(item["jobStatus"], jobStatus)
}

func stringEquals(value types.AttributeValue, expected string) bool {
	text, ok := value.(*types.AttributeValueMemberS)
	return ok && text != nil && text.Value == expected
}

func errorCategory(err error) string {
	if errors.Is(err, context.DeadlineExceeded) {
		return "Timeout"
	}
	if errors.Is(err, context.Canceled) {
		return "Canceled"
	}
	var apiErr smithy.APIError
	if errors.As(err, &apiErr) {
		switch apiErr.ErrorCode() {
		case "ResourceNotFoundException":
			return "TableMissing"
		case "AccessDenied", "AccessDeniedException":
			return "AccessDenied"
		case "ExpiredToken", "ExpiredTokenException", "InvalidClientTokenId",
			"UnrecognizedClientException", "InvalidSignatureException",
			"SignatureDoesNotMatch", "MissingAuthenticationTokenException":
			return "AuthenticationError"
		}
	}
	return "RequestFailed"
}

func fail(category string) int {
	fmt.Fprintln(os.Stderr, category)
	return 1
}
