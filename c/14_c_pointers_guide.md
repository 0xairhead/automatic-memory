# Pointers in C

## Table of Contents

- [What are Pointers?](#what-are-pointers)
- [Why Use Pointers?](#why-use-pointers)
- [Memory and Addresses](#memory-and-addresses)
- [Basic Pointer Syntax](#basic-pointer-syntax)
- [Declaring Pointers](#declaring-pointers)
- [Pointer Operations](#pointer-operations)
- [Null Pointers](#null-pointers)
- [Pointer Arithmetic](#pointer-arithmetic)
- [Pointers and Arrays](#pointers-and-arrays)
- [Pointers and Strings](#pointers-and-strings)
- [Pointers and Functions](#pointers-and-functions)
- [Pointer to Pointer](#pointer-to-pointer)
- [Dynamic Memory Allocation](#dynamic-memory-allocation)
- [Function Pointers](#function-pointers)
- [Const Pointers](#const-pointers)
- [Void Pointers](#void-pointers)
- [Common Pointer Mistakes](#common-pointer-mistakes)
- [Practical Examples](#practical-examples)
- [Common Pitfalls](#common-pitfalls)
- [Best Practices](#best-practices)
- [Quick Reference](#quick-reference)

---

## What are Pointers?

**Pointers** are variables that store memory addresses of other variables. Instead of holding a value directly, a pointer "points to" the location in memory where the value is stored.

```
Regular Variable:
int num = 42;
[Memory: 0x1000] → 42

Pointer Variable:
int *ptr = &num;
[Memory: 0x2000] → 0x1000 (address of num)
```

**Analogy:** A pointer is like a street address that tells you where a house is located, rather than being the house itself.

---

## Why Use Pointers?

1. **Dynamic Memory Allocation** - Allocate memory at runtime
2. **Efficient Array/String Handling** - Pass large data without copying
3. **Function Modifications** - Allow functions to modify variables
4. **Data Structures** - Build linked lists, trees, graphs
5. **Low-level Programming** - Direct memory access and hardware control
6. **Performance** - Avoid copying large structures

---

## Memory and Addresses

Every variable in C has a memory address where it's stored.

### Understanding Memory Addresses

```c
#include <stdio.h>

int main() {
    int num = 42;
    float price = 99.99;
    char grade = 'A';
    
    printf("Value of num: %d\n", num);
    printf("Address of num: %p\n", (void*)&num);
    printf("\n");
    
    printf("Value of price: %.2f\n", price);
    printf("Address of price: %p\n", (void*)&price);
    printf("\n");
    
    printf("Value of grade: %c\n", grade);
    printf("Address of grade: %p\n", (void*)&grade);
    
    return 0;
}
```

**Output (addresses will vary):**
```
Value of num: 42
Address of num: 0x7ffd5c9b8a0c

Value of price: 99.99
Address of price: 0x7ffd5c9b8a08

Value of grade: A
Address of grade: 0x7ffd5c9b8a07
```

### Memory Layout

```
High Memory Addresses
    ↑
    |  Stack (local variables, function calls)
    |
    |  Heap (dynamic memory allocation)
    |
    |  BSS (uninitialized global/static variables)
    |
    |  Data (initialized global/static variables)
    |
    ↓  Text (program code)
Low Memory Addresses
```

---

## Basic Pointer Syntax

### The Two Key Operators

**1. Address-of Operator (`&`)** - Gets the address of a variable

```c
int num = 10;
&num  // Address of num
```

**2. Dereference Operator (`*`)** - Accesses the value at an address

```c
int *ptr = &num;
*ptr  // Value at the address stored in ptr
```

### Basic Example

```c
#include <stdio.h>

int main() {
    int num = 42;
    int *ptr;
    
    ptr = &num;  // ptr now holds the address of num
    
    printf("Value of num: %d\n", num);
    printf("Address of num: %p\n", (void*)&num);
    printf("Value of ptr: %p\n", (void*)ptr);
    printf("Value pointed to by ptr: %d\n", *ptr);
    
    return 0;
}
```

**Output:**
```
Value of num: 42
Address of num: 0x7ffd5c9b8a0c
Value of ptr: 0x7ffd5c9b8a0c
Value pointed to by ptr: 42
```

### Visual Representation

```
Memory Layout:

Address     Variable    Value
0x1000      num         42
0x2000      ptr         0x1000

ptr "points to" num:
ptr → num
```

---

## Declaring Pointers

### Basic Declaration Syntax

```c
datatype *pointer_name;
```

### Examples

```c
int *ptr1;       // Pointer to int
float *ptr2;     // Pointer to float
char *ptr3;      // Pointer to char
double *ptr4;    // Pointer to double
```

### Multiple Declarations

```c
// ⚠️ Careful with syntax!
int *p1, *p2, *p3;  // Three pointers to int

// ❌ WRONG: Only p1 is a pointer!
int *p1, p2, p3;    // p1 is pointer, p2 and p3 are ints
```

### Declaration and Initialization

```c
#include <stdio.h>

int main() {
    int num = 100;
    
    // Declare and initialize
    int *ptr = &num;
    
    // Or separate
    int *ptr2;
    ptr2 = &num;
    
    printf("*ptr = %d\n", *ptr);
    printf("*ptr2 = %d\n", *ptr2);
    
    return 0;
}
```

---

## Pointer Operations

### 1. Assignment

```c
int num = 42;
int *ptr = &num;  // Assign address
```

### 2. Dereferencing

```c
int value = *ptr;  // Get value at address
*ptr = 100;        // Modify value at address
```

### 3. Complete Example

```c
#include <stdio.h>

int main() {
    int x = 10;
    int *ptr = &x;
    
    printf("Initial x: %d\n", x);
    printf("*ptr: %d\n", *ptr);
    
    // Modify through pointer
    *ptr = 20;
    
    printf("After *ptr = 20:\n");
    printf("x: %d\n", x);
    printf("*ptr: %d\n", *ptr);
    
    // Modify directly
    x = 30;
    
    printf("After x = 30:\n");
    printf("x: %d\n", x);
    printf("*ptr: %d\n", *ptr);
    
    return 0;
}
```

**Output:**
```
Initial x: 10
*ptr: 10
After *ptr = 20:
x: 20
*ptr: 20
After x = 30:
x: 30
*ptr: 30
```

### Swapping Values with Pointers

```c
#include <stdio.h>

void swap(int *a, int *b) {
    int temp = *a;
    *a = *b;
    *b = temp;
}

int main() {
    int x = 5, y = 10;
    
    printf("Before swap: x = %d, y = %d\n", x, y);
    swap(&x, &y);
    printf("After swap: x = %d, y = %d\n", x, y);
    
    return 0;
}
```

**Output:**
```
Before swap: x = 5, y = 10
After swap: x = 10, y = 5
```

---

## Null Pointers

A **null pointer** points to nothing. Always initialize pointers to NULL if not assigning an address immediately.

### NULL Pointer Basics

```c
#include <stdio.h>

int main() {
    int *ptr = NULL;  // Good practice: initialize to NULL
    
    if (ptr == NULL) {
        printf("Pointer is NULL\n");
    }
    
    // ❌ DANGEROUS: Don't dereference NULL pointer!
    // printf("%d\n", *ptr);  // CRASH!
    
    // ✅ SAFE: Check before dereferencing
    if (ptr != NULL) {
        printf("%d\n", *ptr);
    }
    
    return 0;
}
```

### Safe Pointer Usage

```c
#include <stdio.h>

void printValue(int *ptr) {
    if (ptr == NULL) {
        printf("NULL pointer received\n");
        return;
    }
    printf("Value: %d\n", *ptr);
}

int main() {
    int num = 42;
    int *p1 = &num;
    int *p2 = NULL;
    
    printValue(p1);  // OK
    printValue(p2);  // Handled safely
    
    return 0;
}
```

**Output:**
```
Value: 42
NULL pointer received
```

---

## Pointer Arithmetic

Pointers can be incremented, decremented, and compared.

### Basic Arithmetic

```c
#include <stdio.h>

int main() {
    int arr[5] = {10, 20, 30, 40, 50};
    int *ptr = arr;
    
    printf("ptr points to: %d\n", *ptr);
    
    ptr++;  // Move to next element
    printf("After ptr++: %d\n", *ptr);
    
    ptr += 2;  // Move 2 elements forward
    printf("After ptr += 2: %d\n", *ptr);
    
    ptr--;  // Move back one element
    printf("After ptr--: %d\n", *ptr);
    
    return 0;
}
```

**Output:**
```
ptr points to: 10
After ptr++: 20
After ptr += 2: 40
After ptr--: 30
```

### How Pointer Arithmetic Works

```c
#include <stdio.h>

int main() {
    int arr[3] = {10, 20, 30};
    int *ptr = arr;
    
    printf("Address of arr[0]: %p\n", (void*)&arr[0]);
    printf("Address of arr[1]: %p\n", (void*)&arr[1]);
    printf("Address of arr[2]: %p\n", (void*)&arr[2]);
    
    printf("\nptr: %p\n", (void*)ptr);
    printf("ptr + 1: %p\n", (void*)(ptr + 1));
    printf("ptr + 2: %p\n", (void*)(ptr + 2));
    
    printf("\nDifference: %ld bytes\n", (char*)(ptr + 1) - (char*)ptr);
    
    return 0;
}
```

**Output:**
```
Address of arr[0]: 0x7ffd5c9b8a00
Address of arr[1]: 0x7ffd5c9b8a04
Address of arr[2]: 0x7ffd5c9b8a08

ptr: 0x7ffd5c9b8a00
ptr + 1: 0x7ffd5c9b8a04
ptr + 2: 0x7ffd5c9b8a08

Difference: 4 bytes
```

**Key Point:** `ptr + 1` moves by `sizeof(type)` bytes, not just 1 byte!

### Subtracting Pointers

```c
#include <stdio.h>

int main() {
    int arr[5] = {10, 20, 30, 40, 50};
    int *ptr1 = &arr[1];
    int *ptr2 = &arr[4];
    
    printf("Distance between ptr2 and ptr1: %ld elements\n", ptr2 - ptr1);
    
    return 0;
}
```

**Output:**
```
Distance between ptr2 and ptr1: 3 elements
```

---

## Pointers and Arrays

Arrays and pointers are closely related in C. An array name is essentially a pointer to its first element.

### Array Name as Pointer

```c
#include <stdio.h>

int main() {
    int arr[5] = {10, 20, 30, 40, 50};
    
    printf("arr: %p\n", (void*)arr);
    printf("&arr[0]: %p\n", (void*)&arr[0]);
    
    printf("\n*arr: %d\n", *arr);
    printf("arr[0]: %d\n", arr[0]);
    
    printf("\n*(arr + 1): %d\n", *(arr + 1));
    printf("arr[1]: %d\n", arr[1]);
    
    return 0;
}
```

**Output:**
```
arr: 0x7ffd5c9b8a00
&arr[0]: 0x7ffd5c9b8a00

*arr: 10
arr[0]: 10

*(arr + 1): 20
arr[1]: 20
```

### Array Access Methods

```c
#include <stdio.h>

int main() {
    int arr[5] = {10, 20, 30, 40, 50};
    int *ptr = arr;
    
    // Method 1: Array notation
    printf("arr[2] = %d\n", arr[2]);
    
    // Method 2: Pointer notation with array name
    printf("*(arr + 2) = %d\n", *(arr + 2));
    
    // Method 3: Pointer notation with pointer variable
    printf("*(ptr + 2) = %d\n", *(ptr + 2));
    
    // Method 4: Array notation with pointer
    printf("ptr[2] = %d\n", ptr[2]);
    
    return 0;
}
```

**Output:**
```
arr[2] = 30
*(arr + 2) = 30
*(ptr + 2) = 30
ptr[2] = 30
```

### Iterating Array with Pointer

```c
#include <stdio.h>

int main() {
    int arr[5] = {10, 20, 30, 40, 50};
    int *ptr;
    
    printf("Using pointer to iterate:\n");
    for (ptr = arr; ptr < arr + 5; ptr++) {
        printf("%d ", *ptr);
    }
    printf("\n");
    
    return 0;
}
```

**Output:**
```
Using pointer to iterate:
10 20 30 40 50
```

### Key Differences: Array vs Pointer

```c
int arr[5];
int *ptr;

// ✅ Valid
ptr = arr;        // OK: pointer can point to array
ptr++;            // OK: can increment pointer

// ❌ Invalid
arr = ptr;        // ERROR: can't reassign array name
arr++;            // ERROR: can't increment array name
```

---

## Pointers and Strings

Strings in C are arrays of characters, so pointers work naturally with them.

### String Initialization

```c
#include <stdio.h>

int main() {
    // Method 1: Array
    char str1[] = "Hello";
    
    // Method 2: Pointer (points to string literal)
    char *str2 = "World";
    
    printf("str1: %s\n", str1);
    printf("str2: %s\n", str2);
    
    // Modify array (OK)
    str1[0] = 'h';
    printf("str1: %s\n", str1);
    
    // ❌ Modify string literal (UNDEFINED BEHAVIOR)
    // str2[0] = 'w';  // MAY CRASH!
    
    return 0;
}
```

### Iterating String with Pointer

```c
#include <stdio.h>

int main() {
    char str[] = "Hello";
    char *ptr = str;
    
    printf("Characters: ");
    while (*ptr != '\0') {
        printf("%c ", *ptr);
        ptr++;
    }
    printf("\n");
    
    return 0;
}
```

**Output:**
```
Characters: H e l l o
```

### String Functions with Pointers

```c
#include <stdio.h>

int stringLength(char *str) {
    int len = 0;
    while (*str != '\0') {
        len++;
        str++;
    }
    return len;
}

void stringCopy(char *dest, char *src) {
    while (*src != '\0') {
        *dest = *src;
        dest++;
        src++;
    }
    *dest = '\0';  // Don't forget null terminator!
}

int main() {
    char str1[] = "Hello";
    char str2[20];
    
    printf("Length: %d\n", stringLength(str1));
    
    stringCopy(str2, str1);
    printf("Copied: %s\n", str2);
    
    return 0;
}
```

**Output:**
```
Length: 5
Copied: Hello
```

---

## Pointers and Functions

Pointers enable functions to modify variables and work with large data efficiently.

### Pass by Value vs Pass by Reference

```c
#include <stdio.h>

// Pass by value - doesn't modify original
void increment1(int x) {
    x++;
}

// Pass by reference - modifies original
void increment2(int *x) {
    (*x)++;
}

int main() {
    int num = 10;
    
    increment1(num);
    printf("After increment1: %d\n", num);  // Still 10
    
    increment2(&num);
    printf("After increment2: %d\n", num);  // Now 11
    
    return 0;
}
```

**Output:**
```
After increment1: 10
After increment2: 11
```

### Returning Multiple Values

```c
#include <stdio.h>

void divide(int a, int b, int *quotient, int *remainder) {
    *quotient = a / b;
    *remainder = a % b;
}

int main() {
    int q, r;
    
    divide(17, 5, &q, &r);
    
    printf("17 / 5 = %d remainder %d\n", q, r);
    
    return 0;
}
```

**Output:**
```
17 / 5 = 3 remainder 2
```

### Passing Arrays to Functions

```c
#include <stdio.h>

int sumArray(int *arr, int size) {
    int sum = 0;
    for (int i = 0; i < size; i++) {
        sum += arr[i];
    }
    return sum;
}

void modifyArray(int *arr, int size) {
    for (int i = 0; i < size; i++) {
        arr[i] *= 2;
    }
}

int main() {
    int numbers[] = {1, 2, 3, 4, 5};
    int size = 5;
    
    printf("Sum: %d\n", sumArray(numbers, size));
    
    modifyArray(numbers, size);
    
    printf("After modification: ");
    for (int i = 0; i < size; i++) {
        printf("%d ", numbers[i]);
    }
    printf("\n");
    
    return 0;
}
```

**Output:**
```
Sum: 15
After modification: 2 4 6 8 10
```

---

## Pointer to Pointer

A pointer that stores the address of another pointer.

### Basic Example

```c
#include <stdio.h>

int main() {
    int num = 42;
    int *ptr = &num;
    int **pptr = &ptr;
    
    printf("Value of num: %d\n", num);
    printf("Value using *ptr: %d\n", *ptr);
    printf("Value using **pptr: %d\n", **pptr);
    
    printf("\nAddress of num: %p\n", (void*)&num);
    printf("Value of ptr: %p\n", (void*)ptr);
    printf("Value of pptr: %p\n", (void*)pptr);
    
    return 0;
}
```

**Output:**
```
Value of num: 42
Value using *ptr: 42
Value using **pptr: 42

Address of num: 0x7ffd5c9b8a0c
Value of ptr: 0x7ffd5c9b8a0c
Value of pptr: 0x7ffd5c9b8a10
```

### Visual Representation

```
Memory Layout:

Address     Variable    Value
0x1000      num         42
0x2000      ptr         0x1000  (points to num)
0x3000      pptr        0x2000  (points to ptr)

pptr → ptr → num
```

### Modifying Pointer Through Double Pointer

```c
#include <stdio.h>

void changePointer(int **pptr, int *newTarget) {
    *pptr = newTarget;
}

int main() {
    int x = 10, y = 20;
    int *ptr = &x;
    
    printf("*ptr = %d\n", *ptr);
    
    changePointer(&ptr, &y);
    
    printf("*ptr = %d\n", *ptr);
    
    return 0;
}
```

**Output:**
```
*ptr = 10
*ptr = 20
```

### 2D Array with Pointer to Pointer

```c
#include <stdio.h>

int main() {
    int rows = 3, cols = 3;
    int data[3][3] = {
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9}
    };
    
    int *ptrs[3];
    for (int i = 0; i < 3; i++) {
        ptrs[i] = data[i];
    }
    
    int **pptr = ptrs;
    
    printf("Accessing 2D array through pointer to pointer:\n");
    for (int i = 0; i < rows; i++) {
        for (int j = 0; j < cols; j++) {
            printf("%d ", pptr[i][j]);
        }
        printf("\n");
    }
    
    return 0;
}
```

**Output:**
```
Accessing 2D array through pointer to pointer:
1 2 3
4 5 6
7 8 9
```

---

## Dynamic Memory Allocation

Allocate memory at runtime using `malloc`, `calloc`, `realloc`, and `free`.

### malloc - Memory Allocation

```c
#include <stdio.h>
#include <stdlib.h>

int main() {
    int *ptr;
    int n = 5;
    
    // Allocate memory for 5 integers
    ptr = (int*)malloc(n * sizeof(int));
    
    if (ptr == NULL) {
        printf("Memory allocation failed\n");
        return 1;
    }
    
    // Use the allocated memory
    for (int i = 0; i < n; i++) {
        ptr[i] = (i + 1) * 10;
    }
    
    printf("Values: ");
    for (int i = 0; i < n; i++) {
        printf("%d ", ptr[i]);
    }
    printf("\n");
    
    // Free the memory
    free(ptr);
    ptr = NULL;  // Good practice
    
    return 0;
}
```

**Output:**
```
Values: 10 20 30 40 50
```

### calloc - Contiguous Allocation (Initialized to Zero)

```c
#include <stdio.h>
#include <stdlib.h>

int main() {
    int *ptr;
    int n = 5;
    
    // Allocate and initialize to 0
    ptr = (int*)calloc(n, sizeof(int));
    
    if (ptr == NULL) {
        printf("Memory allocation failed\n");
        return 1;
    }
    
    printf("Values (initialized to 0): ");
    for (int i = 0; i < n; i++) {
        printf("%d ", ptr[i]);
    }
    printf("\n");
    
    free(ptr);
    
    return 0;
}
```

**Output:**
```
Values (initialized to 0): 0 0 0 0 0
```

### realloc - Resize Allocated Memory

```c
#include <stdio.h>
#include <stdlib.h>

int main() {
    int *ptr;
    int n = 3;
    
    ptr = (int*)malloc(n * sizeof(int));
    
    for (int i = 0; i < n; i++) {
        ptr[i] = i + 1;
    }
    
    printf("Original: ");
    for (int i = 0; i < n; i++) {
        printf("%d ", ptr[i]);
    }
    printf("\n");
    
    // Resize to 5 elements
    n = 5;
    ptr = (int*)realloc(ptr, n * sizeof(int));
    
    if (ptr == NULL) {
        printf("Reallocation failed\n");
        return 1;
    }
    
    // Initialize new elements
    for (int i = 3; i < n; i++) {
        ptr[i] = i + 1;
    }
    
    printf("After realloc: ");
    for (int i = 0; i < n; i++) {
        printf("%d ", ptr[i]);
    }
    printf("\n");
    
    free(ptr);
    
    return 0;
}
```

**Output:**
```
Original: 1 2 3
After realloc: 1 2 3 4 5
```

### Complete Dynamic Array Example

```c
#include <stdio.h>
#include <stdlib.h>

int main() {
    int *arr;
    int size;
    
    printf("Enter array size: ");
    scanf("%d", &size);
    
    arr = (int*)malloc(size * sizeof(int));
    
    if (arr == NULL) {
        printf("Memory allocation failed\n");
        return 1;
    }
    
    printf("Enter %d elements:\n", size);
    for (int i = 0; i < size; i++) {
        scanf("%d", &arr[i]);
    }
    
    printf("You entered: ");
    for (int i = 0; i < size; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");
    
    // Calculate sum
    int sum = 0;
    for (int i = 0; i < size; i++) {
        sum += arr[i];
    }
    printf("Sum: %d\n", sum);
    
    free(arr);
    
    return 0;
}
```

---

## Function Pointers

Pointers can also point to functions, enabling callbacks and dynamic function calls.

### Basic Function Pointer

```c
#include <stdio.h>

int add(int a, int b) {
    return a + b;
}

int multiply(int a, int b) {
    return a * b;
}

int main() {
    // Declare function pointer
    int (*operation)(int, int);
    
    // Point to add function
    operation = add;
    printf("5 + 3 = %d\n", operation(5, 3));
    
    // Point to multiply function
    operation = multiply;
    printf("5 * 3 = %d\n", operation(5, 3));
    
    return 0;
}
```

**Output:**
```
5 + 3 = 8
5 * 3 = 15
```

### Calculator with Function Pointers

```c
#include <stdio.h>

int add(int a, int b) { return a + b; }
int subtract(int a, int b) { return a - b; }
int multiply(int a, int b) { return a * b; }
int divide(int a, int b) { return b != 0 ? a / b : 0; }

int calculate(int a, int b, int (*operation)(int, int)) {
    return operation(a, b);
}

int main() {
    int x = 10, y = 5;
    
    printf("%d + %d = %d\n", x, y, calculate(x, y, add));
    printf("%d - %d = %d\n", x, y, calculate(x, y, subtract));
    printf("%d * %d = %d\n", x, y, calculate(x, y, multiply));
    printf("%d / %d = %d\n", x, y, calculate(x, y, divide));
    
    return 0;
}
```

**Output:**
```
10 + 5 = 15
10 - 5 = 5
10 * 5 = 50
10 / 5 = 2
```

### Array of Function Pointers

```c
#include <stdio.h>

void morning() { printf("Good morning!\n"); }
void afternoon() { printf("Good afternoon!\n"); }
void evening() { printf("Good evening!\n"); }
void night() { printf("Good night!\n"); }

int main() {
    void (*greetings[4])() = {morning, afternoon, evening, night};
    
    int hour;
    printf("Enter hour (0-23): ");
    scanf("%d", &hour);
    
    if (hour >= 6 && hour < 12) {
        greetings[0]();
    } else if (hour >= 12 && hour < 18) {
        greetings[1]();
    } else if (hour >= 18 && hour < 22) {
        greetings[2]();
    } else {
        greetings[3]();
    }
    
    return 0;
}
```

---

## Const Pointers

There are several ways to use `const` with pointers.

### Types of Const Pointers

```c
#include <stdio.h>

int main() {
    int x = 10, y = 20;
    
    // 1. Pointer to constant (can't change value through pointer)
    const int *ptr1 = &x;
    // *ptr1 = 15;  // ERROR: can't modify value
    ptr1 = &y;      // OK: can change pointer
    
    // 2. Constant pointer (can't change where it points)
    int *const ptr2 = &x;
    *ptr2 = 15;     // OK: can modify value
    // ptr2 = &y;   // ERROR: can't change pointer
    
    // 3. Constant pointer to constant (can't change either)
    const int *const ptr3 = &x;
    // *ptr3 = 15;  // ERROR: can't modify value
    // ptr3 = &y;   // ERROR: can't change pointer
    
    printf("ptr1 points to: %d\n", *ptr1);
    printf("ptr2 points to: %d\n", *ptr2);
    printf("ptr3 points to: %d\n", *ptr3);
    
    return 0;
}
```

### Practical Use Case

```c
#include <stdio.h>
#include <string.h>

// Function can't modify the string
void printString(const char *str) {
    // str[0] = 'X';  // ERROR: can't modify
    printf("%s\n", str);
}

// Function can't modify the array
int sumArray(const int *arr, int size) {
    int sum = 0;
    for (int i = 0; i < size; i++) {
        sum += arr[i];
        // arr[i] = 0;  // ERROR: can't modify
    }
    return sum;
}

int main() {
    char message[] = "Hello";
    printString(message);
    
    int numbers[] = {1, 2, 3, 4, 5};
    printf("Sum: %d\n", sumArray(numbers, 5));
    
    return 0;
}
```

---

## Void Pointers

A **void pointer** (`void*`) is a generic pointer that can point to any data type.

### Basic Void Pointer

```c
#include <stdio.h>

int main() {
    int x = 10;
    float y = 3.14;
    char c = 'A';
    
    void *ptr;
    
    ptr = &x;
    printf("Integer: %d\n", *(int*)ptr);
    
    ptr = &y;
    printf("Float: %.2f\n", *(float*)ptr);
    
    ptr = &c;
    printf("Char: %c\n", *(char*)ptr);
    
    return 0;
}
```

**Output:**
```
Integer: 10
Float: 3.14
Char: A
```

### Generic Swap Function

```c
#include <stdio.h>
#include <string.h>

void swap(void *a, void *b, size_t size) {
    char temp[size];
    memcpy(temp, a, size);
    memcpy(a, b, size);
    memcpy(b, temp, size);
}

int main() {
    int x = 5, y = 10;
    printf("Before: x=%d, y=%d\n", x, y);
    swap(&x, &y, sizeof(int));
    printf("After: x=%d, y=%d\n", x, y);
    
    float a = 1.5, b = 2.5;
    printf("\nBefore: a=%.1f, b=%.1f\n", a, b);
    swap(&a, &b, sizeof(float));
    printf("After: a=%.1f, b=%.1f\n", a, b);
    
    return 0;
}
```

**Output:**
```
Before: x=5, y=10
After: x=10, y=5

Before: a=1.5, b=2.5
After: a=2.5, b=1.5
```

---

## Common Pointer Mistakes

### 1. Uninitialized Pointer

```c
// ❌ BAD: Wild pointer
int *ptr;
*ptr = 10;  // CRASH! Points to random memory

// ✅ GOOD: Initialize before use
int *ptr = NULL;
int x = 10;
ptr = &x;
*ptr = 20;
```

### 2. Dangling Pointer

```c
// ❌ BAD: Pointer to freed memory
int *ptr = (int*)malloc(sizeof(int));
*ptr = 10;
free(ptr);
printf("%d\n", *ptr);  // UNDEFINED BEHAVIOR!

// ✅ GOOD: Set to NULL after freeing
int *ptr = (int*)malloc(sizeof(int));
*ptr = 10;
free(ptr);
ptr = NULL;
if (ptr != NULL) {
    printf("%d\n", *ptr);
}
```

### 3. Memory Leak

```c
// ❌ BAD: Forgot to free
void createArray() {
    int *arr = (int*)malloc(100 * sizeof(int));
    // ... use arr ...
    // Forgot to free(arr)!
}

// ✅ GOOD: Always free
void createArray() {
    int *arr = (int*)malloc(100 * sizeof(int));
    if (arr == NULL) return;
    // ... use arr ...
    free(arr);
}
```

### 4. Returning Pointer to Local Variable

```c
// ❌ BAD: Returns address of local variable
int* getNumber() {
    int x = 10;
    return &x;  // DANGER! x is destroyed after return
}

// ✅ GOOD: Use static or dynamic allocation
int* getNumber() {
    static int x = 10;
    return &x;
}

// ✅ BETTER: Dynamic allocation
int* getNumber() {
    int *ptr = (int*)malloc(sizeof(int));
    *ptr = 10;
    return ptr;  // Caller must free!
}
```

### 5. Buffer Overflow

```c
// ❌ BAD: Writing beyond allocated memory
int *arr = (int*)malloc(5 * sizeof(int));
for (int i = 0; i < 10; i++) {
    arr[i] = i;  // Writing beyond bounds!
}

// ✅ GOOD: Respect bounds
int *arr = (int*)malloc(5 * sizeof(int));
for (int i = 0; i < 5; i++) {
    arr[i] = i;
}
```

---

## Practical Examples

### Example 1: Dynamic String

```c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* createString(const char *str) {
    char *newStr = (char*)malloc(strlen(str) + 1);
    if (newStr != NULL) {
        strcpy(newStr, str);
    }
    return newStr;
}

int main() {
    char *message = createString("Hello, World!");
    
    if (message != NULL) {
        printf("Message: %s\n", message);
        free(message);
    }
    
    return 0;
}
```

### Example 2: Linked List Node

```c
#include <stdio.h>
#include <stdlib.h>

typedef struct Node {
    int data;
    struct Node *next;
} Node;

Node* createNode(int data) {
    Node *newNode = (Node*)malloc(sizeof(Node));
    if (newNode != NULL) {
        newNode->data = data;
        newNode->next = NULL;
    }
    return newNode;
}

void printList(Node *head) {
    Node *current = head;
    while (current != NULL) {
        printf("%d -> ", current->data);
        current = current->next;
    }
    printf("NULL\n");
}

void freeList(Node *head) {
    Node *current = head;
    while (current != NULL) {
        Node *temp = current;
        current = current->next;
        free(temp);
    }
}

int main() {
    Node *head = createNode(1);
    head->next = createNode(2);
    head->next->next = createNode(3);
    
    printList(head);
    
    freeList(head);
    
    return 0;
}
```

**Output:**
```
1 -> 2 -> 3 -> NULL
```

### Example 3: Matrix Operations

```c
#include <stdio.h>
#include <stdlib.h>

int** createMatrix(int rows, int cols) {
    int **matrix = (int**)malloc(rows * sizeof(int*));
    for (int i = 0; i < rows; i++) {
        matrix[i] = (int*)malloc(cols * sizeof(int));
    }
    return matrix;
}

void fillMatrix(int **matrix, int rows, int cols) {
    int value = 1;
    for (int i = 0; i < rows; i++) {
        for (int j = 0; j < cols; j++) {
            matrix[i][j] = value++;
        }
    }
}

void printMatrix(int **matrix, int rows, int cols) {
    for (int i = 0; i < rows; i++) {
        for (int j = 0; j < cols; j++) {
            printf("%3d ", matrix[i][j]);
        }
        printf("\n");
    }
}

void freeMatrix(int **matrix, int rows) {
    for (int i = 0; i < rows; i++) {
        free(matrix[i]);
    }
    free(matrix);
}

int main() {
    int rows = 3, cols = 4;
    int **matrix = createMatrix(rows, cols);
    
    fillMatrix(matrix, rows, cols);
    printMatrix(matrix, rows, cols);
    
    freeMatrix(matrix, rows);
    
    return 0;
}
```

**Output:**
```
  1   2   3   4 
  5   6   7   8 
  9  10  11  12
```

---

## Common Pitfalls

### 1. Pointer Type Mismatch

```c
// ❌ BAD
float x = 3.14;
int *ptr = (int*)&x;  // Wrong type!
printf("%d\n", *ptr);  // Garbage value

// ✅ GOOD
float x = 3.14;
float *ptr = &x;
printf("%.2f\n", *ptr);
```

### 2. Modifying String Literals

```c
// ❌ BAD: String literal in read-only memory
char *str = "Hello";
str[0] = 'h';  // MAY CRASH!

// ✅ GOOD: Use array for modifiable strings
char str[] = "Hello";
str[0] = 'h';  // OK
```

### 3. Pointer Arithmetic on Non-Arrays

```c
// ❌ BAD
int x = 10;
int *ptr = &x;
ptr++;  // Where does it point now? Undefined!

// ✅ GOOD: Only use pointer arithmetic with arrays
int arr[5] = {1, 2, 3, 4, 5};
int *ptr = arr;
ptr++;  // Now points to arr[1]
```

### 4. Double Free

```c
// ❌ BAD
int *ptr = (int*)malloc(sizeof(int));
free(ptr);
free(ptr);  // CRASH! Double free

// ✅ GOOD
int *ptr = (int*)malloc(sizeof(int));
free(ptr);
ptr = NULL;  // Prevent double free
```

### 5. Using Freed Memory

```c
// ❌ BAD
int *ptr = (int*)malloc(sizeof(int));
*ptr = 42;
free(ptr);
printf("%d\n", *ptr);  // Use after free!

// ✅ GOOD
int *ptr = (int*)malloc(sizeof(int));
*ptr = 42;
printf("%d\n", *ptr);
free(ptr);
ptr = NULL;
```

---

## Best Practices

1. **Always initialize pointers** - Set to NULL if no immediate target
2. **Check for NULL** before dereferencing
3. **Free allocated memory** when done
4. **Set pointers to NULL** after freeing
5. **Use const** when pointer shouldn't modify data
6. **Avoid pointer arithmetic** except with arrays
7. **Don't return addresses** of local variables
8. **Match malloc with free** - no memory leaks
9. **Check malloc return value** - it can fail
10. **Use sizeof** for portable allocation

```c
// ✅ GOOD PRACTICES EXAMPLE
#include <stdio.h>
#include <stdlib.h>

int* createArray(int size) {
    // Check parameter
    if (size <= 0) return NULL;
    
    // Allocate with sizeof
    int *arr = (int*)malloc(size * sizeof(int));
    
    // Check allocation
    if (arr == NULL) {
        fprintf(stderr, "Memory allocation failed\n");
        return NULL;
    }
    
    // Initialize
    for (int i = 0; i < size; i++) {
        arr[i] = 0;
    }
    
    return arr;
}

int main() {
    int *arr = createArray(5);
    
    // Check before use
    if (arr != NULL) {
        // Use array
        for (int i = 0; i < 5; i++) {
            arr[i] = i + 1;
        }
        
        // Free when done
        free(arr);
        arr = NULL;  // Prevent dangling pointer
    }
    
    return 0;
}
```

---

## Quick Reference

```c
// Declaration
int *ptr;              // Pointer to int
int **pptr;            // Pointer to pointer to int
int *arr[10];          // Array of 10 int pointers
int (*ptr)[10];        // Pointer to array of 10 ints
int (*func)(int, int); // Pointer to function

// Operators
&var    // Address-of operator
*ptr    // Dereference operator

// Assignment
ptr = &var;     // ptr points to var
*ptr = 10;      // Set value at address

// Null pointer
ptr = NULL;
if (ptr == NULL) { }

// Pointer arithmetic
ptr++;          // Next element
ptr--;          // Previous element
ptr + n         // n elements forward
ptr - n         // n elements backward
ptr2 - ptr1     // Distance between pointers

// Arrays and pointers
arr[i]          // Same as *(arr + i)
ptr[i]          // Same as *(ptr + i)

// Dynamic allocation
ptr = (type*)malloc(size);        // Allocate
ptr = (type*)calloc(n, size);     // Allocate and zero
ptr = (type*)realloc(ptr, size);  // Resize
free(ptr);                        // Free

// Const pointers
const int *ptr;        // Pointer to constant
int *const ptr;        // Constant pointer
const int *const ptr;  // Constant pointer to constant

// Void pointer
void *ptr;             // Generic pointer
(type*)ptr             // Must cast to use

// Common patterns
if (ptr != NULL) { *ptr = value; }
for (p = arr; p < arr + n; p++) { }
```

Pointers are powerful but require careful handling - master them to unlock C's full potential!
