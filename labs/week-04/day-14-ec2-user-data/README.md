cat > labs/week-04/day-14-ec2-user-data/README.md <<'MARKDOWN'
# Day 14 — EC2 User Data Go Service

## Outcome

This lab packages a Go health service with an EC2 user-data script and a hardened `systemd` service unit.

No AWS resources were created during this local validation.

## Health service

The Go application provides:

- `GET /healthz`
- HTTP status `200`
- Response body `ok`
- Listen address `:8080`
- Read-header, read, write, and idle timeouts

The service listens on `:8080`, which means it accepts connections through the instance's available network interfaces. This allows an Application Load Balancer to reach the service through the instance's private network interface.

## Security-group rule

The EC2 instance security group should allow inbound TCP traffic on port `8080` only from the Application Load Balancer's security group.

Using the load-balancer security group as the source prevents arbitrary internet clients from connecting directly to the application port. Clients should communicate with the load balancer over HTTPS, and the load balancer should forward accepted requests to the application.

## systemd service

The service runs as the dedicated `cloudeng` system user instead of `root`.

The unit uses:

- `Restart=on-failure`
- `NoNewPrivileges=true`
- `PrivateTmp=true`
- `ProtectHome=true`
- `ProtectSystem=strict`

It starts after `network-online.target` and is enabled for the normal multi-user boot target.

## User-data process

The user-data script:

1. Enables strict Bash error handling.
2. Defines an explicit root-owned Go build cache because cloud-init does not guarantee that `HOME` is set.
3. Installs Go through the Amazon Linux package manager.
4. Creates the restricted `cloudeng` system user.
5. Builds the Go application.
6. Installs the binary under `/usr/local/bin`.
7. Installs the service unit.
8. Reloads `systemd`.
9. Enables and starts the service.

## Troubleshooting

Inspect EC2 user-data execution with:

```bash
sudo less /var/log/cloud-init-output.log
sudo less /var/log/cloud-init.log
```

## Service logs and health check

Inspect the `systemd` service and its logs on Amazon Linux with:

- `systemctl status cloud-eng-health`
- `sudo journalctl -u cloud-eng-health --no-pager`

Test the health endpoint from the instance with:

- `curl -fsS http://127.0.0.1:8080/healthz`

The expected response is `ok`.

## Secret handling

EC2 user data must not contain AWS credentials, passwords, tokens, private keys, or other secrets because user data can be retrieved from the instance and may appear in logs.

AWS access should use an IAM instance role with temporary credentials. Application secrets should be retrieved at runtime from an appropriate managed service with narrowly scoped permissions.
