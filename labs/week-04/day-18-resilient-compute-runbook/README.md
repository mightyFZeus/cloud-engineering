
# Resilient Compute Runbook

## Architecture used in the Day 16 lab

- Region: eu-west-1.
- Network: default VPC with two public subnets, in eu-west-1a and eu-west-1b.
- Public subnet route: internet-bound traffic routes through an Internet Gateway.
- Load balancer: internet-facing Application Load Balancer with an HTTP port 80 listener.
- ALB security group: inbound HTTP 80 allowed only from my current public address.
- Target group: instance targets using HTTP port 8080, with health-check path /healthz.
- Application security group: inbound TCP 8080 allowed only from the ALB security group; no SSH rule.
- Compute: two Amazon Linux 2023 t3.micro instances managed by an Auto Scaling group.
- Auto Scaling: desired/minimum/maximum capacity 2/2/2, with ELB health checks enabled.
- Instance role: AWSLearningEC2SSMRole; no key pair; IMDSv2 required.
- Startup: user data installs and starts the Go health service.

This lab used public subnets. It did not demonstrate a deployment with application instances in private subnets.


## Architecture sketch

```text
Client browser / curl
        |
        | HTTP 80
        | ALB security group permits my current address
        v
Internet-facing ALB
Enabled in public subnets across eu-west-1a and eu-west-1b
        |
        | Target group: HTTP 8080
        | Application security group permits the ALB security group
        |
        +---------------------------+
        |                           |
        v                           v
AZ: eu-west-1a                 AZ: eu-west-1b
Public subnet A                Public subnet B
Go service instance A          Go service instance B
        |                           |
        +-- /healthz -> 200, ok -----+

Both public subnet route tables:
0.0.0.0/0 -> Internet Gateway

Health-check traffic:
ALB -> each target on HTTP 8080, path /healthz

Replacement control:
Auto Scaling monitors instance health/state
        -> detects terminated or unhealthy instance
        -> launches replacement using the launch template
        -> registers replacement with the target group

Instance permissions:
Each instance -> IAM instance role -> authorized AWS API access
```

The ALB forwards application requests. Auto Scaling controls instance capacity and replacement. Replacement is a control action, separate from the request path.


## Request and response path

The client resolves the ALB DNS name and sends an HTTP request to its port 80 listener.
The internet-facing ALB uses public subnets with routes to an Internet Gateway, and its security group permits HTTP only from my current public address.
The listener forwards the request through the target group to a healthy Go service instance on port 8080.
The ALB reaches the instance using its private address, and the application security group permits port 8080 only from the ALB security group.
The Go service handles /healthz and returns HTTP 200 with ok.
The response travels back through the ALB to the client, with security groups automatically allowing response traffic for permitted connections because they are stateful.
A public subnet provides an internet route, but reachability also depends on addressing, security rules, and a listening application; the backend's port 8080 was restricted to the ALB.


## Failure diagnosis

| Failure | Symptom | Evidence to inspect | Detector | Automatic recovery | Human repair |
|---|---|---|---|---|---|
| Terminated instance | Group loses one running instance | EC2 state and ASG activity | Auto Scaling | Launches a replacement to restore desired capacity | Investigate unexpected termination |
| Incorrect health-check path | Probe returns 404 although the application runs | Probe path, HTTP status, and working endpoint | ALB target health checks | With ELB checks enabled, ASG may replace unhealthy instances; replacements still fail the same incorrect probe | Correct the target group's health-check path |
| Blocked backend port 8080 | ALB cannot reach the application | Target health reason, security-group rules, and listening socket | ALB target health checks | ASG may replace unhealthy instances, but a shared incorrect security rule remains | Allow TCP 8080 from the ALB security group |
| Failed deployment | Service fails to start or returns errors | Service status, application logs, and deployment version | systemd; ALB checks if the health endpoint fails | systemd may restart the process; ASG may replace unhealthy instances, but a shared broken deployment persists | Repair or roll back the deployment |
| Unavailable database | Requests needing the database fail | Application errors, database health, and connectivity | Application monitoring; ALB only if its health endpoint checks database readiness | Depends on database recovery or failover configuration; replacing application instances does not guarantee recovery | Diagnose database availability, connectivity, or credentials |
| Impaired Availability Zone | Targets in one AZ become unreachable; remaining capacity may be reduced | Target health by AZ, EC2 status, ASG activity, AWS Health, and request errors | ALB health checks, Auto Scaling health signals, and monitoring | ALB can route to healthy targets in the surviving AZ; ASG may restore capacity where launches are possible | Check surviving capacity, launch failures, and dependencies; adjust capacity or placement as needed. Recovery depends on available capacity and configuration; uninterrupted service is not guaranteed |

The impaired-AZ scenario is design analysis, not a tested outage. Day 16 tested instance termination and replacement. Day 17 tested an incorrect probe path locally.

## Ordered troubleshooting checklist

1. Check the client request: correct DNS name, HTTP scheme, port 80, and /healthz path. Confirm the client's current address is allowed by the ALB security group.
2. Check the ALB: Active state, intended security group attached, and HTTP 80 listener forwarding to the correct target group.
3. Check networking: both ALB subnets have routes to an Internet Gateway; inspect network ACLs for permitted request and return traffic.
4. Check target health: inspect registered targets, their Availability Zones, health status, and reason codes.
5. Check backend security rules: application inbound TCP 8080 from the ALB security group, and ALB outbound access to the targets.
6. Check the service: process running, systemd status on EC2, and a listening socket on port 8080.
7. Check the probe: correct /healthz path, health-check port, expected HTTP status, and timeout settings.
8. Check startup and application evidence: user-data/cloud-init errors, service logs, and application errors. Record summarized findings without publishing raw logs or identifiers.

## Day 16 browser timeout

Both ALB targets were healthy, and a direct HTTP curl request from my Mac returned ok.
This confirmed the health endpoint was reachable through the ALB from that client at the time of the test.
It narrowed the investigation toward differences between the browser request and curl, such as scheme, proxy settings, or client behavior.
It did not establish the browser timeout's root cause, so I do not claim that a particular browser or network setting caused it.


## Verified evidence

- [Day 13 traffic-path notes](../../week-03/day-13-vpc-traffic-path/README.md): earlier network-path analysis.
- [Day 16 deployment verification](../../../docs/week-04/day-16-alb-auto-scaling.md): two InService instances across two AZs, two healthy targets, and an ok health response.
- Day 16 deliberate instance termination caused Auto Scaling to launch a replacement, restoring two InService instances and two healthy targets.
- [Day 17 probe evidence](../day-17-health-check-failure/README.md): correct path returned 200/ok, incorrect path returned 404 while the service stayed running, and the corrected probe returned 200/ok.

## Evidence limits

Instance replacement was tested on AWS during Day 16.
The incorrect probe path was tested locally during Day 17.
The impaired-AZ, blocked-port, failed-deployment, and database scenarios are documented analysis rather than additional tested failures.
The browser timeout's root cause remains unconfirmed.

## Cleanup

- Day 16 instances terminated; Auto Scaling group, ALB, target group, launch template, and both temporary security groups deleted.
- Zero Day 16 EBS volumes remained.
- No NAT gateways or Elastic IPs were created.
- Day 17 local service stopped and port 8080 released.
- Day 18 created no new AWS resources.

