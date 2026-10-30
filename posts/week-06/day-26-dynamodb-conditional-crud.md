# Day 26 — Conditional DynamoDB CRUD with Go

**Verified:** October 30, 2026. User-ready drafts; external publication is not verified.

## LinkedIn

Day 26 of my AWS cloud engineering journey: I tested when a DynamoDB write should be rejected.

In this guided Go lab, one fictional job moved from absent to queued, then running, then absent again.

I attempted two writes that should fail: creating the same item again and updating it using an outdated queued-state assumption. Both returned ConditionalCheckFailedException. Strongly consistent reads then confirmed that the existing item stayed unchanged.

That distinction matters: an error alone doesn't prove a condition protected the data. The program checked the specific error category and the resulting state.

The successful update and delete also checked the returned item contents. All six checks passed with temporary profile credentials.

Finally, I deleted the table, verified TableMissing through a separate read-only reader, and logged out. A missing item and a missing table are different outcomes.

Next: private S3 transfers using short-lived presigned URLs.

https://github.com/mightyFZeus/cloud-engineering

#AWS #Golang #DynamoDB

## X / Threads

Day 26: Go + DynamoDB conditional writes. Duplicate creates and stale updates were rejected; reads confirmed unchanged state. Deleted the item and table, verified TableMissing, and logged out.
https://github.com/mightyFZeus/cloud-engineering
#AWS #Golang
