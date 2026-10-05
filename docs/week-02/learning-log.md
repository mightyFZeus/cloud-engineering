# Week 2 Learning Log — Linux Operations, EC2, EBS, and S3

## Outcome

Week 2 connected application code, Linux service management, EC2 access, and S3 authorization. The completed labs used short-lived Amazon Linux 2023 `t3.micro` instances in `eu-west-1`, Session Manager instead of inbound SSH, and explicit cleanup checks after each exercise.

## What I built and verified

The first lab established the EC2 operating baseline. An instance used `AWSLearningEC2SSMRole`, a security group with zero inbound rules, no key pair, and `Project=cloud-eng-journey`. Session Manager connected successfully, the SSM Agent was active, the root filesystem was 21% used, and two TCP sockets were listening. The instance was terminated and no Day 8 EBS volume remained.

The second lab ran a dependency-free Go health endpoint under `systemd`. The service listened only on `127.0.0.1:8080`; `systemctl` reported it active, `ss` showed the expected listener, `curl` returned HTTP 200 with `ok`, and `journalctl` showed the startup events. Sending SIGKILL to the main process tested the recovery policy. `systemd` restarted the process and the health check returned successfully again. The journal also warned about using the shared `nobody` account, showing why a dedicated or dynamic service identity is a better long-term choice.

The third lab attached narrowly scoped S3 permissions to the existing instance role. The bucket kept all four Block Public Access settings enabled and used AES256 server-side encryption. The role policy applied only to one bucket and its objects. Both `aws s3 cp` and `aws s3 sync` succeeded, a downloaded file matched its source, and the final test count was four objects. The instance, EBS volume, objects, bucket, and temporary policy were removed.

## Security and architecture decisions

Session Manager allowed host access without an SSH key or inbound security-group rule. The S3 workload used temporary credentials from the instance role instead of embedded access keys. IAM governed which authenticated operations were permitted, while S3 Block Public Access guarded against public exposure. The temporary S3 policy was removed after the lab so the reusable role did not keep unnecessary permissions.

## Runbook — service active but endpoint unreachable

1. Run `systemctl is-active cloud-eng-health` to determine whether the service manager sees a running process.
2. Run `ss -lnt` and locate the expected address and port. A listener on `127.0.0.1` is reachable only from the instance itself.
3. Run `curl` locally against the health endpoint. This separates application behavior from remote network access.
4. Read recent unit logs with `journalctl -u cloud-eng-health` for startup, bind, permission, and crash evidence.
5. If local HTTP works but a permitted remote path does not, inspect the application bind address, security-group rules, subnet route table, and network ACL in that order. Check DNS only after verifying the destination address and traffic path.
6. Change one demonstrated cause, repeat the failed check, and record the recovery result instead of making several speculative changes.

## Troubleshooting lessons

Process state, socket state, HTTP response, and logs answer different questions. A running process does not prove that the expected port is open, and a local listener does not prove that a remote client has a valid network path. Review also caught an inconsistent S3 object count; comparing it with the known upload layout prevented an unsupported result from being accepted.

## Cost and cleanup

Each EC2 instance was terminated after its lab, and the reported remaining Day 8, Day 9, and Day 10 EBS volume counts were zero. The temporary S3 data, bucket, and inline policy were deleted. The base SSM role and zero-inbound security group remain for later controlled labs.

## Explain it in an interview

I used Session Manager and an EC2 instance role to operate a Linux workload without opening SSH or storing AWS keys. I ran a Go health endpoint under `systemd`, verified it across process, socket, HTTP, and log layers, and tested automatic restart after a forced process failure. I then scoped the role to one private S3 bucket, verified public-access blocking and encryption, tested file transfer and integrity, and removed all temporary resources. The main operational lesson was to attach each claim to a distinct check and to close every lab with permission and cost cleanup.
