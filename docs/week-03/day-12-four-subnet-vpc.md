# Day 12 — Four-Subnet Amazon VPC

## Outcome

Built and inspected a temporary Amazon VPC in `eu-west-1` with two public and two private subnets across two Availability Zones, then removed all lab resources.

## Verified configuration

| Component | Verified behavior |
|---|---|
| Public subnets | Both were associated with a route table containing `0.0.0.0/0` to an internet gateway. |
| Private subnets | Both were associated with a route table containing only the VPC-local route. |
| Automatic public IPv4 | Enabled on both public subnets and disabled on both private subnets. |
| Paid resources | No NAT gateway, Elastic IP, EC2 instance, or load balancer was created. |
| Cleanup | The VPC, four subnets, custom route tables, and internet gateway were deleted. |

## CIDR record

The submitted VPC CIDR conflicted with the recorded subnet CIDRs. The exact VPC CIDR was not retained before the VPC was deleted, so the planned `10.20.0.0/16` VPC with four `/24` subnets could not be verified. The recorded subnet networks were `/20` ranges under `10.0.0.0/16`.

The practical routing and cleanup work was accepted with this exception documented rather than requiring the deleted lab to be rebuilt.

## What determines public and private behavior

A public subnet has a route table with an internet default route to an internet gateway. For direct IPv4 internet communication, a resource in that subnet also needs a public IPv4 address and network security rules that allow the traffic.

The private subnets had no internet default route. A NAT device could provide general outbound IPv4 access, while a VPC endpoint could provide private access to supported AWS services, but neither was required for this lab.

## Cleanup and security

All temporary networking resources were reported deleted. Resource identifiers and the submitted console link were deliberately excluded from this repository.

## References

- [Create a VPC](https://docs.aws.amazon.com/vpc/latest/userguide/create-vpc.html)
- [Configure route tables](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html)
