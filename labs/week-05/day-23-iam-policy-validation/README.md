# Day 23 — IAM Policy Validation

## References

- [AWS policy validation](https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-policy-validation.html)
- [Policy check reference](https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-reference-policy-checks.html)

## Requirement

Permit reading known object keys only under the fictional reports/daily/ prefix. No upload, deletion, or listing permission is required.

## Observed validation results

Basic IAM Access Analyzer validation was used in an unsaved IAM JSON policy editor.

| File | JSON exit status | Errors | Security | Warnings | Suggestions | Finding |
|---|---|---|---|---|---|---|
| broken.json | 0 | 1 | 0 | 0 | 0 | INVALID_ACTION: s3:GetObjet |
| broad.json | 0 | 0 | 0 | 0 | 0 | None |
| scoped.json | 0 | 0 | 0 | 0 | 0 | None |

## Repair and scope reduction

- Broken policy: misspelled action s3:GetObjet.
- Broad policy: corrected action to s3:GetObject, retaining Resource "*".
- Scoped policy: retained s3:GetObject and restricted Resource to arn:aws:s3:::amzn-s3-demo-cloud-eng/reports/daily/*.
- Resource "*" exceeded the daily-report requirement by allowing the statement to apply to S3 objects outside that prefix.
- Zero findings on the broad policy did not prove least privilege.
- Live S3 authorization was not tested.

## What each check establishes

- JSON parsing checks whether the JSON syntax is correctly written.
- AWS policy validation checks policy grammar, valid AWS elements/actions, and defined security/best-practice concerns.
- Policy simulation evaluates whether supplied policies allow or deny an action on a resource with supplied context.
- A real GetObject request tests whether object retrieval actually works using a real identity and resource.
- Validation findings do not replace requirement review, effective-permission analysis, or live testing.

## Cleanup

- Unsaved editor canceled.
- No AWS policy saved or attached.
- No analyzer or AWS resource created.
- No paid custom policy check run.
- No real S3 data accessed.


## Final verification

- All three JSON syntax checks: exit 0.
- README presence check: exit 0.
- Reviewed October 27, 2026: all six assignment criteria supported by saved artifacts and learner-reported AWS validation findings. Live S3 authorization was not tested.
