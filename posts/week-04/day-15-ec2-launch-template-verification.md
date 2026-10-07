# Day 15 Build-in-Public Draft — Diagnosing EC2 User Data

## LinkedIn

Day 15 of my AWS cloud engineering journey turned into a useful first-boot troubleshooting exercise.

I created an EC2 launch template for an Amazon Linux 2023 `t3.micro` instance running my Go health service. The template required IMDSv2, used an IAM instance role for Systems Manager, used no key pair, and attached a security group that accepted port 8080 only from a placeholder load-balancer security group.

The first launch failed during cloud-init. Instead of changing several settings, I traced how far the script had progressed. Go and the restricted service user existed, but the application binary and `systemd` unit did not. The cloud-init log revealed the actual cause: its environment did not define `HOME` or `GOCACHE`, so the Go compiler could not locate a build cache.

I updated the user-data script to create and use an explicit root-owned Go build cache, validated the Bash syntax, created a new launch-template version, and launched a replacement instance.

The replacement completed cloud-init successfully. Through Session Manager, I verified that the service was active, `/healthz` returned `ok`, and the process listened on port 8080. I then terminated both lab instances and deleted the EBS volumes, launch template, and security groups.

The main lesson was to use evidence from each stage of the boot process: cloud-init status, created files, service state, and focused logs. That made the repair specific and reproducible.

#AWS #CloudEngineering #EC2 #Go #Linux #Troubleshooting

## X / Threads

AWS Day 15: diagnosed an EC2 cloud-init failure instead of guessing. Go was installed, but the build stopped because cloud-init had no HOME or GOCACHE. I added an explicit build cache, created a new launch-template version, and verified `active`, `ok`, and `*:8080`. Full cleanup completed. #AWS #EC2 #Go
