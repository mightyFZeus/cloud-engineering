# Day 1 — AWS Account Safety

**Date:** September 14, 2026  
**Status:** Carried forward and completed through the Day 2 catch-up

## Objective

Establish safe account-access habits before deploying any AWS infrastructure.

## Work completed

- Protected the AWS root user with MFA.
- Verified that the root user has no access keys.
- Created a separate identity for routine administration.
- Enabled MFA on the everyday identity and confirmed non-root sign-in.
- Avoided creating long-lived programmatic access keys.

## Key lesson

The AWS root user has unrestricted account access and cannot be limited through IAM policies. It should be reserved for the small set of account-level operations that specifically require root credentials. Routine work belongs to a separate MFA-protected identity with only the permissions it needs.

## Security decision

For this personal learning account, I currently use a separate administrative IAM user as a temporary setup. The long-term target is federated access through IAM Identity Center and short-lived credentials.

## Reference

- [AWS root-user best practices](https://docs.aws.amazon.com/IAM/latest/UserGuide/root-user-best-practices.html)

