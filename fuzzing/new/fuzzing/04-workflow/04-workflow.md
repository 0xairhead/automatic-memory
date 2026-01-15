# Lesson 4: Cleaning Your Room (Workflow) 🧹

Fuzzing creates a **LOT** of mess. If you leave a fuzzer running overnight, you might wake up to:
- 1,000,000 test cases focused on the same path.
- 5,000 "crash reports" that are actually just the same bug hit 5,000 times.

We need to clean up. We need a workflow.

## 1. Minimization (Packing the Suitcase) 🧳

### The Analogy: The Over-packer
Imagine you are going on a weekend trip.
- You have 50 red t-shirts.
- You have 1 blue shirt.
- You have 1 green shirt.

Do you pack all 52 shirts? **No!** You pack **1 red**, **1 blue**, and **1 green**.
The other 49 red shirts are "redundant". They don't give you any new fashion options.

**Corpus Minimization** does exactly this. It looks at all the files your fuzzer generated and deletes the ones that don't trigger anything *new* (new code paths/coverage).

### Deep Dive: How the Code Works 🧐

#### The Smart Packer: `minimizer.py`
In a real fuzzer, we use code coverage to decide if a file is "unique". For this demo, let's look at a simpler version: **Deduplication**. This script removes files that are exactly identical repeats.

```python
# minimizer.py
import hashlib

def get_file_hash(content):
    # Create a unique fingerprint (MD5) for the content
    return hashlib.md5(content.encode()).hexdigest()

def minimize_corpus(files_list):
    unique_fingerprints = set()
    clean_suitcase = []

    print(f"📦 Inspecting {len(files_list)} items...")

    for file_content in files_list:
        fingerprint = get_file_hash(file_content)

        if fingerprint in unique_fingerprints:
            print(f"  🗑️  Throwing away duplicate (Hash: {fingerprint[:6]}...)")
        else:
            print(f"  ✅ Keeping unique item (Hash: {fingerprint[:6]}...)")
            unique_fingerprints.add(fingerprint)
            clean_suitcase.append(file_content)

    return clean_suitcase

# Let's try it!
messy_room = ["Red Shirt", "Blue Shirt", "Red Shirt", "Green Shirt", "Red Shirt"]
clean_pile = minimize_corpus(messy_room)

print(f"\n✨ Final Suitcase: {clean_pile}")
```

In the real world (`afl-cmin`), the tool runs every file through the target binary. If file A and file B trigger the **exact same edges** in the code, it deletes the larger one.

## 2. Triage (Grouping Accidents) 🚑

### The Analogy: The Banana Peel
Imagine 100 people slip on the same banana peel in the hallway.
- Patient 1 comes in with a bruised knee.
- Patient 2 comes in with a bruised knee.
- ...
- Patient 100 comes in with a bruised knee.

Do you have 100 different safety problems to fix? **No!** You have **1 problem** (the banana).
If you tried to "fix" all 100, you'd just be sweeping up the same banana peel 100 times.

**Crash Triage** groups similar crashes together so you know how many *unique* bugs you actually found.

### Deep Dive: How the Code Works 🧐

#### The Doctor: `triage.py`
A triage tool looks at the "stack trace" (where the program crashed) or the error signal. If two crashes happen at the exact same memory address, they are usually the same bug.

```python
# triage.py
def triage_crashes(crash_reports):
    buckets = {}

    print(f"🚑 Triaging {len(crash_reports)} accidents...")

    for report in crash_reports:
        # We assume the "cause" is the last line of the error log
        # e.g., "Error: Segmentation fault at address 0x1234"
        cause = report.split(":")[-1].strip()

        if cause not in buckets:
            buckets[cause] = 0
            print(f"  🆕 Found NEW Bug cause: '{cause}'")
        
        buckets[cause] += 1

    print("\n📋 Final Report:")
    for cause, count in buckets.items():
        print(f"  - Bug '{cause}' happened {count} times.")

# Let's try it!
hospital_log = [
    "Crash 1: SegFault at 0x001", # The Banana
    "Crash 2: SegFault at 0x001", # The Banana
    "Crash 3: BufferOverflow at 0x999", # A different issue (maybe a wet floor?)
    "Crash 4: SegFault at 0x001"  # The Banana again
]

triage_crashes(hospital_log)
```

## Practice Questions 🧠

1.  **Concept Check**: Why do we delete text files during "Minimization" even if they are valid inputs?
    <details>
    <summary>Answer</summary>

    Because valid doesn't mean **useful**. If we have 10,000 inputs that all just print "Hello", we only need to keep one. Tests that re-test the exact same code paths slow down the fuzzer without finding new bugs.
    </details>

2.  **Code Analysis**: In our `minimizer.py` example, we used MD5 hashes to check for uniqueness. Why is this method *worse* than `afl-cmin`'s coverage-based minimization?
    <details>
    <summary>Answer</summary>

    Hash-based minimization only checks if the **files** are identical. Coverage-based minimization checks if the **behavior** is identical.
    
    Example: 
    - File A: "Hello World"
    - File B: "Hello World!"
    
    Our script would keep both (different hashes). But `afl-cmin` might delete one if they both trigger the exact same print function in the target program.
    </details>

3.  **Critical Thinking**: You ran a fuzzer for 24 hours and found 5,000 crashes. After running a triage tool, it says you have "1 Unique Crash". Is this a failure?
    <details>
    <summary>Answer</summary>

    **No!** Finding even one bug is a success. The triage tool saved you time by telling you that you don't need to investigate 5,000 reports; you just need to fix that one bug. This is exactly why we triage!
    </details>
