import subprocess
import random
import string
import time
import sys

def generate_input():
    length = 5
    # Pure random guessing
    return ''.join(random.choices(string.ascii_letters + '!', k=length))

def fuzz():
    print("Starting DUMB fuzzer...")
    count = 0
    start_time = time.time()
    while True:
        count += 1
        data = generate_input()
        
        process = subprocess.Popen(
            ['./maze'], 
            stdin=subprocess.PIPE, 
            stdout=subprocess.PIPE, 
            stderr=subprocess.PIPE,
            text=True
        )
        stdout, stderr = process.communicate(input=data)
        
        if process.returncode != 0:
            print(f"CRASH FOUND! Input: {repr(data)}")
            break
        
        if count % 1000 == 0:
            elapsed = time.time() - start_time
            print(f"Runs: {count} | Speed: {count/elapsed:.2f} exec/s")

        if time.time() - start_time > 10:
             print("Timeout: DUMB fuzzer failed to find the crash in 10s.")
             break

if __name__ == "__main__":
    fuzz()
