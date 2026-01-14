# Phase 5: Sanitizers & Crash Triage (at Scale)

Finding a "crash" (segmentation fault) is easy. Finding a subtle but exploitable memory corruption bug is hard. Handling **10,000 crashes** without drowning is even harder.

## 1. The Superpower: Sanitizers
Normal C/C++ programs often "work" even when they are broken.
*   **Buffer Overflow**: Writing 1 byte past an array might not crash immediately. It might just corrupt a variable used later.
*   **Use-After-Free**: Reading freed memory might returning "stale" data instead of crashing.

### Enter AddressSanitizer (ASAN)
ASAN recompiles your program to surround every object with **Redzones (Poisoned Memory)**.
*   If you touch a Redzone: **INSTANT CRASH**.
*   **Cost**: ~2x-3x CPU slowdown, 3x Memory usage.

**Other Sanitizers:**
*   **MSAN (MemorySanitizer)**: Detects reading uninitialized memory.
*   **UBSAN (UndefinedBehaviorSanitizer)**: Detects integer overflows, null pointer dereferences.

---

## 2. Scale Consideration: Crash Storms
Scale Scenario: You deploy a new fuzzer to 500 nodes.
*   5 minutes later: **Alert! 50,000 crashes found!**
*   Reality: It is probably **1 bug** triggered 50,000 ways.

### Deduplication
You cannot manually review 50,000 logs. You need **Automated Deduplication**.
*   **Method**: Crash Stacking (Stack Trace Hashing).
*   **Concept**: If two crashes happen in the same function call chain, they are likely the same bug.

#### Algorithm
1.  Run crash input.
2.  Get Stack Trace.
3.  Ignore volatile addresses (offsets).
4.  Hash the top 3-5 function names.
5.  `Hash("Process->Parse->DecodeImage")` = `0xA1B2C3`.
6.  Store crash in folder `0xA1B2C3`.
7.  If folder exists, discard (or increment counter).

---

## 3. Scale Consideration: Minimization
Fuzzers generate weird, bloated inputs.
*   **Crash Input**: A 5MB PDF file with 10,000 pages, but the bug is just a single byte in the header.
*   **Developer Reaction**: "I'm not debugging a 5MB file."

### The Minimizer (afl-tmin)
An automated tool that attempts to delete bytes from the file.
1.  Delete byte 100. Does it still crash? Yes? -> **Keep deleted.**
2.  Delete byte 101. No crash? -> **Restore byte.**
3.  Repeat until minimal.

**Result**: 5MB file -> 48 byte file. Developers are happy.

---

## 4. Hands-on: Manual Deduplication
To understand how Google's **ClusterFuzz** works, let's look at two stack traces.

**Crash A:**
```
#0 0x00007ffff in __asan_report_error
#1 0x000040112 in parse_header (parser.c:45)
#2 0x000040220 in process_pdf (main.c:112)
#3 0x000040500 in main
```

**Crash B:**
```
#0 0x00007ffff in __asan_report_error
#1 0x000040112 in parse_header (parser.c:45)
#2 0x000040220 in process_pdf (main.c:112)
#3 0x000040500 in main
```

**Crash C:**
```
#0 0x00007ffff in __asan_report_error
#1 0x000040888 in render_image (graphics.c:20) <--- DIFFERENT!
#2 0x000040220 in process_pdf (main.c:115)
#3 0x000040500 in main
```

**Analysis**:
*   **Crash A & B**: Same top frame (`parse_header`). **Signature**: `parse_header`. Count = 2.
*   **Crash C**: Different top frame (`render_image`). **Signature**: `render_image`. Count = 1.
*   **Report**: "2 Unique Bugs Found."
