# Functions in C

## Table of Contents

1. [What is a Function?](#what-is-a-function)
2. [Function Components](#function-components)
3. [Types of Functions](#types-of-functions)
4. [Function Declaration vs Definition](#function-declaration-vs-definition)
5. [Function Parameters](#function-parameters)
6. [Return Statement](#return-statement)
7. [Common Function Examples](#common-function-examples)
8. [Recursion](#recursion)
9. [Arrays as Function Parameters](#arrays-as-function-parameters)
10. [Function with Multiple Return Values (Using Pointers)](#function-with-multiple-return-values-using-pointers)
11. [Variable Scope in Functions](#variable-scope-in-functions)
12. [Function Pointers](#function-pointers)
13. [Inline Functions (C99)](#inline-functions-c99)
14. [Common Pitfalls](#common-pitfalls)
15. [Practical Function Library Example](#practical-function-library-example)
16. [Best Practices](#best-practices)
17. [Function Summary](#function-summary)

---

## What is a Function?

A **function** is a reusable block of code that performs a specific task. Functions help organize code, improve readability, and reduce repetition.

```c
// Instead of repeating code:
printf("Hello, Alice\n");
printf("Hello, Bob\n");
printf("Hello, Charlie\n");

// Use a function:
void greet(char name[]) {
    printf("Hello, %s\n", name);
}

greet("Alice");
greet("Bob");
greet("Charlie");
```

---

## Function Components

```c
return_type function_name(parameter_list) {
    // Function body
    return value;  // Optional, depends on return_type
}
```

**Components:**
1. **Return Type** - Data type of the value returned (int, float, void, etc.)
2. **Function Name** - Identifier for the function
3. **Parameters** - Input values (optional)
4. **Function Body** - Code to execute
5. **Return Statement** - Value to return (if not void)

---

## Types of Functions

### 1. Functions Without Parameters and Without Return Value

```c
#include <stdio.h>

void greet() {
    printf("Hello, World!\n");
}

int main() {
    greet();  // Call the function
    greet();  // Can call multiple times
    
    return 0;
}
```

### 2. Functions Without Parameters but With Return Value

```c
#include <stdio.h>

int getRandomNumber() {
    return 42;  // Always returns 42
}

int main() {
    int num = getRandomNumber();
    printf("Number: %d\n", num);
    
    return 0;
}
```

### 3. Functions With Parameters but Without Return Value

```c
#include <stdio.h>

void printSum(int a, int b) {
    printf("Sum: %d\n", a + b);
}

int main() {
    printSum(5, 3);   // 8
    printSum(10, 20); // 30
    
    return 0;
}
```

### 4. Functions With Parameters and Return Value

```c
#include <stdio.h>

int add(int a, int b) {
    return a + b;
}

int main() {
    int result = add(5, 3);
    printf("Result: %d\n", result);  // 8
    
    return 0;
}
```

---

## Function Declaration vs Definition

### Function Declaration (Prototype)

Tells the compiler about the function's existence.

```c
return_type function_name(parameter_types);
```

### Function Definition

Contains the actual implementation.

```c
return_type function_name(parameters) {
    // Function body
}
```

### Example

```c
#include <stdio.h>

// Function declarations (prototypes)
int add(int a, int b);
int subtract(int a, int b);
void printResult(int result);

int main() {
    int sum = add(10, 5);
    printResult(sum);
    
    int diff = subtract(10, 5);
    printResult(diff);
    
    return 0;
}

// Function definitions
int add(int a, int b) {
    return a + b;
}

int subtract(int a, int b) {
    return a - b;
}

void printResult(int result) {
    printf("Result: %d\n", result);
}
```

---

## Function Parameters

### Pass by Value

C uses **pass by value** - a copy of the argument is passed.

```c
#include <stdio.h>

void modify(int x) {
    x = 100;  // Only modifies the copy
    printf("Inside function: %d\n", x);
}

int main() {
    int num = 10;
    modify(num);
    printf("After function: %d\n", num);  // Still 10
    
    return 0;
}
```

**Output:**
```
Inside function: 100
After function: 10
```

### Pass by Reference (Using Pointers)

To modify the original value, pass a pointer.

```c
#include <stdio.h>

void modify(int *x) {
    *x = 100;  // Modifies the original
}

int main() {
    int num = 10;
    modify(&num);  // Pass address
    printf("After function: %d\n", num);  // 100
    
    return 0;
}
```

**Output:**
```
After function: 100
```

### Multiple Parameters

```c
#include <stdio.h>

int calculate(int a, int b, char op) {
    switch(op) {
        case '+': return a + b;
        case '-': return a - b;
        case '*': return a * b;
        case '/': return b != 0 ? a / b : 0;
        default: return 0;
    }
}

int main() {
    printf("%d\n", calculate(10, 5, '+'));  // 15
    printf("%d\n", calculate(10, 5, '*'));  // 50
    
    return 0;
}
```

---

## Return Statement

### Returning a Value

```c
int square(int n) {
    return n * n;
}
```

### Multiple Return Statements

```c
int max(int a, int b) {
    if (a > b) {
        return a;
    } else {
        return b;
    }
}

// Or using ternary operator
int max2(int a, int b) {
    return (a > b) ? a : b;
}
```

### Early Return

```c
int divide(int a, int b) {
    if (b == 0) {
        printf("Error: Division by zero\n");
        return 0;  // Early return
    }
    return a / b;
}
```

### void Functions (No Return Value)

```c
void printMessage() {
    printf("Hello\n");
    return;  // Optional in void functions
}
```

---

## Common Function Examples

### Example 1: Check Even or Odd

```c
#include <stdio.h>

int isEven(int num) {
    return num % 2 == 0;
}

int main() {
    int number = 7;
    
    if (isEven(number)) {
        printf("%d is even\n", number);
    } else {
        printf("%d is odd\n", number);
    }
    
    return 0;
}
```

### Example 2: Factorial

```c
#include <stdio.h>

unsigned long long factorial(int n) {
    unsigned long long result = 1;
    
    for (int i = 1; i <= n; i++) {
        result *= i;
    }
    
    return result;
}

int main() {
    int num = 5;
    printf("Factorial of %d = %llu\n", num, factorial(num));
    
    return 0;
}
```

### Example 3: Check Prime Number

```c
#include <stdio.h>

int isPrime(int n) {
    if (n <= 1) return 0;
    
    for (int i = 2; i * i <= n; i++) {
        if (n % i == 0) {
            return 0;  // Not prime
        }
    }
    
    return 1;  // Prime
}

int main() {
    int num = 17;
    
    if (isPrime(num)) {
        printf("%d is prime\n", num);
    } else {
        printf("%d is not prime\n", num);
    }
    
    return 0;
}
```

### Example 4: Find Maximum of Three Numbers

```c
#include <stdio.h>

int max3(int a, int b, int c) {
    int max = a;
    
    if (b > max) max = b;
    if (c > max) max = c;
    
    return max;
}

int main() {
    printf("Maximum: %d\n", max3(10, 25, 15));  // 25
    
    return 0;
}
```

### Example 5: Calculate Power

```c
#include <stdio.h>

double power(double base, int exponent) {
    double result = 1.0;
    
    for (int i = 0; i < exponent; i++) {
        result *= base;
    }
    
    return result;
}

int main() {
    printf("2^5 = %.0f\n", power(2, 5));      // 32
    printf("3^4 = %.0f\n", power(3, 4));      // 81
    printf("1.5^3 = %.2f\n", power(1.5, 3));  // 3.38
    
    return 0;
}
```

---

## Recursion

A function that calls itself.

### Example 1: Factorial (Recursive)

```c
#include <stdio.h>

int factorial(int n) {
    if (n <= 1) {
        return 1;  // Base case
    }
    return n * factorial(n - 1);  // Recursive call
}

int main() {
    printf("Factorial of 5 = %d\n", factorial(5));  // 120
    
    return 0;
}
```

### Example 2: Fibonacci (Recursive)

```c
#include <stdio.h>

int fibonacci(int n) {
    if (n <= 1) {
        return n;  // Base case
    }
    return fibonacci(n - 1) + fibonacci(n - 2);  // Recursive call
}

int main() {
    printf("Fibonacci sequence: ");
    for (int i = 0; i < 10; i++) {
        printf("%d ", fibonacci(i));
    }
    printf("\n");
    
    return 0;
}
```

### Example 3: Sum of Digits (Recursive)

```c
#include <stdio.h>

int sumDigits(int n) {
    if (n == 0) {
        return 0;  // Base case
    }
    return (n % 10) + sumDigits(n / 10);  // Recursive call
}

int main() {
    printf("Sum of digits of 12345 = %d\n", sumDigits(12345));  // 15
    
    return 0;
}
```

### Example 4: GCD (Recursive - Euclidean Algorithm)

```c
#include <stdio.h>

int gcd(int a, int b) {
    if (b == 0) {
        return a;  // Base case
    }
    return gcd(b, a % b);  // Recursive call
}

int main() {
    printf("GCD of 48 and 18 = %d\n", gcd(48, 18));  // 6
    
    return 0;
}
```

---

## Arrays as Function Parameters

### Passing Arrays

```c
#include <stdio.h>

void printArray(int arr[], int size) {
    for (int i = 0; i < size; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");
}

int main() {
    int numbers[] = {10, 20, 30, 40, 50};
    int size = sizeof(numbers) / sizeof(numbers[0]);
    
    printArray(numbers, size);
    
    return 0;
}
```

### Modifying Arrays

```c
#include <stdio.h>

void doubleValues(int arr[], int size) {
    for (int i = 0; i < size; i++) {
        arr[i] *= 2;  // Modifies original array
    }
}

int main() {
    int numbers[] = {1, 2, 3, 4, 5};
    int size = sizeof(numbers) / sizeof(numbers[0]);
    
    printf("Before: ");
    for (int i = 0; i < size; i++) printf("%d ", numbers[i]);
    printf("\n");
    
    doubleValues(numbers, size);
    
    printf("After: ");
    for (int i = 0; i < size; i++) printf("%d ", numbers[i]);
    printf("\n");
    
    return 0;
}
```

**Output:**
```
Before: 1 2 3 4 5 
After: 2 4 6 8 10 
```

### Array Functions

```c
#include <stdio.h>

int sumArray(int arr[], int size) {
    int sum = 0;
    for (int i = 0; i < size; i++) {
        sum += arr[i];
    }
    return sum;
}

int findMax(int arr[], int size) {
    int max = arr[0];
    for (int i = 1; i < size; i++) {
        if (arr[i] > max) {
            max = arr[i];
        }
    }
    return max;
}

float average(int arr[], int size) {
    return (float)sumArray(arr, size) / size;
}

int main() {
    int numbers[] = {45, 23, 78, 12, 67};
    int size = sizeof(numbers) / sizeof(numbers[0]);
    
    printf("Sum: %d\n", sumArray(numbers, size));
    printf("Max: %d\n", findMax(numbers, size));
    printf("Average: %.2f\n", average(numbers, size));
    
    return 0;
}
```

---

## Function with Multiple Return Values (Using Pointers)

```c
#include <stdio.h>

void calculate(int a, int b, int *sum, int *product) {
    *sum = a + b;
    *product = a * b;
}

int main() {
    int x = 5, y = 3;
    int sum, product;
    
    calculate(x, y, &sum, &product);
    
    printf("Sum: %d\n", sum);        // 8
    printf("Product: %d\n", product); // 15
    
    return 0;
}
```

---

## Variable Scope in Functions

### Local Variables

```c
#include <stdio.h>

void myFunction() {
    int x = 10;  // Local to myFunction
    printf("Inside function: %d\n", x);
}

int main() {
    myFunction();
    // printf("%d", x);  // ❌ Error: x not accessible here
    
    return 0;
}
```

### Global Variables

```c
#include <stdio.h>

int globalVar = 100;  // Global variable

void modifyGlobal() {
    globalVar = 200;  // Can access and modify
}

int main() {
    printf("Before: %d\n", globalVar);  // 100
    modifyGlobal();
    printf("After: %d\n", globalVar);   // 200
    
    return 0;
}
```

### Static Variables

```c
#include <stdio.h>

void counter() {
    static int count = 0;  // Initialized only once
    count++;
    printf("Count: %d\n", count);
}

int main() {
    counter();  // Count: 1
    counter();  // Count: 2
    counter();  // Count: 3
    
    return 0;
}
```

---

## Function Pointers

```c
#include <stdio.h>

int add(int a, int b) {
    return a + b;
}

int subtract(int a, int b) {
    return a - b;
}

int main() {
    // Function pointer
    int (*operation)(int, int);
    
    operation = add;
    printf("Add: %d\n", operation(10, 5));  // 15
    
    operation = subtract;
    printf("Subtract: %d\n", operation(10, 5));  // 5
    
    return 0;
}
```

---

## Inline Functions (C99)

Suggests to the compiler to insert the function code directly.

```c
#include <stdio.h>

inline int square(int x) {
    return x * x;
}

int main() {
    printf("%d\n", square(5));  // 25
    
    return 0;
}
```

---

## Common Pitfalls

### 1. Missing Return Statement

```c
// ❌ BAD: No return statement
int getValue() {
    int x = 10;
    // Missing return!
}

// ✅ GOOD
int getValue() {
    int x = 10;
    return x;
}
```

### 2. Returning Local Variable Address

```c
// ❌ BAD: Returns address of local variable
int* getArray() {
    int arr[5] = {1, 2, 3, 4, 5};
    return arr;  // Undefined behavior!
}

// ✅ GOOD: Use static or dynamic allocation
int* getArray() {
    static int arr[5] = {1, 2, 3, 4, 5};
    return arr;
}
```

### 3. Modifying Array Size in Function

```c
// ❌ BAD: sizeof doesn't work as expected
void printArray(int arr[]) {
    int size = sizeof(arr) / sizeof(arr[0]);  // Wrong!
    // arr is a pointer, not an array
}

// ✅ GOOD: Pass size as parameter
void printArray(int arr[], int size) {
    for (int i = 0; i < size; i++) {
        printf("%d ", arr[i]);
    }
}
```

### 4. Function Declaration Mismatch

```c
// ❌ BAD: Declaration and definition don't match
int add(int a, int b);  // Declaration

float add(int a, int b) {  // Different return type!
    return a + b;
}

// ✅ GOOD: Match exactly
int add(int a, int b);  // Declaration

int add(int a, int b) {  // Same return type
    return a + b;
}
```

---

## Practical Function Library Example

```c
#include <stdio.h>

// Math functions
int add(int a, int b) { return a + b; }
int subtract(int a, int b) { return a - b; }
int multiply(int a, int b) { return a * b; }
float divide(int a, int b) { return b != 0 ? (float)a / b : 0; }

// Utility functions
void swap(int *a, int *b) {
    int temp = *a;
    *a = *b;
    *b = temp;
}

int isEven(int n) { return n % 2 == 0; }
int isOdd(int n) { return n % 2 != 0; }

int max(int a, int b) { return (a > b) ? a : b; }
int min(int a, int b) { return (a < b) ? a : b; }

// String functions
int stringLength(char str[]) {
    int len = 0;
    while (str[len] != '\0') {
        len++;
    }
    return len;
}

void reverseString(char str[]) {
    int len = stringLength(str);
    for (int i = 0; i < len / 2; i++) {
        char temp = str[i];
        str[i] = str[len - 1 - i];
        str[len - 1 - i] = temp;
    }
}

int main() {
    // Test math functions
    printf("10 + 5 = %d\n", add(10, 5));
    printf("10 - 5 = %d\n", subtract(10, 5));
    printf("10 * 5 = %d\n", multiply(10, 5));
    printf("10 / 5 = %.2f\n", divide(10, 5));
    
    // Test utility functions
    int x = 10, y = 20;
    printf("\nBefore swap: x=%d, y=%d\n", x, y);
    swap(&x, &y);
    printf("After swap: x=%d, y=%d\n", x, y);
    
    printf("\n7 is %s\n", isEven(7) ? "even" : "odd");
    printf("Max of 10 and 20: %d\n", max(10, 20));
    
    // Test string functions
    char str[] = "Hello";
    printf("\nLength of '%s': %d\n", str, stringLength(str));
    reverseString(str);
    printf("Reversed: %s\n", str);
    
    return 0;
}
```

---

## Best Practices

1. **Use descriptive function names** - `calculateTotal()` not `calc()`
2. **Keep functions small** - One function, one task
3. **Use function prototypes** - Declare before use
4. **Document functions** - Add comments explaining purpose
5. **Limit parameters** - Max 3-5 parameters
6. **Validate input** - Check for invalid values
7. **Use const for read-only parameters** - `void print(const int arr[])`
8. **Avoid global variables** - Pass parameters instead
9. **Return early** - Exit on error conditions
10. **Test functions independently** - Unit testing

```c
// ✅ GOOD: Well-documented function
/**
 * Calculates the factorial of a number
 * @param n The number to calculate factorial for
 * @return The factorial of n, or 1 if n <= 0
 */
unsigned long long factorial(int n) {
    if (n <= 0) return 1;  // Early return
    
    unsigned long long result = 1;
    for (int i = 1; i <= n; i++) {
        result *= i;
    }
    return result;
}
```

---

## Function Summary

| Concept | Example | Description |
|---------|---------|-------------|
| **Declaration** | `int add(int, int);` | Function prototype |
| **Definition** | `int add(int a, int b) { }` | Implementation |
| **Call** | `add(5, 3);` | Execute function |
| **Return** | `return value;` | Return a value |
| **void** | `void func()` | No return value |
| **Parameters** | `func(int a, int b)` | Input values |
| **Pass by value** | `func(x)` | Copy of value |
| **Pass by reference** | `func(&x)` | Address of value |
| **Recursion** | `fact(n) = n * fact(n-1)` | Function calls itself |
| **Array parameter** | `func(int arr[], int size)` | Pass array |

Functions are the building blocks of modular, maintainable C code - master them to write professional programs!