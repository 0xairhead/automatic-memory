# Variables in C

## Table of Contents

- [What is a Variable?](#what-is-a-variable)
- [Variable Declaration and Initialization](#variable-declaration-and-initialization)
- [Data Types in C](#data-types-in-c)
- [Variable Naming Rules](#variable-naming-rules)
- [Variable Scope](#variable-scope)
- [Storage Classes](#storage-classes)
- [Constants](#constants)
- [Type Casting](#type-casting)
- [Sizeof Operator](#sizeof-operator)
- [Variable Examples](#variable-examples)
- [Common Mistakes](#common-mistakes)
- [Format Specifiers for printf/scanf](#format-specifiers-for-printfscanf)
- [Best Practices](#best-practices)
- [Summary](#summary)

---

## What is a Variable?

A **variable** is a named storage location in memory that holds a value. Think of it as a labeled box that stores data.

```c
int age = 25;  // Variable named 'age' storing the value 25
```

---

## Variable Declaration and Initialization

### Declaration
Telling the compiler about a variable and its type:

```c
int number;          // Declared but not initialized
float price;         // Contains garbage value
char letter;         // Uninitialized
```

### Initialization
Assigning a value when declaring:

```c
int number = 42;           // Declaration + Initialization
float price = 19.99;       // Good practice
char letter = 'A';         // Always initialize!
```

### Assignment
Giving a value to an already declared variable:

```c
int number;      // Declaration
number = 42;     // Assignment
number = 100;    // Re-assignment
```

---

## Data Types in C

### Integer Types

```c
// Signed integers (can be positive or negative)
char c = 'A';              // 1 byte  (-128 to 127)
short s = 32000;           // 2 bytes (-32,768 to 32,767)
int i = 100000;            // 4 bytes (-2.1B to 2.1B)
long l = 1000000000L;      // 4 or 8 bytes
long long ll = 9223372036854775807LL;  // 8 bytes

// Unsigned integers (only positive)
unsigned char uc = 255;    // 0 to 255
unsigned int ui = 4000000000U;  // 0 to 4.3B
```

### Floating Point Types

```c
float f = 3.14f;           // 4 bytes (6-7 decimal digits)
double d = 3.14159265359;  // 8 bytes (15-16 decimal digits)
long double ld = 3.141592653589793238L;  // 10-16 bytes
```

### Character Type

```c
char letter = 'A';         // Single character
char digit = '5';          // Character, not number 5!
char newline = '\n';       // Escape sequence
```

### Boolean Type (C99 and later)

```c
#include <stdbool.h>

bool isActive = true;      // true or false
bool isComplete = false;
```

---

## Variable Naming Rules

### Rules (MUST follow)

```c
// ✅ Valid variable names
int age;
int user_age;
int userAge;
int age2;
int _temp;
int MAX_SIZE;

// ❌ Invalid variable names
int 2age;          // Cannot start with digit
int user-age;      // Cannot use hyphen
int int;           // Cannot use keywords
int user age;      // Cannot have spaces
```

### Conventions (should follow)

```c
// Snake case (common in C)
int user_age;
int max_buffer_size;

// Camel case
int userAge;
int maxBufferSize;

// Constants (uppercase with underscore)
#define MAX_SIZE 100
const int BUFFER_SIZE = 1024;
```

---

## Variable Scope

### Local Variables
Declared inside a function or block:

```c
void myFunction() {
    int x = 10;  // Local to myFunction
    // x is only accessible here
}  // x is destroyed here

int main() {
    int y = 20;  // Local to main
    // Cannot access x here!
    
    if (y > 10) {
        int z = 30;  // Local to this if block
    }
    // Cannot access z here!
}
```

### Global Variables
Declared outside all functions:

```c
int globalVar = 100;  // Global variable

void function1() {
    printf("%d", globalVar);  // Can access
}

void function2() {
    globalVar = 200;  // Can modify
}

int main() {
    printf("%d", globalVar);  // Can access
}
```

### Static Variables
Retain their value between function calls:

```c
void counter() {
    static int count = 0;  // Initialized only once
    count++;
    printf("%d\n", count);
}

int main() {
    counter();  // Prints: 1
    counter();  // Prints: 2
    counter();  // Prints: 3
}
```

---

## Storage Classes

### auto (default)
Local variables, automatically destroyed:

```c
void func() {
    auto int x = 10;  // 'auto' is default, rarely written
}
```

### static
Preserves value between calls, local scope:

```c
void func() {
    static int count = 0;
    count++;
}
```

### extern
Declares a variable defined elsewhere:

```c
// file1.c
int globalVar = 100;

// file2.c
extern int globalVar;  // References the one in file1.c
```

### register
Suggests storing in CPU register (rarely used):

```c
register int counter;  // Hint for performance
```

---

## Constants

### Using #define

```c
#define PI 3.14159
#define MAX_SIZE 100

int main() {
    float area = PI * radius * radius;
    int array[MAX_SIZE];
}
```

### Using const

```c
const int MAX_USERS = 50;
const float PI = 3.14159;
const char* MESSAGE = "Hello";

// MAX_USERS = 100;  // ❌ Error: cannot modify const
```

---

## Type Casting

### Implicit Casting (Automatic)

```c
int x = 10;
float y = x;  // int → float (automatic)
printf("%f", y);  // 10.000000
```

### Explicit Casting (Manual)

```c
float f = 3.7;
int i = (int)f;  // Explicit cast: 3.7 → 3
printf("%d", i);  // 3

int a = 5, b = 2;
float result = (float)a / b;  // 2.5 (not 2)
```

---

## Sizeof Operator

Get the size of a variable or type in bytes:

```c
printf("int: %zu bytes\n", sizeof(int));           // 4
printf("char: %zu bytes\n", sizeof(char));         // 1
printf("float: %zu bytes\n", sizeof(float));       // 4
printf("double: %zu bytes\n", sizeof(double));     // 8

int x = 10;
printf("x: %zu bytes\n", sizeof(x));               // 4
```

---

## Variable Examples

### Basic Variables

```c
#include <stdio.h>

int main() {
    // Integer variables
    int age = 25;
    int year = 2024;
    
    // Floating point variables
    float price = 19.99;
    double pi = 3.14159265359;
    
    // Character variable
    char grade = 'A';
    
    // Boolean variable
    bool isActive = true;
    
    // Print variables
    printf("Age: %d\n", age);
    printf("Price: %.2f\n", price);
    printf("Grade: %c\n", grade);
    printf("Pi: %f\n", pi);
    
    return 0;
}
```

### Multiple Variables

```c
// Same type, same line
int x = 5, y = 10, z = 15;

// Different types
int age = 30;
float height = 5.9;
char initial = 'J';
```

### Swapping Variables

```c
int a = 5, b = 10, temp;

temp = a;    // temp = 5
a = b;       // a = 10
b = temp;    // b = 5

printf("a = %d, b = %d", a, b);  // a = 10, b = 5
```

---

## Common Mistakes

### Uninitialized Variables

```c
int x;  // ❌ Contains garbage value
printf("%d", x);  // Undefined behavior!

int y = 0;  // ✅ Always initialize
printf("%d", y);  // Safe: prints 0
```

### Type Mismatch

```c
int x = 3.7;  // ⚠️ Warning: 3.7 truncated to 3
printf("%d", x);  // Prints: 3

float f = 10 / 3;  // ❌ Integer division: 3
float g = 10.0 / 3;  // ✅ Float division: 3.333...
```

### Overflow

```c
char c = 127;
c = c + 1;  // ❌ Overflow: wraps to -128

unsigned char uc = 255;
uc = uc + 1;  // ❌ Overflow: wraps to 0
```

---

## Format Specifiers for printf/scanf

```c
int i = 42;
float f = 3.14;
double d = 3.14159;
char c = 'A';
char str[] = "Hello";
unsigned int u = 100;
long l = 1000000;

printf("%d", i);      // Integer
printf("%f", f);      // Float (default 6 decimals)
printf("%.2f", f);    // Float with 2 decimals
printf("%lf", d);     // Double
printf("%c", c);      // Character
printf("%s", str);    // String
printf("%u", u);      // Unsigned int
printf("%ld", l);     // Long
printf("%p", &i);     // Pointer address
printf("%x", i);      // Hexadecimal
printf("%o", i);      // Octal
```

---

## Best Practices

1. **Always initialize variables** before use
2. **Use meaningful names** - `age` not `a`
3. **Declare variables close to first use**
4. **Use const** for values that don't change
5. **Avoid global variables** when possible
6. **Check for overflow** in calculations
7. **Use appropriate types** - don't waste memory
8. **Comment complex variable purposes**

```c
// ✅ Good
const int MAX_STUDENTS = 100;
int studentCount = 0;
float averageGrade = 0.0;

// ❌ Bad
int a;
int x = 1000000000;  // Might overflow
float f;  // What does this represent?
```

---

## Summary

| Concept | Example | Description |
|---------|---------|-------------|
| **Declaration** | `int x;` | Create variable |
| **Initialization** | `int x = 10;` | Declare + assign |
| **Assignment** | `x = 20;` | Change value |
| **Local** | Inside function | Limited scope |
| **Global** | Outside functions | Accessible everywhere |
| **Constant** | `const int X = 10;` | Cannot change |
| **Static** | `static int x;` | Retains value |
| **Type Cast** | `(int)3.7` | Convert type |

Variables are fundamental to C programming - master them and you'll have a solid foundation!