# GCC Build Essentials Setup Guide

Quick reference for installing development tools across macOS, Fedora, and Ubuntu.

## Table of Contents

- [macOS](#macos)
- [Fedora](#fedora)
- [Ubuntu / Debian](#ubuntu--debian)
- [Quick Verification (All Systems)](#quick-verification-all-systems)
- [Common Development Libraries](#common-development-libraries)
- [Notes](#notes)

---

## macOS

### Install Command Line Tools (Recommended)
```bash
xcode-select --install
```

### Check Installation
```bash
gcc --version
make --version
git --version
```

### Install Homebrew (Optional)
```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

### Install GNU GCC via Homebrew (Optional)
```bash
brew install gcc
brew install cmake pkg-config autoconf automake
```

### Use GNU GCC
```bash
gcc-14 --version
g++-14 --version
```

---

## Fedora

### Update System
```bash
sudo dnf update
```

### Install Development Tools (Full)
```bash
sudo dnf groupinstall "Development Tools"
```

### Install Essential Tools (Minimal)
```bash
sudo dnf install gcc gcc-c++ make
```

### Additional Common Packages
```bash
sudo dnf install kernel-devel kernel-headers
sudo dnf install glibc-devel
sudo dnf install cmake
sudo dnf install pkgconfig
sudo dnf install git
```

### Verify Installation
```bash
gcc --version
g++ --version
make --version
```

---

## Ubuntu / Debian

### Update Package List
```bash
sudo apt update
```

### Install Build Essential (Recommended)
```bash
sudo apt install build-essential
```

### Install Essential Tools (Minimal)
```bash
sudo apt install gcc g++ make
```

### Additional Common Packages
```bash
sudo apt install linux-headers-$(uname -r)
sudo apt install libc6-dev
sudo apt install cmake
sudo apt install pkg-config
sudo apt install git
sudo apt install gdb
sudo apt install autoconf automake libtool
```

### Verify Installation
```bash
gcc --version
g++ --version
make --version
```

---

## Quick Verification (All Systems)

After installation on any system, verify with:

```bash
gcc --version
g++ --version
make --version
git --version
cmake --version
```

---

## Common Development Libraries

### macOS
```bash
brew install openssl readline sqlite3 xz zlib
```

### Fedora
```bash
sudo dnf install openssl-devel readline-devel sqlite-devel xz-devel zlib-devel
```

### Ubuntu
```bash
sudo apt install libssl-dev libreadline-dev libsqlite3-dev liblzma-dev zlib1g-dev
```

---

## Notes

- **macOS**: The `gcc` command actually points to Clang (Apple's compiler). For GNU GCC, use Homebrew and version-specific commands like `gcc-14`.
- **Fedora**: Uses DNF package manager (replaced YUM in recent versions).
- **Ubuntu/Debian**: Uses APT package manager. Commands work on both Ubuntu and Debian-based distros.
- Always run system updates before installing development tools.
- The "Development Tools" group (Fedora) and "build-essential" package (Ubuntu) include most common tools needed for compiling software.