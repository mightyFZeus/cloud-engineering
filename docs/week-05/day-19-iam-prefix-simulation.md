# Day 19 — Least-privilege S3 prefix simulation

Passed October 19, 2026: six assignment criteria verified through saved policy/README review, console screenshot evidence, and learner checkpoints.

## Verified evidence

- JSON syntax and README checks exited 0. Coach parsed the saved JSON and checked its exact action, resource, and prefix condition scope.
- Only the custom identity policy was selected; no permissions boundary supplied.
- GetObject reports/sample.txt: allowed.
- GetObject private/sample.txt: implicit deny.
- PutObject reports/sample.txt: implicit deny.
- ListBucket with explicit String context s3:prefix=reports/: allowed.
- ListBucket with explicit String context s3:prefix=private/: implicit deny.
- Learner explained that the policy grants reading reports but no upload permission.
- No policy attached to a real identity; no AWS resources created.

## Interpretation

GetObject uses an object resource restricted to reports/*. ListBucket uses the bucket resource plus a prefix condition. With only this policy supplied, requests without a matching allow are implicitly denied; this policy contains no explicit deny.

The initial GetObject simulation used the wildcard resource and returned implicit deny. Correcting the simulation input to the fictional reports object returned allow. The policy was not broadened.

Simulation does not issue real S3 operations or prove live access. Other policies, organization controls, boundaries, resource policies, and service context can affect real requests.

## Artifacts and safe commands

- [Policy](../../labs/week-05/day-19-iam-prefix-simulation/policy.json)
- [Results and explanation](../../labs/week-05/day-19-iam-prefix-simulation/README.md)
- Commands: python3 -m json.tool for local JSON syntax, test -s for README presence; console Custom mode for simulated authorization. No raw console output or real identifiers retained.

## References

- [IAM policy simulator](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_testing-policies.html)
- [How to simulate policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/policies_policy-simulator-how-to.html)
- [S3 policy conditions](https://docs.aws.amazon.com/AmazonS3/latest/userguide/amazon-s3-policy-keys.html)
