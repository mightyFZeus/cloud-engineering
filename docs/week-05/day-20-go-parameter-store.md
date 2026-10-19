# Day 20 — Go Parameter Store reader

Passed October 19, 2026: six assignment criteria verified from learner checkpoints and review of the saved source, module files, and README.

## Verified evidence

- Region eu-west-1 and budget check confirmed. Existing aws-learning profile authenticated with aws login; authentication check exited 0.
- One Standard String parameter /cloud-eng/day20/environment created with non-secret value learning.
- AWS SDK for Go v2 config.LoadDefaultConfig, explicit Region, 15-second context timeout, and ssm.GetParameter used without static credentials. Environment credential overrides cleared in a subshell for Go runs.
- gofmt completed; build exited 0. Nonempty main.go, go.mod, go.sum, and README reviewed.
- Successful read printed parameter check: PASS with exit 0.
- Deliberately absent name returned typed ParameterNotFound with exit 1.
- DeleteParameter exited 0. A post-delete Go read returned ParameterNotFound with exit 1.
- README check and source-session logout exited 0. No long-term keys or EC2 instances created.

## Explanation and limits

Temporary credentials establish the authenticated identity; IAM policy evaluation governs access. The non-secret value learning can use String. SecureString with KMS encryption was explained for secret handling but was not deployed or tested. This lab exercised successful live SSM access, not a live test of Day 19's S3 policy.

A sign-in endpoint connectivity error occurred during cleanup. A diagnostic request later returned HTTP 405 and the delete retry succeeded. The 405 established endpoint reachability, not token validity; the initial connectivity cause was not established.

## Safe commands and artifacts

Command summary: aws login; authentication check with identity output discarded; go mod init/go get/go mod tidy; gofmt/go build; profile-selected local Go runs; scoped aws ssm delete-parameter; aws logout. No raw AWS response, credential cache, real ARN, account ID, or login URL is saved.

[Lab source and evidence](../../labs/week-05/day-20-go-parameter-store/README.md)

## References

- [Go SDK configuration and credentials](https://docs.aws.amazon.com/sdk-for-go/v2/developer-guide/configure-gosdk.html)
- [Working with Parameter Store](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-working-with.html)
