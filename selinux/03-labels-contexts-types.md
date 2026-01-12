# Phase 2: Labels, Contexts & Types (Week 2)


## Table of Contents

- [Concepts](#concepts)
    - [The Security Context](#the-security-context)
    - [Common Types](#common-types)
    - [The "Mismatch" Problem](#the-mismatch-problem)
- [Hands-On Exercises](#hands-on-exercises)
    - [1. View Security Contexts](#1-view-security-contexts)
    - [2. The "Break and Fix" Lab](#2-the-break-and-fix-lab)
- [Key Takeaway](#key-takeaway)

---

## Concepts

### The Security Context
Every process, file, and directory in SELinux has a **security context**. It looks like this:

`user:role:type:level`

Example: `unconfined_u:object_r:httpd_sys_content_t:s0`

1.  **User (`user_u`)**: The SELinux user (mapped to Linux user). rarely changed in Targeted policy.
2.  **Role (`role_r`)**: Role Based Access Control (RBAC). rarely changed for files.
3.  **Type (`type_t`)**: **This is the most important part!** Type Enforcement (TE) rules say which process types can access which file types.
4.  **Level (`s0`)**: For MLS/MCS (sensitivity levels). Usually ignored in basics.

### Common Types
*   `httpd_t`: The process type for the Apache web server.
*   `httpd_sys_content_t`: The file type for web content in `/var/www/html`.
*   `var_log_t`: Log files.
*   `etc_t`: Configuration files.
*   `ssh_home_t`: Files in `~/.ssh`.

### The "Mismatch" Problem
Most SELinux issues are simply **labeling problems**.
*   Process (`httpd_t`) tries to read a file (`admin_home_t`).
*   Policy says: `httpd_t` can only read `httpd_sys_content_t`.
*   Result: **Access Denied**.

## Hands-On Exercises

### 1. View Security Contexts
The `-Z` flag is your friend.

```bash
# View file contexts
ls -Z /var/www/html
# Output: system_u:object_r:httpd_sys_content_t:s0 ...

# View process contexts
ps -ZC httpd
# or
ps auxZ | grep httpd
```

### 2. The "Break and Fix" Lab
This demonstrates a classic labeling issue.

**Step 1: Create a file in a non-standard location**
```bash
# Create a dummy index file in /tmp (or your home dir)
echo "<h1>Moved Content</h1>" > /tmp/index.html
```

**Step 2: Move it to the webroot**
*Crucial*: The `mv` command preserves the original label (`user_tmp_t` or `admin_home_t`).
```bash
sudo mv /tmp/index.html /var/www/html/test.html
```

**Step 3: Check the label**
```bash
ls -Z /var/www/html/test.html
# You will likely see something that isn't httpd_sys_content_t
```

**Step 4: Access it (and fail)**
If you try to `curl localhost/test.html`, you will likely get a 403 Forbidden (check Apache error logs!).

**Step 5: Fix it with `restorecon`**
`restorecon` resets the label to what the system *thinks* it should be (based on policy).
```bash
sudo restorecon -v /var/www/html/test.html
ls -Z /var/www/html/test.html
# Now it should be httpd_sys_content_t
```

**Step 6: Access it (Success!)**
Try `curl` again. It works.

## Key Takeaway
> When files are **moved (`mv`)**, they keep their old label.
> When files are **copied (`cp`)**, they inherit the label of the parent directory.
> **Always check `ls -Z` first.**
