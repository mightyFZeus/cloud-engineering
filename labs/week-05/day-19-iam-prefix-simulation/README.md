
# Day 19 — S3 Prefix Policy Simulation

## Validation

JSON syntax validation exit status: 0.

## Observed simulation results

Only the custom lab policy was included; no permissions boundary was supplied.

| Action | Resource or prefix | Actual result |
|---|---|---|
| GetObject | reports/sample.txt | Allowed |
| GetObject | private/sample.txt | Implicit deny |
| PutObject | reports/sample.txt | Implicit deny |
| ListBucket | s3:prefix = reports/ | Allowed |
| ListBucket | s3:prefix = private/ | Implicit deny |

Object tests used the fictional example bucket's object ARNs.
ListBucket tests used its bucket ARN and explicitly supplied s3:prefix.

## Resource scope and policy evaluation

GetObject applies to object resources, so its resource includes reports/*.
ListBucket applies to the bucket resource; the s3:prefix condition restricts the requested listing.
An implicit deny occurs when no applicable allow grants the requested permission.
An explicit deny comes from a matching Deny statement.

## Simulator limitations

These simulations evaluate the supplied policy and context without making real S3 requests.
They do not prove live access: other policies, permissions boundaries, organization controls, and service context can affect real requests.


## Cleanup and learner explanation

No policy was attached to a real identity. No AWS resources were created. README check exited 0. The learner explained that reading reports is allowed but upload permission is absent.
