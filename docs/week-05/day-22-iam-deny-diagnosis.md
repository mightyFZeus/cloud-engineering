# Day 22 — Diagnose an explicit IAM deny

## References

- [IAM policy evaluation logic](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html)
- [How to simulate policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/policies_policy-simulator-how-to.html)

Passed October 27, 2026: all six assignment criteria supported by learner-reported simulator checkpoints, explanations, and review of the saved policies and README. The coach did not independently execute AWS simulations.

## Verified evidence

- Both policy files parse as JSON. The saved baseline matches the Day 19 policy exactly.
- The restricted policy is precisely the baseline plus DenyRestrictedReports: Deny s3:GetObject for the fictional reports/restricted/* object prefix.
- One identity-based policy selected in Custom mode for each reported test.

| Policy | Object key | Actual reported result | Relevant evidence |
|---|---|---|---|
| Baseline | reports/restricted/sample.txt | Allowed | Explicit allow in one statement |
| Restricted | reports/restricted/sample.txt | Denied | Explicit deny; matched DenyRestrictedReports |
| Restricted | reports/sample.txt | Allowed | Explicit allow in one statement |
| Restored baseline | reports/restricted/sample.txt | Allowed | Explicit allow in one statement |

- Learner explained explicit deny precedence, the unaffected ordinary report, and explicit versus implicit deny.
- Scenario analysis identifies assets, actors, the application request boundary, and three risk/control pairs: credential exposure, excess prefix access, and public exposure. Controls are proposed, not deployed.
- Learner reported baseline JSON, restricted JSON, and README checks all exited 0. Saved contents were inspected; a nonempty README alone would not demonstrate the simulations.

## Interpretation and limits

The GetObject statements have no condition requirement. The separate ListBucket allow uses s3:prefix. The narrow content deny does not remove listing permission and does not establish that restricted object names are hidden.

An intentional deny in a real environment should be preserved while approved access is discussed with its policy owner. Restoring the baseline was a controlled simulator reversal.

Simulation evaluates supplied policies and context. Other applicable policies and conditions can affect live requests. This exercise does not verify real S3 authorization, deployed mitigations, or a repository-wide security audit.

The simulator UI initially failed to load services. On October 27 the learner confirmed S3 actions loaded without an error and completed the tests. The original loading failure's root cause was not established.

## Cleanup and artifacts

No real identity policy was attached, no AWS resource was created, and no real security control was changed, as confirmed by the learner. No billable cleanup was required.

- [Lab evidence](../../labs/week-05/day-22-iam-deny-diagnosis/README.md)
- [Baseline policy](../../labs/week-05/day-22-iam-deny-diagnosis/baseline.json)
- [Restricted policy](../../labs/week-05/day-22-iam-deny-diagnosis/restricted.json)
- [Social post drafts](../../posts/week-05/day-22-iam-deny-diagnosis.md)

Safe command summary: local JSON parsing, clipboard copy, README appending, and file-presence checks. No raw AWS response, credential material, actual account identifiers, or real resource ARNs were saved.
