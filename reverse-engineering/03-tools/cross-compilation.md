# Lesson: Cross-Compiling Windows Binaries on macOS

**Goal:** Learn how to compile Windows executables (`.exe`) directly from macOS. This is a critical skill for a reverse engineer because it allows you to:
1.  **Create your own test subjects:** Write C code to test specific hypotheses about how the compiler generates assembly.
2.  **Build custom tools:** Create configuration extractors or loaders that need to run on the target Windows environment.
3.  **Understand the build process:** Seeing how code maps to PE headers helps you deconstruct it later.

---

## 1. The Standard Way: MinGW-w64

MinGW-w64 is the classic, industry-standard toolchain for cross-compiling.

### Installation
```bash
brew install mingw-w64
```

### Compilation
**For 64-bit (x64):**
```bash
x86_64-w64-mingw32-gcc hello.c -o hello.exe
```

**For 32-bit (x86):**
```bash
i686-w64-mingw32-gcc hello.c -o hello.exe
```

*   **Pros:** Produces native binaries, standard GCC flags.
*   **Cons:** Toolchain management can be messy on macOS.

---

## 2. The Modern Way: Zig

Zig is an emergent language that includes a powerful C/C++ compiler (`zig cc`) with built-in cross-compilation support for almost every architecture. It is often cleaner than installing the full MinGW suite.

### Installation
```bash
brew install zig
```

### Compilation
**For 64-bit (x64):**
```bash
zig cc hello.c -target x86_64-windows-gnu -o hello.exe
```

**For 32-bit (x86):**
```bash
zig cc hello.c -target i386-windows-gnu -o hello.exe
```

*   **Pros:** Zero dependency mess (ships with its own libc/linker), extremely fast, reproducible.
*   **Cons:** Newer tool, might require specific version pinning for complex builds.

---

## 3. The CI/Reproduction Way: Docker

If you need to guarantee that your build environment is identical to your team's, use a Docker container.

### Usage
```bash
# Drop into a shell with the toolchain ready
docker run -it --rm -v "$PWD:/src" dockcross/windows-static-x64
```
Inside the container:
```bash
x86_64-w64-mingw32-gcc hello.c -o hello.exe
```

---

## 4. Testing Your Binaries

You don't need to copy files to a Windows machine for every little test.

### Quick Check: Wine
Wine allows you to run Windows binaries on macOS/Linux.
```bash
brew install wine-stable
wine hello.exe
```
*Note: Wine is great for logic tests, but acts differently than a real kernel for advanced malware techniques.*

### Full Verification: VM
For "real" testing (especially if you are using Windows API calls for process injection or kernel interaction), you must use a VM (UTM, VMWare Fusion, or Parallels).

---

## 5. Reverse Engineering Context

When writing code for RE practice, you will often need to link against Windows system libraries.

**Example Code (`test.c`):**
```c
#include <windows.h>
#include <stdio.h>

int main() {
    MessageBoxA(NULL, "Hello from macOS!", "Cross-Compile", MB_OK);
    return 0;
}
```

**Linking:**
Sometimes you need to explicitly link libraries:
```bash
x86_64-w64-mingw32-gcc test.c -o test.exe -midl -luser32 -lkernel32
```
