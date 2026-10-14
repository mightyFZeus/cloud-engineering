
## Failure diagnosis

The probe results below were observed locally. The table combines the Day 16 termination test with analysis of other failure scenarios.

| Failure | Symptom | Evidence to inspect | Detector | Automatic recovery | Human repair |
|---|---|---|---|---|---|
| Terminated instance | Group loses one running instance | EC2 state and ASG activity | Auto Scaling | Launches a replacement to restore desired capacity | Investigate unexpected termination |
| Incorrect health-check path | Probe returns 404 although the application runs | Probe path, HTTP status, and working endpoint | ALB target health checks | With ELB checks enabled, ASG may replace unhealthy instances; replacements still fail the same incorrect probe | Correct the target group's health-check path |
| Blocked backend port 8080 | ALB cannot reach the application | Target health reason, security-group rules, and listening socket | ALB target health checks | ASG may replace unhealthy instances, but a shared incorrect security rule remains | Allow TCP 8080 from the ALB security group |
| Failed deployment | Service fails to start or returns errors | Service status, application logs, and deployment version | systemd; ALB checks if the health endpoint fails | systemd may restart the process; ASG may replace unhealthy instances, but a shared broken deployment persists | Repair or roll back the deployment |
| Unavailable database | Requests needing the database fail | Application errors, database health, and connectivity | Application monitoring; ALB only if its health endpoint checks database readiness | Depends on database recovery or failover configuration; replacing application instances does not guarantee recovery | Diagnose database availability, connectivity, or credentials |

## Observed local probe results

- Correct probe: HTTP 200, ok.
- Incorrect probe: HTTP 404.
- Application remained running: yes.
- Repaired probe: HTTP 200, ok.

## Why replacements do not fix the incorrect path

The health-check path belongs to the shared probe configuration. Replacing an instance does not correct that path, so each replacement can run successfully while still failing the same incorrect probe.

