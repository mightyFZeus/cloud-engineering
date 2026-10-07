# Day 15 — Launch and Verify the Go Service from an EC2 Launch Template

**Date:** Thursday, October 8, 2026

**Due:** 11:00 PM Africa/Lagos

**Timebox:** 75–90 minutes

## Outcome

Create a reusable EC2 launch template from the Day 14 bundle, launch one short-lived Amazon Linux instance, verify that user data starts the Go health service, and delete every temporary resource.

## Learn — 10–15 minutes maximum

- Review [Create an Amazon EC2 launch template](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/create-launch-template.html).
- Review [Session Manager prerequisites](https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-prerequisites.html).

## Build — 45–55 minutes

1. Sign in with the everyday MFA-protected identity, select `eu-west-1`, check the budget dashboard, and review the launch estimate before creating the instance.
2. In the default VPC, create `cloud-eng-day15-alb-sg` with no inbound rules. This represents the future load balancer's identity. Keep normal outbound HTTPS available.
3. Create `cloud-eng-day15-app-sg` with exactly one inbound rule:
   - Type: Custom TCP
   - Port: `8080`
   - Source: `cloud-eng-day15-alb-sg`
   Do not add SSH or public-CIDR inbound rules. Keep outbound HTTPS available so user data and Systems Manager can reach required endpoints.
4. Create launch template `cloud-eng-day15-health` with:
   - latest Amazon Linux 2023 x86_64 AMI offered by the console;
   - `t3.micro`;
   - no key pair;
   - the existing `AWSLearningEC2SSMRole` instance profile;
   - `cloud-eng-day15-app-sg`;
   - an 8 GiB `gp3` root volume with delete-on-termination enabled;
   - instance metadata version set to require IMDSv2;
   - `Project=cloud-eng-journey` tags for the instance and volume; and
   - the complete contents of `labs/week-04/day-14-ec2-user-data/user-data.sh` as user data.
5. Launch one instance from the template in a default-VPC public subnet. Ensure it receives a public IPv4 address for outbound package installation, while its security group still permits no public inbound traffic.
6. Wait for the instance to become a Systems Manager managed node, then connect with Session Manager. Do not use SSH.
7. Run these safe checks inside the session:

   ```bash
   systemctl is-active cloud-eng-health
   curl -fsS http://127.0.0.1:8080/healthz
   ss -lnt | grep ':8080'
   sudo journalctl -u cloud-eng-health -n 20 --no-pager
   sudo tail -n 30 /var/log/cloud-init-output.log
   ```

8. Record only the service status, health response, listening address, and whether cloud-init completed successfully. Do not submit instance IDs, public or private IP addresses, account IDs, ARNs, or raw logs.

## Clean up — 10–15 minutes

1. Terminate the instance and wait until it reaches `Terminated`.
2. Confirm its Day 15 EBS volume was deleted.
3. Delete the launch template.
4. Delete `cloud-eng-day15-app-sg`, then `cloud-eng-day15-alb-sg`.
5. Confirm no Day 15 EC2 instance, EBS volume, launch template, or security group remains.

## Safety and cost rules

- This lab creates one short-lived EC2 instance and may incur a small compute and storage charge. Review the estimate first and terminate it during the same session.
- Create no NAT gateway, Elastic IP, load balancer, target group, or Auto Scaling group today.
- Use Session Manager with the instance role; do not create a key pair or open SSH.
- Do not submit resource IDs, assigned IP addresses, account IDs, ARNs, credentials, tokens, user-data output, or private console URLs.
- If the default VPC is absent or its subnet cannot provide outbound internet access, stop and report the blocker. Do not create a NAT gateway for this lab.

## Submit for verification

```text
Day 15 submission

1. Region; budget and launch estimate checked:
2. AMI family and instance type:
3. Launch template name; IMDSv2 required:
4. Instance role; key pair:
5. Application security-group inbound rule count and source type:
6. Root volume type/size; delete on termination:
7. Session Manager connected: yes/no
8. Service status; health response; listening address:
9. Cloud-init completed successfully: yes/no
10. Explain why port 8080 trusts the ALB security group instead of a public CIDR:
11. Instance terminated; Day 15 EBS volumes remaining:
12. Launch template and both Day 15 security groups deleted: yes/no
13. NAT gateways, Elastic IPs, load balancers, and Auto Scaling groups created:
14. Exact blocker, if any:
```

## Pass criteria

- [x] The launch template used Amazon Linux 2023, `t3.micro`, IMDSv2, the SSM instance role, no key pair, and the Day 14 user data.
- [x] The application security group allowed port `8080` only from the placeholder ALB security group and exposed no SSH or public inbound rule.
- [x] Session Manager connected and the service returned `active` and `ok` while listening on `:8080`.
- [x] Cloud-init completed successfully without credentials or secrets in user data.
- [x] No NAT gateway, Elastic IP, load balancer, target group, or Auto Scaling group was created.
- [x] The instance, EBS volume, launch template, and both temporary security groups were deleted.

## Verification — October 7, 2026

**Result:** Passed, 6/6 criteria.

- The launch template used Amazon Linux 2023 x86_64, `t3.micro`, the existing SSM instance profile, no key pair, an 8 GiB `gp3` root volume with delete-on-termination, required IMDSv2, project tags, and the Day 14 user data.
- The application security group contained one inbound TCP `8080` rule whose source was the placeholder ALB security group. It contained no SSH or public-CIDR inbound rule.
- The first corrected-template launch produced `cloud-init status: error` and an inactive service. The cloud-init log showed that Go could not locate a build cache because neither `GOCACHE`, `XDG_CACHE_HOME`, nor `HOME` was defined.
- The repair explicitly set `HOME=/root`, created a root-owned build-cache directory, and set `GOCACHE` to that directory. A new launch-template version was created and used for a replacement instance.
- On the replacement, cloud-init returned `done`, the service returned `active`, `/healthz` returned `ok`, and the socket listened on `*:8080` without service errors.
- The failed and replacement instances were terminated. Zero Day 15 EBS volumes remained, and the launch template and both temporary security groups were deleted.
- No NAT gateway, Elastic IP, load balancer, target group, or Auto Scaling group was created. The sensitive-data scan excluded resource identifiers and raw logs.
