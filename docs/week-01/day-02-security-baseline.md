# Day 2 — Security and Cost Baseline

**Date:** September 15, 2026  
**Status:** Passed — 8/8 checks

## Verified baseline

- Root MFA is enabled.
- Root access keys are absent.
- The everyday IAM identity has MFA and can sign in without root.
- The everyday identity has no access keys.
- A zero-spend alert triggers at USD 0.01 of actual spend.
- A USD 10 monthly budget alerts at 85% actual spend and 100% forecasted spend.
- Budget email notifications are configured.
- The default course Region is `eu-west-1` (Europe, Ireland).

## What I learned

Budgets provide early cost visibility, but they are notifications rather than guaranteed hard spending caps. Billing data and alerts can also be delayed. Cost safety therefore requires both alerts and disciplined cleanup of billable resources.

Identity safety and cost visibility are prerequisites for every later lab. I should know which identity I am using, which Region I am operating in, and what resources can continue accruing charges after a session ends.

## References

- [AWS IAM security best practices](https://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html)
- [AWS budget templates](https://docs.aws.amazon.com/cost-management/latest/userguide/budget-templates.html)

