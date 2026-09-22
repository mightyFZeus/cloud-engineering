#!/usr/bin/env bash
   set -euo pipefail

   profile="${1:-aws-learning-readonly}"
   expected_region="${AWS_LEARNING_REGION:-eu-west-1}"

   command -v aws >/dev/null 2>&1 || {
     echo "FAIL: AWS CLI is not installed" >&2
     exit 1
   }

   configured_region="$(
     aws configure get region --profile "$profile" 2>/dev/null || true
   )"

   if [[ "$configured_region" != "$expected_region" ]]; then
     echo "FAIL: profile is missing or has an unexpected Region" >&2
     exit 1
   fi

   caller_arn="$(
     aws sts get-caller-identity \
       --profile "$profile" \
       --query Arn \
       --output text
   )"

   case "$caller_arn" in
     *":assumed-role/AWSLearningReadOnlyRole/"*)
       caller_type="assumed-role"
       ;;
     *)
       unset caller_arn
       echo "FAIL: profile is not using the expected assumed role" >&2
       exit 1
       ;;
   esac
   unset caller_arn

   region_count="$(
     aws ec2 describe-regions \
       --profile "$profile" \
       --query 'length(Regions)' \
       --output text \
       --no-cli-pager
   )"

   az_count="$(
     aws ec2 describe-availability-zones \
       --profile "$profile" \
       --region "$expected_region" \
       --filters \
         Name=zone-type,Values=availability-zone \
         Name=state,Values=available \
       --query 'length(AvailabilityZones)' \
       --output text \
       --no-cli-pager
   )"

   bucket_count="$(
     aws s3api list-buckets \
       --profile "$profile" \
       --query 'length(Buckets)' \
       --output text \
       --no-cli-pager
   )"

   printf 'PASS: AWS access baseline\n'
   printf 'Profile: %s\n' "$profile"
   printf 'Caller type: %s\n' "$caller_type"
   printf 'Configured Region: %s\n' "$configured_region"
   printf 'Regions visible: %s\n' "$region_count"
   printf 'Standard AZs available: %s\n' "$az_count"
   printf 'S3 bucket count: %s\n' "$bucket_count"
