import subprocess
import random
import string
import time

# Genetic Algorithm Fuzzer

current_best_input = ""
max_level_reached = 0

def mutate(data):
    # Mutate the existing best input
    if len(data) < 5:
        data += random.choice(string.ascii_letters + '!')
    
    # 20% chance to change a character
    if random.random() < 0.2 and len(data) > 0:
        pos = random.randint(0, len(data)-1)
        chars = list(data)
        chars[pos] = random.choice(string.ascii_letters + '!')
        data = "".join(chars)
        
    return data

def fuzz():
    global current_best_input, max_level_reached
    print("Starting SMART fuzzer...")
    count = 0
    start_time = time.time()
    
    while True:
        count += 1
        
        # Strategy: Mutate our best guess so far
        data = mutate(current_best_input)
        
        process = subprocess.Popen(
            ['./maze'], 
            stdin=subprocess.PIPE, 
            stdout=subprocess.PIPE, 
            stderr=subprocess.PIPE,
            text=True
        )
        stdout, stderr = process.communicate(input=data)
        
        # Check for crash
        if process.returncode != 0:
            print(f"\nCRASH FOUND! Input: {repr(data)}")
            break
            
        # Check for coverage feedback (our "Instrumentation")
        level = 0
        if "Level 1 passed" in stdout: level = 1
        if "Level 2 passed" in stdout: level = 2
        if "Level 3 passed" in stdout: level = 3
        if "Level 4 passed" in stdout: level = 4
        
        if level > max_level_reached:
            print(f"New coverage! Level {level} with input: {repr(data)}")
            max_level_reached = level
            current_best_input = data # Save this as new base

        if count % 1000 == 0:
            elapsed = time.time() - start_time
            print(f"Runs: {count} | Speed: {count/elapsed:.2f} exec/s")

        if time.time() - start_time > 10:
             print("Timeout reached.")
             break

if __name__ == "__main__":
    fuzz()
