# Phase 7: Advanced eBPF

> "Where most people quit &mdash; you won’t."

This phase is about hitting the limits of the technology and knowing how to break through them.

---


## Table of Contents
* [1. The Verifier Deep Dive](#1-the-verifier-deep-dive)
* [2. Performance Engineering](#2-performance-engineering)
* [3. CO-RE (Compile Once, Run Everywhere)](#3-co-re-compile-once-run-everywhere)
* [4. eBPF LSM (Linux Security Modules)](#4-ebpf-lsm-linux-security-modules)
* [5. Emerging Tech: eBPF + WASM](#5-emerging-tech-ebpf--wasm)

---

## 1. The Verifier Deep Dive

The Verifier is a static analyzer that proves your program is safe *without* running it. It is conservative. It will reject safe code if it cannot *prove* it is safe.

### The Limits (and how to bypass them)
1.  **Instruction Limit:** Traditionally 4096 instructions. Modern kernels (5.2+) allow 1 million.
    *   *Solution:* **Tail Calls**. Chain multiple eBPF programs together. Prog A calls Prog B. The instruction counter resets.
2.  **Stack Size:** Limited to 512 bytes.
    *   *Problem:* You can't declare `char large_buffer[1024]`.
    *   *Solution:* Use a **Per-CPU Array Map** as scratch space. Lookup the map, get a pointer, and use that as your buffer.
3.  **Loops:** Historically forbidden. Modern kernels (5.3+) support Bounded Loops (`for` loops where the max iteration count is constant).
    *   *Trick:* `#pragma unroll`. Force the compiler to unroll the loop if the verifier is being stubborn.

---

## 2. Performance Engineering

eBPF is fast, but you can make it slow.

### A. Map Contention
*   **Bad:** A global Hash Map updated by all CPUs. `__sync_fetch_and_add` locks the cache line.
*   **Good:** **Per-CPU Maps**. Each CPU has its own private slice of the map. No locking. User space sums them up later.

### B. Size Matters
*   `bpf_probe_read()` is expensive. Reading 1 byte is cheaper than 100 bytes.
*   Don't copy the whole packet payload if you only need the IP header.

---

## 3. CO-RE (Compile Once, Run Everywhere)

Before CO-RE, you had to compile your C code *on the target machine* because `struct task_struct` changed shape in every kernel version.

### The Magic of BTF (BPF Type Format)
*   **Relocation:** The `libbpf` loader reads the running kernel's BTF.
*   **Adaptation:** If `task_struct->pid` moved from offset 16 to offset 20, `libbpf` patches your bytecode *in memory* before loading it.
*   **Result:** You ship a single binary. It runs on Ubuntu, Fedora, and Arch.

---

## 4. eBPF LSM (Linux Security Modules)
This is the new frontier (Kernel 5.7+).

*   **Old Way:** Trace `openat` and kill the process. (Race condition: The file might already be open).
*   **LSM Way:** eBPF hooks into the *LSM hooks* (same place SELinux/AppArmor live).
*   **Power:** You can return `-EPERM` to *deny* the action safely before it happens.
    ```c
    SEC("lsm/file_open")
    int BPF_PROG(restrict_files, struct file *file) {
        if (is_protected(file))
            return -EPERM; // Access Denied!
        return 0;
    }
    ```

---

## 5. Emerging Tech: eBPF + WASM

WebAssembly (WASM) is gaining traction as a userspace sandbox.
*   **Idea:** Write logic in WASM. WASM calls eBPF functions.
*   **Why?** WASM is easier to distribute and sign than raw ELF binaries. Projects like `bumblebee` are exploring this.

---

## Assignment: The " impossible" Program

1.  **Challenge:** Write a program that traces `sys_execve` but needs a 4KB buffer (larger than the stack).
2.  **Solution:**
    *   Create a `BPF_MAP_TYPE_PERCPU_ARRAY`.
    *   In the program, `map_lookup_elem` index 0.
    *   Use that returned pointer as your scratch buffer.

---

## Next Steps
You have reached the peak of the mountain. You understand the Internals, the Performance hacks, and the Future (LSM).

**Phase 8: Build Real Projects** is the victory lap. I will give you 5 project briefs. You pick one, and you build it.
