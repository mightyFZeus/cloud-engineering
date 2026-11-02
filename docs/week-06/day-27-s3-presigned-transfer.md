# Day 27 — Private S3 Transfers with Presigned URLs

## References

- [Presigned URLs](https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-presigned-url.html)
- [Go SDK S3 examples](https://docs.aws.amazon.com/sdk-for-go/v2/developer-guide/go_s3_code_examples.html)

**Verified:** November 2, 2026

**Roadmap:** Week 6 — storage and databases

**Assessment:** Passed, six criteria supported by learner-run console/program results, explanations, and inspection of saved artifacts.

## Outcome

A guided Go SDK v2 program presigned PUT and GET requests for one fictional object while the S3 bucket remained private. Temporary source-profile credentials signed the requests; the HTTP client used the resulting URLs only in memory.

The bucket was a new general purpose bucket in eu-west-1 with all four public-access blocks enabled, Bucket owner enforced/ACLs disabled, SSE-S3, versioning disabled, Object Lock disabled, and no additional policies/features reported. Its actual name is omitted; use `<PRIVATE_LAB_BUCKET>` in notes.

## Observed evidence

| Evidence | Result |
|---|---|
| Budget, source login, authentication | Checked; login and identity check exit 0 |
| Module and fixture | Correct payload including newline; module/S3 setup exit 0 |
| Formatting, tidy, build | Exit 0 each |
| Configured presign lifetime | 300 seconds |
| Signed upload | HTTP 200; PASS |
| Signed download | HTTP 200; exact payload true; PASS |
| Fresh unsigned request | HTTP 403; private-access check PASS |
| Overall transfer | Exit 0; object reported visible in console |
| Console cleanup | Object removed; bucket empty, deleted, and absent after refresh |
| Read-only SDK verification | BucketMissing from exact NoSuchBucket; exit 1 |
| Environment and logout | Bucket variable unset; logout exit 0 |
| Other resources or policy changes | None reported |
| Saved README | Nonempty check exit 0; substantive evidence inspected |

The saved source matches the reviewed implementation and the payload matches the expected bytes. An offline build and fake-transport protocol checks passed. Live AWS checks were learner-run and were not independently repeated by the coach.

## What the checks establish

The signed requests used the issuer's authorization without making the bucket public. A fresh unsigned request to the same existing object was denied. The program preserves signed methods/headers and verifies exact downloaded content rather than treating HTTP success as sufficient.

URLs are bearer credentials. Their configured lifetime is limited by signing-credential validity; policy or resource changes may also prevent access. The learner explained these limits and explicitly distinguished configured expiry from testing an expired request.

Coach clarification: CLI logout success alone is not a revocation test for previously signed requests. No universal revocation claim is made.

## Safe implementation

The code uses the default credential chain, explicit Region, bounded SDK and HTTPS operations, an explicit upload checksum, disabled redirects, bounded body reads, closed bodies, and fixed safe output categories. URLs, headers, actual bucket names, raw HTTP errors/responses, and credentials are not retained in publishable evidence.

The cleanup mode is separate and read-only. It uses ListObjectsV2 with MaxKeys=1 and classifies only the precise NoSuchBucket error as BucketMissing.

## Cleanup and limits

The single object and bucket were removed in the same lab session. The bucket variable was unset and the source CLI session logged out. Actual cost and study hours were not measured.

The five-minute expiry was configured, not tested by waiting. IAM prefix isolation, production URL-issuer authentication, version recovery, and a complete storage/serverless milestone remain outside this verified result.

- [Lab evidence and command summary](../../labs/week-06/day-27-s3-presigned-transfer/README.md)
- [Program](../../labs/week-06/day-27-s3-presigned-transfer/main.go)
- [Fictional payload](../../labs/week-06/day-27-s3-presigned-transfer/payload.txt)
- [Social drafts](../../posts/week-06/day-27-s3-presigned-transfer.md)
