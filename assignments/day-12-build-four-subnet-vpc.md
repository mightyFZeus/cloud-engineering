# Day 12 — Build and Inspect a Four-Subnet VPC

**Date:** Tuesday, October 6, 2026

**Due:** 11:00 PM Africa/Lagos

**Timebox:** 75–90 minutes

## Outcome

Build the Day 11 network plan in `eu-west-1`, prove which subnets are public or private from their route tables, and delete every temporary network resource.

## Learn — 10–15 minutes maximum

- Review [Create a VPC](https://docs.aws.amazon.com/vpc/latest/userguide/create-vpc.html).
- Read [Configure route tables](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html), focusing on local and default routes.

## Build — 45–55 minutes

1. Sign in with the everyday MFA-protected identity, select `eu-west-1`, and check the budget dashboard.
2. Create a VPC named `cloud-eng-day12` with:
   - IPv4 CIDR: `10.20.0.0/16`
   - IPv6 CIDR: none
   - Tenancy: default
3. Create four subnets using two available Availability Zones. Record the actual AZ choices:
   - `cloud-eng-day12-public-a`: `10.20.0.0/24`
   - `cloud-eng-day12-private-a`: `10.20.1.0/24`
   - `cloud-eng-day12-public-b`: `10.20.2.0/24`
   - `cloud-eng-day12-private-b`: `10.20.3.0/24`
4. Create `cloud-eng-day12-igw` and attach it to the VPC.
5. Create `cloud-eng-day12-public-rt`:
   - Keep the VPC-local route.
   - Add `0.0.0.0/0` with the internet gateway as the target.
   - Explicitly associate both public subnets.
6. Enable automatic public IPv4 assignment on both public subnets. Leave it disabled on both private subnets.
7. Create `cloud-eng-day12-private-rt`:
   - Keep the VPC-local route.
   - Do not add `0.0.0.0/0`.
   - Explicitly associate both private subnets.
8. Record only route destinations, target types, association counts, and automatic public IPv4 settings. Do not submit resource IDs or raw console output.

## Clean up — 10–15 minutes

1. Delete the four subnets and both custom route tables.
2. Detach and delete the internet gateway.
3. Delete `cloud-eng-day12`.
4. Confirm no Day 12 VPC, subnet, custom route table, internet gateway, NAT gateway, Elastic IP, EC2 instance, or load balancer remains.

## Safety and cost rules

- Create no NAT gateway, Elastic IP, EC2 instance, load balancer, or VPC endpoint.
- A VPC, subnet, route table, and internet gateway have no hourly charge by themselves, but check the current console estimate and budget.
- Do not submit account IDs, resource IDs, ARNs, credentials, or IP addresses assigned to real resources.

## Submit for verification

```text
Day 12 submission

1. Region; budget checked:
2. VPC CIDR:
3. Four subnet purpose/AZ/CIDR rows:
4. Public route table: associated subnet count; default-route destination and target type:
5. Private route table: associated subnet count; routes present:
6. Automatic public IPv4 assignment: public subnets; private subnets:
7. Explain what made the two public subnets public and the two private subnets private (2–4 sentences):
8. NAT gateways, Elastic IPs, EC2 instances, and load balancers created:
9. VPC, subnets, custom route tables, and internet gateway deleted: yes/no
10. Exact blocker, if any:
```

## Pass criteria

- [ ] One `10.20.0.0/16` VPC contained four valid `/24` subnets across two Availability Zones.
- [x] Two public subnets were associated with a route table containing `0.0.0.0/0` to an internet gateway.
- [x] Two private subnets were associated with a route table containing no internet default route.
- [x] Automatic public IPv4 assignment was enabled only for the public subnets.
- [x] No NAT gateway, Elastic IP, EC2 instance, or load balancer was created.
- [x] The VPC, subnets, custom route tables, and internet gateway were deleted after verification.

## Verification — October 5, 2026

**Result:** Passed with a documented exception. Five of six criteria were supported, and the routing and cleanup outcome was accepted without rebuilding the deleted lab.

- Both public subnets used a route table with `0.0.0.0/0` targeting an internet gateway.
- Both private subnets used a route table with only the VPC-local route.
- Automatic public IPv4 assignment was enabled on the public subnets and disabled on the private subnets.
- No NAT gateway, Elastic IP, EC2 instance, or load balancer was created.
- The VPC and its temporary network resources were deleted.
- The exact VPC CIDR was not retained before deletion. The submitted VPC CIDR conflicted with the recorded `/20` subnet CIDRs, so the planned `10.20.0.0/16` and `/24` layout remains unverified.
