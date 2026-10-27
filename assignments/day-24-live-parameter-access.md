# Day 24 — Test Live Least-Privilege Parameter Access

## Read first — 10 minutes

1. [IAM roles in the AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-role.html) — focus on role permissions versus role trust, source_profile, role_session_name, and temporary credentials. Use the existing aws login profile; skip the static-access-key examples.
2. [Parameter Store setup and access](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-setting-up.html) — focus on the exact ssm:GetParameter action and parameter resource scope. Other retrieval APIs have separate permissions.

**Date:** October 28, 2026
**Loaded:** October 27, 2026, after Day 23 passed early
**Due:** October 28 at 11:00 PM Africa/Lagos
**Timebox:** 75–90 focused minutes including reading, documentation, and cleanup
**Status:** Passed October 27, 2026, ahead of the due date — six criteria supported by live role tests, saved redacted artifacts, explanations, Go-based absence checks, and cleanup confirmation.

## Outcome

Use one temporary assumed role to read one existing non-secret Standard String parameter and receive AccessDeniedException for a second existing parameter. Verify which role made the requests and distinguish lack of permission from a nonexistent parameter. This is a live SSM authorization lab; it does not test S3 prefix access.

## Verification note

The CLI performed the live role identity and allow/deny tests. Later CLI cleanup reads encountered connection errors. The reviewed Day 20 Go reader used the source profile to confirm ParameterNotFound for both deleted parameters. This verified the required absence without claiming the CLI issue was fixed. Role deletion was confirmed through the console, the temporary local profile was removed, and source logout exited 0.

## Numbered steps

1. Read the references. Confirm eu-west-1 and your budget. Reserve 15 minutes for cleanup. Use the existing everyday MFA-protected console identity and aws-learning CLI profile. Renew with aws login --profile aws-learning. Never create long-term keys or export temporary access keys/session tokens.
2. Create labs/week-05/day-24-live-parameter-access/README.md and a redacted policy-template.json under /Users/mightyzeus/Desktop/cloud-eng. Keep all real identity/resource ARNs only in the AWS console and private AWS configuration outside the repository. Do not paste them into chat. Suggested allocation: 10 minutes reading, 10 parameters, 20 role/configuration, 15 tests, 15 notes, and 15 cleanup.
3. Using the source identity in the console, create exactly two parameters: /cloud-eng/day24/allowed and /cloud-eng/day24/denied. For both choose Standard tier, String type, and the benign value learning. Use standard throughput and project tags where available. If either name already exists, stop and report the collision rather than overwriting unrelated data. Confirm both exist and the source identity can retrieve both; discard identity output and record only success booleans/categories/exit statuses.
4. In IAM, create the temporary role cloud-eng-day24-reader. Use a custom trust policy with only the current learner IAM user's exact principal and sts:AssumeRole, entered privately in the console. The existing source identity must be authorized to assume this role. If the source identity is no longer an IAM user, report that before creating the role so the trust can be adapted. Attach no managed permission policy. Add one inline policy named ReadOneDay24Parameter, containing only an Allow for ssm:GetParameter on the exact allowed parameter ARN. Do not grant GetParameters, GetParametersByPath, GetParameterHistory, DescribeParameters, writes, or wildcard resources.
5. Save only this redacted permission template in the repository. Resolve the account placeholder privately in the AWS editor; keep the published template redacted. Do not save the real trust policy or private profile in the repository.

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "ReadOneDay24Parameter",
      "Effect": "Allow",
      "Action": "ssm:GetParameter",
      "Resource": "arn:aws:ssm:eu-west-1:<ACCOUNT_ID>:parameter/cloud-eng/day24/allowed"
    }
  ]
}
```

6. In the private AWS config, add only a new profile day24-reader with source_profile=aws-learning, the temporary role's real role_arn, role_session_name=cloud-eng-day24, and region=eu-west-1. Configure duration_seconds=900 for a short requested session. Keep the existing aws-learning profile intact. Use the CLI for this lab; no Go code change or credential export is needed. Clear environment credential overrides only within the lab command subshell, as in Day 20.
7. Verify the role without displaying its ARN. Require True from the comparison below, as well as successful command execution. A zero exit status with False is not the intended identity.

```bash
aws sts get-caller-identity --profile day24-reader \
  --query "ends_with(Arn, ':assumed-role/cloud-eng-day24-reader/cloud-eng-day24')" \
  --output text
```

8. Using only --profile day24-reader and eu-west-1, run GetParameter for /cloud-eng/day24/allowed. Query a boolean comparison of Parameter.Value with learning, or discard the response body. Require the expected value check to be True and exit 0. Then GetParameter /cloud-eng/day24/denied, which was already confirmed to exist. Require AccessDeniedException (or equivalent AccessDenied category) and nonzero exit; ParameterNotFound is not a passing denial test. Capture CLI output privately and save only the boolean, error category, and exit status; raw errors can contain real identifiers. The walkthrough will provide a small output-filtering command rather than asking for raw errors.
9. Save a table comparing the source identity's existence checks and the assumed role's allowed/denied results. Explain trust versus permissions, which identity made the requests, and why the source identity's administrative permissions were not added to the role's service permissions. Limit the claim to the exact action/resources/context tested. If assumption or allowed access fails, check profile selection, trust, the exact action/resource, and propagation; stop after 10 minutes of repeated failure and report a sanitized blocker instead of broadening permissions.
10. As the source identity, delete both Day 24 parameters and confirm each returns ParameterNotFound afterward. Delete the role's inline policy and the temporary role. Remove only the added day24-reader profile section; preserve other configuration and credential caches. Log out aws-learning. Do not treat source logout as proof that every cached STS session is revoked. Confirm both parameters and the role are absent, with no lab policy remaining. Even if testing blocks, clean up every resource created before reporting the blocker.
11. Record safe command summaries, exact test outcomes, explanation, and cleanup status. Parse the redacted policy-template.json with python3 -m json.tool and run test -s on README.md; record exit statuses. Review files for real ARNs, account IDs, login/session material, raw errors, and private configuration before submission.

## Cost and security

Use only two short-lived Standard String parameters at standard throughput and one temporary role/inline policy. Check current account pricing/budget before creation. Create no Advanced parameter, custom KMS key, EC2 instance, NAT gateway, load balancer, trail, analyzer, or other paid infrastructure. Delete both parameters and the role in the same study session.

The value learning is intentionally not a secret. Permission failures must not be repaired with broad managed policies. Keep real trust principals, resource ARNs, private config, credential caches, MFA material, raw STS responses, and raw CLI errors outside the repository and chat. The coach needs only booleans, categories, exit statuses, and cleanup confirmations.

## Submission

1. Region/budget checked; lab paths; source profile name:
2. Both Standard String parameters created with non-secret value; source identity existence/read checks:
3. Role trust limited to the intended current identity; only the exact inline GetParameter permission:
4. Assumed-role identity comparison result and exit status:
5. Allowed read value comparison and exit status:
6. Existing denied parameter: error category and nonzero exit status:
7. Trust versus permissions and source-versus-role explanation:
8. Both parameters deleted and post-delete categories; inline policy and role deleted:
9. Added profile removed; source logout result; no long-term keys or other resources:
10. Redacted JSON check and README check exit statuses; blocker if any:

## Pass criteria

- Two known existing Standard String parameters and a temporary role with the exact intended trust and one GetParameter resource permission; safe redacted artifacts.
- CLI uses temporary source/role credentials and an identity comparison explicitly verifies the intended assumed role.
- Allowed parameter value check is True with exit 0; the existing out-of-scope parameter produces AccessDenied with nonzero exit under the same role.
- Explanation distinguishes trust, role permissions, source permissions, missing resource, and the narrow scope of the live result.
- Both parameters and the inline policy/role are removed; added profile removed and source logout confirmed, without claiming all cached STS sessions were revoked.
- Valid redacted JSON, substantive nonempty README, sanitized evidence, and no credentials/actual identifying ARNs/account IDs or unintended resources retained.
