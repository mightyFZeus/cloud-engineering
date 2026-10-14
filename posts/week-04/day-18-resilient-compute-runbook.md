# Day 18 posts

## LinkedIn

Day 18 of my cloud engineering journey: I consolidated my networking and resilient-compute labs into an operations runbook.

It includes a two-AZ architecture sketch, the request and response path, security-group restrictions, and a six-scenario failure table covering instance termination, wrong health-check paths, blocked ports, failed deployments, database availability, and an impaired AZ.

The troubleshooting checklist starts with the client request and works through the ALB, subnet routes, target health, security groups, process, probe path, and logs.

I kept the evidence boundaries clear: instance replacement was tested on AWS; the wrong probe path was tested locally; the other scenarios are analysis. An AZ outage was not tested, and I do not claim uninterrupted recovery.

This was a documentation exercise with no new AWS resources. The earlier compute lab was fully cleaned up.

#AWS #CloudEngineering #Troubleshooting #LearningInPublic

## X / Threads

Day 18: consolidated my 2-AZ ALB/Auto Scaling lab into a runbook: architecture, request path, 6 failure scenarios, and an ordered troubleshooting checklist. Tested evidence stays separate from outage analysis. No new AWS resources. #CloudEngineering
