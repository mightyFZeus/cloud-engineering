# Day 23 — Validate and Narrow an IAM Policy

## Read first — 10 minutes

1. [Validate policies with IAM Access Analyzer](https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-policy-validation.html) — read the console JSON-editor workflow and the four finding categories. Distinguish basic validation from paid custom checks.
2. [IAM policy validation check reference](https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-reference-policy-checks.html#access-analyzer-reference-policy-checks-error-invalid-action) — focus on Invalid action / INVALID_ACTION and why valid JSON can still contain an invalid AWS action.

**Date:** October 28, 2026
**Loaded:** October 27, 2026, after Day 22 passed
**Due:** October 28 at 11:00 PM Africa/Lagos
**Timebox:** 60–75 focused minutes, including reading and documentation
**Status:** Passed October 27, 2026, ahead of the due date — all six criteria supported by saved policies, learner validation results, corrected explanations, and cleanup confirmation.

## Outcome

Catch an action-name error using AWS basic policy validation, repair it, and reduce an overly broad policy to a concrete requirement: read known object keys only under a fictional reports/daily/ prefix. Explain what JSON parsing, semantic validation, simulation, and live testing each establish.

This completes a policy-authoring exercise in Week 5. Live least-privilege authorization remains a separate pending roadmap item.

## Numbered steps

1. Read the references. Sign into the IAM console with your everyday MFA-protected identity. Use the console workflow for this exercise.
2. Create labs/week-05/day-23-iam-policy-validation/README.md, broken.json, broad.json, and scoped.json under the project root /Users/mightyzeus/Desktop/cloud-eng. Preserve Day 22 artifacts. Suggested allocation: 10 minutes reading, 10 preparing files, 25 validating/comparing, 20 documenting, and 10 reviewing/cleaning up.
3. Put the following intentionally invalid AWS action in broken.json. It is valid JSON; s3:GetObjet is a deliberate spelling error. The fictional bucket is not created.

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "ReadDailyReports",
      "Effect": "Allow",
      "Action": "s3:GetObjet",
      "Resource": "*"
    }
  ]
}
```

4. Make broad.json the same policy with only the action corrected to s3:GetObject. Make scoped.json a copy of broad.json with only Resource changed to arn:aws:s3:::amzn-s3-demo-cloud-eng/reports/daily/*. Parse all three with python3 -m json.tool and record the actual exit statuses. Each should be 0. Notice that JSON syntax checking cannot detect the invalid AWS action name.
5. In IAM -> Policies -> Create policy -> JSON, paste broken.json. Inspect the basic validation pane below the editor. Record the Errors count, the Invalid action finding, and INVALID_ACTION if the code is displayed. Record counts for Security, Warnings, and Suggestions as displayed. If findings cannot load or access is denied, stop after 10 minutes and report the exact short blocker; no missing finding counts should be invented.
6. Replace the editor contents with broad.json. Record all four finding counts. The invalid-action error should disappear. Explain why a corrected action with Resource * still exceeds the daily-report requirement. A wildcard need not trigger a warning; record what AWS actually shows.
7. Replace the editor contents with scoped.json. Record the four counts and resolve any Errors without broadening its action or resource. Compare the allowed resource scope manually: s3:GetObject on reports/daily/*, with no upload, delete, or listing permission. Explain any remaining non-error finding, or explicitly record zero if the UI shows zero.
8. Save a three-row table: file, JSON exit status, Errors, Security, Warnings, Suggestions, notable finding/repair. Add the original requirement, before/after action/resource scope, and why basic validation passing is not proof of least privilege or successful live access. In your own words distinguish local JSON syntax checks, AWS semantic validation, policy simulation, and a real authorized API request.
9. Cancel the new-policy editor. Confirm no policy was saved or attached, no real bucket or IAM analyzer was created, no real data was accessed, and no paid custom check was run. Preserve the local files and run test -s on README.md, recording its exit status.

## Cost and security

AWS basic policy validation has no additional charge. Use only the Errors, Security, Warnings, and Suggestions checks. Check for new access and custom policy checks are separate paid features; do not run them for this assignment. Create no analyzer, policy, identity, bucket, or other AWS resource. Use a new unsaved editor, not an existing policy attached to a real identity. Cancel the editor when finished.

Save only fictional policy documents, finding categories/codes/counts, concise explanations, and cleanup status. Exclude account IDs, real ARNs, credentials, tokens, raw console records, and screenshots containing identifying details. No billable infrastructure is needed.

## Submission

1. Lab paths and all three JSON exit statuses:
2. Broken policy Errors count, finding name, and code if shown:
3. Broad policy: Errors / Security / Warnings / Suggestions:
4. Scoped policy: Errors / Security / Warnings / Suggestions:
5. Requirement and exact action/resource reduction:
6. One sentence each: JSON parsing, AWS validation, simulation, live request:
7. Remaining non-error findings and handling, or observed zero:
8. README test -s exit status:
9. Editor canceled; policies saved/attached, analyzers/resources created, paid custom checks run, or real data accessed:
10. Exact blocker, if any:

## Pass criteria

- Three valid JSON files with the intended action typo and only the controlled action/resource edits.
- Basic AWS validation identifies the broken action as an error; repaired broad and scoped versions have zero Errors.
- Actual finding counts are recorded without assuming that all overly broad policies generate warnings.
- Saved requirement and scope comparison reduce access from Resource * to only GetObject on the fictional daily-report prefix.
- Explanation distinguishes syntax, AWS validation, simulation, and live testing; zero findings is not treated as proof of least privilege or real authorization.
- Safe saved evidence, nonempty README, canceled editor, and no saved/attached AWS policy, created resource/analyzer, paid custom check, or real data access.
