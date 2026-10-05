# Day 12 Build-in-Public Draft — Four-Subnet Amazon VPC

## LinkedIn

Day 12 of my AWS cloud engineering journey moved my VPC design from a local plan into AWS.

I created a temporary VPC across two Availability Zones with two public subnets and two private subnets. The public subnets shared a route table with an IPv4 default route to an internet gateway, while the private subnets used a separate route table with only the VPC-local route.

I also enabled automatic public IPv4 assignment only on the public subnets. This reinforced that a subnet becomes public through its routing configuration, while an individual resource still needs a public address and suitable security rules for direct IPv4 internet communication.

I avoided NAT gateways, Elastic IPs, EC2 instances, and load balancers because they were unnecessary for this routing exercise. After checking every association and route, I deleted the VPC, four subnets, custom route tables, and internet gateway.

One documentation lesson stood out: record the final VPC CIDR before cleanup. My submitted VPC and subnet CIDRs conflicted, and the deleted VPC could no longer be checked. I documented that gap instead of presenting the uncertain value as verified.

#AWS #CloudEngineering #AmazonVPC #Networking #LearningInPublic

## X / Threads

AWS Day 12: built a temporary four-subnet VPC across two AZs. Public subnets routed `0.0.0.0/0` to an internet gateway; private subnets kept only the local route. No NAT, EIP, EC2, or load balancer. Everything was cleaned up, and I documented one CIDR evidence gap. #AWS #VPC
