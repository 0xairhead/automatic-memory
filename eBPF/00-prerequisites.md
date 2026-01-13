# Phase 0: Prerequisites (Explained Simply)

## Table of Contents
* [1. The Restaurant Analogy (User vs Kernel)](#1-the-restaurant-analogy-user-vs-kernel)
* [2. Linux Basics (The "Waiters")](#2-linux-basics-the-waiters)
* [3. C Programming (Don't Panic)](#3-c-programming-dont-panic)

---

## 1. The Restaurant Analogy (User vs Kernel)

Before we talk technicals, let's look at how computers actually work using a **Restaurant** analogy.

### User Space = The Dining Room
*   This is where your applications (Chrome, Python, Slack) live.
*   **Safety:** You (the customer) sit at a table. You cannot just walk into the kitchen and grab a steak. You might burn yourself or steal food.
*   **Restricted:** If you want water, you have to *ask* for it.

### Kernel Space = The Kitchen
*   This is where the Operating System (Linux) lives.
*   **Privileged:** The staff has access to the knives, the fire (CPU), and the fridge (Memory).
*   **Critical:** If a waiter slips and falls, the whole service might stop. If the Chef quits, the restaurant closes (System Crash).

### The System Call (The "Order")
*   When your app needs to save a file, it's like asking a waiter: *"Can I have a steak?"*
*   Technically, this is a **System Call** (or **syscall**).
*   Your app pauses, the kernel does the work (talks to the hard drive), and returns the result.

---

## 2. Linux Basics (The "Waiters")

eBPF watches the waiters. To understand it, you need to know what the waiters do.

### A. PID (Process ID)
*   Every customer (program) has a ticket number. That's the **PID**.
*   The kernel uses this to track who is asking for what.

### B. Syscalls (The Menu)
There is a fixed menu of things you can ask the kernel to do.
*   `open()` -> "Open this file for me."
*   `write()` -> "Write this data to the network."
*   `execve()` -> "Run this new program." (This is the most important one for security!).

### C. Everything is a File
*   In Linux, even your mouse and your network connection look like "files".
*   If you can read/write to it, it's a file.

---

## 3. C Programming (Don't Panic)

eBPF programs are written in **C**.
*   *Wait! Don't run away!*
*   You do **not** need to be a C expert. You mostly just need to be able to *read* it.

### The 3 Things You Need to Know:

1.  **Structs:**
    *   Think of a `struct` like a Python Dictionary or a JSON object. It's just a group of variables bundled together.
    *   `struct Task { int pid; char name[16]; }`

2.  **Pointers (`*`):**
    *   A pointer is just an **Address**.
    *   It doesn't hold the value "42". It holds the note: *"The value 42 is stored in Locker #500"*.
    *   In the kernel, we use pointers everywhere to avoid copying massive amounts of data.

3.  **Functions:**
    *   They look like `int my_function(int a) { return a + 1; }`.
    *   Pretty standard stuff.

---

### Summary
*   **User Space:** Safe, restricted (Dining Room).
*   **Kernel Space:** Powerful, dangerous (Kitchen).
*   **Syscalls:** The way User Space asks Kernel Space for help.
