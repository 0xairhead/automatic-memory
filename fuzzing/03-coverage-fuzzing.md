# Phase 3: Coverage-Guided Fuzzing

Blindly mutating files is efficient (high speed), but effective coverage-guided fuzzing (CGF) is "smart." It wants to maximize **Code Coverage**.

## 1. How Coverage Works (The Eyes of the Fuzzer)
To know if a mutation "worked," the fuzzer needs to know if it executed new code.

### Instrumentation
We don't use standard compilers (`gcc` or `clang`). We use **Instrumenting Compilers** (like `afl-clang-fast`).
When you compile `target.c`, the compiler inserts a tiny snippet of code at the start of every **Basic Block** (every `if`, `for`, `while`, or function entry).

```c
// Pseudo-code of what instrumentation looks like
void basic_block_A() {
  shm_trace_map[current_location ^ prev_location]++; 
  prev_location = current_location;
  // ... actual code ...
}
```

### The Bitmap
*   The fuzzer monitors a "Shared Memory Bitmap" (usually 64KB).
*   If a mutation causes a jump to a previously unvisited code block, a byte in the bitmap changes.
*   The fuzzer sees this change and thinks: **"This input is INTERESTING! Save it!"**

---

## 2. Scale Consideration: Edge Explosion
At scale, "more coverage" isn't always good.

### The Problem
*   If you fuzz a loop that runs 1 to 1,000,000 times, a naive fuzzer might think running it 5 times is "different" than running it 6 times.
*   **Path Explosion**: You end up with 1,000,000 test cases that all test the same loop, just with different iteration counts.
*   **Result**: The fuzzer wastes time refining a "boring" path instead of finding new bugs.

### The Solution: Collapsing
*   Modern tools (AFL++) group hit counts into buckets:
    *   1, 2, 3, 4-7, 8-15, 16-31, 32-127, 128+
*   This treats "loop ran 50 times" and "loop ran 60 times" as **the same** coverage, preventing explosion.

---

## 3. Scale Consideration: Saturation
How do you know when to STOP fuzzing?

*   **Coverage Saturation**: The point where investing 1,000 more CPU hours yields 0 new paths.
*   **Visualization**: Plot "Paths Found" over "Time".
    *   Start: Steep vertical climb.
    *   End: Flat horizontal line.
*   **At Scale**: When 100 fuzzers all hit saturation, **restart the cluster** with:
    *   New dictionaries.
    *   New seeds.
    *   Different compilers (SanitizerCoverage vs Edge).

---

## 4. Scale Consideration: Synchronization
If you run 20 copies of AFL++, they act like a hive mind.

*   **Main Node**: Coordinates.
*   **Secondary Nodes**: Workers.
*   **Sync**:
    1.  Worker A finds `input_123`.
    2.  Worker A writes it to its own queue directory.
    3.  Worker B periodicially scans Worker A's directory.
    4.  Worker B runs `input_123`. If it's interesting to B, B imports it.

**ClusterFuzz Strategy**:
It uses a central database (Cloud Storage / S3). Local workers upload new findings to the cloud; other workers download the latest "corpus bundle" periodically.

---

## 5. Tools: AFL++
**AFL++ (AFLplusplus)** is the gold standard for C/C++ fuzzing. It is a community-driven fork of Google's original AFL.

*   **Features:**
    *   **QEMU mode**: Fuzz binaries without source code.
    *   **Custom Mutators**: Plug in your own Python mutators (like Phase 2!).
    *   **Collision-free instrumentation**: Newer "PC-Guard" coverage.

---

## 6. Hands-on: Anatomy of an AFL Compile
This is how you would compile a target for actual coverage-guided fuzzing.

```bash
# 1. Install AFL++
sudo apt-get install aflplusplus

# 2. The Target (vulnerable.c)
# -----------------------------
# #include <stdio.h>
# int main() {
#    char buf[100];
#    read(0, buf, 100);
#    if (buf[0] == 'B') {
#       if (buf[1] == 'U') {
#          if (buf[2] == 'G') {
#             // Crash!
#             char *p = NULL; *p = 0;
#          }
#       }
#    }
# }
# -----------------------------

# 3. Compile with Instrumentation
afl-clang-fast vulnerable.c -o vulnerable_fuzz

# 4. Run the Fuzzer
# -i: Input folder (seeds)
# -o: Output folder (findings)
mkdir inputs
echo "AAA" > inputs/seed.txt
afl-fuzz -i inputs -o outputs ./vulnerable_fuzz
```

**Why this works**:
The fuzzer guesses 'B'. It sees the code took a new branch. It saves 'B'.
It guesses 'BU'. It sees another new branch. It saves 'BU'.
It guesses 'BUG'. CRASH!
