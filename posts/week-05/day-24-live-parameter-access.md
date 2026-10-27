# Day 24 posts

## LinkedIn

Day 24 of my cloud engineering journey: I tested least privilege against real AWS Parameter Store requests.

I created two non-secret parameters and a temporary IAM role with GetParameter permission on just one of them. I confirmed both parameters existed, then verified that the CLI was using the intended assumed role.

The role read the allowed parameter successfully. Reading the other existing parameter returned AccessDeniedException.

This made the identity distinction concrete: the source identity could read both, while the assumed role used its own scoped permissions. Trust controls who can assume a role; permissions control what it can do.

Cleanup also required troubleshooting. The CLI encountered Parameter Store connection errors, so I reused my Go reader to verify deletion. Both names returned ParameterNotFound. The original CLI connection failure's cause remains unconfirmed.

Both parameters, the role and its inline policy were removed. I removed the temporary local profile and logged out of the source session. No long-term access keys or additional AWS resources were created.

Proof of work: https://github.com/mightyFZeus/cloud-engineering/blob/main/docs/week-05/day-24-live-parameter-access.md

#AWS #IAM #Go #CloudEngineering

## X / Threads

Day 24: verified live least privilege. A temporary role read one SSM parameter and got AccessDenied for another existing one. Go confirmed deletion after CLI connection errors. Role, parameters and local profile removed. #AWS #IAM
