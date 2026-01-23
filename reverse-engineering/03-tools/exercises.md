# Exercises: Tools Mastery

## Exercise 4: Structure Reconstruction
1.  **Setup:** Find a simple "Crackme" on [crackmes.one](https://crackmes.one/) (Filter for difficulty: 1-2).
2.  **Task:** In IDA/Ghidra, find the main structure used to hold user info (e.g., Name/Serial).
3.  **Action:** Create a custom struct in the disassembler view and apply it to the code, so `[rax+4]` becomes `user->serial_id`.

## Exercise 6: The "Anti-Debug" Check
1.  **Setup:** Write a C program that searches for the `IsDebuggerPresent()` API.
    ```c
    if (IsDebuggerPresent()) {
        printf("I see you!\n");
        exit(1);
    }
    ```
2.  **Task:** Run this in x64dbg. It will exit immediately.
3.  **Challenge:**
    *   **Method A:** Change the Zero Flag (ZF) in the register view after the check happens to bypass it.
    *   **Method B:** Patch the binary on disk (change `JZ` to `JNZ` or `NOP` out the check).
    *   **Method C (Pro):** Use a debugger plugin (ScyllaHide) to hide the debugger.
