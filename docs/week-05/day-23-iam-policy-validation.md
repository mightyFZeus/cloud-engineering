# Day 23 — Validate and narrow an IAM policy

## References

- [AWS basic policy validation](https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-policy-validation.html)
- [Policy check reference](https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-reference-policy-checks.html)

Passed October 27, 2026, ahead of the October 28 due date: all six criteria supported by learner-reported AWS editor findings, corrected explanations, cleanup confirmation, and review of the saved policy files and README. No independent AWS validation request was made by the coach.

## Verified evidence

Requirement: read known object keys only in a fictional reports/daily/ prefix; no listing, upload, or deletion permission required.

| File | JSON exit | Errors | Security | Warnings | Suggestions | Observed finding |
|---|---|---|---|---|---|---|
| broken.json | 0 | 1 | 0 | 0 | 0 | INVALID_ACTION: s3:GetObjet |
| broad.json | 0 | 0 | 0 | 0 | 0 | None |
| scoped.json | 0 | 0 | 0 | 0 | 0 | None |

Saved-file inspection confirms that broken to broad changes only the action from s3:GetObjet to s3:GetObject. Broad to scoped changes only Resource from * to the fictional reports/daily/* object prefix. All three files parse as JSON.

## Interpretation

The action typo passed local JSON syntax checking but failed AWS action validation. Both repaired variants produced zero findings in the learner's console session. Zero findings did not establish that broad Resource * matched the daily-report requirement; the scope was reduced after reviewing that requirement.

The learner initially confused validation categories with simulation and then corrected the explanation: simulation evaluates whether supplied policies allow or deny an action on a resource; an actual API request tests the operation with a real identity and resource. The saved README includes the corrected distinction among syntax, semantic validation, simulation, and live operations.

The scoped policy's GetObject statement does not grant listing, upload, or deletion. No claim is made about effective live access or all possible permissions from other policies.

## Cleanup and review

- README presence check reported exit 0; its substantive content was inspected.
- Unsaved editor canceled.
- No AWS policy saved or attached; no analyzer or AWS resource created.
- No paid custom check run and no real S3 data accessed.
- Only fictional resource strings, sanitized finding counts/codes, explanations, and cleanup declarations retained.

## Artifacts

- [Lab evidence](../../labs/week-05/day-23-iam-policy-validation/README.md)
- [Broken policy](../../labs/week-05/day-23-iam-policy-validation/broken.json)
- [Broad policy](../../labs/week-05/day-23-iam-policy-validation/broad.json)
- [Scoped policy](../../labs/week-05/day-23-iam-policy-validation/scoped.json)
- [Social drafts](../../posts/week-05/day-23-iam-policy-validation.md)

Safe command summary: local JSON generation/parsing, clipboard copy, README appending, and file-presence checking. Raw AWS responses and identifying details were not retained. Live least-privilege authorization remains pending for a separate lab.
