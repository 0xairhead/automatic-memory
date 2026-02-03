# Module 5, Lesson 5: Cloud Data Protection

## Table of Contents
- [Media Resources](#media-resources)
- [Data Classification & Governance](#data-classification--governance)
  - [Classification Framework](#classification-framework)
  - [Data Discovery & Tagging](#data-discovery--tagging)
- [Encryption Strategies](#encryption-strategies)
  - [Encryption at Rest](#encryption-at-rest)
  - [Encryption in Transit](#encryption-in-transit)
  - [Encryption in Use](#encryption-in-use)
- [Key Management Architecture](#key-management-architecture)
  - [Cloud KMS Services](#cloud-kms-services)
  - [Customer-Managed vs Provider-Managed Keys](#customer-managed-vs-provider-managed-keys)
  - [Hardware Security Modules (HSM)](#hardware-security-modules-hsm)
  - [Key Rotation & Lifecycle](#key-rotation--lifecycle)
- [Cloud Storage Security](#cloud-storage-security)
  - [Object Storage (S3, Blob, GCS)](#object-storage-s3-blob-gcs)
  - [Database Security](#database-security)
- [Data Loss Prevention (DLP)](#data-loss-prevention-dlp)
- [Key Concepts to Remember](#key-concepts-to-remember)
- [Practice Questions](#practice-questions)
- [Next Up](#next-up)

---

Data is the ultimate target. Attackers don't want your compute or network—they want what's stored on them.

---

## Media Resources

**Visual Guide:**

![Cloud Data Protection Architecture](./assets/05-cloud-data-protection.png)

**Audio Lecture:**

🎧 <audio controls src="./assets/05-cloud-data-protection-audio.m4a"></audio>

---

## Data Classification & Governance

Before you can protect data, you must understand what you have and how sensitive it is.

---

### Classification Framework

```
┌─────────────────────────────────────────────────────────────────────┐
│ Data Classification Levels                                          │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  ┌───────────────────────────────────────────────────────────────┐  │
│  │ PUBLIC                                                        │  │
│  │ • Marketing materials, public docs                            │  │
│  │ • No business impact if disclosed                             │  │
│  │ • Controls: Basic access management                           │  │
│  └───────────────────────────────────────────────────────────────┘  │
│                             ▲                                       │
│  ┌───────────────────────────────────────────────────────────────┐  │
│  │ INTERNAL                                                      │  │
│  │ • Internal communications, non-sensitive business data        │  │
│  │ • Minor business impact if disclosed                          │  │
│  │ • Controls: Authentication, basic encryption                  │  │
│  └───────────────────────────────────────────────────────────────┘  │
│                             ▲                                       │
│  ┌───────────────────────────────────────────────────────────────┐  │
│  │ CONFIDENTIAL                                                  │  │
│  │ • Customer data, financial records, IP                        │  │
│  │ • Significant business impact if disclosed                    │  │
│  │ • Controls: Encryption, access logging, DLP                   │  │
│  └───────────────────────────────────────────────────────────────┘  │
│                             ▲                                       │
│  ┌───────────────────────────────────────────────────────────────┐  │
│  │ RESTRICTED / HIGHLY CONFIDENTIAL                              │  │
│  │ • PII, PHI, PCI data, trade secrets                           │  │
│  │ • Severe/regulatory impact if disclosed                       │  │
│  │ • Controls: Strong encryption, MFA, audit, DLP, RBAC          │  │
│  └───────────────────────────────────────────────────────────────┘  │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

**Mapping to regulations:**

| Classification | Regulations | Examples |
|---------------|-------------|----------|
| Restricted | HIPAA, PCI DSS, GDPR | SSN, Credit cards, Health records |
| Confidential | SOX, GDPR | Financial data, Customer records |
| Internal | General policies | Project plans, Meeting notes |
| Public | None specific | Press releases, Product docs |

---

### Data Discovery & Tagging

You can't protect what you don't know exists.

**Automated discovery tools:**

```
┌─────────────────────────────────────────────────────────────────────┐
│ Data Discovery Pipeline                                             │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  ┌──────────────┐    ┌──────────────┐    ┌──────────────┐           │
│  │ Scan Data    │───▶│ Classify     │───▶│ Tag & Label  │           │
│  │ Stores       │    │ Content      │    │              │           │
│  │              │    │              │    │              │           │
│  │ • S3 buckets │    │ • Pattern    │    │ • AWS Tags   │           │
│  │ • Databases  │    │   matching   │    │ • Azure Tags │           │
│  │ • File shares│    │ • ML models  │    │ • GCP Labels │           │
│  │ • SaaS apps  │    │ • Fingerprint│    │              │           │
│  └──────────────┘    └──────────────┘    └──────────────┘           │
│                                                                     │
│  AWS: Macie                                                         │
│  Azure: Purview                                                     │
│  GCP: Cloud DLP / Data Catalog                                      │
│  Third-party: BigID, Varonis, Nightfall                             │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

### Cloud-Native Discovery Tools

#### AWS Macie
*   **Best for:** AWS-centric organizations using S3 heavily.
*   **How it works:** Uses machine learning to automatically discover, classify, and protect sensitive data in AWS.
*   **Key Features:**
    *   Continuously monitors S3 buckets for PII, PHI, and financial data.
    *   Generates findings in Security Hub and EventBridge for automation (e.g., auto-tagging buckets).
    *   Provides a dashboard of your data security posture across accounts.

#### Azure Purview
*   **Best for:** Hybrid and multi-cloud environments requiring unified governance.
*   **How it works:** A unified data governance solution that maps data across your on-prem, multi-cloud, and SaaS estate.
*   **Key Features:**
    *   **Data Map:** Visual graph of data assets and lineage.
    *   **Data Catalog:** Searchable inventory for business and technical users.
    *   **Data Estate Insights:** View of sensitive data across the entire organization (SQL, Blob, AWS S3, etc.).

#### GCP Cloud DLP / Data Catalog
*   **Best for:** Google Cloud workloads and BigQuery analytics.
*   **How it works:** Fully managed service to inspect, classify, and de-identify sensitive data.
*   **Key Features:**
    *   **De-identification:** Redact, mask, or tokenize data *before* it's stored or processed.
    *   **Streaming API:** Inspect data in real-time streams.
    *   **Risk Analysis:** Calculate k-anonymity and l-diversity for privacy compliance.

### Third-Party Solutions

*   **BigID:** Focuses on deep discovery and privacy automation. Excellent for finding "dark data" across fragmented environments.
*   **Varonis:** Strong on-prem roots extended to cloud. Excellent for permission visualization and user behavior analytics alongside data classification.
*   **Nightfall:** API-driven cloud-native DLP. Great for scanning SaaS apps (Slack, Jira, GitHub) where cloud provider tools often lack reach.

**Sensitive data patterns detected:**

```
Pattern Type           Examples Found
─────────────────────────────────────────────────
Credit Card            4532-xxxx-xxxx-1234
SSN                    xxx-xx-4567
Email                  user@company.com
AWS Access Key         AKIAIOSFODNN7EXAMPLE
API Key                api_key=sk_live_abc123
Private Key            -----BEGIN RSA PRIVATE KEY-----
Health Record          Patient ID, Diagnosis codes
```

---

## Encryption Strategies

---

### Encryption at Rest

Data sitting in storage must be encrypted.

```
┌─────────────────────────────────────────────────────────────────────┐
│ Encryption at Rest Models                                           │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  Server-Side Encryption (SSE):                                      │
│  ┌───────────────────────────────────────────────────────────────┐  │
│  │ Client → Cloud Service → Encrypt → Store                      │  │
│  │                ▲                                              │  │
│  │                │                                              │  │
│  │         Cloud manages encryption                              │  │
│  │         (SSE-S3, SSE-KMS, Azure SSE)                          │  │
│  └───────────────────────────────────────────────────────────────┘  │
│                                                                     │
│  Client-Side Encryption (CSE):                                      │
│  ┌───────────────────────────────────────────────────────────────┐  │
│  │ Client → Encrypt → Cloud Service → Store (encrypted)          │  │
│  │    ▲                                                          │  │
│  │    │                                                          │  │
│  │  Client manages encryption                                    │  │
│  │  (Cloud never sees plaintext)                                 │  │
│  └───────────────────────────────────────────────────────────────┘  │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

**AWS S3 encryption options:**

| Option | Key Management | Use Case |
|--------|----------------|----------|
| SSE-S3 | AWS manages | Default, simple |
| SSE-KMS | Customer-managed in KMS | Audit trail, key control |
| SSE-C | Customer provides key | You control keys completely |
| Client-side | Customer encrypts | Maximum control, cloud sees nothing |

**Enforce encryption by default:**

```json
// S3 bucket policy - deny unencrypted uploads
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "DenyUnencryptedUploads",
      "Effect": "Deny",
      "Principal": "*",
      "Action": "s3:PutObject",
      "Resource": "arn:aws:s3:::my-bucket/*",
      "Condition": {
        "StringNotEquals": {
          "s3:x-amz-server-side-encryption": "aws:kms"
        }
      }
    }
  ]
}
```

---

### Encryption in Transit

All data moving over networks should be encrypted with TLS.

```
┌─────────────────────────────────────────────────────────────────────┐
│ Encryption in Transit Points                                        │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  Internet → Load Balancer → Application → Database → Backup         │
│      │            │              │            │          │          │
│      ▼            ▼              ▼            ▼          ▼          │
│    TLS 1.3      TLS 1.2+      mTLS      TLS + Auth   TLS to         │
│   (HTTPS)     (ALB term)   (pod-to-pod) (RDS SSL)   storage         │
│                                                                     │
│  Every hop should be encrypted!                                     │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

**TLS configuration best practices:**

```
✓ TLS 1.2 minimum (TLS 1.3 preferred)
✓ Strong cipher suites only
✓ HSTS headers (force HTTPS)
✓ Certificate validation enabled
✓ Certificate pinning for mobile apps

✗ SSL 2.0, SSL 3.0, TLS 1.0, TLS 1.1 (deprecated)
✗ Weak ciphers (RC4, DES, export ciphers)
✗ Self-signed certs in production
```

**Internal traffic encryption (Zero Trust):**

```yaml
# Istio - enforce mTLS between all pods
apiVersion: security.istio.io/v1beta1
kind: PeerAuthentication
metadata:
  name: default
  namespace: istio-system
spec:
  mtls:
    mode: STRICT  # All traffic must be mTLS
```

---

### Encryption in Use

The final frontier: protecting data while it's being processed.

```
Traditional:                     Confidential Computing:
┌──────────────────────────┐     ┌──────────────────────────┐
│ Memory (plaintext)       │     │ Encrypted Enclave        │
│                          │     │ ┌──────────────────────┐ │
│ ┌──────────────────────┐ │     │ │ Memory (encrypted)   │ │
│ │ Sensitive Data       │ │     │ │                      │ │
│ │ (visible to host,    │ │     │ │ CPU decrypts only    │ │
│ │  hypervisor, admins) │ │     │ │ inside enclave       │ │
│ └──────────────────────┘ │     │ │                      │ │
│                          │     │ │ Host cannot see data │ │
└──────────────────────────┘     │ └──────────────────────┘ │
                                 └──────────────────────────┘
```

**Technologies & Hardware Roots of Trust:**

#### 1. Intel SGX (Software Guard Extensions)
*   **What it is:** Application-layer isolation. Creates "enclaves" – protected regions of memory that even the OS or hypervisor cannot read.
*   **Security Model:** Removes the OS and Hypervisor from the Trusted Computing Base (TCB).
*   **Trade-off:** Requires application refactoring (you must rewrite apps to use the SGX SDK) and has memory size limits.

#### 2. AMD SEV (Secure Encrypted Virtualization)
*   **What it is:** VM-layer isolation. Encrypts the entire virtual machine's memory with a key managed by the AMD Secure Processor.
*   **Security Model:** Protects the VM from the Hypervisor.
*   **Trade-off:** Easier "lift and shift" (no code changes needed), but the Guest OS is still inside the TCB.

**Cloud Provider Implementations:**

#### AWS Nitro Enclaves
*   **Mechanism:** Uses the Nitro Hypervisor to carve out isolated compute environments from EC2 instances.
*   **Features:** No persistent storage, no interactive access (SSH), and only secure local channel communication.
*   **Best For:** Processing highly sensitive data / cryptographic operations where you want to prove that no admin could possibly SSH in and dump memory.

#### Azure Confidential Computing
*   **Mechanism:** extensive support for both Intel SGX (DCsv3-series) and AMD SEV-SNP.
*   **Beat For:** "Confidential Containers" on AKS – running Kubernetes pods in secure enclaves without code changes.

#### GCP Confidential VMs
*   **Mechanism:** Built on AMD SEV.
*   **Features:** "Click to enable" simplicity. Data stays encrypted in memory with no performance degradation.
*   **Best For:** Lift-and-shift of legacy applications that process sensitive data, needing immediate compliance upgrade.

**Use cases:**
- Multi-party computation (joint analysis without sharing data)
- Secure key management (keys never leave enclave)
- Privacy-preserving ML (train on encrypted data)
- Regulated workloads (healthcare, finance)

---

## Key Management Architecture

Keys are the foundation of encryption. Lose or expose the keys, and encryption is useless.

---

### Cloud KMS Services

```
┌─────────────────────────────────────────────────────────────────────┐
│ Cloud Key Management Services                                       │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  AWS KMS                    Azure Key Vault       GCP Cloud KMS     │
│  ├── Symmetric keys         ├── Keys              ├── Symmetric     │
│  ├── Asymmetric keys        ├── Secrets           ├── Asymmetric    │
│  ├── HMAC keys              ├── Certificates      ├── MAC keys      │
│  ├── Multi-region keys      ├── Managed HSM       ├── Import keys   │
│  └── External key store     └── BYOK              └── External KMS  │
│                                                                     │
│  Common features:                                                   │
│  • Automatic key rotation                                           │
│  • Access policies (IAM integration)                                │
│  • Audit logging                                                    │
│  • Hardware-backed (HSM)                                            │
│  • Regional/global options                                          │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

---

### Customer-Managed vs Provider-Managed Keys

| Aspect | Provider-Managed (SSE-S3) | Customer-Managed (SSE-KMS) |
|--------|---------------------------|----------------------------|
| Key creation | Provider | You (in KMS) |
| Key rotation | Automatic | You control |
| Access policy | Provider's policy | Your IAM policy |
| Audit trail | Limited | Full CloudTrail logging |
| Cost | Free | Per-request charges |
| Compliance | May not meet requirements | Full control for compliance |

**When to use Customer-Managed Keys (CMK):**
```
✓ Regulatory requirements (PCI, HIPAA, SOX)
✓ Need to audit every key usage
✓ Need to disable/revoke access immediately
✓ Cross-account data sharing with key-based access
✓ Key material requirements (BYOK)
```

**CMK policy example:**
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "Allow use of the key",
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::123456789012:role/AppRole"
      },
      "Action": [
        "kms:Encrypt",
        "kms:Decrypt",
        "kms:GenerateDataKey"
      ],
      "Resource": "*",
      "Condition": {
        "StringEquals": {
          "kms:ViaService": "s3.us-east-1.amazonaws.com"
        }
      }
    }
  ]
}
```

---

### Hardware Security Modules (HSM)

For the highest security requirements, HSMs provide tamper-resistant key storage.

```
┌─────────────────────────────────────────────────────────────────────┐
│ HSM Options                                                         │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  Cloud-Native HSM:                                                  │
│  ├── AWS CloudHSM              ($1.50/hour per HSM)                 │
│  ├── Azure Dedicated HSM       (Thales Luna)                        │
│  └── GCP Cloud HSM             (KMS with HSM-backing)               │
│                                                                     │
│  Third-Party HSM:                                                   │
│  ├── Thales Luna Network HSM                                        │
│  ├── Utimaco SecurityServer                                         │
│  └── Entrust nShield                                                │
│                                                                     │
│  HSM provides:                                                      │
│  • FIPS 140-2 Level 3 validation                                    │
│  • Tamper-evident/tamper-resistant hardware                         │
│  • Keys never leave the HSM in plaintext                            │
│  • Cryptographic operations performed IN the HSM                    │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

**When to use HSM:**
- Payment card processing (PCI DSS may require)
- Root CA key storage
- Highly regulated industries (finance, government)
- Contractual requirements from customers

---

### Key Rotation & Lifecycle

```
┌─────────────────────────────────────────────────────────────────────┐
│ Key Lifecycle                                                       │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  Create ──▶ Active ──▶ Rotate ──▶ Deprecated ──▶ Destroyed          │
│    │           │          │           │              │              │
│    │           │          │           │              │              │
│    ▼           ▼          ▼           ▼              ▼              │
│  Generate   Encrypt/   Create new   Decrypt only   Permanent        │
│  key        Decrypt    version      (old data)     deletion         │
│                                                                     │
│  AWS KMS automatic rotation:                                        │
│  • Creates new key material annually                                │
│  • Old versions retained for decryption                             │
│  • Same key ID, seamless to applications                            │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

**Enable automatic rotation:**
```bash
# AWS KMS
aws kms enable-key-rotation --key-id alias/my-key

# Verify
aws kms get-key-rotation-status --key-id alias/my-key
```

---

## Cloud Storage Security

---

### Object Storage (S3, Blob, GCS)

Object storage is the most common data breach source in cloud.

**S3 security checklist:**

```
┌─────────────────────────────────────────────────────────────────────┐
│ S3 Bucket Security Layers                                           │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  1. Block Public Access (Account & Bucket level)                    │
│     └── s3:BlockPublicAcls: true                                    │
│     └── s3:IgnorePublicAcls: true                                   │
│     └── s3:BlockPublicPolicy: true                                  │
│     └── s3:RestrictPublicBuckets: true                              │
│                                                                     │
│  2. Bucket Policy (Who can access what)                             │
│     └── Explicit deny for sensitive operations                      │
│     └── Require encryption                                          │
│     └── Require HTTPS                                               │
│                                                                     │
│  3. Encryption                                                      │
│     └── Default encryption enabled (SSE-KMS)                        │
│     └── Bucket key enabled (cost savings)                           │
│                                                                     │
│  4. Versioning                                                      │
│     └── Protect against accidental deletion                         │
│     └── Enable MFA delete for extra protection                      │
│                                                                     │
│  5. Access Logging                                                  │
│     └── S3 server access logs or CloudTrail data events             │
│                                                                     │
│  6. Object Lock (for compliance)                                    │
│     └── WORM - prevent deletion/modification                        │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

**Secure bucket policy:**
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "DenyNonHTTPS",
      "Effect": "Deny",
      "Principal": "*",
      "Action": "s3:*",
      "Resource": [
        "arn:aws:s3:::my-secure-bucket",
        "arn:aws:s3:::my-secure-bucket/*"
      ],
      "Condition": {
        "Bool": {
          "aws:SecureTransport": "false"
        }
      }
    },
    {
      "Sid": "DenyIncorrectEncryption",
      "Effect": "Deny",
      "Principal": "*",
      "Action": "s3:PutObject",
      "Resource": "arn:aws:s3:::my-secure-bucket/*",
      "Condition": {
        "StringNotEquals": {
          "s3:x-amz-server-side-encryption": "aws:kms"
        }
      }
    }
  ]
}
```

---

### Database Security

**Encryption layers for databases:**

```
┌─────────────────────────────────────────────────────────────────────┐
│ Database Encryption Architecture                                    │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  Application Layer:                                                 │
│  ┌───────────────────────────────────────────────────────────────┐  │
│  │ Application-level encryption                                  │  │
│  │ • Encrypt sensitive fields before storing                     │  │
│  │ • Examples: SSN, credit cards, health data                    │  │
│  │ • Use envelope encryption with KMS                            │  │
│  └───────────────────────────────────────────────────────────────┘  │
│                             ▼                                       │
│  Database Layer:                                                    │
│  ┌───────────────────────────────────────────────────────────────┐  │
│  │ Transparent Data Encryption (TDE)                             │  │
│  │ • Database encrypts entire data files                         │  │
│  │ • AWS RDS: Uses KMS                                           │  │
│  │ • Azure SQL: TDE with service-managed or CMK                  │  │
│  └───────────────────────────────────────────────────────────────┘  │
│                             ▼                                       │
│  Storage Layer:                                                     │
│  ┌───────────────────────────────────────────────────────────────┐  │
│  │ Volume/Disk Encryption                                        │  │
│  │ • EBS encryption, Azure Disk Encryption                       │  │
│  │ • Protects data on physical media                             │  │
│  └───────────────────────────────────────────────────────────────┘  │
│                                                                     │
│  All layers together = defense in depth                             │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

**Database access controls:**

```
RDS Security Configuration:

1. Network isolation
   └── Private subnet, no public IP
   └── Security group: only app tier

2. Authentication
   └── IAM database authentication (no passwords)
   └── Strong password policy if using native auth

3. Encryption
   └── Encryption at rest (KMS)
   └── SSL/TLS enforced for connections

4. Audit logging
   └── Database activity streams
   └── CloudTrail for API calls

5. Backup security
   └── Encrypted backups (automatic with encrypted DB)
   └── Cross-region backup for DR
```

---

## Data Loss Prevention (DLP)

DLP prevents sensitive data from leaving your environment.

```
┌─────────────────────────────────────────────────────────────────────┐
│ DLP Architecture                                                    │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  ┌───────────────────────────────────────────────────────────────┐  │
│  │ Data at Rest                                                  │  │
│  │ • Scan storage for sensitive data                             │  │
│  │ • AWS Macie, Azure Purview, GCP DLP                           │  │
│  └───────────────────────────────────────────────────────────────┘  │
│                                                                     │
│  ┌───────────────────────────────────────────────────────────────┐  │
│  │ Data in Motion                                                │  │
│  │ • Inspect network traffic                                     │  │
│  │ • Block/alert on sensitive data exfiltration                  │  │
│  │ • CASB, network DLP                                           │  │
│  └───────────────────────────────────────────────────────────────┘  │
│                                                                     │
│  ┌───────────────────────────────────────────────────────────────┐  │
│  │ Data in Use                                                   │  │
│  │ • Endpoint DLP                                                │  │
│  │ • Prevent copy/paste, screenshots                             │  │
│  │ • USB blocking                                                │  │
│  └───────────────────────────────────────────────────────────────┘  │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

**DLP policy examples:**

```yaml
# Example: Block credit card numbers from leaving via API
Policy: Prevent PCI Data Exfiltration
  Trigger: Outbound API response
  Detection:
    - Pattern: Credit card regex (Visa, MC, Amex)
    - Confidence: High (validated with Luhn)
  Action:
    - Severity: Critical
    - Response: Block and alert
    - Mask data in logs

# Example: Detect sensitive data in S3
Policy: Find Unprotected PII in S3
  Scan: All S3 buckets
  Detection:
    - SSN patterns
    - Passport numbers
    - Driver license
  Action:
    - Tag bucket as "Contains PII"
    - Create finding in Security Hub
    - Notify data owner
```

**AWS Macie example findings:**

```
Finding: S3 bucket contains credit card numbers

Bucket: customer-data-prod
Sensitive data types:
  - Credit Card Number: 1,247 occurrences
  - AWS Secret Access Key: 3 occurrences (!)
  - Email Address: 45,892 occurrences

Risk: HIGH
Recommendation:
  - Encrypt bucket with KMS
  - Rotate exposed AWS credentials immediately
  - Review bucket access policy
```

---

## Key Concepts to Remember

1. **Classify before protecting** - Know what data you have and its sensitivity
2. **Encrypt everything** - At rest, in transit, and consider in-use for sensitive workloads
3. **Control your keys** - Use CMK for regulated data, understand the tradeoffs
4. **Defense in depth for storage** - Block public access + policies + encryption + logging
5. **DLP is your last line** - Detect and prevent data exfiltration
6. **Automate discovery** - You can't protect data you don't know exists

---

## Practice Questions

**Q1:** Your company stores healthcare data (PHI) in AWS. An auditor asks you to demonstrate that the data is encrypted at rest with customer-controlled keys and that you can revoke access immediately if needed. How would you architect this?

<details>
<summary>View Answer</summary>

**Architecture for PHI with CMK:**

```
┌─────────────────────────────────────────────────────────────────────┐
│ HIPAA-Compliant Data Architecture                                   │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  1. KMS Customer Managed Key (CMK):                                 │
│     ├── Create dedicated CMK for PHI                                │
│     ├── Enable automatic annual rotation                            │
│     ├── Key policy restricts access to specific roles               │
│     └── CloudTrail logs every key usage                             │
│                                                                     │
│  2. S3 Configuration:                                               │
│     ├── Default encryption: SSE-KMS with CMK                        │
│     ├── Bucket policy denies non-KMS uploads                        │
│     ├── Versioning + MFA delete enabled                             │
│     └── S3 Access Logging enabled                                   │
│                                                                     │
│  3. RDS Configuration:                                              │
│     ├── Encryption at rest with same/different CMK                  │
│     ├── SSL/TLS enforced for connections                            │
│     ├── Database activity streams for audit                         │
│     └── Automated encrypted backups                                 │
│                                                                     │
│  4. Immediate Revocation Capability:                                │
│     └── Disable the CMK                                             │
│         aws kms disable-key --key-id alias/phi-key                  │
│         • All data immediately inaccessible                         │
│         • Can re-enable if needed                                   │
│         • For permanent: schedule key deletion (7-30 days)          │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

**Key policy for PHI CMK:**
```json
{
  "Statement": [
    {
      "Sid": "Allow PHI Application Access",
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::123456789012:role/PHI-Application-Role"
      },
      "Action": [
        "kms:Decrypt",
        "kms:GenerateDataKey"
      ],
      "Resource": "*",
      "Condition": {
        "StringEquals": {
          "kms:ViaService": [
            "s3.us-east-1.amazonaws.com",
            "rds.us-east-1.amazonaws.com"
          ]
        }
      }
    },
    {
      "Sid": "Allow Security Team to Disable",
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::123456789012:role/Security-Admin"
      },
      "Action": [
        "kms:DisableKey",
        "kms:EnableKey"
      ],
      "Resource": "*"
    }
  ]
}
```

**Evidence for auditor:**
- CloudTrail logs showing key creation and configuration
- KMS key rotation status
- S3 bucket encryption configuration
- RDS encryption configuration
- IAM policies limiting key access
- Procedure document for emergency key disablement

</details>

**Q2:** A developer accidentally committed AWS credentials to a public GitHub repository. The credentials have access to S3 buckets containing customer data. Walk through your incident response.

<details>
<summary>View Answer</summary>

**Incident Response for Exposed Credentials:**

**1. Immediate (0-15 minutes):**
```bash
# Disable the access key IMMEDIATELY
aws iam update-access-key \
  --user-name compromised-user \
  --access-key-id AKIAEXAMPLE \
  --status Inactive

# Or delete it
aws iam delete-access-key \
  --user-name compromised-user \
  --access-key-id AKIAEXAMPLE

# If it's a role, revoke all sessions
aws iam put-role-policy \
  --role-name compromised-role \
  --policy-name RevokeOldSessions \
  --policy-document '{
    "Version": "2012-10-17",
    "Statement": [{
      "Effect": "Deny",
      "Action": "*",
      "Resource": "*",
      "Condition": {
        "DateLessThan": {"aws:TokenIssueTime": "2024-01-15T12:00:00Z"}
      }
    }]
  }'
```

**2. Investigate (15 minutes - 2 hours):**
```
Check CloudTrail for unauthorized access:

aws cloudtrail lookup-events \
  --lookup-attributes AttributeKey=AccessKeyId,AttributeValue=AKIAEXAMPLE \
  --start-time 2024-01-01 \
  --end-time 2024-01-15

Look for:
• Data access (S3 GetObject, ListBucket)
• Data exfiltration (large downloads)
• Privilege escalation (IAM changes)
• Persistence (new users, roles, keys)
• Resource creation (EC2 for cryptomining)
```

**3. Contain & Eradicate (2-4 hours):**
```
If unauthorized access confirmed:
• Block IP addresses used by attacker (if identifiable)
• Rotate ALL credentials that may have been exposed
• Check for persistence mechanisms:
  - New IAM users/roles/policies
  - Lambda functions
  - EC2 instances
  - EventBridge rules
• Review S3 access logs for data accessed
```

**4. Recovery & Notification:**
```
• Issue new credentials to legitimate user
• If customer data was accessed:
  - Legal notification requirements (GDPR 72 hours, etc.)
  - Customer notification
  - Regulatory reporting
• Preserve evidence for potential investigation
```

**5. Post-Incident (Prevention):**
```
• Implement secrets scanning in CI/CD (git-secrets, TruffleHog)
• Enable GitHub secret scanning
• Use temporary credentials (IAM roles, OIDC federation)
• Reduce permissions (least privilege)
• Add MFA requirement for sensitive operations
• Security awareness training
```

</details>

**Q3:** Your organization is implementing a multi-region disaster recovery strategy. Data must be encrypted with customer-managed keys, and you need to ensure encryption works across regions. How do you design the key management?

<details>
<summary>View Answer</summary>

**Multi-Region Key Management Architecture:**

**Option 1: AWS Multi-Region Keys (Recommended)**
```
┌─────────────────────────────────────────────────────────────────────┐
│ Multi-Region Key Architecture                                       │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  Primary Region (us-east-1)      Replica Region (eu-west-1)         │
│  ┌─────────────────────┐         ┌─────────────────────┐            │
│  │ Primary MRK         │ ──────▶ │ Replica MRK         │            │
│  │ mrk-1234abcd        │  Sync   │ mrk-1234abcd        │            │
│  │                     │         │                     │            │
│  │ Same key ID         │         │ Same key ID         │            │
│  │ Same key material   │         │ Same key material   │            │
│  └─────────────────────┘         └─────────────────────┘            │
│           │                               │                         │
│           ▼                               ▼                         │
│  ┌─────────────────────┐         ┌─────────────────────┐            │
│  │ S3 (encrypted)      │  ────▶  │ S3 Replica          │            │
│  │ RDS (encrypted)     │  Cross  │ RDS Read Replica    │            │
│  │                     │  Region │                     │            │
│  └─────────────────────┘         └─────────────────────┘            │
│                                                                     │
│  Benefits:                                                          │
│  • Same key ID works in both regions                                │
│  • No re-encryption needed for cross-region copies                  │
│  • Independent operation during regional failure                    │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

**Create Multi-Region Key:**
```bash
# Create primary key
aws kms create-key \
  --multi-region \
  --description "Multi-region key for DR"

# Create replica in DR region
aws kms replicate-key \
  --key-id mrk-1234567890abcdef0 \
  --replica-region eu-west-1
```

**Option 2: Separate Keys with Re-encryption**
```
If multi-region keys don't meet requirements:

Primary (us-east-1):           DR (eu-west-1):
┌──────────────────┐           ┌──────────────────┐
│ Key A            │           │ Key B            │
│ (CMK for primary)│           │ (CMK for DR)     │
└────────┬─────────┘           └────────┬─────────┘
         │                              │
         ▼                              ▼
┌──────────────────┐           ┌──────────────────┐
│ Data encrypted   │  ──────▶  │ Data re-encrypted│
│ with Key A       │  Lambda   │ with Key B       │
└──────────────────┘  re-encrypt└──────────────────┘

Complexity: Higher (need re-encryption process)
Latency: Higher (decryption + re-encryption)
Use when: Regulatory requires regional key isolation
```

**Key policy for DR:**
```json
{
  "Statement": [
    {
      "Sid": "AllowDRAccess",
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::123456789012:role/DR-Application"
      },
      "Action": [
        "kms:Decrypt",
        "kms:GenerateDataKey"
      ],
      "Resource": "*",
      "Condition": {
        "StringEquals": {
          "kms:CallerAccount": "123456789012"
        }
      }
    }
  ]
}
```

**Testing DR key access:**
- Regular DR drills should verify key access
- Monitor key usage in replica region
- Alert if replica key is disabled

</details>

**Q4:** Explain the difference between envelope encryption and direct encryption. When and why would you use envelope encryption?

<details>
<summary>View Answer</summary>

**Direct Encryption:**
```
┌─────────────────────────────────────────────────────────────────────┐
│ Direct Encryption                                                   │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  Data ──────────▶ KMS ──────────▶ Encrypted Data                    │
│  (plaintext)      (encrypt)       (ciphertext)                      │
│                                                                     │
│  • KMS performs the encryption                                      │
│  • Data travels to KMS                                              │
│  • Limited to 4KB per request                                       │
│  • High latency for large data                                      │
│  • Expensive at scale ($0.03 per 10,000 requests)                   │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

**Envelope Encryption:**
```
┌─────────────────────────────────────────────────────────────────────┐
│ Envelope Encryption                                                 │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  Step 1: Generate Data Key                                          │
│  ┌───────────┐                                                      │
│  │    KMS    │ ──▶ Returns: Plaintext DEK + Encrypted DEK           │
│  └───────────┘                                                      │
│                                                                     │
│  Step 2: Encrypt Data Locally                                       │
│  ┌────────────────────────────────────────────┐                     │
│  │ Your Application                           │                     │
│  │                                            │                     │
│  │ Data ─── Plaintext DEK ───▶ Encrypted Data │                     │
│  │                                            │                     │
│  │ Delete plaintext DEK from memory           │                     │
│  └────────────────────────────────────────────┘                     │
│                                                                     │
│  Step 3: Store Together                                             │
│  ┌────────────────────────────────────────────┐                     │
│  │ Encrypted DEK + Encrypted Data             │                     │
│  │ (stored together)                          │                     │
│  └────────────────────────────────────────────┘                     │
│                                                                     │
│  Decryption: Send Encrypted DEK to KMS ──▶ Get Plaintext DEK        │
│              Use Plaintext DEK locally to decrypt data              │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

**Why Envelope Encryption:**

| Aspect | Direct | Envelope |
|--------|--------|----------|
| Data size limit | 4KB | Unlimited |
| Latency | High (network to KMS) | Low (local encryption) |
| Cost | High at scale | Low (1 KMS call per object) |
| Offline capability | No | Yes (cached DEK) |
| Performance | Slow | Fast |

**When to use Envelope Encryption:**
```
✓ Encrypting files larger than 4KB (basically everything)
✓ High-throughput encryption needs
✓ Encrypting data in applications
✓ S3 encryption (AWS does this automatically with SSE-KMS)
✓ Database field-level encryption

When direct KMS encryption is OK:
○ Encrypting small secrets (API keys, passwords)
○ Low-frequency encryption operations
○ When simplicity is priority over performance
```

**Code example (AWS Encryption SDK):**
```python
import aws_encryption_sdk

# Client handles envelope encryption automatically
client = aws_encryption_sdk.EncryptionSDKClient()

kms_key_provider = aws_encryption_sdk.StrictAwsKmsMasterKeyProvider(
    key_ids=["arn:aws:kms:us-east-1:123456789:key/mrk-xxx"]
)

# Encrypt (generates DEK, encrypts data, packages together)
ciphertext, header = client.encrypt(
    source=plaintext_data,
    key_provider=kms_key_provider
)

# Decrypt (extracts encrypted DEK, calls KMS, decrypts data)
plaintext, header = client.decrypt(
    source=ciphertext,
    key_provider=kms_key_provider
)
```

</details>

**Q5:** You are a security architect for a Global retail company. You have data scattered across AWS S3, on-premise SQL servers, and various SaaS applications like Salesforce. You need a unified view of your data estate to identify where sensitive customer PII is located. Which tool is best suited for this requirement and why?

<details>
<summary>View Answer</summary>

**Answer: Azure Purview (or Microsoft Purview)**

**Why:**
*   **Unified Governance:** Purview is explicitly designed for hybrid and multi-cloud scenarios.
*   **Broad Connectivity:** Unlike AWS Macie (which is AWS S3 focused) or GCP Cloud DLP (GCP focused), Azure Purview has "Data Map" collectors for on-premise SQL, multicloud storage (S3), and SaaS apps (Salesforce).
*   **Holistic View:** It provides a single pane of glass for governance across the disparate environments described.

**Why not others:**
*   **AWS Macie:** primarily scans S3.
*   **GCP Cloud DLP:** primarily for Google Cloud and streaming data.
*   **BigID:** Could also be a correct answer if "Cloud Native" wasn't implied, but Purview is the major cloud-provider offering for this scope.
</details>

**Q6:** Your finance team wants to run fraud detection models on extremely sensitive user transaction data. They want to use the cloud for scalability but are concerned that a malicious cloud admin or hypervisor vulnerability could expose the data while it is being processed in memory. What technology should you recommend?

<details>
<summary>View Answer</summary>

**Answer: Confidential Computing (Encryption in Use)**

**Technology to recommend:**
*   **AWS Nitro Enclaves**, **Azure Confidential Computing (SGX)**, or **GCP Confidential VMs**.

**Reasoning:**
*   **Encryption in Use:** Standard encryption at rest and in transit protects data on disk and network, but data is typically decrypted in RAM for processing.
*   **Isolation:** Confidential computing uses hardware-based execution environments (Trusted Execution Environments - TEEs) to isolate the memory.
*   **Threat Model:** This specifically addresses the "malicious cloud admin" or "hypervisor breakout" threat vectors, as the host system cannot see inside the encrypted memory enclave.
</details>

**Q7:** You need to enforce a policy that *no object* can be uploaded to your company's "sensitive-data" S3 bucket unless it is encrypted with server-side encryption using AWS KMS (SSE-KMS). Construct the S3 bucket policy condition statement to achieve this.

<details>
<summary>View Answer</summary>

**Answer: Deny PutObject where x-amz-server-side-encryption is NOT aws:kms**

```json
{
  "Sid": "DenyIncorrectEncryptionHeader",
  "Effect": "Deny",
  "Principal": "*",
  "Action": "s3:PutObject",
  "Resource": "arn:aws:s3:::sensitive-data/*",
  "Condition": {
    "StringNotEquals": {
      "s3:x-amz-server-side-encryption": "aws:kms"
    }
  }
}
```

*Note: This works because `StringNotEquals` will match if the header is missing OR if the header is present but has a different value (like AES256).*
</details>

**Q8:** A government client has a strict regulatory requirement that all encryption keys must be generated and stored in a device that is FIPS 140-2 Level 3 validated. They also require that the cloud provider have absolutely no visibility into the key generation material. Which key management solution should you choose?

<details>
<summary>View Answer</summary>

**Answer: AWS CloudHSM / Azure Dedicated HSM**

**Reasoning:**
*   **FIPS 140-2 Level 3:** Standard KMS services (like AWS KMS) are typically FIPS 140-2 Level 2 validated (some parts Level 3, but generally considered Level 2 for the service wrapping). CloudHSM provides dedicated hardware that is fully Level 3 validated.
*   **Single Tenancy:** The requirement asks for keys to be stored in a device where the provider has no visibility. CloudHSM gives you a single-tenant hardware appliance where you hold the crypto-officer credentials.
*   **KMS Custom Key Store:** You could also link KMS to CloudHSM (Custom Key Store), but the primary requirement driver here for "pure" isolation usually points directly to the HSM service.
</details>

**Q9:** A legacy healthcare application processes highly sensitive patient data in memory. The application runs on Linux, is written in C++, and the source code is no longer available to be recompiled. The Chief CISO wants to move this to the cloud but requires that the memory be encrypted to protect against hypervisor-level attacks. Which Confidential Computing technology is most appropriate?

<details>
<summary>View Answer</summary>

**Answer: AMD SEV (Secure Encrypted Virtualization) / GCP Confidential VMs**

**Reasoning:**
*   **"No source code available":** Intel SGX (App Enclaves) requires you to use an SDK and recompile the application to partition code into trusted/untrusted parts.
*   **"Lift and Shift":** AMD SEV encrypts the *entire* VM memory transparently to the OS and application. No code changes are required.
*   **Suitability:** This fits the requirement of protecting legacy apps without refactoring.
</details>

**Q10:** You are designing a tokenizer service that accepts credit card numbers and returns a token. The processing of the credit card number must happen in an isolated environment where even the root user of the EC2 instance hosting the service cannot access the plaintext data or the memory. Which AWS service fits this description?

<details>
<summary>View Answer</summary>

**Answer: AWS Nitro Enclaves**

**Reasoning:**
*   **Isolation:** Nitro Enclaves carves out vCPUs and memory from a parent EC2 instance to create a fully isolated environment.
*   **No Access:** The Enclave has no persistent storage, no interactive access (no SSH), and even the root user of the parent instance cannot peer into the enclave's memory.
*   **Communication:** Data is sent via a secure local channel (vsock). This is the classic use case for critical processing like tokenization or crypto operations.
</details>

---

## Next Up

In Lesson 6, we'll cover **Multi-Cloud Security, CASB & Architecture Patterns** — bringing together everything you've learned into real-world enterprise architectures!
