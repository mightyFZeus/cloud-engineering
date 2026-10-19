# Day 20 — Read a Non-Secret Parameter with the Go AWS SDK

**Date:** October 19, 2026
**Due:** 11:00 PM Africa/Lagos
**Timebox:** 60–90 minutes
**Status:** Passed October 19, 2026 — all six criteria verified.

## Outcome

Use Go AWS SDK v2 and your existing temporary-credential profile to retrieve one non-secret Parameter Store setting, handle a missing setting, and delete the temporary parameter. Run on your Mac; no EC2 required.

## Steps

1. Check the budget and eu-west-1 selection. Use your existing MFA/SSO temporary-credential profile. Refresh its normal session if necessary; create no long-term access keys. SSO needs a valid SSO login; an assume-role profile with mfa_serial may require an interactive SDK TokenProvider, which we will configure during the walkthrough. A prior AWS CLI role-cache entry alone is not proof the SDK can authenticate. Do not paste credential files or identity output into chat.
2. Create labs/week-05/day-20-go-parameter-store/ with README.md and main.go. From this new directory initialize a dedicated module, only if no go.mod exists: go mod init example.com/cloud-eng/day20. Add github.com/aws/aws-sdk-go-v2/config and github.com/aws/aws-sdk-go-v2/service/ssm with go get; retain go.mod and go.sum.
3. In Systems Manager -> Parameter Store, create a new Standard-tier String named /cloud-eng/day20/environment with value learning. This is deliberately non-secret. Leave advanced features and higher-throughput settings unchanged. If that exact name already exists, stop and report it rather than overwriting.
4. Write a small Go program accepting a -name flag for the parameter name (default /cloud-eng/day20/environment). Use a context with a 15-second timeout, config.LoadDefaultConfig with config.WithRegion("eu-west-1"), ssm.NewFromConfig, and GetParameter with WithDecryption false. Do not use static credential providers or hard-code credentials. Select your existing profile through AWS_PROFILE when running; a web-console login alone does not authenticate the local SDK.
5. Check errors before dereferencing the response. Confirm a successful response contains value learning and type String; print only a concise parameter check: PASS. For errors, print a sanitized category and exit nonzero. Use errors.As with *ssm/types.ParameterNotFound to distinguish a missing name from access/authentication failures. Do not print raw SDK errors, response objects, ARNs, or credentials.
6. Run gofmt, go build, then the program with the intended profile; record build and run exit status and the safe PASS result.
7. Run against a deliberately absent name /cloud-eng/day20/does-not-exist. Confirm ParameterNotFound and a nonzero program exit. If it exists, choose a different unused lab name. Do not equate AccessDenied or expired credentials with a missing-parameter pass.
8. Delete only /cloud-eng/day20/environment after the successful read. Run the program with its default name again; confirm ParameterNotFound and nonzero exit. Record cleanup and log out of the source session using the established workflow.
9. In README.md explain the default credential chain, why a non-secret String differs from SecureString, and how authentication differs from authorization. Record summarized evidence and cleanup only.

## Authoritative references

- [AWS SDK for Go v2 configuration and credential providers](https://docs.aws.amazon.com/sdk-for-go/v2/developer-guide/configure-gosdk.html)
- [Working with Parameter Store](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-working-with.html)

## Cost and security

Use one short-lived Standard String parameter with ordinary throughput. Check current account settings and pricing; do not enable Advanced tier, higher throughput, KMS features, or Secrets Manager for this lab. Create no compute resources. Existing identity needs ssm:PutParameter, ssm:GetParameter, and ssm:DeleteParameter for the lab parameter; the absent-name test also needs GetParameter permission for that test name. if access is denied, report the sanitized blocker rather than granting broad permissions. The value learning is not a secret; never substitute a password or token.

## Submission

1. Lab directory and Go files:
2. Region; budget checked; temporary-profile session ready:
3. Parameter tier and type:
4. gofmt completed; build exit status:
5. Successful read result; run exit status:
6. Missing-name error category and nonzero exit:
7. Parameter deleted; post-delete error category and nonzero exit:
8. Default credential chain, String/SecureString, authentication/authorization explanation:
9. Long-term keys, EC2, or other paid resources created:
10. Source session logged out:
11. Blocker:

## Pass criteria

- Saved Go code uses AWS SDK v2 default credential chain, explicit Region, and bounded timeout.
- Correct temporary-profile session successfully reads the Standard String without exposing sensitive output.
- Missing parameter yields ParameterNotFound and nonzero exit, distinguished from other failures.
- Formatting and build succeed; module files saved.
- Explanations are accurate and no hard-coded credentials appear.
- Temporary parameter deletion verified by a failed read; no compute, long-term keys, or unneeded resources created; source session logged out.
