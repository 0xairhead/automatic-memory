# SELinux Practical Labs

This file contains consolidated labs to test your skills. Each lab simulates a real-world scenario.

---


## Table of Contents

- [Lab 1: The "Moved Webroot" Scenario](#lab-1-the-moved-webroot-scenario)
- [Lab 2: The SSH Port Switch](#lab-2-the-ssh-port-switch)
- [Lab 3: The "Broken" Container Volume](#lab-3-the-broken-container-volume)
- [Lab 4: Boolean Toggles](#lab-4-boolean-toggles)
- [Lab 5: The "Forbidden" Cron Job (Advanced)](#lab-5-the-forbidden-cron-job-advanced)
- [Lab 6: The "Remote DB" Connection](#lab-6-the-remote-db-connection)
- [Lab 7: Custom Log File Rotation](#lab-7-custom-log-file-rotation)
- [Lab 8: The Forgotten `restorecon`](#lab-8-the-forgotten-restorecon)
- [Lab 9: User Confinement (RBAC) - Advanced](#lab-9-user-confinement-rbac---advanced)
- [Lab 10: Manual Policy Writing (Expert)](#lab-10-manual-policy-writing-expert)

---

## Lab 1: The "Moved Webroot" Scenario
**Goal**: Host a website from a non-standard directory (`/srv/web`) and make it work with SELinux.

### Problem Setup
1.  **Setup**:
    ```bash
    sudo mkdir -p /srv/web
    echo "<h1>Custom Webroot</h1>" | sudo tee /srv/web/index.html
    ```
2.  **Configure Apache**:
    *   Edit `/etc/httpd/conf/httpd.conf`.
    *   Change `DocumentRoot "/var/www/html"` to `DocumentRoot "/srv/web"`.
    *   Change `<Directory "/var/www/html">` to `<Directory "/srv/web">`.
3.  **Start Service**: `sudo systemctl restart httpd`
4.  **Test**: `curl localhost`
    *   *Result*: **403 Forbidden**.

<details>
<summary><b>Click to reveal the Solution (Debug & Fix)</b></summary>

### Solution
1. **Debug**:
    *   Run `ausearch -m AVC -ts recent`.
    *   Run `ls -Z /srv/web`.
2. **Fix**:
    *   **Do NOT use `chcon`**.
    *   Use `semanage fcontext -a -t httpd_sys_content_t "/srv/web(/.*)?"`.
    *   Run `restorecon -Rv /srv/web`.
3. **Verify**: `curl localhost` should return the HTML.
</details>

---

## Lab 2: The SSH Port Switch
**Goal**: Run SSH on port 2222.

### Problem Setup
1.  **Configure SSH**:
    *   Edit `/etc/ssh/sshd_config`.
    *   Change `Port 22` to `Port 2222`.
2.  **Restart SSH**: `sudo systemctl restart sshd`
    *   *Result*: **FAILED**. Systemd will show it failed to bind to the port.

<details>
<summary><b>Click to reveal the Solution (Debug & Fix)</b></summary>

### Solution
1. **Debug**:
    *   `journalctl -u sshd` will say "Permission denied".
    *   `sealert -a /var/log/audit/audit.log` will mention `ssh_port_t`.
2. **Fix**:
    ```bash
    sudo semanage port -a -t ssh_port_t -p tcp 2222
    ```
3. **Verify**: Restart sshd. It should start.
</details>

---

## Lab 3: The "Broken" Container Volume
**Goal**: Share host data with a container correctly.

### Problem Setup
1.  **Setup**:
    ```bash
    mkdir ~/container_data
    echo "Secret Data" > ~/container_data/secret.txt
    ```
2.  **Run (Fail)**:
    ```bash
    podman run --rm -v ~/container_data:/data alpine cat /data/secret.txt
    ```
    *   *Result*: "Permission denied".

<details>
<summary><b>Click to reveal the Solution (Debug & Fix)</b></summary>

### Solution
1. **Fix**: Use the `:Z` flag.
    ```bash
    podman run --rm -v ~/container_data:/data:Z alpine cat /data/secret.txt
    ```
2. **Inspect**:
    *   Check `ls -Z ~/container_data/secret.txt` after the run. It changed!
</details>

---

## Lab 4: Boolean Toggles
**Goal**: Allow a web server to send emails.

### Problem Setup
1.  **Scenario**: Your PHP script tries to send an email using `sendmail`.
2.  **Check Status**:
    ```bash
    getsebool httpd_can_sendmail
    ```
    *   *Likely Output*: `off`.

<details>
<summary><b>Click to reveal the Solution (Fix)</b></summary>

### Solution
1. **Enable it**:
    ```bash
    sudo setsebool -P httpd_can_sendmail 1
    ```
</details>

---

## Lab 5: The "Forbidden" Cron Job (Advanced)
**Goal**: Create a cron job that writes to a user's home directory (and fix the denial).

### Problem Setup
1.  **Setup**: Create a script `/usr/local/bin/backup.sh` that writes to `/home/user/backup.log`.
2.  **Schedule**: Add it to system cron `/etc/cron.d/backup`.
3.  **Observation**: It fails to write. `cron` runs in a specific domain that might not have write access to generic user home files depending on policy.

<details>
<summary><b>Click to reveal the Solution (Hint)</b></summary>

### Solution Hint
4. **Fix**: Investigate `cron_t` and file labels.
    *   System cron jobs often run as `system_cronjob_t`.
    *   User home directories are `user_home_t`.
    *   You might need to use a directory labeled `backup_store_t` or similar, or adjust booleans related to cron and user homes.
</details>

---

## Lab 6: The "Remote DB" Connection
**Goal**: Allow your web server to talk to a remote database.

### Problem Setup
1.  **Scenario**: You have a WordPress site trying to connect to a remote MySQL server (not localhost).
2.  **Failure**: Connection times out or "Permission Denied" in the app logs.

<details>
<summary><b>Click to reveal the Solution (Debug & Fix)</b></summary>

### Solution
1. **Debug**:
    *   `ausearch -m AVC` shows `name_connect` denial to a MySQL port (3306).
2. **Fix**:
    *   This is a common boolean.
    ```bash
    # Check
    getsebool -a | grep http | grep db
    # Set
    sudo setsebool -P httpd_can_network_connect_db 1
    ```
</details>

---

## Lab 7: Custom Log File Rotation
**Goal**: Store application logs in a custom directory and ensure `logrotate` works.

### Problem Setup
1.  **Setup**:
    *   App writes to `/opt/myapp/logs/app.log`.
    *   Directory context is likely `etc_t` or `usr_t` (inherited from `/opt`).
2.  **The Issue**:
    *   `logrotate` runs as `logrotate_t`.
    *   It tries to rename/create files in `/opt/myapp/logs`.
    *   **Denied!** `logrotate_t` can write to `var_log_t`, not generic files.

<details>
<summary><b>Click to reveal the Solution (Fix)</b></summary>

### Solution
1. **Fix**:
    ```bash
    sudo semanage fcontext -a -t var_log_t "/opt/myapp/logs(/.*)?"
    sudo restorecon -Rv /opt/myapp/logs
    ```
</details>

---

## Lab 8: The Forgotten `restorecon`
**Goal**: Understand why `semanage` alone isn't enough.

### Problem Setup
1.  **Setup**: `mkdir /testing`. `ls -Z /testing` -> `default_t`.
2.  **Apply Policy**:
    ```bash
    sudo semanage fcontext -a -t httpd_sys_content_t "/testing(/.*)?"
    ```
3.  **Check**: `ls -Z /testing`.
    *   *Result*: It is **STILL** `default_t`.
    *   *Why?* `semanage` only updates the database, not the filesystem.

<details>
<summary><b>Click to reveal the Solution (Fix)</b></summary>

### Solution
1. **Apply to Disk**:
    ```bash
    sudo restorecon -v /testing
    ls -Z /testing
    ```
    *   *Result*: Now it is `httpd_sys_content_t`.
</details>

---

## Lab 9: User Confinement (RBAC) - Advanced
**Goal**: Create a Linux user that is locked down by SELinux (cannot `sudo`, cannot exec in `/tmp`).

### Problem Setup
1.  **Create User**: `sudo useradd restricted_user`

<details>
<summary><b>Click to reveal the Solution (Steps)</b></summary>

### Solution
2. **Map to SELinux User**:
    *   By default, users are `unconfined_u`. We will map them to `user_u`.
    *   `user_u` is a strict targeted policy: no networking, no executing in home/tmp, no SUID.
    ```bash
    sudo semanage login -a -s user_u restricted_user
    ```
3. **Verify**:
    ```bash
    sudo semanage login -l
    ```
4. **Test**:
    *   Login as the user: `su - restricted_user`
    *   Try to run `sudo ls`: **Denied**.
    *   Copy `/bin/ls` to `/tmp` and try to run it: **Denied** (no exec in tmp).
    *   Check context: `id -Z` (should show `user_u:user_r:user_t:s0`).
</details>

---

## Lab 10: Manual Policy Writing (Expert)
**Goal**: Write a Custom Type Enforcement (.te) file from scratch to allow a specific interaction.

### Problem Setup
1.  **Scenario**: A custom backup script (`backup_t`) needs to read a specific config file type (`myapp_conf_t`).

<details>
<summary><b>Click to reveal the Solution (Policy Code)</b></summary>

### Solution
2. **Create the TE File** (`mybackup.te`):
    ```selinux
    module mybackup 1.0;

    require {
        type backup_t;
        type myapp_conf_t;
        class file { read open getattr };
    }

    # The Rule
    allow backup_t myapp_conf_t:file { read open getattr };
    ```
3. **Compile**:
    ```bash
    checkmodule -M -m -o mybackup.mod mybackup.te
    semodule_package -o mybackup.pp -m mybackup.mod
    ```
4. **Install**:
    ```bash
    sudo semodule -i mybackup.pp
    ```
5. **Verify**:
    ```bash
    semodule -l | grep mybackup
    ```
</details>
