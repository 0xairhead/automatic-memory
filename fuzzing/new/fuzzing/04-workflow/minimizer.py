import os
import subprocess
import hashlib
import sys

def minimize(target_binary, input_dir):
    print(f"[*] Minimizing corpus in '{input_dir}' for target '{target_binary}'...")
    
    seen_behaviors = set()
    files_to_delete = []
    
    files = sorted(os.listdir(input_dir))
    
    for fname in files:
        fpath = os.path.join(input_dir, fname)
        
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
        
        # In real world, we use code coverage (bitmap).
        # Here, we use a hash of stdout to simulate "unique execution path".
        behavior_hash = hashlib.md5(stdout.encode()).hexdigest()
        
        if behavior_hash in seen_behaviors:
            print(f"[-] {fname} is redundant. Deleting.")
            files_to_delete.append(fpath)
        else:
            print(f"[+] {fname} triggered new behavior ({behavior_hash}). Keeping.")
            seen_behaviors.add(behavior_hash)
            
    # Cleanup
    for f in files_to_delete:
        os.remove(f)
        
    print(f"\n[*] Minimization complete. Removed {len(files_to_delete)} files.")

if __name__ == "__main__":
    if len(sys.argv) != 3:
        print("Usage: python3 minimizer.py <target_binary> <input_dir>")
    else:
        minimize(sys.argv[1], sys.argv[2])
