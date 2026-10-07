# Day 16 — Deploy the Go Service Behind an ALB and Auto Scaling Group

**Date:** Friday, October 9, 2026

**Due:** 11:00 PM Africa/Lagos

**Timebox:** 90 minutes

## Outcome

Run two copies of the Go health service across two Availability Zones behind an internet-facing Application Load Balancer, prove that Auto Scaling replaces a terminated instance, and delete every paid resource in the same session.

## Learn — 10–15 minutes maximum

- Review [Create an Application Load Balancer](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-application-load-balancer.html).
- Review [Attach a load balancer target group to an Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/attach-load-balancer-asg.html).

## Build — 55–65 minutes

1. Sign in with the everyday MFA-protected identity, select `eu-west-1`, check the budget dashboard, and review the estimated cost of one ALB plus two short-lived `t3.micro` instances.
2. In the default VPC, create:
   - `cloud-eng-day16-alb-sg`, allowing inbound HTTP port `80` only from **My IP** and retaining normal outbound access; and
   - `cloud-eng-day16-app-sg`, allowing inbound TCP port `8080` only from `cloud-eng-day16-alb-sg`, with no SSH or public-CIDR inbound rule.
3. Create target group `cloud-eng-day16-tg`:
   - target type: Instances;
   - protocol and port: HTTP `8080`;
   - health-check path: `/healthz`;
   - VPC: default VPC; and
   - do not register instances manually.
4. Create launch template `cloud-eng-day16-health` using the corrected Day 14 user data:
   - Amazon Linux 2023 x86_64 and `t3.micro`;
   - no key pair;
   - existing `AWSLearningEC2SSMRole` instance profile;
   - `cloud-eng-day16-app-sg`;
   - 8 GiB `gp3`, delete on termination;
   - IMDSv2 required; and
   - `Project=cloud-eng-journey` instance and volume tags.
   Do not add a subnet, public-IP setting, or network-interface block to the template.
5. Create internet-facing ALB `cloud-eng-day16-alb`:
   - IPv4;
   - default VPC;
   - two default public subnets in different Availability Zones;
   - `cloud-eng-day16-alb-sg`; and
   - HTTP port `80` listener forwarding to `cloud-eng-day16-tg`.
6. Create Auto Scaling group `cloud-eng-day16-asg`:
   - use the launch template's latest/default version;
   - use the same two public subnets;
   - attach `cloud-eng-day16-tg`;
   - turn on Elastic Load Balancing health checks;
   - health-check grace period: 300 seconds;
   - desired capacity: 2, minimum: 2, maximum: 2; and
   - no scaling policy for this lab.
7. Wait until both instances are `InService` and both targets are healthy. Open the ALB DNS name with `/healthz` and record only the `ok` response. Do not submit the DNS name.
8. Terminate one Auto Scaling instance from the EC2 Instances page. Observe the Auto Scaling activity, replacement instance, and target health until the group again has two `InService` instances and two healthy targets.

## Clean up — 15–20 minutes

1. Set the Auto Scaling group's minimum and desired capacities to `0`, then wait for its instances to terminate.
2. Delete the Auto Scaling group.
3. Delete the Application Load Balancer and wait until deletion completes.
4. Delete the target group and launch template.
5. Confirm zero Day 16 EBS volumes remain.
6. Delete `cloud-eng-day16-app-sg`, then `cloud-eng-day16-alb-sg`.
7. Confirm no Day 16 ALB, target group, Auto Scaling group, launch template, EC2 instance, EBS volume, NAT gateway, or Elastic IP remains.

## Safety and cost rules

- An Application Load Balancer, EC2 instances, and EBS volumes can incur charges. Create them in one focused session and clean them up immediately after verification.
- Use only two `t3.micro` instances and one ALB. Create no NAT gateway or Elastic IP.
- Restrict the ALB's HTTP rule to **My IP**. Do not submit that address or the ALB DNS name.
- Instances use Session Manager and an IAM role. Create no key pair and no SSH rule.
- Submit no resource IDs, account IDs, ARNs, assigned IP addresses, raw logs, credentials, tokens, or private console URLs.

## Submit for verification

```text
Day 16 submission

1. Region; budget and estimate checked:
2. ALB security-group inbound rule; application security-group inbound rule:
3. Target-group port and health-check path:
4. Launch-template AMI family, type, instance role, key pair, and IMDSv2 setting:
5. ALB scheme; Availability Zone count; listener:
6. Auto Scaling desired/minimum/maximum; ELB health checks enabled:
7. Initial InService instance count; healthy target count; ALB `/healthz` response:
8. One instance terminated deliberately: yes/no
9. Replacement observed; final InService and healthy counts:
10. Explain what detected the failure and what performed the replacement:
11. ASG, ALB, target group, launch template, instances, and security groups deleted: yes/no
12. Day 16 EBS volumes remaining:
13. NAT gateways and Elastic IPs created:
14. Exact blocker, if any:
```

## Pass criteria

- [ ] The ALB spanned two Availability Zones and accepted HTTP only from the learner's current address.
- [ ] The application security group accepted port `8080` only from the ALB security group.
- [ ] The Auto Scaling group maintained two instances across two subnets and registered them with the target group.
- [ ] Two targets became healthy and the ALB returned `ok` from `/healthz`.
- [ ] Terminating one instance caused Auto Scaling to create a replacement that became healthy.
- [ ] All paid and temporary Day 16 resources were deleted, with zero remaining EBS volumes and no NAT gateway or Elastic IP created.
