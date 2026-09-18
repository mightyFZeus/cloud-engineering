# Day 6 Build-in-Public Draft — Least-Privilege IAM

## LinkedIn

Day 6 of my AWS cloud engineering journey was about turning “read only” into actual least privilege.

Yesterday, I used AWS's broad `ReadOnlyAccess` policy to learn role assumption and policy evaluation. Today, I replaced it with an inline policy that permits only three discovery operations: describing Regions, describing Availability Zones, and listing S3 buckets.

I verified the intended reads through the assumed-role CLI profile, then used IAM policy simulation to confirm that an S3 write and IAM user creation were both implicitly denied—without executing either operation.

One useful nuance: `Resource: "*"` does not automatically make a policy unrestricted. Some list and describe actions cannot be scoped to individual resource ARNs. In this case, least privilege comes from keeping the permitted `Action` list extremely small.

Next, I’m continuing the foundation work and applying the same evidence-first approach to each permission change.

#AWS #CloudEngineering #IAM #LeastPrivilege

## X

AWS Day 6: replaced broad ReadOnlyAccess with a 3-action inline policy. Verified the intended discovery calls, then simulated and confirmed implicit denies for S3 writes and IAM user creation. `Resource: "*"` can still be least privilege when actions are tightly scoped. #AWS #IAM
