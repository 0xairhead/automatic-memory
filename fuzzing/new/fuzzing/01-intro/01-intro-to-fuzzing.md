# Lesson 1: The Blindfolded Monkey (Intro to Fuzzing)

## Imagine a Monkey 🐵
Imagine you put a blindfolded monkey in front of a keyboard. It starts hitting random keys: `agklhja$#%`.

Most of the time, it writes nonsense. But if it types for a million years, eventually it might hit the special self-destruct sequence: `crash`.

**Fuzzing** is just using a fast computer to be that monkey.

## The Experiment
We made two files:
1.  **The Target (`vuln.c`)**: A program that explodes if anyone types "crash".
2.  **The Fuzzer (`fuzzer.py`)**: A "blindfolded monkey" script that types random letters.

## Deep Dive: How the Code Works 🧐

### The Target: `vuln.c`
This is a very simple program that crashes if you type exactly "crash".

```c
// vuln.c
if (strcmp(buffer, "crash\n") == 0) {
    // strcmp = String Compare.
    // It checks if 'buffer' is EXACTLY "crash\n"
    
     printf("Boom!\n");
     abort(); // This command crashes the program intentionally
}
```

### The Fuzzer: `fuzzer.py`
This script acts like the monkey. It knows nothing about the "crash" password.

```python
# fuzzer.py
def generate_input():
    # 1. Pick a random length (e.g., 5 letters)
    length = random.randint(1, 10)
    
    # 2. Pick random letters (e.g., "a", "g", "z")
    # This creates strings like "agz\n" or "hello\n"
    return ''.join(random.choices(string.ascii_lowercase + '\n', k=length))
```

It just calls `generate_input()` over and over again until the program crashes.

## Try it!
1.  **Build the target**: `gcc vuln.c -o vuln`
2.  **Unleash the monkey**: `python3 fuzzer.py`

You will see it trying thousands of random things until BAM! It accidentally types "crash" and the program breaks. That's a Fuzzer!
 
## Practice Questions 🧠

1.  **Concept Check**: In the "blindfolded monkey" analogy, what does the monkey represent in a fuzzing setup?
    <details>
    <summary>Answer</summary>
    The monkey represents the **Fuzzer** (specifically, the random input generator). It blindly generates inputs without understanding the program logic.
    </details>

2.  **Target Analysis**: Looking at `vuln.c`, why does the input "crash" cause the program to fail?
    <details>
    <summary>Answer</summary>
    The code has an explicit check `if (strcmp(buffer, "crash\n") == 0)`. When the input matches "crash" (plus a newline), the program executes the `abort()` function, causing it to crash.
    </details>

3.  **Core Components**: What are the two main files required to run this fuzzing experiment, and what is the role of each?
    <details>
    <summary>Answer</summary>

    1.  **The Target (`vuln.c`)**: The vulnerable program we are trying to break.
    2.  **The Fuzzer (`fuzzer.py`)**: The script that generates random inputs to throw at the target.
    
    </details>
