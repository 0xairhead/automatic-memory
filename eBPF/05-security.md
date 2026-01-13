# Phase 5: eBPF for Security

> "This is your sweet spot."

eBPF has revolutionized Linux security. It allows us to observe and *block* malicious behavior in real-time without modifying the kernel or the application.

---


## Table of Contents
* [1. Why eBPF for Security?](#1-why-ebpf-for-security)
* [2. Runtime Security (The Falco Model)](#2-runtime-security-the-falco-model)
* [3. Hands-On: Detecting Access to `/etc/shadow`](#3-hands-on-detecting-access-to-etcshadow)
* [4. Container & Kubernetes Awareness](#4-container--kubernetes-awareness)
* [5. The Big Players](#5-the-big-players)

---

## 1. Why eBPF for Security?

Traditional security (Antivirus, LD_PRELOAD) is easy to bypass.
*   **Tamper Proof:** Attackers in user space cannot modify eBPF programs running in the kernel.
*   **Deep Visibility:** You see the *source of truth* (syscalls), not just the logs the application *chose* to write.
*   **Enforcement:** Modern eBPF can kill processes (SIGKILL) or drop packets before the malicious action completes.

---

## 2. Runtime Security (The Falco Model)

The core idea is **Behavioral Monitoring**. We don't look for signatures (file hashes); we look for bad behavior.

### Detecting a "Reverse Shell"
A reverse shell usually looks like:
1.  A network connection (`connect`).
2.  Followed by a shell spawned (`execve`) redirecting stdin/stdout.
3.  Often to a non-standard port.

With eBPF, we trace `sys_connect` and `sys_execve` and correlate them in maps.

---

## 3. Hands-On: Detecting Access to `/etc/shadow`

Let's write a program that alerts whenever someone tries to read the password file.

### The Code (`file_monitor.bpf.c`)

```c
#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_tracing.h>

// Trace file opening
SEC("kprobe/do_sys_openat2")
int BPF_KPROBE(detect_shadow, int dfd, const char *filename) {
    char p[20];
    
    // Read the filename string from kernel memory
    long len = bpf_probe_read_kernel_str(&p, sizeof(p), filename);
    
    // Check if it matches "/etc/shadow"
    // Use a simple naive check for demo purposes (real apps use Maps for prefixes)
    char shadow[] = "/etc/shadow";
    
    // Compare string (manual loop needed in older kernels, simplified here)
    // Note: String comparison in eBPF is tricky/limited.
    // In production, we send the filename to User Space to check.
    
    if (len > 0) {
        bpf_printk("File Access: %s\n", p);
    }
    return 0;
}

char LICENSE[] SEC("license") = "GPL";
```
*Note: String matching in eBPF is notoriously hard due to verifier limits. The industry pattern is: Send the event to user space (Perf Buffer) -> User space checks the string.*

---

## 4. Container & Kubernetes Awareness

eBPF sees everything, but it sees it as PIDs (Process IDs).
In Kubernetes, we care about **Pods** and **Namespaces**.

### Making eBPF "Container-Aware"
How do we know PID 1234 belongs to the "Payment Service"?
1.  **Cgroups:** Every container runs in a cgroup. eBPF can read the cgroup ID (`bpf_get_current_cgroup_id()`).
2.  **Correlation:** The user space agent (Falco/Tetragon) queries the Docker/containerd socket: "Hey, what container has cgroup ID X?" -> "Oh, that's the Payment Pod."

---

## 5. The Big Players

*   **Falco:** The standard. Uses eBPF to send syscalls to user space. A Rules Engine (YAML) checks for patterns.
    *   *Rule:* "Terminal shell in container"
*   **Tracee (Aqua Security):** Pure eBPF tracing. Very strong on forensics.
*   **Tetragon (Cilium/Isovalent):** **Inline Enforcement.**
    *   Unlike Falco (which alerts), Tetragon can *block* the syscall in the kernel. If a hacker runs `curl evil.com`, Tetragon kills the process *before the network packet leaves*.

---

## Assignment: Install Falco (The Easy Way)

1.  **Install Falco** in your Linux VM.
2.  **Trigger a Rule:**
    ```bash
    # Writing to a binary directory usually triggers an alert
    sudo touch /bin/hack
    ```
3.  **Check Logs:**
    ```bash
    tail -f /var/log/syslog | grep falco
    ```

---

## Next Steps
We've covered Observability and Security. Now, let's look at the infrastructure that runs the world.
**Phase 6: eBPF for DevOps & Cloud** (Networking, Load Balancing, and replacing iptables).
