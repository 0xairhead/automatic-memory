# Phase 3: eBPF Maps (The Brains)

> "Maps are how eBPF becomes powerful."

Without Maps, eBPF programs are just fire-and-forget triggers. **Maps** allow your programs to have *memory*, keep *state*, and *communicate* with user space.

---


## Table of Contents
* [1. What are Maps?](#1-what-are-maps)
* [2. Hands-On: Counting Execs per User](#2-hands-on-counting-execs-per-user)
* [3. User Space Side (Conceptual)](#3-user-space-side-conceptual)
* [4. Theory: Perf Events (Ring Buffer)](#4-theory-perf-events-ring-buffer)

---

## 1. What are Maps?

Think of a Map as a shared Key-Value store that lives in the kernel.
*   **eBPF Program (Kernel):** Can Read/Write.
*   **User App (Python/Go/C):** Can Read/Write (usually Reads metrics).

### Crucial Map Types
1.  `BPF_MAP_TYPE_HASH`: A standard hash map. Good for sparse data (e.g., "Packets per IP address").
2.  `BPF_MAP_TYPE_ARRAY`: An array. Super fast, but fixed size. Good for global config or simple counters.
3.  `BPF_MAP_TYPE_RINGBUF`: A high-performance circular buffer. Use this to send *events* (like logs) to user space efficiently. Avoid `bpf_trace_printk`!

---

## 2. Hands-On: Counting Execs per User

Let's build a program that counts how many times each User ID (UID) runs a command.

### Kernel Side (`counter.bpf.c`)

```c
#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>

// 1. Define the Map
struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __uint(max_entries, 1024);
    __type(key, __u32);   // Key: UID
    __type(value, __u64); // Value: Count
} exec_counts SEC(".maps");

SEC("tracepoint/syscalls/sys_enter_execve")
int count_execs(void *ctx) {
    __u32 uid;
    __u64 *val, init_val = 1;

    // Get current User ID
    uid = bpf_get_current_uid_gid();

    // Look up the map
    val = bpf_map_lookup_elem(&exec_counts, &uid);
    if (!val) {
        // Not found? Initialize it.
        bpf_map_update_elem(&exec_counts, &uid, &init_val, BPF_ANY);
    } else {
        // Found? Increment it atomically.
        __sync_fetch_and_add(val, 1);
    }

    return 0;
}

char LICENSE[] SEC("license") = "GPL";
```

### Explaining the Code
*   `SEC(".maps")`: Defines the map structure so `libbpf` knows how to create it.
*   `bpf_get_current_uid_gid()`: Helper to get the UID of the process triggering the event.
*   `bpf_map_lookup_elem()`: Returns a pointer to the value in the map. **Always check for NULL.**
*   `__sync_fetch_and_add()`: We must use atomic operations because this function might run on multiple CPUs simultaneously!

---

## 3. User Space Side (Conceptual)

In Python (using BCC) or raw C (`libbpf`), you would:
1.  Load the program.
2.  Attach to `tracepoint`.
3.  Every 1 Second: Iterate through the map and print:
    *   UID 1000: 5 execs
    *   UID 0 (Root): 12 execs

### Quick Check with `bpftool`

You don't *need* to write a user space app right now. `bpftool` can dump maps for you.

1.  **Compile & Load (like before):**
    ```bash
    clang -O2 -target bpf -c counter.bpf.c -o counter.bpf.o
    sudo bpftool prog load counter.bpf.o /sys/fs/bpf/counter
    sudo bpftool prog attach /sys/fs/bpf/counter
    ```

2.  **Generate Traffic:**
    Open a new terminal and type `ls`, `id`, `whoami` a few times.

3.  **Read the Map:**
    ```bash
    # Find the map ID
    sudo bpftool map list
    # Dump the map content
    sudo bpftool map dump name exec_counts
    ```

    *Result:*
    ```json
    [{
        "key": 1000,
        "value": 4
    }, {
        "key": 0,
        "value": 1
    }]
    ```

---

## 4. Theory: Perf Events (Ring Buffer)

Maps are for **State** (Counters, Config).
Ring Buffers are for **Events** (Logs, Alerts).

If you want to send a message *"Alert! Suspicious process started!"*, don't use a Hash Map. Use a Ring Buffer. It works like a queue:
*   Kernel Pushes Data -> Ring Buffer -> User Space Polls/Consumes Data.

---

## Next Steps
You have mastered the core mechanics!
*   Phase 1: Hooks
*   Phase 2: Code
*   Phase 3: State (Maps)

Now we move to **Phase 4: Observability**. We will use these maps to visualize latency histograms and track what's slowing down your system.
