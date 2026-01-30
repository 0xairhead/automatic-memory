# Lesson 05: Code Constructs

## Learning Objectives

By the end of this lesson, you will be able to:

*   Explain the **goals of binary analysis** when source code is unavailable.
*   Distinguish between **Global** and **Local** variables and identify them in IDA Pro.
*   Apply practical analysis techniques such as **renaming variables** and **identifying standard library functions**.
*   Understand **arithmetic operations** at the assembly level, specifically division and sign extension (`CDQ`).
*   Perform **Two's Complement** calculations to verify negative values.

---

## 1. Goals of Binary Analysis

The primary goal of this lesson is to teach you how to analyze unknown binaries to determine their functionality and purpose without access to the original source code.

*   **Micro vs. Macro**: While you must initially examine code instruction-by-instruction to understand compiler behavior, the ultimate goal is to identify "big picture" functionality (e.g., "This function modifies the registry" or "This function opens a network socket") rather than getting lost in every single line of assembly.

---

## 2. Variable Scope and Identification in IDA

Distinguishing between variable types is a fundamental skill in reverse engineering.

### Global Variables
*   **Definition**: Variables accessible throughout the entire program.
*   **Location**: Typically defined in the `.data` (initialized) or `.bss` (uninitialized) sections.
*   **Identification**: Referenced by a hardcoded memory address (e.g., `mov eax, dword_414000`).
*   **Visuals**:
    *   **Initialized**: You will see their data value immediately (e.g., `255`).
    *   **Uninitialized**: You will see a question mark (`?`) in the data section.

### Local Variables
*   **Definition**: Variables accessible only within a specific function or block.
*   **Location**: Typically stored on the **Stack**.
*   **Identification**: Referenced using offsets relative to the Base Pointer (**EBP**) or Stack Pointer (**ESP**) (e.g., `ebp - 4` or `esp + 8`).
*   **Note**: Compiler optimizations may sometimes place local variables directly into **registers** instead of the stack.

---

## 3. Practical Analysis Techniques

### Renaming
*   **Hotkey**: `N`
*   **Best Practice**: As soon as you identify the purpose of a variable (e.g., a loop counter or a username buffer), rename it. This propagates the meaning throughout the entire IDA database, making the overall flow much easier to understand.

### Identifying Standard Functions
*   **Context Clues**: You can often identify standard library calls (like `printf` or `scanf`) by looking at the arguments pushed onto the stack *before* the call.
*   **Example**: If you see an instruction pushing an offset to a string format like `"%d"`, the subsequent `call` is almost certainly `printf` or `scanf`.

---

## 4. Arithmetic and Register Manipulation

### Division (`div` / `idiv`)
Division involves specific register pairs depending on operand size. For 32-bit operations:
*   The instruction divides the 64-bit value in **EDX:EAX** by the operand.
*   **Quotient**: Stored in **EAX**.
*   **Remainder**: Stored in **EDX**.

### Sign Extension (`CDQ`)
The `CDQ` (Convert Double to Quad) instruction is critical for **signed division**.
*   **Function**: It sign-extends **EAX** into **EDX**. This means it fills EDX with the sign bit of EAX (0s for positive, 1s for negative).
*   **Why it matters**:
    *   **Positive Numbers**: EDX must be zeroed out manually before division.
    *   **Negative Numbers**: Simply moving a value into EAX is not enough. You must use `CDQ` to ensure EDX has the correct sign bits, otherwise, the division result will be incorrect.

---

## 5. Two's Complement Refresher

Two's Complement is the standard method for representing negative integers in binary.

### How to Convert (Negative Hex to Decimal)
1.  **Binary**: Convert the hex digits to binary.
2.  **Invert**: If the most significant bit is 1 (indicating a negative number), flip all bits (0 becomes 1, 1 becomes 0).
3.  **Add One**: Add 1 to the result. This gives you the magnitude of the negative number.

---

## Summary

In this lesson, we moved from basic instructions to recognizing higher-level code constructs. By learning to distinguish global from local variables and understanding the nuances of signed arithmetic instructions like `CDQ`, you are building the mental model necessary to reconstruct the original programmer's intent from raw assembly.

---

## Knowledge Check

1.  **Where are Global Variables typically stored?**
    <details>
    <summary>Answer</summary>
    In the `.data` (initialized) or `.bss` (uninitialized) sections.
    </details>

2.  **What instruction is often used before valid signed division to set up the EDX register?**
    <details>
    <summary>Answer</summary>
    `CDQ` (Convert Double to Quad).
    </details>

3.  **If `div` is executed, where is the Remainder stored?**
    <details>
    <summary>Answer</summary>
    In the **EDX** register.
    </details>

4.  **How are Local Variables typically referenced in assembly?**
    <details>
    <summary>Answer</summary>
    By offsets relative to the Base Pointer (`EBP`) or Stack Pointer (`ESP`), e.g., `[EBP-4]`.
    </details>
