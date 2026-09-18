# Day 5 Build-in-Public Draft — IAM Roles

## LinkedIn

Day 5 of my AWS cloud engineering journey focused on IAM roles, temporary credentials, and policy evaluation.

I created a read-only learning role, configured a second AWS CLI profile to assume it through my temporary login session, and confirmed that the active principal was an assumed role rather than a long-lived user credential.

The useful part was testing authorization safely. Instead of attempting a real write operation, I used IAM policy simulation. The role could list S3 buckets, while `s3:PutObject` returned an implicit deny.

My main takeaway: a permission policy defines what a role can do, while its trust policy defines who can assume it. Also, an implicit deny means no policy granted the action; an explicit deny is stronger and overrides an allow.

This was intentionally a broad read-only learning policy. The next step is moving toward resource-scoped, least-privilege policies for real workloads.

#AWS #CloudEngineering #IAM #LearningInPublic

## X

AWS Day 5: created and assumed a read-only IAM role with temporary credentials, then used policy simulation to verify an allowed S3 read and an implicitly denied write—without making the write request. Next: resource-scoped least privilege. #AWS #IAM
