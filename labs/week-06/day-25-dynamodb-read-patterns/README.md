# Day 25 — DynamoDB Keys and Go Read Patterns

## References

- [DynamoDB Query key conditions](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Query.KeyConditionExpressions.html)
- [DynamoDB examples for Go v2](https://docs.aws.amazon.com/sdk-for-go/v2/developer-guide/go_dynamodb_code_examples.html)

## Access patterns

1. Retrieve one job using a known user ID and job ID.
2. Retrieve the jobs belonging to one known user.

## Key design

| Attribute | Type | Shape | Example |
|---|---|---|---|
| PK | String | USER#<demo-id> | USER#demo-a |
| SK | String | JOB#<job-id> | JOB#001 |

The partition key groups a user's jobs.
The sort key identifies and orders jobs within that group.
Together, PK and SK uniquely identify an item.


## Verified evidence — October 30, 2026

Guided lab. AWS console and runtime results were reported by the learner; the coach inspected the saved source, fixture, and module files. The coach did not make authenticated AWS requests.

| Check | Observed result |
|---|---|
| Region and budget | eu-west-1; budget checked before table creation |
| Browser login recovery | Earlier browser flow returned HTTP 400; a fresh private-window login using a newly generated remote-flow link completed, exit 0 |
| Source authentication | aws-learning identity check exit 0; identity output discarded |
| Disposable table | cloud-eng-day25-jobs; Active; PK String and SK String; on-demand, Standard |
| Seed data | Exactly the three fictional items recorded in fixture.json were reported created and visible |
| Fixture JSON validation | Exit 0; saved data also matched the expected three-item fixture |
| Module setup and dependencies | Initialization and installation exit 0; go.mod and go.sum saved |
| Formatting, module tidy, build | Each exit 0 |
| Exact GetItem content check | item check: PASS |
| Exact Query content/order check | query check: PASS |
| Overall live reader | Exit 0 |
| Console cleanup | Table deleted and absent after refresh; no backup created |
| SDK post-delete check | TableMissing; exit 1 |
| Source logout | Exit 0 |
| Other resources or optional features | None reported |
| Saved README presence check | Exit 0; substantive contents also reviewed |

## Fixture

| PK | SK | jobStatus |
|---|---|---|
| USER#demo-a | JOB#001 | queued |
| USER#demo-a | JOB#002 | complete |
| USER#demo-b | JOB#001 | queued |

## Read behavior and explanations

- GetItem supplies the complete composite key. The same job ID can occur under different users, so the partition key and sort key together identify the intended item.
- Query requires partition-key equality. This reader selects USER#demo-a and uses begins_with on the JOB# sort-key prefix, then checks both expected jobs and their statuses in ascending sort-key order.
- Base-table reads use ConsistentRead=true. Every configuration/API operation has a 15-second context bound.
- A Query response is limited to a page of up to 1 MB. If LastEvaluatedKey is nonempty, this tiny-fixture reader returns IncompletePage and a nonzero exit. It does not implement a general paginator, and large-page behavior was not exercised.
- Query selection does not establish IAM tenant isolation. This lab did not prove that the source identity was denied access to demo-b.
- A missing item would require an empty-item check. After the entire table was deleted, ResourceNotFoundException was mapped to the safe TableMissing category.

The learner explained that GetItem needs both keys, Query selects a user's jobs, a nonempty LastEvaluatedKey yields IncompletePage, and querying demo-a does not prove IAM blocks demo-b.

## Safe command summary

All relative paths were resolved from /Users/mightyzeus/Desktop/cloud-eng.

- Created a separate Go module and installed SDK v2 configuration and DynamoDB packages.
- Formatted the source, tidied the module, and built the executable outside the repository.
- Cleared credential environment overrides in a subshell and selected AWS_PROFILE=aws-learning for the reader.
- Ran the same read-only reader after table deletion and recorded only its safe category and exit status.
- Discarded identity/logout output and recorded exit statuses. No login URLs, codes, private configuration, raw AWS responses, account IDs, or credentials are stored here.

## Failure and recovery

AWS's browser sign-in page returned HTTP 400 after session selection. A new remote-flow authorization link opened in a fresh private browser session completed successfully. Authentication then passed. This establishes a successful recovery path; it does not prove the original failure's root cause.

## Cost, cleanup, and limits

DynamoDB requests and storage can incur charges; on-demand is pay-per-request. The budget was checked. Actual billed cost and focused study time were not measured.

The disposable table and its items were deleted in the same successful lab session. Console absence and the SDK TableMissing result support cleanup. No backup or other resource was reported created. Source logout succeeded; no claim of universal session revocation is made.

This lesson covers key design and Go reads. Go writes, recovery, general pagination, IAM tenant isolation, and production-scale operation were not demonstrated.

## Verified artifacts

- [Go reader](main.go)
- [Fictional fixture](fixture.json)
- [Learning note](../../../docs/week-06/day-25-dynamodb-read-patterns.md)
