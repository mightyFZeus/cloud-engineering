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
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go"
)

func main() {
	os.Exit(run())
}

func run() int {
	name := flag.String("name", "/cloud-eng/day20/environment", "Parameter name")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("eu-west-1"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "ConfigurationError")
		return 1
	}

	client := ssm.NewFromConfig(cfg)
	result, err := client.GetParameter(ctx, &ssm.GetParameterInput{
		Name:           aws.String(*name),
		WithDecryption: aws.Bool(false),
	})
	if err != nil {
		category := "RequestFailed"
		var missing *types.ParameterNotFound
		var apiErr smithy.APIError

		switch {
		case errors.As(err, &missing):
			category = "ParameterNotFound"
		case errors.Is(err, context.DeadlineExceeded):
			category = "Timeout"
		case errors.As(err, &apiErr):
			switch apiErr.ErrorCode() {
			case "AccessDenied", "AccessDeniedException":
				category = "AccessDenied"
			case "ExpiredToken", "ExpiredTokenException",
				"InvalidClientTokenId", "UnrecognizedClientException",
				"InvalidSignatureException", "SignatureDoesNotMatch":
				category = "AuthenticationError"
			}
		}

		fmt.Fprintln(os.Stderr, category)
		return 1
	}

	if result.Parameter == nil ||
		result.Parameter.Type != types.ParameterTypeString ||
		aws.ToString(result.Parameter.Value) != "learning" {
		fmt.Fprintln(os.Stderr, "UnexpectedParameter")
		return 1
	}

	fmt.Println("parameter check: PASS")
	return 0
}
