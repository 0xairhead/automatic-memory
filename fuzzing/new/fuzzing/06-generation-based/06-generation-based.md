# Lesson 6: Speaking the Language (Grammar Fuzzing) 🗣️

## The Guard at the Door 💂
Some programs have a strict Guard at the door.
- If you say "Glarb!", the Guard kicks you out.
- If you say "Bloop!", the Guard kicks you out.
- You must say "Hello, my name is..." to get in.

## Monkeys Can't Speak 🙊
Our blindfolded monkey (Random Fuzzer) will never accidentally type "Hello, my name is...". It just babbles. The Guard blocks it every time.

## Teaching the Monkey to Speak 🦜
**Grammar Fuzzing** is like teaching the monkey a few words.
We give it a rule book (Grammar):
1.  Start with "Hello".
2.  Then say a name.

Now the monkey says: "Hello, my name is BOMB".
The Guard says "Come on in!"... and then the BOMB explodes inside! 💥

**Demo**: `python3 grammar_fuzzer.py` knows how to write proper XML tags so it can sneak past the checks and find the bug.
