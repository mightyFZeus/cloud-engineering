# Day 4 — Regions, Availability Zones, and Shared Responsibility

**Date:** September 17, 2026  
**Status:** Passed — 6/6 checks

## Live infrastructure map

The AWS CLI reported three standard Availability Zones for this account in `eu-west-1`:

```text
AWS account
└── Region: eu-west-1 (Ireland)
    ├── eu-west-1a → euw1-az2
    ├── eu-west-1b → euw1-az3
    └── eu-west-1c → euw1-az1
```

AZ names can be account-relative, while AZ IDs identify the same physical location across accounts. Recording both makes infrastructure comparisons clearer.

## Availability lesson

A Region is a geographical area containing multiple isolated Availability Zones. A Go API can survive one AZ failure only when it actually runs in at least two AZs and traffic is routed away from unhealthy instances. Multi-AZ architecture does not protect against a complete regional failure; that requires a multi-Region design.

## Shared responsibility: EC2 versus Lambda

The AWS Shared Responsibility Model shifts the operational burden depending on the compute service selected.

With EC2, AWS manages physical data centers, hardware, underlying networking, and virtualization. The customer operates at the infrastructure level and remains responsible for the guest operating system, OS patching, application dependencies, network security groups, IAM access, and data.

With Lambda, AWS absorbs more operational overhead by managing the infrastructure, host operating system, capacity scaling, and managed runtime platform. The customer still owns function code and dependencies, execution configuration, least-privilege IAM permissions, secrets, logging and monitoring, application security, and data protection. Serverless reduces operational work, but it does not eliminate the customer's security responsibility.

## References

- [AWS Regions and Availability Zones](https://docs.aws.amazon.com/global-infrastructure/latest/regions/aws-regions-availability-zones.html)
- [AWS Shared Responsibility Model](https://aws.amazon.com/compliance/shared-responsibility-model/)

