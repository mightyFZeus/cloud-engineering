# Day 25 — DynamoDB Keys and Go Read Patterns

## References

- [Query key conditions](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Query.KeyConditionExpressions.html)
- [Go SDK DynamoDB examples](https://docs.aws.amazon.com/sdk-for-go/v2/developer-guide/go_dynamodb_code_examples.html)

**Verified:** October 30, 2026

**Roadmap:** Week 6 — storage and databases

**Assessment:** Passed, six criteria supported by the learner's console/runtime checkpoints and inspection of saved artifacts.

## Outcome

Designed a String PK/SK composite key for two access patterns: retrieve one known user/job pair, and retrieve one user's jobs. In a guided lab, the Go SDK v2 reader used temporary source-profile credentials to perform exact GetItem and ordered Query checks against one disposable on-demand Standard table in eu-west-1.

The fixture contains two jobs for demo-a and one for demo-b. Both users can have JOB#001 because the complete composite keys differ.

## Verification

| Evidence | Result |
|---|---|
| Budget, source login, authentication | Checked; login and identity check exit 0 |
| Table setup | Expected String composite keys; on-demand Standard; Active |
| Fixture | Three intended fictional items reported visible; saved JSON matched exactly |
| Formatting, module tidy, build | Exit 0 each |
| GetItem | Exact PK/SK/status check passed |
| Query | Exactly two expected demo-a jobs, expected statuses and ascending order; passed |
| Overall reader | Exit 0 |
| Table deletion | Absent after console refresh; no backup created |
| Post-delete SDK read | TableMissing, exit 1 |
| Source logout | Exit 0 |
| Other resources/features | None reported |
| README | Nonempty check exit 0; substantive evidence inspected |

The actual saved reader matches the reviewed implementation. An offline build against the installed dependencies also succeeded. AWS runtime tests were learner-run and were not independently repeated by the coach.

## Engineering decisions

GetItem uses both key attributes. Query uses partition-key equality and a sort-key prefix, with expression placeholders, strongly consistent base-table reads, and bounded contexts. Content checks reject missing or incorrect items, unexpected query content, and incorrect ordering.

The reader detects a nonempty LastEvaluatedKey and fails with IncompletePage. The fixture did not test the 1 MB boundary, and the program does not implement general pagination. Querying demo-a is a selection pattern and does not prove IAM denies access to demo-b.

## Troubleshooting

A browser HTTP 400 initially blocked login. A fresh private-window login using a newly generated remote-flow link completed, followed by a successful source identity check. The recovery was observed; the original root cause was not established.

## Cost and cleanup

The table was short-lived and used three fictional items and a few reads. DynamoDB requests and storage can incur charges. Actual cost and focused hours were not measured.

The learner deleted the table and confirmed console absence. The reader then produced the safe TableMissing category with exit 1. No backup was created, no other resources were reported, and source logout succeeded.

## Safe artifacts

- [Lab evidence and command summary](../../labs/week-06/day-25-dynamodb-read-patterns/README.md)
- [Reader](../../labs/week-06/day-25-dynamodb-read-patterns/main.go)
- [Fictional fixture](../../labs/week-06/day-25-dynamodb-read-patterns/fixture.json)
- [Social drafts](../../posts/week-06/day-25-dynamodb-read-patterns.md)

Go writes, recovery, IAM tenant isolation, general pagination, and a complete storage/serverless milestone remain outside this verified lesson.
