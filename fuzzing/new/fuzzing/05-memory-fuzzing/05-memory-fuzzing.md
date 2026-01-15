# Lesson 5: The Speed of Light (In-Process Fuzzing) ⚡

## 1. Starting the Car 🚗

### The Old Way: "Forking"
Our old fuzzers (like the ones running simple python scripts) were like this:
1.  **Start the car** (Launch the program process).
2.  **Drive 1 inch** (Feed it ONE input file).
3.  **Turn off the car** (Kill the process).
4.  **Repeat**.

This is **SLOW**. Starting a process takes time. Even if it's fast (0.01s), that limits us to ~100 execs/sec.

### The New Way: Persistent Mode 🏎️
**libFuzzer** (In-Process Fuzzing) acts like keeping the engine running.
1.  **Start the car ONCE**.
2.  **Drive, Drive, Drive!** (Call the test function 1,000,000 times inside the same process using a loop).
3.  **Never turn it off** until it crashes.

Because we don't restart the OS process every time, we can reach speeds of **100,000 execs/sec**.

## 2. Deep Dive: How the Code Works 🧐

To achieve this speed, we don't just "run a program". We write a special function called the **Fuzz Target**.
The fuzzer calls this function over and over again with different random data.

```cpp
// harness.cc
#include <stdint.h>
#include <stddef.h>

// This is the specific function name libFuzzer looks for:
extern "C" int LLVMFuzzerTestOneInput(const uint8_t *Data, size_t Size) {
    
    // 1. "Data" is the random bytes from the fuzzer (the "monkey")
    // 2. We pass it directly to our target function
    bool result = ParseJPEGImage(Data, Size);
    
    return 0;  // Return 0 to say "I'm done, send me the next one!"
}
```

The fuzzer handles the loop. We just provide the logic for **one** test case.

## 3. Magic Glasses (AddressSanitizer) 👓

We also use special glasses called **AddressSanitizer (ASAN)**.

### The Problem: Silent Bugs
In C/C++, you can sometimes write past the end of an array, and the program **won't crash**. It just corrupts memory silently. Use-After-Free bugs are also often silent.

```c
char buffer[10];
buffer[11] = 'A'; // Writing outside bounds!
// In normal C, this might NOT crash. It just overwrites whatever was next.
```

### The Solution: Redzones
ASAN surrounds every variable with "poisoned" memory (Redzones). If you touch the poison, ASAN pauses the program and yells **"ERROR!"** immediately.

**Demo**: When we compile with `-fsanitize=address`, the compiler adds checks around every memory access. It makes the fuzzer slightly slower (2x), but it catches **way more bugs**.

## Practice Questions 🧠

1.  **Concept Check**: Why is "In-Process Fuzzing" (libFuzzer) faster than "File-Based Fuzzing" (AFL classic)?
    <details>
    <summary>Answer</summary>

    Because it avoids the overhead of creating a new generic operating system process (fork/exec) for every single test case. It reuses the same process memory.
    </details>

2.  **Critical Thinking**: If `LLVMFuzzerTestOneInput` runs in a loop inside the same process, what happens if your target function has a memory leak (doesn't free memory)?
    <details>
    <summary>Answer</summary>

    The fuzzer typically eventually crashes with an "Out of Memory" (OOM) error! Since the process never restarts, the leaked memory piles up until the RAM is full. You must fix memory leaks to use in-process fuzzing effectively.
    </details>

3.  **Tooling**: What does AddressSanitizer (ASAN) help you find that a normal crash report might miss?
    <details>
    <summary>Answer</summary>

    It finds **Memory Corruption** bugs that don't immediately cause a segmentation fault, such as buffer overflows (writing slightly out of bounds) or Use-After-Free (accessing memory after it's been deleted).
    </details>
