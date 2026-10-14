# Milestone 2 retrospective

Core practical evidence complete; external social publication is not verified.

## LinkedIn

Networking and resilient compute retrospective: I can now connect the architecture to the evidence.

I planned subnet boundaries, studied traffic paths, packaged a Go service for EC2, and deployed two instances across two AZs behind an ALB and Auto Scaling group. I terminated an instance and observed automatic replacement restore two healthy targets.

I also tested a wrong health-check path locally: the probe returned 404 while the service stayed running. Correcting the probe restored success. That exercise showed why replacing instances cannot fix configuration shared by all replacements.

My consolidated runbook includes the actual lab architecture, request and response flow, six failure scenarios, and a troubleshooting checklist. It identifies which results were observed and which are analysis. The public-subnet lab is documented as built; AZ failure recovery remains untested.

All paid compute lab resources were deleted, with zero remaining lab EBS volumes. No NAT gateways or Elastic IPs were created.

Next: return to least-privilege IAM and verify allowed and denied requests using policy simulation.

#AWS #CloudEngineering #Go #LearningInPublic

## X / Threads

Networking + resilient compute: built a 2-AZ ALB/ASG Go service, proved instance replacement, diagnosed a wrong probe locally, and wrote a 6-scenario runbook. Paid lab resources deleted. AZ outage recovery is analysis, not a tested claim. Next: IAM. #AWS
