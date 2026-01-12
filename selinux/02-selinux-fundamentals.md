# Phase 1: SELinux Fundamentals (Week 1)


## Table of Contents

- [Concepts](#concepts)
    - [What is SELinux and Why Does it Exist?](#what-is-selinux-and-why-does-it-exist)
    - [SELinux Modes](#selinux-modes)
    - [Policy Types](#policy-types)
    - [The Decision Process](#the-decision-process)
- [Hands-On Exercises](#hands-on-exercises)
    - [1. Check SELinux Status](#1-check-selinux-status)
    - [2. Switch Modes (Temporarily)](#2-switch-modes-temporarily)
    - [3. Understanding the Config File](#3-understanding-the-config-file)
- [Key Takeaway](#key-takeaway)

---

## Concepts

### What is SELinux and Why Does it Exist?
**SELinux (Security-Enhanced Linux)** is a security architecture for Linux systems that allows administrators to have more control over who can access the system.
- **MAC (Mandatory Access Control)**: Unlike DAC (where users control permissions), MAC policies are defined by the system administrator and users cannot override them.
- **Containment**: It confines processes to their own "domains". If a web server is hacked, the hacker is trapped in the web server's domain, unable to access `/home` or `/etc/shadow`.

### SELinux Modes
It's critical to know which mode your system is in:
1. **Enforcing**: The default. Policies are enforced, and access is denied if not explicitly allowed. Violations are logged.
2. **Permissive**: Policies are checked, but *not* enforced. Violations are only logged. Great for troubleshooting.
3. **Disabled**: SELinux is turned off completely. Requires a reboot to re-enable.

### Policy Types
- **Targeted**: The default in Red Hat/CentOS/Fedora. Only specific "targeted" network daemons (httpd, named, dhcpd, etc.) are protected. Everything else runs in an `unconfined` domain.
- **MLS (Multi-Level Security)**: Used in highly secure environments (government/military). Much more complex.

### The Decision Process
**"Deny by Default"**: If there isn't a rule explicitly allowing an action, SELinux denies it.

## Hands-On Exercises

### 1. Check SELinux Status
See if you are running Enforcing, Permissive, or Disabled.

```bash
# Detailed status
sestatus

# Simple status (Enforcing/Permissive/Disabled)
getenforce
```

### 2. Switch Modes (Temporarily)
You can switch between Enforcing and Permissive on the fly (does not survive reboot). **Note**: You cannot switch to/from Disabled without a reboot and config change.

```bash
# Set to Permissive (0)
sudo setenforce 0
getenforce
# Output should be: Permissive

# Set back to Enforcing (1)
sudo setenforce 1
getenforce
# Output should be: Enforcing
```

### 3. Understanding the Config File
To make changes persistent, you edit `/etc/selinux/config`.

```bash
cat /etc/selinux/config
```
*Look for `SELINUX=enforcing`.*

## Key Takeaway
> **SELinux is "Deny by Default", unlike traditional Linux permissions.**
> Always check `getenforce` first when troubleshooting a "Permission Denied" error on a Red Hat-based system.
