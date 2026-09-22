# Week 1 Learning Log — AWS Foundations

**Dates:** September 14–20, 2026  
**Status:** In progress  
**Planned time:** 10 hours

## Goals

- [x] Establish a secure account and cost baseline.
- [x] Use AWS CLI temporary credentials without long-lived access keys.
- [x] Explain Regions, Availability Zones, and shared responsibility.
- [x] Assume an IAM role and inspect policy-evaluation outcomes.
- [x] Replace broad read-only access with a narrowly scoped policy.
- [ ] Complete the Week 1 recall review.

## Verified progress

| Day | Outcome | Evidence |
|---:|---|---|
| 1 | Root access reserved for exceptional account actions | Root MFA and absence of root access keys verified |
| 2 | Secure everyday identity and cost guardrails | MFA, non-root sign-in, zero-spend alert, and monthly budget verified |
| 3 | Temporary CLI authentication | Named profile, caller identity, Region query, and logout verified |
| 4 | Global infrastructure and responsibility model | Three-AZ map and EC2/Lambda comparison verified |
| 5 | IAM roles and policy evaluation | Assumed-role profile, safe read, and simulated denied write verified |
| 6 | Least-privilege IAM policy | Three permitted discovery actions worked; S3 write and IAM administration remained denied |
| 7 | AWS access baseline automation | Script checked the role and Region, returned narrow counts, rejected a missing profile, and passed syntax review |

## Most important lessons

- A Region contains multiple isolated Availability Zones, but resilience requires deploying the workload across those AZs and routing around failures.
- Serverless reduces infrastructure operations; it does not remove responsibility for application code, dependencies, permissions, secrets, monitoring, security, or data.
- Temporary role credentials reduce the risk created by reusable long-lived secrets.
- Permission policies answer what an identity may do; trust policies answer who may assume a role.
- Lack of an allow produces an implicit deny, while an explicit deny overrides applicable allows.
- `Resource: "*"` can be necessary for list and describe actions that do not support individual resource ARNs; least privilege can still be enforced with a narrowly limited `Action` list.

## Troubleshooting

A temporary AWS Sign-In endpoint connection failure prevented the CLI from refreshing credentials. The dependent configuration steps were paused, authentication was re-established, and identity checks were rerun before continuing.

## Security and cost

No credentials, account identifiers, complete ARNs, or raw identity output are stored in this repository. No billable workload resources were created during Days 1–7. The read-only learning role remains available for later labs.

## Next actions

- Finish the Week 1 recall review.
- Begin Linux, EC2, EBS, and S3 operations in Week 2.
- Keep all billable resources short-lived and document cleanup evidence.
