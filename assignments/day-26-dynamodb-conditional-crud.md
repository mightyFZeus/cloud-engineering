# Day 26 — Conditional DynamoDB CRUD with Go

## Read first — 10 minutes

1. [DynamoDB condition expressions](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.ConditionExpressions.html) — focus on conditional puts and updates: DynamoDB evaluates a write condition against the item identified by the complete key.
2. [DynamoDB examples for Go v2](https://docs.aws.amazon.com/sdk-for-go/v2/developer-guide/go_dynamodb_code_examples.html) — focus only on PutItem, UpdateItem, DeleteItem, and GetItem. Reuse the small SDK/configuration helpers from Day 25 instead of implementing the full movie example.

**Date:** October 31, 2026
**Loaded:** October 30, 2026, after Day 25 verification
**Due:** October 31 at 11:00 PM Africa/Lagos
**Timebox:** 75–90 focused minutes, including cleanup and documentation
**Status:** Assigned for October 31; loaded October 30. No submission received.
**Roadmap section:** Week 6 — storage and databases

## Outcome

Extend the previous Go read-pattern work with create, guarded update, and delete operations on one fictional job. Prove that a duplicate create and stale update are rejected by their conditions, then distinguish a missing item from a deleted table. Use one small disposable table and temporary credentials.

This exercise is not a backup/restore, general pagination, concurrent-client load test, or IAM tenant-isolation test. It adds conditional item writes to the earlier GetItem/Query work.

## Numbered steps

1. Complete the reading. Write the required state sequence in the new README before creating anything: absent -> queued -> running -> absent. Record why a duplicate create and an update expecting the old queued state should fail after the appropriate state has changed.
2. Prepare labs/week-06/day-26-dynamodb-conditional-crud/README.md, main.go, fixture.json, and Go module files under /Users/mightyzeus/Desktop/cloud-eng. The fixture is only PK=USER#demo-a, SK=JOB#001, jobStatus=queued. Preserve Day 25 files. Suggested allocation: 10 minutes reading, 5 design/setup, 30 guided Go changes/build, 10 table creation, 10 execution, and 25 cleanup/notes.
3. Reuse the Day 25 SDK/default-credential pattern: config.LoadDefaultConfig, explicit eu-west-1, externally selected AWS_PROFILE=aws-learning, a -table flag defaulting to cloud-eng-day26-jobs, bounded configuration/API contexts, and safe error categories. Check the budget and renew aws-learning using aws login with the explicit profile and Region. If the browser HTTP 400 returns, a newly generated --remote link in a fresh private browser window was the successful Day 25 recovery path; its original root cause remains unproven. Store no credentials or private AWS configuration in the lab. Clear credential environment overrides in the execution subshell as in Day 25.
4. Implement this fixed sequence in Go using the complete PK/SK key and expression aliases/placeholders:
   - PutItem creates the queued fixture only when attribute_not_exists(#pk) AND attribute_not_exists(#sk). A conditional put checks the single item identified by both key values; it does not ban every other job sharing that user's partition.
   - A strongly consistent GetItem checks the exact String PK, SK, and queued status.
   - Repeat the same conditional PutItem with jobStatus=unexpected. Require ConditionalCheckFailedException, then strongly consistently read the item and prove it remains queued. Never count a timeout, credential failure, AccessDenied, or any arbitrary nonzero result as a successful condition guard.
   - UpdateItem uses SET #status = :next with #status = :expected, where expected is queued and next is running. Use ReturnValues ALL_NEW and verify the exact key/status in the returned Attributes. The condition also prevents an accidental upsert when the expected queued item does not exist.
   - Attempt a stale UpdateItem using expected=queued and next=complete. Require ConditionalCheckFailedException; a strongly consistent GetItem must still show running. This is a deterministic stale-state check, not a multi-client concurrency test.
   - DeleteItem removes that one item. Use ReturnValues ALL_OLD and verify that the returned key/status describe the running fixture; do not treat a successful delete of an already absent item as evidence that this fixture was deleted.
   - A strongly consistent GetItem must return no Item. Missing items produce a successful GetItem response without item data; they do not produce ResourceNotFoundException.
5. Print concise labels only after checking the corresponding evidence: create check: PASS; duplicate guard: PASS; update check: PASS; stale guard: PASS; delete check: PASS; absent item check: PASS. Exit 0 only if every check succeeds. Only the duplicate/stale checkpoints treat the specifically identified ConditionalCheckFailedException as expected. All other unexpected failures must exit nonzero with a safe category. Never print raw errors or complete responses.
6. Run gofmt and build before creating the table. In the console, create only cloud-eng-day26-jobs with PK String and SK String, on-demand capacity, Standard class, learning tags, AWS-owned encryption, and deletion protection off. If that exact table already exists, stop before changing it. Keep the table empty initially. Create no index, replica, stream, backup/PITR feature, integration, or customer-managed KMS key. Wait for Active.
7. Run the program once with the temporary source profile. Record the six actual PASS labels and overall exit status. If a run stops partway through, do not overwrite the remaining item or remove its conditions to obtain PASS. Record the stopping category; inspect the disposable table state privately before any controlled retry. Bound troubleshooting to 10 minutes and prioritize cleanup if blocked.
8. Explain in your own words: why an unconditional PutItem can overwrite an item; what ConditionalCheckFailedException means at the two expected checkpoints; why the update checks the prior state; and how item absence differs from table absence. Save a compact operation/expected/actual table and note that the duplicate/stale attempts did not modify the item.
9. Delete only this temporary table through the console, without a final backup. Wait until the table is absent. Use the existing Day 25 read-only reader with its -table flag pointed at cloud-eng-day26-jobs to verify the deletion. Require TableMissing / ResourceNotFoundException and nonzero exit; UnexpectedItem, AccessDenied, a timeout, or a network error is not table-absence proof. This cleanup verification is a read-only GetItem attempt, not another run of the write sequence.
10. Log out aws-learning. Preserve the reviewed code, module files, fictional fixture, summarized results, explanations, table-deletion confirmation, post-delete category/exit status, and logout status. Run test -s on README.md and record the exit status. Inspect the files for secrets, real ARNs/account IDs, credentials, private configuration, and raw console/API output before submission.

## Cost and security

DynamoDB charges for requests and storage; on-demand is pay-per-request, not a guarantee of free usage. Check the account budget and [current pricing](https://aws.amazon.com/dynamodb/pricing/). A condition-rejected write can still consume capacity, so keep the exercise to the specified single fixture and a few calls. Delete the table in the same session, including when a checkpoint blocks.

Use temporary source credentials through the default SDK chain. Create no EC2 instance, new role, access key, backup, reserved commitment, index, replica, or other supporting resource. The source profile's authority is not claimed to be limited to the fictional user's partition. Save only safe labels, categories, exit statuses, and fictional data. No real report/job data is needed.

## Submission

1. README/code/module/fixture paths; planned state sequence:
2. Region, budget check, table PK/SK types, on-demand Standard settings, and initial empty/Active state:
3. gofmt/build result:
4. Six actual check labels and overall run exit status:
5. Duplicate and stale-update exception category; preserved queued/running states:
6. Put overwrite, expected-state update, missing-item versus missing-table explanations:
7. Table deleted/absent; read-only SDK post-delete category and nonzero exit:
8. Source logout result; other resources or features created:
9. README test -s exit status:
10. Exact blocker, if any:

## Pass criteria

- One intended on-demand Standard table with the expected String composite key and one fictional job; no unplanned resources/features.
- Reviewed Go uses the default credential chain, bounded calls, full keys, condition expressions/placeholders, and safe output; formatting/build succeed.
- Conditional create succeeds; the duplicate produces ConditionalCheckFailedException and the queued item is preserved.
- The guarded queued-to-running update succeeds; the stale queued-state update is condition-rejected and running state is preserved.
- DeleteItem evidence identifies the removed fixture; a strongly consistent GetItem then confirms an absent item without confusing it with a missing table; all six content/guard checks pass with overall exit 0.
- Explanations and substantive safe files are saved; table absence is confirmed by console and a read-only SDK TableMissing result, source logout succeeds, and no unintended resources or credentials remain.
