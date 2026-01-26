# Exercises: Windows Internals

## Exercise 1: Kernel Structures (Paper Exercise)
*   **Task:** Draw the Windows process structures on a whiteboard.
    *   `EPROCESS` points to -> `KPROCESS`
    *   `PEB` (Process Environment Block) location in user mode vs kernel mode.
    *   **Question:** If I wanted to hide a process from Task Manager, which linked list in the kernel (`ActiveProcessLinks`) would I need to modify?

## Exercise 2: The Anti-Debug (PEB Access)
1.  **Setup:** Write the following C program (`anti_debug.c`) that checks for debuggers using two methods.
    ```c
    #include <windows.h>
    #include <stdio.h>

    // Function to check PEB manually via Inline Assembly
    int check_peb_manual() {
        unsigned char being_debugged = 0;
    #ifdef __x86_64__
        // x64: PEB is at GS:[0x60], BeingDebugged is at offset 0x02
        __asm__ ("movq %%gs:0x60, %%rax; movb 2(%%rax), %0" : "=r" (being_debugged) : : "%rax");
    #else
        // x86: PEB is at FS:[0x30], BeingDebugged is at offset 0x02
        __asm__ ("movl %%fs:0x30, %%eax; movb 2(%%eax), %0" : "=r" (being_debugged) : : "%eax");
    #endif
        return (int)being_debugged;
    }

    int main() {
        printf("Checking environment...\n");
        if (IsDebuggerPresent()) printf("[API] Debugger Detected!\n");
        if (check_peb_manual()) printf("[PEB] Debugger Detected!\n");
        return 0;
    }
    ```
2.  **Compile:**
    *   **macOS (Cross-Compile):** `x86_64-w64-mingw32-gcc code/anti_debug.c -o challenges/anti_debug.exe`
3.  **Task:** Open `challenges/anti_debug.exe` in your disassembler.
    *   Find the call to `IsDebuggerPresent`. This is an imported function from `kernel32.dll`.
    *   Find the manual check. Look for the "magic" segment registers:
        *   **x64:** `GS:[0x60]`
        *   **x86:** `FS:[0x30]`
    *   **Challenge:** Patch the binary to make it print "Safe" even when a debugger is attached (or force the check function to always return 0).

<details>
<summary><strong>Reveal Anti-Debug Solution Walkthrough</strong></summary>

### 1. Understanding the PEB
The **Process Environment Block (PEB)** is a data structure in user mode memory that contains information about the running process. The operating system uses it, but so can the process itself.

One of the most famous fields is `BeingDebugged` (a single byte at offset `0x02`), which the kernel sets to `1` if a debugger is attached.

### 2. Analysis: Finding the Check
In the disassembly Graph View, look for the `check_peb_manual` function.

**x64 Assembly signature:**
```assembly
mov rax, gs:60h     ; Load PEB address into RAX
movzx eax, byte ptr [rax+2] ; Load the BeingDebugged byte
```

**x86 Assembly signature:**
```assembly
mov eax, fs:30h     ; Load PEB address into EAX
movzx eax, byte ptr [eax+2]
```

### 3. Patching Strategies
*   **Strategy A (The API):** Find the `CALL IsDebuggerPresent`. After the call, you will see a `TEST EAX, EAX` (checking the return value) and a conditional jump (`JZ` or `JNZ`). Change the jump to force the "Not Detected" path.
*   **Strategy B (The PEB):** In `check_peb_manual`, simply replace the instructions that read the flag with `XOR EAX, EAX` (set return value to 0) and `RET`. This neutralizes the check entirely.
</details>
