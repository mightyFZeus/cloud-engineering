# Day 13 — Trace and Diagnose VPC Traffic Paths

**Date:** Wednesday, October 7, 2026

**Due:** 11:00 PM Africa/Lagos

**Timebox:** 60–75 minutes

## Outcome

Create an architecture diagram and evidence-based troubleshooting table that trace traffic from a client to a load balancer and private application workload across two Availability Zones.

## Learn — 10–15 minutes maximum

- Review [Route tables](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html).
- Review [Security group rules](https://docs.aws.amazon.com/vpc/latest/userguide/security-group-rules.html), especially stateful response traffic and referencing one security group from another.

## Build — 40–50 minutes

1. Create `labs/week-03/day-13-vpc-traffic-path/README.md`.
2. Add a Mermaid flowchart showing:
   - a client on the internet;
   - DNS;
   - an internet gateway;
   - an internet-facing application load balancer in two public subnets across two Availability Zones;
   - an application workload in two private subnets;
   - the load-balancer and application security groups;
   - the public and private route-table behavior;
   - an optional NAT path for general outbound IPv4 traffic; and
   - an S3 gateway endpoint path that avoids NAT for S3 access.
3. Under the diagram, write the request and response path as numbered hops. Explain where routes apply and where security groups apply.
4. Add a troubleshooting table with these five failures:
   - the public route table is missing its internet-gateway default route;
   - the load-balancer security group does not allow HTTPS from clients;
   - the application security group does not allow the application port from the load-balancer security group;
   - a private workload tries to use an internet gateway directly without a public IPv4 address; and
   - a custom network ACL blocks required return traffic.
5. Give each failure four columns: symptom, evidence to inspect, root cause, and specific repair. Base each repair on the evidence rather than changing several controls at once.
6. Add a short comparison explaining when a NAT device is needed and when an S3 gateway endpoint can avoid it.
7. Run these local checks:

   ```bash
   test -s labs/week-03/day-13-vpc-traffic-path/README.md
   rg -n 'internet gateway|security group|route table|NAT|S3 gateway endpoint|network ACL' labs/week-03/day-13-vpc-traffic-path/README.md
   ```

## Safety and cost rules

- Create no AWS resources for this assignment.
- Do not include account IDs, resource IDs, ARNs, credentials, public IP addresses, or private console URLs.
- Label NAT as optional in the design; do not create a NAT gateway for this documentation exercise.

## Submit for verification

```text
Day 13 submission

1. README path:
2. Diagram contains two AZs, public/private subnets, IGW, ALB, workloads, route tables, and security groups: yes/no
3. NAT and S3 gateway endpoint paths shown: yes/no
4. Request and response path documented: yes/no
5. Failure-table row count:
6. One failure: symptom; evidence; root cause; repair:
7. test -s exit status:
8. rg check found all required concepts: yes/no
9. AWS resources created:
10. Exact blocker, if any:
```

## Pass criteria

- [x] The Mermaid diagram shows the required two-AZ request path and security boundaries.
- [x] The request and response paths correctly distinguish routing from security filtering.
- [x] NAT and S3 gateway endpoint paths are accurately explained.
- [x] The troubleshooting table contains all five required failures and evidence-based repairs.
- [x] Both local checks pass.
- [x] No AWS resources or sensitive identifiers were created or recorded.

## Verification — October 5, 2026

**Result:** Passed, 6/6 criteria.

- The README contains a two-AZ Mermaid diagram with public and private subnets, an internet gateway, load-balancer nodes, private workloads, route tables, and security groups.
- The numbered request path distinguishes route selection from stateful security-group filtering.
- Optional NAT and S3 gateway endpoint paths are shown and correctly explained.
- The troubleshooting table contains five required failures with symptoms, evidence, root causes, and specific repairs.
- `test -s` and the required `rg` check both exited `0`.
- No AWS resources or sensitive identifiers were created or recorded.
