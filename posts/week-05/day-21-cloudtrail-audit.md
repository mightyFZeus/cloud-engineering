# Day 21 posts

## LinkedIn

Day 21 of my cloud engineering journey: I traced my Parameter Store lab through AWS CloudTrail Event history.

I found the real PutParameter and DeleteParameter events in eu-west-1 and privately matched both to the temporary setting from Day 20. The audit timeline showed deletion 19 minutes 42 seconds after creation, with no API error code on either matched operation.

I saved only safe fields: operation, UTC time, service, Region, identity type, readOnly flag, and error classification. I excluded raw records and identifying details.

The useful distinction: CloudTrail records API activity; my Go program's PASS output records that its application checks succeeded. A missing event alone does not prove that no request occurred—filters, Region, event coverage, and delivery delay matter.

This was a read-only review of existing Event history. I created no trail, CloudTrail Lake store, or other AWS resource and changed no logging settings.

#AWS #CloudTrail #CloudEngineering #LearningInPublic

## X / Threads

Day 21: found my Parameter Store create/delete events in CloudTrail and built a sanitized UTC timeline. CloudTrail records API activity; Go PASS records application checks. Missing events need context. No resources or logging changes. #AWS #CloudTrail
