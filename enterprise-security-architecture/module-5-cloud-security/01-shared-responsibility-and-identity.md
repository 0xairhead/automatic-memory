# Module 5, Lesson 1: Shared Responsibility Model & Cloud Identity

## Table of Contents
- [Media Resources](#media-resources)
- [The Shared Responsibility Model](#the-shared-responsibility-model)
  - [The Apartment Analogy](#the-apartment-analogy)
  - [IaaS vs PaaS vs SaaS: Who Owns What?](#iaas-vs-paas-vs-saas-who-owns-what)
  - [AWS, Azure, and GCP Specifics](#aws-azure-and-gcp-specifics)
  - [Compliance Inheritance](#compliance-inheritance)
- [Cloud Identity & Access Management](#cloud-identity--access-management)
  - [Understanding Cloud IAM](#understanding-cloud-iam)
  - [Service Accounts & Workload Identity](#service-accounts--workload-identity)
  - [Cross-Account Access Patterns](#cross-account-access-patterns)
  - [Just-in-Time Privileged Access](#just-in-time-privileged-access)
  - [Federation with On-Premises Directories](#federation-with-on-premises-directories)
  - [Identity Protocols Reference](#identity-protocols-reference)
- [Key Concepts to Remember](#key-concepts-to-remember)
- [Practice Questions](#practice-questions)
- [Next Up](#next-up)

---

Welcome to Cloud Security! This foundational lesson covers the most critical concept you'll encounter: who is responsible for what in the cloud.

## Media Resources

**Visual Guide:**

![Shared Responsibility Model](./assets/shared-responsibility-model.png)

**Audio Lecture:**

🎧 [Shared Responsibility & Identity Management (Audio)](./assets/01-shared-responsibility-audio.m4a)

---

## The Shared Responsibility Model

Here's the uncomfortable truth about cloud security: **the cloud provider is NOT responsible for all your security.** Many organizations learn this the hard way after a breach.

### The Apartment Analogy

Think of cloud computing like renting an apartment:

| Responsibility | Landlord (Cloud Provider) | Tenant (You) |
|---------------|---------------------------|--------------|
| Building structure | Yes | No |
| Locks on main entrance | Yes | No |
| Fire suppression system | Yes | No |
| Your apartment's lock | No | Yes |
| Who you give keys to | No | Yes |
| What you store inside | No | Yes |
| Leaving windows open | No | Yes |

If someone breaks in because you left your door unlocked, that's on you—not the landlord. The same applies to cloud security.

### IaaS vs PaaS vs SaaS: Who Owns What?

The responsibility split changes dramatically based on the service model:

```
                    YOU MANAGE
                        │
    ┌───────────────────┼───────────────────┐
    │                   │                   │
    ▼                   ▼                   ▼
┌────────┐        ┌────────┐         ┌────────┐
│  IaaS  │        │  PaaS  │         │  SaaS  │
├────────┤        ├────────┤         ├────────┤
│ Data   │        │ Data   │         │ Data   │
│ Apps   │        │ Apps   │         │  ---   │
│ Runtime│        │  ---   │         │  ---   │
│ OS     │        │  ---   │         │  ---   │
│  ---   │        │  ---   │         │  ---   │
│  ---   │        │  ---   │         │  ---   │
└────────┘        └────────┘         └────────┘
                        │
                PROVIDER MANAGES
```

**IaaS (Infrastructure as a Service)** - EC2, Azure VMs, GCP Compute Engine
- **You manage:** Data, Applications, Runtime, Middleware, OS patching
- **Provider manages:** Virtualization, Servers, Storage, Networking, Physical DC

**PaaS (Platform as a Service)** - AWS Elastic Beanstalk, Azure App Service, Cloud Run
- **You manage:** Data, Applications
- **Provider manages:** Everything else including runtime and OS

**SaaS (Software as a Service)** - Salesforce, Microsoft 365, Workday
- **You manage:** Data, User access, Configuration
- **Provider manages:** The entire application stack

### AWS, Azure, and GCP Specifics

While the "Shared Responsibility" concept is universal, each provider explains it in their own unique flavor. Understanding these nuances helps you "speak the language" of that cloud.

#### 1. AWS Model: "Of" vs. "In" (The Classic)

AWS uses the clearest, most binary definition. They split the world into two distinct distinct piles:

**Security OF the Cloud (AWS's Job)**
*   **Think:** " Concrete and Steel."
*   **What it is:** The global infrastructure that runs the services.
*   **Includes:** Physical data centers, the servers rack-mounted inside them, the cabling, the air conditioning, and the virtualization software (Hypervisor) that creates the VMs.
*   **Analogy:** Hertz Rental Car. Hertz guarantees the car has an engine, tires, and safe brakes.

**Security IN the Cloud (Your Job)**
*   **Think:** "Data and Configuration."
*   **What it is:** Everything you put *into* that infrastructure.
*   **Includes:** Your customer data, your encryption keys, your operating system patches (for EC2), and your firewall rules.
*   **Analogy:** The Driver. You are responsible for wearing a seatbelt, not speeding, and locking the car when you park.

#### 2. Azure's Model: The Sliding Scale

Microsoft visualizes responsibility not as a binary switch, but as a sliding scale that changes based on the service type. They group responsibilities into three buckets:

1.  **Always Yours:**
    *   **Data & Endpoints:** No matter what, you own your data and the devices (laptops/phones) accessing it.
    *   **Accounts & Access:** You always control who logs in.

2.  **Always Theirs:**
    *   **Physical Hosts & Network:** You will never repair a server or splice a fiber optic cable in an Azure data center.

3.  **Varies (The "Slider"):**
    *   **OS & Network Controls:** This moves based on the service. In IaaS (VMs), it's your job. In PaaS (App Service), it's their job. In SaaS (Office 365), it's entirely their job.

#### 3. GCP's Model: "Shared Fate" (The Partner)

Google takes a more modern, holistic approach. They noticed that "Shared Responsibility" often felt like "Shared Blame" (e.g., "Well, you left the door open, so it's your fault").

**What is Shared Fate?**
*   **Philosophy:** "If you fail, we fail. So let's help you not fail."
*   **Active Assistance:** Instead of just drawing a line in the sand, Google actively provides tools to ensure you stay secure on your side of the line.
*   **How they do it:**
    *   **Secure Blueprints:** Pre-made, gold-standard architecture templates so you don't build from scratch.
    *   **Guardrails:** Policy intelligence tools that warn you *before* you deploy something insecure.
    *   **Default Security:** Turning on encryption and safety features by default, rather than making you hunt for the "on" switch.

### Compliance Inheritance

Here's what many architects miss: **You inherit some compliance controls, but not all.**

```
Example: PCI DSS Compliance on AWS

AWS provides:
✓ Physical security (Requirement 9)
✓ Network infrastructure (parts of Req 1)
✓ Some encryption capabilities (Req 3, 4)

You must still handle:
✗ Cardholder data protection (how you store/encrypt)
✗ Access control (who can see the data)
✗ Vulnerability management (your application patches)
✗ Security testing (your code and configurations)
```

**Key insight:** When a compliance auditor comes, you need TWO things:
1. The cloud provider's compliance attestation (SOC 2 report, ISO cert)
2. YOUR evidence of controls on top of the provider's

---

## Cloud Identity & Access Management

If Shared Responsibility is the "what" of cloud security, IAM is the "who." Every cloud breach investigation starts with one question: **Who had access?**

### Understanding Cloud IAM

Cloud IAM is fundamentally different from traditional on-premises identity:

**Traditional (On-Prem):**
```
User → Active Directory → Application
         └── LDAP/Kerberos ──┘
```

**Cloud IAM:**
```
User/Service → Cloud IAM → 1000s of Services & APIs
                   │
         ┌────────┴────────┐
         ▼                 ▼
    Fine-grained      Policies attached to
    permissions       resources AND identities
```

**AWS IAM Core Concepts:**
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "s3:GetObject",
      "Resource": "arn:aws:s3:::my-bucket/*",
      "Condition": {
        "IpAddress": {"aws:SourceIp": "192.168.1.0/24"}
      }
    }
  ]
}
```

**Azure IAM (Entra ID + RBAC):**
- Entra ID (formerly Azure AD) for identity
- Role-Based Access Control for authorization
- Managed Identities for workloads

**GCP IAM:**
- Resource hierarchy (Org → Folder → Project → Resource)
- Permissions granted at any level inherit downward
- Service accounts for workload identity

### Service Accounts & Workload Identity

**The problem:** How does your application authenticate to cloud services?

**Bad approach (but common):**
```
# Hardcoded credentials - NEVER DO THIS
AWS_ACCESS_KEY_ID=AKIAIOSFODNN7EXAMPLE
AWS_SECRET_ACCESS_KEY=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
```

**Good approach - Workload Identity:**

```
┌─────────────────┐     ┌──────────────────┐     ┌─────────────┐
│ Your Application│────▶│ Cloud Metadata   │────▶│ Cloud IAM   │
│ (no credentials)│     │ Service          │     │             │
└─────────────────┘     └──────────────────┘     └─────────────┘
                              │
                              ▼
                        Temporary, auto-rotated
                        credentials injected
```

**AWS:** IAM Roles + Instance Profiles, EKS IRSA (IAM Roles for Service Accounts)
**Azure:** Managed Identities (System-assigned or User-assigned)
**GCP:** Workload Identity, Service Account Keys (avoid if possible)

### Cross-Account Access Patterns

Enterprise environments have multiple accounts/subscriptions. How do you manage access across them?

**AWS Cross-Account Access:**
```
┌─────────────────────────────────────────────────────────────┐
│ Account A (Development)                                     │
│  └── Role: "DeveloperRole"                                  │
│       └── Trust Policy: Allow Account B to assume this role │
└─────────────────────────────────────────────────────────────┘
                              │
                              │ AssumeRole
                              ▼
┌─────────────────────────────────────────────────────────────┐
│ Account B (Security/Audit)                                  │
│  └── User/Role needs: sts:AssumeRole permission             │
└─────────────────────────────────────────────────────────────┘
```

**Azure Cross-Tenant:**
- Azure Lighthouse for managed service providers
- B2B collaboration in Entra ID
- Cross-tenant access policies

**GCP Cross-Project:**
- Shared VPC for network resources
- IAM policies at organization/folder level
- Service account impersonation

### Just-in-Time Privileged Access

**The principle:** No one should have standing privileged access. Access should be:
- Requested when needed
- Approved (human or automated)
- Time-limited
- Fully audited

**Implementation patterns:**

```
Normal state:           JIT Request:              Active session:
┌──────────────┐       ┌──────────────┐          ┌──────────────┐
│ Admin User   │       │ Admin User   │          │ Admin User   │
│ - Read only  │  ──▶  │ - Requests   │  ──▶     │ - Full admin │
│ - No prod    │       │   elevation  │          │ - 4 hours    │
└──────────────┘       │ - Justifies  │          │ - Logged     │
                       └──────────────┘          └──────────────┘
                              │
                              ▼
                       Approval workflow
                       (manager, ticket, auto)
```

**Tools:**
- AWS: IAM Identity Center + Permission Sets
- Azure: Privileged Identity Management (PIM)
- GCP: Privileged Access Manager (PAM)
- Third-party: CyberArk, HashiCorp Boundary, Teleport

### Federation with On-Premises Directories

Most enterprises already have Active Directory. You don't want to recreate all users in the cloud.

**Federation flow:**
```
┌───────────┐     ┌─────────────┐     ┌─────────────┐     ┌───────────┐
│ User      │────▶│ Corporate   │────▶│ IdP (ADFS,  │────▶│ Cloud     │
│           │     │ Portal      │     │ Okta, etc.) │     │ Provider  │
└───────────┘     └─────────────┘     └─────────────┘     └───────────┘
                                             │
                                             ▼
                                       SAML Assertion
                                       or OIDC Token
```

**Key protocols:**
- **SAML 2.0:** Enterprise standard, XML-based, mature
- **OIDC:** Modern, JSON/REST-based, mobile-friendly
- **SCIM:** For user provisioning/deprovisioning sync

**Architecture considerations:**
- High availability of IdP (it becomes a critical path)
- Conditional access policies (MFA, device compliance, location)
- Break-glass accounts for IdP outages (secured, monitored, rarely used)

### Identity Protocols Reference

This lesson covers several key protocols and mechanisms. Here is a quick reference guide:

#### 1. Federation Protocols
Used to "federate" (connect) your existing corporate identity (like Active Directory) with cloud providers so you don't have to recreate users.

*   **SAML 2.0 (Security Assertion Markup Language):**
    *   **Description:** The mature, enterprise standard XML-based protocol.
    *   **Use Case:** The most common way to log users into the cloud console using their corporate credentials. It passes "assertions" (XML documents) stating who the user is.
*   **OIDC (OpenID Connect):**
    *   **Description:** A more modern, JSON/REST-based protocol built on top of OAuth 2.0.
    *   **Use Case:** Often preferred for mobile-friendly or modern web applications, serving a similar authentication purpose as SAML but with lighter-weight JSON tokens.

#### 2. Synchronization Protocols
*   **SCIM (System for Cross-domain Identity Management):**
    *   **Description:** A protocol specifically for user provisioning.
    *   **Use Case:** It handles the *lifecycle* of a user. If you hire someone, SCIM automatically creates their account in the cloud. Critical for security: if you fire someone and disable them in Active Directory, SCIM automatically detects this and deactivates their cloud access.

#### 3. Legacy Protocols
*   **LDAP / Kerberos:**
    *   **Context:** Used in *traditional on-premises* Active Directory environments (User -> Active Directory). This contrasts with how Cloud IAM works (API-based, fine-grained policies).

#### 4. Cloud-Native Mechanisms
While not "protocols" in the strict networking sense, these are the identity mechanisms specific to each cloud provider:

*   **AWS:** **IAM Roles** (for assuming identity) and **Instance Profiles**.
*   **Azure:** **Entra ID** (formerly Azure AD) and **Managed Identities**.
*   **GCP:** **Workload Identity** and **Services Accounts**.

---

## Key Concepts to Remember

1. **Cloud providers secure the cloud; you secure what's IN the cloud**
2. **Responsibility shifts based on service model** - IaaS = more on you, SaaS = less on you
3. **Never use long-lived credentials** - Use workload identity, managed identities, IAM roles
4. **Least privilege + Just-in-Time** - No standing access, time-limited elevation
5. **Federate, don't duplicate** - Use your existing IdP, sync via SCIM
6. **Compliance is shared** - You inherit some controls, but must prove your own

---

## Practice Questions

**Q1:** Your company is running a containerized application on AWS ECS. A security incident reveals that an attacker gained access to your S3 buckets. The ECS tasks were using an IAM role with `s3:*` permissions on all buckets. Under the shared responsibility model, who is responsible for this breach?

<details>
<summary>View Answer</summary>

**The customer (you) is responsible.**

This is a classic example of over-privileged access—a customer responsibility. AWS secured the infrastructure, but:
- The customer created an IAM role with excessive permissions (`s3:*` on all buckets)
- The customer failed to follow least privilege principles
- The customer should have scoped permissions to specific buckets and actions needed

**The fix:**
- Scope IAM policies to specific S3 buckets
- Use only required actions (e.g., `s3:GetObject`, `s3:PutObject`)
- Add conditions (VPC endpoint, source IP, etc.)
- Implement S3 bucket policies as a second layer of defense

</details>

**Q2:** You're designing a multi-account AWS architecture for a healthcare company. They need to be HIPAA compliant. The CTO asks: "AWS is HIPAA compliant, so we're automatically compliant, right?" How do you respond?

<details>
<summary>View Answer</summary>

**No, that's a dangerous misconception.**

Your response should cover:

1. **AWS's HIPAA compliance means:**
   - AWS will sign a Business Associate Agreement (BAA)
   - Certain AWS services are HIPAA-eligible
   - AWS maintains controls for their portion of responsibility

2. **You still must:**
   - Configure services correctly (encryption at rest and in transit)
   - Implement access controls (who can access PHI)
   - Enable logging and monitoring (CloudTrail, VPC Flow Logs)
   - Manage your application security
   - Train your employees
   - Have incident response procedures
   - Conduct risk assessments

3. **Compliance inheritance:**
   - You can point to AWS's SOC 2 and HIPAA attestations for physical/infrastructure controls
   - You must provide your own evidence for data handling, access control, and application-level security

</details>

**Q3:** A developer asks you to create a service account with a downloadable JSON key for their application running on GKE (Google Kubernetes Engine). They say it's "just for testing." What's wrong with this approach, and what would you recommend instead?

<details>
<summary>View Answer</summary>

**Problems with downloadable service account keys:**
- Keys don't expire automatically
- Keys can be copied, shared, or leaked (git commits, logs, screenshots)
- No built-in rotation mechanism
- Difficult to audit where the key is being used
- Violates security best practices

**Recommended approach - Workload Identity:**

1. Enable Workload Identity on the GKE cluster
2. Create a Kubernetes ServiceAccount
3. Create a GCP Service Account with minimal required permissions
4. Bind them together:
```bash
gcloud iam service-accounts add-iam-policy-binding \
  GSA_NAME@PROJECT_ID.iam.gserviceaccount.com \
  --role roles/iam.workloadIdentityUser \
  --member "serviceAccount:PROJECT_ID.svc.id.goog[NAMESPACE/KSA_NAME]"
```
5. Annotate the Kubernetes ServiceAccount

**Benefits:**
- No keys to manage or leak
- Credentials auto-rotate (short-lived tokens)
- Clear audit trail
- Per-pod identity

</details>

**Q4:** Explain why "break-glass" accounts are necessary when using federated identity, and describe three security controls you would implement for them.

<details>
<summary>View Answer</summary>

**Why break-glass accounts are necessary:**

When you federate identity to your corporate IdP (Okta, Azure AD, ADFS), that IdP becomes a single point of failure. If the IdP goes down:
- No one can log into the cloud console
- No one can manage cloud resources
- You can't fix the IdP issue if it requires cloud access
- Business continuity is compromised

**Three essential security controls:**

1. **Extremely strong authentication:**
   - Long, complex passwords (30+ characters) stored in a physical safe
   - Hardware security keys (YubiKeys) required for MFA
   - Multiple keys distributed geographically

2. **Alerting and monitoring:**
   - Immediate alert (PagerDuty, SMS) on any break-glass login
   - CloudTrail/Azure Activity Log monitored for these accounts
   - Automated ticket creation for mandatory review
   - SIEM correlation rules for break-glass usage

3. **Minimal permissions with time constraints:**
   - Start with minimal permissions, elevate only when needed
   - Sessions limited to shortest practical duration
   - Automatic password rotation after any use
   - Quarterly access review and testing

**Additional controls:** Store credentials in separate locations, require two-person access (split password/MFA), document approved use cases, and conduct annual break-glass drills.

</details>

**Q5: The "Always-On" Admin**
**Scenario:** You have hired a new Lead Cloud Architect. To ensure they can do their job effectively, you create a user account for them in AWS and attach the `AdministratorAccess` policy (granting full control). This permission is permanent so they don't have to ask for approval every time they need to work. According to the lesson's security principles, what is the flaw in this approach, and which specific IAM concept resolves it?

<details>
<summary>View Answer</summary>

**The flaw is "Standing Privileged Access."**

Granting permanent (24/7) administrative access violates the principle of Least Privilege and expands the attack surface. If this architect's credentials are compromised, the attacker has immediate, unfettered access to the entire environment.

**The solution is Just-in-Time (JIT) Access.**

Instead of permanent permissions, the architect should have a standard user account. When they need to perform administrative tasks, they should request elevated access that is:
1.  **Time-limited** (e.g., valid for only 4 hours).
2.  **Approved** (either automatically based on policy or by a peer).
3.  **Audited** (logging exactly why and when access was elevated).

*Recommended Tools:* AWS IAM Identity Center (Permission Sets), Azure Privileged Identity Management (PIM), or GCP Privileged Access Manager (PAM).

</details>

**Q6: The Federation Dilemma**
**Scenario:** A company with 5,000 employees uses Microsoft Active Directory (AD) on-premises. They are migrating to Google Cloud. The IT manager suggests exporting a CSV list of all users from AD and importing them into Google Cloud Identity manually to get started quickly. Why is this considered a bad practice, and what architecture should be used instead?

<details>
<summary>View Answer</summary>

**Bad Practice:**
Manually recreating users creates two separate identities for every employee ("Identity Islands"). This leads to:
*   **Drift:** If an employee leaves the company and is disabled in AD, their Cloud account remains active and dangerous.
*   **Management Overhead:** You must manage passwords and policies in two places.

**Correct Architecture: Federation.**
You should **"Federate, don't duplicate"**. The company should connect Google Cloud to their existing Active Directory using standard protocols like **SAML 2.0** or **OIDC**. This allows the on-premises AD to remain the "source of truth." When a user is disabled in AD, their access to the cloud is automatically revoked. You can also use **SCIM** to handle the syncing of user provisioning and deprovisioning.

</details>

**Q7: The Hardcoded VM Credential**
**Scenario:** You are reviewing code for an application running on an Azure Virtual Machine. The application needs to read sensitive files from a storage account. You notice the developer has hardcoded a "Connection String" (containing a secret key) in the application's config file. The developer argues this is secure because the config file is inside the private VM. How do you refute this, and what is the correct Azure-native solution?

<details>
<summary>View Answer</summary>

**Refutation:**
Hardcoded credentials are a security failure because:
*   If the code is ever committed to a repository (like GitHub), the secret is leaked.
*   The key is "long-lived" (it doesn't rotate automatically), making it valid indefinitely if stolen.

**Correct Solution:**
Use **Managed Identities** (specifically a System-assigned or User-assigned Managed Identity).
1.  Assign an identity to the VM itself.
2.  Grant that identity access to the storage account via Azure RBAC.
3.  The application requests a token from the local Azure instance metadata service.
4.  No secrets or keys are ever stored in the code or config files, and the cloud provider handles the credential rotation automatically.

</details>

**Q8: The Multi-Account User Problem**
**Scenario:** Your organization uses AWS and has separated environments into two accounts: `Development` and `Production`. A senior developer needs access to debug logs in the `Production` account. The System Administrator decides to create a new IAM User named `dev_admin_prod` directly inside the `Production` account for this developer to use.

**Why is this approach incorrect, and what is the secure alternative?**

<details>
<summary>View Answer</summary>

*   **The Problem:** This violates the principle of **"Federate, don't duplicate"** and creates identity sprawl. The developer now has two sets of long-term credentials to manage (one for Dev, one for Prod). If the developer leaves the company, you must remember to revoke credentials in multiple places.
*   **The Solution:** Use **Cross-Account Access** patterns (specifically IAM Roles). The developer should log into their primary account (or Identity Provider) and "assume a role" in the `Production` account. This grants temporary access without creating a new permanent user or credential in the target account.

</details>

**Q9: The "Valid" Login from an Unknown Device**
**Scenario:** You have successfully federated your Cloud Identity with your on-premises Active Directory using SAML 2.0. A hacker manages to steal a user’s valid username and password via a phishing attack. The hacker successfully logs into the cloud console from a laptop in a different country. The identity system recognized the credentials as valid, so it let them in.

**Which specific identity control mentioned in the lesson was missing that could have prevented this breach?**

<details>
<summary>View Answer</summary>

*   **The Missing Control:** **Conditional Access Policies**.
*   **Explanation:** Federation and strong passwords are not enough. You must implement policies that evaluate the *context* of the login attempt. Source explicitly lists "Conditional access policies" such as **MFA**, **device compliance**, and **location**. If a policy had been set to block logins from unknown devices or unexpected geographic locations, the valid credentials would have been useless to the hacker.

</details>

**Q10: The Accidental Admin (Hierarchy Risk)**
**Scenario:** You are using Google Cloud Platform (GCP). To help a junior team member organize resources, you grant them `Project Editor` permissions at the **Folder** level (which contains 10 different Projects). The next day, you realize this user has full modify access to a critical database in a Project they shouldn't be touching.

**How did this happen, according to the GCP Identity model?**

<details>
<summary>View Answer</summary>

*   **The Cause:** **Permission Inheritance.**
*   **Explanation:** In GCP's resource hierarchy (Org → Folder → Project → Resource), permissions granted at a higher level automatically inherit downward. By granting the role at the Folder level, you inadvertently propagated that access to every single Project and Resource inside that folder.
*   **The Lesson:** Always apply **Least Privilege** by granting permissions at the lowest possible level of the hierarchy necessary for the task.

</details>

**Q11: Why are break-glass accounts essential when using federated identity?**

<details>
<summary>View Answer</summary>

Break-glass accounts are essential when using federated identity because the Identity Provider (IdP) becomes a **single point of failure**.

If your organization relies exclusively on an external IdP (like Okta, Azure AD, or ADFS) to log in to the cloud, and that service goes offline, the following critical issues occur:

*   **Total Lockout:** No one can log into the cloud console.
*   **Loss of Control:** You cannot manage cloud resources or fix the underlying issue if the resolution requires cloud access.
*   **Business Continuity Failure:** Critical operations may halt because administrative access is severed.

**What is a Break-Glass Account?**
A break-glass account is a highly secured **local** account that bypasses the federation process, allowing entry when the standard digital keys fail.

**Essential Security Controls**
Because these accounts bypass your standard security checks (like your corporate SSO), the sources recommend securing them with the following controls:
*   **Physical Security:** Use long, complex passwords (30+ characters) stored in a physical safe, along with hardware security keys (like YubiKeys) distributed geographically.
*   **Immediate Alerting:** Configure systems to trigger immediate alerts (via SMS or PagerDuty) the moment a break-glass login is detected.
*   **Procedural Rigor:** Require "two-person access" (split passwords), conduct quarterly access reviews, and run annual "break-glass drills" to ensure the process works during an emergency.

</details>

**Q12: What are the security risks of using service account keys?**

<details>
<summary>View Answer</summary>

The security risks of using service account keys (specifically downloadable keys, such as JSON files) include:

*   **Indefinite Validity:** Unlike modern identity tokens, these keys do not expire automatically. They are considered "long-lived credentials," meaning they remain valid until someone manually deletes them.
*   **High Risk of Leakage:** Because these keys are often saved as files, they are easily copied, shared, or accidentally exposed. Common leak vectors include accidental commits to source code repositories (like GitHub), inclusion in system logs, or visibility in screenshots.
*   **Lack of Rotation:** There is no built-in mechanism to rotate these keys. If a key is stolen, the attacker has persistent access because the credential does not change automatically.
*   **Audit Challenges:** It is difficult to track exactly where a specific key is being used, making it hard to detect if unauthorized parties are utilizing it.

**The Recommended Alternative**
To mitigate these risks, avoid using downloadable keys. Instead, use **Workload Identity** (such as AWS IAM Roles or Azure Managed Identities).

</details>

**Q13: What are some third-party tools for federation access and what is the role of tools like PagerDuty?**

<details>
<summary>View Answer</summary>

**Third-Party Tools for Federation & Access**
The sources distinguish between **Identity Providers** (who hold the user directory) and **Privileged Access tools** (which manage permissions dynamically).

**Identity Providers (IdPs) for Federation**
These are the directories you "federate" (connect) to your cloud so you don't have to recreate users manually:
*   **Okta:** A common cloud-based IdP used to manage corporate identities.
*   **ADFS (Active Directory Federation Services):** The traditional on-premises solution for bridging Active Directory with the cloud.
*   **Azure AD (Entra ID):** Can be federated with other clouds.

**Just-in-Time (JIT) & Privileged Access Tools**
For managing temporary, elevated access, these tools are often used:
*   **CyberArk**
*   **HashiCorp Boundary**
*   **Teleport**

**The Role of PagerDuty**
PagerDuty is **not** an identity or federation tool. It serves a critical role in **Monitoring and Alerting**, specifically for **"Break-Glass" accounts**.

**The Specific Use Case:**
When you use a "Break-Glass" account (an emergency account that bypasses your federation tools), it is a high-risk event. You must implement "Immediate Alerting" to notify the security team the instant a login is detected. PagerDuty is the mechanism used to trigger that immediate alert.

</details>

**Q14: Explain the difference between service account keys and workload identity.**

<details>
<summary>View Answer</summary>

**Service Account Keys (The "Bad" Approach)**
This method involves generating a downloadable credential file (often a JSON file) that an application uses to log in.
*   **Static & Long-Lived:** These keys do not expire automatically.
*   **High Leakage Risk:** Easily copied, shared, or leaked in code/logs.
*   **Management Burden:** You must manually rotate and distribute keys.

**Workload Identity (The "Good" Approach)**
This method allows the application to authenticate using temporary credentials generated by the cloud platform itself.
*   **Dynamic & Short-Lived:** The cloud provider issues a temporary token that expires automatically.
*   **Secure Storage:** You do not handle or store any secret files.
*   **Auto-Rotation:** The cloud provider handles rotation.

**Analogy**
*   **Service Account Keys** are like a **physical metal key**. If lost, it works forever until the lock is changed.
*   **Workload Identity** is like a **digital badge** that changes its code every shift.

</details>

**Q15: The Rental Car Analogy (AWS)**
**Scenario:** You launch an EC2 instance on AWS. Six months later, the underlying hard drive fails, causing data loss. At the same time, hackers breach the instance because the operating system wasn't patched.
**According to the AWS "Of vs. In" model, who is responsible for which failure?**

<details>
<summary>View Answer</summary>

*   **Hard Drive Failure:** **AWS.** This is "Security OF the Cloud." AWS guarantees the physical infrastructure and basic compute/storage hardware (the "car").
*   **Hackers/Unpatched OS:** **You.** This is "Security IN the Cloud." You are the "driver." AWS gives you the car, but if you drive it into a wall (or don't lock the doors/patch the OS), that is your responsibility.

</details>

**Q16: The Sliding Scale (Azure)**
**Scenario:** You are migrating an application from an on-premises server to Azure. You are debating between using an **Azure Virtual Machine (IaaS)** or **Azure App Service (PaaS)**.
**How does your responsibility for "Network Controls" change between these two choices according to Azure's model?**

<details>
<summary>View Answer</summary>

*   **Virtual Machine (IaaS):** **You own it.** You must configure the detailed firewall rules (Network Security Groups), subnetting, and virtual network integration yourself.
*   **App Service (PaaS):** **Shared/Microsoft.** Microsoft handles the underlying network connectivity and basic DDoS protection. You only configure high-level access restrictions (like "allow traffic only from this IP"). The "slider" has moved towards the provider.

</details>

**Q17: Shared Fate vs. Shared Responsibility (GCP)**
**Scenario:** A startup is choosing a cloud provider. They are worried they don't have enough security experts to properly configure everything from scratch. They read about Google's "Shared Fate" model.
**How is "Shared Fate" fundamentally different from the traditional "Shared Responsibility" model in a way that helps this startup?**

<details>
<summary>View Answer</summary>

*   **Traditional Model:** Draws a line. "This is your side, this is my side. Good luck." (Passive)
*   **Shared Fate Model:** "We are in this together." (Active)
*   **Difference:** Google proactively helps you secure *your* side of the line by providing **Secure Blueprints** (pre-made safe templates) and **Guardrails** (automatic warnings/blocks for bad configs). It changes the dynamic from "Provider vs. Customer" to "Partners."

</details>

**Quick Review of Key Identity Principles Tested:**
*   **Cross-Account:** Use roles/impersonation, not duplicate users.
*   **Context:** Identity is more than just a password; verify location and device (Conditional Access).
*   **Hierarchy:** Be careful where you apply permissions; they often trickle down.

---

## Next Up

In Lesson 2, we'll dive into **Cloud Network Security** — VPC design, security groups, micro-segmentation, and building defense-in-depth at the network layer!
