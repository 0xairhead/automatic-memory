# Syllabus: Reverse Engineering Mastery

This syllabus is designed to provide a comprehensive path to mastering Reverse Engineering. It covers all the skills and technologies required for advanced analysis and research, structured into a logical learning path.

## **Phase 1: Foundations of Architecture & Assembly**
**Goal:** Master the low-level language of binaries.
*   **x86 & x64 Architecture:** 
    *   Registers (General purpose, segment, eflags).
    *   Memory Management (Virtual memory, Paging, Stack vs. Heap).
    *   Calling Conventions (cdecl, stdcall, fastcall, x64 calling conventions).
*   **Assembly Language:**
    *   Instructions: Data movement (`mov`, `push`, `pop`), Arithmetic (`add`, `sub`, `xor`), Control Flow (`jmp`, `call`, `ret`, `cmp`, `jcc`).
    *   Manual Decompilation: Translating assembly blocks back to C logic.

## **Phase 2: Master the Toolset**
**Goal:** Gain proficiency with industry-standard analysis tools.
*   **Disassemblers & Decompilers:**
    *   **IDA Pro:** Navigation, Cross-references, Type system, IDAPython scripting.
    *   **Ghidra:** Project management, Decompiler analysis, Java scripting API.
*   **Debuggers:**
    *   **x64dbg:** User-mode debugging, breakpoints (hardware/software), patching.
    *   **WinDbg:** Kernel-mode vs User-mode, specialized extensions (bang commands).
    *   **GDB:** Linux debugging, PEDA/GEF extensions.
*   **System Monitors:**
    *   **Sysinternals Suite:** Procmon (filtering events), Process Explorer (handles, DLLs), TCPView.
*   **Binary Utilities:**
    *   **CFF Explorer / PEBear:** Inspecting PE headers.
    *   **binwalk:** Analyzing firmware/blobs.

## **Phase 3: Analysis Techniques**
**Goal:** Analyze real-world malware and complex binaries.
*   **Static Analysis:** 
    *   Identifying strings, imports, exports, hashes.
    *   Recognizing standard cryptographic constants and compression algorithms.
*   **Dynamic Analysis:**
    *   Detonation in sandboxes (Cuckoo/CAPE).
    *   Behavioral monitoring (registry changes, file system drops).
    *   Network traffic correlation.
*   **Advanced Topics:**
    *   **Unpacking:** Manual unpacking of packed executables (UPX, custom packers).
    *   **De-obfuscation:** Handling control flow flattening, dead code, opaque predicates.
    *   **Anti-Analysis:** Bypassing anti-debugger and anti-VM checks.

## **Phase 4: File Formats & Internals**
**Goal:** Deep dive into how executables are structured.
*   **PE-COFF (Windows):** DOS Header, PE Header, Sections (.text, .data, .rsrc), IAT/EAT.
*   **ELF (Linux):** Headers, Segments vs Sections, Dynamic linking, GOT/PLT.

## **Phase 5: Networking & Traffic Analysis**
**Goal:** Understand the "wire" side of reverse engineering.
*   **Protocols:** 
    *   **TCP/UDP:** Handshakes, flags, stream reassembly.
    *   **HTTP:** Methods, headers, status codes, common malware variations.
*   **Tools:**
    *   **Wireshark:** Filtering (display/capture filters), following streams, extracting objects.
    *   **Network Forensics:** Identifying C2 channels, beaconing patterns, custom protocols.

## **Phase 6: Practical Application**
**Goal:** Build a portfolio and apply skills to real-world scenarios.
*   **Malware Analysis Reports:** Write detailed technical reports on analyzed samples (IoCs, TTPs, functionality).
*   **Tool Development:** Write scripts (Python) to extract malware configurations or automate tasks.
*   **Threat Modeling:** Perform security design reviews of theoretical systems.

## **Recommended Resources**
*   **Books:** *Practical Malware Analysis*, *The IDA Pro Book*, *Secrets of Reverse Engineering*.
*   **Practice:** Crackmes.one, Flare-On Challenge archives, MalwareTrafficAnalysis.net.
