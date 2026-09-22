# Day 7 Build-in-Public Draft — AWS Baseline Automation

## LinkedIn

Day 7 of my AWS cloud engineering journey turned a few manual checks into a small Bash baseline script.

The script checks that my CLI profile is set to the expected Region and is using the intended assumed role. It then reports only counts for visible Regions, available standard Availability Zones, and S3 buckets. My successful run returned 17 Regions, three AZs in `eu-west-1`, and zero buckets, without printing an account ID or ARN.

I also tested the failure path: a nonexistent profile was rejected with a nonzero exit status. During review, I found that a single space before the shebang prevented direct execution. I corrected it, reran `bash -n`, and confirmed the missing-profile test still failed as intended.

My takeaway: a useful access check should verify the identity and configuration it depends on, fail clearly when they are wrong, and keep its output narrow. The `length(...)` queries limit what appears in the terminal, although the filtering happens in the CLI after AWS returns the data.

Next, I’m moving into a short-lived Linux EC2 lab with Session Manager and explicit cleanup.

#AWS #CloudEngineering #Bash #AWSCLI

## X

AWS Day 7: built a Bash check for the expected role and Region. It reported counts only: 17 Regions, 3 AZs, 0 S3 buckets. A missing profile failed as expected; review caught and fixed a shebang typo. Next: a short-lived Linux EC2 lab. #AWS #Bash
