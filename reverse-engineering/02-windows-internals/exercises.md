# Exercises: Windows Internals

## Exercise 7: Kernel Structures (Paper Exercise)
*   **Task:** Draw the Windows process structures on a whiteboard.
    *   `EPROCESS` points to -> `KPROCESS`
    *   `PEB` (Process Environment Block) location in user mode vs kernel mode.
    *   **Question:** If I wanted to hide a process from Task Manager, which linked list in the kernel (`ActiveProcessLinks`) would I need to modify?
