# Milestone 1 Build-in-Public Draft — Safe AWS Operator

## LinkedIn

My first AWS cloud engineering milestone is technically complete: I can launch, inspect, secure, troubleshoot, and clean up a small EC2 workload using temporary credentials.

Across the foundation labs, I configured safer account access, worked with assumed roles, narrowed IAM permissions, and automated baseline checks. I then launched Amazon Linux through Session Manager with no SSH key or inbound rule.

On that host, I ran a Go health endpoint as a `systemd` service. I verified the process, listening socket, HTTP response, and logs separately, then killed the process and confirmed that `systemd` restored it.

The final lab gave the instance role access to one private S3 bucket. I verified Block Public Access and encryption, tested `cp`, `sync`, and file integrity, and removed the instance, storage, bucket, objects, and temporary permission.

The biggest lesson is that operational confidence comes from specific evidence: identity, policy scope, process state, network listener, application response, logs, and cleanup each require their own check.

Next, I’m moving into CIDR planning, VPC routing, and network failure diagnosis.

#AWS #CloudEngineering #Linux #IAM #S3

## X

AWS Milestone 1: operated a short-lived Linux EC2 workload through Session Manager, ran and recovered a Go `systemd` service, used a scoped instance role with private S3, and verified cleanup. Next: CIDR planning and VPC routing. #AWS #CloudEngineering
