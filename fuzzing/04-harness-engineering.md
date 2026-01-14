# Phase 4: Harness Engineering for Scale

Scale isn't just about adding more CPUs. It's about **making every CPU cycle count**. A bad harness is the #1 reason purely effective fuzzing campaigns fail.

## 1. What is a "Harness"?
A harness is a small C/C++ function that acts as a bridge between the Fuzzer (which speaks "Random Bytes") and the Target "Library (which speaks "API calls").

Standard Interface:
```c
// The Fuzzer calls this function millions of times
int LLVMFuzzerTestOneInput(const uint8_t *Data, size_t Size) {
  // Your job: Pass 'Data' to the target library
  Target_ProcessData(Data, Size);
  return 0;
}
```

---

## 2. Fork Server vs. Persistent Mode
How does the fuzzer reset the program state after every run?

### The "Fork Server" (Old Way)
1.  Fuzzer pauses at `main()`.
2.  Fuzzer `fork()`s the entire process.
3.  Child process runs the input.
4.  Child process exits.
5.  **Cost:** `fork()` is expensive! It takes milliseconds. This caps you at ~500 execs/sec.

### Persistent Mode (The Scale Way)
Instead of killing the process, we **loop inside the process**.
1.  Fuzzer starts process.
2.  Process runs a `while(__AFL_LOOP(10000))` loop.
3.  Inside the loop, run the target function.
4.  Reset the state manually (free memory, reset counters).
5.  **Cost:** Microseconds. Speed goes up to **5,000 - 10,000 execs/sec**.

---

## 3. Scale Consideration: Global State Pollution
Persistent mode is dangerous. If you don't clean up, **Run #5** might crash because of something **Run #4** changed.

*   **Example**: A static counter `static int user_count = 0;`.
*   **Run 1**: `user_count` becomes 1.
*   **Run 2**: `user_count` becomes 2.
*   **Run 1000**: Code crashes because `user_count` > 999.
*   **False Positive**: The input for Run 1000 isn't malicious; the dirty state caused the crash.

**Fix**: You must write a `Reset()` function that manually clears all globals/statics after every iteration.

---

## 4. Scale Consideration: Mocking
Real code does slow things: Network requests, Disk I/O, crypto, sleep().

*   **The Problem**: If your function calls `ConnectToServer()`, it might wait 100ms for a timeout.
    *   100ms = 10 execs/sec. **Unacceptable.**
*   **The Fix**: **Mocking**. Replace the slow function with a fake one.

```c
// Real Code
// void CheckLicense() { connect(google.com); ... }

// Fuzz Harness Override
void CheckLicense() {
  // Pretend we checked and it's fine. Return instantly.
  return; 
}
```

---

## 5. Hands-on: Optimizing a Slow Harness (Before vs After)

### The Slow Harness (Bad)
```c
// harness_slow.c
#include "my_lib.h"

int main(int argc, char** argv) {
  // 1. Startup is slow (Parsing config files)
  Config* cfg = LoadConfig("config.ini"); 
  
  // 2. Reading from disk is slow
  char buffer[1024];
  FILE* f = fopen(argv[1], "rb");
  fread(buffer, 1, 1024, f);
  
  // 3. Execution
  ProcessData(cfg, buffer);
  
  return 0;
}
```
**Why it's bad**: It re-reads the config and opens a file every single time. `fork()` overhead kills performance.

### The Fast Harness (Persistent Mode)
```c
// harness_fast.c
#include "my_lib.h"

// 1. Setup happens ONCE
Config* cfg;
void Initialize() {
  cfg = LoadConfig("config.ini");
}

int LLVMFuzzerTestOneInput(const uint8_t *Data, size_t Size) {
  static int initialized = 0;
  if (!initialized) {
    Initialize();
    initialized = 1;
  }

  // 2. No disk I/O. Use 'Data' directly from memory.
  // 3. Persistent Loop is automatic with LLVMFuzzerTestOneInput
  ProcessData(cfg, Data, Size);
  
  return 0;
}
```
**Why it's good**: Configuration loads once. No file I/O. No forking. **100x speedup.**
