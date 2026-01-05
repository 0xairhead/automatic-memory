# Module 6, Lesson 1: Secure SDLC & Security Requirements

## Table of Contents
- [Introduction](#introduction)
- [Media Resources](#media-resources)
- [Part 1: Security in the Software Development Lifecycle](#part-1-security-in-the-software-development-lifecycle)
- [Part 2: Security Requirements Engineering](#part-2-security-requirements-engineering)
- [Part 3: Secure Design Principles](#part-3-secure-design-principles)
- [Part 4: Security Architecture Review](#part-4-security-architecture-review)
- [Part 5: Building a Security Champions Program](#part-5-building-a-security-champions-program)
- [Practice Questions](#practice-questions)
- [Summary](#summary)

---

## Introduction

Application security starts long before the first line of code is written. The most cost-effective time to address security is during requirements and design - finding and fixing a vulnerability in production costs 30-100x more than catching it during design. This lesson covers how to integrate security throughout the software development lifecycle, from gathering security requirements to conducting architecture reviews.

---

## Media Resources

### Recommended Videos
- "Secure Software Development Lifecycle" - OWASP Foundation
- "Threat Modeling: Designing for Security" - Microsoft Security
- "Building Security Champions" - SANS AppSec
- "Security Requirements with Abuse Cases" - SafeCode

### Recommended Reading
- "Threat Modeling: Designing for Security" - Adam Shostack
- OWASP Software Assurance Maturity Model (SAMM)
- NIST SP 800-160: Systems Security Engineering
- Microsoft Security Development Lifecycle (SDL)

---

## Part 1: Security in the Software Development Lifecycle

### The Assembly Line Analogy

Think of software development like manufacturing a car:

| Manufacturing Stage | SDLC Stage | Security Activity |
|---------------------|------------|-------------------|
| Design blueprints | Requirements | Security requirements, abuse cases |
| Engineering specs | Design | Threat modeling, secure architecture |
| Parts fabrication | Development | Secure coding, code review |
| Quality inspection | Testing | Security testing (SAST, DAST, pentest) |
| Assembly line | Deployment | Secure configuration, hardening |
| Dealer service | Operations | Monitoring, patching, incident response |

**Key Insight:** You wouldn't wait until the car is assembled to check if the brakes work. Similarly, you shouldn't wait until deployment to think about security.

### The Cost of Late Detection

```
Relative Cost to Fix Security Vulnerabilities:
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  Phase              Relative Cost    Example                │
│  ─────────────────────────────────────────────────────      │
│  Requirements          1x           Add auth requirement    │
│  Design                5x           Redesign data flow      │
│  Development          10x           Refactor auth module    │
│  Testing              20x           Fix and retest          │
│  Production           30x           Hotfix + incident       │
│  Post-breach         100x+          Breach response + legal │
│                                                             │
│  Visualization:                                             │
│                                                             │
│  Cost │                                           ████      │
│       │                                     ████  ████      │
│       │                               ████  ████  ████      │
│       │                         ████  ████  ████  ████      │
│       │                   ████  ████  ████  ████  ████      │
│       │             ████  ████  ████  ████  ████  ████      │
│       │       ████  ████  ████  ████  ████  ████  ████      │
│       │  ██   ████  ████  ████  ████  ████  ████  ████      │
│       └──────────────────────────────────────────────►      │
│         Req  Design  Dev  Test  Prod  Breach                │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### Secure SDLC Models

```
Microsoft Security Development Lifecycle (SDL):
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  Training ──► Requirements ──► Design ──► Implementation    │
│     │              │            │              │            │
│     │         Security      Threat         Secure           │
│     │         Requirements  Modeling       Coding           │
│     │         Abuse Cases   Attack         Static           │
│     │                       Surface        Analysis         │
│     │                       Review                          │
│     │                                                       │
│     │              │            │              │            │
│     └──────────────┼────────────┼──────────────┼───────────►│
│                    ▼            ▼              ▼            │
│              Verification ──► Release ──► Response          │
│                    │            │              │            │
│               Dynamic       Final         Incident          │
│               Testing       Security      Response          │
│               Fuzz          Review        Plan              │
│               Testing       Pen Test      Monitoring        │
│                                                             │
└─────────────────────────────────────────────────────────────┘

OWASP SAMM (Software Assurance Maturity Model):
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  Business Function    Security Practices                    │
│  ─────────────────────────────────────────────────────      │
│  Governance           • Strategy & Metrics                  │
│                       • Policy & Compliance                 │
│                       • Education & Guidance                │
│                                                             │
│  Design               • Threat Assessment                   │
│                       • Security Requirements               │
│                       • Security Architecture               │
│                                                             │
│  Implementation       • Secure Build                        │
│                       • Secure Deployment                   │
│                       • Defect Management                   │
│                                                             │
│  Verification         • Architecture Assessment             │
│                       • Requirements-driven Testing         │
│                       • Security Testing                    │
│                                                             │
│  Operations           • Incident Management                 │
│                       • Environment Management              │
│                       • Operational Management              │
│                                                             │
│  Maturity Levels: 0 (Implicit) → 1 → 2 → 3 (Measured)      │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### Security Gates in SDLC

```
Security Quality Gates:
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  Gate 1: Requirements Sign-off                              │
│  ─────────────────────────────────────────────────────      │
│  ✓ Security requirements documented                         │
│  ✓ Abuse cases identified                                   │
│  ✓ Data classification completed                            │
│  ✓ Compliance requirements mapped                           │
│  ✓ Risk assessment completed                                │
│                                                             │
│  Gate 2: Design Approval                                    │
│  ─────────────────────────────────────────────────────      │
│  ✓ Threat model completed and reviewed                      │
│  ✓ Security architecture approved                           │
│  ✓ Authentication/authorization design reviewed             │
│  ✓ Data flow diagrams with trust boundaries                │
│  ✓ Third-party components evaluated                         │
│                                                             │
│  Gate 3: Code Complete                                      │
│  ─────────────────────────────────────────────────────      │
│  ✓ SAST scan completed, critical issues resolved            │
│  ✓ Security code review completed                           │
│  ✓ No hard-coded secrets                                    │
│  ✓ Dependency scan completed (SCA)                          │
│  ✓ Unit tests include security test cases                   │
│                                                             │
│  Gate 4: Release Approval                                   │
│  ─────────────────────────────────────────────────────      │
│  ✓ DAST/penetration testing completed                       │
│  ✓ All critical/high findings remediated                    │
│  ✓ Security documentation complete                          │
│  ✓ Incident response procedures in place                    │
│  ✓ Final security review sign-off                           │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

---

## Part 2: Security Requirements Engineering

### Functional vs Security Requirements

```
Requirement Types:
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  Functional Requirement        Security Requirement         │
│  ─────────────────────────    ─────────────────────────    │
│  "Users can log in"           "Failed login attempts are    │
│                                locked after 5 tries"        │
│                                                             │
│  "Users can upload files"     "File uploads are scanned     │
│                                for malware and limited      │
│                                to approved types"           │
│                                                             │
│  "System stores customer      "Customer data is encrypted   │
│   data"                        at rest with AES-256"        │
│                                                             │
│  "API returns user profile"   "API enforces authorization   │
│                                to prevent IDOR"             │
│                                                             │
│  "Password reset via email"   "Reset tokens expire in 15    │
│                                minutes and are single-use"  │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### Security Requirements Categories

```
OWASP ASVS Security Requirements Framework:
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  Category                 Example Requirements              │
│  ─────────────────────────────────────────────────────      │
│  V1: Architecture         Verify security architecture      │
│                           patterns are documented           │
│                                                             │
│  V2: Authentication       Verify passwords require          │
│                           minimum 12 characters             │
│                                                             │
│  V3: Session Management   Verify sessions timeout after     │
│                           30 minutes of inactivity          │
│                                                             │
│  V4: Access Control       Verify principle of least         │
│                           privilege is enforced             │
│                                                             │
│  V5: Validation           Verify all input is validated     │
│                           against a whitelist               │
│                                                             │
│  V6: Cryptography         Verify only approved algorithms   │
│                           are used (no MD5, SHA1)           │
│                                                             │
│  V7: Error Handling       Verify errors don't leak          │
│                           sensitive information             │
│                                                             │
│  V8: Data Protection      Verify sensitive data is          │
│                           classified and protected          │
│                                                             │
│  V9: Communications       Verify TLS 1.2+ for all           │
│                           external communications           │
│                                                             │
│  V10: Malicious Code      Verify no backdoors or            │
│                           malicious functions               │
│                                                             │
│  V11: Business Logic      Verify business logic cannot      │
│                           be bypassed                       │
│                                                             │
│  V12: Files & Resources   Verify file uploads are           │
│                           restricted and validated          │
│                                                             │
│  V13: API & Web Service   Verify APIs require               │
│                           authentication                    │
│                                                             │
│  V14: Configuration       Verify security headers           │
│                           are properly configured           │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### Abuse Cases and Misuse Cases

```
From Use Case to Abuse Case:
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  Use Case: "User logs into account"                         │
│                                                             │
│  Abuse Cases:                                               │
│  ┌─────────────────────────────────────────────────────┐   │
│  │ Attacker brute-forces password                       │   │
│  │ → Mitigate: Rate limiting, account lockout           │   │
│  ├─────────────────────────────────────────────────────┤   │
│  │ Attacker uses credential stuffing with leaked DB     │   │
│  │ → Mitigate: Breach detection, MFA, password history  │   │
│  ├─────────────────────────────────────────────────────┤   │
│  │ Attacker steals session token                        │   │
│  │ → Mitigate: HTTPOnly cookies, token binding          │   │
│  ├─────────────────────────────────────────────────────┤   │
│  │ Attacker performs session fixation                   │   │
│  │ → Mitigate: Regenerate session ID after login        │   │
│  ├─────────────────────────────────────────────────────┤   │
│  │ Attacker intercepts credentials over network         │   │
│  │ → Mitigate: TLS everywhere, HSTS                     │   │
│  └─────────────────────────────────────────────────────┘   │
│                                                             │
│  Use Case: "User uploads profile picture"                   │
│                                                             │
│  Abuse Cases:                                               │
│  ┌─────────────────────────────────────────────────────┐   │
│  │ Attacker uploads malware disguised as image          │   │
│  │ → Mitigate: Validate magic bytes, AV scan            │   │
│  ├─────────────────────────────────────────────────────┤   │
│  │ Attacker uploads web shell (.php, .jsp)              │   │
│  │ → Mitigate: Whitelist extensions, no execute perms   │   │
│  ├─────────────────────────────────────────────────────┤   │
│  │ Attacker uploads very large file (DoS)               │   │
│  │ → Mitigate: File size limits, quota enforcement      │   │
│  ├─────────────────────────────────────────────────────┤   │
│  │ Attacker uploads image with XSS in metadata          │   │
│  │ → Mitigate: Strip metadata, sanitize filenames       │   │
│  ├─────────────────────────────────────────────────────┤   │
│  │ Attacker uses path traversal in filename             │   │
│  │ → Mitigate: Generate random filenames, no user input │   │
│  └─────────────────────────────────────────────────────┘   │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### Security Requirements Template

```markdown
## Security Requirement Template

**ID:** SEC-REQ-001
**Title:** Authentication Rate Limiting
**Category:** Authentication (ASVS V2)
**Priority:** High
**Compliance:** PCI-DSS 8.1.6, NIST 800-53 AC-7

### Description
The system shall limit failed authentication attempts to prevent
brute-force attacks.

### Acceptance Criteria
1. After 5 failed attempts, account is locked for 30 minutes
2. Failed attempts are logged with timestamp and source IP
3. Account lockout notification is sent to user
4. Lockout counter resets after successful authentication
5. Administrative unlock capability exists

### Abuse Cases Addressed
- Brute-force password attacks
- Credential stuffing attacks
- Automated login enumeration

### Verification Method
- Unit tests for lockout logic
- Integration tests for notification
- Penetration testing validation

### Related Requirements
- SEC-REQ-002: Strong password policy
- SEC-REQ-003: Multi-factor authentication
```

---

## Part 3: Secure Design Principles

### Core Security Design Principles

```
Fundamental Security Principles:
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  1. DEFENSE IN DEPTH                                        │
│  ─────────────────────────────────────────────────────      │
│  Multiple layers of security controls                       │
│                                                             │
│  ┌─────────────────────────────────────────────┐           │
│  │ ┌─────────────────────────────────────────┐ │           │
│  │ │ ┌─────────────────────────────────────┐ │ │           │
│  │ │ │ ┌─────────────────────────────────┐ │ │ │           │
│  │ │ │ │         Application             │ │ │ │           │
│  │ │ │ │           Data                  │ │ │ │           │
│  │ │ │ └─────────────────────────────────┘ │ │ │           │
│  │ │ │       Application Controls          │ │ │           │
│  │ │ └─────────────────────────────────────┘ │ │           │
│  │ │         Network Controls                │ │           │
│  │ └─────────────────────────────────────────┘ │           │
│  │           Perimeter Controls                │           │
│  └─────────────────────────────────────────────┘           │
│                                                             │
│  2. LEAST PRIVILEGE                                         │
│  ─────────────────────────────────────────────────────      │
│  Grant minimum permissions necessary                        │
│                                                             │
│  ✗ Bad:  Application runs as root                          │
│  ✓ Good: Application runs as dedicated service account      │
│          with only required file/network permissions        │
│                                                             │
│  3. FAIL SECURE (FAIL CLOSED)                              │
│  ─────────────────────────────────────────────────────      │
│  When errors occur, default to secure state                │
│                                                             │
│  ✗ Bad:  if (authCheck() == ERROR) { grantAccess(); }     │
│  ✓ Good: if (authCheck() == SUCCESS) { grantAccess(); }   │
│          else { denyAccess(); logFailure(); }              │
│                                                             │
│  4. SEPARATION OF DUTIES                                    │
│  ─────────────────────────────────────────────────────      │
│  No single user should control all aspects                 │
│                                                             │
│  ✗ Bad:  Developer deploys own code to production          │
│  ✓ Good: Developer → Code Review → QA → Ops deploys        │
│                                                             │
│  5. ECONOMY OF MECHANISM (KISS)                            │
│  ─────────────────────────────────────────────────────      │
│  Keep security mechanisms simple and verifiable            │
│                                                             │
│  ✗ Bad:  Custom encryption algorithm                       │
│  ✓ Good: Standard AES-256-GCM from vetted library          │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### Additional Design Principles

```
More Security Design Principles:
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  6. COMPLETE MEDIATION                                      │
│  ─────────────────────────────────────────────────────      │
│  Check authorization on every access                        │
│                                                             │
│  ✗ Bad:  Check auth once, cache forever                    │
│  ✓ Good: Verify authorization on each request              │
│                                                             │
│  7. OPEN DESIGN                                             │
│  ─────────────────────────────────────────────────────      │
│  Security shouldn't depend on obscurity                    │
│                                                             │
│  ✗ Bad:  "Attackers won't find our admin page at /x7k2"   │
│  ✓ Good: Admin page protected by authentication + authz    │
│                                                             │
│  8. PSYCHOLOGICAL ACCEPTABILITY                             │
│  ─────────────────────────────────────────────────────      │
│  Security shouldn't make the system unusable               │
│                                                             │
│  ✗ Bad:  30-character password requirement                 │
│  ✓ Good: Passphrase support + MFA option                   │
│                                                             │
│  9. MINIMIZE ATTACK SURFACE                                 │
│  ─────────────────────────────────────────────────────      │
│  Reduce entry points and exposure                          │
│                                                             │
│  ✗ Bad:  All ports open, debug endpoints in prod           │
│  ✓ Good: Only required ports, no debug in prod             │
│                                                             │
│  10. SECURE DEFAULTS                                        │
│  ─────────────────────────────────────────────────────      │
│  Out-of-the-box configuration should be secure             │
│                                                             │
│  ✗ Bad:  Default admin password is "admin"                 │
│  ✓ Good: Force password change on first login              │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### Threat Modeling Integration

```
Threat Modeling in Design Phase:
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  STRIDE Threat Model:                                       │
│                                                             │
│  Threat              Security Property   Design Response    │
│  ─────────────────────────────────────────────────────      │
│  Spoofing            Authentication      MFA, certificates  │
│  Tampering           Integrity           Signatures, MACs   │
│  Repudiation         Non-repudiation     Audit logs, sigs   │
│  Info Disclosure     Confidentiality     Encryption, ACLs   │
│  Denial of Service   Availability        Rate limiting      │
│  Elevation of Priv   Authorization       RBAC, least priv   │
│                                                             │
│  Data Flow Diagram with Trust Boundaries:                   │
│                                                             │
│  ┌─────────────────────────────────────────────────────┐   │
│  │ External Zone (Untrusted)                            │   │
│  │                                                      │   │
│  │    ┌──────┐                                         │   │
│  │    │ User │                                         │   │
│  │    └──┬───┘                                         │   │
│  │       │ HTTPS                                       │   │
│  ├───────┼─────────────────────────────────────────────┤   │
│  │ DMZ   │ Trust Boundary 1                            │   │
│  │       ▼                                             │   │
│  │    ┌──────┐         ┌─────┐                        │   │
│  │    │ WAF  │────────►│ LB  │                        │   │
│  │    └──────┘         └──┬──┘                        │   │
│  │                        │                            │   │
│  ├────────────────────────┼────────────────────────────┤   │
│  │ App Zone               │ Trust Boundary 2           │   │
│  │                        ▼                            │   │
│  │                   ┌─────────┐                       │   │
│  │                   │ Web App │                       │   │
│  │                   └────┬────┘                       │   │
│  │                        │                            │   │
│  ├────────────────────────┼────────────────────────────┤   │
│  │ Data Zone              │ Trust Boundary 3           │   │
│  │                        ▼                            │   │
│  │                   ┌─────────┐                       │   │
│  │                   │   DB    │                       │   │
│  │                   └─────────┘                       │   │
│  └─────────────────────────────────────────────────────┘   │
│                                                             │
│  Key Design Questions at Each Boundary:                     │
│  • What data crosses this boundary?                        │
│  • How is data validated/sanitized?                        │
│  • How is the connection authenticated?                    │
│  • What happens if this component is compromised?          │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

---

## Part 4: Security Architecture Review

### Architecture Review Process

```
Security Architecture Review Lifecycle:
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  Step 1: INFORMATION GATHERING                              │
│  ─────────────────────────────────────────────────────      │
│  □ Architecture diagrams (network, data flow, deployment)   │
│  □ Technology stack documentation                           │
│  □ Security requirements and constraints                    │
│  □ Compliance requirements (PCI, HIPAA, SOC 2)             │
│  □ Threat model (if available)                             │
│  □ Previous security assessments                           │
│                                                             │
│  Step 2: REVIEW SCOPE DEFINITION                            │
│  ─────────────────────────────────────────────────────      │
│  □ Define system boundaries                                 │
│  □ Identify crown jewel assets                             │
│  □ Determine trust boundaries                              │
│  □ List external dependencies                              │
│  □ Agree on review criteria and depth                      │
│                                                             │
│  Step 3: ARCHITECTURE ANALYSIS                              │
│  ─────────────────────────────────────────────────────      │
│  □ Authentication & authorization review                    │
│  □ Data protection assessment                              │
│  □ Network security architecture                           │
│  □ Third-party integration security                        │
│  □ Logging and monitoring capabilities                     │
│  □ Resilience and availability design                      │
│                                                             │
│  Step 4: THREAT MODELING                                    │
│  ─────────────────────────────────────────────────────      │
│  □ Identify threat actors and motivations                  │
│  □ Map attack vectors to entry points                      │
│  □ Analyze each component for STRIDE threats               │
│  □ Assess existing controls                                │
│  □ Identify gaps and weaknesses                            │
│                                                             │
│  Step 5: FINDINGS AND RECOMMENDATIONS                       │
│  ─────────────────────────────────────────────────────      │
│  □ Document findings with severity ratings                 │
│  □ Provide specific remediation recommendations            │
│  □ Prioritize based on risk                                │
│  □ Present to stakeholders                                 │
│  □ Track remediation progress                              │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### Architecture Review Checklist

```
Security Architecture Review Checklist:
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  AUTHENTICATION                                             │
│  ─────────────────────────────────────────────────────      │
│  □ Strong authentication mechanism (MFA available)          │
│  □ Secure credential storage (bcrypt, Argon2)              │
│  □ Session management properly implemented                  │
│  □ Password policy enforced                                │
│  □ Account lockout implemented                             │
│  □ Secure password reset flow                              │
│                                                             │
│  AUTHORIZATION                                              │
│  ─────────────────────────────────────────────────────      │
│  □ Role-based or attribute-based access control            │
│  □ Principle of least privilege applied                    │
│  □ Authorization checked on every request                  │
│  □ Direct object reference protection (IDOR)               │
│  □ Administrative functions properly protected             │
│  □ API authorization implemented                           │
│                                                             │
│  DATA PROTECTION                                            │
│  ─────────────────────────────────────────────────────      │
│  □ Data classification scheme defined                      │
│  □ Encryption at rest for sensitive data                   │
│  □ Encryption in transit (TLS 1.2+)                        │
│  □ Key management strategy defined                         │
│  □ PII/PHI handling compliant with regulations             │
│  □ Data retention and disposal policies                    │
│                                                             │
│  INPUT VALIDATION                                           │
│  ─────────────────────────────────────────────────────      │
│  □ All input validated server-side                         │
│  □ Parameterized queries for database access               │
│  □ Output encoding implemented                             │
│  □ File upload restrictions in place                       │
│  □ Content-Type validation                                 │
│                                                             │
│  LOGGING & MONITORING                                       │
│  ─────────────────────────────────────────────────────      │
│  □ Security events logged                                  │
│  □ Logs protected from tampering                           │
│  □ Sensitive data not logged                               │
│  □ Centralized log aggregation                             │
│  □ Alerting on security events                             │
│  □ Audit trail for compliance                              │
│                                                             │
│  ERROR HANDLING                                             │
│  ─────────────────────────────────────────────────────      │
│  □ Generic error messages to users                         │
│  □ Detailed errors logged securely                         │
│  □ Fail-secure behavior implemented                        │
│  □ Exception handling doesn't leak info                    │
│                                                             │
│  CONFIGURATION                                              │
│  ─────────────────────────────────────────────────────      │
│  □ Security headers implemented                            │
│  □ Debug mode disabled in production                       │
│  □ Default credentials changed                             │
│  □ Unnecessary features disabled                           │
│  □ Secrets externalized from code                          │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### Architecture Review Finding Template

## Security Architecture Finding

**Finding ID:** SAR-2024-001
**Title:** Insecure Direct Object Reference in Document API
**Severity:** High
**CVSS Score:** 7.5 (AV:N/AC:L/PR:L/UI:N/S:U/C:H/I:N/A:N)

### Description
The document API endpoint `/api/documents/{id}` does not verify
that the authenticated user has authorization to access the
requested document. Any authenticated user can access any
document by iterating through document IDs.

### Affected Components
- Document Service API
- `/api/documents/{id}` endpoint
- Mobile application document viewer

### Attack Scenario
1. Attacker authenticates as normal user
2. Attacker requests `/api/documents/1`
3. API returns document belonging to another user
4. Attacker iterates through IDs to harvest all documents

### Business Impact
- Confidential documents exposed to unauthorized users
- Potential regulatory violations (GDPR, HIPAA)
- Reputational damage if exploited

### Recommendation
1. Implement authorization check on document access:
   ```python
   def get_document(document_id, user):
       doc = Document.get(document_id)
       if doc.owner_id != user.id and not user.has_role('admin'):
           raise AuthorizationError("Access denied")
       return doc
   ```
2. Use unpredictable document identifiers (UUIDs)
3. Add authorization integration tests
4. Enable API access logging for forensics

### References
- OWASP IDOR: https://owasp.org/www-project-web-security-testing-guide/
- CWE-639: Authorization Bypass Through User-Controlled Key

---

## Part 5: Building a Security Champions Program

### The Security Champions Model

```
Security Champions Program Structure:
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│                    ┌──────────────────┐                    │
│                    │  Security Team   │                    │
│                    │   (Central)      │                    │
│                    └────────┬─────────┘                    │
│                             │                              │
│              ┌──────────────┼──────────────┐              │
│              │              │              │              │
│              ▼              ▼              ▼              │
│        ┌─────────┐    ┌─────────┐    ┌─────────┐        │
│        │Champion │    │Champion │    │Champion │        │
│        │ Team A  │    │ Team B  │    │ Team C  │        │
│        └────┬────┘    └────┬────┘    └────┬────┘        │
│             │              │              │              │
│        ┌────┴────┐    ┌────┴────┐    ┌────┴────┐        │
│        │ Dev Team│    │ Dev Team│    │ Dev Team│        │
│        │    A    │    │    B    │    │    C    │        │
│        │ (8 devs)│    │(12 devs)│    │ (6 devs)│        │
│        └─────────┘    └─────────┘    └─────────┘        │
│                                                             │
│  Ratio: ~1 Champion per 8-15 developers                    │
│                                                             │
└─────────────────────────────────────────────────────────────┘

Champion Responsibilities:
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  Primary (within their team):                              │
│  • First point of contact for security questions            │
│  • Conduct initial security code reviews                    │
│  • Advocate for security requirements                       │
│  • Triage security tool findings                           │
│  • Participate in threat modeling sessions                  │
│                                                             │
│  Secondary (with security team):                           │
│  • Attend monthly security champion meetings                │
│  • Complete advanced security training                      │
│  • Share knowledge back to development team                 │
│  • Provide feedback on security tools/processes             │
│  • Escalate complex security issues                        │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### Champion Selection and Training

```
Security Champion Profile:
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  Ideal Characteristics:                                     │
│  ─────────────────────────────────────────────────────      │
│  ✓ Respected by peers (technical credibility)               │
│  ✓ Interest in security (volunteer, not voluntold)          │
│  ✓ Strong communication skills                              │
│  ✓ 2+ years with the organization                          │
│  ✓ Good understanding of team's technology stack            │
│  ✓ Time allocated (10-20% of work time)                    │
│                                                             │
│  Training Curriculum:                                       │
│  ─────────────────────────────────────────────────────      │
│  Month 1: Foundations                                       │
│  • OWASP Top 10 deep dive                                  │
│  • Secure coding basics                                    │
│  • Security tools overview (SAST, DAST, SCA)               │
│                                                             │
│  Month 2: Intermediate                                      │
│  • Threat modeling workshop                                │
│  • Security code review techniques                         │
│  • Common vulnerability patterns                           │
│                                                             │
│  Month 3: Advanced                                          │
│  • Security architecture principles                        │
│  • Advanced attack techniques                              │
│  • Incident response basics                                │
│                                                             │
│  Ongoing:                                                   │
│  • Monthly security team meetings                          │
│  • Quarterly security training updates                     │
│  • Annual security conference attendance                   │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### Measuring Champion Program Success

```
Security Champions Program KPIs:
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  Input Metrics (Activity):                                  │
│  ─────────────────────────────────────────────────────      │
│  • Number of trained champions: 12/15 teams covered         │
│  • Training completion rate: 95%                           │
│  • Champion meeting attendance: 85%                        │
│  • Security consultations handled by champions: 234/month  │
│                                                             │
│  Output Metrics (Results):                                  │
│  ─────────────────────────────────────────────────────      │
│  • Vulnerabilities found in design review: +40% YoY        │
│  • Time to remediate vulnerabilities: -35%                 │
│  • Security tickets escalated to security team: -50%       │
│  • Developer security survey scores: 4.2/5.0              │
│                                                             │
│  Outcome Metrics (Impact):                                  │
│  ─────────────────────────────────────────────────────      │
│  • Production vulnerabilities: -60% YoY                    │
│  • Security-related release delays: -45%                   │
│  • Mean time to fix critical vulns: 5 days → 2 days        │
│  • Security incidents from code defects: -70%              │
│                                                             │
│  Champion Satisfaction:                                     │
│  ─────────────────────────────────────────────────────      │
│  • Champion retention rate: 85%                            │
│  • NPS from champions: +45                                 │
│  • Career advancement rate: 30% promoted within 2 years    │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

---

## Practice Questions

### Question 1
You're joining a company that has no formal secure SDLC. Development teams ship code without security review, and there's been a recent string of production vulnerabilities. How would you prioritize implementing security in the SDLC?

<details>
<summary>Show Answer</summary>

**Phased approach to implementing secure SDLC:**

**Phase 1: Quick Wins (0-3 months)**
1. Deploy automated SAST in CI/CD pipeline with blocking on critical issues
2. Implement Software Composition Analysis (SCA) for known vulnerabilities
3. Add basic security requirements to existing templates
4. Establish security incident process for production issues
5. Start tracking security metrics (vulnerabilities found, time to fix)

**Phase 2: Foundation Building (3-6 months)**
1. Launch security champions program (identify 2-3 volunteers)
2. Implement security gates at key milestones
3. Begin threat modeling for high-risk applications
4. Develop secure coding guidelines for primary languages
5. Add DAST scanning for web applications

**Phase 3: Maturation (6-12 months)**
1. Integrate security requirements into product backlog process
2. Implement full architecture review process
3. Expand security champions to all teams
4. Establish bug bounty or vulnerability disclosure program
5. Advanced training for development teams

**Key success factors:**
- Get executive sponsorship and budget
- Start with carrots before sticks (enable, don't block)
- Measure and communicate improvements
- Build relationships with development teams
- Be pragmatic about what's achievable

</details>

### Question 2
A development team complains that security requirements are too vague ("the application must be secure") to implement. How would you help them create actionable security requirements?

<details>
<summary>Show Answer</summary>

**Approach to creating actionable security requirements:**

1. **Replace vague with specific:**
   - Vague: "Application must be secure"
   - Specific: "Authentication must use bcrypt with cost factor 12 minimum"
   - Specific: "Session tokens must expire after 30 minutes of inactivity"
   - Specific: "All user input must be validated against defined schemas"

2. **Use established frameworks:**
   - Map to OWASP ASVS requirements by category
   - Reference CIS benchmarks for infrastructure
   - Align with compliance requirements (PCI-DSS, HIPAA)

3. **Create abuse cases for each user story:**
   - For every "user can do X," ask "how could attacker abuse X?"
   - Document specific mitigations as requirements

4. **Provide templates with examples:**
   ```
   Security Requirement Template:
   - ID: SEC-001
   - Category: Authentication
   - Requirement: Failed login attempts locked after 5 tries
   - Acceptance Criteria:
     * Counter increments on failure
     * Account locked for 30 min after 5 failures
     * Lockout notification sent to user
   - Test Method: Unit test + integration test
   ```

5. **Include acceptance criteria:**
   - Testable conditions that prove the requirement is met
   - Link to specific test cases

6. **Hold requirements workshops:**
   - Work directly with team on first few features
   - Build their capability to derive requirements themselves

</details>

### Question 3
During an architecture review, you discover the application stores API keys in a config file committed to the git repository. The team says it's necessary because the application needs the keys at runtime. How do you address this?

<details>
<summary>Show Answer</summary>

**Root cause:** Developers conflate "application needs secrets at runtime" with "secrets must be in code repository."

**Remediation approach:**

1. **Immediate mitigation:**
   - Rotate all exposed API keys immediately
   - Remove config file from repository
   - Add pattern to .gitignore
   - Use git-filter-branch or BFG to remove from history

2. **Implement secrets management:**

   **Option A: Environment variables**
   ```bash
   # Injected at runtime, not in code
   export API_KEY=actual_secret_value
   ./start_application
   ```

   **Option B: Secrets manager**
   ```python
   # Application retrieves at startup
   import boto3
   client = boto3.client('secretsmanager')
   secret = client.get_secret_value(SecretId='api-key')
   ```

   **Option C: HashiCorp Vault**
   ```python
   import hvac
   client = hvac.Client(url='https://vault:8200')
   secret = client.secrets.kv.read_secret_version(path='api-key')
   ```

3. **CI/CD integration:**
   - Secrets injected by pipeline, not stored in repo
   - Different secrets per environment (dev/staging/prod)
   - Audit trail of secret access

4. **Prevention controls:**
   - Pre-commit hooks to scan for secrets (git-secrets, trufflehog)
   - CI pipeline scanning for exposed secrets
   - Regular repository scanning

5. **Education:**
   - Train team on why this matters (credential theft, breach)
   - Show examples of real breaches from exposed secrets
   - Document approved patterns for secrets management

</details>

### Question 4
How would you design a security architecture review process that scales across 20 development teams with limited security staff (3 people)?

<details>
<summary>Show Answer</summary>

**Scalable architecture review process:**

1. **Tiered review approach:**

   | Tier | Criteria | Review Type | Reviewer |
   |------|----------|-------------|----------|
   | 1 | Low risk, no sensitive data | Self-assessment checklist | Dev team |
   | 2 | Medium risk, internal data | Security champion review | Champion |
   | 3 | High risk, PII/financial | Full security team review | Security |
   | 4 | Critical, external-facing | Deep dive + pentest | Security + external |

2. **Risk classification criteria:**
   ```
   High/Critical triggers:
   - Handles PII, PHI, or financial data
   - External/internet-facing
   - Authentication/authorization changes
   - Cryptographic implementation
   - New third-party integrations
   - Compliance scope (PCI, HIPAA)
   ```

3. **Self-service tools:**
   - Architecture review questionnaire (risk classification)
   - Automated threat modeling tool (OWASP Threat Dragon)
   - Security checklist by application type
   - Pre-approved architecture patterns library

4. **Security champions enablement:**
   - Train champions to conduct Tier 2 reviews
   - Provide review templates and criteria
   - Office hours with security team for guidance
   - Escalation path for complex issues

5. **Time-boxing:**
   - Tier 2: 2-4 hours max
   - Tier 3: 1-2 days max
   - Schedule reviews in advance, not ad-hoc

6. **Continuous improvement:**
   - Track common findings, build automation
   - Update checklists based on real vulnerabilities
   - Measure review bottleneck impact

This approach lets 3 security staff support 20 teams by handling only ~20% of reviews directly while ensuring all applications receive appropriate scrutiny.

</details>

### Question 5
A product manager argues that adding security requirements will slow down development and delay the product launch. How do you respond?

<details>
<summary>Show Answer</summary>

**Business-aligned response:**

1. **Acknowledge the concern:**
   "I understand launch timing is critical. Let's find the right balance."

2. **Quantify the cost of NOT doing security:**
   - Average data breach cost: $4.45M (2023 IBM report)
   - Breach causes average 23-day business disruption
   - Post-breach security remediation is 30-100x more expensive
   - Example: Equifax spent $1.4B+ on breach remediation

3. **Show that security doesn't have to slow things down:**
   - Security integrated in sprint work, not separate waterfall phase
   - Automated testing catches issues without manual overhead
   - Security requirements clarify scope, reducing rework
   - Companies with mature DevSecOps deploy MORE frequently

4. **Offer risk-based prioritization:**
   "Let's focus on critical security requirements for MVP:
   - Authentication and authorization (non-negotiable)
   - Data protection for sensitive fields
   - Input validation on external interfaces

   We can address medium/low items in fast-follows."

5. **Propose parallel tracks:**
   - Development continues while security reviews happen
   - Security testing runs in CI/CD automatically
   - Threat modeling done during design, not blocking code

6. **Frame as business enabler:**
   - Security certifications (SOC 2) unlock enterprise customers
   - Security features are competitive differentiators
   - Customer trust is a business asset

7. **Get executive alignment:**
   - Present risk acceptance decision to leadership
   - Document the business decision and who approved
   - Ensure product manager understands they own the risk

</details>

---

## Summary

Secure SDLC integration is about embedding security into existing development processes, not creating separate security processes:

### Key Principles
- **Shift left:** Find and fix issues early when they're cheapest to address
- **Automate:** Use tools to scale security without blocking development
- **Enable:** Provide developers with knowledge, tools, and support
- **Measure:** Track metrics to demonstrate improvement and identify gaps

### Security Requirements
- Derive from **abuse cases** and **misuse cases**
- Use frameworks like **OWASP ASVS** for comprehensive coverage
- Make requirements **specific, testable, and actionable**
- Include **acceptance criteria** and **verification methods**

### Secure Design
- Apply fundamental principles: **defense in depth, least privilege, fail secure**
- Integrate **threat modeling** into design reviews
- Use **trust boundaries** to identify where security controls are needed
- Document security architecture decisions

### Architecture Reviews
- Scale through **tiered reviews** and **risk classification**
- Leverage **security champions** to extend coverage
- Use **checklists** and **templates** for consistency
- Track and remediate findings systematically

### Security Champions
- Build a network of security-minded developers
- Provide **training, time, and recognition**
- Measure program effectiveness through **outcomes, not just activities**

---

## Next Up

**Module 6, Lesson 2: OWASP Top 10 & Web Application Vulnerabilities** - Deep dive into the most critical web application security risks and their mitigations.
