# Day 21 — Parameter Store Audit

## Evidence source

CloudTrail Event history in eu-west-1, covering October 19, 2026.
Both events were privately matched to the Day 20 parameter using requestParameters.name.
Only sanitized audit fields are recorded here.

## Observed events

| Event | Time UTC | Source | Region | Identity type | readOnly | Error code |
|---|---|---|---|---|---|---|
| PutParameter | 2026-10-19T19:49:15Z | ssm.amazonaws.com | eu-west-1 | IAMUser | false | none |
| DeleteParameter | 2026-10-19T20:08:57Z | ssm.amazonaws.com | eu-west-1 | IAMUser | false | none |

## Timeline

- Creation: 19:49:15 UTC, or 20:49:15 Lagos.
- Deletion: 20:08:57 UTC, or 21:08:57 Lagos.
- Deletion followed creation by 19 minutes 42 seconds.
- Both matched operations contained no API error code.

## Related evidence

[Day 20 Go reader and cleanup](../../../docs/week-05/day-20-go-parameter-store.md)


## Audit evidence versus application output

CloudTrail records AWS API activity, including the operation, time, Region, and identity type.
The Go program's PASS result records that its own response checks succeeded.
Application and service logs describe program behavior and errors; they serve a different purpose from API audit records.

## Event history scope and limitations

Event history provides the past 90 days of recorded management events in the selected Region.
It does not contain every data event or application log.
Viewing existing Event history requires no new trail or CloudTrail Lake store.

A missing event alone does not prove that no request occurred.
Check the Region, time range, filters, event coverage, and delivery delay.
AWS does not record ParameterNotFound exceptions from GetParameter or DeleteParameter in CloudTrail.


## Final explanation and cleanup

The learner explained that CloudTrail shows the API action, time, and identity, whereas Go PASS shows the program's checks succeeded. An event may be absent due to delivery delay. No new AWS resources, trail, Lake store, or logging configuration changes were made. The Day 20 parameter remains deleted. README presence check exited 0.
