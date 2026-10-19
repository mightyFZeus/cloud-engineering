# Day 19 — Simulate Least-Privilege S3 Prefix Access

**Date:** October 19, 2026
**Due:** 11:00 PM Africa/Lagos
**Timebox:** 60–75 minutes
**Status:** Passed October 19, 2026 — all six criteria verified.

## Outcome

Draft an identity policy that permits reading objects only under reports/ in one example bucket, then verify allowed and denied outcomes with the AWS IAM policy simulator. No real S3 requests or bucket creation required.

## Steps

1. Read the two authoritative links below for 10 minutes. Sign in using the normal MFA-protected identity. Do not create or attach a role or policy.
2. Create labs/week-05/day-19-iam-prefix-simulation/ with policy.json and README.md.
3. Write a policy with Version 2012-10-17. Allow s3:GetObject on arn:aws:s3:::amzn-s3-demo-cloud-eng/reports/*. Allow s3:ListBucket on arn:aws:s3:::amzn-s3-demo-cloud-eng only with StringLike condition s3:prefix equal to reports/*. These are fictional example resources. Do not use actual account IDs or deployment identifiers.
4. Validate JSON locally with python3 -m json.tool labs/week-05/day-19-iam-prefix-simulation/policy.json; record exit status.
5. Open the IAM policy simulator and select Custom mode. Paste only your draft policy; include no existing identity policies. If Custom mode is unavailable, report the exact blocker rather than attaching policies to a real identity.
6. Simulate GetObject for reports/sample.txt in the example bucket: expect allowed. Simulate GetObject for private/sample.txt: expect denied. Simulate PutObject for reports/sample.txt: expect denied. Set the object resource separately for each request.
7. Simulate ListBucket on the example bucket with context key s3:prefix (String) set to reports/: expect allowed. Repeat with private/: expect denied. Supply the context value deliberately; do not interpret missing context as a tested prefix restriction.
8. Record a five-row table of action, example resource/key or prefix, expected result, and actual simulator result. Explain bucket-level versus object-level resources and implicit deny in 3–4 sentences.
9. Explain that simulation evaluates supplied policies/context and does not prove live access in the presence of other policies, boundaries, SCPs, or endpoint policies. State no policy was attached and no AWS resources were created.

## References

- [IAM policy simulator](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_testing-policies.html)
- [S3 policy condition keys and prefix restrictions](https://docs.aws.amazon.com/AmazonS3/latest/userguide/amazon-s3-policy-keys.html)

## Cost and security

Simulation only; create no bucket, EC2, role, or attached policy. Use fictional example S3 resource strings only. Publish no raw identity details, credentials, account IDs, or real resource ARNs. If access is denied, report the sanitized error; do not broaden permissions yourself.

## Submission

1. Policy and README paths:
2. JSON validation exit status:
3. GetObject reports/ result:
4. GetObject private/ result:
5. PutObject reports/ result:
6. ListBucket with reports/ prefix result:
7. ListBucket with private/ prefix result:
8. Explanation of resource scope and implicit deny:
9. Simulator limitations:
10. Policies attached or AWS resources created:
11. Blocker:

## Pass criteria

- Valid JSON; exact read-only object prefix and conditional bucket listing.
- Five simulated decisions match the intended permissions.
- Listing tests supply the s3:prefix context explicitly.
- Correct distinction between object resource and bucket resource.
- Simulation limitations accurately stated.
- No real policy attachment or resource creation; safe saved evidence.
