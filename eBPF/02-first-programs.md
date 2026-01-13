# Phase 2: First Programs (Hello World)

## Table of Contents
* [1. The "Easy Mode": `bpftrace`](#1-the-easy-mode-bpftrace)
* [2. The "Pro Mode": C Code Explained](#2-the-pro-mode-c-code-explained)

---

## 1. The "Easy Mode": `bpftrace`

Before we write complex code, let's just feel the power.
`bpftrace` is a tool that lets you write "One-Liners". It handles all the compiling and loading for you.

### Try this (if you have Linux):
```bash
sudo bpftrace -e 'tracepoint:syscalls:sys_enter_execve { printf("New Process: %s\n", comm); }'
```

### Breaking it down:
1.  **`tracepoint:syscalls:sys_enter_execve`**: This is the **Hook**. We are saying "Wait for the `execve` system call (New Process) to start".
2.  **`{ ... }`**: This is the **Action**. Run this code when the hook fires.
3.  **`comm`**: This is a built-in variable that holds the **Command Name** (e.g., "ls", "curl").
4.  **`printf`**: Print it to the screen.

**That's it.** You just intercepted a kernel event and printed data.

---

## 2. The "Pro Mode": C Code Explained

Real production tools (like Cilium or Falco) use C because it's more robust.
Here is the exact same program in C, explained line-by-line.

```c
// Import standard kernel types
#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>

// 1. SEC(...)
// This tells the loader WHERE to put this code.
// We want to attach to the 'execve' tracepoint.
SEC("tracepoint/syscalls/sys_enter_execve")

// 2. The Function
// This is the function that runs every time the event happens.
int handle_execve(void *ctx) {
    
    // Define a simple message
    char msg[] = "Hello from eBPF!\n";
    
    // 3. bpf_trace_printk
    // This is like 'printf' for the kernel.
    // It writes the message to a special file: /sys/kernel/debug/tracing/trace_pipe
    bpf_trace_printk(msg, sizeof(msg));
    
    return 0; // Always return 0 (Success)
}

// 4. License
// You MUST declare a license. 
// "GPL" gives you access to all helper functions.
char LICENSE[] SEC("license") = "GPL";
```

### How to Run This (Conceptual)
1.  **Compile** this C file into an object file (`.o`).
2.  **Use a Loader Tool** (like `bpftool`) to send that `.o` file to the kernel.
3.  **Read the logs** (`cat /sys/kernel/debug/tracing/trace_pipe`).

---

### You did it!
You now understand the two ways to use eBPF:
1.  **Scripting (`bpftrace`)**: For quick debugging.
2.  **Programming (C)**: For building permanent tools.

Next, in **Phase 3**, we will learn how to send *structured data* (like numbers and lists) instead of just text messages.
