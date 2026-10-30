# Day 25 — DynamoDB Keys and Go Read Patterns

**Verified:** October 30, 2026. User-ready drafts; external publication is not verified.

## LinkedIn

Day 25 of my AWS cloud engineering journey: I designed DynamoDB keys around the reads my application needs.

In this guided Go lab, I started with two access patterns: fetch one known job, and fetch the jobs belonging to one user. A String partition key grouped each user's jobs, while the sort key distinguished jobs within that group.

I seeded three fictional items and used the AWS SDK for Go v2 with temporary profile credentials. GetItem and Query both passed checks for the expected keys, statuses, and query order.

Two useful lessons: a successful request still needs content validation, and selecting one user's partition doesn't prove IAM prevents access to another user's data. My reader also detects an unfinished query page instead of treating it as complete.

Cleanup was part of the result: I deleted the table, verified TableMissing with a nonzero exit, and logged out.

Next: conditional writes that reject duplicate creates and stale updates.

https://github.com/mightyFZeus/cloud-engineering

#AWS #Golang #DynamoDB

## X / Threads

Day 25: designed DynamoDB keys, verified Go GetItem + Query with temporary credentials, then deleted the table and confirmed TableMissing. Query scope isn't IAM isolation.
https://github.com/mightyFZeus/cloud-engineering
#AWS #Golang
