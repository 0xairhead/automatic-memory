# Lesson 9: The Fuzzing Gauntlet (Capstone) 🥊

Welcome to the final exam.
Below are three challenges. Each one requires a different technique you've learned.

## Challenge 1: The Glass Jaw (Classic Overflow) 🥛
**Difficulty**: Easy
**Concept**: Coverage-guided Fuzzing (Lesson 2 / 3)

This program has a bug deep inside a few if-statements. A random fuzzer might miss it, but AFL (Coverage-guided) will find it quickly.

```c
// glass_jaw.c
#include <stdio.h>
#include <string.h>

int main() {
    char key[100];
    char buffer[10];
    
    printf("Enter key: ");
    if (fgets(key, sizeof(key), stdin)) {
        if (key[0] == 'F') {
            if (key[1] == 'U') {
                if (key[2] == 'Z') {
                    if (key[3] == 'Z') {
                        // The Bug: Copying a large input into a small buffer
                        strcpy(buffer, key); 
                        printf("You win!\n");
                    }
                }
            }
        }
    }
    return 0;
}
```
**Goal**: Find the crash.
**Hint**: Use **AFL++** or **Honggfuzz**. They will quickly learn the path `F-U-Z-Z` and then trigger the overflow.

---

## Challenge 2: The Gatekeeper (Structured Input) ⛩️
**Difficulty**: Medium
**Concept**: Grammar / Structure-Aware Fuzzing (Lesson 6 / 7)

This parser rejects anything that isn't valid "XML-like" data. AFL will struggle to get past the first check.

```python
# gatekeeper.py
import sys

def parse_xml(data):
    if not data.startswith("<root>") or not data.endswith("</root>"):
        return "Rejected: Invalid Format"
    
    if "<admin>true</admin>" in data:
        # The Bug: We shouldn't allow admin access!
        raise Exception("CRASH: Admin Access Gained!")
        
    return "Accepted"

if __name__ == "__main__":
    print(parse_xml(sys.stdin.read()))
```
**Goal**: Bypass the format check and trigger the exception.
**Hint**: Use a **Grammar Fuzzer** (Lesson 6) that knows the rules: `<start> -> <root>...`. If you rely on bit-flipping, you'll be blocked at the door forever.

---

## Challenge 3: The Silent Killer (Memory Corruption) 🥷
**Difficulty**: Hard
**Concept**: AddressSanitizer (ASAN) (Lesson 5)

This program has a serious bug, but it usually doesn't crash! It just writes data one byte too far.

```c
// silent_killer.c
#include <stdlib.h>
#include <string.h>

int main(int argc, char **argv) {
    char *data = malloc(10);
    
    // The Bug: Writing 11 bytes into a 10-byte buffer.
    // Without ASAN, this might just overwrite metadata and keep running.
    data[10] = 'X'; 
    
    free(data);
    return 0;
}
```
**Goal**: Detect the bug.
**Hint**: Compile this with `gcc -fsanitize=address -g silent_killer.c`. Run it, and watch **ASAN** scream at you.

---

## Congratulations! 🏆
You have conceptually completed the Fuzzing Course.
1.  You started as a blindfolded monkey.
2.  You became a safe-cracker (Coverage-guided).
3.  You learned to speak languages (Grammar).
4.  You became a Lego Master (Structure-aware).
5.  And you learned to catch the silent killers (ASAN).

**Go forth and break things.**
