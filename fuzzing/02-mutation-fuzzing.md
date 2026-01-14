# Phase 2: Mutation-Based Fuzzing & Scale

In Phase 1, we learned that fuzzing at scale requires thousands of cores. In Phase 2, we learn **what** those cores primarily do: **Mutation**.

## 1. What is Mutation?
Mutation is the art of taking a valid input (like an image or a PDF) and slightly corrupting it to see if the parser breaks.

### Common Strategies
1.  **Bit Flipping**: Turning a `0` to a `1`. Good for testing flag checks.
2.  **Arithmetic**: Adding/subtracting small numbers (e.g., changing `0xFF` to `0x00`). Good for integer overflows.
3.  **Block Operations**: Deleting, cloning, or shuffling blocks of data. Good for buffer overflows.
4.  **Dictionary Injection**: Inserting specific dangerous strings (e.g., `NaN`, `%s`, `\0`) into the file.

---

## 2. Scale Consideration: Seed Quality vs. Quantity
When you have 10,000 cores, you might think "I'll just throw random trash at the target." **This fails.**

### Garbage In, Garbage Out
*   **The Problem**: If your seed is a 10MB video file, and your fuzzer runs at 100 execs/sec, it will take *years* to mutate the header bytes enough to pass the "Is this a valid video?" check.
*   **The Fix**: Small, high-quality seeds.
    *   **Minimize**: Use tools like `afl-tmin` to shrink seeds to the smallest possible size that still triggers the target code functionality.
    *   **Diversity**: 10 small, different seeds are better than 1 giant seed.

### Throughput at Scale
*   **Small Seed (1KB)**: 5,000 execs/sec.
*   **Large Seed (1MB)**: 50 execs/sec.
*   **Result**: The small seed fuzzing cluster is **100x more efficient**.

---

## 3. Scale Consideration: Determinism
In a single-player mode, random is fine. In a distributed system, **determinism is king**.

### The "Reproduction" Nightmare
Imagine Worker Node #8492 finds a crash. It sends the crash input to the central database.
You download it, run it on your laptop... and **nothing happens**.

### Why?
*   **Randomness**: Maybe the fuzzer used a random number generator for a decision *inside* the harness.
*   **Time**: Maybe the code depends on `time.now()`.

### The Solution: PRNG with Fixed Seeds
All scalable fuzzers (AFL, libFuzzer) use **Pseudo-Random Number Generators (PRNG)**.
*   If you tell the fuzzer "Run seed `12345`", it will generating the *exact same* sequence of mutations every time.
*   This guarantees that if Node #8492 crashed, **you will crash too**.

---

## 4. Hands-on: Simple Mutator Script
This Python script demonstrates the core logic of a mutation engine.

```python
import random
import sys

def load_file(filename):
    with open(filename, "rb") as f:
        return bytearray(f.read())

def mutate_bitflip(data):
    """Flips a single random bit."""
    if not data: return data
    idx = random.randint(0, len(data) - 1)
    bit = random.randint(0, 7)
    data[idx] ^= (1 << bit)
    return data

def mutate_arithmetic(data):
    """Adds a small integer to a random byte."""
    if not data: return data
    idx = random.randint(0, len(data) - 1)
    val = random.choice([-1, 1, 10, -10])
    # Keep it within byte range 0-255
    data[idx] = (data[idx] + val) % 256
    return data

def fuzz(seed_file, iterations=10):
    original_data = load_file(seed_file)
    print(f"[+] Loaded seed: {len(original_data)} bytes")
    print(f"[+] Generating {iterations} mutations...\n")

    for i in range(iterations):
        # Always copy so we don't destroy the original seed for the next run
        mutated_data = bytearray(original_data) 
        
        # Pick a random strategy
        strategy = random.choice([mutate_bitflip, mutate_arithmetic])
        mutated_data = strategy(mutated_data)
        
        # In a real fuzzer, we would 'execute' the target here
        # For now, let's just print the hex to see the change
        print(f"Run #{i+1} ({strategy.__name__}): {mutated_data.hex()[:40]}...")

if __name__ == "__main__":
    # Create a dummy seed if none exists
    with open("seed.bin", "wb") as f:
        f.write(b"HELLO_FUZZING_WORLD")
    
    fuzz("seed.bin")
```

### Try it yourself
1.  Save this code as `simple_fuzzer.py`.
2.  Run `python3 simple_fuzzer.py`.
3.  Observe how the output bytes change slightly in each run.
