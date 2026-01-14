import os
import subprocess
import hashlib
import sys

def triage(target_binary, crash_dir):
    print(f"[*] Triaging crashes in '{crash_dir}'...")
    
    buckets = {}
    
    files = sorted(os.listdir(crash_dir))
    
    for fname in files:
        fpath = os.path.join(crash_dir, fname)
        
        # Run target
        with open(fpath, 'rb') as f:
            data = f.read()
            
        process = subprocess.Popen(
            [target_binary],
            stdin=subprocess.PIPE,
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
            text=True
        )
        stdout, stderr = process.communicate(input=data.decode(errors='ignore'))
        
        # Use the first line of output as the "crash signature" for this demo.
        # In reality, we would use stack traces (GDB) or unique crash paths (AFL).
        signature = "Unknown"
        if "Triggering" in stdout:
             signature = stdout.split('\n')[0].strip()
        else:
             # If it didn't print our debug trigger, maybe use return code
             signature = f"Return Code {process.returncode}"

        if signature not in buckets:
            buckets[signature] = []
        buckets[signature].append(fname)
        
    print("\n[*] Triage Report:")
    for sig, crashes in buckets.items():
        print(f"\nSignature: {sig}")
        print(f"  Count: {len(crashes)}")
        print(f"  Examples: {crashes[:3]}...")

if __name__ == "__main__":
    if len(sys.argv) != 3:
        print("Usage: python3 triage.py <target_binary> <crash_dir>")
    else:
        triage(sys.argv[1], sys.argv[2])
