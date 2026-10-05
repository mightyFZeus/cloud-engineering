# Day 11 — VPC CIDR Planning and Validation

## Outcome

Designed four non-overlapping `/24` subnets inside the planned `10.20.0.0/16` VPC, distributed them across two Availability Zones, and validated the plan with a dependency-free Go program before creating any AWS networking resources.

## Network plan

| Purpose | Availability Zone | CIDR | Intended default route |
|---|---|---|---|
| Public A | `eu-west-1a` | `10.20.0.0/24` | Internet gateway |
| Private A | `eu-west-1a` | `10.20.1.0/24` | None; VPC-local route only |
| Public B | `eu-west-1b` | `10.20.2.0/24` | Internet gateway |
| Private B | `eu-west-1b` | `10.20.3.0/24` | None; VPC-local route only |

The `/16` leaves sixteen address bits available. Moving to `/24` uses eight of those bits to identify a subnet, allowing 256 separate `/24` networks. The chosen networks use different third octets, so their address ranges do not overlap.

## Validator

The lab source is stored in `labs/week-03/day-11-vpc-cidr/`. The Go program uses `net/netip` to check that every subnet:

- parses as a valid prefix;
- uses a `/24` prefix;
- begins inside the `10.20.0.0/16` VPC; and
- does not overlap a previously checked subnet.

`gofmt -d` reported no formatting changes. The saved plan exited successfully after validating all four subnets. An isolated test changed `private-a` to the same prefix as `public-a`; the validator identified the overlap and exited nonzero. The unchanged saved program then passed again.

## Routing and security model

A public subnet has a default route to an internet gateway, but a resource also needs a public IPv4 address for direct IPv4 internet communication. A private subnet has no direct internet-gateway route and commonly uses NAT for general outbound IPv4 internet access, while VPC endpoints can provide private access to supported AWS services. Security groups are stateful allow-only controls associated with resources or network interfaces. Network ACLs are stateless subnet-level filters that support allow and deny rules, so return traffic must be permitted explicitly.

## Cost and security

This was a local planning exercise. No AWS resources were created and the lab incurred no AWS resource cost. The saved README and Go source contain no credentials, account identifiers, ARNs, personal addresses, or private environment details.

## References

- [What is Amazon VPC?](https://docs.aws.amazon.com/vpc/latest/userguide/what-is-amazon-vpc.html)
- [VPC subnet basics](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-subnet-basics.html)
