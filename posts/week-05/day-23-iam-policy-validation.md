# Day 23 posts

## LinkedIn

Day 23 of my cloud engineering journey: I compared JSON syntax checking with AWS IAM policy validation.

My first policy was valid JSON but contained a misspelled action: s3:GetObjet. AWS basic policy validation identified it as INVALID_ACTION.

After correcting the action to s3:GetObject, the policy had zero findings. Its resource was still "*", which exceeded the requirement to read only daily reports.

I narrowed the resource to a fictional reports/daily/ prefix and validated it again. The scoped version also had zero findings.

The lesson: matching permissions to a requirement still requires reviewing the action and resource, even after validation passes.

I also clarified what each check establishes: JSON parsing checks structure, AWS validation checks policy elements and defined concerns, simulation evaluates allow/deny decisions, and a real request tests the operation against AWS.

I canceled the unsaved editor. No AWS policy was saved or attached, no analyzer or resource was created, no paid custom check ran, and no real S3 data was accessed.

Proof of work: https://github.com/mightyFZeus/cloud-engineering/blob/main/docs/week-05/day-23-iam-policy-validation.md

#AWS #IAM #CloudEngineering #LearningInPublic

## X / Threads

Day 23: valid JSON contained an invalid IAM action. AWS validation caught it. After fixing it, a broad policy had zero findings, so I narrowed its resource to the required report prefix. Validation alone does not prove least privilege. #AWS #IAM
