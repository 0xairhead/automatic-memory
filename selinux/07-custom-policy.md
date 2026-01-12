# Phase 6: Writing Custom SELinux Policy (Week 6)


## Table of Contents

- [Concepts](#concepts)
    - [The Golden Rule](#the-golden-rule)
    - [Tools](#tools)
- [Hands-On Exercises](#hands-on-exercises)
    - [1. Generating Policy with `audit2allow`](#1-generating-policy-with-audit2allow)
- [Advanced: Writing Hand-Coded Policy](#advanced-writing-hand-coded-policy)
- [Key Takeaway](#key-takeaway)

---

## Concepts
Sometimes, `semanage` and booleans aren't enough. If you have a custom application doing something unique, you might need to create a **Type Enforcement (TE) module**.

### The Golden Rule
> **Custom policy should be your last resort.**
> Check for Booleans first. Check for correct Labeling second. Only write policy if it's truly a new, legitimate access requirement.

### Tools
*   `audit2allow`: Reads audit logs and generates allow rules.
*   `checkmodule`: Checks if your policy file is valid.
*   `semodule`: Installs the compiled policy package.

## Hands-On Exercises

### 1. Generating Policy with `audit2allow`
Let's assume you have a denial in your audit log (e.g., your custom app `myapp` trying to access `/var/lib/myapp`).

**Step 1: Check the denials**
```bash
sudo ausearch -m AVC -ts recent
```

**Step 2: Dry Run - See what simple policy would look like**
```bash
sudo ausearch -m AVC -ts recent | audit2allow
```
*Output will look like:*
```
allow httpd_t some_type_t:file read;
```

**Step 3: Generate the Module Package**
This creates two files: `myapp_fix.te` (Type Enforcement source) and `myapp_fix.pp` (Policy Package binary).
```bash
sudo ausearch -m AVC -ts recent | audit2allow -M myapp_fix
```

**Step 4: Inspect the Source**
**CRITICAL**: Always read the `.te` file. `audit2allow` is "dumb"—it allows *everything* that was blocked. If your app was hacked and trying to read `/etc/shadow`, `audit2allow` effectively blindly allows it.
```bash
cat myapp_fix.te
```

**Step 5: Load the Module**
```bash
sudo semodule -i myapp_fix.pp
```

## Advanced: Writing Hand-Coded Policy
For production, you often write `.te` files manually to be precise.

```selinux
module myapp 1.0;

require {
    type httpd_t;
    type var_log_t;
    class file { read write };
}

# Allow httpd to read/write var_log_t (just an example)
allow httpd_t var_log_t:file { read write };
```

## Key Takeaway
> **`audit2allow` is a double-edged sword.** It fixes the problem, but validates that "whatever the app tried to do is okay." **Review every line.**
