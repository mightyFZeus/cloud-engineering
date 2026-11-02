# Day 28 — Recover an Overwritten S3 Object with Versioning

## Read first — 10 minutes

1. [Restoring previous S3 versions](https://docs.aws.amazon.com/AmazonS3/latest/userguide/RestoringPreviousVersions.html) — focus on making a previous version current, the difference between copying and permanently deleting a current version, and the console Versions workflow.
2. [S3 examples for Go v2](https://docs.aws.amazon.com/sdk-for-go/v2/developer-guide/go_s3_code_examples.html) — focus on PutObject and GetObject; reuse the small configuration, bounded-read, and safe-output helpers from Day 27. Replace example raw-error/resource logging with fixed categories.

**Date:** November 3, 2026
**Loaded:** November 2, 2026, after Day 27 verification
**Due:** November 3 at 11:00 PM Africa/Lagos
**Timebox:** 75–90 focused minutes, including the initial versioning propagation interval, cleanup, and notes
**Status:** Assigned for November 3; loaded November 2. No submission received.
**Roadmap section:** Week 6 — storage and databases

## Outcome

Create two versions of one fictional object, prove that the replacement is current, then recover the original as the current object. Use Go for the two writes and exact-byte reads; use the console to permanently delete only the replacement's current version. Finish by deleting every remaining version/delete marker and the bucket.

The console rollback is deliberately limited to a disposable lab fixture. Copying an earlier version into a new current version preserves all versions, but that in-place workflow requires CLI/SDK rather than the console. This lesson uses the simpler, documented console rollback so its first-enablement waiting period and full cleanup fit within 90 minutes. It demonstrates one small overwrite recovery, not an independent backup, disaster recovery readiness, a measured RPO/RTO, or a completed serverless milestone.

## Numbered steps

1. Complete the reading. Write the expected state sequence in README: original bytes current -> replacement bytes current while original remains noncurrent -> original bytes current after permanently deleting exactly the replacement version. Explain why ordinary object deletion and version-specific permanent deletion differ.
2. Prepare labs/week-06/day-28-s3-version-recovery/README.md, main.go, original.txt, replacement.txt, and Go module files under /Users/mightyzeus/Desktop/cloud-eng. Use key reports/day28/config.txt. The exact fictional payloads are day28 original config followed by a newline, and day28 overwritten config followed by a newline. Reuse Day 27's SDK dependencies/helpers. Suggested allocation: 10 minutes reading, 25 guided code/build, 10 bucket setup, 15 propagation plus explanatory notes, 10 seed/recover/verify, and 15 cleanup/final notes.
3. Implement explicit, separate modes before creating resources:
   - Default mode verify is read-only. A bounded GetObject without VersionId compares the entire current object's bytes with the selected known payload (original or replacement), including the newline, and prints only a fixed check label/boolean. Selecting an expected payload must not write/reseed anything.
   - Mode seed writes the original payload, confirms its exact current bytes, writes the replacement to the same key, and confirms the exact replacement is now current. Inspect the two PutObject VersionId results only in memory: require both nonempty, non-null, and distinct. Print only two distinct versions: true, original seed check: PASS, and replacement current check: PASS after those assertions. Do not print or save either Version ID.
   - Mode check-bucket is read-only SDK ListObjectsV2 with MaxKeys=1. Exact NoSuchBucket becomes BucketMissing with nonzero exit. A successful listing means BucketPresent even if empty; other failures are not absence proof.
   - Use config.LoadDefaultConfig, explicit eu-west-1, external AWS_PROFILE=aws-learning, bounded configuration/API contexts, bounded body reads, closed response bodies, and safe fixed error categories. Read the actual new bucket name only from CLOUD_ENG_DAY28_BUCKET; use `<PRIVATE_LAB_BUCKET>` in notes. Suppress SDK diagnostics that could expose private names or IDs. Create no presigned URL or automatic lifecycle operation in this program.
4. Run gofmt and build before creating the bucket. Check eu-west-1 and the budget, and renew the normal temporary source login. Privately choose a new globally unique learning bucket name with a random suffix and no account ID/personal information; do not modify an existing bucket if the name is unavailable. Set the private bucket environment variable locally without sharing or saving its value.
5. In the console, create exactly one general purpose bucket in eu-west-1, shared global namespace, all four Block Public Access settings enabled, Bucket owner enforced/ACLs disabled, SSE-S3, versioning ENABLED, Object Lock disabled, and learning-project tags. Keep it initially empty. Add no bucket/IAM policy, KMS key, CORS/website feature, replication, lifecycle rule, backup, logging/inventory/analytics integration, or supporting resource. Use only S3 Standard for both tiny versions.
6. After first enabling versioning, wait 15 minutes before any object PUT/DELETE, following [AWS's first-enablement guidance](https://docs.aws.amazon.com/AmazonS3/latest/userguide/manage-versioning-examples.html). Use this interval to write the recovery-limit and cleanup explanations; it is included in the timebox. Newly enabled versioning may otherwise produce intermittent NoSuchKey visibility errors. Do not repeatedly seed to work around the propagation interval.
7. After the interval, run seed mode once with the source profile and credential environment overrides cleared. Require its exact-byte checks, two-distinct-versions boolean, and exit 0. Inspect the one object in the console's Versions view: expect two versions and identify the replacement as the current/latest version. Record only the count and current/noncurrent relationship. Keep actual bucket names, Version IDs, raw API output, and screenshots containing them private. If a seed stops partway through, inspect the disposable fixture privately before a controlled retry rather than blindly writing extra versions.
8. In the object's Versions view, select only the current/latest replacement version. Follow the version-specific Delete workflow and its permanently delete confirmation. Do not use an ordinary object delete that creates a delete marker, and do not delete the original noncurrent version. This permanently removes the replacement fixture; that deleted version cannot be recovered by versioning. The retained original should become current. If the UI does not clearly identify the intended current version, stop and report a short description without exposing Version IDs.
9. Run read-only verify mode expecting the original payload, with no VersionId in GetObject. Require an exact original-byte match, restored current check: PASS, and exit 0. A version-specific read of the old payload alone would not prove that the original was restored as current. Record the actual recovery method and result, and explain that copying an earlier version would preserve the replacement whereas this chosen lab rollback permanently deleted it. Bound unexpected troubleshooting to 10 minutes and prioritize cleanup if blocked.
10. In the console, empty this lab bucket completely, including all remaining object versions and any delete markers, then verify the Versions view has no remaining versions/markers. Delete the bucket and confirm its disappearance. Run only the read-only bucket-check mode; require BucketMissing / NoSuchBucket and nonzero exit. An empty current-object view, NoSuchKey, AccessDenied, generic 404, or connection error does not establish complete cleanup. Unset CLOUD_ENG_DAY28_BUCKET and log out aws-learning. Complete cleanup even if seeding or recovery blocks.
11. Save safe code/module files, the fictional payloads, settings confirmations with the placeholder, actual check summaries/exit statuses, version counts only, explanations, and cleanup/logout status. Explain versioning's limits: permanently deleted versions are not restored by this method, retention can be affected by authorized deletion/lifecycle, and keeping versions in one account/bucket is not an independent backup. Do not invent recovery-time/loss targets or cost figures from this single fixture. Run test -s on README.md and record the exit status; inspect artifacts for credentials, URLs, actual bucket names/ARNs, Version IDs, account IDs, raw errors/responses, or private configuration.

## Cost and security

S3 charges can include requests, storage for each retained version, and data transfer. Check the budget and [current pricing](https://aws.amazon.com/s3/pricing/). Keep this to two tiny S3 Standard versions of one object and a few reads. Object versions continue using storage until permanently removed; deleting only the visible current object can leave charged versions and delete markers behind. Empty every version/marker and delete the bucket in the same session.

Use temporary credentials through the SDK default chain. No actual bucket name or Version ID belongs in repository files or chat; use placeholders and booleans/counts. Create no IAM role/policy, access key, KMS key, Object Lock retention, replica, backup, lifecycle rule, or supporting resource. Source logout is session hygiene, not evidence of a comprehensive credential-revocation procedure.

## Submission

1. Lab paths, expected state sequence, gofmt/build exit status:
2. Region/budget, one initially empty new bucket, safe private settings, Versioning Enabled, and initial 15-minute propagation interval observed:
3. Two distinct non-null Version IDs confirmed privately; only a boolean/count submitted:
4. Original seed and replacement-current exact-byte checks; seed exit status:
5. Recovery method: latest replacement version permanently deleted; original retained and current:
6. Read-only current-object restoration check, exact original-byte result, and exit status:
7. Ordinary deletion versus version-specific deletion, copy-versus-delete rollback, and versioning/backup limits:
8. All versions/delete markers removed; bucket empty/deleted/absent; read-only SDK category and nonzero exit:
9. Environment variable unset; source logout result; additional resources/features created:
10. README test -s exit status and exact safe blocker, if any:

## Pass criteria

- One intended private general purpose bucket with versioning enabled, the specified settings, and the initial propagation interval observed; no unplanned resources/features.
- Reviewed Go has separated write/read-only modes, temporary default credentials, explicit Region, bounded operations, exact payload checks, and no sensitive output; formatting/build pass.
- Two distinct non-null versions are established and the original/replacement current-byte checks pass with exit 0, without recording actual Version IDs.
- The specific replacement version is permanently deleted through the documented console rollback; a subsequent read-only GetObject without VersionId proves the original bytes are current and exits 0.
- Learner explains delete markers, permanent-version deletion, preservation by copying, and why one small overwrite recovery is not independent backup/disaster recovery evidence or an invented RPO/RTO result.
- Complete version/delete-marker removal, bucket deletion confirmed by console and SDK NoSuchBucket, environment cleanup/logout, and substantive safe evidence files are verified.
