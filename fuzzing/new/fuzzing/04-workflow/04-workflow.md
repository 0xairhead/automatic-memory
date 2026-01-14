# Lesson 4: Cleaning Your Room (Workflow) 🧹

Fuzzing creates a LOT of mess. Millions of files. We need to clean up.

## 1. Minimization (Packing the Suitcase) 🧳
Imagine you are going on a trip.
- You have 50 red t-shirts.
- You have 1 blue shirt.
- You have 1 green shirt.

Do you pack all 52 shirts? No! You pack **1 red**, **1 blue**, and **1 green**.
**Corpus Minimization** deletes the 49 extra red shirts. It keeps only the files that do something *new*.

**Demo**: `python3 minimizer.py` checks which files are unique and throws away the copies.

## 2. Triage (Grouping Accidents) 🚑
Imagine 100 people slip on the same banana peel.
Do you have 100 different problems? No! You have **1 problem** (the banana).

**Crash Triage** groups all the crashes that happened because of the "banana" into one bucket.
This way, you don't have to fix 100 bugs, you just fix the banana.

**Demo**: `python3 triage.py` looks at your crash reports and sorts them into piles.
