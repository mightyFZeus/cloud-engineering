# Day 3 — AWS CLI with Temporary Credentials

**Date completed:** September 17, 2026  
**Status:** Passed — 8/8 checks

## Objective

Use AWS from a local terminal without creating a long-lived IAM access key.

## Configuration

- AWS CLI: version 2.36.47
- Authentication: browser-based `aws login`
- Profile: `aws-learning`
- Region: `eu-west-1`
- Output: JSON
- Principal: non-root IAM user

## Commands practiced

```bash
aws --version
aws login --profile aws-learning --region eu-west-1
aws configure get region --profile aws-learning
aws sts get-caller-identity --profile aws-learning
aws ec2 describe-regions --profile aws-learning --no-cli-pager
aws logout --profile aws-learning
```

The output of `get-caller-identity` is intentionally not recorded because it contains account and identity identifiers.

## What I learned

A named profile is a reusable collection of AWS CLI configuration and authentication-provider settings. It does not have to contain permanent credentials.

`sts get-caller-identity` identifies the account and principal signing a request. It does not prove that the principal can perform unrelated service actions.

Browser-based `aws login` is safer than a long-lived access key because it creates expiring credentials without requiring a reusable secret. The cached session can also be ended with `aws logout`.

## Reference

- [AWS CLI login with console credentials](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sign-in.html)

