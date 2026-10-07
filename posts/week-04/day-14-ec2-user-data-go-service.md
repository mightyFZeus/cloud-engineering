# Day 14 Build-in-Public Draft — EC2 User Data and Go

## LinkedIn

Day 14 of my AWS cloud engineering journey focused on turning a working Go service into a repeatable EC2 boot process.

I packaged a small `/healthz` service with an EC2 user-data script and a hardened `systemd` unit. The service listens on port 8080 for an Application Load Balancer target and includes HTTP read-header, read, write, and idle timeouts.

The service runs under a dedicated system user and restarts after failures. Its unit also enables `NoNewPrivileges`, a private temporary directory, home-directory protection, and strict filesystem protection.

The user-data script uses strict Bash error handling, installs Go through the Amazon Linux package manager, builds and installs the binary, reloads `systemd`, and enables the service. I validated the Go code and Bash syntax locally and documented how to inspect cloud-init and service logs on Amazon Linux.

For networking, the EC2 security group should accept port 8080 only from the load balancer security group. AWS access should come from an IAM instance role, and credentials or secrets should never be embedded in user data.

No AWS resources were created for this packaging lab.

#AWS #CloudEngineering #Go #EC2 #Linux #Systemd

## X / Threads

AWS Day 14: packaged a Go `/healthz` service as EC2 user data with strict Bash handling and a hardened `systemd` unit. The app listens on port 8080 for an ALB, while the instance SG should trust only the ALB SG. Local Go and Bash checks passed; no AWS resources created. #AWS #Go
