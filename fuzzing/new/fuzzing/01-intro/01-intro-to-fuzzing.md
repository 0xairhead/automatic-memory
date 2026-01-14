# Lesson 1: The Blindfolded Monkey (Intro to Fuzzing)

## Imagine a Monkey 🐵
Imagine you put a blindfolded monkey in front of a keyboard. It starts hitting random keys: `agklhja$#%`.

Most of the time, it writes nonsense. But if it types for a million years, eventually it might hit the special self-destruct sequence: `crash`.

**Fuzzing** is just using a fast computer to be that monkey.

## The Experiment
We made two files:
1.  **The Target (`vuln.c`)**: A program that explodes if anyone types "crash".
2.  **The Fuzzer (`fuzzer.py`)**: A "blindfolded monkey" script that types random letters.

## Try it!
1.  **Build the target**: `gcc vuln.c -o vuln`
2.  **Unleash the monkey**: `python3 fuzzer.py`

You will see it trying thousands of random things until BAM! It accidentally types "crash" and the program breaks. That's a Fuzzer!
