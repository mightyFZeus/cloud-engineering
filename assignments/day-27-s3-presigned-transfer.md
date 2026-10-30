# Day 27 — Private S3 Transfers with Presigned URLs

## Read first — 10 minutes

1. [Download and upload with presigned URLs](https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-presigned-url.html) — focus on the signed operation/object, bearer-token handling, and the earlier of configured expiry or signing-credential expiry.
2. [S3 examples for Go v2](https://docs.aws.amazon.com/sdk-for-go/v2/developer-guide/go_s3_code_examples.html) — focus only on PresignPutObject, PresignGetObject, and using their HTTP requests. Adapt the API shapes; replace example URL/raw-error logging with fixed safe categories.

**Date:** October 31, 2026
**Loaded:** October 30, 2026, after Day 26 verification
**Due:** October 31 at 11:00 PM Africa/Lagos
**Timebox:** 75–90 focused minutes, including cleanup and notes
**Status:** Assigned for October 31; loaded October 30. No submission received.
**Roadmap section:** Week 6 — storage and databases

## Outcome

Keep one S3 bucket private while using Go-generated, five-minute presigned PUT and GET requests to transfer one tiny fictional payload. Verify exact downloaded bytes, show that an unsigned request to the existing object is denied, and delete the object and bucket. Presigned URLs stay only in program memory and are never shown in chat, logs, files, or screenshots.

This demonstrates a private-object transfer capability. It does not establish IAM prefix isolation, build an authenticated URL-issuing backend, perform a backup/restore, or test actual expiration by waiting five minutes. The configured expiry is five minutes; the underlying temporary credentials can end access sooner.

## Numbered steps

1. Complete the reading. In the new README, write the intended flow: temporary source credentials -> presign one object operation -> ordinary HTTPS client uses that signed request -> S3 evaluates the request. Explain why possession of the signed URL matters even when the bucket is private.
2. Prepare labs/week-06/day-27-s3-presigned-transfer/README.md, main.go, payload.txt, and Go module files under /Users/mightyzeus/Desktop/cloud-eng. Use the fictional payload bytes day27 private transfer followed by a newline, and the object key reports/day27/sample.txt. Reuse the previous SDK configuration/sanitized-error pattern and add the S3 service dependency. Suggested allocation: 10 minutes reading, 30 guided Go/build, 15 bucket/authentication setup, 10 requests, and 25 cleanup/notes.
3. Implement Go using config.LoadDefaultConfig, explicit eu-west-1, and externally selected AWS_PROFILE=aws-learning. Read the actual bucket name only from the temporary environment variable CLOUD_ENG_DAY27_BUCKET; require a nonempty value. Use `<PRIVATE_LAB_BUCKET>` in saved documentation and never hard-code or print the actual name. Use bounded SDK/presigning contexts and an HTTP client timeout, such as 15 seconds per operation.
4. Add a transfer mode that presigns PutObject and GetObject for exactly reports/day27/sample.txt with Expires=5*time.Minute. Keep the resulting URLs and any signing-token/query data in memory. Use the presigner's supplied method, URL, and SignedHeader values. Match any signed content-type or checksum requirements when constructing the upload request. Do not alter a signed query or send an ACL header. Use a plain HTTP client, with automatic redirects disabled for this small lab, and safe fixed failure categories. Never print an error value from HTTP: url.Error can include the complete secret URL. Close every response body and avoid logging raw response/error bodies or headers.
5. The transfer mode must perform this sequence:
   - Send the presigned PUT with the exact payload bytes; require HTTP 200.
   - Send the presigned GET; require HTTP 200 and compare a bounded body read with the exact payload, including the newline. HTTP success or matching length alone is insufficient.
   - Build a fresh unsigned GET for the same verified-existing object's URL by removing all query parameters from a copy in memory. Send no Authorization header, signing headers, or session token. Require HTTP 403 after the successful signed download.
   - Print only safe statuses/check labels: upload check: PASS; download check: PASS; unsigned GET status: 403; private access check: PASS; presign expiry configured: 300 seconds. Exit 0 only when all checks succeed. Report the configured expiry without claiming the URLs were actually tested after expiry.
6. Add a separate read-only cleanup-check mode using SDK ListObjectsV2 with MaxKeys=1 for that bucket. Print BucketMissing and exit nonzero only when the exact service error is NoSuchBucket. A successful listing means the bucket still exists, even when empty. AccessDenied, NoSuchKey, a generic HTTP 404, a timeout, or a connection failure is not proof that the bucket is gone. This mode must not upload an object or create/recreate a bucket.
7. Run gofmt and build before creating the AWS bucket. Check eu-west-1 and the budget; renew aws-learning with the normal temporary-login workflow. Choose a globally unique learning bucket name with a random suffix privately, with no account ID or personal information. If the name is unavailable, choose a new one; do not reuse or modify an existing bucket. Set only the private temporary environment variable locally; do not send its value to chat or save it in repository files.
8. In the S3 console, create exactly one general purpose bucket in eu-west-1 using the shared global namespace. Confirm all four Block Public Access settings enabled, Object Ownership=Bucket owner enforced with ACLs disabled, SSE-S3 encryption, versioning disabled, Object Lock disabled, and learning-project tags. Keep the bucket initially empty. Add no bucket policy, CORS configuration, website hosting, replication, logging/inventory/analytics feature, access point, backup, or other integration. Use ordinary S3 Standard storage for the one object.
9. Run transfer mode once with the temporary source profile and credential environment overrides cleared in a subshell. Record only the actual upload/download statuses, exact-payload boolean, unsigned 403, safe check labels, configured expiry, and overall exit status. Confirm the object exists privately in the console. If a request fails, retain only its safe category/status; bound troubleshooting to 10 minutes, preserve the private settings, and prioritize cleanup. Do not publish a presigned URL to troubleshoot it.
10. Delete only the lab object through the console, verify the bucket is empty, then delete the lab bucket and confirm its disappearance. Run the read-only cleanup-check mode and require BucketMissing / NoSuchBucket with nonzero exit. Do not rerun transfer mode after deletion. Unset CLOUD_ENG_DAY27_BUCKET and log out aws-learning. Cleanup is required even after a blocked or partially successful run.
11. Save the code, module files, fictional payload, bucket-settings confirmations using the placeholder, safe results, and cleanup status. Explain why presigned access does not make the bucket public, why URLs must be protected, and why credential validity can shorten the configured five-minute lifetime. State explicitly that expiry was configured rather than tested by waiting, and that source logout alone is not evidence that every previously signed request was revoked. Run test -s on README.md and record its exit status. Inspect all saved files for URLs, signing query parameters, credentials, real bucket names/ARNs, account IDs, raw HTTP/SDK errors, and private configuration before submission.

## Cost and security

S3 can charge for requests, storage, and data transfer; checking the budget and [current pricing](https://aws.amazon.com/s3/pricing/) comes before resource creation. Keep this to one tiny S3 Standard object, a few requests, and one disposable bucket. Delete both in the same session. Presigning does not require making the bucket public or changing IAM policies. Create no role, access key, KMS key, compute, replica, backup, or supporting resource.

Treat every presigned URL as a temporary bearer credential. Never print, persist, paste, screenshot, or commit one, and never print complete HTTP errors that may embed it. Save no raw AWS/HTTP responses. Keep the actual bucket name private and use only `<PRIVATE_LAB_BUCKET>` in publishable notes. The source profile's authority is not claimed to be limited to this one prefix.

## Submission

1. Lab paths, fictional payload/key, gofmt completion, and build exit status:
2. Region/budget, one new bucket confirmed, all four public-access blocks, ownership/ACL setting, SSE-S3, versioning/Object Lock settings, and initial emptiness:
3. Configured presign lifetime and temporary source-profile confirmation:
4. Actual PUT status, GET status, exact-payload result, unsigned GET status, check labels, and overall exit status:
5. Private bucket versus presigned access, bearer-token handling, and configured-versus-tested-expiry explanations:
6. Code safeguards preventing URLs, headers, raw errors, real names/identifiers, or credentials from reaching saved output:
7. Object deleted; bucket empty/deleted/absent; read-only post-delete SDK category and nonzero exit:
8. Bucket environment variable unset; source logout result; other resources/policy changes:
9. README test -s exit status:
10. Exact safe blocker/status, if any:

## Pass criteria

- Exactly one new private general purpose bucket with the required safe settings and one fictional object; no unplanned policy/resource/features and actual bucket name absent from publishable artifacts.
- Reviewed Go uses the SDK default temporary-credential chain, explicit Region, bounded operations, correctly configured five-minute presigning, and safe in-memory HTTP request handling; formatting/build pass.
- Presigned PUT and GET return HTTP 200, and the downloaded bytes exactly match the fixture; signed headers/methods are preserved.
- A fresh unsigned GET for the same existing object returns 403, and all transfer checks complete with exit 0 without revealing a URL or sensitive output.
- Learner accurately explains bearer possession, private-object authorization, and configured lifetime versus credential lifetime; no unperformed expiry, IAM-isolation, or production-authentication claim.
- Object and bucket cleanup is corroborated by console and a read-only SDK NoSuchBucket result; environment cleanup/logout are confirmed, and substantive safe code/payload/README evidence is saved.
