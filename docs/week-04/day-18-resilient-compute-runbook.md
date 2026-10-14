# Day 18 — Resilient compute runbook

Passed October 14, 2026: six criteria verified by reading the saved runbook and reviewing the learner checkpoints.

- Two-AZ public-subnet topology accurately documents the actual Day 16 lab, ALB listener, target port, health checks, security groups, IAM role, and replacement control.
- Seven-sentence request/response explanation is correct; learner explained backend port 8080 accepts only the ALB security group.
- Six failure rows distinguish symptoms, evidence, detection, recovery, and repair. The AZ row explicitly limits recovery by capacity and configuration.
- Ordered checklist proceeds from client reachability to backend process, probe, and startup evidence.
- Tested instance replacement and local probe results are separated from untested outage/dependency scenarios; browser timeout root cause remains unknown.
- Nonempty README check exited 0; no new AWS resources created; prior Day 16 cleanup remains confirmed.
- Coach corrected two relative evidence links; all three links now resolve locally.

See the [runbook](../../labs/week-04/day-18-resilient-compute-runbook/README.md).

## References

- [ALB health checks](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/target-group-health-checks.html)
- [Auto Scaling health checks](https://docs.aws.amazon.com/autoscaling/ec2/userguide/health-checks-overview.html)
