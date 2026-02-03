# Module 5, Lesson 6: Multi-Cloud Security, CASB & Architecture Patterns

## Table of Contents
- [Media Resources](#media-resources)
- [Multi-Cloud & Hybrid Security](#multi-cloud--hybrid-security)
  - [Why Multi-Cloud?](#why-multi-cloud)
  - [Consistent Policy Enforcement](#consistent-policy-enforcement)
  - [Centralized Identity Management](#centralized-identity-management)
  - [Unified Logging & Monitoring](#unified-logging--monitoring)
  - [Hybrid Connectivity Security](#hybrid-connectivity-security)
- [Cloud Access Security Brokers (CASB)](#cloud-access-security-brokers-casb)
  - [What is CASB?](#what-is-casb)
  - [CASB Deployment Modes](#casb-deployment-modes)
  - [CASB Use Cases](#casb-use-cases)
- [Cloud Security Architecture Patterns](#cloud-security-architecture-patterns)
  - [Landing Zone Design](#landing-zone-design)
  - [Account/Subscription Structure](#accountsubscription-structure)
  - [Security Baseline Automation](#security-baseline-automation)
  - [Immutable Infrastructure](#immutable-infrastructure)
  - [GitOps for Security](#gitops-for-security)
- [Bringing It All Together](#bringing-it-all-together)
- [Key Concepts to Remember](#key-concepts-to-remember)
- [Practice Questions](#practice-questions)
- [Module 5 Summary](#module-5-summary)

---

This final lesson ties together everything you've learned into enterprise-grade architecture patterns.

---

## Media Resources

**Visual Guide:**

![Enterprise Cloud Security Architecture](./assets/06-multicloud-casb-architecture.png)

**Audio Lecture:**

🎧 [🎧 Listen to Audio](./assets/06-multicloud-casb-audio.m4a)

---

## Multi-Cloud & Hybrid Security

---

### Why Multi-Cloud?

Organizations adopt multi-cloud for various reasons:

```
┌─────────────────────────────────────────────────────────────────────┐
│ Multi-Cloud Drivers                                                 │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  Strategic:                        Practical:                       │
│  ├── Avoid vendor lock-in          ├── Best-of-breed services       │
│  ├── Negotiating leverage          ├── M&A inheritance              │
│  ├── Regulatory requirements       ├── Developer preference         │
│  └── Business continuity           └── Geographic requirements      │
│                                                                     │
│  Reality check:                                                     │
│  Multi-cloud increases complexity significantly                     │
│  Security teams must now be experts in 2-3 platforms                │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

**The multi-cloud security challenge:**

```
Single Cloud:                    Multi-Cloud:
┌──────────────────┐             ┌────────────────────┐
│ AWS              │             │ AWS    Azure GCP   │
│ ├── IAM          │             │ ├── IAM   AD  IAM  │
│ ├── VPC          │             │ ├── VPC VNet VPC   │
│ ├── CloudTrail   │             │ ├── CT  Monitor CL │
│ ├── GuardDuty    │             │ ├── GD  Sent  SCC  │
│ └── Config       │             │ └── Cfg Pol  ???   │
│                  │             │                    │
│ One set of skills│             │ 3x the complexity  │
│ One set of tools │             │ 3x the tools       │
└──────────────────┘             └────────────────────┘
```

---

### Consistent Policy Enforcement

The key to multi-cloud security is **abstraction** — define policies once, enforce everywhere.

**Policy-as-Code across clouds:**

```
┌─────────────────────────────────────────────────────────────────────┐
│ Unified Policy Architecture                                         │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  ┌─────────────────────────────────────────────────────────────┐    │
│  │ Central Policy Repository (Git)                             │    │
│  │                                                             │    │
│  │ policies/                                                   │    │
│  │ ├── common/                                                 │    │
│  │ │   ├── encryption-required.rego                            │    │
│  │ │   ├── no-public-storage.rego                              │    │
│  │ │   └── mfa-required.rego                                   │    │
│  │ └── cloud-specific/                                         │    │
│  │     ├── aws/                                                │    │
│  │     ├── azure/                                              │    │
│  │     └── gcp/                                                │    │
│  └─────────────────────────────────────────────────────────────┘    │
│                              │                                      │
│              ┌───────────────┼───────────────┐                      │
│              ▼               ▼               ▼                      │
│       ┌───────────┐   ┌───────────┐   ┌───────────┐                 │
│       │ AWS       │   │ Azure     │   │ GCP       │                 │
│       │ Config/   │   │ Policy    │   │ Org       │                 │
│       │ SCP       │   │           │   │ Policy    │                 │
│       └───────────┘   └───────────┘   └───────────┘                 │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

**Example: No public storage (OPA Rego):**

```rego
# Abstract policy - works for any cloud
package storage.public_access

deny[msg] {
    storage := input.resource
    storage.type == "storage_bucket"
    storage.public_access == true
    msg := sprintf("Storage %v must not be publicly accessible", [storage.name])
}
```

**Tools for cross-cloud policy:**
- **Open Policy Agent (OPA)** - Cloud-agnostic policy engine
- **HashiCorp Sentinel** - Policy as code for Terraform
- **Checkov** - IaC security scanning for all clouds
- **CSPM platforms** - Prisma Cloud, Wiz, Lacework

---

### Centralized Identity Management

**The goal:** Single source of truth for identity across all clouds.

```
┌─────────────────────────────────────────────────────────────────────┐
│ Centralized Identity Architecture                                   │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│                    ┌─────────────────┐                              │
│                    │ Central IdP     │                              │
│                    │ (Okta, Azure AD,│                              │
│                    │  Google, Ping)  │                              │
│                    └────────┬────────┘                              │
│                             │                                       │
│              ┌──────────────┼──────────────┐                        │
│              │ SAML/OIDC    │ SAML/OIDC    │ SAML/OIDC              │
│              ▼              ▼              ▼                        │
│       ┌───────────┐  ┌───────────┐  ┌───────────┐                   │
│       │ AWS IAM   │  │ Azure AD  │  │ GCP IAM   │                   │
│       │ Identity  │  │ (Guest or │  │ Workload  │                   │
│       │ Center    │  │  B2B)     │  │ Identity  │                   │
│       └───────────┘  └───────────┘  └───────────┘                   │
│                                                                     │
│  Benefits:                                                          │
│  • Single user lifecycle (provision/deprovision)                    │
│  • Consistent MFA policy                                            │
│  • Centralized access reviews                                       │
│  • Single audit trail for authentication                            │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

**SCIM for automated provisioning:**

```
User lifecycle in central IdP:

1. HR creates user → IdP → SCIM → AWS, Azure, GCP accounts created
2. User joins team → IdP group → SCIM → Cloud permissions granted
3. User leaves company → IdP disabled → SCIM → All cloud access revoked

No manual intervention needed!
```

---

### Unified Logging & Monitoring

**Centralized SIEM architecture:**

```
┌─────────────────────────────────────────────────────────────────────┐
│ Multi-Cloud Logging Architecture                                    │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  AWS                    Azure                  GCP                  │
│  ├── CloudTrail         ├── Activity Log       ├── Cloud Audit      │
│  ├── VPC Flow Logs      ├── NSG Flow Logs      ├── VPC Flow Logs    │
│  ├── GuardDuty          ├── Defender           ├── SCC              │
│  └── Config             └── Monitor            └── Asset Inventory  │
│       │                      │                      │               │
│       └──────────────────────┼──────────────────────┘               │
│                              ▼                                      │
│                    ┌─────────────────┐                              │
│                    │ Log Aggregation │                              │
│                    │ (S3, EventHub,  │                              │
│                    │  Pub/Sub)       │                              │
│                    └────────┬────────┘                              │
│                             │                                       │
│                             ▼                                       │
│                    ┌─────────────────┐                              │
│                    │     SIEM        │                              │
│                    │ (Splunk, Elastic│                              │
│                    │  Sentinel, etc.)│                              │
│                    └────────┬────────┘                              │
│                             │                                       │
│              ┌──────────────┼──────────────┐                        │
│              ▼              ▼              ▼                        │
│       ┌───────────┐  ┌───────────┐  ┌───────────┐                   │
│       │ Detection │  │ Dashboards│  │ Compliance│                   │
│       │ Rules     │  │           │  │ Reports   │                   │
│       └───────────┘  └───────────┘  └───────────┘                   │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

**Normalizing log formats:**

```
AWS CloudTrail event:                    Normalized format:
{                                        {
  "eventSource": "s3.amazonaws.com",       "source": "aws",
  "eventName": "GetObject",                "service": "storage",
  "userIdentity": {...},                   "action": "read",
  "sourceIPAddress": "1.2.3.4"             "actor": "user@company.com",
}                                          "source_ip": "1.2.3.4"
                                         }

Azure Activity Log event:                Normalized format:
{                                        {
  "operationName": "Get Blob",             "source": "azure",
  "caller": "user@company.com",            "service": "storage",
  "callerIpAddress": "1.2.3.4"             "action": "read",
}                                          "actor": "user@company.com",
                                           "source_ip": "1.2.3.4"
                                         }
```

---

### Hybrid Connectivity Security

**Connecting on-premises to cloud securely:**

```
┌─────────────────────────────────────────────────────────────────────┐
│ Hybrid Connectivity Options                                         │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  Option 1: VPN (encrypted over internet)                            │
│  ┌─────────────┐    IPsec Tunnel    ┌─────────────┐                 │
│  │ On-Premises │ ================== │ Cloud VPC   │                 │
│  │ Firewall    │    (encrypted)     │ VPN Gateway │                 │
│  └─────────────┘                    └─────────────┘                 │
│  Pros: Quick to set up, inexpensive                                 │
│  Cons: Internet latency, bandwidth limits                           │
│                                                                     │
│  Option 2: Direct Connect / ExpressRoute / Interconnect             │
│  ┌─────────────┐    Private Line    ┌─────────────┐                 │
│  │ On-Premises │ ────────────────── │ Cloud Edge  │                 │
│  │ Router      │    (dedicated)     │ Location    │                 │
│  └─────────────┘                    └─────────────┘                 │
│  Pros: Consistent latency, high bandwidth                           │
│  Cons: Expensive, weeks to provision                                │
│                                                                     │
│  Security for both:                                                 │
│  • Encrypt traffic even on "private" connections                    │
│  • Firewall at cloud entry point                                    │
│  • Network segmentation (don't expose everything)                   │
│  • Monitor for unusual traffic patterns                             │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

---

## Cloud Access Security Brokers (CASB)

---

### What is CASB?

CASB sits between users and cloud services, providing visibility and control over SaaS applications.

```
┌─────────────────────────────────────────────────────────────────────┐
│ CASB Position in Architecture                                       │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  Users                                                              │
│    │                                                                │
│    │  ┌─────────────────────────────────────────────────────────┐   │
│    └──│                      CASB                               │   │
│       │                                                         │   │
│       │  ┌──────────┐ ┌──────────┐ ┌──────────┐ ┌──────────┐    │   │
│       │  │Visibility│ │ Threat   │ │ DLP      │ │Compliance│    │   │
│       │  │(discover │ │Protection│ │(data     │ │(policy   │    │   │
│       │  │shadow IT)│ │(malware) │ │ control) │ │ enforce) │    │   │
│       │  └──────────┘ └──────────┘ └──────────┘ └──────────┘    │   │
│       │                                                         │   │
│       └─────────────────────────────────────────────────────────┘   │
│              │                    │                     │           │
│              ▼                    ▼                     ▼           │
│       ┌───────────┐        ┌───────────┐         ┌───────────┐      │
│       │ Salesforce│        │ Box       │         │ Office365 │      │
│       └───────────┘        └───────────┘         └───────────┘      │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

---

### CASB Deployment Modes

**1. API Mode (Out-of-band):**
```
┌─────────────────────────────────────────────────────────────────────┐
│ API-Based CASB                                                      │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  User ─────────────────────────▶ SaaS Application                   │
│                                        │                            │
│                                        │ API                        │
│                                        ▼                            │
│                                 ┌─────────────┐                     │
│                                 │    CASB     │                     │
│                                 │ (scanning,  │                     │
│                                 │  analysis)  │                     │
│                                 └─────────────┘                     │
│                                                                     │
│  Pros:                          Cons:                               │
│  • Easy deployment              • Not real-time                     │
│  • No user impact               • Can't block inline                │
│  • Works with any device        • API rate limits                   │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

**2. Proxy Mode (Inline):**
```
┌─────────────────────────────────────────────────────────────────────┐
│ Proxy-Based CASB                                                    │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  Forward Proxy:                 Reverse Proxy:                      │
│                                                                     │
│  User → CASB → SaaS             User → CASB → SaaS                  │
│  (agent on device)              (SAML redirect)                     │
│                                                                     │
│  Pros:                          Pros:                               │
│  • Real-time blocking           • Agentless                         │
│  • Full visibility              • Works for SAML apps               │
│  • Works with any app           • Managed devices not required      │
│                                                                     │
│  Cons:                          Cons:                               │
│  • Requires agent               • Only works with SAML-enabled apps │
│  • Latency impact               • Complex setup                     │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

---

### CASB Use Cases

**1. Shadow IT Discovery:**
```
CASB Analysis Report:
─────────────────────────────────────────────
Cloud Apps Discovered: 1,847
  ├── Sanctioned: 45
  ├── Unsanctioned: 1,802
  └── High Risk: 234

Top Unsanctioned Apps:
1. Dropbox Personal (3,421 users) - Risk: Medium
2. WeTransfer (892 users) - Risk: High
3. Grammarly (2,104 users) - Risk: Low
4. ChatGPT (1,567 users) - Risk: High (data leakage)

Action: Block high-risk, educate on alternatives
```

**2. Data Loss Prevention:**
```yaml
CASB DLP Policy: Prevent PII Sharing

Triggers:
  - File upload to non-sanctioned apps
  - Sharing with external users
  - Download to unmanaged device

Detection:
  - Credit card numbers
  - Social Security numbers
  - Patient health information

Actions:
  - Block and notify user
  - Alert security team
  - Log for compliance
```

**3. Threat Protection:**
```
Threats Detected This Month:
────────────────────────────────────────
• Compromised account logins: 23
  └── Impossible travel detected

• Malware in cloud storage: 7
  └── Files quarantined

• Ransomware behavior: 2
  └── Mass file encryption blocked

• OAuth app risks: 15
  └── Excessive permissions requested
```

---

## Cloud Security Architecture Patterns

---

### Landing Zone Design

A **landing zone** is a pre-configured, secure environment for workloads.

```
┌─────────────────────────────────────────────────────────────────────┐
│ AWS Landing Zone Architecture                                       │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  ┌─────────────────────────────────────────────────────────────┐    │
│  │ Management Account (Root)                                   │    │
│  │ • AWS Organizations                                         │    │
│  │ • Service Control Policies (SCPs)                           │    │
│  │ • Consolidated billing                                      │    │
│  └─────────────────────────────────────────────────────────────┘    │
│                              │                                      │
│       ┌──────────────────────┼──────────────────────┐               │
│       ▼                      ▼                      ▼               │
│  ┌─────────────┐       ┌─────────────┐       ┌─────────────┐        │
│  │ Security OU │       │ Workload OU │       │ Sandbox OU  │        │
│  │             │       │             │       │             │        │
│  │ ┌─────────┐ │       │ ┌─────────┐ │       │ ┌─────────┐ │        │
│  │ │Log      │ │       │ │Prod     │ │       │ │Dev Team │ │        │
│  │ │Archive  │ │       │ │Account  │ │       │ │Accounts │ │        │
│  │ └─────────┘ │       │ └─────────┘ │       │ └─────────┘ │        │
│  │ ┌─────────┐ │       │ ┌─────────┐ │       │             │        │
│  │ │Security │ │       │ │Staging  │ │       │             │        │
│  │ │Tooling  │ │       │ │Account  │ │       │             │        │
│  │ └─────────┘ │       │ └─────────┘ │       │             │        │
│  │ ┌─────────┐ │       │ ┌─────────┐ │       │             │        │
│  │ │Audit    │ │       │ │Dev      │ │       │             │        │
│  │ │Account  │ │       │ │Account  │ │       │             │        │
│  │ └─────────┘ │       │ └─────────┘ │       │             │        │
│  └─────────────┘       └─────────────┘       └─────────────┘        │
│                                                                     │
│  Shared Services:                                                   │
│  ├── Transit Gateway (central networking)                           │
│  ├── Centralized logging (CloudTrail, VPC Flow Logs)                │
│  ├── Security Hub (aggregated findings)                             │
│  └── Identity Center (centralized IAM)                              │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

**AWS Control Tower** automates landing zone setup with:
- Account factory
- Guardrails (preventive and detective)
- Dashboard for compliance status

---

### Account/Subscription Structure

**Principles:**
```
1. Workload Isolation
   └── Separate accounts/subscriptions per workload or environment
   └── Blast radius containment

2. Environment Separation
   └── Dev, Staging, Prod in different accounts
   └── Different security controls per environment

3. Security Account Isolation
   └── Logging account (append-only, no delete)
   └── Security tooling account
   └── Audit account (read-only access to all)

4. Sandbox for Experimentation
   └── Limited budget
   └── No production connectivity
   └── Auto-cleanup policies
```

**Example structure:**

```
Organization
├── Core OU
│   ├── Management Account
│   ├── Logging Account
│   ├── Security Account
│   └── Networking Account (Transit Gateway)
│
├── Workloads OU
│   ├── Application A
│   │   ├── app-a-dev
│   │   ├── app-a-staging
│   │   └── app-a-prod
│   │
│   └── Application B
│       ├── app-b-dev
│       ├── app-b-staging
│       └── app-b-prod
│
└── Sandbox OU
    ├── sandbox-team-1
    └── sandbox-team-2
```

---

### Security Baseline Automation

Every account should have security controls deployed automatically.

**Terraform security baseline module:**

```hcl
# security-baseline/main.tf

module "cloudtrail" {
  source = "./modules/cloudtrail"

  s3_bucket_name     = var.logging_bucket
  enable_log_file_validation = true
  is_multi_region_trail = true
}

module "config" {
  source = "./modules/config"

  config_rules = [
    "s3-bucket-public-read-prohibited",
    "s3-bucket-ssl-requests-only",
    "encrypted-volumes",
    "iam-password-policy",
    "root-account-mfa-enabled"
  ]
}

module "guardduty" {
  source = "./modules/guardduty"

  enable_s3_protection = true
  enable_kubernetes_protection = true
  finding_publishing_frequency = "FIFTEEN_MINUTES"
}

module "security_hub" {
  source = "./modules/security-hub"

  enabled_standards = [
    "aws-foundational-security-best-practices",
    "cis-aws-foundations-benchmark"
  ]
}

module "iam_baseline" {
  source = "./modules/iam"

  require_mfa           = true
  password_max_age      = 90
  password_reuse_prevention = 24
}
```

**Automatic deployment:**
```
New Account Created
        │
        ▼
┌───────────────────┐
│ Account Factory   │
│ (Control Tower or │
│  custom pipeline) │
└────────┬──────────┘
         │
         ▼
┌───────────────────┐
│ Security Baseline │
│ Terraform/CFN     │
│ Automatically     │
│ Applied           │
└───────────────────┘
         │
         ▼
Account ready with:
✓ CloudTrail
✓ Config Rules
✓ GuardDuty
✓ Security Hub
✓ IAM Baseline
```

---

### Immutable Infrastructure

**The principle:** Never modify running infrastructure. Replace it.

```
Traditional (Mutable):              Immutable:
┌─────────────────────┐             ┌─────────────────────┐
│ Server v1           │             │ Server v1           │
│                     │             │                     │
│ Deploy app v1       │             │ App v1 baked in     │
│ Patch OS            │             │ (AMI/Image)         │
│ Update app to v2    │             └─────────────────────┘
│ Install more patches│                      │
│ Hotfix v2.1         │                      │ Replace
│ ???                 │                      ▼
│ Drift happens       │             ┌─────────────────────┐
└─────────────────────┘             │ Server v2           │
                                    │                     │
                                    │ App v2 baked in     │
                                    │ (new AMI/Image)     │
                                    └─────────────────────┘
```

**Security benefits:**
- No configuration drift
- Known state at all times
- Easy rollback (deploy previous image)
- Reduced attack surface (no SSH access needed)
- Simpler forensics (compare to golden image)

**Implementation:**

```hcl
# Packer - Build immutable AMI
source "amazon-ebs" "app" {
  ami_name      = "app-${var.version}-${local.timestamp}"
  instance_type = "t3.medium"
  source_ami    = data.amazon-ami.ubuntu.id
}

build {
  sources = ["source.amazon-ebs.app"]

  provisioner "shell" {
    scripts = [
      "scripts/harden-os.sh",
      "scripts/install-app.sh",
      "scripts/security-scan.sh"
    ]
  }
}
```

---

### GitOps for Security

**Everything in Git, everything automated:**

```
┌─────────────────────────────────────────────────────────────────────┐
│ GitOps Security Workflow                                            │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  Developer                                                          │
│     │                                                               │
│     │ 1. Submit PR                                                  │
│     ▼                                                               │
│  ┌─────────────────┐                                                │
│  │ Git Repository  │                                                │
│  │ (GitHub, GitLab)│                                                │
│  └────────┬────────┘                                                │
│           │                                                         │
│           │ 2. Automated checks                                     │
│           ▼                                                         │
│  ┌─────────────────────────────────────────────────────────────┐    │
│  │ CI Pipeline                                                 │    │
│  │ ├── Terraform fmt/validate                                  │    │
│  │ ├── tfsec (security scanner)                                │    │
│  │ ├── checkov (policy as code)                                │    │
│  │ ├── OPA/Conftest (custom policies)                          │    │
│  │ └── terraform plan (diff review)                            │    │
│  └─────────────────────────────────────────────────────────────┘    │
│           │                                                         │
│           │ 3. Require approvals                                    │
│           ▼                                                         │
│  ┌─────────────────┐                                                │
│  │ Security Review │ ◀── Required for sensitive changes             │
│  │ (human approval)│                                                │
│  └────────┬────────┘                                                │
│           │                                                         │
│           │ 4. Merge & Deploy                                       │
│           ▼                                                         │
│  ┌─────────────────┐                                                │
│  │ CD Pipeline     │                                                │
│  │ terraform apply │                                                │
│  └────────┬────────┘                                                │
│           │                                                         │
│           │ 5. Drift detection                                      │
│           ▼                                                         │
│  ┌─────────────────┐                                                │
│  │ Continuous      │                                                │
│  │ Reconciliation  │ ◀── Alert if actual != desired state           │
│  └─────────────────┘                                                │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

---

## Bringing It All Together

Here's a complete enterprise cloud security architecture:

```
┌──────────────────────────────────────────────────────────────────────────────┐
│ Enterprise Cloud Security Architecture                                       │
├──────────────────────────────────────────────────────────────────────────────┤
│                                                                              │
│  ┌────────────────────────────────────────────────────────────────────────┐  │
│  │ IDENTITY LAYER                                                         │  │
│  │ Central IdP → Federation → AWS/Azure/GCP IAM → JIT Access → MFA        │  │
│  └────────────────────────────────────────────────────────────────────────┘  │
│                                         │                                    │
│  ┌────────────────────────────────────────────────────────────────────────┐  │
│  │ NETWORK LAYER                                                          │  │
│  │ On-Prem ← Direct Connect → Transit GW → VPCs → Micro-segmentation      │  │
│  │                               ↓                                        │  │
│  │                    ┌──────────────────────┐                            │  │
│  │                    │ Centralized Firewall │                            │  │
│  │                    │ (Inspection VPC)     │                            │  │
│  │                    └──────────────────────┘                            │  │
│  └────────────────────────────────────────────────────────────────────────┘  │
│                                         │                                    │
│  ┌────────────────────────────────────────────────────────────────────────┐  │
│  │ WORKLOAD LAYER                                                         │  │
│  │ ┌─────────────┐  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐     │  │
│  │ │ VMs         │  │ Containers  │  │ Serverless  │  │ SaaS        │     │  │
│  │ │ (CWPP)      │  │ (K8s+CWPP)  │  │ (IAM+Code)  │  │ (CASB)      │     │  │
│  │ └─────────────┘  └─────────────┘  └─────────────┘  └─────────────┘     │  │
│  └────────────────────────────────────────────────────────────────────────┘  │
│                                         │                                    │
│  ┌────────────────────────────────────────────────────────────────────────┐  │
│  │ DATA LAYER                                                             │  │
│  │ Classification → Encryption (CMK) → DLP → Access Logging → Backup      │  │
│  └────────────────────────────────────────────────────────────────────────┘  │
│                                         │                                    │
│  ┌────────────────────────────────────────────────────────────────────────┐  │
│  │ SECURITY OPERATIONS                                                    │  │
│  │ ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌──────────┐   │  │
│  │ │ CSPM     │  │ SIEM     │  │ SOAR     │  │ Vuln Mgmt│  │ Threat   │   │  │
│  │ │ (posture)│  │ (detect) │  │(respond) │  │ (scan)   │  │ Intel    │   │  │
│  │ └──────────┘  └──────────┘  └──────────┘  └──────────┘  └──────────┘   │  │
│  └────────────────────────────────────────────────────────────────────────┘  │
│                                         │                                    │
│  ┌────────────────────────────────────────────────────────────────────────┐  │
│  │ GOVERNANCE                                                             │  │
│  │ Policy as Code → Guardrails → Compliance Reporting → Audit Trail       │  │
│  └────────────────────────────────────────────────────────────────────────┘  │
│                                                                              │
└──────────────────────────────────────────────────────────────────────────────┘
```

---

## Key Concepts to Remember

1. **Multi-cloud requires abstraction** - Policy-as-code, centralized identity, unified logging
2. **CASB extends security to SaaS** - Visibility, DLP, threat protection for cloud apps
3. **Landing zones provide guardrails** - Secure by default, hard to misconfigure
4. **Account separation limits blast radius** - Workload isolation, environment separation
5. **Automate security baselines** - Every account gets CloudTrail, Config, GuardDuty
6. **Immutable infrastructure prevents drift** - Replace, don't modify
7. **GitOps enables security review** - All changes through PR, automated scanning

---

## Practice Questions

**Q1:** Your organization is adopting a multi-cloud strategy (AWS primary, Azure for specific workloads, GCP for ML). The CISO is concerned about maintaining consistent security. What architecture would you propose?

<details>
<summary>View Answer</summary>

**Proposed Architecture:**

**1. Centralized Identity:**
```
┌─────────────────────────────────────────────────────────────┐
│ Okta (or Azure AD) as Central IdP                           │
├─────────────────────────────────────────────────────────────┤
│ • All users authenticate through Okta                       │
│ • SAML federation to AWS IAM Identity Center                │
│ • SAML federation to Azure AD (B2B)                         │
│ • SAML federation to GCP Cloud Identity                     │
│ • SCIM provisioning to sync users/groups                    │
│ • Consistent MFA policy (hardware keys for admins)          │
│ • Centralized access reviews and certification              │
└─────────────────────────────────────────────────────────────┘
```

**2. Unified Policy Enforcement:**
```
┌─────────────────────────────────────────────────────────────┐
│ Policy-as-Code Repository (Git)                             │
├─────────────────────────────────────────────────────────────┤
│ policies/                                                   │
│ ├── opa/                    (cloud-agnostic Rego policies)  │
│ ├── terraform/              (IaC security - Checkov)        │
│ └── cloud-native/                                           │
│     ├── aws-scp/            (Service Control Policies)      │
│     ├── azure-policy/       (Azure Policy definitions)      │
│     └── gcp-org-policy/     (Organization Policies)         │
└─────────────────────────────────────────────────────────────┘
```

**3. Centralized Logging & Security Operations:**
```
AWS CloudTrail ─────┐
Azure Activity ─────┼───▶ Centralized SIEM (Splunk/Elastic)
GCP Cloud Audit ────┘              │
                                   ▼
                    ┌─────────────────────────────────┐
                    │ Unified Detection Rules         │
                    │ • Normalized log format         │
                    │ • Cross-cloud correlation       │
                    │ • Single pane of glass          │
                    └─────────────────────────────────┘
```

**4. Multi-Cloud CSPM:**
```
• Deploy Wiz, Prisma Cloud, or similar
• Single dashboard for all three clouds
• Consistent compliance benchmarks
• Unified risk prioritization
```

**5. Network Architecture:**
```
                    ┌─────────────────────────────────┐
On-Premises ────────│ Hub Network (AWS Transit GW)    │
                    │ Centralized firewall inspection │
                    └────────────┬────────────────────┘
                                 │
        ┌────────────────────────┼────────────────────────┐
        ▼                        ▼                        ▼
   AWS VPCs              Azure VNet (VPN)           GCP VPC
   (native)              (site-to-site)            (VPN/Interconnect)
```

**6. Governance:**
- Cloud Center of Excellence (CCoE) with security representation
- Standard landing zones for each cloud
- Approved patterns catalog
- Security review gates in deployment pipelines

</details>

**Q2:** A business unit wants to use a new SaaS application for project management. They've already started a pilot. How do you handle this from a security perspective, and what role does CASB play?

<details>
<summary>View Answer</summary>

**Shadow IT Discovery & Response Process:**

**1. Discovery (via CASB):**
```
CASB Alert: New unsanctioned application detected
────────────────────────────────────────────────
Application: ProjectFlow (project management SaaS)
Users: 45 employees
Data uploaded: 234 files
Department: Marketing
Risk Score: Medium
Concerns:
  • Not on approved vendor list
  • Data residency unknown
  • SSO not configured (password-based)
  • OAuth token with broad permissions granted
```

**2. Assessment (not blocking immediately):**
```
Step 1: Reach out to business unit
  • "We noticed you're piloting ProjectFlow"
  • "Let's work together to evaluate it properly"
  • Don't be adversarial—they have a legitimate need

Step 2: Security questionnaire to vendor
  • SOC 2 Type II report?
  • Data encryption (at rest, in transit)?
  • Data residency options?
  • SSO/SAML support?
  • Security certifications?
  • Incident response procedures?

Step 3: Technical assessment
  • CASB API connection for monitoring
  • DLP policy testing
  • SSO integration feasibility
  • Data export/portability
```

**3. Decision Framework:**
```
┌───────────────────────────────────────────────────────────────┐
│ If vendor passes assessment:                                  │
│ ├── Add to sanctioned app list                                │
│ ├── Configure SSO integration                                 │
│ ├── Apply CASB policies (DLP, access control)                 │
│ ├── Document in cloud services catalog                        │
│ └── Onboard remaining users properly                          │
│                                                               │
│ If vendor fails assessment:                                   │
│ ├── Explain risks to business unit                            │
│ ├── Propose alternatives (approved tools)                     │
│ ├── Assist with data migration                                │
│ ├── Set deadline for transition                               │
│ └── Block via CASB after transition period                    │
└───────────────────────────────────────────────────────────────┘
```

**4. CASB Controls for Approved Apps:**
```yaml
CASB Policy for ProjectFlow:
  Access Control:
    - Require SSO authentication
    - Block personal account usage
    - Session timeout: 8 hours

  DLP:
    - Block upload of PII, PCI data
    - Scan shared links for sensitive content
    - Alert on mass downloads

  Threat Protection:
    - Scan uploaded files for malware
    - Detect anomalous behavior (mass delete)
    - Alert on impossible travel

  Visibility:
    - Log all file sharing
    - Monitor OAuth app permissions
    - Track external collaborators
```

**5. Ongoing Governance:**
- Quarterly access reviews
- Annual vendor security reassessment
- CASB risk score monitoring
- User education on approved tools

</details>

**Q3:** Your organization allows employees to access Microsoft 365 from both managed corporate laptops and unmanaged personal devices. You need a solution that allows full access from corporate devices but restricts personal devices to "web-only" access (blocking downloads) in real-time. Which CASB deployment mode is required?

<details>
<summary>View Answer</summary>

**Answer: Reverse Proxy Mode**

**Reasoning:**
*   **Real-time Control:** API mode is out-of-band and cannot block downloads inline in real-time.
*   **Unmanaged Devices:** Forward Proxy requires an agent on the endpoint, which you cannot force on unmanaged/personal devices.
*   **Solution:** Reverse Proxy sits in the authentication path (via SAML). When a user logs in, the IdP redirects them through the CASB. The CASB can then inspect the device posture (managed vs unmanaged) and dynamically rewrite the session to block downloads if the device is unmanaged.
</details>

**Q4:** A fintech company needs to connect their on-premises mainframe to their AWS VPC to process transaction data. The connection requires consistent low latency, high bandwidth (10 Gbps+), and must not traverse the public internet. Which connectivity option should they choose?

<details>
<summary>View Answer</summary>

**Answer: AWS Direct Connect (or Azure ExpressRoute / GCP Interconnect)**

**Reasoning:**
*   **Performance:** VPNs over the internet are subject to variable latency and jitter ("internet weather"). Dedicated connections provide consistent, SLA-backed performance.
*   **Bandwidth:** VPN tunnels typically top out at ~1.25 Gbps per tunnel. Direct Connect supports 10 Gbps, 100 Gbps, and aggregated links.
*   **Security:** The requirement "must not traverse the public internet" rules out VPNs (even though VPNs are encrypted, they ride over the public internet). Direct Connect uses private fiber circuits.
</details>

**Q5:** You have 50 VPCs in AWS and need to inspect all east-west traffic (VPC-to-VPC) and north-south traffic (VPC-to-Internet) using a central fleet of Next-Gen Firewalls. Which network architecture pattern supports this with the least management overhead?

<details>
<summary>View Answer</summary>

**Answer: Hub and Spoke with Transit Gateway (Inspection VPC pattern)**

**Reasoning:**
*   **Hub and Spoke:** Connecting 50 VPCs with VPC Peering (Full Mesh) would require $N(N-1)/2$ connections (1,225 peers), which is unmanageable. Transit Gateway acts as a central hub.
*   **Centralized Inspection:** You can create a dedicated "Security Protocol" or "Inspection VPC" attached to the Transit Gateway.
*   **Routing:** Route tables in the TGW force traffic from Spoke VPCs to the Inspection VPC firewall fleet before forwarding it to the destination (another VPC or Internet).
*   **Scalability:** This allows you to manage one firewall cluster for the entire organization rather than deploying firewalls in every VPC.
</details>

**Q6:** You're designing the landing zone for a regulated financial services company moving to AWS. What guardrails would you implement, and how would you enforce them?

<details>
<summary>View Answer</summary>

**Financial Services Landing Zone Design:**

**1. Organizational Structure:**
```
Organization Root
│
├── Core OU (Deny SCPs - no workloads here)
│   ├── Management Account
│   │   └── Organizations, Billing, IAM Identity Center
│   ├── Log Archive Account
│   │   └── S3 (immutable), CloudTrail, VPC Flow Logs
│   ├── Security Tooling Account
│   │   └── Security Hub, GuardDuty (delegated admin)
│   └── Network Account
│       └── Transit Gateway, Direct Connect, DNS
│
├── Workloads OU
│   ├── Production OU (strictest controls)
│   │   ├── Payment Processing Account
│   │   ├── Core Banking Account
│   │   └── Customer Data Account
│   │
│   ├── Non-Production OU
│   │   ├── Staging Accounts
│   │   └── Development Accounts
│   │
│   └── Data OU
│       ├── Data Lake Account
│       └── Analytics Account
│
└── Sandbox OU (isolated, auto-cleanup)
    └── Developer Sandbox Accounts
```

**2. Service Control Policies (Preventive Guardrails):**

```json
// SCP: Deny non-approved regions
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Deny",
      "Action": "*",
      "Resource": "*",
      "Condition": {
        "StringNotEquals": {
          "aws:RequestedRegion": ["us-east-1", "us-west-2"]
        }
      }
    }
  ]
}

// SCP: Require encryption
{
  "Effect": "Deny",
  "Action": [
    "s3:CreateBucket"
  ],
  "Resource": "*",
  "Condition": {
    "StringNotEquals": {
      "s3:x-amz-server-side-encryption": "aws:kms"
    }
  }
}

// SCP: Deny public S3
{
  "Effect": "Deny",
  "Action": [
    "s3:PutBucketPublicAccessBlock",
    "s3:DeletePublicAccessBlock"
  ],
  "Resource": "*"
}

// SCP: Require IMDSv2
{
  "Effect": "Deny",
  "Action": "ec2:RunInstances",
  "Resource": "*",
  "Condition": {
    "StringNotEquals": {
      "ec2:MetadataHttpTokens": "required"
    }
  }
}
```

**3. Detective Guardrails (AWS Config Rules):**

```yaml
Required Config Rules:
  # Encryption
  - s3-bucket-server-side-encryption-enabled
  - rds-storage-encrypted
  - encrypted-volumes
  - dynamodb-table-encrypted-kms

  # Network Security
  - vpc-flow-logs-enabled
  - vpc-sg-open-only-to-authorized-ports
  - no-unrestricted-route-to-igw

  # Identity & Access
  - iam-password-policy
  - root-account-mfa-enabled
  - iam-user-mfa-enabled
  - iam-no-inline-policy-check

  # Logging & Monitoring
  - cloudtrail-enabled
  - cloud-trail-encryption-enabled
  - guardduty-enabled-centralized

  # Data Protection
  - s3-bucket-public-read-prohibited
  - s3-bucket-public-write-prohibited
  - rds-instance-public-access-check
```

**4. Mandatory Security Baseline (Terraform):**

```hcl
# Deployed to every account automatically
module "security_baseline" {
  source = "git::https://github.com/company/tf-security-baseline"

  enable_cloudtrail           = true
  cloudtrail_s3_bucket        = var.central_logging_bucket
  enable_config               = true
  enable_guardduty            = true
  guardduty_master_account_id = var.security_account_id
  enable_security_hub         = true
  enable_access_analyzer      = true
  enable_macie                = true  # For PII detection

  # PCI DSS specific
  enable_pci_config_rules     = true
}
```

**5. Network Guardrails:**

```
• All egress through centralized firewall
• No direct internet access from production
• Private endpoints for all AWS services
• TLS 1.2+ required for all connections
• VPC Flow Logs to central log archive
```

**6. Compliance Automation:**

```
Daily:
  • Security Hub compliance score
  • Config rule compliance status
  • GuardDuty findings review

Weekly:
  • IAM Access Analyzer findings
  • Macie sensitive data findings

Monthly:
  • SOC 2 evidence collection
  • PCI DSS control validation
  • Access certification reviews

Quarterly:
  • Penetration testing
  • External audit prep
```

</details>

**Q7:** Explain how you would implement GitOps for infrastructure security, including the security scanning pipeline and approval workflows.

<details>
<summary>View Answer</summary>

**GitOps Security Implementation:**

**1. Repository Structure:**
```
infrastructure-repo/
├── .github/
│   └── workflows/
│       ├── pr-checks.yml          # PR validation
│       ├── plan.yml               # Terraform plan
│       └── apply.yml              # Terraform apply (after merge)
│
├── modules/                       # Reusable, security-reviewed modules
│   ├── vpc/
│   ├── eks/
│   ├── rds/
│   └── s3/
│
├── environments/
│   ├── dev/
│   │   ├── main.tf
│   │   └── terraform.tfvars
│   ├── staging/
│   └── production/
│
├── policies/                      # OPA/Rego policies
│   ├── deny-public-s3.rego
│   ├── require-encryption.rego
│   └── require-tags.rego
│
└── CODEOWNERS                     # Required reviewers
```

**2. CODEOWNERS for Security Reviews:**
```
# CODEOWNERS
# Security team must approve changes to IAM, networking, encryption

*.tf                    @platform-team
*/iam.tf                @security-team @platform-team
*/networking.tf         @security-team @network-team
*/kms.tf                @security-team
environments/production/ @security-team @platform-lead

# Any changes to policies require security approval
policies/               @security-team
```

**3. PR Validation Pipeline:**
```yaml
# .github/workflows/pr-checks.yml
name: Security Checks

on:
  pull_request:
    paths:
      - '**.tf'
      - 'policies/**'

jobs:
  security-scan:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Terraform Format Check
        run: terraform fmt -check -recursive

      - name: Terraform Validate
        run: |
          terraform init -backend=false
          terraform validate

      - name: tfsec Security Scan
        uses: aquasecurity/tfsec-action@v1.0.0
        with:
          soft_fail: false

      - name: Checkov Scan
        uses: bridgecrewio/checkov-action@v12
        with:
          directory: .
          framework: terraform
          output_format: sarif
          soft_fail: false

      - name: OPA Policy Check
        run: |
          conftest test . --policy policies/ --all-namespaces

      - name: Terraform Plan
        run: |
          terraform plan -out=tfplan
          terraform show -json tfplan > tfplan.json

      - name: Scan Plan for Sensitive Changes
        run: |
          # Check if plan modifies IAM, KMS, or security groups
          python scripts/classify-changes.py tfplan.json

      - name: Post Plan to PR
        uses: actions/github-script@v6
        with:
          script: |
            // Post terraform plan output as PR comment
            // Highlight security-relevant changes
```

**4. Policy Examples (OPA/Rego):**
```rego
# policies/deny-public-s3.rego
package terraform.s3

deny[msg] {
    resource := input.resource_changes[_]
    resource.type == "aws_s3_bucket_public_access_block"
    resource.change.after.block_public_acls == false
    msg := sprintf("S3 bucket %s must block public ACLs", [resource.address])
}

# policies/require-encryption.rego
package terraform.encryption

deny[msg] {
    resource := input.resource_changes[_]
    resource.type == "aws_ebs_volume"
    resource.change.after.encrypted == false
    msg := sprintf("EBS volume %s must be encrypted", [resource.address])
}

deny[msg] {
    resource := input.resource_changes[_]
    resource.type == "aws_db_instance"
    resource.change.after.storage_encrypted == false
    msg := sprintf("RDS instance %s must have storage encryption", [resource.address])
}
```

**5. Approval Workflow:**
```yaml
# Required checks before merge
branch_protection:
  required_status_checks:
    - terraform-fmt
    - terraform-validate
    - tfsec
    - checkov
    - opa-policies
    - terraform-plan

  required_reviews:
    - count: 2
    - dismiss_stale: true
    - require_code_owner_review: true

  # For production changes
  required_reviewers:
    - security-team (if IAM/KMS/network changes detected)
```

**6. Apply Pipeline (Post-Merge):**
```yaml
# .github/workflows/apply.yml
name: Terraform Apply

on:
  push:
    branches: [main]
    paths:
      - 'environments/**'

jobs:
  apply:
    runs-on: ubuntu-latest
    environment: production  # Requires manual approval
    steps:
      - uses: actions/checkout@v4

      - name: Terraform Apply
        run: terraform apply -auto-approve

      - name: Notify Security Team
        uses: slackapi/slack-github-action@v1
        with:
          payload: |
            {
              "text": "Infrastructure deployed to production",
              "attachments": [
                {
                  "color": "#36a64f",
                  "fields": [
                    {"title": "Commit", "value": "${{ github.sha }}"},
                    {"title": "Author", "value": "${{ github.actor }}"}
                  ]
                }
              ]
            }
```

**7. Drift Detection:**
```yaml
# .github/workflows/drift-detection.yml
name: Drift Detection

on:
  schedule:
    - cron: '0 */6 * * *'  # Every 6 hours

jobs:
  detect-drift:
    runs-on: ubuntu-latest
    steps:
      - name: Terraform Plan (detect drift)
        run: |
          terraform plan -detailed-exitcode -out=drift.tfplan
        continue-on-error: true

      - name: Alert on Drift
        if: steps.plan.outcome == 'failure'
        run: |
          # Parse drift and alert
          # Create issue or Slack alert
```

</details>

---

## Module 5 Summary

Congratulations! You've completed the Cloud Security module. Here's what you've learned:

```
┌─────────────────────────────────────────────────────────────────────────────┐
│ Module 5: Cloud Security - Complete                                         │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│ Lesson 1: Shared Responsibility & Identity                                  │
│ ├── Cloud provider vs customer responsibilities                             │
│ ├── IaaS/PaaS/SaaS security models                                          │
│ ├── Cloud IAM architecture                                                  │
│ ├── Workload identity & service accounts                                    │
│ └── Federation & JIT access                                                 │
│                                                                             │
│ Lesson 2: Cloud Network Security                                            │
│ ├── VPC design patterns                                                     │
│ ├── Security Groups & NACLs                                                 │
│ ├── Private endpoints                                                       │
│ ├── Transit Gateway & hub-spoke                                             │
│ └── DDoS protection                                                         │
│                                                                             │
│ Lesson 3: CSPM & CWPP                                                       │
│ ├── Misconfiguration detection                                              │
│ ├── Compliance benchmarks                                                   │
│ ├── Policy-as-code                                                          │
│ ├── Runtime workload protection                                             │
│ └── Vulnerability management                                                │
│                                                                             │
│ Lesson 4: Container, Kubernetes & Serverless                                │
│ ├── Container image security                                                │
│ ├── Kubernetes RBAC & Pod Security                                          │
│ ├── Network policies                                                        │
│ ├── Secrets management                                                      │
│ └── Serverless security model                                               │
│                                                                             │
│ Lesson 5: Cloud Data Protection                                             │
│ ├── Data classification                                                     │
│ ├── Encryption at rest, in transit, in use                                  │
│ ├── Key management (KMS, HSM)                                               │
│ ├── Storage security (S3, databases)                                        │
│ └── Data loss prevention                                                    │
│                                                                             │
│ Lesson 6: Multi-Cloud, CASB & Architecture Patterns                         │
│ ├── Multi-cloud security strategy                                           │
│ ├── CASB for SaaS security                                                  │
│ ├── Landing zone design                                                     │
│ ├── Security baseline automation                                            │
│ └── GitOps for infrastructure                                               │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

**You're now prepared to:**
- Architect secure cloud environments from the ground up
- Implement defense-in-depth across all cloud layers
- Enforce consistent security policies across multi-cloud
- Automate security controls and compliance
- Protect data throughout its lifecycle
- Secure modern workloads (containers, Kubernetes, serverless)
