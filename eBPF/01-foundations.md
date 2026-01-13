# Phase 1: Foundations (The "Why")

## Table of Contents
* [1. The "JavaScript" Analogy](#1-the-javascript-analogy)
* [2. How eBPF Works (The Safety Net)](#2-how-ebpf-works-the-safety-net)
* [3. The Lifecycle](#3-the-lifecycle)

---

## 1. The "JavaScript" Analogy

This is the best way to understand eBPF:
**eBPF is to the Linux Kernel what JavaScript is to a Web Browser.**

### HTML (The Old Way)
*   In the 90s, if you wanted to change a webpage, you had to reload the whole page.
*   **Linux Equivalent:** If you wanted to change the Kernel, you had to write a "Kernel Module" and reboot (or risk crashing).

### JavaScript (The New Way)
*   Now, websites run little scripts (JS) *inside* the browser to verify forms or show animations. It's safe. If the script crashes, your browser doesn't explode.
*   **eBPF Equivalent:** You can upload little scripts *into* the running Kernel. They can see what's happening, but they are sandboxed so they can't crash the machine.

---

## 2. How eBPF Works (The Safety Net)

If we let anyone run code in the Kernel (The Kitchen), they might burn the place down.
So, eBPF has a bouncer called **The Verifier**.

### The Verifier
Before your program runs, the Kernel checks it:
1.  **"Will you finish?"** (No infinite loops allowed. You can't freeze the computer).
2.  **"Are you reading safe memory?"** (You can't read passwords from another app's memory).

If you fail these checks, the Kernel says: *"Reject! Safe code only."*

### JIT (Just-In-Time)
Once the Bouncer (Verifier) says you are cool, the **JIT Compiler** translates your script into native machine code (0s and 1s) so it runs *super fast*.

---

## 3. The Lifecycle

How do you actually use this?

1.  **Write Code:** You write a small C program.
2.  **Compile:** You turn it into "Bytecode" (an intermediate language).
3.  **Load:** You ask the Kernel to take it.
4.  **Attach:** You say *"Run this script every time `execve` happens"*.
5.  **Data:** The script writes data to a **Map** (like a shared mailbox).
6.  **Read:** Your User Space app reads the mailbox to see what happened.

---

### Summary
*   **eBPF** lets you change Kernel behavior safest and dynamically.
*   **The Verifier** ensures you don't crash the system.
*   **Maps** are how you get data out.
