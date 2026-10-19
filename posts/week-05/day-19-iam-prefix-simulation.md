# Day 19 posts

## LinkedIn

Day 19 of my cloud engineering journey: I tested a least-privilege S3 policy with the AWS IAM policy simulator.

The policy allowed reading objects under reports/ in one fictional bucket and allowed listing only with a matching reports/ prefix. It granted no upload permission.

Five simulations confirmed the boundary: reading reports was allowed; reading private objects and uploading reports were implicitly denied; listing reports was allowed and listing private was implicitly denied.

One useful correction during testing: the simulated resource initially remained *. Setting the specific object resource produced the intended allow result without widening the policy.

I documented bucket-level versus object-level permissions, implicit deny, and simulator limitations. These results evaluate the supplied policy and context; they do not prove live S3 access.

No policy was attached to a real identity, and no AWS resources were created.

#AWS #IAM #CloudEngineering #LearningInPublic

## X / Threads

Day 19: simulated a read-only S3 reports/ policy. Read reports + list reports allowed; read private, upload reports, and list private implicitly denied. Exact resources and prefix context matter. No policy attached or resources created. #AWS #IAM
