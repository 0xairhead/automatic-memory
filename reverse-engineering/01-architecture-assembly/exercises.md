# Exercises: Architecture & Assembly

## Exercise 1: The "Secret" String
1.  **Setup:** Write a C program that takes a password as a command-line argument.
    ```c
    #include <stdio.h>
    #include <string.h>
    int main(int argc, char *argv[]) {
        if (argc < 2) return 1;
        if (strcmp(argv[1], "Flamingo!23") == 0) {
            printf("Access Granted\n");
        } else {
            printf("Access Denied\n");
        }
        return 0;
    }
    ```
2.  **Compile:** `gcc secret.c -o secret.exe` (on Windows) or standard ELF on Linux.
3.  **Task:** Open it in IDA/Ghidra.
    *   Can you find the `strcmp` call?
    *   Where is "Flamingo!23" stored? (Data section? Stack?)
    *   **Challenge:** Patch the binary (using a hex editor) so that entering "WRONG" prints "Access Granted".

## Exercise 2: The Loop
1.  **Setup:** Write a C program with a `for` loop that performs an XOR operation on an array of numbers.
2.  **Task:** Open in a disassembler.
    *   Identify the loop counter (usually a register like `ECX` or a stack var).
    *   Identify the XOR instruction.
    *   **Challenge:** Reconstruct the C code primarily by looking *only* at the graph view, not the decompiler.
