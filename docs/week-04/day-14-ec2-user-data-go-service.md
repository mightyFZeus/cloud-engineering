# Day 14 — EC2 User Data Go Service

## Outcome

Packaged the Day 9 Go health service as a locally validated EC2 user-data deployment bundle suitable for an Application Load Balancer target.

## Application

The Go service provides `GET /healthz`, returns HTTP `200` with `ok`, listens on `:8080`, and configures read-header, read, write, and idle timeouts. Listening on all instance interfaces allows an ALB to reach the service over the instance's private network interface.

## Service isolation

The `systemd` unit runs as the dedicated `cloudeng` user, starts after `network-online.target`, and uses `Restart=on-failure`. It enables `NoNewPrivileges`, `PrivateTmp`, `ProtectHome`, and strict filesystem protection.

The instance security group should allow TCP port `8080` only from the ALB security group. This keeps the application port unavailable to arbitrary internet sources while allowing the load balancer to perform health checks and forward requests.

## User-data process

The strict-mode Bash script installs Go with `dnf`, creates the restricted service user, builds and installs the binary, installs the service unit, reloads `systemd`, and enables the service. It downloads no executable script from an untrusted source and embeds no credentials.

## Verification

- All four required files are present.
- `gofmt` produced no remaining changes.
- `go test` completed successfully as a compile check.
- `bash -n` exited `0`.
- `systemd-analyze` was unavailable on the macOS validation host; the unit was inspected statically for every required setting.
- The README documents cloud-init logs, `journalctl`, the health check, IAM instance roles, and secret handling.
- The sensitive-data scan was clean.
- No AWS resources were created.

## References

- [EC2 user data](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/user-data.html)
- [Create an EC2 launch template](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/create-launch-template.html)
