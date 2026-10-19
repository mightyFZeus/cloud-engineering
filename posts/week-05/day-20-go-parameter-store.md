# Day 20 posts

## LinkedIn

Day 20 of my cloud engineering journey: I used the AWS SDK for Go v2 to read a non-secret setting from Systems Manager Parameter Store.

The program used temporary credentials from my existing AWS login profile through the default credential chain. It checked a Standard String parameter, used a 15-second timeout, and printed a concise PASS result without exposing credentials or raw responses.

I tested the failure path too: an absent parameter produced a typed ParameterNotFound error and a nonzero exit status.

During cleanup, the CLI initially failed to reach the sign-in endpoint. A later connectivity check received an HTTP response, and retrying deletion succeeded. The initial failure's root cause remained unconfirmed.

After deleting the parameter, the Go program returned ParameterNotFound again. I then logged out of the source profile. No long-term access keys or EC2 instances were created.

Next: find the lab's API activity in CloudTrail.

#AWS #Go #CloudEngineering #LearningInPublic

## X / Threads

Day 20: Go AWS SDK v2 read a Standard String parameter using temporary profile credentials. Tested success and typed ParameterNotFound errors, deleted the parameter, verified absence, and logged out. No long-term keys or EC2 created. #AWS #Go
