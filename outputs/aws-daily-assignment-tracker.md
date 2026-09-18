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
5. Meaningful progress produces a build-in-public post draft; milestone completion produces a retrospective post.
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
| 5 | Sep 18 | Create and assume a read-only IAM role | 11:00 PM | ⬜ Assigned | Pending | Not yet |

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

- [ ] `AWSLearningReadOnlyRole` exists with `ReadOnlyAccess` and all three tags.
- [ ] `aws-learning-readonly` successfully assumes the role through `aws-learning`.
- [ ] The role can perform the safe read-only S3 listing.
- [ ] Policy simulation allows `s3:ListAllMyBuckets` and denies `s3:PutObject` without executing either action.
- [ ] User/role, permission/trust policy, and implicit/explicit deny are explained accurately.
- [ ] The source session is logged out and no access keys, workload resources, secrets, or identifiers are submitted.
