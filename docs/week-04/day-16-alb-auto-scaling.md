# Day 16 — ALB and Auto Scaling

Passed October 13, 2026: 6/6 assignment criteria supported by console screenshots and learner checkpoints.

## Verified configuration and outcomes

- Region eu-west-1; budget and launch estimate confirmed checked.
- Internet-facing IPv4 ALB across eu-west-1a and eu-west-1b; HTTP 80 forwarded to an instance target group on HTTP 8080, health path /healthz.
- ALB inbound HTTP restricted to the learner's current address; application inbound 8080 restricted to the ALB security group, no SSH.
- Launch template: Amazon Linux 2023 x86_64, t3.micro, AWSLearningEC2SSMRole, no key pair, IMDSv2 required, 8 GiB gp3 delete on termination, Project tags on instances and volumes, corrected user data.
- Auto Scaling desired/minimum/maximum 2/2/2, ELB checks enabled, 300-second grace period.
- Two InService instances and two healthy targets; direct HTTP curl returned ok.
- Deliberate termination followed by replacement restored two InService instances, two healthy targets, and ok.

## Explanation and troubleshooting

Auto Scaling observes EC2 state and health and restores desired capacity using the launch template. ALB health checks send HTTP requests to /healthz to assess application readiness. With ELB checks enabled, Auto Scaling can also use unhealthy target reports to replace instances; this lab directly demonstrated termination recovery, not application-failure replacement.

A browser timed out while a direct HTTP request with curl succeeded. The browser-specific root cause was not established. No configuration change was proven to fix it.

Safe command summary: curl with an HTTP health endpoint, connection timeout, overall timeout, and proxy bypass. The actual deployment address and raw output are omitted.

## Cleanup

All lab instances terminated; ASG, ALB, target group, launch template, and both security groups deleted. Zero Day 16 EBS volumes remain. No NAT gateways or Elastic IPs created.

## References

- https://docs.aws.amazon.com/autoscaling/ec2/userguide/health-checks-overview.html
- https://docs.aws.amazon.com/elasticloadbalancing/latest/application/target-group-health-checks.html
