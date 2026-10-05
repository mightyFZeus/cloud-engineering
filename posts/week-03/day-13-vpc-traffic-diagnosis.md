# Day 13 Build-in-Public Draft — VPC Traffic Diagnosis

## LinkedIn

Day 13 of my AWS cloud engineering journey focused on tracing traffic through a two-Availability-Zone VPC and diagnosing failures from evidence.

I diagrammed the path from an internet client to an internet-facing Application Load Balancer in public subnets, then to application workloads in private subnets. The design labels the internet gateway, public and private route tables, load-balancer and application security groups, an optional NAT path, and an S3 gateway endpoint.

The exercise reinforced the separation between routing and filtering. Route tables choose the next network hop, while security groups decide whether traffic can reach an associated resource. Security groups are stateful, whereas network ACLs are stateless and require return traffic to be allowed explicitly.

I also created a five-scenario troubleshooting table covering a missing internet-gateway route, blocked HTTPS, blocked load-balancer-to-application traffic, incorrect direct internet-gateway use from a private workload, and a network ACL blocking return traffic. For every failure, I documented the symptom, evidence to inspect, root cause, and one specific repair.

No AWS resources were required for this lab.

#AWS #CloudEngineering #AmazonVPC #Networking #Troubleshooting

## X / Threads

AWS Day 13: diagrammed a two-AZ path from an internet client through an ALB to private workloads, then documented five evidence-based failure diagnoses. Routes choose the next hop; security groups filter resource traffic; NACLs are stateless. No AWS resources created. #AWS #VPC
