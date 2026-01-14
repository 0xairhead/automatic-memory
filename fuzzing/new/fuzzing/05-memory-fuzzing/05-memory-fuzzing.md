# Lesson 5: The Speed of Light (In-Process Fuzzing) ⚡

## Starting the Car 🚗
Our old fuzzers were like this:
1.  Start the car (Launch program).
2.  Drive 1 inch (Test input).
3.  Turn off the car (Exit program).
4.  Repeat.

This is SLOW.

## Keeping it Running 🏎️
**libFuzzer** (In-Process Fuzzing) is like keeping the engine running.
1.  Start the car ONCE.
2.  Drive, Drive, Drive, Drive! (Test thousands of inputs).
3.  Never turn it off until it crash.

Because we don't restart the car every time, we can test **100,000 inputs per second**!

## Magic Glasses (AddressSanitizer) 👓
We also use special glasses called **AddressSanitizer (ASAN)**.
Normally, if a program makes a small mistake (like writing on the wrong line of a notebook), it keeps going.
With ASAN glasses, if the program makes even a TINY mistake, it screams **"ERROR!"** immediately.
