# Phase 6: eBPF for DevOps & Cloud

> "eBPF replaces agents, sidecars, and iptables."

If you work in Platform Engineering or Cloud Native, this is the phase that changes your architecture. We are moving from "User Space Agents" to "Kernel Native Intelligence."

---


## Table of Contents
* [1. The Death of iptables](#1-the-death-of-iptables)
* [2. Networking Layers: XDP vs TC](#2-networking-layers-xdp-vs-tc)
* [3. Kubernetes & CNI (Cilium)](#3-kubernetes--cni-cilium)
* [4. "Sidecarless" Service Mesh](#4-sidecarless-service-mesh)
* [5. Operations: Production Debugging](#5-operations-production-debugging)

---

## 1. The Death of iptables

Kubernetes `kube-proxy` relies heavily on `iptables`.
*   **The Problem:** `iptables` is a linear list of rules. If you have 5,000 services, every packet might check 5,000 rules. O(N).
*   **The eBPF Solution:** eBPF uses **Hash Maps**. Lookup is O(1).
*   **Result:** Companies like Cloudflare and Facebook dropped standard IP routing for eBPF years ago because it is orders of magnitude faster at scale.

---

## 2. Networking Layers: XDP vs TC

There are two main places to hook networking in eBPF:

### A. XDP (eXpress Data Path)
*   **Where:** Runs inside the Network Driver (NIC), *before* the kernel even allocates an `sk_buff` (socket buffer).
*   **Speed:** Millions of packets per second.
*   **Use Cases:** Anti-DDoS (dropping bad traffic instantly), Load Balancing.
*   **Limitation:** It is *so* early you don't have the full packet context (no TCP socket info yet).

### B. TC (Traffic Control)
*   **Where:** Runs in the kernel networking stack (Ingress/Egress).
*   **Speed:** Slower than XDP, but faster than standard userspace.
*   **Use Cases:** Observability, Container Policies, modifications.

---

## 3. Kubernetes & CNI (Cilium)

**Cilium** is the poster child for eBPF in Kubernetes.
It replaces `kube-proxy` entirely.

### How it works:
1.  **Service Map:** Instead of writing iptables rules, Cilium updates an eBPF Map: `{ ServiceIP -> [PodIP_1, PodIP_2] }`.
2.  **Load Balancing:** When a packet hits the NIC, the eBPF program reads the Map, picks a backend Pod, rewrites the destination IP, and forwards it.
3.  **Efficiency:** It avoids "Hairpinning" (context switching between kernel/user space).

---

## 4. "Sidecarless" Service Mesh

Traditional Service Mesh (Istio/Linkerd) injects a Proxy (Envoy) into *every* Pod.
*   **Cost:** Memory usage explodes. Latency increases (2x TCP stack traversal per hop).

**eBPF Mesh:**
*   Moves the logic (mTLS, Retries, Observability) into the Kernel.
*   One eBPF program serves *all* Pods on the Node.
*   **Outcome:** 90% reduction in memory overhead.

---

## 5. Operations: Production Debugging

Imagine a production app is stalling. You cannot restart it to add logging.
eBPF allows you to:
1.  **Dynamic Logging:** Inject a print statement into a live function safely.
2.  **Health Checks:** Monitor if the *app* is actually processing requests, not just if the PID is alive.

---

## Assignment: Explore the Network

1.  **Install `bpftool`** (you have this).
2.  **Check Loaded Network Programs:**
    ```bash
    # Show programs attached to XDP or TC
    sudo bpftool net list
    ```
3.  **Stretch Goal:**
    If you have a Kubernetes cluster (Kind/Minikube), install Cilium and run:
    ```bash
    cilium status --verbose
    ```

---

## Next Steps
We have covered the entire "User" stack.
**Phase 7: Advanced eBPF** is where we look at the Verifier's edge cases, complex loop unrolling, and perform deep performance tuning. This is the "Guru" level.
