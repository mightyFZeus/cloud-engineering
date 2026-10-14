# Day 17 — Health-check failure diagnosis

Passed October 14, 2026: six assignment criteria verified from learner checkpoints and the saved lab README.

## Observed evidence

- Go service started locally on :8080.
- /healthz returned HTTP 200 and ok.
- /wrong-healthz returned HTTP 404 while /healthz still returned ok and the process remained running.
- Correcting the requested path returned HTTP 200 and ok again; no code change was needed.
- Five failure-table rows documented termination, incorrect probe path, blocked backend port, failed deployment, and unavailable database.
- Learner explanation: replacement instances face the same wrong health-check path; repair the path instead.
- Service stopped, port 8080 released, README nonempty check exited 0, no AWS resources created.

## Scope and command summary

The HTTP path failure and repair were tested locally with curl. The table's other scenarios are analysis, not new AWS experiments. Day 16 separately verified actual instance termination and replacement. No ALB application-failure replacement was tested on Day 17.

Commands used: go run for the existing service, curl to inspect status/body, lsof to check the listener, Ctrl+C to stop the service, and test -s to check the README.

## References

- [ALB health checks](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/target-group-health-checks.html)
- [Auto Scaling health checks](https://docs.aws.amazon.com/autoscaling/ec2/userguide/health-checks-overview.html)
