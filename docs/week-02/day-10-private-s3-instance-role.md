# Day 10 — Private S3 Access from an EC2 Instance Role

## Outcome

Used an EC2 instance role to copy and synchronize files with one private S3 bucket in `eu-west-1`. Verified the bucket's public-access protections and encryption, confirmed downloaded content matched its source, and removed every temporary resource.

## Configuration and evidence

- The short-lived instance ran Amazon Linux 2023 on `t3.micro` and carried the tag `Project=cloud-eng-journey`.
- `AWSLearningEC2SSMRole` supplied temporary credentials to the AWS CLI. No access key or SSH key was created.
- The security group had zero inbound rules, and the instance was reached through Session Manager.
- All four S3 Block Public Access settings were `true`.
- Default server-side encryption used `AES256`.
- The temporary `AWSLearningDay10S3` inline policy allowed only the required bucket-level and object-level operations for one bucket.
- `aws s3 cp` and `aws s3 sync` succeeded through the instance role.
- The download matched its source, and the verified object count was four: one object under `manual/` and three under `sync/`.

## Access and security model

The instance obtained temporary credentials automatically from its attached IAM role, so the workload needed no stored AWS keys. The inline policy restricted access to one bucket and the objects inside it. IAM permissions determine which authenticated principals may call S3 operations, while Block Public Access prevents configurations that would expose the bucket publicly. `cp` performed a specific transfer; `sync` compared a directory and destination and copied new or changed files.

## Review correction

An initially submitted object count did not match the four objects created by the prescribed commands. Review kept the lab open until the actual recorded count of four was confirmed.

## Cleanup

The S3 objects and bucket were deleted, and the temporary inline policy was removed from the role. The EC2 instance was terminated and zero Day 10 EBS volumes remained. The reusable base SSM role and zero-inbound security group were retained for later labs.

## References

- [Blocking public access to S3 storage](https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-control-block-public-access.html)
- [IAM roles for Amazon EC2](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/iam-roles-for-amazon-ec2.html)
