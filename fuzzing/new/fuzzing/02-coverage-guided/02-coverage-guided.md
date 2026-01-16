# Lesson 2: The Safe Cracker (Coverage-guided Fuzzing) 🔓

## The Problem: The 5-Dial Safe
Imagine a safe with 5 dials (A-Z). The combination is **M-A-Z-E-!**.
There are 11 million possible combinations ($26^5$).

### 1. The Monkey Approach (Dumb Fuzzing) 🙉
A blindfolded monkey spins all 5 dials at random and pulls the handle.
- Try: `Q-W-E-R-T` -> Locked.
- Try: `A-S-D-F-G` -> Locked.
- Try: `M-A-Z-E-?` -> Locked.

The monkey doesn't know how close it was. It just knows "It didn't open".
It will take millions of tries to open the safe.

### 2. The Safe Cracker Approach (Coverage-guided) 🕵️
A professional safe cracker puts their ear against the safe.
- Try: `Q-W-E-R-T`. *Silence*. (The first dial is wrong).
- Try: `M-W-E-R-T`. *CLICK!* (The fuzzer hears a click!).

**This "CLICK" is Code Coverage.**
The program executed a **new line of code** inside the `if (input[0] == 'M')` block.

Because the Safe Cracker heard the click, they know: **"The first letter IS definitely M!"**
Now they stop spinning the first dial. They leave it at 'M' and work on the second one.

- Try: `M-A-X-X-X`. *CLICK!* (Second dial is correct!).
- Try: `M-A-Z-X-X`. *CLICK!* (Third dial is correct!).

Instead of 11 million tries, they solve it in about $26 \times 5 = 130$ tries.

## Deep Dive: How the Code Works 🧐

### The Target: `maze.c`
This program is a "Maze". You can't just teleport to the end; you have to pass each door one by one.

```c
// maze.c
if (buffer[0] == 'M') {             // Door 1
    printf("Level 1 passed\n");     // "CLICK!" Sound (Feedback)
    
    if (buffer[1] == 'a') {         // Door 2
        printf("Level 2 passed\n"); // "CLICK!" Sound
        
        if (buffer[2] == 'z') {     // Door 3
             // ... and so on ...
        }
    }
}
```

If you send "M", it prints "Level 1 passed".
If you send "Q", it prints nothing.

### The Smart Fuzzer: `smart_fuzzer.py`
This script listens for those "Level 1 passed" messages.

**1. The "Ear" (Feedback Loop)**
It reads the output of the program to see how far it got.

```python
# smart_fuzzer.py - inside the loop
stdout, stderr = process.communicate(input=data)

# Check the output "sounds"
level = 0
if "Level 1 passed" in stdout: level = 1
if "Level 2 passed" in stdout: level = 2

# The Critical "Aha!" Moment
if level > max_level_reached:
    # We found a new door! 
    # SAVE this input ("M") so we can build on it later.
    current_best_input = data 
```

**2. The "Hands" (Mutation)**
Instead of starting from zero every time, it takes the `current_best_input` and changes it slightly.

```python
def mutate(data):
    # If our best input is "M", let's try answering "M" + a random letter
    # Try: "Ma", "Mb", "Mc"...
    if len(data) < 5:
        data += random.choice(string.ascii_letters + '!')
    return data
```

**Summary of the Loop:**
1. Start with empty string `""`.
2. Try `"X"` (Fail). Try `"M"` (Success! "Level 1 passed").
3. Save `"M"` as the new best input.
4. Mutate `"M"` -> try `"Ma"` (Success! "Level 2 passed").
5. Save `"Ma"` as the new best input.
6. ... and so on until `"Maze!"`.

## Run the Demo
Watch the `smart_fuzzer.py` crack the safe step-by-step:

```bash
cd 02-coverage-guided
gcc maze.c -o maze
python3 smart_fuzzer.py
```

## Practice Questions 🧠

1.  **Analogy Check**: In the "Safe Cracker" analogy, what corresponds to the "click" sound in a software program?
    <details>
    <summary>Answer</summary>

    The "click" corresponds to **Code Coverage** (or executing a new path/line of code). It tells the fuzzer that the input triggered something new, even if it didn't cause a crash yet.

    </details>

2.  **Efficiency**: Why is the "Safe Cracker" approach (coverage-guided) faster than the "Monkey" approach (random fuzzing) for solving the maze?
    <details>
    <summary>Answer</summary>

    The "Safe Cracker" learns from partial success. When it finds a correct character (like 'M'), it saves it and builds upon it ("mutates" it), whereas the "Monkey" starts from scratch every time.

    </details>

3.  **Mechanism**: When `smart_fuzzer.py` detects that a new level has been passed, what does it do with that input?
    <details>
    <summary>Answer</summary>

    It **saves** the input as the `current_best_input`. Future inputs are created by "mutating" (slightly changing) this best input, rather than successful random guessing.
    
    </details>
