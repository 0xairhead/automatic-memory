# Architecture & Assembly Lesson

## 1. The CPU Layout: Registers
Think of registers as the "hands" of the CPU. If RAM is the workbench, the registers are what the CPU is currently holding.

### **General Purpose Registers (x86 - 32-bit / x64 - 64-bit)**
You will see these constantly. In 64-bit mode, they start with 'R' (e.g., RAX). In 32-bit, they start with 'E' (e.g., EAX).
*   **RAX/EAX**: Accumulator. Used for arithmetic and **return values**.
*   **RBX/EBX**: Base. General storage.
*   **RCX/ECX**: Counter. Used for loops.
*   **RDX/EDX**: Data.
*   **RSI/ESI**: Source Index.
*   **RDI/EDI**: Destination Index.
*   **RSP/ESP**: Stack Pointer. Points to the top of the stack.
*   **RBP/EBP**: Base Pointer. Points to the base of the stack frame.

### **Instruction Pointer**
*   **EIP/RIP**: Points to the **next instruction** to execute. You cannot control this directly (usually), but `jmp`, `call`, and `ret` instructions control it.

## 1.1 CPU Flags (The EFLAGS Register)
When the CPU does math or compares values, it updates the **EFLAGS** register. These flags decide where code jumps (`je`, `jne`, `jg`).
*   **ZF (Zero Flag):** Set to 1 if the result of an operation is Zero (e.g., `cmp eax, eax` or `sub eax, eax`). Used by `je` (Jump if Equal).
*   **CF (Carry Flag):** Set to 1 if an unsigned arithmetic operation overflows (e.g., 0xFF + 1). Used by `jb` (Jump if Below).
*   **SF (Sign Flag):** Set to 1 if the result is negative. Used by `jl` (Jump if Less).

## 2. Memory: Stack vs. Heap
*   **Stack**:
    *   **LIFO** (Last In, First Out).
    *   Used for local variables, function arguments, and return addresses.
    *   Grows **downwards** (towards lower memory addresses).
    *   "Pushing" onto the stack decreases ESP. "Popping" increases ESP.
    *   **Stack Frame:** Every function gets a "frame" to store its data.
        *   **Prologue:** `push rbp; mov rbp, rsp` (Saves old base pointer, sets new one).
        *   **Local Vars:** `sub rsp, 0x10` (Makes space).
        *   **Epilogue:** `leave; ret` (Restores stack and returns).
*   **Heap**:
    *   Dynamic memory (`malloc`, `new`).
    *   Managed manually or by the OS.
    *   Used for persistent data that must outlive a function call.

## 2.1 Memory Segments
A program isn't just one big blob. It's divided into sections:
*   **`.text` (Code):** Where the instructions live. Read-Only/Executable.
*   **`.data`:** Initialized global variables (e.g., `int x = 10;`). Read/Write.
*   **`.bss`:** Uninitialized global variables (e.g., `int y;`). Read/Write. (Takes up no space on disk).
*   **`.rdata` / `.rodata`:** Read-only data (e.g., string literals "Hello World").

## 3. Calling Conventions
How do functions talk to each other? Who cleans up the stack?

### **System V AMD64 ABI (Linux x64 Standard)**
*   The first 6 integer arguments are passed in registers: **RDI, RSI, RDX, RCX, R8, R9**.
*   Remaining arguments are pushed onto the stack.
*   **Caller** cleans up the stack.
*   *Key indicator*: You see `mov edi, 0x...` before a `call`.

### **cdecl (C Declaration - 32-bit)**
*   Standard for 32-bit C functions.
*   Arguments pushed onto the stack **Right-to-Left**.
*   **Caller** cleans up the stack.

### **stdcall (Windows 32-bit)**
*   Standard for Windows API functions.
*   Arguments pushed **Right-to-Left**.
*   **Callee** cleans up the stack.

### **fastcall**
*   First few arguments passed in **registers** (usually ECX, EDX) for speed.
*   Remaining arguments on stack.
*   *Common in modern x64 and optimized code.*

---
## Practice Time
We will now compile a simple C program and look at what it actually looks like to the computer.
