# Day 22 — Explicit IAM Deny Diagnosis

## References

- [IAM evaluation logic](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html)
- [Policy simulator guide](https://docs.aws.amazon.com/IAM/latest/UserGuide/policies_policy-simulator-how-to.html)

## Simulation scope

AWS IAM Policy Simulator in Custom mode, with one identity-based policy selected.
Resources used fictional S3 bucket and object names.

## Observed results

| Policy | Action | Object key | Actual result | Result details | Matched deny Sid |
|---|---|---|---|---|---|
| Baseline | s3:GetObject | reports/restricted/sample.txt | Allowed | Explicit allow in one statement | Not applicable |
| Restricted | s3:GetObject | reports/restricted/sample.txt | Denied | Explicit deny in one statement | DenyRestrictedReports |
| Restricted | s3:GetObject | reports/sample.txt | Allowed | Explicit allow in one statement | Not applicable |
| Restored baseline | s3:GetObject | reports/restricted/sample.txt | Allowed | Explicit allow in one statement | Not applicable |

## JSON validation

- Baseline JSON validation exit status: 0.
- Restricted JSON validation exit status: 0.

## Access-denial diagnosis

- Principal scope: fictional identity-based policies in Custom simulation mode.
- Action: s3:GetObject.
- Allowed scope: reports/*.
- Denied scope in the restricted policy: reports/restricted/*.
- Matched deny statement: DenyRestrictedReports.
- Root cause: the restricted object matched an explicit deny, which overrides allow.
- The ordinary report remained allowed because it was outside the denied prefix.
- Restoring the baseline restored simulated access to the restricted object.
- Implicit deny means no applicable allow grants access; explicit deny means a matching Deny statement blocks access.

## Scenario threat model

- Assets: report contents and AWS credentials.
- Actors: an authorized reader and someone attempting unauthorized access.
- Entry point: an application making an authenticated S3 API request.
- Trust boundary: AWS evaluates the request against applicable permissions.

| Risk | Proposed mitigation |
|---|---|
| Exposed AWS credentials | Use temporary credentials and revoke exposed credentials or sessions through the appropriate identity mechanism. |
| Access to unintended report folders | Restrict IAM permissions to the required prefix. |
| Public report exposure | Enable S3 Block Public Access and avoid public bucket policies. |

These mitigations are proposed controls, not controls implemented on a real bucket during this lab.

The GetObject deny restricts object contents. It does not remove the baseline ListBucket permission, so restricted object names may still appear in permitted listings.

## Simulator limitations and approved access

An Allowed simulation does not guarantee a successful real request: other policies and conditions may apply.

An intentional real deny should remain in place. Ask the policy owner about approved access rather than removing the restriction merely to make a request succeed.

## Cleanup and scope

- No policy was attached to a real identity.
- No AWS resources were created.
- No real security controls were changed.
- No billable resources require cleanup from this lab.


## Final verification

- Baseline JSON check: exit 0.
- Restricted JSON check: exit 0.
- README presence check: exit 0.
- Reviewed October 27, 2026: six assignment criteria passed using saved artifacts and learner-reported console results. Live S3 authorization was not tested.
