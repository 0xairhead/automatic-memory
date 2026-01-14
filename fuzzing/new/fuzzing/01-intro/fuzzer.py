import subprocess
import random
import string
import time

def generate_input():
    # Occasionally generate the crash string to ensure demo works quickly
    if random.random() < 0.1:
        return "crash\n"
    length = random.randint(1, 10)
    return ''.join(random.choices(string.ascii_lowercase + '\n', k=length))

def fuzz():
    print("Starting fuzzer...")
    count = 0
    start_time = time.time()
    while True:
        count += 1
        data = generate_input()
        print(f"Trying input: {repr(data)}")
        
        # Run the target program
        process = subprocess.Popen(
            ['./vuln'], 
            stdin=subprocess.PIPE, 
            stdout=subprocess.DEVNULL, 
            stderr=subprocess.PIPE,
            text=True
        )
        stdout, stderr = process.communicate(input=data)
        
        if process.returncode != 0:
            print(f"CRASH FOUND! Input: {repr(data)}")
            print(f"Return code: {process.returncode}")
            break
        
        if count % 100 == 0:
            print(f"Runs: {count}")

        if time.time() - start_time > 10:
             print("Timeout reached without crash (unlikely with this setup).")
             break

if __name__ == "__main__":
    fuzz()
