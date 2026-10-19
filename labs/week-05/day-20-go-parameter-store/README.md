
# Day 20 — Go Parameter Store Reader

## Verified results

- Region: eu-west-1; budget checked.
- Authentication: aws login with the aws-learning profile.
- Parameter: /cloud-eng/day20/environment.
- Tier: Standard; type: String; non-secret value: learning.
- gofmt completed; build exit status: 0.
- Successful read: parameter check: PASS; exit status: 0.
- Missing-name test: ParameterNotFound; exit status: 1.
- Parameter deletion exit status: 0.
- Post-delete read: ParameterNotFound; exit status: 1.

## Troubleshooting

The CLI initially could not reach the sign-in token endpoint.
A diagnostic HTTP request later returned 405, confirming endpoint connectivity.
Retrying deletion succeeded. The initial connectivity failure's root cause was not established.

## Implementation

The Go program uses AWS SDK v2's default credential chain, an explicit Region,
a 15-second timeout, typed missing-parameter handling, and sanitized output.
It contains no hard-coded credentials.


## Credentials, settings, and cleanup

The default credential chain loaded temporary credentials from the aws-learning login profile. IAM policies determine whether the authenticated identity may read the parameter. The value learning is non-secret, so String is appropriate. A password requires secret handling, such as SecureString with KMS encryption and appropriately restricted access; SecureString was discussed but not tested here.

Authentication establishes identity; authorization determines which actions it may perform. Environment credential overrides were cleared inside a temporary subshell for the Go runs so the selected login profile was used.

README check exited 0. Source-profile logout exited 0. No long-term access keys or EC2 instances were created. The temporary parameter was deleted and its absence verified by the typed ParameterNotFound result.
