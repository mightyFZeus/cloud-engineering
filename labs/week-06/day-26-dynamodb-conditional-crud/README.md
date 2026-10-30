# Day 26 — Conditional DynamoDB CRUD with Go

## References

- [DynamoDB condition expressions](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.ConditionExpressions.html)
- [DynamoDB Go SDK examples](https://docs.aws.amazon.com/sdk-for-go/v2/developer-guide/go_dynamodb_code_examples.html)

## Planned item

- PK: USER#demo-a
- SK: JOB#001
- Initial jobStatus: queued

## Planned state sequence

absent -> queued -> running -> absent

## Expected guards

- Creating the same item again must fail without overwriting it.
- Updating queued to running requires the current status to be queued.
- After the status becomes running, an update expecting queued must fail.
- After deleting the item, GetItem should return no item data.


## Verified evidence — October 30, 2026

Guided lab. The learner ran the AWS console and SDK checks; the coach inspected saved code, the fixture, and module files. No authenticated AWS requests were made by the coach.

| Check | Observed result |
|---|---|
| Region and budget | eu-west-1; reused the budget check completed earlier the same day |
| Temporary source session | aws-learning remote-flow login exit 0; identity check exit 0 with identity output discarded |
| Table setup | cloud-eng-day26-jobs; PK/SK String; on-demand Standard; Active and initially empty |
| Fixture | USER#demo-a / JOB#001 / queued; JSON validation exit 0; saved object matched exactly |
| Go module setup | Exit 0; separate Day 26 module reusing Day 25 dependency versions |
| Formatting, module tidy, build | Exit 0 each |
| Overall live run | All six checks passed; exit 0 |
| Console cleanup | Table deleted and absent after refresh; no backup created |
| Read-only post-delete check | Day 25 reader pointed at the Day 26 table: TableMissing, exit 1 |
| Source logout | Exit 0 |
| Other resources or optional features | None reported |
| Saved README presence check | Exit 0; substantive contents also reviewed |

## Operation evidence

| Operation | Expected result | Observed result |
|---|---|---|
| Conditional create | Create queued job only if the complete key is absent; exact strong read | create check: PASS |
| Duplicate create | ConditionalCheckFailedException; item still queued | Exact category reported; duplicate guard: PASS |
| Guarded update | queued to running; ALL_NEW contains expected String keys/status | update check: PASS |
| Stale update | ConditionalCheckFailedException; item still running | Exact category reported; stale guard: PASS |
| Guarded delete | ALL_OLD identifies the removed running fixture | delete check: PASS |
| Get missing item | Successful response with no Item data | absent item check: PASS |

The actual source matches the independently reviewed template. The template was formatted and built offline against cached dependencies without execution. The learner then copied, formatted, tidied, built, and executed it against the disposable table.

## Explanations

- An unconditional PutItem replaces an existing item with the same complete primary key. Here that means both PK and SK. The conditional create checks absence for that one item rather than excluding every job with the same partition key.
- ConditionalCheckFailedException means the write precondition was false. At the two intentional rejection checkpoints, the write was blocked; strongly consistent reads also verified that queued or running content stayed unchanged.
- Checking jobStatus=queued before setting running prevents the tested stale update. The update condition also avoids an upsert when that expected state is absent.
- The stale attempt expected queued after the job had already become running. It was rejected without changing running to complete.
- DeleteItem additionally requires the running state. ALL_OLD proves which fixture was removed; a successful no-op deletion is insufficient.
- A missing item produces a successful GetItem response without item data. A missing table produces ResourceNotFoundException, which the read-only reader maps to TableMissing.
- These are deterministic state checks. No multi-client concurrency, transaction, retry/idempotency, load, or IAM tenant-isolation guarantee was established.

The learner explained that failed conditions blocked the writes, checking queued prevents a stale update, missing items return no data, and TableMissing means the table does not exist.

## Safe command summary

All relative paths were resolved from /Users/mightyzeus/Desktop/cloud-eng.

- Created separate lab files and module path; copied the earlier dependency versions.
- Formatted and built the reviewed writer outside the repository before table creation.
- Renewed the temporary source profile and verified authentication while suppressing identity output.
- Cleared credential environment overrides in a subshell and selected AWS_PROFILE=aws-learning for the writer.
- After console table deletion, ran the separate Day 25 read-only executable with the Day 26 table flag; no write sequence was rerun for cleanup verification.
- Logged out the source profile and recorded only its exit status.

No login links/codes, private AWS configuration, raw AWS responses, credentials, actual account IDs, or resource ARNs are retained.

## Cost and cleanup

The budget was checked earlier the same day. On-demand requests and stored data can incur charges; condition-rejected writes can also consume capacity. Actual billed cost and focused hours were not measured.

The program removed its item; the learner then deleted the table without a backup. Refreshed console absence and the SDK TableMissing result support cleanup. Other resources/features were reported as none, and source logout succeeded.

## Artifacts

- [Go writer](main.go)
- [Fictional fixture](fixture.json)
- [Verified note](../../../docs/week-06/day-26-dynamodb-conditional-crud.md)
