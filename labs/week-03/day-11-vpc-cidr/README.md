# Day 11 — VPC CIDR Plan

## VPC

- Region: eu-west-1
- VPC CIDR: 10.20.0.0/16

## Subnets

| Purpose | Availability Zone | CIDR | Intended default route |
|---|---|---|---|
| Public A | eu-west-1a | 10.20.0.0/24 | 0.0.0.0/0 to Internet Gateway |
| Private A | eu-west-1a | 10.20.1.0/24 | No default route; local VPC route only |
| Public B | eu-west-1b | 10.20.2.0/24 | 0.0.0.0/0 to Internet Gateway |
| Private B | eu-west-1b | 10.20.3.0/24 | No default route; local VPC route only |

## Networking explanation

A public subnet has a default route to an internet gateway, but a resource also needs a public IPv4 address for direct IPv4 internet communication. A private subnet has no direct route to an internet gateway and commonly uses a NAT device for general outbound IPv4 internet access, although VPC endpoints can provide private access to supported AWS services. Security groups are stateful allow-only firewalls associated with resources or network interfaces, so permitted response traffic is automatically allowed. Network ACLs are stateless subnet-level filters that support both allow and deny rules, which means return traffic must also be permitted explicitly. Internet-facing components such as public load balancers belong in public subnets, while application workloads and databases are commonly placed in private subnets.
