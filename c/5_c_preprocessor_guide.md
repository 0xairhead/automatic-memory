# C Preprocessor

## Table of Contents

- [What is the C Preprocessor?](#what-is-the-c-preprocessor)
- [Preprocessor Directives](#preprocessor-directives)
- [1. #include - File Inclusion](#1-include---file-inclusion)
- [2. #define - Macro Definition](#2-define---macro-definition)
- [3. #undef - Undefine Macro](#3-undef---undefine-macro)
- [4. Conditional Compilation](#4-conditional-compilation)
- [5. Include Guards](#5-include-guards)
- [6. Predefined Macros](#6-predefined-macros)
- [7. Stringification (#)](#7-stringification-)
- [8. Token Pasting (##)](#8-token-pasting-)
- [9. #error and #warning](#9-error-and-warning)
- [10. #pragma](#10-pragma)
- [Practical Examples](#practical-examples)
- [View Preprocessor Output](#view-preprocessor-output)
- [Macros vs Functions](#macros-vs-functions)
- [Best Practices](#best-practices)
- [Common Preprocessor Patterns](#common-preprocessor-patterns)
- [Summary](#summary)

---

## What is the C Preprocessor?

The **preprocessor** is a program that processes your source code **before** compilation. It handles all directives that start with `#`.

```
Source Code (.c) → [PREPROCESSOR] → Expanded Code → [COMPILER] → Object Code
```

### Preprocessor Flow

```
program.c  →  Preprocessor  →  program.i  →  Compiler  →  program.o
             (handles #...)      (expanded)    (compile)    (binary)
```

---

## Preprocessor Directives

All preprocessor directives start with `#` and do NOT end with a semicolon.

```c
#include <stdio.h>     // ✅ No semicolon
#define MAX 100        // ✅ No semicolon
```

---

## 1. #include - File Inclusion

Includes the contents of another file into your source code.

### System Headers (Angle Brackets)

```c
#include <stdio.h>      // Standard I/O
#include <stdlib.h>     // Standard library
#include <string.h>     // String functions
#include <math.h>       // Math functions
```

Searches in **system directories** (e.g., `/usr/include/`)

### User Headers (Quotes)

```c
#include "myheader.h"   // User-defined header
#include "utils.h"      // Your custom file
#include "../lib/helper.h"  // Relative path
```

Searches in **current directory first**, then system directories.

### What Happens During #include

```c
// main.c
#include "header.h"

int main() {
    return 0;
}
```

```c
// header.h
#define MAX 100
void myFunction();
```

**After preprocessing:**
```c
// Expanded main.c
#define MAX 100
void myFunction();

int main() {
    return 0;
}
```

---

## 2. #define - Macro Definition

Creates macros (text substitution).

### Simple Macros (Constants)

```c
#define PI 3.14159
#define MAX_SIZE 100
#define BUFFER_SIZE 1024
#define TRUE 1
#define FALSE 0

int main() {
    float area = PI * radius * radius;  // Replaced with 3.14159
    int array[MAX_SIZE];                // Replaced with 100
}
```

### Function-like Macros

```c
#define SQUARE(x) ((x) * (x))
#define MAX(a, b) ((a) > (b) ? (a) : (b))
#define MIN(a, b) ((a) < (b) ? (a) : (b))
#define ABS(x) ((x) < 0 ? -(x) : (x))

int main() {
    int result = SQUARE(5);        // Expands to: ((5) * (5))
    int max = MAX(10, 20);         // Expands to: ((10) > (20) ? (10) : (20))
    int absolute = ABS(-15);       // Expands to: ((-15) < 0 ? -(-15) : (-15))
}
```

### ⚠️ Macro Pitfalls

```c
// ❌ BAD: No parentheses
#define SQUARE(x) x * x

int result = SQUARE(2 + 3);  // Expands to: 2 + 3 * 2 + 3 = 11 (not 25!)

// ✅ GOOD: Use parentheses
#define SQUARE(x) ((x) * (x))

int result = SQUARE(2 + 3);  // Expands to: ((2 + 3) * (2 + 3)) = 25
```

```c
// ❌ BAD: Side effects
#define SQUARE(x) ((x) * (x))

int a = 5;
int result = SQUARE(a++);  // Expands to: ((a++) * (a++))
// a is incremented TWICE! Unpredictable behavior
```

### Multi-line Macros

```c
#define SWAP(a, b) do { \
    int temp = a;       \
    a = b;              \
    b = temp;           \
} while(0)

int main() {
    int x = 5, y = 10;
    SWAP(x, y);
    printf("x=%d, y=%d", x, y);  // x=10, y=5
}
```

---

## 3. #undef - Undefine Macro

Removes a previously defined macro.

```c
#define MAX 100
printf("%d", MAX);  // 100

#undef MAX
// printf("%d", MAX);  // ❌ ERROR: MAX not defined

#define MAX 200
printf("%d", MAX);  // 200
```

---

## 4. Conditional Compilation

### #ifdef / #ifndef

```c
#define DEBUG

#ifdef DEBUG
    printf("Debug mode is ON\n");
#endif

#ifndef RELEASE
    printf("Not in release mode\n");
#endif
```

### #if, #elif, #else, #endif

```c
#define VERSION 2

#if VERSION == 1
    printf("Version 1\n");
#elif VERSION == 2
    printf("Version 2\n");
#else
    printf("Unknown version\n");
#endif
```

### defined() Operator

```c
#if defined(DEBUG) && defined(VERBOSE)
    printf("Debug and verbose mode\n");
#endif

#if !defined(RELEASE)
    printf("Not release build\n");
#endif
```

### Practical Example: Platform-Specific Code

```c
#ifdef _WIN32
    #include <windows.h>
    #define CLEAR_SCREEN "cls"
#elif defined(__linux__)
    #include <unistd.h>
    #define CLEAR_SCREEN "clear"
#elif defined(__APPLE__)
    #include <unistd.h>
    #define CLEAR_SCREEN "clear"
#else
    #error "Unsupported platform"
#endif
```

---

## 5. Include Guards

Prevents multiple inclusion of the same header file.

### Problem Without Include Guards

```c
// math_utils.h
#define PI 3.14159
void calculate();

// main.c
#include "math_utils.h"
#include "math_utils.h"  // ❌ ERROR: PI redefined!
```

### Solution: Include Guards

```c
// math_utils.h
#ifndef MATH_UTILS_H
#define MATH_UTILS_H

#define PI 3.14159
void calculate();

#endif  // MATH_UTILS_H
```

**How it works:**
1. First inclusion: `MATH_UTILS_H` not defined → define it and include content
2. Second inclusion: `MATH_UTILS_H` already defined → skip content

### Modern Alternative: #pragma once

```c
// math_utils.h
#pragma once  // Simpler, but not standard C

#define PI 3.14159
void calculate();
```

---

## 6. Predefined Macros

The C preprocessor provides several built-in macros:

### Standard Predefined Macros

```c
#include <stdio.h>

int main() {
    printf("File: %s\n", __FILE__);        // Current file name
    printf("Line: %d\n", __LINE__);        // Current line number
    printf("Date: %s\n", __DATE__);        // Compilation date
    printf("Time: %s\n", __TIME__);        // Compilation time
    printf("Standard: %ld\n", __STDC__);   // 1 if standard C
    
    return 0;
}
```

**Output:**
```
File: main.c
Line: 5
Date: Dec 22 2024
Time: 14:30:15
Standard: 1
```

### Useful for Debugging

```c
#define DEBUG_PRINT(x) printf("DEBUG [%s:%d]: %s = %d\n", \
                              __FILE__, __LINE__, #x, x)

int main() {
    int value = 42;
    DEBUG_PRINT(value);
    // Output: DEBUG [main.c:10]: value = 42
}
```

---

## 7. Stringification (#)

Converts macro parameter to a string literal.

```c
#define TO_STRING(x) #x

printf("%s\n", TO_STRING(Hello));     // "Hello"
printf("%s\n", TO_STRING(123));       // "123"
printf("%s\n", TO_STRING(a + b));     // "a + b"
```

### Practical Use: Printing Variable Names

```c
#define PRINT_VAR(var) printf(#var " = %d\n", var)

int main() {
    int age = 25;
    int count = 100;
    
    PRINT_VAR(age);      // Prints: age = 25
    PRINT_VAR(count);    // Prints: count = 100
}
```

---

## 8. Token Pasting (##)

Concatenates two tokens into one.

```c
#define CONCAT(a, b) a##b

int xy = 100;
int result = CONCAT(x, y);  // Expands to: xy
printf("%d", result);        // 100
```

### Creating Variable Names

```c
#define CREATE_VAR(name, num) int name##num

CREATE_VAR(value, 1) = 10;   // int value1 = 10;
CREATE_VAR(value, 2) = 20;   // int value2 = 20;

printf("%d %d", value1, value2);  // 10 20
```

---

## 9. #error and #warning

### #error - Stop Compilation

```c
#if !defined(PLATFORM)
    #error "PLATFORM must be defined!"
#endif

#if VERSION < 2
    #error "Minimum version 2 required"
#endif
```

### #warning - Issue Warning

```c
#ifndef BUFFER_SIZE
    #warning "BUFFER_SIZE not defined, using default"
    #define BUFFER_SIZE 1024
#endif
```

---

## 10. #pragma

Compiler-specific directives (non-standard).

```c
#pragma once  // Include guard (modern compilers)

#pragma pack(1)  // Structure packing

#pragma GCC diagnostic ignored "-Wunused-variable"  // Ignore warnings
```

---

## Practical Examples

### Example 1: Debug Macro

```c
#ifdef DEBUG
    #define LOG(msg) printf("[DEBUG] %s:%d - %s\n", __FILE__, __LINE__, msg)
#else
    #define LOG(msg)  // Empty in release build
#endif

int main() {
    LOG("Program started");
    int x = 10;
    LOG("Processing data");
    return 0;
}
```

### Example 2: Assert Macro

```c
#ifdef DEBUG
    #define ASSERT(condition) \
        if (!(condition)) { \
            fprintf(stderr, "Assertion failed: %s, file %s, line %d\n", \
                    #condition, __FILE__, __LINE__); \
            exit(1); \
        }
#else
    #define ASSERT(condition)
#endif

int main() {
    int x = 5;
    ASSERT(x > 0);   // Passes
    ASSERT(x > 10);  // Fails in debug mode
}
```

### Example 3: Platform Detection

```c
#if defined(_WIN32) || defined(_WIN64)
    #define PLATFORM "Windows"
    #define PATH_SEPARATOR '\\'
#elif defined(__linux__)
    #define PLATFORM "Linux"
    #define PATH_SEPARATOR '/'
#elif defined(__APPLE__)
    #define PLATFORM "macOS"
    #define PATH_SEPARATOR '/'
#else
    #define PLATFORM "Unknown"
    #define PATH_SEPARATOR '/'
#endif

int main() {
    printf("Running on: %s\n", PLATFORM);
    printf("Path separator: %c\n", PATH_SEPARATOR);
}
```

### Example 4: Configuration System

```c
// config.h
#define VERSION_MAJOR 1
#define VERSION_MINOR 2
#define VERSION_PATCH 3

#define FEATURE_NETWORKING 1
#define FEATURE_GRAPHICS 0

#if VERSION_MAJOR >= 2
    #define NEW_API_AVAILABLE
#endif

// main.c
#include "config.h"

int main() {
    printf("Version: %d.%d.%d\n", VERSION_MAJOR, VERSION_MINOR, VERSION_PATCH);
    
    #ifdef FEATURE_NETWORKING
        printf("Networking enabled\n");
    #endif
    
    #ifdef NEW_API_AVAILABLE
        printf("Using new API\n");
    #endif
}
```

---

## View Preprocessor Output

### Using GCC

```bash
# Generate preprocessed file (.i)
gcc -E program.c -o program.i

# View preprocessed output directly
gcc -E program.c
```

### Example

```c
// program.c
#include <stdio.h>
#define MAX 100

int main() {
    int arr[MAX];
    return 0;
}
```

```bash
gcc -E program.c
```

**Output shows expanded code:**
```c
// ... thousands of lines from stdio.h ...

int main() {
    int arr[100];
    return 0;
}
```

---

## Macros vs Functions

| Feature | Macros | Functions |
|---------|--------|-----------|
| **Type Safety** | ❌ No type checking | ✅ Type checked |
| **Debugging** | ❌ Hard to debug | ✅ Easy to debug |
| **Code Size** | ⚠️ Increases (inline) | ✅ Smaller |
| **Speed** | ✅ Fast (no call overhead) | ⚠️ Function call overhead |
| **Side Effects** | ❌ Can cause issues | ✅ Safe |
| **Recursion** | ❌ Not possible | ✅ Possible |

### When to Use Macros

```c
// ✅ Constants
#define MAX_SIZE 100
#define PI 3.14159

// ✅ Simple operations
#define SQUARE(x) ((x) * (x))
#define MAX(a, b) ((a) > (b) ? (a) : (b))

// ✅ Conditional compilation
#ifdef DEBUG
#endif

// ✅ Code generation
#define CREATE_GETTER(type, name) \
    type get_##name() { return name; }
```

### When to Use Functions

```c
// ✅ Complex logic
int calculateFactorial(int n) {
    if (n <= 1) return 1;
    return n * calculateFactorial(n - 1);
}

// ✅ Type safety needed
double divide(double a, double b) {
    if (b == 0) {
        fprintf(stderr, "Division by zero!\n");
        return 0.0;
    }
    return a / b;
}
```

---

## Best Practices

1. **Use UPPERCASE for macros** - `#define MAX_SIZE 100`
2. **Always use parentheses in macro expressions** - `#define SQUARE(x) ((x) * (x))`
3. **Prefer const over #define for constants** - `const int MAX = 100;`
4. **Use include guards in all headers**
5. **Avoid complex macros** - use inline functions instead
6. **Don't use macros with side effects**
7. **Document macro behavior**
8. **Use conditional compilation sparingly**

```c
// ✅ GOOD
#define SQUARE(x) ((x) * (x))
const int MAX_SIZE = 100;

// ❌ BAD
#define square(x) x*x  // No parentheses, lowercase
#define max 100        // Lowercase constant
```

---

## Common Preprocessor Patterns

### Configuration Header

```c
// config.h
#ifndef CONFIG_H
#define CONFIG_H

#define APP_NAME "MyApp"
#define APP_VERSION "1.0.0"
#define DEBUG_MODE 1
#define MAX_CONNECTIONS 100

#if DEBUG_MODE
    #define LOG_LEVEL 3
#else
    #define LOG_LEVEL 1
#endif

#endif
```

### Platform Abstraction

```c
// platform.h
#ifdef _WIN32
    #define EXPORT __declspec(dllexport)
    #define IMPORT __declspec(dllimport)
#else
    #define EXPORT __attribute__((visibility("default")))
    #define IMPORT
#endif
```

---

## Summary

| Directive | Purpose | Example |
|-----------|---------|---------|
| `#include` | Include files | `#include <stdio.h>` |
| `#define` | Define macros | `#define MAX 100` |
| `#undef` | Undefine macro | `#undef MAX` |
| `#ifdef` | If defined | `#ifdef DEBUG` |
| `#ifndef` | If not defined | `#ifndef HEADER_H` |
| `#if` | Conditional | `#if VERSION > 1` |
| `#else` | Alternative | `#else` |
| `#elif` | Else if | `#elif VERSION == 2` |
| `#endif` | End conditional | `#endif` |
| `#error` | Compilation error | `#error "Error message"` |
| `#pragma` | Compiler directive | `#pragma once` |

The preprocessor is a powerful tool that runs before compilation - master it to write flexible, maintainable C code!