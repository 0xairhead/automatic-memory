# Windows Internals Lesson

## 1. User Mode vs. Kernel Mode
Windows isolates critical system components from user applications.
*   **User Mode (Ring 3):** Where your applications (Chrome, Word, Notepad) run. They cannot touch hardware directly. If they crash, only the app dies.
*   **Kernel Mode (Ring 0):** Where the OS kernel and drivers run. They have full access to hardware and memory. If they crash, you get a Blue Screen of Death (BSOD).

## 2. Key Data Structures (The "Alphabet Soup")
To manage processes, Windows uses specific data structures.

### **EPROCESS (Executive Process Block)**
*   Lives in **Kernel Mode**.
*   The "God Object" for a process. Contains everything the OS needs to know (PID, creation time, token, list of threads).
*   Contains the `ActiveProcessLinks` (a circular linked list of all processes on the system). *This is a common target for rootkits (Direct Kernel Object Manipulation).*

### **PEB (Process Environment Block)**
*   Lives in **User Mode**.
*   Contains data the *application* needs to know about itself (e.g., loaded modules/DLLs, command line arguments, environment variables, "BeingDebugged" flag).
*   **Access:** Can be accessed via the `FS` register (x86) or `GS` register (x64).
    *   x86: `FS:[0x30]`
    *   x64: `GS:[0x60]`

### **TEB (Thread Environment Block)**
*   Lives in **User Mode**.
*   One for *every thread*. Contains thread-local storage, stack base/limit, and the Last Error code (`GetLastError()`).
*   **Access:**
    *   x86: `FS:[0x18]`
    *   x64: `GS:[0x30]`

## 3. Objects & Handles
Windows is "Object-Oriented" at the kernel level.
*   **Object:** A system resource (File, Process, Thread, Mutex, Registry Key).
*   **Handle:** An indirect "pointer" or "ticket" given to a user mode application to access an Object.
*   **Handle Table:** The kernel keeps a table mapping these integer Handles to the actual Kernel Objects. applications cannot just "guess" a pointer; they must request a Handle.

## 4. The Windows API (Win32 API)
Unlike Linux (which uses syscalls like `int 0x80` or `syscall` directly), Windows apps talk to **Subsystem DLLs** (kernel32.dll, user32.dll, ntdll.dll).
1.  **Application:** Calls `CreateFile` in `kernel32.dll`.
2.  **Subsystem:** `kernel32.dll` calls `NtCreateFile` in `ntdll.dll` (The Native API).
3.  **Transition:** `ntdll.dll` executes the `syscall` instruction to jump to Kernel Mode.
4.  **Kernel:** The kernel executes the real work.

## 5. Summary
| Concept | Location | Access | Notes |
| :--- | :--- | :--- | :--- |
| **EPROCESS** | Kernel | Ring 0 Only | "God Object", ActiveProcessLinks |
| **PEB** | User | Ring 3 (FS/GS) | Loaded Modules, BeingDebugged |
| **Handle** | User/Kernel | Ring 3 (Integer) | Ticket to access Kernel Objects |
| **Native API** | User | ntdll.dll | The bridge to the Kernel |
