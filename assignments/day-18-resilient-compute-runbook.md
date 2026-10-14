# Day 18 — Consolidate the Resilient Compute Runbook

**Date:** October 14, 2026
**Due:** 11:00 PM Africa/Lagos
**Timebox:** 60–75 minutes
**Status:** Assigned

## Outcome

Create a concise portfolio runbook linking your verified VPC, ALB, Auto Scaling, and local troubleshooting evidence. Complete the remaining availability analysis without rebuilding paid infrastructure.

## Numbered steps

1. Spend 10 minutes reading the two references below. Review Days 13, 16, and 17 notes.
2. Create labs/week-04/day-18-resilient-compute-runbook/README.md. Record the actual Day 16 topology: two public default subnets across two AZs, internet-facing ALB, target group, two instances, application security group, and instance role. Explicitly state that this lab used public subnets; do not describe it as a private-subnet production deployment.
3. Add a labeled architecture sketch (plain text is sufficient). Show browser HTTP 80 to ALB, ALB HTTP 8080 to targets, /healthz checks, and ASG replacement. Include both AZs, subnet route to IGW, and security-group source restrictions. Separate request traffic from replacement control actions.
4. Explain the request and response path in 5–7 sentences. Explain which control restricts inbound access and why a public subnet alone does not make an application publicly reachable.
5. Reuse your Day 17 failure table and add an impaired-AZ row. Include evidence, detector, possible routing/capacity recovery, and human action. Describe capacity limitations; do not claim you tested an AZ outage or guaranteed uninterrupted service.
6. Add an ordered troubleshooting checklist: client HTTP/port/source address, ALB state/listener, public subnet routes, target health, backend security group, process/listener, probe path, user-data/service logs. Explain why the successful curl result narrowed but did not prove the browser timeout's cause.
7. Add the tested deployment evidence and cleanup summary, linking relative paths to existing notes. Distinguish observed results from design analysis; state no new AWS resources were created.
8. Run test -s labs/week-04/day-18-resilient-compute-runbook/README.md and record its exit status. Review for AWS IDs, complete ARNs, deployment addresses, raw logs, and secrets before submitting.

## References

- [Auto Scaling health checks](https://docs.aws.amazon.com/autoscaling/ec2/userguide/health-checks-overview.html)
- [ALB target health checks](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/target-group-health-checks.html)

## Cost and security

Local documentation only: create no AWS resources. Use labels rather than real resource identifiers or deployment addresses. Preserve existing lab notes; do not overwrite them.

## Submission

1. README path:
2. Two-AZ topology sketch included:
3. Request/response explanation included:
4. Failure-table row count, including impaired AZ:
5. One AZ failure row (evidence, detection, possible recovery, limitations, repair):
6. Ordered troubleshooting checklist included:
7. Observed evidence separated from analysis:
8. test -s exit status:
9. AWS resources created:
10. Blocker:

## Pass criteria

- Sketch accurately represents the actual public-subnet lab and both AZs.
- Request/response explanation correctly covers listener, targets, routes, and security groups.
- Six failure rows correctly distinguish detection, recovery, limits, and human action.
- Checklist proceeds from client reachability to backend/process/probe evidence.
- Claims match verified Days 13, 16, and 17 evidence; no untested outage claim.
- Nonempty saved runbook, safe contents, and no new AWS resources.
