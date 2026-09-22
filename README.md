# AWS Cloud Engineering Journey

This repository documents my hands-on path from application developer to AWS cloud engineer. I have five years of React Native experience, 1.5 years of Go experience, and working familiarity with Linux.

The course is built around roughly 10 focused hours per week. Every daily milestone is completed, submitted, reviewed, corrected when necessary, and only then marked as passed.

## Learning goals

- Operate AWS accounts safely with non-root identities and temporary credentials.
- Build practical skills in Linux, networking, IAM, EC2, S3, containers, Terraform, CI/CD, observability, and troubleshooting.
- Run Kubernetes locally before deploying a short-lived Amazon EKS environment.
- Practice EKS Pod Identity, rollouts, rollbacks, monitoring, and failure diagnosis.
- Complete three portfolio projects with architecture, security, cost, deployment, and cleanup documentation.

## Current progress

| Day | Topic | Status | Notes |
|---:|---|---|---|
| 1 | AWS account security and cost guardrails | Completed through Day 2 | [Write-up](docs/week-01/day-01-account-safety.md) |
| 2 | Verify the security and billing baseline | Passed | [Write-up](docs/week-01/day-02-security-baseline.md) |
| 3 | AWS CLI with temporary credentials | Passed | [Write-up](docs/week-01/day-03-aws-cli.md) |
| 4 | Regions, Availability Zones, and shared responsibility | Passed | [Write-up](docs/week-01/day-04-global-infrastructure.md) |
| 5 | IAM roles and policy evaluation | Passed | [Write-up](docs/week-01/day-05-iam-roles.md) |
| 6 | Least-privilege IAM policy | Passed | [Write-up](docs/week-01/day-06-least-privilege.md) |
| 7 | AWS access baseline automation | Passed | [Write-up](docs/week-01/day-07-baseline-automation.md) · [Post draft](posts/week-01/day-07-baseline-automation.md) |
| 8 | First Linux EC2 instance and cleanup | Passed | [Write-up](docs/week-02/day-08-ec2-session-manager.md) · [Post draft](posts/week-02/day-08-ec2-session-manager.md) |

## Repository map

- [`docs/week-01/`](docs/week-01/) — verified daily learning notes and command summaries.
- [`docs/week-02/`](docs/week-02/) — EC2 and Linux operations notes.
- [`posts/week-01/`](posts/week-01/) — reviewed build-in-public drafts based on verified work.
- [`posts/week-02/`](posts/week-02/) — Week 2 build-in-public drafts.
- [`outputs/aws-cloud-engineer-12-week-plan.md`](outputs/aws-cloud-engineer-12-week-plan.md) — complete 12-week roadmap.
- [`outputs/aws-daily-assignment-tracker.md`](outputs/aws-daily-assignment-tracker.md) — assignment evidence and pass criteria.
- [`outputs/aws-learning-log-template.md`](outputs/aws-learning-log-template.md) — reusable learning-log format.
- [`outputs/aws-build-in-public-post-template.md`](outputs/aws-build-in-public-post-template.md) — reusable social-post format.

## Security rule

This repository must never contain AWS credentials, account IDs, complete ARNs, MFA secrets, login URLs, private keys, Terraform state, or other secrets. Examples use placeholders and non-sensitive identifiers only.
