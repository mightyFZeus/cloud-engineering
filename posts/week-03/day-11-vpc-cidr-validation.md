# Day 11 Build-in-Public Draft — VPC CIDR Validation in Go

## LinkedIn

Day 11 of my AWS cloud engineering journey moved into VPC address planning.

Before creating networking resources, I divided a planned `/16` VPC into four non-overlapping `/24` subnets: one public and one private subnet in each of two Availability Zones.

I wrote a small Go validator using the standard `net/netip` package. It checks that every subnet has the expected prefix size, belongs to the VPC, and does not overlap another subnet. The valid plan passed with four subnets.

I then changed one private subnet to use the same CIDR as a public subnet. The validator identified the overlap and exited with a failure status. Restoring the correct CIDR returned the plan to a successful result.

The networking lesson was that a CIDR does not make a subnet public. Public or private behavior comes from its route table. A public subnet routes to an internet gateway, while the private subnets in this plan have only the VPC-local route.

No AWS resources were created for this exercise. Next, I’ll build and inspect the design in Amazon VPC, then remove it.

#AWS #CloudEngineering #Networking #Go #VPC

## X

AWS Day 11: planned four `/24` subnets across two AZs and built a Go validator for VPC containment and overlap. A deliberate duplicate CIDR failed as expected; restoring it passed. No AWS resources created. Next: build the VPC. #AWS #Go #Networking
