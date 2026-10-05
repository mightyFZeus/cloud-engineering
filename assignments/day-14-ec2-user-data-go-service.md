# Day 14 — Package the Go Health Service for EC2 User Data

**Date:** Thursday, October 8, 2026

**Due:** 11:00 PM Africa/Lagos

**Timebox:** 60–90 minutes

## Outcome

Create and locally validate an EC2 user-data deployment bundle that installs the Day 9 Go health service as a hardened `systemd` service suitable for an Application Load Balancer target.

## Learn — 10–15 minutes maximum

- Review [Run commands when you launch an EC2 instance with user data](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/user-data.html).
- Review [Boot modes and user-data execution logs on Amazon Linux](https://docs.aws.amazon.com/linux/al2023/ug/boot-modes.html).

## Build — 45–60 minutes

1. Create `labs/week-04/day-14-ec2-user-data/` containing:
   - `main.go`
   - `cloud-eng-health.service`
   - `user-data.sh`
   - `README.md`
2. In `main.go`, implement:
   - `GET /healthz` returning HTTP `200` and `ok`;
   - a server listening on `:8080`, so an ALB can reach it through a security-group-controlled private address; and
   - sensible HTTP read, write, and idle timeouts.
3. In the service unit:
   - use a dedicated `cloudeng` system user;
   - start `/usr/local/bin/cloud-eng-health`;
   - set `Restart=on-failure`;
   - start after `network-online.target`; and
   - enable `NoNewPrivileges`, `PrivateTmp`, `ProtectHome`, and `ProtectSystem` hardening.
4. In `user-data.sh`:
   - start with `#!/bin/bash` and `set -euo pipefail`;
   - install Go through `dnf`;
   - create the dedicated system user if absent;
   - write or install the Go source and service unit without downloading from an untrusted URL;
   - build and install the binary;
   - run `systemctl daemon-reload`; and
   - enable and start the service.
5. In the README, explain:
   - why the service listens on all instance interfaces for an ALB target;
   - why the instance security group should accept port `8080` only from the ALB security group;
   - where to inspect user-data output and service logs; and
   - why user data must not contain credentials or secrets.
6. Run the local validation commands:

   ```bash
   gofmt -w labs/week-04/day-14-ec2-user-data/main.go
   go test labs/week-04/day-14-ec2-user-data/main.go
   bash -n labs/week-04/day-14-ec2-user-data/user-data.sh
   systemd-analyze verify labs/week-04/day-14-ec2-user-data/cloud-eng-health.service
   ```

   If `systemd-analyze` is unavailable on macOS, record that as `not available locally`; the other three checks must pass.

## Safety and cost rules

- Create no AWS resources for this assignment.
- Do not place AWS credentials, account IDs, resource IDs, private keys, tokens, or secrets in user data or documentation.
- Do not use `curl | bash` or download executable code from an untrusted URL.

## Submit for verification

```text
Day 14 submission

1. Lab directory:
2. Four required files present: yes/no
3. Health endpoint and listen address:
4. Dedicated service user; restart policy:
5. Hardening settings enabled:
6. gofmt completed: yes/no
7. go test exit status:
8. bash -n exit status:
9. systemd-analyze result:
10. Explain the ALB-to-instance security-group rule:
11. AWS resources created:
12. Exact blocker, if any:
```

## Pass criteria

- [ ] All four required files exist and contain the required configuration.
- [ ] The health service listens on `:8080`, returns `200 ok`, and uses HTTP timeouts.
- [ ] The service unit uses the dedicated user, restart policy, network ordering, and all four hardening settings.
- [ ] User data uses strict Bash mode, installs dependencies, builds the binary, and enables the service without embedding secrets.
- [ ] `gofmt`, `go test`, and `bash -n` pass; the `systemd-analyze` result or macOS limitation is recorded.
- [ ] The README accurately explains ALB security-group scoping, logs, and secret handling.
