# Lesson 2: Playing "Hot or Cold" (Coverage-guided Fuzzing)

## The Problem with Monkeys 🙈
The monkey approach works for simple passwords like "crash". But what if the password is a maze?
- Do `M` -> "Good job!"
- Do `a` -> "Good job!"
- Do `z` -> "Good job!"

A random monkey will never guess "M-a-z-e-!" in the right order. It's too hard!

## Playing "Hot or Cold" 🔥
**Coverage-guided Fuzzing** is like playing "Hot or Cold".
1.  The fuzzer guesses "Q". The computer says "Cold" (ignored).
2.  The fuzzer guesses "M". The computer says "**HOT!**" (You passed level 1!).
3.  The fuzzer remembers "M" and tries "Ma", "Mb", "Mc"...
4.  It finds "Ma" is "**HOTTER!**".

It learns step-by-step!

## The Experiment
We built a `maze.c` (the game) and a `smart_fuzzer.py` (the player).
Run it and watch it solve the puzzle:

```bash
cd 02-coverage-guided
gcc maze.c -o maze
python3 smart_fuzzer.py
```
