# Day 15 — EC2 Launch Template Verification

## Outcome

Created a reusable EC2 launch template, diagnosed a first-boot failure from cloud-init evidence, repaired the user-data environment, and verified the Go health service on a replacement Amazon Linux 2023 instance.

## Launch configuration

The template used Amazon Linux 2023 x86_64, `t3.micro`, the existing SSM instance profile, no key pair, an 8 GiB `gp3` root volume with delete-on-termination, required IMDSv2, and project tags. The application security group accepted TCP port `8080` only from a placeholder ALB security group and exposed no SSH or public-CIDR inbound rule.

## Failure diagnosis and repair

The first corrected-template launch reported `cloud-init status: error`, and the service was inactive. File inspection showed that Go and the service user existed, but the binary and unit file had not been installed. The cloud-init log then identified the cause: the Go build could not locate its cache because the cloud-init environment defined neither `HOME` nor `GOCACHE`.

The user-data script was changed to set `HOME=/root`, create a root-owned cache directory, and point `GOCACHE` to it. The Bash syntax and sensitive-data checks passed before a new launch-template version was created.

## Successful verification

The replacement instance completed cloud-init with `status: done`. Session Manager connected without SSH or a key pair. The service was `active`, `/healthz` returned `ok`, and the listening socket was `*:8080` with no service errors.

## Cleanup

- The failed and replacement instances were terminated.
- Zero Day 15 EBS volumes remained.
- The launch template and both Day 15 security groups were deleted.
- No NAT gateway, Elastic IP, load balancer, target group, or Auto Scaling group was created.
- Resource identifiers, assigned addresses, and raw logs were excluded from the repository.

## References

- [Create an EC2 launch template](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/create-launch-template.html)
- [EC2 user data](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/user-data.html)
- [Session Manager prerequisites](https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-prerequisites.html)
