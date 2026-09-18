# Day 5 — IAM Roles and Policy Evaluation

**Date:** September 18, 2026  
**Status:** Passed — 6/6 checks

## Objective

Create a role, assume it with temporary credentials through a second AWS CLI profile, and verify the difference between allowed and denied actions without executing a write request.

## Verified work

- Created `AWSLearningReadOnlyRole` with the AWS-managed `ReadOnlyAccess` policy and the required learning tags.
- Configured the `aws-learning-readonly` profile to use `aws-learning` as its source profile.
- Confirmed that the second profile operated as an assumed-role session.
- Successfully performed a read-only S3 bucket-list operation; the account contained zero buckets.
- Used IAM policy simulation rather than a real write request.
- Verified that `s3:ListAllMyBuckets` was allowed and `s3:PutObject` received an implicit deny.
- Logged out of the source profile and created no workload resources or long-lived access keys.

## What I learned

An IAM user is a persistent identity that can have its own credentials. A role has no permanent credentials and is assumed temporarily by an authorized principal.

A permission policy determines what the role may do after it is assumed. A trust policy determines which principals are allowed to assume the role.

Implicit deny is AWS's default when no applicable policy allows an action. An explicit deny is written into a policy and overrides an allow from another applicable policy.

## Troubleshooting note

An initial AWS Sign-In endpoint connection failure interrupted the temporary-login workflow. The configuration steps were paused until authentication worked again so an empty account identifier could not be saved into the role profile. After temporary authentication was re-established, the role assumption and validation checks succeeded.

## Security and cost

The account identifier and complete role ARN were inspected locally and were not recorded. No real S3 write request was made. IAM roles and policy simulation did not create billable infrastructure; the learning role was retained for later exercises.

## References

- [IAM roles](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html)
- [Policy evaluation logic](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html)
- [AWS CLI role configuration](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-role.html)
