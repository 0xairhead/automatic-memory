# Phase 0: Prerequisites (Optional but Helpful)


## Goal
Be comfortable with the Linux internals that SELinux builds upon. Before diving into SELinux, ensure you have a solid grasp of these fundamental concepts.

---

## Table of Contents

- [Goal](#goal)
- [Key Concepts](#key-concepts)
    - [1. Linux Permissions (DAC - Discretionary Access Control)](#1-linux-permissions-dac---discretionary-access-control)
    - [2. Processes, PIDs, and Services](#2-processes-pids-and-services)
    - [3. Filesystems, Sockets, and Ports](#3-filesystems-sockets-and-ports)
    - [4. Basic Troubleshooting Tools](#4-basic-troubleshooting-tools)
- [Hands-On Refresher](#hands-on-refresher)
- [Why This Matters](#why-this-matters)

---

## Key Concepts

### 1. Linux Permissions (DAC - Discretionary Access Control)
Standard Linux permissions are based on user identity and group membership.
- **R/W/X**: Read, Write, Execute permissions.
- **Owner/Group/Others**: The three categories of users permissions apply to.
- **Root**: The superuser who bypasses these checks.

### 2. Processes, PIDs, and Services
- **Process**: A running instance of a program.
- **PID**: Process ID, a unique number for each process.
- **Systemd**: The init system used by most modern distributions (RHEL, Fedora, Rocky) to manage services (daemons).
    *   **SELinux Connection**: Systemd is responsible for launching services and assigning them their initial **SELinux Security Context** (domain). If a service starts with the wrong context, it's often an issue in the systemd unit file or the executable's file label.
    *   **Unit Files**: Configuration files (e.g., `httpd.service`) that define how a service starts.

### 3. Filesystems, Sockets, and Ports
- **Filesystem**: How data is stored (ext4, xfs). SELinux labels are stored as extended attributes on the filesystem.
- **Sockets**: Endpoints for communication.
- **Ports**: Network endpoints (e.g., Port 80 for HTTP).

### 4. Basic Troubleshooting Tools
You should be comfortable using:
- `ps aux`: specific process information.
- `ls -l`: list files with permissions.
- `systemctl status <service>`: check service health.
- `journalctl -xe`: view system logs.
- `strace`: trace system calls (advanced but useful).

## Hands-On Refresher
If you need a quick refresher, run these commands to inspect your system state:

```bash
# Check running processes
ps aux | head

# Check file permissions
ls -la /etc/passwd

# Check listening ports
ss -tunlp

# Check service status (e.g., sshd)
systemctl status sshd
```

## Why This Matters
SELinux adds a layer *on top* of these. If DAC denies access, SELinux isn't even checked. You need to distinguish between a "Permission denied" caused by `chmod` (DAC) and one caused by an AVC denial (MAC/SELinux).
