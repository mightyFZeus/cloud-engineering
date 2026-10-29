# Day 25 — DynamoDB Keys and Go Read Patterns

## Read first — 10 minutes

1. [DynamoDB Query key conditions](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Query.KeyConditionExpressions.html) — focus on partition-key equality, sort-key conditions, sorted results, and the 1 MB pagination limit.
2. [DynamoDB examples for Go v2](https://docs.aws.amazon.com/sdk-for-go/v2/developer-guide/go_dynamodb_code_examples.html) — focus only on GetItem and Query. This assignment uses a smaller fixture than the full movie example.

**Date:** October 29, 2026
**Loaded:** October 27, 2026, after Day 24 passed early
**Original due:** October 28 at 11:00 PM Africa/Lagos
**Due:** October 29 at 11:00 PM Africa/Lagos
**Timebox:** 75–90 focused minutes, including cleanup and notes
**Status:** Catch-up carried forward on October 29; no submission received for October 28. Complete this same lesson before starting another day.
**Roadmap section:** Week 6 — storage and databases

## Outcome

Design a composite key from two access patterns, then use Go to retrieve one fictional job and query one user's jobs from a short-lived DynamoDB table. Reuse the familiar Go SDK default-credential pattern. Table creation, three seed writes, and deletion happen through the console; Go implements only GetItem and Query in this lesson.

## Numbered steps

1. Read the references. Before creating a table, write two access patterns in your README: retrieve a job using a known user/job pair; retrieve the jobs belonging to one known user. Choose String keys PK and SK, and explain why both are useful. Plan the key shapes USER#<demo-id> and JOB#<job-id>.
2. Prepare labs/week-06/day-25-dynamodb-read-patterns/README.md, main.go, fixture.json, and Go module files under /Users/mightyzeus/Desktop/cloud-eng. Suggested allocation: 20 minutes reading/design, 15 table/seed, 25 Go code, 10 verification, and 20 cleanup/notes.
3. Confirm eu-west-1 and budget. Renew aws-learning with normal aws login and use the source profile directly. Day 24's role/profile were removed and are not prerequisites. The earlier SSM CLI connection issue does not establish whether DynamoDB requests will work; report the actual Go result without inventing a cause or exporting credentials.
4. In DynamoDB console, create only cloud-eng-day25-jobs with partition key PK (String), sort key SK (String), on-demand capacity, and Standard table class. If that exact table already exists, stop before modifying it. Add learning-project tags. Use default AWS-owned encryption, no customer-managed key, secondary index, replica/global table, stream, backup/PITR feature, or additional integration. Keep deletion protection off for this disposable lab. Wait for Active.
5. Record and create exactly these three fictional items using the console item editor. fixture.json documents the same three logical items; do not paste a whole JSON array into an editor expecting one item.

| PK | SK | jobStatus |
|---|---|---|
| USER#demo-a | JOB#001 | queued |
| USER#demo-a | JOB#002 | complete |
| USER#demo-b | JOB#001 | queued |

6. Implement a small Go reader using AWS SDK for Go v2, config.LoadDefaultConfig, explicit eu-west-1, a -table flag, bounded API contexts, and sanitized error categories. Reuse Day 20's credential/error-handling pattern; add the DynamoDB service dependency. Keep table writes and lifecycle operations out of this program. GetItem must supply both key attributes for demo-a/JOB#001 and assert returned PK, SK, and jobStatus=queued. A successful HTTP response with no matching item is not a passing content check.
7. Query using partition equality PK=USER#demo-a plus begins_with(SK, JOB#), with expression placeholders. Assert the result contains exactly JOB#001/queued and JOB#002/complete for demo-a in ascending sort-key order, with no demo-b item. Use ConsistentRead=true for base-table reads to avoid ambiguity immediately after seeding. Use Query rather than Scan or a non-key filter to satisfy the access pattern.
8. Explain the 1 MB query-page limit and LastEvaluatedKey. Either use the SDK Query paginator or explicitly detect a nonempty LastEvaluatedKey and fail with an incomplete-page category for this tiny fixture. Do not claim general pagination testing when only three small items were used.
9. Run gofmt and build. With AWS_PROFILE=aws-learning and credential environment overrides cleared in a subshell, require item check PASS and query check PASS with exit 0. Record only concise check summaries, expected fictional keys/statuses, and exit status. Explain why GetItem needs the complete composite key and Query needs partition-key equality. Querying one user's partition is a data-access pattern, not an IAM tenant-isolation guarantee.
10. Delete only the temporary table through the console, without creating a final backup. Confirm it disappears after deletion finishes. Re-run the Go reader and require ResourceNotFoundException (or the program's corresponding sanitized table-missing category) with nonzero exit. Log out aws-learning. If testing blocks, delete the table before reporting the sanitized blocker; stop repeated troubleshooting after 10 minutes instead of leaving storage running.
11. Save access patterns, key rationale, fixture, code, observed checks, capacity mode, pagination limits, and cleanup status. Run test -s on README.md and record the exit status. Inspect files for secrets, real ARNs/account IDs, login material, raw errors, and private config before submission. No actual user data is needed.

## Cost and security

DynamoDB charges for requests and stored data. On-demand is pay-per-request, not a promise of free usage. Check [current pricing](https://aws.amazon.com/dynamodb/pricing/) and the account budget before creating the table. Limit this lab to three dummy writes and a few reads; delete the table in the same session. Create no backups, indexes, replicas, integrations, reserved commitments, EC2 instances, or other supporting infrastructure.

Use temporary source credentials through the SDK's default chain. Store no keys, real resource ARNs, account IDs, raw API responses/errors, session material, or private configuration in the repository or chat. Preserve evidence as safe booleans, check labels, fictional fixture data, sanitized categories, and cleanup confirmations.

## Submission

1. Two access patterns and chosen PK/SK types/shapes:
2. Region/budget, table name, capacity mode, class, and Active status:
3. Three expected fictional items seeded:
4. Files present, gofmt completed, and build exit status:
5. GetItem check result, Query check result, and run exit status:
6. Complete-key versus partition-query explanation:
7. Pagination handling/limitation and data-pattern versus IAM-isolation explanation:
8. Table deleted and absent; post-delete Go category and nonzero exit:
9. Source logout result; other AWS resources created:
10. README check exit status and exact blocker, if any:

## Pass criteria

- Two access patterns motivate a correct String PK/SK design before table creation; fixture is fictional and consistent.
- Exactly one intended on-demand Standard table with the expected keys and three items; no unplanned features/resources.
- Reviewed Go code uses the default credential chain, bounded calls, correct GetItem/Query inputs, and safe output without Scan substitution.
- Actual exact-item and exact-two-item Query checks pass with exit 0; content validation goes beyond HTTP success or item count alone.
- Learner explains complete-key lookup, partition equality, pagination limits, and the difference between key selection and authorization isolation.
- Table absence corroborated by console and the SDK table-missing result, source logout confirmed, safe substantive README/code/fixture saved, and no credentials or unintended resources retained.
