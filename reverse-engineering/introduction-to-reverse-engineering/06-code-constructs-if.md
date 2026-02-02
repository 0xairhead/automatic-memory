# Lesson 06: Code Constructs: IF Statements

## Learning Objectives

By the end of this lesson, you will be able to:

*   Explain the mechanics of **Control Flow** alterations using Comparisons (`CMP`) and Jumps.
*   Distinguish between **Signed** and **Unsigned** conditional jumps and what they imply about data types.
*   Recognize **Inverted Logic** in assembly code (e.g., how `if (a > b)` usually compiles to a check for `<=`).
*   Identify complex constructs such as **Logical Operators** and **Nested If Statements** in assembly.
*   Apply practical analysis techniques like **Back-Tracing** and **Text View** in IDA Pro.

---

## 1. The Mechanics of Comparisons and Jumps

Control flow is altered based on conditions using two primary steps: a comparison followed by a jump.

### The Compare Instruction (`CMP`)
The `CMP` instruction compares two operands (values) by subtracting the second operand from the first.
*   The result of this subtraction is **not stored** as a value.
*   Instead, it is used to set specific bits in the **EFLAGS register** (such as the Zero Flag or Carry Flag).

### Conditional vs. Unconditional Jumps
*   **Conditional Jump**: An `if` statement relies on this. It decides whether to branch based on the status of the flags set by the previous comparison.
*   **Unconditional Jump (`JMP`)**: Forces execution to a specific location regardless of any condition. This is often used to "skip over" an `else` block after the "true" branch has finished executing.

---

## 2. Signed vs. Unsigned Data

Assembly language does not have explicit variable types (like `int` or `unsigned int`). Reverse engineers must infer them from the instructions used.

### Signed Instructions
Used for signed integers (can be positive or negative).
*   **Examples**: `JL` (Jump Less), `JG` (Jump Greater), `JLE` (Jump Less or Equal), `JGE` (Jump Greater or Equal).

### Unsigned Instructions
Used for unsigned integers (always positive).
*   **Examples**: `JB` (Jump Below), `JA` (Jump Above).

### Context Clues
Identifying whether a program uses a signed jump (e.g., `JG`) or an unsigned jump (e.g., `JA`) provides a strong clue as to whether the original variables were signed or unsigned.

---

## 3. Inverted Logic in Assembly

A critical lesson for reverse engineering is that logic in assembly often appears as the **inverse** of the high-level source code.

### Jumping to Else
If the source code checks `if (a > b)`, the assembly often checks the opposite condition `if (a <= b)`.

**Reasoning**:
1.  The compiler wants to execute the "True" block if `a > b`.
2.  In linear assembly, it is often more efficient to check if the condition is **FALSE** (`a <= b`) and **JUMP** over the "True" block to the `else` (or exit).
3.  If the jump is *not* taken (i.e., `a` was indeed greater than `b`), the program naturally **falls through** into the "True" code block.

---

## 4. Analyzing Complex Constructs

Here is how more complex logic translates into assembly:

### Simple If/Else
Involves:
1.  A Comparison (`CMP`).
2.  A Conditional Jump to the `else` block (inverting logic).
3.  The "True" block instructions.
4.  An Unconditional Jump (`JMP`) to skip the `else` block instructions.
5.  The `else` block instructions.

### Logical Operators (AND/OR) and Else-If
*   **Short-Circuit Evaluation**: In a logical `AND` condition (e.g., `if A && B`), the assembly checks the first condition. If it fails, it immediately jumps to the end/else without evaluating the second condition.
*   **Else-If Chains**: Multiple comparisons followed by jumps to different locations typically indicate `else-if` blocks.

### Nested If Statements
These appear as a hierarchy where failing an outer comparison jumps to an `else` block, while passing it allows the code to "fall through" to a second, inner comparison.

---

## 5. Practical Analysis Techniques

### Graph vs. Text View
*   **Graph View**: Excellent for visualizing the overall flow and shapes of blocks.
*   **Text View**: Essential for verifying the specific linear order of instructions and jumps, answering *why* a block is located where it is.

### Back-Tracing
To understand a comparison, you often need to trace backwards.
*   *Example*: If you see `CMP EAX, 5`, you must scroll up to find where `EAX` got its value to understand *what* is being compared to 5.

### Renaming
As you identify logic, rename variables (e.g., `var_4` -> `user_count`) and locations immediately. This progressively clarifies the code.

---

## Knowledge Check

1.  **Which instruction is primarily used to compare two values before a jump?**
    <details>
    <summary>Answer</summary>
    **CMP** (Compare). It subtracts the operands to set flags.
    </details>

2.  **If you see a `JG` (Jump Greater) instruction, is the data likely Signed or Unsigned?**
    <details>
    <summary>Answer</summary>
    **Signed**. Unsigned comparisons would use `JA` (Jump Above).
    </details>

3.  **True or False: Assembly logic usually checks the exact same condition as the source code.**
    <details>
    <summary>Answer</summary>
    **False**. Assembly often checks the **inverse** condition to jump over the "true" block.
    </details>

4.  **What does "Short-Circuit Evaluation" mean in the context of an `AND` condition?**
    <details>
    <summary>Answer</summary>
    It means if the **first** condition fails, the code jumps immediately to failure/else **without** checking the second condition.
    </details>