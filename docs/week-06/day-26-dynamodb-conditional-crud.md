# Day 26 — Conditional DynamoDB CRUD with Go

## References

- [Condition expressions](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.ConditionExpressions.html)
- [Go SDK DynamoDB examples](https://docs.aws.amazon.com/sdk-for-go/v2/developer-guide/go_dynamodb_code_examples.html)

**Verified:** October 30, 2026

**Roadmap:** Week 6 — storage and databases

**Assessment:** Passed, six criteria supported by learner-run AWS results and inspection of the saved artifacts.

## Outcome

A guided Go SDK v2 program created, read, conditionally updated, and deleted one fictional DynamoDB job. Its state sequence was absent -> queued -> running -> absent. Together with Day 25's Query checks, this completes the roadmap's small Go CRUD/query exercise.

The source used the default temporary-credential chain, explicit eu-west-1, complete String PK/SK keys, expression placeholders, strongly consistent GetItem reads, bounded contexts, and safe error categories.

## Observed checks

| Evidence | Result |
|---|---|
| Source login and identity check | Exit 0 each; identity output discarded |
| Setup | One new on-demand Standard table, initially empty and Active, expected String keys |
| Fixture/module checks | Correct one-job fixture; JSON and module setup exit 0 |
| Formatting, tidy, build | Exit 0 each |
| Create | Exact queued item check passed |
| Duplicate create | ConditionalCheckFailedException; queued item preserved; guard passed |
| Update | queued to running, checked ALL_NEW; passed |
| Stale update | ConditionalCheckFailedException; running item preserved; guard passed |
| Delete | Guarded running-state removal, checked ALL_OLD; passed |
| Missing item | Successful strong read with no Item; passed |
| Overall writer | Six PASS checks; exit 0 |
| Table cleanup | Deleted and absent after refresh; no backup |
| Read-only SDK cleanup verification | Day 25 reader returned TableMissing with exit 1 |
| Source logout and other resources | Exit 0; none reported |
| Saved README | Nonempty check exit 0; substantive evidence inspected |

Actual source matched the reviewed template, which had also built offline without execution. AWS outcomes were reported by the learner and were not independently repeated by the coach.

## Why the rejected writes mattered

PutItem can overwrite an existing complete primary key. The conditional create prevents that replacement. At the duplicate and stale-update checkpoints, only the specifically identified ConditionalCheckFailedException counted as an expected rejection. Follow-up strong reads checked that the original state was preserved.

The update required queued before changing to running. The later attempt still expected queued, so it failed after the state had changed. DeleteItem required running and returned the old item for verification.

A missing item and a missing table are distinct: GetItem can succeed with no Item, whereas a missing table raises ResourceNotFoundException, safely categorized as TableMissing.

## Safe execution and cleanup

The executable was built before the table existed. The source profile used temporary credentials and credential environment overrides were cleared in the execution subshell. Stored evidence contains only fictional data, check labels, known error categories, and exit statuses.

The budget check from earlier the same day was reused. The item and table were removed in the same lab session. No backup or additional resource was reported, and logout succeeded. Actual cost and study hours were not measured.

## Scope

This demonstrates a small deterministic conditional-write sequence, not multi-client concurrency correctness, production retry/idempotency, transactions, recovery, or IAM tenant isolation. The storage/serverless milestone remains incomplete.

- [Lab evidence](../../labs/week-06/day-26-dynamodb-conditional-crud/README.md)
- [Writer](../../labs/week-06/day-26-dynamodb-conditional-crud/main.go)
- [Fictional fixture](../../labs/week-06/day-26-dynamodb-conditional-crud/fixture.json)
- [Social drafts](../../posts/week-06/day-26-dynamodb-conditional-crud.md)
