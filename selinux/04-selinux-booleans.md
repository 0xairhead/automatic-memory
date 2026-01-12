# Phase 3: SELinux Booleans (Week 3)


## Table of Contents

- [Concepts](#concepts)
    - [What are Booleans?](#what-are-booleans)
    - [Common Examples](#common-examples)
- [Hands-On Exercises](#hands-on-exercises)
    - [1. List All Booleans](#1-list-all-booleans)
    - [2. Check a Specific Boolean](#2-check-a-specific-boolean)
    - [3. Change a Boolean (Runtime vs Persistent)](#3-change-a-boolean-runtime-vs-persistent)
- [Scenario: The "Bad Gateway" Optimization](#scenario-the-bad-gateway-optimization)
- [Key Takeaway](#key-takeaway)

---

## Concepts

### What are Booleans?
Booleans are **on/off switches** (true/false) built into the SELinux policy. They allow administrators to change specific parts of the security policy at runtime without writing any code.

Think of them as "Safe Flexibility". The policy writers knew you might want to do certain things (like let a web server connect to a database), so they made it a toggle.

### Common Examples
*   `httpd_can_network_connect`: Allows Apache/Nginx to make outbound network connections (e.g., for proxying).
*   `httpd_enable_homedirs`: Allows Apache to read `~/public_html`.
*   `ftpd_anon_write`: Allows anonymous FTP users to upload files.

## Hands-On Exercises

### 1. List All Booleans
There are hundreds. Use `getsebool` and `grep` to find what you need.

```bash
# List all
getsebool -a

# Find web server related ones
getsebool -a | grep httpd
```

### 2. Check a Specific Boolean
```bash
getsebool httpd_can_network_connect
# Output: httpd_can_network_connect --> off
```

### 3. Change a Boolean (Runtime vs Persistent)

**Temporary Change (Reset on reboot)**:
```bash
sudo setsebool httpd_can_network_connect on
```

**Persistent Change (Survives reboot)**:
Use the `-P` flag. *Note: This takes a few seconds as it recompiles the policy.*
```bash
sudo setsebool -P httpd_can_network_connect on
```

## Scenario: The "Bad Gateway" Optimization
**Problem**: You set up Nginx as a reverse proxy to a backend app on port 8080. You get a "502 Bad Gateway".
**Diagnosis**: `ausearch` shows a denial for `name_connect`.
**Fix**: `httpd_t` isn't allowed to connect to network ports by default.
**Solution**:
```bash
sudo setsebool -P httpd_can_network_connect 1
```

## Key Takeaway
> **Booleans allow you to tweak policy without being a policy expert.**
> Before writing custom policy, always look for a boolean that might solve your problem.
