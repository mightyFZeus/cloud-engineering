# Day 6 — Replacing Broad Access with Least Privilege

**Date:** September 18, 2026  
**Status:** Passed — 6/6 checks

## Objective

Replace a broad AWS-managed read-only policy with a narrowly scoped inline policy and verify the resulting effective permissions.

## Verified work

- Created the inline role policy `AWSLearningDiscoveryReadOnlyPolicy`.
- Detached the AWS-managed `ReadOnlyAccess` policy and confirmed that no managed policies remained attached.
- Confirmed that the role could still describe Regions, describe Availability Zones, and list S3 buckets.
- Observed 17 Regions, three Availability Zones in `eu-west-1`, and zero S3 buckets.
- Used policy simulation to verify that `ec2:DescribeRegions` and `s3:ListAllMyBuckets` were allowed.
- Verified that `s3:PutObject` and `iam:CreateUser` received implicit denies.
- Logged out of the temporary source session and created no workload resources or access keys.

## What I learned

The new policy is narrower because it permits only three specified API actions, while the AWS-managed `ReadOnlyAccess` policy permits read operations across many AWS services.

The `Action` element defines which API operations are allowed or denied. The `Resource` element defines which AWS resources those operations may target.

Some list and describe actions do not support permissions for individual resource ARNs, so they require `Resource: "*"`. The wildcard does not grant every AWS action; the policy remains narrow because its `Action` list contains only three operations.

## Security and cost

The broad managed policy was removed only after the narrow inline policy existed. No real write or IAM-administration request was made; denied operations were tested through simulation. No billable infrastructure, long-lived access key, account identifier, or complete ARN was created or recorded.

## References

- [Managed policies and inline policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_managed-vs-inline.html)
- [IAM policy testing with the policy simulator](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_testing-policies.html)
- [IAM policy elements](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements.html)
