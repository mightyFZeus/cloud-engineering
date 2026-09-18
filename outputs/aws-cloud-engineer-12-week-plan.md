# AWS Cloud Engineer — 12-Week Study Plan

**Start:** Monday, September 14, 2026  
**Finish:** Sunday, December 6, 2026  
**Capacity:** 10 hours per week (about 120 hours total)  
**Starting point:** 5 years of React Native, 1.5 years of Go, comfortable but not expert with Linux  
**Target:** Build a practical AWS cloud engineering foundation and credible, job-relevant portfolio evidence.

## What success looks like after 12 weeks

This is not enough time to master all of AWS. It is enough time to become effective with the core workflow and prove it publicly. By December 6, you should be able to:

- Explain and use IAM roles, policies, temporary credentials, and least privilege.
- Build and troubleshoot VPCs, subnets, routes, security groups, DNS, and load-balanced services.
- Operate EC2, S3, DynamoDB, Lambda, API Gateway, ECR, ECS/Fargate, CloudWatch, CloudTrail, and Systems Manager at a practical level.
- Deploy and troubleshoot a Go workload on local Kubernetes and a short-lived Amazon EKS cluster.
- Provision repeatable AWS infrastructure with Terraform.
- Deploy a Go workload through GitHub Actions using OIDC instead of long-lived AWS keys.
- Monitor a workload, respond to an alarm, diagnose a deliberately introduced failure, and follow a runbook.
- Present at least one polished capstone with a diagram, deployment instructions, security/cost decisions, tests, and a short demo.

## Weekly rhythm (10 hours)

- **Monday — 2h:** concepts and guided reading/video
- **Wednesday — 2h:** focused lab
- **Friday — 2h:** focused lab or debugging drill
- **Saturday — 4h:** integrate the week's work, document it, clean up AWS resources, and review

Default allocation: **2h learning + 6h hands-on + 1h documentation + 1h recall/review**. If a course is long, watch only the modules named in this plan. Do not substitute passive watching for the hands-on deliverable.

## Non-negotiable guardrails

- [ ] Never create root-user access keys; protect the root user with MFA.
- [ ] Prefer temporary credentials, roles, and federation over long-lived IAM access keys.
- [ ] Create both a zero-spend alert and a small monthly cost budget before the first lab.
- [ ] Add `Project`, `Environment`, and `Owner` tags to resources where supported.
- [ ] Never commit `.env`, credentials, private keys, Terraform state, or plan files.
- [ ] End every hands-on session by stopping or destroying resources that accrue hourly charges.
- [ ] Treat NAT gateways, load balancers, public IPv4 addresses, RDS, and continuously running Fargate tasks as paid resources. Create them only for a planned lab window, then destroy them.

Useful setup references: [IAM security best practices](https://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html), [root-user best practices](https://docs.aws.amazon.com/IAM/latest/UserGuide/root-user-best-practices.html), and [AWS budget templates](https://docs.aws.amazon.com/cost-management/latest/userguide/budget-templates.html).

---

## Week 1 — Sep 14–20: AWS foundation and a safe account

**Outcome:** You can navigate AWS safely, explain its global structure and responsibility model, and use the CLI without treating the root user as an everyday identity.

### Learn (2h)

- Regions, Availability Zones, edge locations, shared responsibility, elasticity, high availability, and basic pricing.
- IAM identities, policies, roles, STS, MFA, and temporary versus long-lived credentials.
- Watch/read selected modules from [AWS Cloud Practitioner Essentials](https://aws.amazon.com/training/learn-about/cloud-practitioner/): introduction, compute, global infrastructure, security, pricing, and Well-Architected.

### Build (6h)

- Secure the root user with MFA and verify there are no root access keys.
- Set up an administrative identity for normal work. Use IAM Identity Center/temporary credentials if the account is already in AWS Organizations or you are not relying on new-account Free Tier credits. Otherwise, use a separate administrative IAM user with MFA temporarily; do not create access keys.
- Install AWS CLI v2 (2.32.0 or newer); use `aws login` with the non-root console identity to create an `aws-learning` named profile backed by temporary credentials, with `eu-west-1` as the default Region. Do not create a long-lived IAM access key.
- Run `aws sts get-caller-identity`, list Regions, and inspect the caller's identity.
- Create a zero-spend budget and a small monthly budget with actual and forecast alerts.
- Create a public repository named `aws-cloud-engineering-journey` using the documentation structure below.

### Document and test (2h)

- Draw a one-page Region/AZ/account/IAM mental model.
- Write a 150-word explanation of shared responsibility using an EC2 example and a Lambda example.
- Record five AWS CLI commands you used and what each proved.

### Done checklist

- [x] Root MFA enabled and no root access keys exist.
- [x] Daily work uses a non-root identity and `aws sts get-caller-identity` succeeds.
- [x] Zero-spend and monthly budget alerts are active.
- [ ] Learning repository and Week 1 log exist.
- [x] I can explain Region versus AZ and AWS versus customer responsibility without notes.
- [ ] I can explain role versus user without notes.

---

## Week 2 — Sep 21–27: Linux operations, EC2, EBS, and S3

**Outcome:** You can operate a Linux VM, inspect a service, move data with S3, and cleanly remove the environment.

### Learn (2h)

- EC2 instance lifecycle, AMIs, instance types, EBS, security groups, instance profiles, user data, and connection choices.
- Refresh Linux: permissions, users/groups, processes, ports, filesystems, `systemd`, `journalctl`, `curl`, `ss`, `top`, `df`, and `du`.
- Use the [EC2 getting-started tutorial](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EC2_GetStarted.html) and compare SSH with [Session Manager connection requirements](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/connect.html).

### Build (6h)

- Launch one small Amazon Linux instance with a project tag and an IAM instance role.
- Prefer Session Manager; if you use SSH, allow port 22 only from your current IP.
- Install and run a tiny Go HTTP health service under `systemd`.
- Use `journalctl`, `curl`, and `ss` to verify and troubleshoot it.
- Create a private S3 bucket; use `aws s3 cp` and `aws s3 sync`; verify public access remains blocked.
- Stop/terminate the instance and delete the bucket after preserving code and notes.

### Document and test (2h)

- Write a runbook: “service is running but endpoint is unreachable.”
- Capture the commands and evidence you would use to distinguish process, port, firewall/security-group, routing, and DNS problems.

### Done checklist — Milestone 1: safe AWS operator

- [ ] I can start, inspect, stop, and terminate EC2 deliberately.
- [ ] My Go service is managed by `systemd` and survives a process restart.
- [ ] I used an instance role for S3 access rather than embedding credentials.
- [ ] The S3 bucket was private and all billable lab resources were cleaned up.
- [ ] Week 2 runbook and learning log are committed.
- [ ] Milestone 1 build-in-public post was drafted and published or scheduled.

---

## Week 3 — Sep 28–Oct 4: Networking fundamentals and VPC

**Outcome:** You can reason about how traffic reaches—or fails to reach—an AWS workload.

### Learn (3h)

- IPv4, CIDR/subnetting, TCP/UDP, ports, DNS, public versus private IP addresses, and stateful versus stateless filtering.
- VPCs, public/private subnets, route tables, internet gateways, NAT, VPC endpoints, security groups, and network ACLs.
- Read [What is Amazon VPC?](https://docs.aws.amazon.com/vpc/latest/userguide/what-is-amazon-vpc.html), [VPC basics](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-subnet-basics.html), and scan the [VPC examples](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-examples-intro.html).

### Build (5h)

- On paper first, subnet `10.20.0.0/16` into at least four non-overlapping `/24` networks across two AZs.
- In the console or CLI, create a VPC, two public subnets, two private subnets, route tables, and an internet gateway.
- Do **not** leave a NAT gateway running. Learn the design, but only create one if you can destroy it in the same session.
- Launch a temporary test instance only if needed; deliberately break and repair one route or security-group rule.

### Document and test (2h)

- Draw the traffic path from a phone to an ALB to a private workload and back.
- Answer: What makes a subnet public? Why does a private subnet need NAT for IPv4 egress? When can an S3 gateway endpoint avoid NAT?

### Done checklist

- [ ] Four valid, non-overlapping subnet CIDRs are documented.
- [ ] The diagram labels AZs, routes, IGW, optional NAT, and security boundaries.
- [ ] I can explain security groups versus NACLs and public versus private subnets.
- [ ] I broke and repaired connectivity using evidence rather than random changes.
- [ ] All temporary resources were destroyed.

---

## Week 4 — Oct 5–11: Compute, load balancing, scaling, and availability

**Outcome:** You can deploy a small highly available service and explain its failure behavior.

### Learn (2h)

- Launch templates, Auto Scaling groups, Application Load Balancers, target groups, health checks, horizontal scaling, and Multi-AZ thinking.
- Compare EC2, Lambda, and Fargate: control, operational load, startup behavior, scaling, and cost.
- Use the architecture and cleanup sections of AWS's [public/private VPC production tutorial](https://docs.aws.amazon.com/vpc/latest/userguide/create-a-vpc-with-private-subnets-and-nat-gateways-using-aws-cli.html). Do not copy its always-on footprint blindly; the page explicitly identifies paid components.

### Build (6h)

- Package the Week 2 Go health service with user data or an AMI/launch template.
- Deploy two instances across two AZs behind an ALB and an Auto Scaling group.
- Verify health checks, terminate one instance, and observe replacement/recovery.
- Change a health-check path or security-group rule to create a failure, then repair it.
- Destroy the ALB, Auto Scaling group, instances, and any NAT gateways the same day.

### Document and test (2h)

- Create a failure table: instance dies, AZ impaired, health endpoint fails, bad deployment, database unavailable.
- For each failure, state what detects it, what recovers automatically, and what still needs a human.

### Done checklist — Milestone 2: networking and resilient compute

- [ ] The service ran behind an ALB in two AZs.
- [ ] A terminated instance was automatically replaced.
- [ ] I can trace a request through DNS/ALB/security groups/routes/instance/port.
- [ ] A failure table and architecture diagram are committed.
- [ ] All paid lab resources were destroyed.
- [ ] Milestone 2 build-in-public post was drafted and published or scheduled.

---

## Week 5 — Oct 12–18: IAM, encryption, audit, and secrets

**Outcome:** You can grant narrowly scoped access, explain an authorization failure, and find evidence of an AWS API action.

### Learn (3h)

- IAM policy evaluation, identity versus resource policies, explicit deny, role trust policies, STS, and least privilege.
- KMS concepts, encryption at rest/in transit, Secrets Manager versus Parameter Store, CloudTrail, and S3 security.
- Read [IAM security best practices](https://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html) and [S3 security best practices](https://docs.aws.amazon.com/AmazonS3/latest/userguide/security-best-practices.html).

### Build (5h)

- Create a role that can read only one prefix in one S3 bucket; test allowed and denied actions.
- Use a Go program with [AWS SDK for Go v2](https://docs.aws.amazon.com/sdk-for-go/) to call S3 or DynamoDB with the active role/profile.
- Store one non-secret setting in Parameter Store. If you test Secrets Manager, delete the secret immediately after the lab and understand its recovery/deletion behavior.
- Find your API actions in CloudTrail event history.
- Run IAM policy validation/Access Analyzer where available and reduce an over-broad policy.

### Document and test (2h)

- Write one IAM troubleshooting note from an `AccessDenied` message: principal, action, resource, context/conditions, and relevant policies.
- Add a threat model with assets, actors, entry points, and three mitigations.

### Done checklist

- [ ] One least-privilege policy permits an intended action and blocks an unintended one.
- [ ] My Go code uses the default credential chain, not hard-coded keys.
- [ ] I found a lab action in CloudTrail.
- [ ] Threat model and IAM failure analysis are committed.

---

## Week 6 — Oct 19–25: Storage, databases, backup, and recovery

**Outcome:** You can choose between object, block, relational, and NoSQL storage based on workload needs.

### Learn (3h)

- S3 versioning, lifecycle, encryption, pre-signed URLs, replication concepts, and consistency.
- EBS snapshots; RDS basics, backups, Multi-AZ versus read replicas; DynamoDB partitions, keys, indexes, consistency, capacity modes, TTL, and point-in-time recovery.
- Work through the concepts and first operations in [Getting started with DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GettingStartedDynamoDB.html).

### Build (5h)

- Design an access pattern first, then create a small DynamoDB table in on-demand mode.
- Build Go create/read/update/delete and query operations using AWS SDK for Go v2.
- Create pre-signed S3 upload/download URLs from Go; keep the bucket private.
- Enable and test one recovery feature such as S3 versioning, then restore a deleted/overwritten object.
- Delete test data and resources after documenting the result.

### Document and test (2h)

- Make a decision record comparing S3, EBS, RDS, and DynamoDB for four concrete scenarios.
- State RPO and RTO for your mini application and explain the chosen backup method.

### Done checklist

- [ ] DynamoDB keys were derived from access patterns, not chosen after coding.
- [ ] Go CRUD/query operations work without embedded credentials.
- [ ] A private S3 object was accessed through a time-limited pre-signed URL.
- [ ] I successfully restored one item/object in a recovery exercise.
- [ ] Storage decision record is committed.

---

## Week 7 — Oct 26–Nov 1: Serverless and event-driven systems with Go

**Outcome:** You can build and operate a small event-driven API rather than only deploying servers.

### Learn (2h)

- Lambda execution model, concurrency, timeouts, retries, idempotency, cold starts, and execution roles.
- API Gateway, SQS, SNS, EventBridge, dead-letter queues, and synchronous versus asynchronous work.
- Start with AWS's [Go Lambda samples](https://docs.aws.amazon.com/lambda/latest/dg/lambda-samples.html), [API Gateway/Lambda/DynamoDB tutorial](https://docs.aws.amazon.com/lambda/latest/dg/services-apigateway-tutorial.html), and [S3-triggered Lambda tutorial](https://docs.aws.amazon.com/lambda/latest/dg/with-s3-example.html).

### Build (7h)

- Build a Go Lambda API behind API Gateway with DynamoDB persistence.
- Add input validation, structured logs, correlation/request IDs, and a least-privilege execution role.
- Add one asynchronous path: API or S3 event → SQS → Go Lambda worker; configure retry behavior and a DLQ.
- Force one failure, inspect CloudWatch logs/DLQ, fix the cause, and redrive or replay the message.

### Document and test (1h)

- Update the architecture diagram and record why the async boundary exists.
- Write tests for the handler/business logic and document one failure/recovery sequence.

### Done checklist — Milestone 3: secure serverless workload

- [ ] Go API works through API Gateway and persists data in DynamoDB.
- [ ] An asynchronous event reaches a Go worker through SQS.
- [ ] A failed message reached the DLQ and was diagnosed/reprocessed.
- [ ] Logs carry a request/correlation ID.
- [ ] Tests, diagram, security notes, and cleanup steps are committed.
- [ ] Milestone 3 build-in-public post was drafted and published or scheduled.

---

## Week 8 — Nov 2–8: Infrastructure as code with Terraform

**Outcome:** You can create, change, review, and destroy AWS infrastructure from versioned code.

### Learn (3h)

- Terraform providers, resources, data sources, variables, locals, outputs, modules, dependency graph, state, drift, and the `init/validate/plan/apply/destroy` workflow.
- Complete the core [HashiCorp AWS getting-started tutorials](https://developer.hashicorp.com/terraform/tutorials/aws-get-started).
- Optional 45-minute comparison: AWS CDK also supports [infrastructure definitions in Go](https://docs.aws.amazon.com/cdk/v2/guide/develop.html), but Terraform is the primary IaC tool in this plan.

### Build (6h)

- Rebuild the Week 7 serverless stack or a reduced VPC stack in Terraform.
- Use variables, outputs, tags, provider/version constraints, formatting, and a small reusable module.
- Add `.gitignore` rules for state, plan files, local variables, and credentials.
- Run a plan before every apply. Change one resource, inspect the plan, apply, detect one manual drift, and reconcile it.
- Destroy and recreate the stack from code.

### Document and test (1h)

- Document state risks and how a team would use protected remote state and locking.
- Save a redacted plan excerpt and explain create/update/replace/destroy indicators.

### Done checklist

- [ ] `terraform fmt` and `terraform validate` pass.
- [ ] A reviewed plan creates the expected resources.
- [ ] The stack can be destroyed and recreated from code.
- [ ] No state, plan, secret, or credential file is committed.
- [ ] I detected and reconciled one deliberate drift.

---

## Week 9 — Nov 9–15: Containers and Kubernetes fundamentals

**Outcome:** You can package a Go service, deploy it to a local Kubernetes cluster, update it safely, and troubleshoot common workload failures.

### Learn (3h)

- Images, containers, registries, layers, Dockerfiles, Compose, ports, health checks, and container security basics.
- Kubernetes control plane and worker nodes; Pods, Deployments, ReplicaSets, Services, namespaces, labels/selectors, ConfigMaps, Secrets, probes, resource requests/limits, rollouts, and desired state.
- Use Docker's [build and share a containerized application](https://docs.docker.com/get-started/tutorials/run-an-app/) and the official [Kubernetes Basics tutorial](https://kubernetes.io/docs/tutorials/kubernetes-basics/).

### Build (6h)

- Write a multi-stage Dockerfile for a Go API, run as a non-root user, and add a health endpoint.
- Run and test the image locally; scan dependencies/image with an available scanner.
- Create a local cluster with `kind` or `minikube`; verify it with `kubectl cluster-info` and `kubectl get nodes`.
- Write manifests for a Namespace, Deployment, ClusterIP Service, and ConfigMap. Add readiness/liveness probes and CPU/memory requests and limits.
- Scale the Deployment, publish a new image tag, perform a rolling update, inspect rollout history, and execute a rollback.
- Deliberately cause an `ImagePullBackOff`, `CrashLoopBackOff`, or failing readiness probe. Diagnose with `kubectl get`, `describe`, `logs`, and `events`, then repair it.
- Delete the local cluster after preserving manifests and evidence.

### Document and test (1h)

- Draw Pod → Deployment → Service traffic and ownership relationships.
- Compare EC2, ECS/Fargate, Kubernetes/EKS, and Lambda for the same Go API.

### Done checklist

- [ ] Go image builds reproducibly with a multi-stage Dockerfile and non-root runtime.
- [ ] Go API runs on a local Kubernetes Deployment and is reachable through a Service/port-forward.
- [ ] Probes and CPU/memory requests and limits are configured.
- [ ] I completed a scale, rolling update, and rollback.
- [ ] I diagnosed one Kubernetes failure from events, `describe`, and logs.
- [ ] Manifests, diagram, troubleshooting note, and cleanup steps are committed.

---

## Week 10 — Nov 16–22: Amazon EKS and secure delivery

**Outcome:** Your Go workload runs on a short-lived Amazon EKS cluster, uses AWS identity safely, and its image is built and pushed without long-lived AWS keys in GitHub.

### Learn (3h)

- CI versus CD, artifacts, immutable images, environment promotion, rollback, approvals, and pipeline security.
- EKS control plane/data plane, managed node groups versus Fargate/Auto Mode, VPC CNI, Kubernetes RBAC versus AWS IAM, access entries, ECR, and EKS Pod Identity.
- Learn how [GitHub Actions uses OIDC with AWS](https://docs.github.com/en/actions/how-tos/secure-your-work/security-harden-deployments/oidc-in-aws) to avoid long-lived cloud credentials.
- Review [EKS setup and local-practice options](https://docs.aws.amazon.com/eks/latest/userguide/setting-up.html) and [EKS Pod Identity](https://docs.aws.amazon.com/eks/latest/userguide/pod-identities.html).

### Build (6h)

- Create a GitHub Actions workflow for `go test`, lint/static analysis, and container build.
- Configure a narrowly scoped AWS OIDC trust role restricted to your repository/branch or environment.
- On a protected event, push a versioned image to a private ECR repository.
- Create one short-lived EKS learning cluster using an AWS getting-started path. Use a small managed node group or another explicitly chosen compute mode; record the cost model before creation.
- Connect with `kubectl`, deploy the Week 9 manifests using the ECR image, and verify the application with port-forwarding or another intentionally chosen exposure method.
- Configure one Kubernetes service account and EKS Pod Identity for narrowly scoped access to one AWS resource, using the Go SDK's default credential chain.
- Destroy the cluster, nodes, load balancers, volumes, and other paid resources immediately after verification.
- **Stretch only:** let GitHub Actions deploy manifests to EKS through OIDC and restricted cluster access. The core requirement is secure CI image delivery plus a verified manual EKS deployment.

### Document and test (1h)

- Draw both identity flows: GitHub OIDC → AWS STS → ECR, and Kubernetes service account → EKS Pod Identity → AWS service.
- Record cluster creation/deletion timestamps, validation evidence, and the resources checked during cleanup.

### Done checklist — Milestone 4: Kubernetes on AWS and secure delivery

- [ ] Pull requests run Go and Terraform checks.
- [ ] GitHub Actions uses OIDC to push a versioned image to ECR and contains no long-lived AWS secret.
- [ ] The trust policy is restricted to the intended repository and ref/environment.
- [ ] The Go workload ran successfully on EKS from declarative Kubernetes manifests.
- [ ] A Kubernetes service account received narrowly scoped AWS access through EKS Pod Identity.
- [ ] I completed and documented a Kubernetes rollout/rollback.
- [ ] The EKS cluster and all paid supporting resources were destroyed and verified absent.
- [ ] Milestone 4 build-in-public post was drafted and published or scheduled.

---

## Week 11 — Nov 23–29: Observability, operations, incidents, and recovery

**Outcome:** You can detect, investigate, and recover from a realistic failure with evidence.

### Learn (3h)

- Metrics, logs, traces, dashboards, alarms, baselines, SLI/SLO thinking, CloudWatch, CloudTrail, Systems Manager, and audit versus application logs.
- Kubernetes operations: `kubectl logs`, events, `describe`, resource usage, probes, rollout status, and common states such as Pending, CrashLoopBackOff, and ImagePullBackOff.
- Scaling, backups, RTO/RPO, Multi-AZ and Multi-Region tradeoffs, and runbooks.
- Review the [AWS Well-Architected Framework](https://docs.aws.amazon.com/wellarchitected/latest/framework/welcome.html) and [CloudWatch alarm getting started](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/alarm-getting-started.html).

### Build (6h)

- Add structured application logs, custom/useful metrics, a CloudWatch dashboard, and at least two actionable alarms.
- Add Kubernetes workload signals to the runbook. Use local Kubernetes for repeated failure drills; use EKS/Container Insights only if its cost is understood and the cluster will be destroyed the same session.
- Add notification to an email/SNS endpoint if desired; verify and then avoid noisy alerts.
- Run a game day: introduce one application/configuration failure and one permissions/network failure.
- Diagnose with a timeline and evidence; recover using your runbook; measure time to detect and recover.
- Verify backup/restore or redeployment from Terraform after deleting a non-production resource.

### Document and test (1h)

- Write a blameless incident report: impact, timeline, detection, root cause, contributing factors, resolution, and prevention.

### Done checklist

- [ ] Dashboard answers traffic, errors, latency, and resource-health questions.
- [ ] Alarms are actionable and were tested.
- [ ] Two injected failures were diagnosed using logs/metrics/events.
- [ ] Restore or full redeployment was demonstrated.
- [ ] Incident report includes preventive actions, not just the immediate fix.

---

## Week 12 — Nov 30–Dec 6: Capstone hardening and career packaging

**Outcome:** One project is understandable, deployable, operable, and demonstrable by another engineer.

### Build and polish (7h)

- Choose one portfolio project below and finish its smallest credible version; do not attempt every stretch goal.
- Make deployment and destruction reproducible with Terraform and a `Makefile` or task runner.
- Add tests, CI, OIDC deployment, least-privilege roles, meaningful logs, a dashboard/alarm, and a rollback/recovery procedure.
- Run a Well-Architected mini-review covering operational excellence, security, reliability, performance, cost, and sustainability.
- Remove secrets and state, run dependency/secret scans, and verify a clean deployment from a fresh clone.

### Present and assess (3h)

- Write a strong README: problem, architecture diagram, tradeoffs, setup, validation, operations, security, cost, cleanup, and future work.
- Record a 3–5 minute demo showing deployment, a request/event, monitoring, one failure/recovery, and teardown.
- Write three resume bullets using action + AWS technology + measurable outcome.
- Review the current [AWS Certified CloudOps Engineer – Associate (SOA-C03) guide](https://docs.aws.amazon.com/aws-certification/latest/sysops-administrator-associate-03.html). Take an official practice question set/pretest if available to you, but book the exam only when your practice results are consistently strong and you have continued hands-on work.

### Done checklist — Milestone 5: portfolio-ready cloud engineering proof

- [ ] A fresh clone can be validated and deployed by following the README.
- [ ] Architecture, trust boundaries, request/event flow, and failure behavior are documented.
- [ ] CI passes and AWS deployment uses OIDC.
- [ ] Dashboard, alarms, runbook, incident report, and recovery evidence exist.
- [ ] Cost estimate and complete teardown instructions exist.
- [ ] Three-to-five-minute demo and three resume bullets are complete.
- [ ] No credentials, secrets, private keys, Terraform state, or sensitive identifiers are in Git history.
- [ ] Milestone 5 build-in-public retrospective was drafted and published or scheduled.

---

## Three portfolio projects

Build **one** to polished depth during the 12 weeks. Add the other two afterward. One excellent, operated project is more persuasive than three unfinished repositories.

### 1. Mobile Media Processing Pipeline — best fit for your background

**Story:** A React Native app securely uploads an image through a time-limited URL. AWS processes it asynchronously and the app displays status/metadata.

**Core architecture:** React Native → authentication → API Gateway → Go Lambda → pre-signed S3 upload → S3 event/SQS → Go Lambda worker → DynamoDB status/metadata; CloudWatch for logs, metrics, alarms; Terraform for infrastructure.

**Minimum credible version**

- [ ] Authentication or a clearly documented temporary demo-auth boundary.
- [ ] Private S3 bucket; client never receives AWS credentials.
- [ ] Go API issues short-lived pre-signed upload URLs.
- [ ] SQS decouples processing and a DLQ captures failures.
- [ ] DynamoDB status can be read by the mobile client.
- [ ] Terraform deploy/destroy, tests, CI, OIDC, dashboard, and alarm.
- [ ] Demo includes a poisoned message reaching the DLQ and successful recovery.

**Stretch:** thumbnails, CloudFront delivery, Cognito, object lifecycle, malware/content checks, or Step Functions. Add only one after the core is polished.

### 2. Production-style Go API on ECS/Fargate or Amazon EKS

**Story:** A containerized Go service runs across two AZs, deploys securely from GitHub, scales, emits useful telemetry, and can roll back.

**Core architecture:** GitHub Actions/OIDC → ECR → either ECS/Fargate or EKS → Go API; private application subnets; DynamoDB or RDS; CloudWatch; Terraform.

**Minimum credible version**

- [ ] Multi-stage, non-root Go container with health/readiness endpoint.
- [ ] Versioned ECR image and at least two workload replicas/tasks during the demo.
- [ ] If using EKS: declarative manifests, probes, requests/limits, Service, rolling update/rollback, and Pod Identity.
- [ ] If using ECS: task definition/service, health checks, and separate least-privilege task/execution roles.
- [ ] CI tests and OIDC deployment with a documented rollback.
- [ ] Dashboard/alarms plus one failure-injection incident report.
- [ ] Cost estimate and a `destroy` path; paid infrastructure is not left running.

**Stretch:** autoscaling, blue/green deployment, RDS backup/restore, WAF, HTTPS/custom domain, or load testing. Choose one.

### 3. Serverless Reliability Lab / Incident Simulator

**Story:** A deliberately fault-injectable serverless order or job system demonstrates operations—not just happy-path application code.

**Core architecture:** API Gateway → Go Lambda → DynamoDB; EventBridge/SQS → worker Lambda → DLQ; CloudWatch dashboard/alarms; CloudTrail audit; Terraform; GitHub Actions OIDC.

**Minimum credible version**

- [ ] Normal request and asynchronous job paths work end to end.
- [ ] Feature flags/config can inject timeout, permission, malformed-message, and throttling-style failures safely.
- [ ] Correlation IDs connect API, queue, worker, and logs.
- [ ] Alarms and runbooks exist for at least three failure modes.
- [ ] One backup/restore or complete rebuild exercise is recorded.
- [ ] Incident reports include detection and recovery times.

**Stretch:** X-Ray tracing, automated remediation, canary checks, multi-account design document, or AWS Config rules. Choose one.

## Portfolio quality checklist (apply to every finished project)

- [ ] The README begins with the problem and a working architecture diagram.
- [ ] The repo contains app code, IaC, tests, CI, deployment, rollback, and cleanup instructions.
- [ ] No screenshots are used as a substitute for reproducible infrastructure.
- [ ] Security decisions name the principal, action, resource, and trust boundary.
- [ ] Cost section identifies continuously billed resources and an estimated demo cost.
- [ ] Operations section shows signals, alarms, runbook, and a real failure/recovery exercise.
- [ ] Tradeoffs explain why services were selected and name at least one credible alternative.
- [ ] A short demo video proves the system, observability, recovery, and teardown.

## Documentation system

Suggested repository structure:

```text
aws-cloud-engineering-journey/
├── README.md
├── weeks/
│   ├── week-01-foundations/
│   │   ├── learning-log.md
│   │   ├── lab-notes.md
│   │   └── diagram.png
│   └── ...
├── projects/
│   └── chosen-capstone/
├── runbooks/
├── incidents/
├── decisions/
└── diagrams/
```

For each week, use the companion `aws-learning-log-template.md`. Keep notes compact and evidence-based. The valuable material is not a transcript of a course; it is what you built, how you verified it, what failed, what evidence led to the fix, and what tradeoff you made.

### Weekly publishing target

- One learning log (roughly 400–800 useful words).
- One diagram or clearly annotated request/traffic flow.
- One deployable lab or code/IaC change.
- One failure and troubleshooting entry.
- One short “explain it in an interview” answer.

## Build-in-public posting workflow

Every progress check-in has two outputs: the checklist is updated, and a social post is drafted from the evidence you provide. Use the companion `aws-build-in-public-post-template.md`.

### When to create a post

- **Progress post:** whenever you report one or more meaningful completed checklist items. Group small related items into one useful post instead of posting every command separately.
- **Milestone post:** after Weeks 2, 4, 7, 10, and 12. This is a fuller retrospective showing what you can now do.
- **Project/demo post:** whenever a public repository, architecture diagram, demo, incident report, or measurable result becomes available.

### What I will generate after your check-in

- A polished LinkedIn-style post of roughly 120–220 words.
- A short X/Threads-style version of roughly 280 characters when possible.
- A suggested title/hook, 3–5 restrained hashtags, and placeholders for a repository, diagram, or demo link.
- A privacy/security check that removes account IDs, resource identifiers, IP addresses, secrets, tokens, email addresses, and unredacted logs.

### Evidence to send with progress

Use this compact format:

```text
Week/Milestone:
Completed checkboxes:
What I built:
One problem or failure:
How I solved/verified it:
Evidence link or result:
Tone/platform preference (optional):
```

The post should be honest and specific. It should explain the problem, the engineering decision, evidence that it worked, and the next step. Avoid presenting a tutorial follow-along as production experience, inflating results, or listing AWS services without explaining what they accomplished.

## Certification direction

Do not spend the first month memorizing for an exam. Use the Cloud Practitioner material as a vocabulary course; your experience makes the foundational exam optional.

- **Closest role match:** AWS Certified CloudOps Engineer – Associate (SOA-C03), focused on operating, monitoring, securing, networking, and maintaining AWS workloads.
- **Broader first credential:** AWS Certified Solutions Architect – Associate (SAA-C03), useful if you want more architecture breadth.
- **Realistic timing:** the current CloudOps exam guide describes about a year of AWS operations experience as the target. Treat these 12 weeks as the practical foundation, continue operating the portfolio for another 1–3 months, and use consistent practice results—not a calendar date—to decide when to sit the exam.

## Progress dashboard

| Milestone | Due | Status |
|---|---:|---|
| 1. Safe AWS operator (Weeks 1–2) | Sep 27 | ⬜ Not started |
| 2. Networking and resilient compute (Weeks 3–4) | Oct 11 | ⬜ Not started |
| 3. Secure serverless workload (Weeks 5–7) | Nov 1 | ⬜ Not started |
| 4. IaC, Kubernetes/EKS, and secure delivery (Weeks 8–10) | Nov 22 | ⬜ Not started |
| 5. Operable, portfolio-ready capstone (Weeks 11–12) | Dec 6 | ⬜ Not started |

When reporting progress, use: **`Week N — completed: ...; blocked: ...; evidence: ...`**. The checklist/dashboard will be updated and a build-in-public post will be drafted from the completed work.
