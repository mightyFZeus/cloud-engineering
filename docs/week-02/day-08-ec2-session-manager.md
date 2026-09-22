# Day 8 — First Linux EC2 Instance

## Outcome

Launched one short-lived Amazon Linux 2023 `t3.micro` in `eu-west-1`, connected through Session Manager, inspected the host, then terminated the instance.

## Configuration and evidence

- The EC2 instance used `AWSLearningEC2SSMRole` with `AmazonSSMManagedInstanceCore`.
- The project tag was `Project=cloud-eng-journey`.
- The security group had no inbound rules and an IPv4 outbound rule allowing all traffic. No key pair was created.
- Session Manager connected, the SSM Agent was active, the root filesystem was 21% used, and two TCP sockets were listening.
- The kernel release was not recorded before termination.
- The instance was reported terminated with zero Day 8 EBS volumes remaining. The role and security group were retained for the next lab.

## Lessons

The AMI supplies the operating system image, the `t3.micro` instance type supplies 2 vCPUs and 1 GiB of memory, and the EBS root volume holds the boot filesystem. The security group controls network traffic; the instance role grants AWS API permissions to software on the VM. Session Manager uses outbound HTTPS from the SSM Agent, so an inbound SSH rule is unnecessary. In `/etc/os-release`, `ID_LIKE=fedora` does not change the OS identity: this instance ran Amazon Linux 2023.

## Cost and cleanup

The console launch estimate was reported as zero, but the final bill was not verified during the lab. EC2, EBS, and public IPv4 usage should be checked in Billing. No Day 8 EBS volume remained after termination.

## Next step

Run a local-only Go health endpoint as a `systemd` service and verify its status, port, logs, and restart behavior.
