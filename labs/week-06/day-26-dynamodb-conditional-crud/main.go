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
	jobSK          = "JOB#001"
)

var errUnexpectedResponse = errors.New("unexpected response")

type jobLab struct {
	client    *dynamodb.Client
	tableName string
}

func main() {
	os.Exit(run())
}

func run() int {
	tableName := flag.String("table", "cloud-eng-day26-jobs", "Disposable DynamoDB table")
	flag.Parse()
	if *tableName == "" {
		return fail("TableNameRequired")
	}

	// Select the existing temporary-credential profile with AWS_PROFILE externally.
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	cancel()
	if err != nil {
		return fail(errorCategory(err))
	}
	lab := jobLab{client: dynamodb.NewFromConfig(cfg), tableName: *tableName}

	if err := lab.putJob("queued"); err != nil {
		return fail(errorCategory(err))
	}
	item, err := lab.getJob()
	if err != nil {
		return fail(errorCategory(err))
	}
	if !matchesJob(item, "queued") {
		return fail("UnexpectedCreatedItem")
	}
	fmt.Println("create check: PASS")

	// Only the expected conditional failure counts as a successful guard.
	err = lab.putJob("unexpected")
	if err == nil {
		return fail("DuplicateWasNotBlocked")
	}
	if !isConditionalFailure(err) {
		return fail(errorCategory(err))
	}
	item, err = lab.getJob()
	if err != nil {
		return fail(errorCategory(err))
	}
	if !matchesJob(item, "queued") {
		return fail("DuplicateChangedItem")
	}
	fmt.Println("duplicate guard category: ConditionalCheckFailedException")
	fmt.Println("duplicate guard: PASS")

	item, err = lab.updateJob("queued", "running")
	if err != nil {
		return fail(errorCategory(err))
	}
	if !matchesJob(item, "running") {
		return fail("UnexpectedUpdatedItem")
	}
	fmt.Println("update check: PASS")

	_, err = lab.updateJob("queued", "complete")
	if err == nil {
		return fail("StaleUpdateWasNotBlocked")
	}
	if !isConditionalFailure(err) {
		return fail(errorCategory(err))
	}
	item, err = lab.getJob()
	if err != nil {
		return fail(errorCategory(err))
	}
	if !matchesJob(item, "running") {
		return fail("StaleUpdateChangedItem")
	}
	fmt.Println("stale guard category: ConditionalCheckFailedException")
	fmt.Println("stale guard: PASS")

	item, err = lab.deleteJob()
	if err != nil {
		return fail(errorCategory(err))
	}
	if !matchesJob(item, "running") {
		return fail("UnexpectedDeletedItem")
	}
	fmt.Println("delete check: PASS")

	item, err = lab.getJob()
	if err != nil {
		return fail(errorCategory(err))
	}
	if len(item) != 0 {
		return fail("ItemStillPresent")
	}
	fmt.Println("absent item check: PASS")
	return 0
}

func jobKey() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"PK": &types.AttributeValueMemberS{Value: userPK},
		"SK": &types.AttributeValueMemberS{Value: jobSK},
	}
}

func (lab jobLab) putJob(status string) error {
	item := jobKey()
	item["jobStatus"] = &types.AttributeValueMemberS{Value: status}
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	_, err := lab.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:           aws.String(lab.tableName),
		Item:                item,
		ConditionExpression: aws.String("attribute_not_exists(#pk) AND attribute_not_exists(#sk)"),
		ExpressionAttributeNames: map[string]string{
			"#pk": "PK",
			"#sk": "SK",
		},
	})
	return err
}

func (lab jobLab) getJob() (map[string]types.AttributeValue, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	out, err := lab.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName:      aws.String(lab.tableName),
		Key:            jobKey(),
		ConsistentRead: aws.Bool(true),
	})
	if err != nil {
		return nil, err
	}
	if out == nil {
		return nil, errUnexpectedResponse
	}
	return out.Item, nil
}

func (lab jobLab) updateJob(expected, next string) (map[string]types.AttributeValue, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	out, err := lab.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName:           aws.String(lab.tableName),
		Key:                 jobKey(),
		UpdateExpression:    aws.String("SET #status = :next"),
		ConditionExpression: aws.String("#status = :expected"),
		ExpressionAttributeNames: map[string]string{
			"#status": "jobStatus",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":expected": &types.AttributeValueMemberS{Value: expected},
			":next":     &types.AttributeValueMemberS{Value: next},
		},
		ReturnValues: types.ReturnValueAllNew,
	})
	if err != nil {
		return nil, err
	}
	if out == nil {
		return nil, errUnexpectedResponse
	}
	return out.Attributes, nil
}

func (lab jobLab) deleteJob() (map[string]types.AttributeValue, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	out, err := lab.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName:           aws.String(lab.tableName),
		Key:                 jobKey(),
		ConditionExpression: aws.String("#status = :expected"),
		ExpressionAttributeNames: map[string]string{
			"#status": "jobStatus",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":expected": &types.AttributeValueMemberS{Value: "running"},
		},
		ReturnValues: types.ReturnValueAllOld,
	})
	if err != nil {
		return nil, err
	}
	if out == nil {
		return nil, errUnexpectedResponse
	}
	return out.Attributes, nil
}

func matchesJob(item map[string]types.AttributeValue, status string) bool {
	return stringEquals(item["PK"], userPK) &&
		stringEquals(item["SK"], jobSK) &&
		stringEquals(item["jobStatus"], status)
}

func stringEquals(value types.AttributeValue, expected string) bool {
	text, ok := value.(*types.AttributeValueMemberS)
	return ok && text != nil && text.Value == expected
}

func isConditionalFailure(err error) bool {
	var conditionErr *types.ConditionalCheckFailedException
	if errors.As(err, &conditionErr) {
		return true
	}
	var apiErr smithy.APIError
	return errors.As(err, &apiErr) && apiErr.ErrorCode() == "ConditionalCheckFailedException"
}

func errorCategory(err error) string {
	if errors.Is(err, errUnexpectedResponse) {
		return "UnexpectedResponse"
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return "Timeout"
	}
	if errors.Is(err, context.Canceled) {
		return "Canceled"
	}
	var apiErr smithy.APIError
	if errors.As(err, &apiErr) {
		switch apiErr.ErrorCode() {
		case "ConditionalCheckFailedException":
			return "ConditionalCheckFailedException"
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
