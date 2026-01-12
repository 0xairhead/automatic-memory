# SELinux Lab Environment Setup


This guide helps you create the necessary environment to practice SELinux. You can run these labs on maintaining **macOS (via Lima)** or a **Native Fedora/RHEL** system.

---

## Table of Contents

- [Part 1: macOS Users (Recommended: Lima)](#part-1-macos-users-recommended-lima)
    - [1. Installation](#1-installation)
    - [2. Start the VM](#2-start-the-vm)
    - [3. Verification](#3-verification)
    - [4. Connect](#4-connect)
- [Part 2: Native Linux Users (Fedora / RHEL / CentOS)](#part-2-native-linux-users-fedora--rhel--centos)
- [Part 3: Common Provisioning (REQUIRED)](#part-3-common-provisioning-required)
- [Part 4: The macOS Workflow (Lima Only)](#part-4-the-macos-workflow-lima-only)
- [Part 5: Lima Cheat Sheet](#part-5-lima-cheat-sheet)
- [Part 6: SELinux Verification](#part-6-selinux-verification)
- [Troubleshooting (Lima)](#troubleshooting-lima)

---

## Part 1: macOS Users (Recommended: Lima)

[Lima](https://github.com/lima-vm/lima) ("Linux on Mac") is the most efficient way to get a running RHEL-like system on your Mac.

### 1. Installation
Using Homebrew:
```bash
brew install lima
```

### 2. Start the VM
We use **Fedora** because it is the upstream for RHEL and has robust SELinux support.
```bash
limactl start template://fedora
```
*   **Prompt**: If asked to "Open an editor to review the configuration", press `Enter` to accept defaults.

### 3. Verification
```bash
limactl list
# Should show 'fedora' as 'Running'
```

### 4. Connect
```bash
limactl shell fedora
```

---

## Part 2: Native Linux Users (Fedora / RHEL / CentOS)

If you are already running Linux (Fedora Workstation, RHEL 9, AlmaLinux, Rocky):

1.  **Ensure SELinux is Enforcing**:
    ```bash
    sestatus
    # Must say: Current mode: enforcing
    ```
    If it says `disabled`, edit `/etc/selinux/config` to set `SELINUX=enforcing` and **reboot**.

2.  **Work Directory**:
    Create a directory for your labs:
    ```bash
    mkdir -p ~/selinux-labs
    cd ~/selinux-labs
    ```

---

## Part 3: Common Provisioning (REQUIRED)

**Whether you are using Lima or Native Fedora**, you must install the tools required for the labs (Apache, Audit tools, Policy tools).

Run the following commands inside your Linux terminal:

```bash
# 1. Update system
sudo dnf update -y

# 2. Install Lab Dependencies
# httpd: Web server for webroot labs
# setroubleshoot-server: For 'sealert'
# policycoreutils-python-utils: For 'semanage', 'audit2allow'
# setools-console: For 'seinfo', 'sesearch'
# checkpolicy: For compiling custom modules
sudo dnf install -y httpd \
    setroubleshoot-server \
    policycoreutils-python-utils \
    setools-console \
    audit \
    checkpolicy \
    bash-completion

# 3. Enable Apache (for web labs)
sudo systemctl enable --now httpd
```

---

## Part 4: The macOS Workflow (Lima Only)

**Do not write code inside the terminal.**

1.  **Edit on Mac**: Open your `selinux` folder in VS Code *on macOS*.
2.  **Run on Linux**: Keep a terminal window open, shelled into Lima.
    *   `limactl shell fedora`
    *   `cd /Users/phoenix/Desktop/selinux` (Lima creates this path automatically).

---

## Part 5: Lima Cheat Sheet

| Action | Command | Description |
| :--- | :--- | :--- |
| **Start** | `limactl start <name>` | Boots the VM. |
| **Stop** | `limactl stop <name>` | Shuts down the VM. |
| **Delete** | `limactl delete <name>` | Destroys the VM (fixes broken SELinux states). |
| **Shell** | `limactl shell <name>` | Opens a terminal inside the VM. |

---

## Part 6: SELinux Verification

Run this to confirm you are ready to start **Lab 1**:

```bash
# 1. Check Mode
getenforce
# Output: Enforcing

# 2. Check Tools
which semanage
# Output: /usr/sbin/semanage

# 3. Check Web Server
systemctl is-active httpd
# Output: active
```

---

## Troubleshooting (Lima)

**"Permission Denied in the Shared Folder"**
Lima mounts your Mac home directory. Sometimes root cannot write to it.
*   **Fix**: Only *read* lab instructions/scripts from the shared folder.
*   **Action**: If a lab requires creating a system file (e.g., custom policy), do it in `/tmp` or `~` inside the VM, not in the shared `/Users/...` folder.
