# Lesson 3: The Super Robot (AFL++) 🤖

## The Professional Tool
We wrote our own fun fuzzers in Python, but real hackers use tools like **AFL++**.
Think of AFL++ as a Super Robot version of our "Hot or Cold" player.

## Motion Sensors (Instrumentation) 🚨
How does AFL++ know it's getting "Hotter"?
It cheats! It puts **Motion Sensors** on every door in your program (this is called "Instrumentation").

When you compile with `afl-cc` instead of `gcc`, it adds these sensors.
Now, when the fuzzer tries an input, the program whispers: *"Hey, he just opened the red door!"*

## The Workflow
1.  **Build** with sensors: `afl-cc harness.c -o harness`
2.  **Give a hint** (Seed): `echo "AAAA" > inputs/seed.txt`
3.  **Run the Robot**: `afl-fuzz -i inputs -o out -- ./harness`

Watch the robot dashboard screen. "Saved Crashes" means the robot won!
