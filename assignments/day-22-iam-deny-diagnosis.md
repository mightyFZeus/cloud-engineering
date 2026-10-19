# Day 22 — Diagnose an Explicit IAM Deny in Simulation

**Date:** October 19, 2026
**Due:** 11:00 PM Africa/Lagos
**Timebox:** 60–75 minutes
**Status:** Assigned

## Outcome

Demonstrate an explicit deny overriding an allow, diagnose the matched policy statement, and write a small security analysis for the fictional report-reader scenario. Simulation only; no real permissions changed.

## Steps

1. Read the two references below for 10 minutes. Create labs/week-05/day-22-iam-deny-diagnosis/README.md and two new policy files baseline.json and restricted.json. Preserve the original Day 19 policy unchanged.
2. Copy Day 19 policy.json into baseline.json. In the IAM console Policy simulator -> Custom mode, include only the baseline policy, no permissions boundary. Add GetObject with exact fictional object resource arn:aws:s3:::amzn-s3-demo-cloud-eng/reports/restricted/sample.txt. Simulate and record the actual result; expect allowed.
3. Create restricted.json as a copy of the baseline plus one statement: Sid DenyRestrictedReports, Effect Deny, Action s3:GetObject, Resource arn:aws:s3:::amzn-s3-demo-cloud-eng/reports/restricted/*. Validate both files with python3 -m json.tool; record exit statuses. This deny restricts GetObject content access only. The baseline ListBucket permission still allows reports/ listings, which can reveal restricted key names; do not claim the names are hidden.
4. Replace the simulator policy with restricted.json. Do not leave both policies loaded accidentally. Keep the exact same action and restricted object resource; simulate and record the decision and matched deny statement Sid. Expect explicit deny.
5. With the restricted policy, change only the object resource to arn:aws:s3:::amzn-s3-demo-cloud-eng/reports/sample.txt. Expect allowed. This checks that the narrow deny does not block other report objects.
6. Restore the baseline policy in the simulator and test reports/restricted/sample.txt again. Expect allowed. This is a lab reversal, not a recommendation to remove an intentional security restriction from a real environment.
7. Save a four-row result table with policy version, action, object key, actual decision, and relevant matched Sid if shown. Write a simulated access-denial troubleshooting note: principal scope (fictional custom identity policy), action, resource pattern, context, applicable statements, root cause, and observed lab reversal. Contrast an explicit deny with Day 19's absence of a matching allow.
8. Add a small design threat model for the fictional S3 report-reader scenario: assets, actors, entry points/trust boundaries, and three concrete risks with mitigations. Consider exposed credentials, unintended prefix access, and public data exposure. Clearly label proposed controls versus verified simulations. This does not claim a repository-wide security audit.
9. Explain simulator limitations and why a genuine intentional deny must not be removed merely to make a request succeed. Confirm no identity policy was attached, no real bucket was accessed, and no AWS resource or organization control was changed. Run test -s on README.md and record exit status.

## References

- [IAM policy evaluation logic](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html)
- [How to simulate policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/policies_policy-simulator-how-to.html)

## Cost and security

Use fictional S3 resource strings and Custom simulator mode only. Create no bucket, role, or attached policy. Do not change a real explicit deny, permissions boundary, or organization policy. Publish no account IDs, real ARNs, identity details, credentials, or raw console records.

## Submission

1. README, baseline.json, and restricted.json paths:
2. Both JSON validation exit statuses:
3. Baseline restricted-object result:
4. Restricted-policy restricted-object result and matched deny Sid:
5. Restricted-policy ordinary-report result:
6. Restored-baseline restricted-object result:
7. Root cause and explicit-versus-implicit deny explanation:
8. Threat-model risks/mitigations count and one example:
9. Simulator limitations and intentional-deny handling:
10. test -s exit status:
11. Real policies attached, resources created, or controls changed:
12. Blocker:

## Pass criteria

- Baseline preserved; new policy has the exact narrow explicit deny and valid JSON.
- Four simulations demonstrate allow, explicit deny, allowed ordinary report, and restored allow.
- Deny result is tied to the matching statement and correct action/resource.
- Troubleshooting note distinguishes explicit deny from missing allow and states safe remediation limits.
- Scenario threat model includes assets, actors, entry points/boundaries, and three grounded risk/mitigation pairs.
- Safe saved evidence, simulator limitations, no changes to real policies or resources.
