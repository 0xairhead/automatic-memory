# Phase 5: Managing Contexts Properly (Week 5)


## Table of Contents

- [Concepts](#concepts)
    - [`chcon` vs `semanage`](#chcon-vs-semanage)
    - [Best Practice Workflow](#best-practice-workflow)
- [Hands-On Exercises](#hands-on-exercises)
    - [1. The "Right Way" to Label a Custom Directory](#1-the-right-way-to-label-a-custom-directory)
    - [2. Adding a Custom Port](#2-adding-a-custom-port)
- [Key Takeaway](#key-takeaway)

---

## Concepts
By now you know that incorrect labels break things. But how do you change them *correctly*?

### `chcon` vs `semanage`
This is a critical distinction that trips up beginners.

*   **`chcon` (Change Context)**: Changes the label on the file *right now*.
    *   **Pro**: Fast, easy.
    *   **Con**: NOT persistent. If the file system is relabeled (common during updates or scheduled jobs), your change is lost.
*   **`semanage fcontext`**: Updates the **policy database**.
    *   **Pro**: Permanent. Survives relabeling.
    *   **Con**: Doesn't touch the file header immediately; you must run `restorecon` afterward.

### Best Practice Workflow
**Never use `chcon` for production fixes.** Always use `semanage fcontext`.

## Hands-On Exercises

### 1. The "Right Way" to Label a Custom Directory
Imagine you want to serve web content from `/web` instead of `/var/www/html`.

**Step 1: Check default**
```bash
mkdir /web
touch /web/index.html
ls -Z /web
# Likely: default_t
```

**Step 2: Update Policy Database**
Tell SELinux: "Any file inside `/web` should allow be `httpd_sys_content_t`."
```bash
sudo semanage fcontext -a -t httpd_sys_content_t "/web(/.*)?"
```
*Note: This uses regex. `(/.*)?` means "the directory itself and anything recursively inside it".*

**Step 3: Verify the Policy Entry**
```bash
sudo semanage fcontext -l | grep /web
```

**Step 4: Apply the Changes**
The file on disk hasn't changed yet. Run `restorecon` to match the disk to the policy.
```bash
sudo restorecon -Rv /web
```

**Step 5: Verify**
```bash
ls -Z /web/index.html
# Should be httpd_sys_content_t
```

### 2. Adding a Custom Port
If you run SSH on port 2222, SELinux will block it because `sshd_t` is only allowed on `ssh_port_t` (22).

```bash
# View allowed ports
semanage port -l | grep ssh

# Allow SSH on 2222
sudo semanage port -a -t ssh_port_t -p tcp 2222
```

## Key Takeaway
> **`chcon` is for testing.**
> **`semanage` + `restorecon` is for production.**
