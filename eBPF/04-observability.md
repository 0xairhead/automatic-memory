# Phase 4: Observability with eBPF

> "This is where eBPF shines."

Observability is about **answering unknown unknowns**. Why is the database slow? Why did that request timeout? eBPF gives you deep, low-overhead visibility without changing your application code.

---


## Table of Contents
* [1. Metrics vs. Logs vs. Traces](#1-metrics-vs-logs-vs-traces)
* [2. Using `bpftrace` for Instant Observability](#2-using-bpftrace-for-instant-observability)
* [3. Network Observability (TCP/IP)](#3-network-observability-tcpip)
* [4. Modern Observability Tools](#4-modern-observability-tools)

---

## 1. Metrics vs. Logs vs. Traces

eBPF can generate all three, but **Metrics** (Histograms) and **Traces** (Spans) are the strongest use cases.

### The Power of Histograms
Averages are lies.
*   **Average:** "APIs respond in 50ms."
*   **P99 (Histogram):** "99% of requests are fast, but 1% take 10 seconds."

eBPF helps us build histograms efficiently *in the kernel* to avoid sending millions of events to user space.

---

## 2. Using `bpftrace` for Instant Observability

You don't always need to write C. `bpftrace` is the industry standard for ad-hoc debugging.

### Scenario: "My specific function is slow."

Let's say you want to measure how long the `do_sys_open` function takes in the kernel (opening files).

```bash
# Measure latency of 'do_sys_open'
sudo bpftrace -e '
kprobe:do_sys_open
{
    @start[tid] = nsecs;
}

kretprobe:do_sys_open
/@start[tid]/
{
    $duration = nsecs - @start[tid];
    @us = hist($duration / 1000);
    delete(@start[tid]);
}
'
```

*   **@start[tid]:** A map storing the start time keyed by Thread ID.
*   **kretprobe:** Fires when the function *returns*.
*   **hist():** A helper that automatically buckets the data into a power-of-2 histogram.

**Output:**
```text
@us:
[0, 1]                 0 |                                                    |
[2, 4]               100 |@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@                      |
[4, 8]               300 |@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@|
[8, 16]               50 |@@@@@@@@@                                           |
```
*This tells you exactly how fast file opens are happening, with zero application changes.*

---

## 3. Network Observability (TCP/IP)

You can trace individual packet drops, TCP retransmits, or connection latency.

*   **TCP Connect Latency:** Hook `tcp_v4_connect` and `tcp_rcv_state_process`.
*   **Packet Drops:** Hook `kfree_skb`. This is magical—it tells you *where* and *why* the kernel dropped a packet.

---

## 4. Modern Observability Tools

Don't reinvent the wheel. These open-source tools use everything we just learned:

| Tool | Focus | Under the Hood |
| :--- | :--- | :--- |
| **Pixie** | Auto-telemetry (HTTP/DNS/SQL traces) | Uses eBPF to parse protocol traffic (L7) directly in the kernel/userspace buffer. |
| **Cilium Hubble** | Kubernetes Network Visibility | Uses XDP and TC (Traffic Control) hooks to map flows to Pods. |
| **Parca** | Continuous Profiling | Uses eBPF to walk stack traces (`perf_event_open`) to find CPU hotspots. |
| **BCC Tools** | Scripting | A collection of Python scripts (`execsnoop`, `biosnoop`) that pre-date reliable libbpf. |

---

## Assignment: Profile Your System

1.  **Install `bpftrace`** (if you haven't).
2.  **Run `execsnoop`** (or writing the one-liner from Phase 2).
3.  **Run `opensnoop`** (traces file opens).
4.  **Visualize:** Use the histogram example above on a function you care about (or just `vfs_read`).

---

## Next Steps
We have data. Now let's protect it.
**Phase 5: eBPF for Security** is where we start detecting hackers, rootkits, and bad behavior.
