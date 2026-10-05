# Day 13 — VPC Traffic Path and Diagnosis

## Outcome

Created a two-Availability-Zone VPC traffic diagram and an evidence-based troubleshooting table for an internet-facing load balancer serving private application workloads.

## Request path

1. A client resolves the load balancer hostname through DNS.
2. HTTPS traffic reaches the internet-facing load balancer through the VPC internet gateway and public-subnet route.
3. The load-balancer security group evaluates inbound HTTPS traffic.
4. The load balancer forwards the request to a workload in a private subnet.
5. The application security group permits the application port from the load-balancer security group.
6. Stateful security groups allow response traffic for the accepted connection.

Route tables select a next hop according to the destination. Security groups filter traffic at associated resources and network interfaces.

## Egress choices

A private workload can use a NAT device for general outbound IPv4 access when its route table directs the default route to that NAT device. An S3 gateway endpoint provides a private path to Amazon S3 and can avoid NAT for supported S3 traffic.

## Failure diagnosis

The lab documented five failures: a missing public default route, missing HTTPS access on the load-balancer security group, a blocked application port, incorrect direct internet-gateway use by a private workload, and a network ACL blocking return traffic. Each entry records a symptom, the evidence to inspect, a root cause, and one specific repair.

## Verification and cleanup

- The README existence check exited `0`.
- The required-concepts search exited `0`.
- No AWS resources were created.
- No sensitive identifiers were recorded.

## References

- [Route tables](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html)
- [Security group rules](https://docs.aws.amazon.com/vpc/latest/userguide/security-group-rules.html)
