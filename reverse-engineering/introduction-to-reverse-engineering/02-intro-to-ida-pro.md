# Lesson 02: Introduction to IDA Pro

## Learning Objectives

By the end of this lesson, you will be able to:

*   Explain the role of **IDA Pro** in the reverse engineering workflow.
*   Differentiate between the **Freeware**, **Demo**, and **Pro** versions of IDA.
*   Describe how IDA **loads and maps** a binary into memory for analysis.
*   Understand the **two-pass analysis** process (Instruction Parsing and Semantic Analysis).
*   Manage the **IDA Database (.idb)** and understand why we analyze the database, not the raw binary.
*   Correctly **save and pack** your work to prevent data loss.

---

## 1. Introduction to IDA Pro

**IDA Pro (Interactive Disassembler)** is the industry-standard tool for static analysis. It is a "multipass" disassembler that attempts to convert binary machine code into human-readable assembly language. Beyond simple translation, IDA analyzes the code to identify functions, variables, data structures, and cross-references, effectively reconstructing the program's logic.

---

## 2. Versions and Licensing

Navigating the different versions of IDA can be confusing. For this course, it is critical to use a version that supports **saving your work**.

### The Lab Version (Freeware 5.0)
The course virtual machine comes pre-installed with an older **Freeware version (IDA 5.0)**.
*   **Pros**: It is fully functional for 32-bit x86 analysis and, crucially, **allows you to save your database**.
*   **Cons**: It lacks the modern user interface and support for newer architectures (x64) found in recent versions.
*   **Usage**: You **must** use this version for your labs to submit your `.idb` files.

### The Demo Version
Hex-Rays provides a free Demo of the latest IDA Pro.
*   **Pros**: Modern interface, latest processor support.
*   **Cons**: **Saving is disabled.** You cannot save your analysis or create a database.
*   **Usage**: Good for quick practice, but useless for coursework submission.

### The Commercial (Pro) Version
*   **Cost**: Extremely expensive (thousands of dollars per seat).
*   **Usage**: Standard for professional malware analysts and vulnerability researchers.

---

## 3. Loading a Binary

To begin analysis, you must load a binary file (e.g., `.exe` or `.dll`) into IDA.

### The Loading Process
1.  **Open**: Drag and drop the file onto the IDA icon or use **File > Open**.
2.  **Detection**: IDA uses **Loaders** to automatically detect the file format. For Windows software, it should detect a **PE (Portable Executable)** format.
3.  **Processor Selection**: IDA attempts to guess the processor architecture (e.g., Intel 80x86).
4.  **Options**: In almost all cases, **accept the defaults**. IDA's auto-detection is excellent.

> **Note**: If IDA misidentifies the file (e.g., thinks it's a raw binary instead of a PE file), the resulting disassembly will likely be garbage.

---

## 4. The Analysis Process

Once you click "OK," IDA begins its "under the hood" work. It doesn't just read the file; it reconstructs the program's memory image.

### Step 1: Memory Mapping
IDA reads the executable from the disk and maps it into **virtual memory**, just like the operating system's loader would when running the program. This ensures that memory addresses in IDA match what they potentially would be during execution.

### Step 2: Instruction Parsing (Pass 1)
IDA performs a **Recursive Descent** disassembly. It starts at the program's entry point and follows the code flow, decoding bytes into assembly instructions. It identifies:
*   Instruction types (mnemonics)
*   Operand lengths
*   Code locations vs. data locations

### Step 3: Semantic Analysis (Pass 2)
After the initial disassembly, IDA performs a second pass to add context:
*   **Cross-References (Xrefs)**: Linking calls to functions and jumps to labels.
*   **Function Identification**: Grouping instructions into function blocks.
*   **Argument Detection**: Guessing function parameters and local variables.

---

## 5. The IDA Database (.idb)

This is the most important concept to grasp: **IDA does not edit the original binary file.**

When you load `malware.exe`, IDA creates a **database** to store its analysis. All your comments, labeled functions, and colored graphs are stored in this database, not in the `.exe`.

### Temporary Files
While you are working, IDA creates several temporary files in the directory:
*   `.id0`, `.id1`: B-tree database components.
*   `.nam`: Name information.
*   `.til`: Type library information.

> **Warning**: Never delete or rename these files while IDA is open. Doing so will corrupt your current session.

### The .idb File
When you save your project, IDA combines all those temporary files into a single **`.idb` file** (or `.i64` for 64-bit). This acts as a "save state" for your analysis.
*   You can send an `.idb` file to a colleague, and they will see exactly what you see, including your custom comments and renames.
*   **For this course, you will submit the `.idb` file.**

---

## 6. Saving Your Work

**IDA Pro does NOT autosave.** If IDA crashes or you close it without saving, your work is gone forever.

### How to Save
1.  Go to **File > Save** (or `Ctrl+S`) periodically.
2.  When closing IDA, you will see a dialog box with options:
    *   **Don't save**: Deletes temporary files. All work is lost.
    *   **Pack database (Store)**: The recommended default. Archives temporary files into the `.idb`.
    *   **Pack database (Deflate)**: Same as Store, but compresses the file to save disk space.

---

## Summary

IDA Pro is a powerful tool with a unique workflow. It separates the **raw binary** (read-only) from the **analysis database** (read/write). Understanding this distinction—and remembering to save that database—is the first step to becoming a competent reverse engineer.

---

## Knowledge Check

1.  **Why is the "Demo" version of IDA Pro unsuitable for this course?**
    <details>
    <summary>Answer</summary>
    It disables the "Save" feature, meaning you cannot create or submit the required .idb database files.
    </details>

2.  **True or False: When you add a comment in IDA Pro, it modifies the original .exe file.**
    <details>
    <summary>Answer</summary>
    **False**. IDA modifies its own *database* (.idb), leaving the original binary untouched.
    </details>

3.  **What is the default action you should choose when closing IDA to ensure your work is kept?**
    <details>
    <summary>Answer</summary>
    "Pack database (Store)". This compiles the temporary working files into a single .idb file.
    </details>