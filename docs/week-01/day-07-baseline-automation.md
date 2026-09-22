# Day 7 — AWS Access Baseline Automation

## Outcome

Created [`scripts/aws-baseline-check.sh`](../../scripts/aws-baseline-check.sh) to check the expected assumed role and `eu-west-1` Region, then print only non-sensitive counts.

## Evidence

- The successful AWS run was reported with 17 visible Regions, three available standard AZs, and zero S3 buckets.
- `bash -n` passed.
- A nonexistent profile produced the expected error and exit status 1.
- A leading space before the shebang prevented direct execution of the saved copy. The header was corrected and the local syntax and missing-profile checks were rerun.
- The source login session was reported logged out, and no AWS resources were created.

## What the script protects

`set -e` stops on an unhandled command failure, `set -u` catches unset variables, and `pipefail` catches failures inside pipelines. The caller ARN is captured only to identify the assumed role, then unset and never printed. The `length(...)` queries keep terminal output to counts; they do not reduce the data AWS sends to the CLI or change the role's permissions.

## Next step

Launch one short-lived Amazon Linux EC2 instance, connect through Session Manager without opening an inbound port, inspect the host, and verify cleanup.
