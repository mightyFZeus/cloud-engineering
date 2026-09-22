# Day 9 — Go Health Service Under systemd

## Outcome

Built a dependency-free Go health service, ran it under `systemd` on a short-lived Amazon Linux 2023 `t3.micro` in `eu-west-1`, verified its local endpoint and restart behavior, then cleaned up the instance and storage.

## Configuration and evidence

- The source is saved at `labs/week-02/day-09-health/main.go`; the unit is saved beside it as `cloud-eng-health.service`.
- The instance used `AWSLearningEC2SSMRole`, a security group with zero inbound rules, no key pair, and the tag `Project=cloud-eng-journey`.
- Go `go1.26.8-X:nodwarf5` built the binary successfully.
- The service ran as the non-root `nobody` user with `Restart=on-failure`.
- `systemctl` reported `active`, `curl` returned HTTP 200 with `ok`, and `ss` showed a listener on `127.0.0.1:8080`.
- The journal recorded the service start and application listener. It also warned that the shared `nobody` account is unsafe for a long-lived service; a dedicated or dynamic service identity is the appropriate follow-up.
- After the main process received SIGKILL, `systemd` restarted it and the health response returned.

## Troubleshooting model

`systemctl` shows whether the service process is running or stopped. `ss` shows whether a process has opened the expected listening socket. `curl` verifies the HTTP status and response body, while `journalctl` exposes startup failures, crashes, and application errors. Together, these tools separate process, listener, protocol, and application-log failures.

## Security and cleanup

The endpoint bound only to the loopback address, and no inbound security-group rule exposed port 8080. The source and unit contain no credentials or AWS identifiers. The instance was terminated and zero Day 9 EBS volumes remained.

## Next step

Use an EC2 instance role to read and write a narrowly scoped private S3 bucket, verify public access remains blocked, and delete all temporary resources.
