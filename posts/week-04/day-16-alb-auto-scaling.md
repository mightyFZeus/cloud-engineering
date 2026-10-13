# Day 16 posts

## LinkedIn

Day 16 of my cloud engineering journey: I deployed my Go health service behind an Application Load Balancer and an Auto Scaling group across two Availability Zones.

I restricted the ALB to HTTP from my current address and allowed backend port 8080 only from the ALB security group. The instances used an IAM role, IMDSv2, and no SSH key pair.

Both targets became healthy, and the health endpoint returned ok. I then terminated one instance deliberately. Auto Scaling launched a replacement, restoring two InService instances and two healthy targets.

A useful troubleshooting lesson: a browser timeout did not establish that the application was down. Both target health checks passed, and a direct HTTP curl test worked. The browser-specific cause remained unconfirmed.

I deleted the ALB, Auto Scaling group, target group, launch template, instances, and temporary security groups. Zero lab EBS volumes remained; no NAT gateways or Elastic IPs were created.

Next: practice diagnosing an application health-check failure.

#AWS #CloudEngineering #Go #LearningInPublic

## X / Threads

Day 16: Go service behind an ALB + Auto Scaling across 2 AZs. Two healthy targets, then a deliberate instance termination: replacement restored capacity and health. All lab resources deleted; 0 EBS volumes remain. #AWS #CloudEngineering
