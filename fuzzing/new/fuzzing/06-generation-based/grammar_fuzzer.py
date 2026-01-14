import random
import sys
import subprocess
import time

# Simple Grammar
# S -> <tag key="KEY">VAL</tag>
# KEY -> safe | BOMB
# VAL -> string
# STR -> string

def generate_input():
    # 50% chance to pick the dangerous key
    if random.random() < 0.5:
        key = "BOMB"
    else:
        key = "safe"
    
    # Generate a random value
    # If key is BOMB, we want a long value to trigger overflow
    val_len = random.randint(1, 100)
    val = "A" * val_len
    
    data = "some_data"
    
    return f'<tag key="{key}">{data}</tag>'

def fuzz():
    print("Starting GRAMMAR fuzzer...")
    count = 0
    start_time = time.time()
    
    while True:
        count += 1
        data = generate_input()
        
        process = subprocess.Popen(
            ['./xml_parser'], 
            stdin=subprocess.PIPE, 
            stdout=subprocess.PIPE, 
            stderr=subprocess.PIPE,
            text=True
        )
        stdout, stderr = process.communicate(input=data)
        
        # Check for crash (return code -11 usually SIGSEGV)
        if process.returncode != 0:
            print(f"\nCRASH FOUND! Input: {data}")
            break
            
        if count % 1000 == 0:
            elapsed = time.time() - start_time
            print(f"Runs: {count} | Speed: {count/elapsed:.2f} exec/s")

        if time.time() - start_time > 10:
             print("Timeout reached.")
             break

if __name__ == "__main__":
    fuzz()
