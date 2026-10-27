# Day 24 — Live Least-Privilege Parameter Access

## References

- [IAM roles in the AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-role.html)
- [Parameter Store permissions](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-setting-up.html)

## Setup

- Region: eu-west-1; budget checked.
- Two Standard String parameters, both containing the non-secret value learning.
- Source profile: aws-learning, authenticated using aws login.
- Temporary role: cloud-eng-day24-reader.
- Trust limited to the source IAM user; identifying details kept private.
- Inline policy: ReadOneDay24Parameter.
- Permission: ssm:GetParameter on the allowed parameter only.
- Managed policies attached: zero.
- Role profile: day24-reader; requested session duration: 900 seconds.
- Assumed-role identity comparison: True; exit status: 0.
- Published permission template contains an account placeholder.

## Live results

| Profile | Parameter name | Result | Exit status |
|---|---|---|---|
| aws-learning | /cloud-eng/day24/allowed | Value matched learning | 0 |
| aws-learning | /cloud-eng/day24/denied | Value matched learning | 0 |
| day24-reader | /cloud-eng/day24/allowed | Value matched learning | 0 |
| day24-reader | /cloud-eng/day24/denied | AccessDeniedException | 254 |

Both parameters were confirmed to exist before testing the restricted role.

## Explanation

The trust policy controls who can assume the role. The permission policy controls what the role can do.

The assumed role uses its own scoped permissions; the source identity's administrative permissions are not automatically added.

AccessDenied indicates an unauthorized request. ParameterNotFound indicates the requested parameter is missing in the tested account and Region.

These results cover the tested GetParameter action and two parameter names.

## Cleanup verification

Both parameters were deleted through the console.

CLI post-delete checks encountered EndpointConnectionError with exit 255. An unauthenticated curl check received HTTP 400 with exit 0; this established an HTTP response, not authenticated API success.

The existing Day 20 Go reader then checked both names using aws-learning in eu-west-1:

| Parameter | Go result | Exit status |
|---|---|---|
| /cloud-eng/day24/allowed | ParameterNotFound | 1 |
| /cloud-eng/day24/denied | ParameterNotFound | 1 |

The temporary role was deleted and confirmed absent in the refreshed IAM console. Its inline policy was removed with the role.

The temporary local profile was removed. Source-profile logout exited 0. Logout is not proof that every cached STS session was revoked.

The CLI connectivity failure's underlying cause remains unconfirmed.


## Final verification and safe command summary

- Redacted JSON syntax check: exit 0.
- README presence check: exit 0.
- No long-term access keys or additional AWS resources created, as confirmed by the learner.
- Command summary: profile login, filtered identity/parameter checks, private clipboard preparation, private role-profile setup/removal, console deletion, unauthenticated HTTP probe, existing Go reader, source logout, JSON validation, and README check.
- Reviewed October 27, 2026: six criteria passed using saved artifacts and learner-reported results. CLI connectivity remains an unresolved limitation; Go verified parameter absence.
