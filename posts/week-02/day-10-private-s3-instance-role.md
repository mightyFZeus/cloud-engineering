# Day 10 Build-in-Public Draft — Private S3 Access from EC2

## LinkedIn

Day 10 of my AWS cloud engineering journey connected EC2 workload identity with private S3 access.

I launched a short-lived Amazon Linux 2023 instance with an IAM role and connected through Session Manager. The security group had zero inbound rules, and I created no SSH key or stored AWS access key.

A temporary inline policy limited the role to the required operations on one private bucket and its objects. I confirmed all four S3 Block Public Access settings were enabled and that default encryption used AES256. From the instance, I tested both `aws s3 cp` and `aws s3 sync`, downloaded an object, and verified that it matched its source.

The useful distinction was that IAM controls which authenticated identities may access the bucket, while Block Public Access protects against public exposure. The instance role supplied short-lived credentials automatically.

I finished by deleting the objects, bucket, temporary policy, instance, and EBS storage. Next, I’m moving into VPC address planning and subnet validation.

#AWS #CloudEngineering #S3 #IAM #EC2

## X

AWS Day 10: used an EC2 role to access one private S3 bucket with no stored keys or inbound rules. Verified Block Public Access, encryption, `cp`, `sync`, and file integrity, then removed every temporary resource. Next: VPC subnet planning. #AWS #S3 #IAM
