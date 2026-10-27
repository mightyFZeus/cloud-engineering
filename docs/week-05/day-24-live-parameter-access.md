# Day 24 — Test live least-privilege Parameter Store access

## References

- [IAM roles in the AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-role.html)
- [Parameter Store permissions](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-setting-up.html)

Passed October 27, 2026, ahead of the October 28 due date: all six criteria supported by learner-reported live results, code/template inspection, explanations, and cleanup evidence. The coach did not execute authenticated AWS requests or inspect private AWS configuration.

## Setup and credential flow

- eu-west-1 and budget checked. Two Standard String parameters contained the benign value learning.
- Source aws-learning profile authenticated with aws login. Real identity/resource ARNs remained in the AWS console, clipboard, and private AWS configuration outside the repository.
- Temporary role cloud-eng-day24-reader trusted the source IAM user. No managed policies attached.
- One inline Allow: ssm:GetParameter on the exact allowed parameter resource. The saved policy-template.json uses an account placeholder and is not a literal deployment document.
- The day24-reader CLI profile used source_profile=aws-learning, session name cloud-eng-day24, and requested duration 900 seconds. The role identity suffix comparison returned True with exit 0.

## Observed live evidence

| Caller profile | Parameter | Actual result | Exit |
|---|---|---|---|
| aws-learning | /cloud-eng/day24/allowed | Value matched learning | 0 |
| aws-learning | /cloud-eng/day24/denied | Value matched learning | 0 |
| day24-reader | /cloud-eng/day24/allowed | Value matched learning | 0 |
| day24-reader | /cloud-eng/day24/denied | AccessDeniedException | 254 |

The denied parameter existed and was readable by the source before the role test. The result therefore demonstrated a real authorization failure for that caller/action/resource/context, rather than a missing-name test.

The learner explained that trust controls assumption, permissions control role actions, and source administrative permissions are not automatically added to an assumed role. AccessDenied and ParameterNotFound were correctly distinguished. The claim covers GetParameter on two tested resources; no live S3 prefix authorization or complete IAM security audit is claimed.

## Cleanup and connectivity limitation

The learner deleted both parameters through the console. CLI absence checks repeatedly returned EndpointConnectionError with exit 255; a filtered diagnostic identified the Parameter Store endpoint. An unauthenticated curl request returned HTTP 400 with exit 0, establishing an HTTP response through curl rather than authenticated API success.

The existing Day 20 Go reader was inspected and rebuilt. With AWS_PROFILE=aws-learning and credential environment overrides cleared, its GetParameter calls in eu-west-1 returned typed ParameterNotFound and exit 1 for both Day 24 names. This supports parameter absence. It does not establish that the CLI connectivity issue was repaired; its underlying cause remains unknown.

The role was deleted and confirmed absent after refreshing IAM. Console role deletion also removes associated inline policies. The exact temporary profile block was removed locally, and source logout exited 0. No global cache deletion was requested; source logout is not proof that every cached STS role session was revoked.

No long-term access keys or additional AWS resources were created, as confirmed by the learner. The redacted JSON and README checks reported exit 0. Local review found the expected exact template and substantive safe evidence.

## Safe command summary and artifacts

Used profile login/authentication checks, output-filtered source/role reads, private clipboard policy preparation, local profile append/removal, console resource deletion, a credential-free HTTP probe, the existing Go reader for absence checks, source logout, and local JSON/file checks. Raw responses, errors, trust documents, private configuration, account IDs, and actual ARNs were not retained in the repository.

- [Saved evidence](../../labs/week-05/day-24-live-parameter-access/README.md)
- [Redacted permission template](../../labs/week-05/day-24-live-parameter-access/policy-template.json)
- [Reviewed Go reader](../../labs/week-05/day-20-go-parameter-store/main.go)
- [Social drafts](../../posts/week-05/day-24-live-parameter-access.md)
