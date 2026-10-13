# Day 17 — Diagnose and Repair an Application Health-Check Failure

**Date:** October 13, 2026
**Due:** 11:00 PM Africa/Lagos
**Timebox:** 60–75 minutes
**Status:** Assigned

## Outcome

Use your Go service locally to distinguish a working application from an incorrect health-check path. Produce a failure table and repair evidence without rebuilding AWS resources.

## Steps

1. Create labs/week-04/day-17-health-check-failure/README.md. Read the two references below for 10 minutes.
2. Inspect your existing Day 14 main.go. Confirm its handler path and listening port. Check port 8080 is free with lsof -nP -iTCP:8080 -sTCP:LISTEN; if occupied, identify the process before stopping anything.
3. In a Mac terminal, from the project root run go run labs/week-04/day-14-ec2-user-data/main.go. Keep this terminal open.
4. In another terminal run curl -sS -o /dev/null -w '%{http_code}\n' http://127.0.0.1:8080/healthz. Record the status, then curl -fsS http://127.0.0.1:8080/healthz and record the response.
5. Simulate a misconfigured health-check path: curl -sS -o /dev/null -w '%{http_code}\n' http://127.0.0.1:8080/wrong-healthz. Record the result. Confirm the service is still running and /healthz still succeeds.
6. Repair the probe by changing the requested path back to /healthz; verify HTTP 200 and ok. This simulates correcting an ALB probe configuration locally; it does not test an actual ALB or AWS replacement.
7. Write a five-row failure table: terminated instance, incorrect health-check path, blocked backend port, failed deployment, unavailable database. For each include symptom, evidence, detector, automatic recovery if any, and human repair. Explain why replacing instances cannot fix a shared incorrect health-check path.
8. Stop the local Go service with Ctrl+C. Confirm lsof reports no lab listener on port 8080. Save only summarized results and explanations in README.md.

## Learning links

- [ALB target health checks](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/target-group-health-checks.html)
- [Auto Scaling health checks](https://docs.aws.amazon.com/autoscaling/ec2/userguide/health-checks-overview.html)

## Cost and security

No AWS resources required; AWS cost zero for this exercise. Use local requests only. The existing service binds :8080; work on a trusted network, keep the session short, and stop it after testing. Publish no deployment DNS names, AWS identifiers, raw logs, or credentials.

## Submission

1. README path:
2. Correct-path HTTP status and response:
3. Incorrect-path HTTP status:
4. Was the application still running during the failed probe?:
5. Repaired probe HTTP status and response:
6. Failure-table row count:
7. Why repeated replacements would not fix the incorrect path:
8. Local service stopped and port released:
9. AWS resources created:
10. Blocker:

## Pass criteria

- Correct path returns 200 and ok.
- Incorrect path returns 404 while the working endpoint remains available.
- Corrected probe returns 200 and ok again.
- Five failure rows distinguish detection, recovery, and human repair accurately.
- Explanation identifies configuration shared across replacement instances as the cause of repeated failure.
- Local service stopped; no AWS resources created.
