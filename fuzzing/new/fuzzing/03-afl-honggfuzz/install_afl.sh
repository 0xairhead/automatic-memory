#!/bin/bash
set -e

echo "[*] Cloning AFLplusplus..."
if [ ! -d "AFLplusplus" ]; then
    git clone https://github.com/AFLplusplus/AFLplusplus
else
    echo "[-] AFLplusplus already exists, skipping clone."
fi

cd AFLplusplus

echo "[*] Building AFLplusplus (source-only to minimize deps)..."
# We use source-only to avoid needing python/qemu/unicorn dependencies if they are missing
make source-only

echo "[+] Build complete!"
echo "    afl-fuzz is at: $(pwd)/afl-fuzz"
echo "    afl-cc is at:   $(pwd)/afl-cc"
