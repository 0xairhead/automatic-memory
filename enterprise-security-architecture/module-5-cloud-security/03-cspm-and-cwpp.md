# Module 5, Lesson 3: Cloud Security Posture Management (CSPM) & Cloud Workload Protection (CWPP)

## Table of Contents
- [Media Resources](#media-resources)
- [The Cloud Misconfiguration Problem](#the-cloud-misconfiguration-problem)
- [Cloud Security Posture Management (CSPM)](#cloud-security-posture-management-cspm)
  - [What CSPM Does](#what-cspm-does)
  - [Compliance Benchmarks](#compliance-benchmarks)
  - [Policy-as-Code](#policy-as-code)
  - [Drift Detection](#drift-detection)
- [Cloud Workload Protection Platforms (CWPP)](#cloud-workload-protection-platforms-cwpp)
  - [What CWPP Protects](#what-cwpp-protects)
  - [Runtime Protection](#runtime-protection)
  - [Vulnerability Management](#vulnerability-management)
- [CSPM vs CWPP: The Complete Picture](#cspm-vs-cwpp-the-complete-picture)
- [Cloud-Native Application Protection Platform (CNAPP)](#cloud-native-application-protection-platform-cnapp)
- [Tool Landscape](#tool-landscape)
- [Key Concepts to Remember](#key-concepts-to-remember)
- [Practice Questions](#practice-questions)
- [Next Up](#next-up)

---

Most cloud breaches aren't sophisticated hacks—they're misconfigurations. This lesson covers the tools that find them before attackers do.

---

## Media Resources

**Visual Guide:**

![Cloud Security Architecture](./assets/03-cspm-cwpp-architecture.png)

**Audio Lecture:**

🎧 <audio controls src="./assets/03-cspm-cwpp-audio.m4a"></audio>

---

## The Cloud Misconfiguration Problem

**Alarming statistics:**
- 65-70% of cloud security incidents involve misconfigurations
- Average time to detect a cloud breach: 197 days
- A single misconfigured S3 bucket can expose millions of records

**Real-world examples:**

| Company | Misconfiguration | Impact |
|---------|-----------------|--------|
| Capital One | Overly permissive IAM role + SSRF | 100M+ customer records |
| Facebook | Unencrypted S3 bucket | 540M user records exposed |
| Microsoft | Open Elasticsearch | 250M support records |
| Twitch | Server misconfiguration | 128GB source code leaked |

**The challenge:** Cloud environments change constantly. A developer can spin up resources in seconds. How do you ensure security keeps pace?

```
Manual Security Review:
Day 1: Audit complete ✓
Day 2: Developer creates new S3 bucket (public)
Day 3: Another bucket...
Day 30: 500 new resources, unknown security state

CSPM Approach:
Every resource → Automatically scanned → Instant alerting → Auto-remediation
```

---

## Cloud Security Posture Management (CSPM)

Think of CSPM as a continuous security audit that never sleeps.

---

### What CSPM Does

```
┌─────────────────────────────────────────────────────────────────────┐
│ CSPM Capabilities                                                   │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐               │
│  │ Discovery    │  │ Assessment   │  │ Remediation  │               │
│  │              │  │              │  │              │               │
│  │ • Inventory  │  │ • Benchmark  │  │ • Auto-fix   │               │
│  │ • Shadow IT  │  │ • Compliance │  │ • Playbooks  │               │
│  │ • Multi-cloud│  │ • Risk score │  │ • Ticketing  │               │
│  └──────┬───────┘  └──────┬───────┘  └──────┬───────┘               │
│         │                 │                 │                       │
│         └─────────────────┼─────────────────┘                       │
│                           ▼                                         │
│                  ┌──────────────┐                                   │
│                  │ Continuous   │                                   │
│                  │ Monitoring   │                                   │
│                  │              │                                   │
│                  │ • Drift      │                                   │
│                  │ • Alerts     │                                   │
│                  │ • Reporting  │                                   │
│                  └──────────────┘                                   │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

**Core functions:**

1. **Asset Discovery & Inventory**
   - What cloud resources exist across all accounts/subscriptions?
   - Which ones are internet-facing?
   - What's the relationship between resources?

2. **Configuration Assessment**
   - Is this S3 bucket public?
   - Is encryption enabled on this database?
   - Is MFA enforced for root accounts?

3. **Compliance Mapping**
   - How do we stack up against CIS benchmarks?
   - Are we meeting PCI-DSS requirements?
   - Generate audit reports automatically

4. **Risk Prioritization**
   - Which misconfigurations are most critical?
   - What's internet-exposed vs internal-only?
   - Correlate with threat intelligence

---

### Compliance Benchmarks

CSPM tools check against industry-standard benchmarks:

**CIS Benchmarks (Center for Internet Security):**
```
CIS AWS Foundations Benchmark v1.5

1. Identity and Access Management
   1.1  Avoid the use of root account
   1.2  Ensure MFA is enabled for root account
   1.3  Ensure credentials unused for 90+ days are disabled
   1.4  Ensure access keys are rotated every 90 days
   ...

2. Logging
   2.1  Ensure CloudTrail is enabled in all regions
   2.2  Ensure log file validation is enabled
   ...

3. Monitoring
   3.1  Ensure unauthorized API calls trigger alerts
   ...
```

**Mapping to compliance frameworks:**
```
CIS Control 1.4          ────▶  PCI-DSS 8.2.4
(Rotate access keys)            (Change passwords every 90 days)

CIS Control 2.1          ────▶  HIPAA § 164.312(b)
(Enable CloudTrail)             (Audit controls)

CIS Control 3.x          ────▶  SOC 2 CC6.1
(Monitoring)                    (Security monitoring)
```

---

### Policy-as-Code

Modern CSPM enables you to define security policies as code, versioned alongside your infrastructure.

**Open Policy Agent (OPA) / Rego:**
```rego
# Deny public S3 buckets
deny[msg] {
    resource := input.resource.aws_s3_bucket[name]
    resource.acl == "public-read"
    msg := sprintf("S3 bucket '%s' has public read access", [name])
}

# Require encryption on EBS volumes
deny[msg] {
    resource := input.resource.aws_ebs_volume[name]
    not resource.encrypted
    msg := sprintf("EBS volume '%s' is not encrypted", [name])
}
```

**HashiCorp Sentinel:**
```python
# Require all EC2 instances to have approved AMIs
import "tfplan/v2" as tfplan

allowed_amis = [
    "ami-0123456789abcdef0",
    "ami-0987654321fedcba0"
]

ec2_instances = filter tfplan.resource_changes as _, rc {
    rc.type is "aws_instance" and
    rc.mode is "managed" and
    rc.change.actions contains "create"
}

main = rule {
    all ec2_instances as _, instance {
        instance.change.after.ami in allowed_amis
    }
}
```

**Azure Policy:**
```json
{
  "if": {
    "allOf": [
      {
        "field": "type",
        "equals": "Microsoft.Storage/storageAccounts"
      },
      {
        "field": "Microsoft.Storage/storageAccounts/allowBlobPublicAccess",
        "equals": true
      }
    ]
  },
  "then": {
    "effect": "deny"
  }
}
```

---

### Drift Detection

Infrastructure should match its defined state. Drift is when reality diverges from intention.

```
Terraform State (Desired):          Actual AWS Config:
┌─────────────────────────┐         ┌─────────────────────────┐
│ S3 Bucket: my-bucket    │         │ S3 Bucket: my-bucket    │
│ - Public Access: false  │   !=    │ - Public Access: true   │ ← DRIFT!
│ - Encryption: AES-256   │         │ - Encryption: AES-256   │
│ - Versioning: enabled   │         │ - Versioning: enabled   │
└─────────────────────────┘         └─────────────────────────┘
```

**Why drift happens:**
- Console changes by developers bypassing IaC
- Emergency fixes not backported to code
- Third-party tools modifying resources
- Malicious actors

**CSPM drift detection workflow:**
```
1. Scan cloud environment
2. Compare against:
   - Terraform/CloudFormation state
   - Previous baseline
   - Policy definitions
3. Detect deviations
4. Alert or auto-remediate
```

---

## Cloud Workload Protection Platforms (CWPP)

While CSPM focuses on configuration, CWPP protects the actual workloads (VMs, containers, serverless).

---

### What CWPP Protects

```
CWPP Protection Layers:

┌─────────────────────────────────────────────────────────────────────┐
│ Workload Types                                                      │
├─────────────────────────────────────────────────────────────────────┤
│  ┌───────────────┐  ┌───────────────┐  ┌───────────────┐            │
│  │ Virtual       │  │ Containers    │  │ Serverless    │            │
│  │ Machines      │  │               │  │               │            │
│  │               │  │ ┌───────────┐ │  │ ┌───────────┐ │            │
│  │ ┌───────────┐ │  │ │ Container │ │  │ │ Function  │ │            │
│  │ │    App    │ │  │ │           │ │  │ │           │ │            │
│  │ ├───────────┤ │  │ ├───────────┤ │  │ └───────────┘ │            │
│  │ │    OS     │ │  │ │ Container │ │  │               │            │
│  │ └───────────┘ │  │ │ Runtime   │ │  │ No OS/Runtime │            │
│  │               │  │ ├───────────┤ │  │ to manage     │            │
│  └───────────────┘  │ │ Host OS   │ │  │               │            │
│                     │ └───────────┘ │  └───────────────┘            │
│                     └───────────────┘                               │
└─────────────────────────────────────────────────────────────────────┘
```

---

### Runtime Protection

CWPP monitors workloads while they're running:

**Behavioral Analysis:**
```
Normal behavior baseline:           Detected anomaly:
┌─────────────────────────┐         ┌─────────────────────────┐
│ Web server process      │         │ Web server process      │
│ - Reads /var/www        │         │ - Reads /etc/shadow     │ ← ALERT!
│ - Writes /var/log       │         │ - Opens /dev/tcp        │ ← ALERT!
│ - Connects to port 443  │         │ - Spawns /bin/sh        │ ← ALERT!
│ - Spawns child workers  │         │                         │
└─────────────────────────┘         └─────────────────────────┘
```

**File Integrity Monitoring (FIM):**
```
Monitored files:
/etc/passwd          - CHANGED → Alert!
/etc/ssh/sshd_config - CHANGED → Alert!
/usr/bin/sudo        - Hash mismatch → Critical Alert!
```

**Network Monitoring:**
```
Expected connections:        Unexpected:
App → Database (5432)        App → 185.x.x.x:4444  ← C2 callback?
App → Redis (6379)           App → TOR exit node   ← Data exfil?
App → S3 endpoint            App → Crypto mining   ← Cryptojacker?
```

---

### Vulnerability Management

CWPP continuously scans workloads for vulnerabilities:

```
┌─────────────────────────────────────────────────────────────────────┐
│ Vulnerability Scan Results                                          │
├─────────────────────────────────────────────────────────────────────┤
│ Image: my-app:v2.3.1                                                │
│ Base: ubuntu:22.04                                                  │
├──────────────────────────────┬────────────────────────────┬─────────┤
│ Vulnerability                │ Package                    │Severity │
├──────────────────────────────┼────────────────────────────┼─────────┤
│ CVE-2023-44487 (HTTP/2 DoS)  │ libnghttp2-14              │ HIGH    │
│ CVE-2023-4911 (glibc LPE)    │ libc6                      │ CRITICAL│
│ CVE-2023-38545 (curl heap)   │ curl                       │ HIGH    │
│ CVE-2022-48174 (busybox)     │ busybox                    │ CRITICAL│
└──────────────────────────────┴────────────────────────────┴─────────┘
│ Recommendations:                                                    │
│ - Update base image to ubuntu:22.04.3                               │
│ - Upgrade libnghttp2 to 1.52.0-1ubuntu0.1                           │
│ - Review if busybox is necessary                                    │
└─────────────────────────────────────────────────────────────────────┘
```

**Shift-left integration:**
```
Developer commits → CI/CD Pipeline → Vulnerability Scan → Pass/Fail Gate

┌──────────────────────────────────────────────┐
│ Pipeline Stage: Security Scan                │
├──────────────────────────────────────────────┤
│ ✓ No CRITICAL vulnerabilities                │
│ ✓ No HIGH vulnerabilities in base image      │
│ ⚠ 3 MEDIUM vulnerabilities (allowed)         │
│                                              │
│ Result: PASSED                               │
└──────────────────────────────────────────────┘
```

---

## CSPM vs CWPP: The Complete Picture

| Aspect | CSPM | CWPP |
|--------|------|------|
| **Focus** | Configuration & posture | Workload runtime security |
| **What it checks** | Is MFA enabled? Is encryption on? | Is this process malicious? |
| **When it acts** | Before/during deployment | During runtime |
| **Protects against** | Misconfiguration, compliance gaps | Malware, intrusions, vulnerabilities |
| **Example finding** | "S3 bucket is public" | "Suspicious process spawned" |

**They're complementary:**
```
┌─────────────────────────────────────────────────────────────────────┐
│ Complete Cloud Security                                             │
│                                                                     │
│  CSPM                          CWPP                                 │
│  ┌─────────────────────┐       ┌─────────────────────┐              │
│  │ Configuration       │       │ Runtime             │              │
│  │ • Is it set up      │       │ • Is it behaving    │              │
│  │   correctly?        │       │   correctly?        │              │
│  │                     │       │                     │              │
│  │ Pre-deployment      │       │ Post-deployment     │              │
│  │ • Prevent misconfig │       │ • Detect attacks    │              │
│  └─────────────────────┘       └─────────────────────┘              │
│            │                            │                           │
│            └────────────┬───────────────┘                           │
│                         ▼                                           │
│              ┌─────────────────────┐                                │
│              │ CNAPP               │                                │
│              │ (Cloud-Native App   │                                │
│              │ Protection Platform)│                                │
│              │                     │                                │
│              │ Unified: CSPM + CWPP│                                │
│              │ + CIEM + IaC Scan   │                                │
│              └─────────────────────┘                                │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

---

## Cloud-Native Application Protection Platform (CNAPP)

The industry is moving away from separate tools (CSPM for config, CWPP for workloads) toward a unified platform: **CNAPP**.

### The Problem: Tool Sprawl & Context Fatigue

Using separate tools creates "noise" without context:
1.  **CSPM says:** "Security Group open on Port 22" (High Severity)
2.  **CWPP says:** "Vulnerable Package `lib-ssh` found" (High Severity)
3.  **CIEM says:** "EC2 instance has Admin privileges" (High Severity)

Security analysts have to manually correlate these. If they don't, they might waste time fixing a vulnerability on a machine that has no permissions and is not exposed to the internet.

### The Solution: Contextual Risk

CNAPP unifies these signals to focus on the **Attack Path**:

```
Attack Path Visualization:

[ Internet ] ──(port 22)──▶ [ EC2 Instance ] ──(Admin Role)──▶ [ S3 Bucket ]
     ▲                            ▲                                 ▲
(CSPM Finding)              (CWPP Finding)                    (Target)
  "Exposed"                  "Vulnerable"                   "Sensitive Data"
```

**The CNAPP Verdict:**
"This is not just 3 separate 'High' alerts. This is a **CRITICAL** attack path where an expliotable vulnerability directly exposes sensitive data."

### CNAPP Pillars

A true CNAPP consolidates multiple capabilities:

1.  **CSPM** (Infrastructure Configuration)
2.  **CWPP** (Workload Runtime)
3.  **CIEM** (Cloud Infrastructure Entitlement Management) - *The Identity piece*
4.  **KSPM** (Kubernetes Security Posture Management)
5.  **DSPM** (Data Security Posture Management) - *Scanning actual data for PII/Secrets*
6.  **IaC Scanning** (Shift-left security)

---

## Tool Landscape

**Cloud-Native Tools:**
| Provider | CSPM | CWPP |
|----------|------|------|
| AWS | Security Hub, Config | GuardDuty, Inspector |
| Azure | Defender for Cloud | Defender for Cloud |
| GCP | Security Command Center | Security Command Center |

**Third-Party Platforms:**
| Vendor | Strengths |
|--------|-----------|
| Wiz | Agentless, fast scanning, graph-based context |
| Prisma Cloud (Palo Alto) | Comprehensive CNAPP, strong compliance |
| Lacework | Behavioral analytics, anomaly detection |
| Orca Security | Agentless, side-scanning technology |
| CrowdStrike | Strong EDR heritage, Falcon Cloud Security |
| Aqua Security | Container/K8s specialty |
| Sysdig | Container security, Falco-based |

**Open Source:**
- **Prowler** - AWS/Azure/GCP security assessments
- **ScoutSuite** - Multi-cloud security auditing
- **CloudSploit** - Cloud security scans
- **Falco** - Runtime security (CNCF project)
- **Trivy** - Vulnerability scanner

---

## Key Concepts to Remember

1. **Misconfigurations are the #1 cloud risk** - CSPM catches them before attackers
2. **Policy-as-Code enables guardrails** - Prevent violations, don't just detect
3. **Drift is dangerous** - Continuous monitoring catches unauthorized changes
4. **CWPP protects what's running** - Behavioral analysis catches attacks in progress
5. **CNAPP unifies everything** - Modern platforms combine CSPM, CWPP, and more
6. **Shift-left security** - Scan in CI/CD, not just production

---

## Practice Questions

**Q1:** Your CSPM tool reports 2,847 findings across your AWS environment. The security team is overwhelmed. How would you prioritize remediation?

<details>
<summary>View Answer</summary>

**Prioritization framework:**

1. **Internet Exposure + Severity:**
   - Critical findings on internet-facing resources → Fix immediately
   - High severity on internal resources → Fix within days
   - Medium/Low → Queue for regular patching

2. **Prioritization matrix:**
   ```
   Priority 1 (Fix now):
   - Public S3 buckets with sensitive data
   - Security groups allowing 0.0.0.0/0 on sensitive ports
   - Root account without MFA
   - Publicly exposed databases

   Priority 2 (Fix this week):
   - Unencrypted resources with sensitive data
   - Overly permissive IAM policies
   - Missing logging/monitoring

   Priority 3 (Fix this month):
   - Best practice deviations (encryption at rest)
   - Unused credentials
   - Non-critical compliance gaps

   Priority 4 (Backlog):
   - Informational findings
   - Dev/test environment findings
   ```

3. **Automated triage:**
   - Use risk scoring that factors in: severity, exposure, data sensitivity, compliance impact
   - Group findings by root cause (one fix may resolve 50 findings)
   - Assign to teams based on resource ownership

4. **Quick wins:**
   - Enable auto-remediation for low-risk, high-confidence fixes
   - Use AWS Config rules with automatic remediation
   - Implement preventive controls (SCPs, Azure Policy) to stop new violations

</details>

**Q2:** A developer argues that CWPP agents add overhead and latency to their containers. How do you address this concern while maintaining security?

<details>
<summary>View Answer</summary>

**Acknowledge the concern, then address it:**

1. **Modern CWPP options:**
   - **Agentless scanning:** Tools like Wiz, Orca scan cloud snapshots without agents
   - **eBPF-based agents:** Falco, Cilium have minimal overhead (1-3% CPU)
   - **Sidecar vs DaemonSet:** DaemonSet approach shares one agent across all pods

2. **Measured impact:**
   ```
   Traditional agent overhead: 5-15% CPU
   Modern eBPF agent overhead: 1-3% CPU
   Agentless: 0% runtime overhead
   ```

3. **Risk-based approach:**
   - Production: Full CWPP coverage (the overhead is worth it)
   - Staging: Full coverage (catch issues before prod)
   - Dev: Agentless scanning, or sampling

4. **Compromise architecture:**
   ```
   ┌─────────────────────────────────────────────┐
   │ Kubernetes Cluster                          │
   │                                             │
   │  DaemonSet: CWPP Agent (1 per node)         │
   │  • Not in each pod (lower overhead)         │
   │  • Uses eBPF for syscall monitoring         │
   │  • Resource limits defined                  │
   │                                             │
   │  Sidecar (high-security pods only):         │
   │  • Financial transactions                   │
   │  • PII processing                           │
   └─────────────────────────────────────────────┘
   ```

5. **Business case:**
   - 2% performance overhead vs potential breach cost
   - Container attacks are increasing 300%+ year over year
   - Regulatory requirements may mandate runtime protection

</details>

**Q3:** You're evaluating CSPM tools for a multi-cloud environment (AWS, Azure, GCP). What key capabilities would you require, and how would you structure the evaluation?

<details>
<summary>View Answer</summary>

**Key capability requirements:**

1. **Multi-cloud coverage:**
   - Single pane of glass for all three clouds
   - Normalized findings (not just raw cloud APIs)
   - Consistent policy language across clouds

2. **Technical capabilities:**
   ```
   Must Have:
   ✓ Agentless deployment option
   ✓ CIS benchmark support for all clouds
   ✓ Custom policy creation (OPA/Rego or similar)
   ✓ API-first architecture
   ✓ IaC scanning (Terraform, CloudFormation, ARM)
   ✓ RBAC with SSO integration
   ✓ Auto-remediation capabilities

   Nice to Have:
   ○ Attack path analysis
   ○ Graph-based visualization
   ○ Kubernetes security
   ○ Identity risk analysis (CIEM)
   ○ Data security posture (DSPM)
   ```

3. **Evaluation structure:**

   **Phase 1: RFI (2 weeks)**
   - Send requirements questionnaire
   - Review vendor responses
   - Shortlist 3-4 vendors

   **Phase 2: POC (4 weeks per vendor)**
   - Connect to representative accounts in each cloud
   - Evaluate:
     - Time to value (deployment ease)
     - Finding accuracy (false positive rate)
     - Coverage completeness
     - Alert actionability
     - Integration with existing tools (SIEM, ticketing)

   **Phase 3: Scoring**
   | Criteria | Weight |
   |----------|--------|
   | Multi-cloud coverage | 25% |
   | Finding accuracy | 20% |
   | Ease of use | 15% |
   | Integration capabilities | 15% |
   | Remediation automation | 15% |
   | Total cost of ownership | 10% |

4. **Red flags:**
   - Agent required for basic scanning
   - Cloud-specific consoles (not unified)
   - No custom policy support
   - Slow scan times (hours vs minutes)

</details>

**Q4:** Your CWPP tool detects a process on an EC2 instance attempting to access the instance metadata service (169.254.169.254) and then making calls to an external IP. What attack does this suggest, and how would you respond?

<details>
<summary>View Answer</summary>

**Attack identification: SSRF leading to credential theft**

This is a classic Server-Side Request Forgery (SSRF) attack pattern, similar to the Capital One breach:

```
Attack chain:
1. Attacker exploits SSRF vulnerability in application
2. Application makes request to http://169.254.169.254/latest/meta-data/
3. Attacker retrieves IAM role credentials from metadata service
4. Credentials sent to external IP (attacker's server)
5. Attacker uses stolen credentials from outside your environment
```

**Immediate response:**

1. **Contain:**
   - Isolate the instance (change security group to deny all)
   - Don't terminate yet (preserve forensic evidence)
   - Block the external IP at network level (NACL, firewall)

2. **Investigate:**
   - Check CloudTrail for IAM role usage from unusual IPs
   - Review what permissions the role has
   - Identify the SSRF vulnerability in the application
   - Check for lateral movement

3. **Remediate:**
   - Rotate/invalidate the IAM role credentials
   - Patch the SSRF vulnerability
   - Enable IMDSv2 (requires session tokens, blocks SSRF):
     ```bash
     aws ec2 modify-instance-metadata-options \
       --instance-id i-1234567890abcdef0 \
       --http-tokens required \
       --http-endpoint enabled
     ```

4. **Prevent (long-term):**
   - Enforce IMDSv2 across all instances (SCP or Config rule)
   - Implement WAF with SSRF rules
   - Network segmentation (app servers shouldn't reach metadata if not needed)
   - Least privilege IAM roles
   - Enable GuardDuty for automated detection

**IMDSv2 explanation:**
```
IMDSv1 (vulnerable):
curl http://169.254.169.254/latest/meta-data/  ← Direct access, SSRF works

IMDSv2 (protected):
TOKEN=$(curl -X PUT http://169.254.169.254/latest/api/token \
  -H "X-aws-ec2-metadata-token-ttl-seconds: 21600")
curl http://169.254.169.254/latest/meta-data/ \
  -H "X-aws-ec2-metadata-token: $TOKEN"  ← Requires PUT + headers, SSRF blocked
```

</details>

**Q5:** You are reviewing a Rego policy for your OPA-based CSPM. What does the following policy enforce, and why might it be critical for a production environment?

```rego
deny[msg] {
    input.resource.type == "aws_security_group"
    some i
    input.resource.ingress[i].cidr == "0.0.0.0/0"
    input.resource.ingress[i].from_port == 22
    msg = "SSH access from the internet is not allowed"
}
```

<details>
<summary>View Answer</summary>

**Answer:**
It denies any AWS Security Group that allows ingress on port 22 (SSH) from `0.0.0.0/0` (any IP on the internet).

**Why it's critical:**
Leaving SSH open to the world invites constant brute-force attacks. If a credential is weak or compromised, attackers can gain direct shell access. SSH should always be restricted to a VPN, bastion host, or specific trusted IP ranges.
</details>

**Q6:** Your CSPM alerts you to "Drift Detected" on a critical production database. The alert shows `Storage Encrypted: False` (Actual) vs `Storage Encrypted: True` (Expected/State). What does this discrepancy indicate about how the change was made, and what is the risk?

<details>
<summary>View Answer</summary>

**Indication:**
The change was likely made **manually** (via the Cloud Console or CLI) or by a separate unauthorized script, bypassing the official Infrastructure-as-Code (Terraform/CloudFormation) pipeline. If it were changed in code, the "Expected" state would have updated too.

**Risk:**
1.  **Security/Compliance:** Sensitive data is now unencrypted at rest, violating policies/regulations (e.g., HIPAA, PCI-DSS).
2.  **Process:** The "source of truth" (IaC) is broken. The next automated deployment might accidentally revert this setting (good in this case) or fail due to conflict.
</details>

**Q7:** You currently use Tool A for CSPM (scanning config) and Tool B for CWPP (scanning runtime). You are considering moving to a unified CNAPP (Cloud-Native Application Protection Platform). How does a CNAPP use "context" to prioritize risk better than your separate tools?

<details>
<summary>View Answer</summary>

**Contextual Prioritization:**
Separate tools see risks in isolation. A CNAPP connects the dots:

*   **CSPM view:** "This VM has port 80 open to the internet." (High Risk?)
*   **CWPP view:** "This VM has the 'Log4Shell' vulnerability in a jar file." (Critical Risk?)

A CNAPP combines these:
"This VM has a Critical vulnerability AND it is exposed to the internet." -> **EMERGENCY Priority.**

Conversely, if the VM is purely internal and the vulnerable library is never loaded into memory (runtime), the CNAPP might downgrade the priority, saving the team from chasing a "ghost" risk.
</details>

**Q8:** During a CI/CD build pipeline, your scanner detects a "Critical" vulnerability in a container image. A developer argues that they should be allowed to deploy anyway and "fix it next week" to meet a deadline. Using the "Shift-Left" concept and the cost of defects, what is the counter-argument?

<details>
<summary>View Answer</summary>

**Counter-argument:**
In the "Shift-Left" model, fixing a defect (security bug) in the **Build/Design phase** is significantly cheaper (often 100x cheaper) than fixing it in **Production**.

*   **Build Phase:** You just update the package version in the Dockerfile and rebuild (Minutes).
*   **Production Phase:** You have to patch specific running servers, deal with potential downtime, handle customer notifications if breached, and perform emergency regression testing (Days/Weeks).
*   **Risk:** Deploying a known Critical vulnerability intentionally invalidates compliance (SOC2/PCI) and leaves the door wide open for immediate exploitation.
</details>

**Q9:** Your team selects an **Agentless-only** CWPP tool because it's easier to deploy (no agents on servers). What specific security capability have you likely sacrificed compared to an agent-based approach, and why does it matter?

<details>
<summary>View Answer</summary>

**Sacrificed Capability: Real-time Runtime Protection (Blocking).**

*   **Agentless** scanners work by taking "snapshots" of disk volumes. They can see what is *installed* (vulnerabilities) and what configurations are set.
*   **Agents** run inside the kernel/OS. They can see what is *executing* in memory right now.

**Why it matters:**
If a piece of "fileless" malware runs purely in memory, or if an attacker exploits a zero-day vulnerability to run a malicious command *right now*, an Agentless snapshot (which might run every 24 hours) will miss it completely. An Agent can block that process execution immediately.
</details>

---

## Next Up

In Lesson 4, we'll dive into **Container, Kubernetes & Serverless Security** — securing the modern workloads that CWPP protects, with specific techniques for each platform!
