# Phase 8: Build Real Projects (Mandatory)

> "This is how you actually *learn* eBPF."

Tutorials are over. It's time to build.
Pick **ONE** of these projects and build it. Do not just read about it.

---


## Table of Contents
* [Level 1: The "Auditor" (Recommended Start)](#level-1-the-auditor-recommended-start)
* [Level 2: The "Profiler"](#level-2-the-profiler)
* [Level 3: The "Firewall" (Networking)](#level-3-the-firewall-networking)
* [Level 4: The "Hidden File Detector" (Security)](#level-4-the-hidden-file-detector-security)
* [Level 5: The "LSM Enforcer" (Guru Mode)](#level-5-the-lsm-enforcer-guru-mode)

---

## Level 1: The "Auditor" (Recommended Start)
**Goal:** Log every time a user runs `sudo`.
*   **Hook:** `sys_execve`.
*   **Logic:** Check if the `uid` is 0 (root) OR if the command line is `sudo ...`.
*   **Output:** `[ALERT] User 1000 ran 'sudo rm -rf /'`
*   **Why:** Teaches you string parsing, user context, and basic logging.

---

## Level 2: The "Profiler"
**Goal:** Build a poor-man's CPU profiler.
*   **Hook:** `perf_event` (Sampling frequency: 49Hz).
*   **Logic:** Every time it fires, record the current Instruction Pointer (IP) in a specialized Map (Stack Trace Map).
*   **Output:** A Flame Graph compatible text file.
*   **Why:** Teaches you stack walking and high-frequency event handling.

---

## Level 3: The "Firewall" (Networking)
**Goal:** Drop packets from a specific "bad" IP address.
*   **Hook:** `XDP` (Attach to `lo` or `eth0`).
*   **Logic:** Look up the Source IP in a `BPF_MAP_TYPE_HASH`. If found, `return XDP_DROP`.
*   **User Space:** A Python script to add/remove IPs from the Ban List map.
*   **Why:** Teaches you XDP, map interactions from user space, and raw packet handling.

---

## Level 4: The "Hidden File Detector" (Security)
**Goal:** Detect when a rootkit tries to hide a process.
*   **Hook:** `getdents64` (Directory listing syscall) AND `openat`.
*   **Logic:**
    1.  Maintain a list of active PIDs in a Map (via `fork`/`exec` hooks).
    2.  When `ps` or `ls /proc` runs (`getdents64`), trace what it returns.
    3.  If a PID exists in *reality* (Map) but is missing from the *directory listing*, ALERT.
*   **Why:** This is advanced malware detection logic.

---

## Level 5: The "LSM Enforcer" (Guru Mode)
**Goal:** Prevent `npm` from connecting to the internet.
*   **Hook:** `lsm/socket_connect`.
*   **Logic:**
    1.  Check `current->comm` (Process Name).
    2.  If it equals "npm", `return -EPERM`.
*   **Why:** This is how you implement "Zero Trust" policies for supply chain security.

---

## What Now?

1.  **Pick a Project.**
2.  **Initialize a Repo:**
    ```bash
    mkdir ebpf-project
    cd ebpf-project
    # Use libbpf-bootstrap or similar starter
    ```
3.  **Code.**

If you get stuck, I am here. Tell me which one you chose, and I will help you design the `structs` and maps.
