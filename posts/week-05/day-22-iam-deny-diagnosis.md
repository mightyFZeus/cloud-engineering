# Day 22 posts

## LinkedIn

Day 22 of my cloud engineering journey: I tested how an explicit IAM deny overrides an allow.

Using the IAM Policy Simulator in Custom mode, I compared four requests:

- The baseline allowed reading a restricted report.
- A narrow deny blocked that same object.
- An ordinary report remained readable.
- Restoring the baseline restored simulated access.

The denied request matched DenyRestrictedReports. That gave me a concrete distinction: an implicit deny means no applicable allow grants access; an explicit deny means a matching restriction blocks it.

I also documented three risks and proposed controls covering exposed credentials, excessive prefix access, and public data exposure. One useful detail: denying GetObject content access does not automatically hide object names from permitted listings.

These were simulations using fictional resources. No policy was attached to a real identity, no AWS resources were created, and no real security controls changed. A simulated allow does not prove that a real request will succeed.

Proof of work: https://github.com/mightyFZeus/cloud-engineering/blob/main/docs/week-05/day-22-iam-deny-diagnosis.md

#AWS #IAM #CloudEngineering #LearningInPublic

## X / Threads

Day 22: tested IAM explicit deny in Custom simulation. A narrow deny blocked restricted reports while ordinary reports stayed readable. Restoring the baseline restored access. Simulation does not prove live access. No real permissions changed. #AWS #IAM
