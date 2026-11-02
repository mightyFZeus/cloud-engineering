# Day 27 — Private S3 Transfers with Presigned URLs

## References

- [Presigned URL permissions and lifetime](https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-presigned-url.html)
- [Go SDK S3 examples](https://docs.aws.amazon.com/sdk-for-go/v2/developer-guide/go_s3_code_examples.html)
- [AWS CLI sign-out behavior](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sign-in.html)

## Planned flow

Temporary source credentials -> presign an object operation
-> ordinary HTTPS client sends the signed request -> S3.

## Planned checks

- Presigned upload succeeds.
- Presigned download returns the exact payload bytes.
- An unsigned request to the existing object is denied.
- Object and bucket are deleted afterward.

## URL handling

Presigned URLs stay in program memory.
Save only safe statuses, check labels, and exit statuses.


## Verified evidence — November 2, 2026

Guided lab. The learner ran the AWS console and live program checks; the coach inspected the saved source, payload, and module files. The coach did not make authenticated AWS requests or read private AWS configuration.

The actual bucket name is intentionally omitted. Resource notes use `<PRIVATE_LAB_BUCKET>`.

| Check | Observed result |
|---|---|
| Region and budget | eu-west-1; budget checked that day |
| Temporary source profile | aws-learning; remote-flow login and identity check exit 0 |
| New bucket | General purpose, shared global namespace, initially empty |
| Public access | All four Block Public Access settings enabled |
| Ownership and ACLs | Bucket owner enforced; ACLs disabled |
| Encryption | SSE-S3 |
| Versioning and Object Lock | Both disabled |
| Additional policies/features | None reported |
| Saved payload | Exact fictional bytes in payload.txt, including final newline |
| Module setup and S3 installation | Exit 0 each |
| Formatting, module tidy, build | Exit 0 each |
| Presign expiry configuration | 300 seconds |
| Signed PUT | HTTP 200; upload check: PASS |
| Signed GET | HTTP 200; exact payload match: true; download check: PASS |
| Fresh unsigned GET | HTTP 403; private access check: PASS |
| Overall transfer run | Exit 0 |
| Console observation | reports/day27/sample.txt reported visible after success |
| Console cleanup | Object deleted; bucket confirmed empty, then deleted and absent after refresh |
| Read-only SDK cleanup | BucketMissing; exit 1 |
| Bucket environment variable | Unset, confirmed without printing its value |
| Source logout | Exit 0 |
| Other resources or policy changes | None reported |
| Saved README presence check | Exit 0; substantive contents also inspected |

## Request and privacy behavior

- The program reads the actual bucket only from CLOUD_ENG_DAY27_BUCKET. It keeps presigned URLs and signing parameters in memory.
- PresignPutObject and PresignGetObject each configure a five-minute lifetime. The upload binds the expected body, length, content type, and CRC32; the HTTP request preserves the presigner's method and signed headers.
- The signed download checks the complete bounded response body against the exact fictional bytes. Matching status or length alone would not pass.
- The unsigned check creates a fresh request for the same verified-existing object with query parameters removed and no signing or Authorization headers. Its 403 did not require changing the private bucket settings.
- Calls and the HTTP client have bounded 15-second timeouts. Redirects are not followed. Response bodies are closed.
- SDK diagnostics are suppressed. Fixed error categories prevent raw HTTP errors, which can embed a signed URL, from reaching output.
- The read-only check-bucket branch calls ListObjectsV2 with MaxKeys=1. Only the exact NoSuchBucket service category becomes BucketMissing; an arbitrary 404 or AccessDenied is not deletion proof.

## Learner explanations and coach clarification

The learner explained that signed requests carry authorization while the unsigned request lacked permission; URL possession permits the authorized operation while valid; and credential expiry/revocation or resource/policy changes can prevent access before the configured five minutes.

The learner also correctly stated that expiration was not tested by waiting and attempting a later request. The bucket remained private.

Coach clarification: successful CLI logout alone is not proof that all previously generated URLs were revoked. Local sign-out and an actual revocation test are different evidence. Object deletion or policy denial can also make a request fail without demonstrating that its configured expiration time elapsed.

## Verification scope

The actual source matches the reviewed template. An offline SDK compatibility build succeeded. Separate fake-transport tests checked 300-second presigns, body/checksum and signed-header handling, fresh unsigned requests, content-mismatch rejection, safe errors, body closure, and exact cleanup categories. Those checks used dummy test inputs, not authenticated AWS requests.

Live results above were reported by the learner. Actual expiry, IAM prefix isolation, a production authenticated URL issuer, and independent backup/recovery were not demonstrated.

## Safe command summary

All relative paths were resolved from /Users/mightyzeus/Desktop/cloud-eng.

- Prepared the fictional payload and separate module; added the S3 service dependency.
- Copied the reviewed source, formatted it, tidied dependencies, and built the executable outside the repository.
- Renewed temporary source login and suppressed identity output.
- Generated the new learning bucket name privately and copied it to the console without printing or storing it in repository files.
- Cleared credential environment overrides in a subshell and selected AWS_PROFILE=aws-learning for the live program.
- After console deletion, ran only check-bucket mode; no transfer was rerun for absence verification.
- Unset CLOUD_ENG_DAY27_BUCKET, confirmed its absence without printing a value, and logged out.

## Cost and cleanup

S3 requests, storage, and data transfer can incur charges. The budget was checked. Actual billed cost and focused hours were not measured.

One tiny fictional object and its bucket were deleted in the same lab session. Console absence and the SDK NoSuchBucket category support cleanup. The bucket environment variable was unset and source logout succeeded. No other resources or policy changes were reported.

## Artifacts

- [Go transfer and cleanup checker](main.go)
- [Fictional payload](payload.txt)
- [Verified note](../../../docs/week-06/day-27-s3-presigned-transfer.md)
