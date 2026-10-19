# Day 21 — Audit Parameter Store Activity in CloudTrail

**Date:** October 19, 2026
**Due:** 11:00 PM Africa/Lagos
**Timebox:** 60–75 minutes
**Status:** Assigned

## Outcome

Find the real Day 20 parameter creation and deletion in CloudTrail Event history, build a sanitized timeline, and distinguish audit evidence from application output. Use existing events; no new parameter or trail needed.

## Steps

1. Spend 10 minutes reviewing the two references below. Sign into the console using your everyday MFA-protected identity; select eu-west-1. The CLI was logged out after Day 20; browser access is sufficient for this lab.
2. Create labs/week-05/day-21-cloudtrail-audit/README.md locally. Record that the source is regional CloudTrail Event history for recorded management events over the past 90 days.
3. Open CloudTrail -> Event history. Set a time range covering Day 20. Clear any restrictive default filters, then filter Event name = PutParameter. Open the event whose requestParameters.name matches /cloud-eng/day20/environment; do not download or paste its raw JSON.
4. Privately inspect eventName, eventSource, eventTime, awsRegion, userIdentity.type, readOnly, and errorCode/errorMessage if present. Save only operation name, UTC event time, Region, identity type (not username/ARN), readOnly value or not present, and error code or none. Mark a successful response using absence of an API error together with the matched creation operation; do not infer success only from a null responseElements field.
5. Replace the name filter with DeleteParameter and find the matching Day 20 delete. Record the same safe fields. Compare creation and deletion order using UTC times; Lagos is UTC+1. Distinguish private matching of requestParameters from the sanitized public summary.
6. Build a two-row audit table and a short timeline. Explain the difference between CloudTrail API activity, Go's PASS/exit status, and application/service logs. A missing event alone does not prove an API call never happened; check Region, time range, filters, event coverage, and delivery delay.
7. Explain why creating a trail or CloudTrail Lake is unnecessary for this exercise, and why Event history is not a record of every data event or application log. AWS does not record ParameterNotFound exceptions from GetParameter or DeleteParameter in CloudTrail; do not expect the missing-name tests to produce those events. If the successful create/delete events have not appeared, wait briefly, refresh, and report the blocker rather than creating additional resources.
8. Link your Day 20 note and record cleanup: no new AWS resources, existing parameter remains deleted. Run test -s labs/week-05/day-21-cloudtrail-audit/README.md and record exit status. Inspect the saved text for identifiers, raw logs, login URLs, secrets, and unnecessary addresses.

## References

- [Working with CloudTrail Event history](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/view-cloudtrail-events.html)
- [Logging Systems Manager API calls](https://docs.aws.amazon.com/systems-manager/latest/userguide/monitoring-cloudtrail-logs.html)

## Cost and security

Viewing Event history has no CloudTrail charge. Create no trail, Lake store, data-event selector, S3 bucket, or parameter. Raw events contain account/identity details, IPs, request IDs, and resource identifiers; inspect privately and save only the specified safe fields. Do not modify log configuration or delete audit records.

## Submission

1. README path; Region:
2. PutParameter safe summary: UTC time, eventSource, identity type, readOnly, error code/none:
3. DeleteParameter safe summary: same fields:
4. Both events privately matched to the Day 20 parameter: yes/no
5. Timeline order and timezone interpretation:
6. Audit versus application-output explanation:
7. Event history scope and missing-event limitations:
8. test -s exit status:
9. AWS resources created or log configuration changed:
10. Blocker:

## Pass criteria

- Correct Region/time range and two actual management events matched privately to the lab parameter.
- Sanitized fields faithfully reflect observed events, including any absent field.
- Timeline correctly relates creation and deletion in UTC.
- Explains audit evidence versus application output and limitations of event absence.
- No new resources, trail/Lake, or logging changes.
- Nonempty saved README with safe contents and evidence links.
