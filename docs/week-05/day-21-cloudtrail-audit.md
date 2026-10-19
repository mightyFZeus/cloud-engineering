# Day 21 — Parameter Store audit in CloudTrail

Passed October 19, 2026: six criteria verified by learner event-field checkpoints and inspection of the saved audit README.

## Verified evidence

- Regional CloudTrail Event history in eu-west-1 covering October 19.
- Both events privately matched to the Day 20 parameter using requestParameters.name.
- PutParameter: 2026-10-19T19:49:15Z; ssm.amazonaws.com; eu-west-1; IAMUser; readOnly=false; no error code.
- DeleteParameter: 2026-10-19T20:08:57Z; ssm.amazonaws.com; eu-west-1; IAMUser; readOnly=false; no error code.
- UTC-to-Lagos conversion adds one hour: creation 20:49:15 and deletion 21:08:57. Elapsed interval verified as 19 minutes 42 seconds.
- Learner distinguished API audit evidence from Go's application checks and identified delivery delay as one reason an event may be missing.
- README check exited 0; no new resources, trail, Lake store, or logging changes. Day 20 parameter remains deleted.

## Review and limitations

Coach removed an identical duplicate timeline section caused by repeated appending. The relative Day 20 evidence link resolves. Raw event JSON was not copied into the repository.

Event history covers recorded management events for the selected Region over the past 90 days. It is not a complete application log or data-event archive. Missing events need investigation of Region, filters, time range, coverage, and delay. AWS does not record ParameterNotFound exceptions from GetParameter or DeleteParameter in CloudTrail.

## Artifact

[Sanitized audit timeline](../../labs/week-05/day-21-cloudtrail-audit/README.md)

## References

- [CloudTrail Event history](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/view-cloudtrail-events.html)
- [Systems Manager CloudTrail integration](https://docs.aws.amazon.com/systems-manager/latest/userguide/monitoring-cloudtrail-logs.html)
