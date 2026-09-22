# AWS Daily Assignment Tracker

**Schedule:** Assignment at 12:00 PM; reminders at 8:00 PM and 10:00 PM Africa/Lagos; submission due by 11:00 PM daily  
**Started:** Monday, September 14, 2026  
**Target pace:** 60–90 focused minutes per day, approximately 10 hours per week  
**Roadmap:** `aws-cloud-engineer-12-week-plan.md`

## Verification workflow

1. The daily assignment contains an outcome, exact steps, learning links, safety notes, submission evidence, and pass/fail criteria.
2. Submit evidence in this task by 11:00 PM. Never include credentials, tokens, full AWS account IDs, unredacted ARNs, private IPs, or secrets.
3. The submission is marked **Passed**, **Needs correction**, **Incomplete**, or **Not submitted**.
4. A failed or incomplete prerequisite is corrected before unrelated work is added.
5. When a day passes, save its build-in-public post draft and load the next assignment before reporting completion. Milestone completion also produces a retrospective post.
6. The 8:00 PM and 10:00 PM reminders stay quiet once that day's assignment has been submitted and verified as Passed.

## Status key

- ⬜ Assigned
- 🟡 Submitted / under review
- ✅ Passed
- 🔧 Needs correction
- ⏸ Not submitted / carried forward

## Assignment log

| Day | Date | Assignment | Due | Status | Verification | Social post |
|---:|---|---|---:|---|---|---|
| 1 | Sep 14 | Secure AWS learning account and establish cost guardrails | 11:00 PM | ⏸ Not submitted | Completed through the Day 2 catch-up | Root-lesson drafts delivered |
| 2 | Sep 15 | Complete Day 1 account-security and cost-guardrail setup | 11:00 PM | ✅ Passed | 8/8 criteria passed | Draft delivered |
| 3 | Sep 16 | Install AWS CLI v2 and authenticate with temporary credentials | 11:00 PM | ✅ Passed | 8/8 criteria passed | Draft delivered |
| 4 | Sep 17 | Map AWS Regions/AZs and explain shared responsibility | 11:00 PM | ✅ Passed | 6/6 criteria passed | Draft delivered |
| 5 | Sep 18 | Create and assume a read-only IAM role | 11:00 PM | ✅ Passed | 6/6 criteria passed | Draft saved |
| 6 | Sep 18 | Replace broad read-only access with least privilege | Before study ends | ✅ Passed | 6/6 criteria passed | Draft saved |
| 7 | Sep 18 | Automate a safe AWS access baseline check | Before study ends | ✅ Passed Sep 22 | 6/6 criteria passed | Draft saved |
| 8 | Sep 22 | Launch, inspect, and clean up a Linux EC2 instance | 11:00 PM | ✅ Passed | 6/6 criteria passed | Draft saved |

---

## Day 1 — Secure AWS learning account and establish cost guardrails

**Date:** Monday, September 14, 2026  
**Due:** 11:00 PM Africa/Lagos  
**Timebox:** 60–90 minutes  
**Outcome:** Your AWS learning account has safe sign-in boundaries and alerts before you create billable infrastructure.

### Learn (15 minutes maximum)

- Read the opening recommendations in [AWS root-user best practices](https://docs.aws.amazon.com/IAM/latest/UserGuide/root-user-best-practices.html).
- Review [AWS budget templates](https://docs.aws.amazon.com/cost-management/latest/userguide/budget-templates.html).

### Do (45–65 minutes)

1. Sign in to your AWS learning account. If you do not have one, create it first.
2. On the root user, enable MFA and verify that no root access keys exist. Do not create one.
3. Create a non-root identity for everyday administration. Use IAM Identity Center with MFA and temporary credentials if the account is already in AWS Organizations or you are not relying on new-account Free Tier credits. If this is a standalone account with active Free Tier credits, do not create an AWS Organization only for this exercise: use a separate administrative IAM user with MFA temporarily, create no access keys, and record that limitation.
4. Sign out of root and sign back in using the non-root identity. Do not continue daily work as root.
5. Create a **zero-spend budget** that alerts when Free Tier usage is exceeded.
6. Create a small **monthly cost budget** of your choosing (suggested: USD 5–10) with email alerts. Include an actual-spend alert and, if the console offers it, a forecasted-spend alert.
7. Choose and record one default AWS Region for the course. Consistency matters more than finding a perfect Region; confirm that a tutorial's services are available before deploying.
8. Write a short safety note with: the identity you will use daily, your chosen Region, your budget amount/thresholds, and why root credentials must not be used for routine work.

### Cost and security warning

- This assignment creates no compute resources.
- AWS Budgets are alerts, not guaranteed hard spending caps, and billing data/notifications can be delayed.
- Never send screenshots or command output containing credentials, a complete account ID, a complete ARN, recovery details, QR codes, or MFA setup secrets.

### Submit by 11:00 PM

Copy this block into the chat and complete it:

```text
Day 1 submission
1. Root MFA enabled: yes/no
2. Root access keys absent: yes/no
3. Everyday identity type: IAM Identity Center / IAM user / other
4. MFA enabled on everyday identity: yes/no
5. Zero-spend budget active: yes/no
6. Monthly budget amount and alert thresholds: [amount; thresholds]
7. Default Region: [region code]
8. Why I should not use root daily (2–4 sentences):
9. Blocker or question:
```

Screenshots are optional. Redact sensitive identifiers before attaching them.

### Pass criteria

- [x] Root MFA is enabled.
- [x] Root has no access keys.
- [x] A separate everyday identity exists, has MFA, and can sign in.
- [x] A zero-spend budget and a small monthly budget are active.
- [x] Default Region `eu-west-1` and the zero-spend `$0.01 actual` threshold are recorded.
- [x] The USD 10 monthly-budget alerts are recorded: 85% actual and 100% forecasted.
- [x] The explanation shows an accurate understanding of why root is reserved for root-only tasks.

---

## Day 2 — Catch up: complete the account-security prerequisite

**Date:** Tuesday, September 15, 2026  
**Due:** 11:00 PM Africa/Lagos  
**Timebox:** 60–90 minutes  
**Outcome:** Finish and submit the Day 1 safety foundation before any AWS resource labs begin.

### Learn (15 minutes maximum)

- Read the first recommendations in [AWS root-user best practices](https://docs.aws.amazon.com/IAM/latest/UserGuide/root-user-best-practices.html).
- Review [AWS budget templates](https://docs.aws.amazon.com/cost-management/latest/userguide/budget-templates.html).

### Do (45–65 minutes)

1. Finish every unchecked item in the Day 1 pass criteria above.
2. Verify root MFA is enabled and that the root user has no access keys.
3. Verify you can sign in with a separate everyday identity protected by MFA.
4. Verify both the zero-spend and small monthly budgets are active and have notification recipients.
5. Record your default Region and a short explanation of why root is reserved for root-only tasks.
6. Submit the completed evidence form below. If one step is impossible, name the exact screen, step, or error instead of writing only “blocked.”

### Cost and security warning

- Create no compute, database, load-balancer, NAT, EKS, or other workload resources today.
- Do not submit account IDs, complete ARNs, credentials, MFA QR codes, recovery information, or secrets.
- Budget alerts are not guaranteed hard spending caps and billing notifications can be delayed.

### Submit by 11:00 PM

```text
Day 2 catch-up submission
1. Root MFA enabled: yes/no
2. Root access keys absent: yes/no
3. Everyday identity type: IAM Identity Center / IAM user / other
4. MFA enabled on everyday identity: yes/no
5. Successful non-root sign-in: yes/no
6. Zero-spend budget active: yes/no
7. Monthly budget amount and alert thresholds: [amount; thresholds]
8. Budget notification recipient confirmed: yes/no
9. Default Region: [region code]
10. Why I should not use root daily (2–4 sentences):
11. Exact blocker, if any:
```

Screenshots are optional and must be redacted.

### Pass criteria

- [x] Root MFA is enabled and root access keys are absent.
- [x] A separate MFA-protected everyday identity can sign in without access keys.
- [x] Zero-spend and USD 10 monthly budgets are active with notifications configured.
- [x] Default Region `eu-west-1` is submitted.
- [x] Zero-spend alert is confirmed at `$0.01 actual spend`.
- [x] The USD 10 monthly-budget alerts are submitted: 85% actual and 100% forecasted.
- [x] An accurate root-user explanation is submitted.
- [x] No sensitive identifiers or secrets are included in the evidence.

### Verification notes

- Submitted identity: temporary administrative IAM user with MFA and no access keys.
- Non-root sign-in was confirmed.
- Default Region confirmed as `eu-west-1` (Europe, Ireland).
- Zero-spend alert confirmed at `$0.01 actual spend`.
- USD 10 monthly-budget alerts confirmed at 85% actual spend and 100% forecasted spend.
- **Result:** Passed on September 15, 2026. All eight verification criteria are complete.

---

## Day 3 — AWS CLI with temporary credentials

**Date:** Wednesday, September 16, 2026  
**Due:** 11:00 PM Africa/Lagos  
**Timebox:** 60–90 minutes  
**Outcome:** Use AWS from your local terminal through a named profile and an expiring browser-authenticated session—without creating an IAM access key.

### Learn (15 minutes maximum)

- Read [Installing or updating AWS CLI v2](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html), selecting the macOS instructions.
- Read [Login for local development using console credentials](https://docs.aws.amazon.com/en_en/cli/latest/userguide/cli-configure-sign-in.html). Focus on `aws login`, named profiles, temporary credentials, and `aws logout`.

### Do (45–60 minutes)

1. In your local terminal, check the installation:

   ```bash
   aws --version
   ```

2. If AWS CLI is missing or older than **2.32.0**, install or update AWS CLI v2 using the official macOS instructions. Run `aws --version` again afterward.
3. Confirm that your browser is signed in with your everyday IAM user—not root. Then create a browser-authenticated named profile:

   ```bash
   aws login --profile aws-learning --region eu-west-1
   ```

4. Complete the browser flow and select the everyday IAM-user session. Do not select or sign in as root.
5. Set and inspect the non-sensitive profile configuration:

   ```bash
   aws configure set region eu-west-1 --profile aws-learning
   aws configure set output json --profile aws-learning
   aws configure list-profiles
   aws configure get region --profile aws-learning
   aws configure get output --profile aws-learning
   ```

6. Verify the active AWS principal. Inspect the output locally, but do not share the complete account ID, ARN, or user ID:

   ```bash
   aws sts get-caller-identity --profile aws-learning
   ```

7. Make one read-only service call:

   ```bash
   aws ec2 describe-regions \
     --profile aws-learning \
     --query 'Regions[].RegionName' \
     --output table \
     --no-cli-pager
   ```

8. Write three short notes explaining:
   - what a named profile is;
   - what `sts get-caller-identity` proves;
   - why `aws login` is safer than a long-lived IAM access key.
9. End the session and remove its cached login credentials:

   ```bash
   aws logout --profile aws-learning
   ```

### Troubleshooting

- If `aws login` is unknown, your CLI is older than 2.32.0; update AWS CLI v2.
- If login returns an authorization error, verify that the everyday IAM identity can use the AWS managed `SignInLocalDevelopmentAccess` policy. Your administrative IAM group may already allow the required actions.
- If the browser displays multiple identities, explicitly choose the non-root IAM user.

### Safety rules

- Do **not** create an IAM access key.
- Do **not** run the interactive plain `aws configure` flow that asks for an access-key ID and secret key.
- Do not paste the complete output of `get-caller-identity`; redact the account number, user ID, username, and ARN.
- This assignment creates no workload resources and should not produce infrastructure charges.

### Submit by 11:00 PM

```text
Day 3 submission
1. AWS CLI version: [version]
2. Authentication method: aws login / other
3. Named profile: [name]
4. Configured Region: [region]
5. Output format: [format]
6. get-caller-identity: success/failure; principal type only (IAM user/role/root)
7. describe-regions: success/failure
8. No long-lived access key created: yes/no
9. aws logout completed: yes/no
10. Named profile means:
11. get-caller-identity proves:
12. Temporary login is safer because:
13. Exact blocker, if any:
```

### Pass criteria

- [x] AWS CLI v2.36.47 is installed.
- [x] The `aws-learning` profile uses `aws login`, not a long-lived access key.
- [x] The profile records `eu-west-1`.
- [x] The profile records JSON as its output-format setting.
- [x] `sts get-caller-identity` succeeds as the non-root everyday IAM user.
- [x] The read-only Region query succeeds.
- [x] The temporary session is logged out successfully.
- [x] The three explanations are technically accurate.
- [x] No secrets or unredacted identifiers are submitted.

### Submission review — September 17, 2026

- Verified: AWS CLI v2.36.47, `aws login`, `aws-learning`, `eu-west-1`, non-root IAM caller, no long-lived access key, and successful logout.
- Missing: output format was left as `[format]` and `describe-regions` was left as `success/failure`.
- Correction: `sts get-caller-identity` proves which account and principal signed the request; it does not prove which service actions that principal is authorized to perform.
- Follow-up verified: `describe-regions` succeeded and the corrected explanations are technically accurate.
- Follow-up verified: the `aws-learning` profile output format is `json`.
- **Result:** Passed on September 17, 2026. All eight verification criteria are complete.

---

## Day 4 — Regions, Availability Zones, and shared responsibility

**Date:** Thursday, September 17, 2026  
**Due:** 11:00 PM Africa/Lagos  
**Timebox:** 60–90 minutes  
**Outcome:** Use live AWS metadata to map `eu-west-1`, explain how a multi-AZ design survives an AZ failure, and distinguish AWS responsibilities from customer responsibilities for EC2 and Lambda.

### Learn (15 minutes maximum)

- Read [AWS Regions and Availability Zones](https://docs.aws.amazon.com/global-infrastructure/latest/regions/aws-regions-availability-zones.html).
- Read [AWS Shared Responsibility Model](https://aws.amazon.com/compliance/shared-responsibility-model/), focusing on “security of the cloud” versus “security in the cloud.”

### Do (45–65 minutes)

1. Start a new temporary CLI session as your everyday IAM user:

   ```bash
   aws login --profile aws-learning --region eu-west-1
   ```

2. List all Regions and their opt-in status:

   ```bash
   aws ec2 describe-regions \
     --all-regions \
     --profile aws-learning \
     --query 'Regions[].[RegionName,OptInStatus]' \
     --output table \
     --no-cli-pager
   ```

3. List the standard Availability Zones available to your account in Ireland:

   ```bash
   aws ec2 describe-availability-zones \
     --region eu-west-1 \
     --profile aws-learning \
     --query 'AvailabilityZones[?ZoneType==`availability-zone`].[ZoneName,ZoneId,State]' \
     --output table \
     --no-cli-pager
   ```

4. Record the number of standard AZs returned and their Zone names and Zone IDs. These are infrastructure identifiers, not credentials.
5. Draw this small text model using your actual CLI results:

   ```text
   AWS account
   └── Region: eu-west-1 (Ireland)
       ├── AZ name → AZ ID
       ├── AZ name → AZ ID
       └── AZ name → AZ ID
   ```

6. Under the model, answer in 3–5 sentences:
   - What is the difference between a Region and an Availability Zone?
   - Why can a Go API deployed across two AZs survive one AZ failure?
   - What failure would a two-AZ, single-Region design not protect against?
7. Write **120–180 words** comparing shared responsibility for:
   - **EC2:** AWS protects facilities, hardware, networking, and virtualization; you manage the guest OS, patches, application, security-group rules, IAM, and data.
   - **Lambda:** AWS manages more of the compute platform and managed runtime; you still manage code, dependencies, IAM permissions, configuration, and data.
8. End the temporary session:

   ```bash
   aws logout --profile aws-learning
   ```

### Safety and cost rules

- The commands are read-only and create no AWS resources.
- Do not create EC2, load balancer, database, Lambda, networking, or other infrastructure today.
- Do not submit account IDs, user IDs, ARNs, credentials, tokens, or browser-login URLs.
- AZ names and AZ IDs may be submitted.

### Submit by 11:00 PM

```text
Day 4 submission
1. aws login: success/failure
2. describe-regions --all-regions: success/failure
3. eu-west-1 standard AZ count:
4. AZ map:
   - [AZ name] -> [AZ ID]
   - [AZ name] -> [AZ ID]
   - [AZ name] -> [AZ ID]
5. Region vs AZ and failure explanation (3–5 sentences):
6. EC2 vs Lambda shared-responsibility explanation (120–180 words):
7. aws logout: success/failure
8. Resources created: none / list them
9. Exact blocker, if any:
```

### Pass criteria

- [x] Both read-only CLI queries succeed through the `aws-learning` profile.
- [x] The submitted three-AZ `eu-west-1` AZ-name-to-ID map matches the reported CLI result.
- [x] Region, AZ, multi-AZ resilience, and remaining regional risk are explained accurately.
- [x] The EC2/Lambda comparison assigns AWS and customer responsibilities accurately and is 120–180 words.
- [x] The temporary session is logged out.
- [x] No resources, secrets, or sensitive account identifiers are submitted.

### Submission review — September 17, 2026

- Verified: login, both read-only queries, three standard AZs, submitted AZ mapping, logout, and no resources created.
- Item 5 correctly distinguishes a Region from an AZ and recognizes AZ fault isolation, but it does not yet explain the required two-AZ Go API design or the remaining risk of a complete regional failure.
- Item 6 was left blank; a 120–180-word EC2-versus-Lambda shared-responsibility comparison is required.
- Follow-up verified: item 5 now accurately covers a two-AZ Go API with traffic failover and the remaining risk of a complete regional failure.
- Remaining: item 6 is technically correct but is only about 45 words; expand it to the required 120–180 words.
- Second item 6 draft is approximately 72 words. Its EC2 comparison is accurate, but “only responsibility” for Lambda is too narrow: the customer also owns dependencies, IAM, configuration and secrets, observability, and application security.
- Third item 6 draft meets the length requirement at exactly 120 words. One accuracy correction remains: Lambda customers do not “only” own the listed items; they also retain responsibility for dependencies, least-privilege access, observability, and application security.
- Final item 6 revision accurately assigns EC2 and Lambda responsibilities and contains 129 words.
- **Result:** Passed on September 17, 2026. All six verification criteria are complete.

---

## Day 5 — IAM roles, temporary credentials, and policy evaluation

**Date:** Friday, September 18, 2026  
**Due:** 11:00 PM Africa/Lagos  
**Timebox:** 60–90 minutes  
**Outcome:** Create a cost-free read-only IAM role, assume it through a second CLI profile, and prove—without making a real write request—that policies allow reading but deny writing.

### Learn (15 minutes maximum)

- Read [Compare IAM identities and credentials](https://docs.aws.amazon.com/IAM/latest/UserGuide/introduction_identity-management.html), focusing on IAM users versus roles.
- Skim [Using an IAM role in the AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-role.html), focusing on `role_arn`, `source_profile`, and temporary credentials.

### Do (45–65 minutes)

1. Start a temporary session with your everyday identity:

   ```bash
   aws login --profile aws-learning --region eu-west-1
   ```

2. In the AWS console, open **IAM → Roles → Create role** and configure:
   - Trusted entity type: **AWS account**
   - Trusted account: **This account**
   - Permission policy: AWS managed **ReadOnlyAccess**
   - Role name: `AWSLearningReadOnlyRole`
   - Tags: `Project=aws-cloud-engineering-journey`, `Environment=learning`, and `Owner=student`
3. Create the role. Do not attach `AdministratorAccess` or create an access key.
4. Back in your terminal, resolve the account ID locally and create a role-based profile. Do not print or submit the account ID or role ARN:

   ```bash
   AWS_LEARNING_ACCOUNT_ID="$(aws sts get-caller-identity \
     --profile aws-learning \
     --query Account \
     --output text)"

   aws configure set \
     role_arn "arn:aws:iam::${AWS_LEARNING_ACCOUNT_ID}:role/AWSLearningReadOnlyRole" \
     --profile aws-learning-readonly

   aws configure set source_profile aws-learning \
     --profile aws-learning-readonly

   aws configure set role_session_name day5-learning \
     --profile aws-learning-readonly

   aws configure set region eu-west-1 \
     --profile aws-learning-readonly

   aws configure set output json \
     --profile aws-learning-readonly
   ```

5. Verify that the second profile assumes the role. Inspect locally; submit only the principal type—not the ARN or account ID:

   ```bash
   aws sts get-caller-identity --profile aws-learning-readonly
   ```

   The ARN should indicate an `assumed-role` session named `day5-learning`.
6. Prove that a read operation works:

   ```bash
   aws s3api list-buckets \
     --profile aws-learning-readonly \
     --query 'length(Buckets)' \
     --output text \
     --no-cli-pager
   ```

7. Simulate—without executing—one read action and one write action:

   ```bash
   aws iam simulate-principal-policy \
     --policy-source-arn "arn:aws:iam::${AWS_LEARNING_ACCOUNT_ID}:role/AWSLearningReadOnlyRole" \
     --action-names s3:ListAllMyBuckets s3:PutObject \
     --profile aws-learning \
     --query 'EvaluationResults[].{Action:EvalActionName,Decision:EvalDecision}' \
     --output table \
     --no-cli-pager
   ```

   Expected: `s3:ListAllMyBuckets` is allowed and `s3:PutObject` is denied, usually as `implicitDeny`.
8. Write 2–4 sentences explaining:
   - IAM user versus IAM role;
   - permission policy versus trust policy;
   - implicit deny versus explicit deny.
9. End the source login session:

   ```bash
   aws logout --profile aws-learning
   unset AWS_LEARNING_ACCOUNT_ID
   ```

### Safety and cost rules

- IAM roles and policy simulation do not incur infrastructure charges.
- Keep `AWSLearningReadOnlyRole` for later course labs; it has no credentials of its own and costs nothing.
- Do not run a real `PutObject`, create a bucket, or create any workload resource.
- Never submit the account ID, complete role/user ARN, temporary credentials, browser-login URL, or cached credential files.

### Submit by 11:00 PM

```text
Day 5 submission
1. Role created: yes/no
2. Attached permission policy: [name]
3. Required tags present: yes/no
4. Role profile name: [name]
5. get-caller-identity: success/failure; principal type only
6. Read-only S3 list call: success/failure; bucket count optional
7. Policy simulation result: ListAllMyBuckets=[decision], PutObject=[decision]
8. Source profile logged out: yes/no
9. IAM user vs role:
10. Permission policy vs trust policy:
11. Implicit deny vs explicit deny:
12. Resources created besides the IAM role: none / list
13. Exact blocker, if any:
```

### Pass criteria

- [x] `AWSLearningReadOnlyRole` exists with `ReadOnlyAccess` and all three tags.
- [x] `aws-learning-readonly` successfully assumes the role through `aws-learning`.
- [x] The role can perform the safe read-only S3 listing.
- [x] Policy simulation allows `s3:ListAllMyBuckets` and denies `s3:PutObject` without executing either action.
- [x] User/role, permission/trust policy, and implicit/explicit deny are explained accurately.
- [x] The source session is logged out and no access keys, workload resources, secrets, or identifiers are submitted.

### Submission review — September 18, 2026

- Verified: role creation, `ReadOnlyAccess`, required tags, role-based profile, assumed-role identity, read-only S3 listing, expected policy-simulation decisions, and source-profile logout.
- Follow-up verified: items 9–11 accurately distinguish user/role, permission/trust policy, and implicit/explicit deny; no additional resources or blockers were reported.
- **Result:** Passed on September 18, 2026. All six verification criteria are complete.

---

## Day 6 — Replace broad read-only access with least privilege

**Date:** Friday, September 18, 2026  
**Due:** Before you stop studying today  
**Timebox:** 60–75 minutes  
**Outcome:** Replace the broad AWS-managed `ReadOnlyAccess` policy with a small inline policy, then prove that only the intended discovery actions are allowed.

### Learn (10 minutes maximum)

- Read [Managed policies and inline policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_managed-vs-inline.html), focusing on the one-to-one relationship between an inline policy and its role.
- Skim [IAM policy testing with the policy simulator](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_testing-policies.html), focusing on testing without making a real service request.

### Do (45–55 minutes)

1. Start a fresh temporary session:

   ```bash
   aws login --profile aws-learning --region eu-west-1
   ```

2. In the AWS console, open **IAM → Roles → AWSLearningReadOnlyRole → Permissions → Add permissions → Create inline policy**. Choose the **JSON** editor and use:

   ```json
   {
     "Version": "2012-10-17",
     "Statement": [
       {
         "Sid": "ReadDiscoveryMetadata",
         "Effect": "Allow",
         "Action": [
           "ec2:DescribeRegions",
           "ec2:DescribeAvailabilityZones",
           "s3:ListAllMyBuckets"
         ],
         "Resource": "*"
       }
     ]
   }
   ```

3. Name the inline policy `AWSLearningDiscoveryReadOnlyPolicy` and create it.
4. Only after the inline policy exists, detach the AWS-managed `ReadOnlyAccess` policy from `AWSLearningReadOnlyRole`. Do not attach `AdministratorAccess` or any replacement broad policy.
5. Verify the role still performs the intended safe reads:

   ```bash
   aws ec2 describe-regions \
     --profile aws-learning-readonly \
     --query 'length(Regions)' \
     --output text \
     --no-cli-pager

   aws ec2 describe-availability-zones \
     --profile aws-learning-readonly \
     --region eu-west-1 \
     --query 'length(AvailabilityZones)' \
     --output text \
     --no-cli-pager

   aws s3api list-buckets \
     --profile aws-learning-readonly \
     --query 'length(Buckets)' \
     --output text \
     --no-cli-pager
   ```

6. Confirm the broad managed policy is gone and the inline policy is present:

   ```bash
   aws iam list-attached-role-policies \
     --role-name AWSLearningReadOnlyRole \
     --profile aws-learning \
     --query 'AttachedPolicies[].PolicyName' \
     --output json \
     --no-cli-pager

   aws iam list-role-policies \
     --role-name AWSLearningReadOnlyRole \
     --profile aws-learning \
     --output json \
     --no-cli-pager
   ```

   Expected: the first command returns an empty list; the second lists `AWSLearningDiscoveryReadOnlyPolicy`.
7. Resolve the account ID locally, validate it without printing it, and simulate four actions:

   ```bash
   AWS_LEARNING_ACCOUNT_ID="$(aws sts get-caller-identity \
     --profile aws-learning \
     --query Account \
     --output text)"

   if [[ "$AWS_LEARNING_ACCOUNT_ID" =~ ^[0-9]{12}$ ]]; then
     echo "Account ID captured successfully"
   else
     echo "Account ID capture failed"
   fi

   aws iam simulate-principal-policy \
     --policy-source-arn "arn:aws:iam::${AWS_LEARNING_ACCOUNT_ID}:role/AWSLearningReadOnlyRole" \
     --action-names \
       ec2:DescribeRegions \
       s3:ListAllMyBuckets \
       s3:PutObject \
       iam:CreateUser \
     --profile aws-learning \
     --query 'EvaluationResults[].{Action:EvalActionName,Decision:EvalDecision}' \
     --output table \
     --no-cli-pager
   ```

   Expected: the two discovery reads are `allowed`; the write and IAM-administration actions are `implicitDeny`.
8. In 3–5 sentences, explain:
   - why this policy is narrower than `ReadOnlyAccess`;
   - why `Resource` is `"*"` even though the policy follows least privilege;
   - the difference between limiting actions and limiting resources.
9. End the source session and clear the local variable:

   ```bash
   aws logout --profile aws-learning
   unset AWS_LEARNING_ACCOUNT_ID
   ```

### Safety and cost rules

- IAM policy changes and simulations create no billable infrastructure.
- Do not create a bucket, EC2 instance, access key, IAM user, or real S3 object.
- Do not submit account IDs, ARNs, credentials, tokens, login URLs, or raw identity output.
- If an intended read fails, stop and report the exact error; do not restore broad admin access as a shortcut.

### Submit for verification

```text
Day 6 submission
1. Inline policy created: yes/no; policy name only
2. AWS-managed ReadOnlyAccess detached: yes/no
3. Attached managed-policy list: empty/not empty; names only if not empty
4. Inline-policy list: [policy names only]
5. describe-regions: success/failure; count optional
6. describe-availability-zones: success/failure; count optional
7. list-buckets: success/failure; count optional
8. Policy simulation: DescribeRegions=[decision], ListAllMyBuckets=[decision], PutObject=[decision], CreateUser=[decision]
9. Why the new policy is narrower:
10. Why Resource is "*":
11. Action scope vs resource scope:
12. Source profile logged out: yes/no
13. Resources created besides the inline policy: none / list
14. Exact blocker, if any:
```

### Pass criteria

- [x] `AWSLearningDiscoveryReadOnlyPolicy` exists as an inline role policy.
- [x] Broad `ReadOnlyAccess` is detached and no other managed policy is attached.
- [x] The role profile successfully performs all three intended safe reads.
- [x] Simulation allows the two intended reads and implicitly denies `s3:PutObject` and `iam:CreateUser`.
- [x] Action scope, resource scope, and the justified use of `Resource: "*"` are explained accurately.
- [x] The source session is logged out and no workload resources, access keys, secrets, or identifiers are submitted.

### Submission review — September 18, 2026

- Verified: inline policy, removal of all managed policies, three successful discovery reads, expected policy-simulation results, source logout, and no additional resources.
- Item 11 correctly distinguishes action scope from resource scope.
- Correction needed: item 9 explains the inline policy's one-to-one ownership, not why its permissions are narrower. Item 10 must explain why these particular list/describe actions require `Resource: "*"` and how the restricted `Action` list still preserves least privilege.
- Follow-up verified: the new policy is narrower because it permits only three API actions; the wildcard resource is required because these list/describe actions do not support individual resource ARNs, while the restricted action list preserves least privilege.
- **Result:** Passed on September 18, 2026. All six verification criteria are complete.

---

## Day 7 — Automate a safe AWS access baseline check

**Date:** Friday, September 18, 2026  
**Due:** Before you stop studying today  
**Timebox:** 60–75 minutes  
**Outcome:** Write and test a Bash script that validates the Day 6 assumed role and reports only non-sensitive AWS baseline information.

### Learn (10 minutes maximum)

- Read the AWS CLI guide to [`--query` output filtering](https://docs.aws.amazon.com/cli/latest/userguide/cli-usage-filter.html), focusing on client-side JMESPath queries.
- Skim the Bash [`set` b-uiltin](https://www.gnu.org/software/bash/manual/html_node/The-Set-Builtin.html), focusing on `-e`, `-u`, and `pipefail`.

### Build (40–50 minutes)

1. Start a new temporary source session:

   ```bash
   aws login --profile aws-learning --region eu-west-1
   ```

2. In the root of this learning project, create `scripts/aws-baseline-check.sh` with the following implementation:

   ```bash
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
   ```

3. Make it executable and run a syntax check:

   ```bash
   chmod +x scripts/aws-baseline-check.sh
   bash -n scripts/aws-baseline-check.sh
   echo $?
   ```

   Expected exit status: `0`.
4. Run the successful path:

   ```bash
   ./scripts/aws-baseline-check.sh
   ```

   Confirm that it reports an assumed role, `eu-west-1`, and only counts—not an account ID or ARN.
5. Prove that it fails closed when given a nonexistent profile:

   ```bash
   if ./scripts/aws-baseline-check.sh missing-profile; then
     echo "FAIL: nonexistent profile was accepted"
   else
     echo "PASS: nonexistent profile was rejected"
   fi
   ```

6. Inspect the script before submission:

   ```bash
   rg -n '(Account|Arn|access.key|secret|token|credential)' \
     scripts/aws-baseline-check.sh || true
   ```

   The `Arn` query used only for local classification is expected. No command may print the ARN or account ID.
7. Write 3–5 sentences explaining:
   - what `set -euo pipefail` protects against;
   - why the ARN is captured, classified, unset, and never printed;
   - how `--query 'length(...)'` reduces output and accidental data exposure.
8. End the temporary source session:

   ```bash
   aws logout --profile aws-learning
   ```

### Safety and cost rules

- This assignment uses only read-only calls and creates no AWS resources.
- Never add `set -x`; shell tracing could expose sensitive values.
- Do not print, submit, log, or commit the caller ARN, account ID, credentials, session cache, or login URL.
- Do not add write permissions to make the script pass.

### Submit for verification

Save the script in the project so it can be inspected, then submit:

```text
Day 7 submission
1. Script path: scripts/aws-baseline-check.sh
2. bash -n exit status:
3. Successful run: yes/no
4. Caller type reported:
5. Configured Region reported:
6. Region count:
7. Standard AZ count:
8. S3 bucket count:
9. Missing-profile test result:
10. What set -euo pipefail protects against:
11. Why the ARN is never printed:
12. Why length(...) queries are used:
13. Source profile logged out: yes/no
14. AWS resources created: none / list
15. Exact blocker, if any:
```

### Pass criteria

- [x] The script exists at the required path, contains no embedded credentials or identifiers, and passes `bash -n`.
- [x] The successful path verifies the assumed role and expected Region without printing an ARN or account ID.
- [x] Region, standard-AZ, and bucket counts are returned through narrow JMESPath queries.
- [x] A nonexistent profile is rejected with a non-zero script exit status.
- [x] Strict mode, local ARN classification, and output minimization are explained accurately.
- [x] The source session is logged out and no AWS resources, access keys, secrets, or identifiers are created or submitted.

### Submission review — September 22, 2026

- The saved script passes `bash -n`. Its leading space before `#!` prevented direct execution; the shebang was corrected, and the direct missing-profile test now exits 1 with the expected message.
- The successful AWS run, returned counts, source logout, and absence of new resources are reported in the submission; they were not rerun after logout.
- Follow-up verified: `-e`, `-u`, and `pipefail` are explained accurately; the ARN is used only to classify the assumed role and is never printed; `length(...)` limits terminal output without changing the data AWS sends to the CLI.
- **Result:** Passed on September 22, 2026. All six verification criteria are complete.

---

## Day 8 — Launch, inspect, and clean up a Linux EC2 instance

**Date:** Tuesday, September 22, 2026  
**Due:** 11:00 PM Africa/Lagos  
**Timebox:** 75–90 minutes  
**Outcome:** Run one short-lived Amazon Linux instance, connect through Session Manager without an inbound SSH rule, inspect the Linux host, and verify cleanup.

### Learn (15 minutes maximum)

- Read the [EC2 getting-started guide](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EC2_GetStarted.html), focusing on the AMI, instance type, security group, storage, and cleanup steps.
- Read the [Session Manager instance-permission guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-getting-started-instance-profile.html) and the [network requirements](https://docs.aws.amazon.com/systems-manager/latest/userguide/setup-create-vpc.html). The instance needs an IAM role and outbound HTTPS access; Session Manager needs no inbound port.

### Build (50–60 minutes)

1. Sign in to the AWS console with your everyday MFA-protected identity. Select `eu-west-1` and confirm your budget alerts are active. Check the EC2 launch estimate and your account's Free Tier or credit eligibility before creating anything; do not assume the instance is free.
2. Create an EC2-trusted IAM role named `AWSLearningEC2SSMRole` and attach only `AmazonSSMManagedInstanceCore` for this lab. Use the role as the instance profile at launch; do not put credentials on the VM.
3. Launch **one** Amazon Linux 2023 instance named `cloud-eng-day8` with the tag `Project=cloud-eng-journey`. Choose a small instance type whose displayed cost fits your budget. Use one default-size root EBS volume with **Delete on termination** enabled and no extra volumes.
4. Use an existing default-VPC public subnet with outbound internet access so the SSM Agent can reach AWS over HTTPS. A public IPv4 address may incur a separate charge. Create a security group with **zero inbound rules** and outbound HTTPS access. Select **Proceed without a key pair** because this lab uses Session Manager. Do not create a NAT gateway, VPC endpoint, Elastic IP, or load balancer for this exercise.
5. Once both EC2 status checks pass, choose **Connect → Session Manager** in the EC2 console. If it is unavailable, check the instance role, SSM Agent status, subnet route, and outbound HTTPS access. Record the exact blocker rather than opening port 22.
6. In the Session Manager shell, run these read-only checks and record only the summarized results:

   ```bash
   cat /etc/os-release | sed -n '1,6p'
   uname -r
   whoami
   df -h /
   systemctl is-active amazon-ssm-agent
   ss -lnt | tail -n +2 | wc -l
   ```

7. Write 4–6 sentences explaining what the AMI, instance type, EBS root volume, security group, and instance role each control. Include why Session Manager works without an inbound SSH rule.

### Clean up (10–15 minutes)

1. End the Session Manager session and **terminate** the instance before stopping study today. A stopped instance can still incur EBS storage charges.
2. Confirm the instance shows `terminated`, the public IPv4 address is released, and no Day 8 EBS volume remains. The IAM role and empty-inbound security group may be reused for the next lab; record that they remain.
3. Sign out of the AWS console.

### Safety and cost rules

- Keep one instance running only for this lab. EC2, EBS, and public IPv4 usage may be billable; budget alerts are not hard spending caps.
- Do not add an inbound SSH or HTTP rule, create a key pair, copy credentials to the instance, or install a workload today.
- Do not submit instance IDs, account IDs, full ARNs, public or private IP addresses, hostnames, login URLs, or raw terminal screenshots.
- If there is no usable default VPC or the launch estimate does not fit the budget, stop before launch and submit the exact blocker. Do not create a NAT gateway as a workaround.

### Submit for verification

```text
Day 8 submission
1. Region:
2. Budget and launch estimate checked: yes/no; estimated short-lab cost:
3. AMI and instance type:
4. Project tag:
5. Instance role and attached policy:
6. Security-group inbound rule count; outbound HTTPS available:
7. Key pair created: yes/no
8. Session Manager connected: yes/no
9. Linux OS/version and kernel version:
10. Root filesystem usage percentage:
11. SSM Agent status; listening TCP socket count:
12. Explain AMI, instance type, EBS, security group, role, and no-inbound Session Manager (4–6 sentences):
13. Instance terminated: yes/no
14. Day 8 EBS volumes remaining: count
15. IAM role and security group retained for the next lab: yes/no
16. Exact blocker, if any:
```

### Pass criteria

- [x] The Region, budget, and launch estimate were checked before creating one short-lived instance.
- [x] The instance used Amazon Linux 2023, an EC2 instance role for Session Manager, no key pair, and a security group with no inbound rules.
- [x] Session Manager connected and the Linux checks returned summarized, non-sensitive evidence.
- [x] The explanation accurately distinguishes the AMI, instance type, EBS volume, security group, and instance role.
- [x] The instance was terminated and no Day 8 EBS volume remains.
- [x] The submission contains no credentials, account or instance identifiers, IP addresses, or secrets.

### Submission review — September 22, 2026

- Reported: `eu-west-1`, budget and estimate checked, Session Manager connected, SSM Agent active, root filesystem 21% used, two listening TCP sockets, instance terminated, and zero Day 8 EBS volumes remaining. These AWS observations were not independently rerun.
- Correction needed: confirm the AMI was Amazon Linux 2023 (its `ID_LIKE=fedora` field does not mean Fedora is the OS), supply the kernel version from the earlier `uname -r` result if available, and provide the missing project tag and security-group inbound/outbound details.
- A key pair was reported as created even though this Session Manager lab called for none. Confirm whether it was attached to the instance, remove the unused EC2 key pair and any downloaded private key, and report that cleanup. No new instance is needed solely to correct this.
- The explanation covers the AMI, EBS, and security group but omits instance type, instance role, and why Session Manager requires no inbound SSH rule. Expand it to 4–6 sentences.
- The estimated cost of zero is a reported launch estimate, not a confirmed final charge. Check Billing later for EC2, EBS, and public IPv4 usage.
- **Result:** Needs correction. Three of six criteria are supported by the submission.

### Follow-up review — September 22, 2026

- Amazon Linux 2023 was confirmed, and the earlier `Fedora` description was corrected. The reported Session Manager connection and summarized Linux checks now support the Linux-evidence criterion; the kernel release was not supplied.
- The latest reply says no key pair was created, correcting the earlier `yes`. The retained security group was reported to have no inbound or outbound rules. Zero outbound rules conflict with the reported successful Session Manager connection, so the group attached to the instance and its outbound rules need to be checked in the EC2 console.
- The expanded explanation covers all requested components, but describes `t3.medium` (2 vCPUs, 4 GiB) while the submitted instance type was `t3.micro` (2 vCPUs, 1 GiB). The project tag is still unreported.
- **Result:** Needs correction. Four of six criteria are supported by the submission.

### Final review — September 22, 2026

- Confirmed in the follow-up: the retained security group has zero inbound rules and an IPv4 `All traffic` outbound rule, which includes HTTPS; the `Project=cloud-eng-journey` tag was set; the launched instance type was `t3.micro`; and no key pair was created. The earlier `t3.medium` example was not the launched instance type; `t3.micro` provides 2 vCPUs and 1 GiB of memory.
- The kernel release was not recorded before termination. The other Linux results and successful Session Manager connection were submitted, so no new instance is required solely to recover that one value.
- **Result:** Passed on September 22, 2026. All six criteria are complete based on the submitted evidence.

