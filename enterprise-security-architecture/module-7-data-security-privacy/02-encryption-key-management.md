# Lesson 2: Encryption & Key Management

## Table of Contents
- [Introduction](#introduction)
- [Encryption Fundamentals](#encryption-fundamentals)
- [Encryption at Rest](#encryption-at-rest)
- [Encryption in Transit](#encryption-in-transit)
- [Encryption in Use](#encryption-in-use)
- [Key Management Architecture](#key-management-architecture)
- [Cloud KMS Services](#cloud-kms-services)
- [Hardware Security Modules](#hardware-security-modules)
- [Key Lifecycle Management](#key-lifecycle-management)
- [Interview Practice Questions](#interview-practice-questions)

## Media Resources

### Recommended Videos
- **Computerphile** - "AES Explained" and "Public Key Cryptography"
- **AWS** - "AWS KMS Deep Dive"
- **Microsoft** - "Azure Key Vault Architecture"
- **HashiCorp** - "Vault Encryption as a Service"

### Official Documentation
- NIST SP 800-57: Key Management Recommendations
- NIST SP 800-131A: Cryptographic Algorithm Transitions
- PCI DSS v4.0: Requirement 3 (Protect Stored Account Data)
- Cloud Security Alliance: Key Management Guide

---

## Introduction

Encryption is your last line of defense. Even if attackers breach all other controls, properly encrypted data remains protected. This lesson covers encryption strategies for data at rest, in transit, and in use, plus the critical infrastructure for managing encryption keys.

> **Key Principle:** Encryption without proper key management is security theater. The keys are the crown jewels.

---

## Encryption Fundamentals

### Symmetric vs Asymmetric Encryption

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                    ENCRYPTION TYPE COMPARISON                                │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  SYMMETRIC ENCRYPTION                    ASYMMETRIC ENCRYPTION              │
│  (Same key for encrypt/decrypt)          (Public/Private key pair)          │
│                                                                             │
│  ┌──────────────────────────────┐       ┌──────────────────────────────┐   │
│  │     ┌─────────┐              │       │  ┌─────────┐  ┌─────────┐    │   │
│  │     │  KEY    │              │       │  │ PUBLIC  │  │ PRIVATE │    │   │
│  │     │  (AES)  │              │       │  │  KEY    │  │  KEY    │    │   │
│  │     └────┬────┘              │       │  └────┬────┘  └────┬────┘    │   │
│  │          │                   │       │       │            │         │   │
│  │    ┌─────┴─────┐             │       │       ▼            ▼         │   │
│  │    ▼           ▼             │       │   Encrypt      Decrypt       │   │
│  │ Encrypt    Decrypt           │       │      OR           OR         │   │
│  │                              │       │   Sign         Verify        │   │
│  └──────────────────────────────┘       └──────────────────────────────┘   │
│                                                                             │
│  Algorithms:                             Algorithms:                        │
│  ├── AES-256 (standard)                  ├── RSA-2048/4096                  │
│  ├── AES-128 (acceptable)                ├── ECDSA (P-256, P-384)           │
│  ├── ChaCha20-Poly1305                   ├── Ed25519                        │
│  └── 3DES (legacy, avoid)                └── X25519 (key exchange)          │
│                                                                             │
│  Performance: FAST (GB/s)                Performance: SLOW (KB/s)           │
│  Key Exchange: Challenging               Key Exchange: Built-in             │
│  Use Case: Bulk data encryption          Use Case: Key exchange, signatures │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Hybrid Encryption (Real-World Pattern)

```
HYBRID ENCRYPTION PATTERN
=========================

Problem: Symmetric is fast but key exchange is hard
         Asymmetric handles key exchange but is slow

Solution: Use both!

┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│  SENDER                                              RECIPIENT              │
│                                                                             │
│  1. Generate random symmetric key (DEK)                                     │
│     ┌────────────────┐                                                      │
│     │ DEK: 256-bit   │                                                      │
│     │ random key     │                                                      │
│     └───────┬────────┘                                                      │
│             │                                                               │
│  2. Encrypt data with DEK (fast)                                            │
│     ┌────────────────┐     ┌────────────────┐                              │
│     │   Plaintext    │────→│   Ciphertext   │                              │
│     │   (large)      │ AES │   (large)      │                              │
│     └────────────────┘     └────────────────┘                              │
│                                                                             │
│  3. Encrypt DEK with recipient's public key (slow, but DEK is small)        │
│     ┌────────────────┐     ┌────────────────┐                              │
│     │      DEK       │────→│ Encrypted DEK  │                              │
│     │   (32 bytes)   │ RSA │   (256 bytes)  │                              │
│     └────────────────┘     └────────────────┘                              │
│                                                                             │
│  4. Send: Encrypted DEK + Ciphertext                                        │
│                                                                             │
│  ═══════════════════════════════════════════════════════════════════════   │
│                                                                             │
│  5. Recipient decrypts DEK with private key                                 │
│     ┌────────────────┐     ┌────────────────┐                              │
│     │ Encrypted DEK  │────→│      DEK       │                              │
│     └────────────────┘ RSA └────────────────┘                              │
│                                   │                                         │
│  6. Decrypt data with DEK         │                                         │
│     ┌────────────────┐     ┌──────┴─────────┐                              │
│     │   Ciphertext   │────→│   Plaintext    │                              │
│     └────────────────┘ AES └────────────────┘                              │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Cryptographic Algorithm Selection

```
ALGORITHM RECOMMENDATIONS (NIST SP 800-131A)
============================================

┌────────────────────────────────────────────────────────────────────────────┐
│ Purpose              │ Recommended              │ Avoid                    │
├──────────────────────┼──────────────────────────┼──────────────────────────┤
│ Symmetric Encryption │ AES-256-GCM              │ DES, 3DES, RC4, Blowfish │
│                      │ AES-128-GCM              │                          │
│                      │ ChaCha20-Poly1305        │                          │
├──────────────────────┼──────────────────────────┼──────────────────────────┤
│ Hashing              │ SHA-256, SHA-384, SHA-512│ MD5, SHA-1               │
│                      │ SHA-3                    │                          │
├──────────────────────┼──────────────────────────┼──────────────────────────┤
│ Key Exchange         │ ECDH (P-256, P-384)      │ DH < 2048-bit            │
│                      │ X25519                   │ RSA key transport        │
│                      │ DH ≥ 3072-bit            │                          │
├──────────────────────┼──────────────────────────┼──────────────────────────┤
│ Digital Signatures   │ ECDSA (P-256, P-384)     │ RSA < 2048-bit           │
│                      │ Ed25519                  │ DSA                      │
│                      │ RSA ≥ 2048-bit           │                          │
├──────────────────────┼──────────────────────────┼──────────────────────────┤
│ Password Hashing     │ Argon2id                 │ MD5, SHA-1, plain SHA-256│
│                      │ bcrypt                   │ without salting          │
│                      │ scrypt                   │                          │
└──────────────────────┴──────────────────────────┴──────────────────────────┘

MINIMUM KEY LENGTHS (2024+):
├── AES: 128-bit (acceptable), 256-bit (recommended)
├── RSA: 2048-bit (minimum), 3072-bit (recommended), 4096-bit (high security)
├── ECDSA/ECDH: P-256 (128-bit security), P-384 (192-bit security)
└── Ed25519/X25519: 256-bit (128-bit security)
```

---

## Encryption at Rest

### Storage Encryption Architecture

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                    ENCRYPTION AT REST LAYERS                                 │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  LAYER 1: FULL DISK ENCRYPTION (FDE)                                        │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  Encrypts entire disk/volume                                        │   │
│  │  Examples: BitLocker, LUKS, AWS EBS encryption                      │   │
│  │  Protection: Physical theft, improper disposal                      │   │
│  │  Limitation: Data decrypted when system running                     │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                              ▲                                              │
│  LAYER 2: FILE/FOLDER ENCRYPTION                                            │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  Encrypts individual files or folders                               │   │
│  │  Examples: EFS, VeraCrypt containers, S3 SSE                        │   │
│  │  Protection: Unauthorized access within system                      │   │
│  │  Limitation: Metadata (filenames) may be visible                    │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                              ▲                                              │
│  LAYER 3: FIELD/COLUMN ENCRYPTION                                           │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  Encrypts specific database columns or fields                       │   │
│  │  Examples: Always Encrypted, Application-level encryption           │   │
│  │  Protection: DBA access, SQL injection data extraction              │   │
│  │  Limitation: Performance impact on encrypted columns                │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                              ▲                                              │
│  LAYER 4: APPLICATION-LEVEL ENCRYPTION                                      │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  Application encrypts before storage                                │   │
│  │  Examples: Client-side encryption, E2EE                             │   │
│  │  Protection: Cloud provider, infrastructure admins                  │   │
│  │  Limitation: Application complexity, key management burden          │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Envelope Encryption Pattern

```
ENVELOPE ENCRYPTION (AWS KMS Pattern)
=====================================

┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│  KEY HIERARCHY                                                              │
│                                                                             │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  ROOT KEY (Customer Master Key - CMK)                               │   │
│  │  ├── Stored in HSM (never leaves)                                   │   │
│  │  ├── Used only to encrypt/decrypt Data Encryption Keys              │   │
│  │  └── AWS/Azure/GCP managed or Customer managed                      │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                              │                                              │
│                              │ Encrypts                                     │
│                              ▼                                              │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  DATA ENCRYPTION KEY (DEK)                                          │   │
│  │  ├── Generated per object/table/file                                │   │
│  │  ├── Used to encrypt actual data                                    │   │
│  │  └── Stored encrypted (wrapped) alongside data                      │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                              │                                              │
│                              │ Encrypts                                     │
│                              ▼                                              │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  DATA                                                               │   │
│  │  └── Your actual files, database records, etc.                      │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
│  ENCRYPTION FLOW:                                                           │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  1. Application requests new DEK from KMS                           │   │
│  │  2. KMS generates DEK, returns plaintext + encrypted DEK            │   │
│  │  3. Application encrypts data with plaintext DEK                    │   │
│  │  4. Application stores: Encrypted Data + Encrypted DEK              │   │
│  │  5. Application discards plaintext DEK (never stored)               │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
│  DECRYPTION FLOW:                                                           │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  1. Application retrieves: Encrypted Data + Encrypted DEK           │   │
│  │  2. Application sends Encrypted DEK to KMS                          │   │
│  │  3. KMS decrypts DEK, returns plaintext DEK                         │   │
│  │  4. Application decrypts data with plaintext DEK                    │   │
│  │  5. Application discards plaintext DEK                              │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
│  BENEFITS:                                                                  │
│  ├── CMK never leaves HSM (high security)                                  │
│  ├── Each object can have unique DEK (limits blast radius)                 │
│  ├── Re-encryption is key rotation (just re-wrap DEKs)                     │
│  └── Performance: Only small DEK needs HSM processing                      │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Database Encryption Options

```
DATABASE ENCRYPTION COMPARISON
==============================

┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│  TRANSPARENT DATA ENCRYPTION (TDE)                                          │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  How: Database encrypts data files at storage layer                 │   │
│  │  Where: SQL Server, Oracle, PostgreSQL, MySQL, Azure SQL            │   │
│  │                                                                     │   │
│  │  ┌─────────┐    ┌─────────┐    ┌─────────────────┐                 │   │
│  │  │  App    │───→│ DB Engine│───→│ Encrypted Files │                 │   │
│  │  │(Clear)  │    │(Clear)  │    │ (At Rest)       │                 │   │
│  │  └─────────┘    └─────────┘    └─────────────────┘                 │   │
│  │                                                                     │   │
│  │  Pro: No application changes, protects backups                      │   │
│  │  Con: DBA sees plaintext, data clear in memory                      │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
│  COLUMN-LEVEL ENCRYPTION                                                    │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  How: Specific columns encrypted, others clear                      │   │
│  │  Where: SQL Server, PostgreSQL (pgcrypto), application layer        │   │
│  │                                                                     │   │
│  │  ┌─────────────────────────────────────────────────────────────┐   │   │
│  │  │ ID │  Name  │      SSN (encrypted)     │  City   │          │   │   │
│  │  ├────┼────────┼──────────────────────────┼─────────┤          │   │   │
│  │  │ 1  │ Alice  │ AES256:x7f8g9h...       │ Boston  │          │   │   │
│  │  │ 2  │ Bob    │ AES256:a1b2c3d...       │ Denver  │          │   │   │
│  │  └─────────────────────────────────────────────────────────────┘   │   │
│  │                                                                     │   │
│  │  Pro: Protects specific sensitive fields, DBA can't see            │   │
│  │  Con: Can't search/sort encrypted columns, app changes needed      │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
│  ALWAYS ENCRYPTED (SQL Server/Azure SQL)                                    │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  How: Client-side encryption, DB never sees plaintext               │   │
│  │                                                                     │   │
│  │  ┌─────────┐    ┌─────────┐    ┌─────────────────┐                 │   │
│  │  │  App    │───→│ DB Engine│───→│ Storage         │                 │   │
│  │  │(Encrypt)│    │(Cipher) │    │ (Cipher)        │                 │   │
│  │  └─────────┘    └─────────┘    └─────────────────┘                 │   │
│  │      ▲                                                              │   │
│  │      │ Keys stored in Azure Key Vault or HSM                        │   │
│  │                                                                     │   │
│  │  Encryption Types:                                                  │   │
│  │  ├── Deterministic: Same plaintext → same ciphertext (allows =)    │   │
│  │  └── Randomized: Different ciphertext each time (more secure)      │   │
│  │                                                                     │   │
│  │  Pro: True separation, DBA can't decrypt                            │   │
│  │  Con: Limited query support, driver requirements                    │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Encryption in Transit

### TLS Architecture

```
TLS 1.3 HANDSHAKE (Simplified)
==============================

┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│  Client                                                Server               │
│  ┌──────┐                                             ┌──────┐             │
│  │      │                                             │      │             │
│  │      │─────── ClientHello ─────────────────────────│      │             │
│  │      │       (supported ciphers, key share)        │      │             │
│  │      │                                             │      │             │
│  │      │←────── ServerHello ─────────────────────────│      │             │
│  │      │       (selected cipher, key share)          │      │             │
│  │      │←────── Certificate ─────────────────────────│      │             │
│  │      │←────── CertificateVerify ───────────────────│      │             │
│  │      │←────── Finished ────────────────────────────│      │             │
│  │      │                                             │      │             │
│  │      │─────── Finished ────────────────────────────│      │             │
│  │      │                                             │      │             │
│  │      │════════ Application Data (encrypted) ═══════│      │             │
│  └──────┘                                             └──────┘             │
│                                                                             │
│  TLS 1.3 Improvements over TLS 1.2:                                         │
│  ├── 1-RTT handshake (was 2-RTT)                                           │
│  ├── 0-RTT resumption (with security trade-offs)                           │
│  ├── Removed weak algorithms (RSA key exchange, CBC, SHA-1)                │
│  ├── Perfect Forward Secrecy mandatory (ECDHE)                             │
│  └── Encrypted handshake (certificate hidden from observers)               │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘

TLS CIPHER SUITE (TLS 1.3 Format):
TLS_AES_256_GCM_SHA384
    │       │       │
    │       │       └── Hash for key derivation
    │       └────────── AEAD cipher for data
    └────────────────── Protocol

Recommended TLS 1.3 Cipher Suites:
├── TLS_AES_256_GCM_SHA384 (preferred)
├── TLS_AES_128_GCM_SHA256
└── TLS_CHACHA20_POLY1305_SHA256
```

### Mutual TLS (mTLS)

```
mTLS (MUTUAL TLS) AUTHENTICATION
================================

Standard TLS:
┌──────────────────────────────────────────────────────────────────────────┐
│  Client authenticates Server only (one-way)                              │
│                                                                          │
│  ┌────────┐                              ┌────────┐                      │
│  │ Client │──── "Who are you?" ─────────→│ Server │                      │
│  │        │←─── Server Certificate ──────│        │                      │
│  │        │     (proves server identity) │        │                      │
│  └────────┘                              └────────┘                      │
│                                                                          │
│  Use: Web browsers, public APIs                                          │
└──────────────────────────────────────────────────────────────────────────┘

Mutual TLS:
┌──────────────────────────────────────────────────────────────────────────┐
│  Both parties authenticate (two-way)                                     │
│                                                                          │
│  ┌────────┐                              ┌────────┐                      │
│  │ Client │──── "Who are you?" ─────────→│ Server │                      │
│  │        │←─── Server Certificate ──────│        │                      │
│  │        │                              │        │                      │
│  │        │←─── "Who are you?" ──────────│        │                      │
│  │        │──── Client Certificate ─────→│        │                      │
│  │        │     (proves client identity) │        │                      │
│  └────────┘                              └────────┘                      │
│                                                                          │
│  Use: Service-to-service, APIs, Zero Trust                               │
└──────────────────────────────────────────────────────────────────────────┘

mTLS Architecture:
┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│                    ┌─────────────────────┐                                  │
│                    │  Certificate        │                                  │
│                    │  Authority (CA)     │                                  │
│                    └──────────┬──────────┘                                  │
│                 Issues certs  │  Issues certs                               │
│              ┌────────────────┴────────────────┐                            │
│              ▼                                 ▼                            │
│  ┌───────────────────┐              ┌───────────────────┐                   │
│  │  Service A        │              │  Service B        │                   │
│  │  ┌─────────────┐  │   mTLS       │  ┌─────────────┐  │                   │
│  │  │ Certificate │  │══════════════│  │ Certificate │  │                   │
│  │  │ Private Key │  │              │  │ Private Key │  │                   │
│  │  └─────────────┘  │              │  └─────────────┘  │                   │
│  └───────────────────┘              └───────────────────┘                   │
│                                                                             │
│  Service Mesh mTLS (Istio/Linkerd):                                         │
│  ├── Automatic certificate provisioning and rotation                        │
│  ├── Transparent encryption (apps don't know)                               │
│  └── Identity-based policies (allow Service A → Service B)                  │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Encryption in Use

### Confidential Computing

```
CONFIDENTIAL COMPUTING ARCHITECTURE
===================================

┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│  TRADITIONAL COMPUTE                    CONFIDENTIAL COMPUTE                │
│  ┌────────────────────────┐            ┌────────────────────────┐          │
│  │  Memory (Plaintext)    │            │  ┌──────────────────┐  │          │
│  │  ┌──────────────────┐  │            │  │ Encrypted Enclave│  │          │
│  │  │ Data in Clear    │  │            │  │ ┌──────────────┐ │  │          │
│  │  │ Code in Clear    │  │            │  │ │ Data (Clear) │ │  │          │
│  │  └──────────────────┘  │            │  │ │ Code (Clear) │ │  │          │
│  │  Visible to:           │            │  │ └──────────────┘ │  │          │
│  │  ├── Hypervisor        │            │  │ Invisible to:    │  │          │
│  │  ├── Host OS           │            │  │ ├── Hypervisor   │  │          │
│  │  ├── Cloud Provider    │            │  │ ├── Host OS      │  │          │
│  │  └── Admins            │            │  │ ├── Cloud Admins │  │          │
│  └────────────────────────┘            │  │ └── Other VMs    │  │          │
│                                        │  └──────────────────┘  │          │
│                                        │  Encrypted Memory      │          │
│                                        └────────────────────────┘          │
│                                                                             │
│  TECHNOLOGIES:                                                              │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  Intel SGX (Software Guard Extensions)                              │   │
│  │  ├── Application-level enclaves                                     │   │
│  │  ├── Limited enclave size (~256MB)                                  │   │
│  │  └── Requires app modification                                      │   │
│  │                                                                     │   │
│  │  AMD SEV (Secure Encrypted Virtualization)                          │   │
│  │  ├── VM-level encryption                                            │   │
│  │  ├── Full VM memory encrypted                                       │   │
│  │  └── Transparent to applications                                    │   │
│  │                                                                     │   │
│  │  ARM TrustZone                                                      │   │
│  │  ├── Secure world / Normal world separation                         │   │
│  │  └── Common in mobile devices                                       │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
│  CLOUD OFFERINGS:                                                           │
│  ├── Azure Confidential Computing (DCsv2, DCsv3 VMs)                       │
│  ├── AWS Nitro Enclaves                                                    │
│  ├── GCP Confidential VMs (AMD SEV)                                        │
│  └── IBM Cloud Data Shield                                                 │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Homomorphic Encryption

```
HOMOMORPHIC ENCRYPTION CONCEPT
==============================

┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│  PROBLEM: Need to compute on encrypted data without decrypting              │
│                                                                             │
│  Traditional:                                                               │
│  ┌─────────┐    ┌─────────┐    ┌─────────┐    ┌─────────┐                  │
│  │Encrypted│───→│ Decrypt │───→│ Compute │───→│ Encrypt │───→ Result       │
│  │  Data   │    │(Expose!)│    │         │    │         │                  │
│  └─────────┘    └─────────┘    └─────────┘    └─────────┘                  │
│                      ▲                                                      │
│                      └── Data exposed during computation                    │
│                                                                             │
│  Homomorphic:                                                               │
│  ┌─────────┐    ┌─────────────────────┐                                    │
│  │Encrypted│───→│ Compute on Cipher   │───→ Encrypted Result               │
│  │  Data   │    │ (Never decrypted!)  │                                    │
│  └─────────┘    └─────────────────────┘                                    │
│                                                                             │
│  TYPES:                                                                     │
│  ├── Partially Homomorphic: One operation (add OR multiply)                │
│  ├── Somewhat Homomorphic: Limited operations before noise buildup         │
│  └── Fully Homomorphic (FHE): Any computation (very slow)                  │
│                                                                             │
│  USE CASES:                                                                 │
│  ├── Privacy-preserving analytics                                          │
│  ├── Secure multi-party computation                                        │
│  ├── Encrypted database queries                                            │
│  └── ML on encrypted data                                                  │
│                                                                             │
│  LIMITATIONS (Current State):                                               │
│  ├── Performance: 1000x-1,000,000x slower than plaintext                   │
│  ├── Ciphertext expansion: 10x-1000x larger                                │
│  └── Complexity: Requires expertise to implement                           │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Key Management Architecture

### Key Hierarchy Design

```
ENTERPRISE KEY HIERARCHY
========================

┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │                    ROOT OF TRUST (HSM)                               │   │
│  │  ┌───────────────────────────────────────────────────────────────┐  │   │
│  │  │  Master Key (MK)                                              │  │   │
│  │  │  ├── Never exported from HSM                                  │  │   │
│  │  │  ├── Used to protect all other keys                           │  │   │
│  │  │  └── Backed up with split knowledge/dual control              │  │   │
│  │  └───────────────────────────────────────────────────────────────┘  │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                      │                                      │
│                                      │ Protects                             │
│                    ┌─────────────────┴─────────────────┐                   │
│                    ▼                                   ▼                   │
│  ┌──────────────────────────────┐   ┌──────────────────────────────┐      │
│  │  KEY ENCRYPTION KEYS (KEK)   │   │  KEY ENCRYPTION KEYS (KEK)   │      │
│  │  (Environment: Production)   │   │  (Environment: Development)  │      │
│  │  ├── KEK-Prod-DB             │   │  ├── KEK-Dev-DB              │      │
│  │  ├── KEK-Prod-Storage        │   │  ├── KEK-Dev-Storage         │      │
│  │  └── KEK-Prod-App            │   │  └── KEK-Dev-App             │      │
│  └──────────────────────────────┘   └──────────────────────────────┘      │
│                    │                                   │                   │
│                    │ Protects                          │                   │
│                    ▼                                   ▼                   │
│  ┌──────────────────────────────┐   ┌──────────────────────────────┐      │
│  │  DATA ENCRYPTION KEYS (DEK)  │   │  DATA ENCRYPTION KEYS (DEK)  │      │
│  │  ├── Customer-DB-Table1      │   │  ├── Test-DB-Table1          │      │
│  │  ├── Customer-DB-Table2      │   │  └── Test-Storage-Bucket     │      │
│  │  ├── S3-Bucket-Invoices      │   │                              │      │
│  │  └── File-Share-HR           │   │                              │      │
│  └──────────────────────────────┘   └──────────────────────────────┘      │
│                    │                                                       │
│                    │ Encrypts                                              │
│                    ▼                                                       │
│  ┌─────────────────────────────────────────────────────────────────────┐  │
│  │  ACTUAL DATA                                                        │  │
│  │  Customer records, files, database content                          │  │
│  └─────────────────────────────────────────────────────────────────────┘  │
│                                                                             │
│  BENEFITS:                                                                  │
│  ├── Separation of environments (prod keys never in dev)                   │
│  ├── Granular access control at each level                                 │
│  ├── Efficient key rotation (rotate KEK, re-wrap DEKs)                     │
│  └── Compliance (prove key segregation for auditors)                       │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Cloud KMS Services

### Cloud KMS Comparison

```
CLOUD KMS SERVICE COMPARISON
============================

┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │                          AWS KMS                                    │   │
│  ├─────────────────────────────────────────────────────────────────────┤   │
│  │  Key Types:                                                         │   │
│  │  ├── AWS Managed Keys (aws/service)                                 │   │
│  │  ├── Customer Managed Keys (CMK)                                    │   │
│  │  └── Customer Owned Keys (CloudHSM)                                 │   │
│  │                                                                     │   │
│  │  Features:                                                          │   │
│  │  ├── Automatic key rotation (annually for CMK)                      │   │
│  │  ├── Key policies + IAM policies                                    │   │
│  │  ├── CloudTrail integration                                         │   │
│  │  ├── Multi-region keys                                              │   │
│  │  └── FIPS 140-2 Level 2 (Level 3 with CloudHSM)                     │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │                       Azure Key Vault                               │   │
│  ├─────────────────────────────────────────────────────────────────────┤   │
│  │  Vault Types:                                                       │   │
│  │  ├── Standard (software-protected)                                  │   │
│  │  └── Premium (HSM-protected, FIPS 140-2 Level 2)                    │   │
│  │                                                                     │   │
│  │  Features:                                                          │   │
│  │  ├── Keys, Secrets, Certificates                                    │   │
│  │  ├── Managed HSM (FIPS 140-2 Level 3)                               │   │
│  │  ├── RBAC + Access Policies                                         │   │
│  │  ├── Soft-delete and purge protection                               │   │
│  │  └── Private endpoint support                                       │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │                       GCP Cloud KMS                                 │   │
│  ├─────────────────────────────────────────────────────────────────────┤   │
│  │  Protection Levels:                                                 │   │
│  │  ├── SOFTWARE                                                       │   │
│  │  ├── HSM (FIPS 140-2 Level 3)                                       │   │
│  │  └── EXTERNAL (External Key Manager - EKM)                          │   │
│  │                                                                     │   │
│  │  Features:                                                          │   │
│  │  ├── Automatic rotation                                             │   │
│  │  ├── Key versions and destruction scheduling                        │   │
│  │  ├── IAM-based access control                                       │   │
│  │  └── Cloud Audit Logs integration                                   │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### BYOK vs HYOK

```
KEY OWNERSHIP MODELS
====================

┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│  PROVIDER-MANAGED KEYS                                                      │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  ┌─────────────┐    ┌─────────────┐    ┌─────────────┐              │   │
│  │  │    Keys     │    │   Cloud     │    │    Data     │              │   │
│  │  │  (Provider  │───→│   KMS       │───→│  Encryption │              │   │
│  │  │   Managed)  │    │             │    │             │              │   │
│  │  └─────────────┘    └─────────────┘    └─────────────┘              │   │
│  │  Pro: Simple, no key management burden                              │   │
│  │  Con: Provider has key access                                       │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
│  BRING YOUR OWN KEY (BYOK)                                                  │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  ┌─────────────┐    ┌─────────────┐    ┌─────────────┐              │   │
│  │  │ Customer    │    │   Cloud     │    │    Data     │              │   │
│  │  │ Generated   │───→│   KMS       │───→│  Encryption │              │   │
│  │  │ Key (Import)│    │ (Imported)  │    │             │              │   │
│  │  └─────────────┘    └─────────────┘    └─────────────┘              │   │
│  │  Pro: Customer controls key generation                              │   │
│  │  Con: Key still stored in cloud KMS                                 │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
│  HOLD YOUR OWN KEY (HYOK)                                                   │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │  ┌─────────────┐         ┌─────────────┐    ┌─────────────┐        │   │
│  │  │ Customer    │────────→│    Cloud    │───→│    Data     │        │   │
│  │  │ On-Prem HSM │ API Call│   Service   │    │  Encrypted  │        │   │
│  │  │ (Key stays) │←────────│             │    │             │        │   │
│  │  └─────────────┘  Result └─────────────┘    └─────────────┘        │   │
│  │  Pro: Key never leaves customer control                             │   │
│  │  Con: Availability depends on customer infra, higher latency        │   │
│  │                                                                     │   │
│  │  Implementations:                                                   │   │
│  │  ├── Azure Key Vault External Key Manager                           │   │
│  │  ├── GCP Cloud External Key Manager (EKM)                           │   │
│  │  └── AWS External Key Store (XKS)                                   │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Hardware Security Modules

### HSM Architecture

```
HARDWARE SECURITY MODULE (HSM) ARCHITECTURE
===========================================

┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│                         HSM APPLIANCE                                       │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │                                                                     │   │
│  │  ┌─────────────────────────────────────────────────────────────┐   │   │
│  │  │                 TAMPER-RESISTANT BOUNDARY                    │   │   │
│  │  │                                                             │   │   │
│  │  │  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐ │   │   │
│  │  │  │  Secure     │  │   Crypto    │  │   Key Storage       │ │   │   │
│  │  │  │  Processor  │  │   Engine    │  │   (Battery-backed)  │ │   │   │
│  │  │  └─────────────┘  └─────────────┘  └─────────────────────┘ │   │   │
│  │  │                                                             │   │   │
│  │  │  ┌─────────────────────────────────────────────────────────┐│   │   │
│  │  │  │  Tamper Detection:                                      ││   │   │
│  │  │  │  ├── Physical intrusion sensors                         ││   │   │
│  │  │  │  ├── Temperature monitoring                             ││   │   │
│  │  │  │  ├── Voltage monitoring                                 ││   │   │
│  │  │  │  └── Response: Zeroize keys on tamper                   ││   │   │
│  │  │  └─────────────────────────────────────────────────────────┘│   │   │
│  │  │                                                             │   │   │
│  │  └─────────────────────────────────────────────────────────────┘   │   │
│  │                                                                     │   │
│  │  ┌─────────────────────────────────────────────────────────────┐   │   │
│  │  │  Network Interface    │    Management Interface             │   │   │
│  │  │  (PKCS#11, JCE, etc.) │    (Admin console)                  │   │   │
│  │  └─────────────────────────────────────────────────────────────┘   │   │
│  │                                                                     │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
│  FIPS 140-2 LEVELS:                                                         │
│  ┌────────────┬────────────────────────────────────────────────────────┐   │
│  │ Level 1    │ Basic security, no physical security                   │   │
│  │ Level 2    │ Tamper-evident seals, role-based auth                  │   │
│  │ Level 3    │ Tamper-resistant, identity-based auth, zeroization     │   │
│  │ Level 4    │ Environmental protection, immediate zeroization        │   │
│  └────────────┴────────────────────────────────────────────────────────┘   │
│                                                                             │
│  CLOUD HSM OPTIONS:                                                         │
│  ├── AWS CloudHSM (FIPS 140-2 Level 3)                                     │
│  ├── Azure Dedicated HSM (FIPS 140-2 Level 3)                              │
│  ├── GCP Cloud HSM (FIPS 140-2 Level 3)                                    │
│  └── On-prem: Thales Luna, Entrust nShield, Utimaco                        │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Key Lifecycle Management

### Key Lifecycle States

```
KEY LIFECYCLE MANAGEMENT
========================

┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│  ┌──────────┐    ┌──────────┐    ┌──────────┐    ┌──────────┐    ┌──────┐ │
│  │ Generate │───→│  Active  │───→│ Rotation │───→│Deprecated│───→│Destroy│ │
│  └──────────┘    └──────────┘    └──────────┘    └──────────┘    └──────┘ │
│                                                                             │
│  GENERATE:                                                                  │
│  ├── Use cryptographically secure random number generator                   │
│  ├── Generate in HSM when possible                                         │
│  ├── Document key purpose, owner, expiry                                   │
│  └── Establish backup and recovery procedures                              │
│                                                                             │
│  ACTIVE:                                                                    │
│  ├── Key available for encrypt and decrypt                                 │
│  ├── Monitor usage and access                                              │
│  ├── Audit key operations                                                  │
│  └── Typical period: 1-3 years (depending on algorithm and use)            │
│                                                                             │
│  ROTATION:                                                                  │
│  ├── Generate new key version                                              │
│  ├── New key used for encrypt                                              │
│  ├── Old key still available for decrypt                                   │
│  └── Re-encrypt data if required                                           │
│                                                                             │
│  DEPRECATED:                                                                │
│  ├── Key available for decrypt only                                        │
│  ├── No new encryptions                                                    │
│  ├── Plan data re-encryption or migration                                  │
│  └── Document dependent systems                                            │
│                                                                             │
│  DESTROY:                                                                   │
│  ├── Cryptographic erasure (zeroization)                                   │
│  ├── Verify all encrypted data migrated or deleted                         │
│  ├── Document destruction                                                  │
│  └── Irreversible - ensure backups if needed                               │
│                                                                             │
│  ROTATION FREQUENCY GUIDELINES:                                             │
│  ┌────────────────────────────────────────────────────────────────────┐    │
│  │ Key Type                    │ Recommended Rotation                 │    │
│  ├─────────────────────────────┼──────────────────────────────────────┤    │
│  │ Symmetric (AES)             │ 1-2 years or per policy              │    │
│  │ Asymmetric (RSA/EC)         │ 2-3 years                            │    │
│  │ TLS certificates            │ 1 year (shorter for security)        │    │
│  │ API keys                    │ 90 days                              │    │
│  │ Service account keys        │ 90-180 days                          │    │
│  │ Root CA keys                │ 10-20 years (or never)               │    │
│  └─────────────────────────────┴──────────────────────────────────────┘    │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### Key Recovery and Escrow

```
KEY RECOVERY ARCHITECTURE
=========================

┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│  SPLIT KNOWLEDGE / DUAL CONTROL                                             │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │                                                                     │   │
│  │  Master Key Split Using Shamir's Secret Sharing:                    │   │
│  │                                                                     │   │
│  │  Master Key ───→ Split into N shares (e.g., 5)                      │   │
│  │                  Require K shares to reconstruct (e.g., 3)          │   │
│  │                                                                     │   │
│  │  ┌──────┐ ┌──────┐ ┌──────┐ ┌──────┐ ┌──────┐                      │   │
│  │  │Share1│ │Share2│ │Share3│ │Share4│ │Share5│                      │   │
│  │  └──┬───┘ └──┬───┘ └──┬───┘ └──┬───┘ └──┬───┘                      │   │
│  │     │        │        │        │        │                           │   │
│  │     ▼        ▼        ▼        ▼        ▼                           │   │
│  │  ┌──────┐ ┌──────┐ ┌──────┐ ┌──────┐ ┌──────┐                      │   │
│  │  │Custd1│ │Custd2│ │Custd3│ │Custd4│ │Custd5│                      │   │
│  │  │(CEO) │ │(CFO) │ │(CIO) │ │(CISO)│ │(Legal)│                     │   │
│  │  └──────┘ └──────┘ └──────┘ └──────┘ └──────┘                      │   │
│  │                                                                     │   │
│  │  Recovery: Any 3 custodians can reconstruct the master key          │   │
│  │  No single person has complete access                               │   │
│  │                                                                     │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
│  ESCROW SCENARIOS:                                                          │
│  ├── Business continuity (key holder incapacitated)                        │
│  ├── Legal/regulatory requirements (lawful intercept)                      │
│  ├── M&A (access to acquired company's encrypted data)                     │
│  └── Disaster recovery (primary HSM destroyed)                             │
│                                                                             │
│  ESCROW BEST PRACTICES:                                                     │
│  ├── Escrow wrapped keys, not plaintext                                    │
│  ├── Geographic separation of key shares                                   │
│  ├── Regular validation of recovery procedures                             │
│  ├── Audit all escrow access                                               │
│  └── Clear governance for escrow release                                   │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## Interview Practice Questions

### Question 1: Encryption Strategy Design
**Q: Design an encryption strategy for a healthcare company storing patient records in AWS. They need to meet HIPAA requirements and want defense in depth.**

**A:**
```
HEALTHCARE ENCRYPTION STRATEGY
==============================

Multi-Layer Encryption Architecture:

LAYER 1: Infrastructure (Volume/Disk Level)
├── EBS Encryption with AWS-managed keys (aws/ebs)
├── Provides baseline protection for all volumes
├── Transparent to applications
└── Protects against physical theft of disks

LAYER 2: Storage (Object/File Level)
├── S3 Server-Side Encryption (SSE-KMS)
│   └── Customer Managed Key (CMK) for PHI buckets
├── Key policy restricts access to authorized services
├── Bucket policy enforces encryption
└── S3 Block Public Access enabled

LAYER 3: Database (TDE + Field Level)
├── RDS encryption enabled (TDE equivalent)
├── Application-level encryption for highly sensitive fields:
│   ├── SSN: Always Encrypted pattern
│   ├── Medical Record Numbers
│   └── Diagnosis codes
└── Keys stored in AWS KMS (separate CMK for PHI)

LAYER 4: Application (End-to-End)
├── Client-side encryption for PHI in transit to S3
├── AWS Encryption SDK for envelope encryption
└── Decrypt only when authorized user views record

Key Management:
┌──────────────────────────────────────────────────────────────────────────┐
│                                                                          │
│  KMS Key Hierarchy:                                                      │
│  ├── CMK: phi-master-key (Customer managed, annual rotation)             │
│  │   ├── Key policy: Only specific IAM roles                             │
│  │   ├── CloudTrail logging mandatory                                    │
│  │   └── Multi-region for DR                                             │
│  │                                                                       │
│  ├── CMK: phi-database-key (RDS encryption)                              │
│  └── CMK: phi-storage-key (S3 encryption)                                │
│                                                                          │
│  Access Control:                                                         │
│  ├── Separate encryption and decryption permissions                      │
│  ├── Decrypt only for authorized roles (clinicians, not admins)          │
│  └── Key usage logged to CloudTrail → SIEM                               │
│                                                                          │
└──────────────────────────────────────────────────────────────────────────┘

Encryption in Transit:
├── TLS 1.3 for all connections
├── VPC endpoints for AWS services (no public internet)
├── mTLS for service-to-service communication
└── Certificate management via ACM

HIPAA Compliance:
├── All PHI encrypted at rest (addressable → we've addressed it)
├── All PHI encrypted in transit (required)
├── Key management documented
├── Access logging enabled
└── Business Associate Agreement (BAA) with AWS
```

---

### Question 2: Key Management Architecture
**Q: Your organization operates in AWS, Azure, and on-premises. Design a key management architecture that provides consistent security across all environments.**

**A:**
```
MULTI-CLOUD KEY MANAGEMENT ARCHITECTURE
=======================================

Architecture Overview:
┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│                    ┌─────────────────────────┐                              │
│                    │   CENTRAL KEY MANAGER   │                              │
│                    │   (HashiCorp Vault)     │                              │
│                    │   ├── Master keys       │                              │
│                    │   ├── Policy engine     │                              │
│                    │   └── Audit logging     │                              │
│                    └───────────┬─────────────┘                              │
│                                │                                            │
│         ┌──────────────────────┼──────────────────────┐                    │
│         │                      │                      │                    │
│         ▼                      ▼                      ▼                    │
│  ┌─────────────┐       ┌─────────────┐       ┌─────────────┐              │
│  │   AWS KMS   │       │Azure KeyVault│      │  On-Prem    │              │
│  │  (Transit)  │       │  (Transit)  │       │    HSM      │              │
│  │             │       │             │       │  (Thales)   │              │
│  └─────────────┘       └─────────────┘       └─────────────┘              │
│         │                      │                      │                    │
│         ▼                      ▼                      ▼                    │
│  ┌─────────────┐       ┌─────────────┐       ┌─────────────┐              │
│  │ AWS Services│       │Azure Services│      │  Legacy     │              │
│  │ S3, RDS, EBS│       │Blob,SQL,Disk│       │  Systems    │              │
│  └─────────────┘       └─────────────┘       └─────────────┘              │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘

Component Details:

1. Central Key Manager (HashiCorp Vault Enterprise)
├── Deployed in primary data center with HA
├── Auto-unseal with cloud KMS or HSM
├── Namespaces for environment separation
├── Transit secrets engine for encryption-as-a-service
├── PKI secrets engine for certificate management
└── Audit device for compliance logging

2. Cloud Integration (Vault Transit)
├── Applications call Vault for encrypt/decrypt
├── Keys never leave Vault
├── Consistent API across all environments
└── Fallback: Native KMS for service-managed encryption

3. Key Distribution Strategy:
├── Master keys: Only in Vault (backed by HSM)
├── Cloud service encryption: Native KMS (wrapped by Vault key)
├── Application encryption: Vault Transit engine
└── TLS certificates: Vault PKI engine

Policy Enforcement:
├── Vault policies for role-based key access
├── Sentinel policies for governance (e.g., no key export)
├── Audit all key operations
└── SIEM integration for anomaly detection

Key Rotation:
├── Vault handles automatic rotation
├── Cloud KMS integrated for service encryption
├── Applications use key versioning (latest for encrypt)
└── Rotation doesn't require re-encryption (envelope pattern)
```

---

### Question 3: TLS Inspection Trade-offs
**Q: The security team wants to implement TLS inspection on all outbound traffic, but the privacy team is concerned. How do you balance these requirements?**

**A:**
```
TLS INSPECTION ARCHITECTURE WITH PRIVACY CONTROLS
=================================================

Technical Architecture:
┌──────────────────────────────────────────────────────────────────────────┐
│                                                                          │
│  ┌──────────┐    ┌───────────────┐    ┌────────────┐    ┌─────────────┐ │
│  │ Endpoint │───→│  TLS Proxy    │───→│  Security  │───→│  Internet   │ │
│  │          │    │  (Decrypt/    │    │  Stack     │    │             │ │
│  │          │    │   Inspect/    │    │  (DLP,IPS) │    │             │ │
│  │          │    │   Re-encrypt) │    │            │    │             │ │
│  └──────────┘    └───────────────┘    └────────────┘    └─────────────┘ │
│                                                                          │
└──────────────────────────────────────────────────────────────────────────┘

Privacy-Preserving Inspection:

1. Categorical Bypass (Don't Inspect):
├── Banking/financial sites (certificate pinning, privacy)
├── Healthcare portals (HIPAA considerations)
├── Government sites
├── Personal webmail (if allowed by policy)
└── Sites with certificate pinning (breaks anyway)

2. User Notification:
├── Policy acknowledgment during onboarding
├── Browser notification via PAC file message
├── Privacy policy clearly states inspection scope
└── Annual reminder/re-acknowledgment

3. Data Handling:
├── Inspected content not logged (only metadata)
├── DLP alerts store only matched patterns, not full content
├── Session data deleted after inspection
└── No inspection of personal accounts (if identifiable)

4. Technical Controls:
├── Proxy logs encrypted at rest
├── Access to decrypted traffic highly restricted
├── Audit all proxy admin access
└── Regular review of bypass list

Governance Framework:
┌──────────────────────────────────────────────────────────────────────────┐
│                                                                          │
│  Inspection Governance:                                                  │
│  ├── Policy approved by Legal, Privacy, HR, Security                     │
│  ├── Clear scope: Corporate devices, corporate network only              │
│  ├── Exclusions documented and reviewed quarterly                        │
│  ├── Incident response: What happens if personal data seen?              │
│  └── Audit: Annual review of inspection justification                    │
│                                                                          │
│  Employee Communication:                                                 │
│  ├── Acceptable Use Policy updated                                       │
│  ├── Training on what is/isn't inspected                                 │
│  ├── Clear guidance: Use personal device for personal banking            │
│  └── Feedback mechanism for concerns                                     │
│                                                                          │
└──────────────────────────────────────────────────────────────────────────┘

Alternative Approaches (If Full Inspection Rejected):
├── Inspect only high-risk categories (file sharing, unknown)
├── Metadata-only analysis (no content inspection)
├── DNS-level filtering (no TLS inspection)
└── Endpoint DLP instead of network inspection
```

---

### Question 4: Encryption Key Compromise Response
**Q: You discover that an encryption key may have been compromised. What is your incident response process?**

**A:**
```
KEY COMPROMISE INCIDENT RESPONSE
================================

IMMEDIATE ACTIONS (First 4 Hours):

1. Assess Scope
├── Which key is compromised? (CMK, DEK, TLS cert?)
├── What data is protected by this key?
├── What's the compromise evidence?
└── Is this confirmed or suspected?

2. Contain
├── Disable compromised key (prevent further encryption)
├── Rotate all derived keys
├── If TLS cert: Revoke immediately (CRL, OCSP)
├── Preserve audit logs for investigation
└── Don't destroy key yet (need for decryption)

3. Notify
├── Security leadership
├── Legal (potential breach notification)
├── Compliance (regulatory implications)
└── Affected system owners

SHORT-TERM ACTIONS (4-48 Hours):

4. Investigate
├── How was key compromised?
│   ├── Stolen from HSM? (very bad - physical security failure)
│   ├── Exported improperly? (policy failure)
│   ├── Extracted from memory? (malware)
│   └── Social engineering?
├── What was done with the key?
├── What data may have been decrypted?
└── Timeline of compromise

5. Re-encrypt Data
├── Generate new key
├── Re-encrypt all data protected by compromised key
├── Update all systems to use new key
└── Verify encryption with new key

6. Remediate Root Cause
├── Close vulnerability that allowed compromise
├── Update access controls
├── Enhance monitoring
└── Update incident response procedures

LONG-TERM ACTIONS (48+ Hours):

7. Destroy Old Key
├── Only after all data re-encrypted
├── Cryptographic erasure (zeroization)
├── Document destruction
└── Update key inventory

8. Post-Incident Review
├── Root cause analysis
├── Control gap assessment
├── Process improvements
├── Training updates

Key Compromise Decision Matrix:
┌─────────────────────────────────────────────────────────────────────────────┐
│ Key Type          │ Compromise Impact     │ Action                         │
├───────────────────┼───────────────────────┼────────────────────────────────┤
│ Root/Master Key   │ CRITICAL - All data   │ Emergency re-key entire infra  │
│ KEK (Key Encrypt) │ HIGH - Multiple DEKs  │ Rotate KEK, re-wrap all DEKs   │
│ DEK (Data Encrypt)│ MEDIUM - Specific data│ Re-encrypt affected data       │
│ TLS Certificate   │ MEDIUM - MITM risk    │ Revoke, issue new cert         │
│ API Key           │ VARIES - Depends on   │ Revoke, issue new, audit usage │
│                   │ API permissions       │                                │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

### Question 5: HSM vs Cloud KMS Decision
**Q: When would you recommend a dedicated HSM (like CloudHSM) versus using the cloud provider's managed KMS service?**

**A:**
```
HSM vs MANAGED KMS DECISION FRAMEWORK
=====================================

USE MANAGED KMS (AWS KMS, Azure Key Vault, GCP Cloud KMS):
┌──────────────────────────────────────────────────────────────────────────┐
│                                                                          │
│  When:                                                                   │
│  ├── Standard compliance requirements (SOC 2, ISO 27001)                 │
│  ├── FIPS 140-2 Level 2 is sufficient                                    │
│  ├── Cost sensitivity (KMS ~$1/key/month vs HSM ~$1/hour)                │
│  ├── Prefer managed service (no operational burden)                      │
│  ├── Tight integration with cloud services needed                        │
│  └── Standard key types and algorithms are sufficient                    │
│                                                                          │
│  Benefits:                                                               │
│  ├── Fully managed (patches, HA, DR handled by provider)                 │
│  ├── Native integration with cloud services                              │
│  ├── Pay-per-use pricing                                                 │
│  ├── Easy key rotation                                                   │
│  └── CloudTrail/audit logging built-in                                   │
│                                                                          │
└──────────────────────────────────────────────────────────────────────────┘

USE DEDICATED HSM (CloudHSM, Dedicated HSM, On-prem HSM):
┌──────────────────────────────────────────────────────────────────────────┐
│                                                                          │
│  When:                                                                   │
│  ├── FIPS 140-2 Level 3 required (regulatory/contractual)                │
│  ├── PCI DSS key management requirements                                 │
│  ├── Custom key types or algorithms needed                               │
│  ├── Key material must never be accessible to cloud provider             │
│  ├── Need to use keys across multiple clouds consistently                │
│  ├── Legacy application integration (PKCS#11, JCE)                       │
│  └── Code signing or PKI root key storage                                │
│                                                                          │
│  Trade-offs:                                                             │
│  ├── Higher cost (~$1-2/hour per HSM, need 2+ for HA)                    │
│  ├── Operational responsibility (backups, HA configuration)              │
│  ├── More complex integration                                            │
│  └── Capacity planning required                                          │
│                                                                          │
└──────────────────────────────────────────────────────────────────────────┘

DECISION MATRIX:
┌─────────────────────────────────────────────────────────────────────────────┐
│                                                                             │
│  Requirement                              │ KMS │ CloudHSM │ On-Prem HSM   │
│  ─────────────────────────────────────────┼─────┼──────────┼───────────────│
│  FIPS 140-2 Level 2                       │  ✓  │    ✓     │      ✓        │
│  FIPS 140-2 Level 3                       │  ✗  │    ✓     │      ✓        │
│  Native cloud service integration         │  ✓  │    △     │      ✗        │
│  Cross-cloud consistency                  │  ✗  │    ✗     │      ✓        │
│  Custom algorithms                        │  ✗  │    ✓     │      ✓        │
│  Key never leaves customer control        │  ✗  │    ✓     │      ✓        │
│  Low operational burden                   │  ✓  │    △     │      ✗        │
│  Cost-effective at scale                  │  ✓  │    ✗     │      ✓        │
│  PKCS#11 / JCE interface                  │  ✗  │    ✓     │      ✓        │
│                                                                             │
│  ✓ = Yes, ✗ = No, △ = Partial                                              │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘

HYBRID APPROACH (Common Pattern):
├── Use CloudHSM for root/master keys (highest security)
├── CloudHSM keys wrap KMS CMKs (key hierarchy)
├── KMS for day-to-day operations (ease of use)
└── Best of both: HSM-level security for critical keys, KMS convenience
```

---

## Key Takeaways

1. **Encryption is Defense in Depth**: Use multiple layers (disk, file, field, application)

2. **Keys are the Crown Jewels**: Strong encryption is useless with poor key management

3. **Envelope Encryption**: Standard pattern for scalable key management

4. **TLS 1.3**: Mandatory for all new deployments, plan TLS 1.2 deprecation

5. **Know Your HSM Options**: Managed KMS for most cases, dedicated HSM for specific compliance

6. **Key Lifecycle**: Plan for generation, rotation, and destruction from day one

---

[← Previous Lesson](./01-data-classification-governance.md) | [Next Lesson →](./03-data-loss-prevention.md) | [Back to Module](./README.md)

---

*This lesson is part of the Enterprise Security Architect learning path.*
