# Variable Scope Rules in C

## Table of Contents

- [What is Scope?](#what-is-scope)
- [Types of Scope in C](#types-of-scope-in-c)
- [Shadowing (Variable Hiding)](#shadowing-variable-hiding)
- [Storage Duration vs Scope](#storage-duration-vs-scope)
- [Scope Rules Summary](#scope-rules-summary)
- [Global vs Local Variables](#global-vs-local-variables)
- [Static Variables in Different Scopes](#static-variables-in-different-scopes)
- [Extern - Sharing Variables Across Files](#extern---sharing-variables-across-files)
- [Common Scope Mistakes](#common-scope-mistakes)
- [Best Practices](#best-practices)
- [Scope Visualization](#scope-visualization)
- [Quick Reference](#quick-reference)

---

## What is Scope?

**Scope** determines where a variable can be accessed in your program. It defines the **visibility** and **lifetime** of variables.

---

## Types of Scope in C

### 1. Block Scope (Local Scope)

Variables declared inside a block `{ }` are only accessible within that block.

```c
#include <stdio.h>

int main() {
    int x = 10;  // Block scope - only in main()
    
    printf("%d\n", x);  // ✅ Works
    
    {
        int y = 20;  // Block scope - only in this inner block
        printf("%d\n", x);  // ✅ Can access x from outer block
        printf("%d\n", y);  // ✅ Works
    }
    
    // printf("%d\n", y);  // ❌ ERROR: y is out of scope
    
    return 0;
}
```

#### if/else Blocks

```c
int main() {
    int x = 10;
    
    if (x > 5) {
        int y = 20;  // Only exists inside if block
        printf("%d\n", y);  // ✅ Works
    }
    
    // printf("%d\n", y);  // ❌ ERROR: y doesn't exist here
    
    return 0;
}
```

#### for Loop Blocks

```c
int main() {
    // C99 and later
    for (int i = 0; i < 5; i++) {  // i only exists in loop
        printf("%d ", i);
    }
    
    // printf("%d", i);  // ❌ ERROR: i is out of scope
    
    // C89 style
    int j;  // Declare outside if needed later
    for (j = 0; j < 5; j++) {
        printf("%d ", j);
    }
    printf("%d", j);  // ✅ Works: j is still in scope
    
    return 0;
}
```

#### while Loop Blocks

```c
int main() {
    int count = 0;
    
    while (count < 5) {
        int temp = count * 2;  // Only exists in while block
        printf("%d\n", temp);
        count++;
    }
    
    // printf("%d", temp);  // ❌ ERROR: temp is out of scope
    
    return 0;
}
```

---

### 2. Function Scope

Variables declared in a function are only accessible within that function.

```c
#include <stdio.h>

void functionA() {
    int x = 10;  // Local to functionA
    printf("A: %d\n", x);
}

void functionB() {
    int x = 20;  // Different variable, local to functionB
    printf("B: %d\n", x);
}

int main() {
    functionA();  // Prints: A: 10
    functionB();  // Prints: B: 20
    
    // printf("%d", x);  // ❌ ERROR: x doesn't exist in main
    
    return 0;
}
```

#### Function Parameters

```c
void greet(char name[]) {  // name is local to greet()
    printf("Hello, %s!\n", name);
}

int main() {
    greet("Alice");
    // printf("%s", name);  // ❌ ERROR: name doesn't exist here
    return 0;
}
```

---

### 3. File Scope (Global Scope)

Variables declared outside all functions are **global** - accessible from anywhere in the file (or across files with `extern`).

```c
#include <stdio.h>

int globalVar = 100;  // File scope - global variable

void function1() {
    printf("Function1: %d\n", globalVar);  // ✅ Can access
    globalVar = 200;  // ✅ Can modify
}

void function2() {
    printf("Function2: %d\n", globalVar);  // ✅ Can access
}

int main() {
    printf("Main: %d\n", globalVar);  // ✅ Can access
    function1();
    function2();
    return 0;
}
```

**Output:**
```
Main: 100
Function1: 100
Function2: 200
```

---

### 4. Function Prototype Scope

Parameters in function prototypes have scope only within that prototype.

```c
// Function prototype
int add(int a, int b);  // a and b exist only in this line

// Function definition - these are different a and b
int add(int a, int b) {
    return a + b;
}

int main() {
    int result = add(5, 3);
    // printf("%d", a);  // ❌ ERROR: a doesn't exist
    return 0;
}
```

---

## Shadowing (Variable Hiding)

When an inner scope variable has the same name as an outer scope variable, it **shadows** (hides) the outer one.

```c
#include <stdio.h>

int x = 100;  // Global x

int main() {
    printf("Global x: %d\n", x);  // 100
    
    int x = 50;  // Local x shadows global x
    printf("Local x: %d\n", x);   // 50
    
    {
        int x = 25;  // Inner x shadows outer local x
        printf("Inner x: %d\n", x);  // 25
    }
    
    printf("Local x again: %d\n", x);  // 50
    
    return 0;
}
```

**Output:**
```
Global x: 100
Local x: 50
Inner x: 25
Local x again: 50
```

### Accessing Shadowed Global Variables

```c
#include <stdio.h>

int x = 100;  // Global

int main() {
    int x = 50;  // Local shadows global
    
    printf("Local x: %d\n", x);      // 50
    printf("Global x: %d\n", ::x);   // ❌ Doesn't work in C (works in C++)
    
    // In C, you cannot access shadowed global variable
    // Best practice: avoid shadowing!
    
    return 0;
}
```

---

## Storage Duration vs Scope

### Automatic Storage Duration

Variables with **block scope** are created and destroyed automatically.

```c
void func() {
    int x = 10;  // Created when function is called
    // ...
}  // x is destroyed when function returns
```

### Static Storage Duration

Variables with **static** keyword exist for the entire program lifetime but maintain their scope.

```c
#include <stdio.h>

void counter() {
    static int count = 0;  // Block scope, but static storage
    count++;
    printf("%d\n", count);
}

int main() {
    counter();  // 1
    counter();  // 2
    counter();  // 3
    
    // printf("%d", count);  // ❌ ERROR: count not in scope
    
    return 0;
}
```

---

## Scope Rules Summary

### Rule 1: Inner Scope Can Access Outer Scope

```c
int x = 10;  // Outer

int main() {
    printf("%d", x);  // ✅ Can access outer x
    
    {
        printf("%d", x);  // ✅ Can still access outer x
    }
    
    return 0;
}
```

### Rule 2: Outer Scope Cannot Access Inner Scope

```c
int main() {
    {
        int x = 10;  // Inner
    }
    
    // printf("%d", x);  // ❌ Cannot access inner x
    
    return 0;
}
```

### Rule 3: Same-Level Scopes Are Independent

```c
int main() {
    {
        int x = 10;
    }
    
    {
        int x = 20;  // Different variable, same name OK
    }
    
    return 0;
}
```

### Rule 4: Variables Exist Until End of Scope

```c
int main() {
    int x = 10;
    
    {
        int y = 20;
        // Both x and y exist here
    }  // y is destroyed here
    
    // Only x exists here
    
    return 0;
}  // x is destroyed here
```

---

## Global vs Local Variables

### Comparison Table

| Feature | Global Variables | Local Variables |
|---------|-----------------|-----------------|
| **Declaration** | Outside all functions | Inside function/block |
| **Scope** | Entire program | Only in declaration block |
| **Lifetime** | Entire program execution | Until block ends |
| **Default Value** | 0 (initialized automatically) | Garbage (uninitialized) |
| **Memory Location** | Data segment | Stack |
| **Access Speed** | Slower | Faster |
| **Best Practice** | Avoid when possible | Prefer these |

### Example Comparison

```c
#include <stdio.h>

int globalCount = 0;  // Global - automatically initialized to 0

void increment() {
    int localCount;  // Local - contains garbage value!
    localCount = 0;  // Must initialize manually
    
    globalCount++;
    localCount++;
    
    printf("Global: %d, Local: %d\n", globalCount, localCount);
}

int main() {
    increment();  // Global: 1, Local: 1
    increment();  // Global: 2, Local: 1 (localCount reset)
    increment();  // Global: 3, Local: 1 (localCount reset)
    return 0;
}
```

---

## Static Variables in Different Scopes

### Static Local Variable

```c
void func() {
    static int x = 0;  // Local scope, but persistent value
    x++;
    printf("%d\n", x);
}

int main() {
    func();  // 1
    func();  // 2
    func();  // 3
    return 0;
}
```

### Static Global Variable

```c
// file1.c
static int count = 0;  // Only visible in file1.c

void increment() {
    count++;
}
```

```c
// file2.c
extern int count;  // ❌ ERROR: count is static in file1.c
```

---

## Extern - Sharing Variables Across Files

### file1.c
```c
int sharedVar = 100;  // Definition

void printFromFile1() {
    printf("File1: %d\n", sharedVar);
}
```

### file2.c
```c
extern int sharedVar;  // Declaration (not definition)

void printFromFile2() {
    printf("File2: %d\n", sharedVar);  // ✅ Can access
}
```

### main.c
```c
extern int sharedVar;

int main() {
    printf("Main: %d\n", sharedVar);  // ✅ Can access
    sharedVar = 200;  // ✅ Can modify
    return 0;
}
```

---

## Common Scope Mistakes

### Mistake 1: Using Variable Outside Scope

```c
if (x > 10) {
    int result = x * 2;
}

printf("%d", result);  // ❌ ERROR: result out of scope
```

**Fix:**
```c
int result;  // Declare outside

if (x > 10) {
    result = x * 2;
}

printf("%d", result);  // ✅ Works
```

### Mistake 2: Uninitialized Local Variables

```c
void func() {
    int x;  // ❌ Contains garbage
    printf("%d", x);  // Undefined behavior!
}
```

**Fix:**
```c
void func() {
    int x = 0;  // ✅ Always initialize
    printf("%d", x);
}
```

### Mistake 3: Excessive Global Variables

```c
// ❌ BAD: Too many globals
int user_count;
int admin_count;
int guest_count;
char current_user[50];

void processUser() {
    // Modifies globals - hard to track bugs
}
```

**Fix:**
```c
// ✅ GOOD: Use structures and pass parameters
typedef struct {
    int user_count;
    int admin_count;
    int guest_count;
    char current_user[50];
} UserData;

void processUser(UserData* data) {
    // Explicit parameter - easier to debug
}
```

---

## Best Practices

1. **Minimize global variables** - use local when possible
2. **Declare variables close to first use** - improves readability
3. **Avoid shadowing** - don't reuse names in nested scopes
4. **Initialize variables** - local variables don't auto-initialize
5. **Use static for persistent local state** - better than globals
6. **Limit variable scope** - smallest scope possible
7. **Use meaningful names** - especially for globals

```c
// ✅ GOOD
int main() {
    // Declare close to use
    int count = getUserCount();
    
    for (int i = 0; i < count; i++) {
        // Process users
    }
    
    // Another variable with limited scope
    double average = calculateAverage();
    
    return 0;
}
```

```c
// ❌ BAD
int count, i;
double average;

int main() {
    count = getUserCount();
    
    for (i = 0; i < count; i++) {
        // Process users
    }
    
    average = calculateAverage();
    
    return 0;
}
```

---

## Scope Visualization

```
┌─────────────────────────────────────────┐
│ FILE SCOPE (Global)                     │
│ int globalVar = 100;                    │
│                                         │
│ ┌─────────────────────────────────────┐ │
│ │ FUNCTION SCOPE (main)               │ │
│ │ int main() {                        │ │
│ │     int x = 10;                     │ │
│ │                                     │ │
│ │     ┌─────────────────────────────┐ │ │
│ │     │ BLOCK SCOPE                 │ │ │
│ │     │ {                           │ │ │
│ │     │     int y = 20;             │ │ │
│ │     │     // Can access: y, x,    │ │ │
│ │     │     //    globalVar         │ │ │
│ │     │ }                           │ │ │
│ │     └─────────────────────────────┘ │ │
│ │                                     │ │
│ │     // Can access: x, globalVar    │ │
│ │     // Cannot access: y            │ │
│ │ }                                   │ │
│ └─────────────────────────────────────┘ │
│                                         │
│ // Can access: globalVar only          │
└─────────────────────────────────────────┘
```

---

## Quick Reference

| Scope Type | Declared Where | Accessible From | Lifetime |
|------------|---------------|-----------------|----------|
| **Block** | Inside `{ }` | Inside that block only | Until block ends |
| **Function** | Inside function | That function only | Until function returns |
| **File (Global)** | Outside all functions | Entire file (or program with extern) | Entire program |
| **Static Local** | Inside function with `static` | That function only | Entire program |
| **Static Global** | Outside functions with `static` | That file only | Entire program |

Understanding scope is crucial for writing bug-free C code!