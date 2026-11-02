package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"hash/crc32"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
	"github.com/aws/smithy-go/logging"
)

const (
	region           = "eu-west-1"
	requestTimeout   = 15 * time.Second
	presignLifetime  = 5 * time.Minute
	objectKey        = "reports/day27/sample.txt"
	payloadText      = "day27 private transfer\n"
	contentType      = "text/plain"
	maxResponseBytes = 1024
)

var (
	errInvalidRequest = errors.New("invalid request")
	errResponseLarge  = errors.New("response too large")
)

// Drop SDK diagnostics, which could otherwise contain a URL or request data.
type quietLogger struct{}

func (quietLogger) Logf(_ logging.Classification, _ string, _ ...interface{}) {}

func main() {
	os.Exit(run())
}

func run() int {
	checkBucket := flag.Bool("check-bucket", false, "Read-only bucket existence check")
	flag.Parse()
	bucket := strings.TrimSpace(os.Getenv("CLOUD_ENG_DAY27_BUCKET"))
	if bucket == "" {
		return fail("MissingBucketName")
	}

	// AWS_PROFILE selects the existing temporary-credential source externally.
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(region),
		config.WithLogger(quietLogger{}),
	)
	cancel()
	if err != nil {
		return fail(errorCategory(err))
	}
	client := s3.NewFromConfig(cfg)
	if *checkBucket {
		return checkBucketExists(client, bucket)
	}

	httpClient := &http.Client{
		Timeout: requestTimeout,
		CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	return transfer(s3.NewPresignClient(client), httpClient, bucket)
}

func transfer(presigner *s3.PresignClient, client *http.Client, bucket string) int {
	payload := []byte(payloadText)

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	put, err := presigner.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket:            aws.String(bucket),
		Key:               aws.String(objectKey),
		Body:              bytes.NewReader(payload),
		ContentLength:     aws.Int64(int64(len(payload))),
		ContentType:       aws.String(contentType),
		ChecksumAlgorithm: types.ChecksumAlgorithmCrc32,
		ChecksumCRC32:     aws.String(payloadCRC32(payload)),
	}, func(options *s3.PresignOptions) {
		options.Expires = presignLifetime
	})
	cancel()
	if err != nil {
		return fail(errorCategory(err))
	}
	if put == nil || put.Method != http.MethodPut {
		return fail("UnexpectedPresignedRequest")
	}
	fmt.Println("presign expiry configured: 300 seconds")

	status, _, err := sendRequest(client, put.Method, put.URL, put.SignedHeader, payload)
	if status != 0 {
		fmt.Printf("upload HTTP status: %d\n", status)
	}
	if err != nil {
		return fail(errorCategory(err))
	}
	if status != http.StatusOK {
		return fail("UploadHTTPFailure")
	}
	fmt.Println("upload check: PASS")

	ctx, cancel = context.WithTimeout(context.Background(), requestTimeout)
	get, err := presigner.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objectKey),
	}, func(options *s3.PresignOptions) {
		options.Expires = presignLifetime
	})
	cancel()
	if err != nil {
		return fail(errorCategory(err))
	}
	if get == nil || get.Method != http.MethodGet {
		return fail("UnexpectedPresignedRequest")
	}

	status, downloaded, err := sendRequest(client, get.Method, get.URL, get.SignedHeader, nil)
	if status != 0 {
		fmt.Printf("download HTTP status: %d\n", status)
	}
	if err != nil {
		return fail(errorCategory(err))
	}
	if status != http.StatusOK {
		return fail("DownloadHTTPFailure")
	}
	matched := bytes.Equal(downloaded, payload)
	fmt.Printf("exact payload match: %t\n", matched)
	if !matched {
		return fail("PayloadMismatch")
	}
	fmt.Println("download check: PASS")

	// Copy only the URL; create a fresh unsigned request without signed headers.
	unsigned, err := url.Parse(get.URL)
	if err != nil {
		return fail("InvalidRequest")
	}
	unsigned.RawQuery = ""
	unsigned.ForceQuery = false
	unsigned.Fragment = ""
	status, _, err = sendRequest(client, http.MethodGet, unsigned.String(), nil, nil)
	if status != 0 {
		fmt.Printf("unsigned GET status: %d\n", status)
	}
	if err != nil {
		return fail(errorCategory(err))
	}
	if status != http.StatusForbidden {
		return fail("UnsignedRequestWasNotDenied")
	}
	fmt.Println("private access check: PASS")
	return 0
}

func sendRequest(client *http.Client, method, target string, headers http.Header, body []byte) (int, []byte, error) {
	parsed, err := url.Parse(target)
	if err != nil || parsed.Scheme != "https" || parsed.Host == "" || parsed.User != nil {
		return 0, nil, errInvalidRequest
	}
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, method, target, bytes.NewReader(body))
	if err != nil {
		return 0, nil, errInvalidRequest
	}
	if headers != nil {
		req.Header = headers.Clone()
		if host := req.Header.Get("Host"); host != "" {
			req.Host = host
			req.Header.Del("Host")
		}
	}
	if method == http.MethodPut && req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", contentType)
	}

	response, err := client.Do(req)
	if err != nil {
		if response != nil && response.Body != nil {
			response.Body.Close()
		}
		return 0, nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		// Error/redirect bodies may contain identifying details. Do not read/log them.
		return response.StatusCode, nil, nil
	}
	data, err := io.ReadAll(io.LimitReader(response.Body, maxResponseBytes+1))
	if err != nil {
		return response.StatusCode, nil, err
	}
	if len(data) > maxResponseBytes {
		return response.StatusCode, nil, errResponseLarge
	}
	return response.StatusCode, data, nil
}

func payloadCRC32(payload []byte) string {
	// Supply the exact body's checksum instead of relying on SDK defaults.
	var checksum [4]byte
	binary.BigEndian.PutUint32(checksum[:], crc32.ChecksumIEEE(payload))
	return base64.StdEncoding.EncodeToString(checksum[:])
}

func checkBucketExists(client *s3.Client, bucket string) int {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	_, err := client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket:  aws.String(bucket),
		MaxKeys: aws.Int32(1),
	})
	if err != nil {
		return fail(errorCategory(err))
	}
	fmt.Println("BucketPresent")
	return 0
}

func errorCategory(err error) string {
	if errors.Is(err, errInvalidRequest) {
		return "InvalidRequest"
	}
	if errors.Is(err, errResponseLarge) {
		return "ResponseTooLarge"
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return "Timeout"
	}
	if errors.Is(err, context.Canceled) {
		return "Canceled"
	}
	var httpErr *url.Error
	if errors.As(err, &httpErr) && httpErr.Timeout() {
		return "Timeout"
	}
	var apiErr smithy.APIError
	if errors.As(err, &apiErr) {
		switch apiErr.ErrorCode() {
		case "NoSuchBucket":
			return "BucketMissing"
		case "NoSuchKey":
			return "ObjectMissing"
		case "AccessDenied", "AccessDeniedException":
			return "AccessDenied"
		case "ExpiredToken", "ExpiredTokenException", "InvalidAccessKeyId",
			"InvalidClientTokenId", "UnrecognizedClientException",
			"InvalidSignatureException", "SignatureDoesNotMatch",
			"MissingAuthenticationTokenException":
			return "AuthenticationError"
		}
	}
	return "RequestFailed"
}

func fail(category string) int {
	fmt.Fprintln(os.Stderr, category)
	return 1
}
