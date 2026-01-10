# Phase 4: Troubleshooting & Auditing (Week 4)

## Concepts
When things break, don't blame SELinux blindly. Use the tools to confirm if SELinux is actively denying access.

### The Audit Log
All SELinux denials (AVCs - Access Vector Cache denials) are logged.
- **Location**: `/var/log/audit/audit.log`
- **Identifier**: Look for `type=AVC`.
- **Key Fields**:
    - `scontext`: Source context (the process)
    - `tcontext`: Target context (the file/resource)
    - `tclass`: The object class (file, dir, tcp_socket)
    - `denied`: The permission that was blocked (read, write, connect)

### Core Tools
1.  **ausearch**: A powerful tool to query the audit logs without simple grep.
2.  **sealert**: Translates raw, cryptic logs into human-readable suggestions.

## Hands-On Workflow

### 1. Trigger a Denial
(If you did the Phase 2 practice, you already have one. If not, create a file in `/root` and try to move it to `/var/www/html` and access it via Apache).

### 2. Check for Denials
This commands shows all AVC denials from "recent" (last 10 mins).

```bash
sudo ausearch -m AVC -ts recent
```

*Raw Output Example:*
```
type=AVC msg=audit(1611234.123:99): avc:  denied  { getattr } for  pid=1234 comm="httpd" path="/var/www/html/test.html" dev="dm-0" ino=12345 scontext=system_u:system_r:httpd_t:s0 tcontext=unconfined_u:object_r:admin_home_t:s0 tclass=file
```

### 3. Translate with `sealert`
This is the "easy mode" for SELinux debugging. It tells you exactly what happened and suggests fixes.

```bash
sudo sealert -a /var/log/audit/audit.log
```

*Output Example:*
```
SELinux is preventing httpd from getattr access on the file /var/www/html/test.html.

*****  Plugin restorecon (99.5 confidence) suggests   ******************

If you want to fix the label.
/var/www/html/test.html default label should be httpd_sys_content_t.
Then you can run:
restorecon -v /var/www/html/test.html
...
```

### 4. The Decision Matrix
Once you see a denial, you have 3 choices:

| Scenario | Solution |
| :--- | :--- |
| **Labeling Issue** | `restorecon` the file (Common) |
| **Legitimate Feature** | Enable a boolean (Common) |
| **Custom App/Port** | Write a custom policy (Rare, last resort) |

## Guided Exercise: Debugging a Denial
Let's practice the complete "Break, Detect, Fix" cycle.

### 1. Preparation (Break it)
We will create a valid web file, then **intentionally break its label** to simulate a problem.

```bash
# Create a test file
echo "SELinux Debug Lab" | sudo tee /var/www/html/debug_test.html

# Apply a WRONG label (admin_home_t is for /root, not web!)
sudo chcon -t admin_home_t /var/www/html/debug_test.html

# Verify it's broken
curl -I localhost/debug_test.html
# Expect: HTTP/1.1 403 Forbidden
```

### 2. Detect (ausearch)
The web server gave us a 403, but was it permissions (DAC) or SELinux (MAC)?

```bash
# Check for denials in the last 10 minutes
sudo ausearch -m AVC -ts recent
```
*Look for:* `denied { getattr }` or `denied { open }` for `comm="httpd"`.

### 3. Analyze (sealert)
Let's get the specific fix suggestions.

```bash
sudo sealert -a /var/log/audit/audit.log
```
*Output Analysis:*
The tool will likely give you a **Confidence Score** (often ~99%).
It will say: *"SELinux is preventing httpd from access..."*
It will suggest: *"If you want to fix the label, run restorecon..."*

### 4. Fix (Without Disabling)
Follow the tool's advice to restore the correct default label.

```bash
# Restore specific file
sudo restorecon -v /var/www/html/debug_test.html

# Verify new label
ls -Z /var/www/html/debug_test.html
# Should be: httpd_sys_content_t

# Verify access
curl -I localhost/debug_test.html
# Expect: HTTP/1.1 200 OK
```
