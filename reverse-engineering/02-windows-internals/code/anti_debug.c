#include <windows.h>
#include <stdio.h>

/*
    Exercise: The Anti-Debug
    Goal: Identify how the program checks if it's being debugged.
*/

// Function to check PEB manually via Inline Assembly (works with MinGW GCC)
int check_peb_manual() {
    unsigned char being_debugged = 0;

#ifdef __x86_64__
    // x64: PEB is at GS:[0x60]
    // BeingDebugged flag is at offset 0x02
    __asm__ (
        "movq %%gs:0x60, %%rax \n\t"
        "movb 2(%%rax), %0"
        : "=r" (being_debugged)
        :
        : "%rax"
    );
#else
    // x86: PEB is at FS:[0x30]
    // BeingDebugged flag is at offset 0x02
    __asm__ (
        "movl %%fs:0x30, %%eax \n\t"
        "movb 2(%%eax), %0"
        : "=r" (being_debugged)
        :
        : "%eax"
    );
#endif
    return (int)being_debugged;
}

int main() {
    printf("--- Secure Application ---\n");
    printf("Checking environment integrity...\n");

    // 1. Standard API Check
    if (IsDebuggerPresent()) {
        printf("[ALERT] IsDebuggerPresent() detected a debugger!\n");
    } else {
        printf("[OK] IsDebuggerPresent() returned false.\n");
    }

    // 2. Manual PEB Check
    if (check_peb_manual()) {
        printf("[ALERT] PEB.BeingDebugged flag detected manual debugger presence!\n");
    } else {
        printf("[OK] PEB checks passed.\n");
    }

    printf("Exiting.\n");
    return 0;
}
