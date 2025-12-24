# Loops in C

## Table of Contents

- [What are Loops?](#what-are-loops)
- [Types of Loops in C](#types-of-loops-in-c)
- [1. for Loop](#1-for-loop)
- [2. while Loop](#2-while-loop)
- [3. do-while Loop](#3-do-while-loop)
- [4. Nested Loops](#4-nested-loops)
- [Loop Control Statements](#loop-control-statements)
- [Loop Comparison](#loop-comparison)
- [Common Loop Patterns](#common-loop-patterns)
- [Practical Examples](#practical-examples)
- [Common Pitfalls](#common-pitfalls)
- [Performance Tips](#performance-tips)
- [Best Practices](#best-practices)
- [Quick Reference](#quick-reference)

---

## What are Loops?

**Loops** allow you to execute a block of code repeatedly until a condition is met. They eliminate the need to write the same code multiple times.

```
Without loop:
printf("Hello\n");
printf("Hello\n");
printf("Hello\n");
printf("Hello\n");
printf("Hello\n");

With loop:
for (int i = 0; i < 5; i++) {
    printf("Hello\n");
}
```

---

## Types of Loops in C

1. **for loop** - When you know the number of iterations
2. **while loop** - When the number of iterations is unknown
3. **do-while loop** - Executes at least once
4. **Nested loops** - Loops inside loops

---

## 1. for Loop

Best used when you know **how many times** to repeat.

### Syntax

```c
for (initialization; condition; increment/decrement) {
    // code to execute
}
```

### Flow Diagram

```
1. Initialization (once)
    ↓
2. Check condition
    ↓
3. Execute body (if condition is true)
    ↓
4. Increment/Decrement
    ↓
Back to step 2
```

### Basic Example

```c
#include <stdio.h>

int main() {
    for (int i = 0; i < 5; i++) {
        printf("Iteration %d\n", i);
    }
    
    return 0;
}
```

**Output:**
```
Iteration 0
Iteration 1
Iteration 2
Iteration 3
Iteration 4
```

### Example: Sum of Numbers

```c
#include <stdio.h>

int main() {
    int sum = 0;
    
    for (int i = 1; i <= 10; i++) {
        sum += i;
    }
    
    printf("Sum of 1 to 10: %d\n", sum);  // 55
    
    return 0;
}
```

### Example: Multiplication Table

```c
#include <stdio.h>

int main() {
    int num = 5;
    
    printf("Multiplication table of %d:\n", num);
    
    for (int i = 1; i <= 10; i++) {
        printf("%d x %d = %d\n", num, i, num * i);
    }
    
    return 0;
}
```

**Output:**
```
Multiplication table of 5:
5 x 1 = 5
5 x 2 = 10
5 x 3 = 15
...
5 x 10 = 50
```

### Counting Backwards

```c
#include <stdio.h>

int main() {
    for (int i = 10; i >= 1; i--) {
        printf("%d ", i);
    }
    printf("\nBlastoff!\n");
    
    return 0;
}
```

**Output:**
```
10 9 8 7 6 5 4 3 2 1 
Blastoff!
```

### Custom Increment

```c
#include <stdio.h>

int main() {
    // Count by 2s
    for (int i = 0; i <= 20; i += 2) {
        printf("%d ", i);
    }
    printf("\n");
    
    // Count by 5s
    for (int i = 0; i <= 50; i += 5) {
        printf("%d ", i);
    }
    printf("\n");
    
    return 0;
}
```

### Infinite for Loop

```c
// ⚠️ Runs forever!
for (;;) {
    printf("Infinite loop\n");
}

// Or
for (int i = 0; i < 10; i--) {  // i keeps decreasing
    printf("Never ends\n");
}
```

---

## 2. while Loop

Best used when the number of iterations is **unknown** beforehand.

### Syntax

```c
while (condition) {
    // code to execute
}
```

### Basic Example

```c
#include <stdio.h>

int main() {
    int i = 0;
    
    while (i < 5) {
        printf("Iteration %d\n", i);
        i++;
    }
    
    return 0;
}
```

**Output:**
```
Iteration 0
Iteration 1
Iteration 2
Iteration 3
Iteration 4
```

### Example: User Input Until Valid

```c
#include <stdio.h>

int main() {
    int num;
    
    printf("Enter a positive number: ");
    scanf("%d", &num);
    
    while (num <= 0) {
        printf("Invalid! Enter a positive number: ");
        scanf("%d", &num);
    }
    
    printf("You entered: %d\n", num);
    
    return 0;
}
```

### Example: Sum Until User Quits

```c
#include <stdio.h>

int main() {
    int num, sum = 0;
    
    printf("Enter numbers to sum (0 to quit):\n");
    
    scanf("%d", &num);
    
    while (num != 0) {
        sum += num;
        printf("Current sum: %d\n", sum);
        printf("Enter next number (0 to quit): ");
        scanf("%d", &num);
    }
    
    printf("Final sum: %d\n", sum);
    
    return 0;
}
```

### Example: Count Digits

```c
#include <stdio.h>

int main() {
    int num, count = 0;
    
    printf("Enter a number: ");
    scanf("%d", &num);
    
    int temp = num;
    
    while (temp != 0) {
        temp /= 10;
        count++;
    }
    
    printf("Number of digits: %d\n", count);
    
    return 0;
}
```

### Infinite while Loop

```c
// ⚠️ Runs forever!
while (1) {
    printf("Infinite loop\n");
}

// Or
while (true) {  // If using stdbool.h
    printf("Infinite loop\n");
}
```

---

## 3. do-while Loop

Executes the body **at least once**, then checks the condition.

### Syntax

```c
do {
    // code to execute
} while (condition);
```

### Basic Example

```c
#include <stdio.h>

int main() {
    int i = 0;
    
    do {
        printf("Iteration %d\n", i);
        i++;
    } while (i < 5);
    
    return 0;
}
```

### Key Difference: Executes At Least Once

```c
#include <stdio.h>

int main() {
    int x = 10;
    
    // while loop - doesn't execute
    while (x < 5) {
        printf("while: This won't print\n");
    }
    
    // do-while loop - executes once
    do {
        printf("do-while: This prints once\n");
    } while (x < 5);
    
    return 0;
}
```

**Output:**
```
do-while: This prints once
```

### Example: Menu System

```c
#include <stdio.h>

int main() {
    int choice;
    
    do {
        printf("\n=== Menu ===\n");
        printf("1. Option 1\n");
        printf("2. Option 2\n");
        printf("3. Option 3\n");
        printf("4. Exit\n");
        printf("Enter choice: ");
        scanf("%d", &choice);
        
        switch (choice) {
            case 1:
                printf("You selected Option 1\n");
                break;
            case 2:
                printf("You selected Option 2\n");
                break;
            case 3:
                printf("You selected Option 3\n");
                break;
            case 4:
                printf("Goodbye!\n");
                break;
            default:
                printf("Invalid choice\n");
        }
    } while (choice != 4);
    
    return 0;
}
```

### Example: Guess the Number

```c
#include <stdio.h>

int main() {
    int secret = 42;
    int guess;
    
    do {
        printf("Guess the number (1-100): ");
        scanf("%d", &guess);
        
        if (guess < secret) {
            printf("Too low!\n");
        } else if (guess > secret) {
            printf("Too high!\n");
        } else {
            printf("Correct! You win!\n");
        }
    } while (guess != secret);
    
    return 0;
}
```

---

## 4. Nested Loops

Loops inside other loops.

### Basic Nested Loop

```c
#include <stdio.h>

int main() {
    for (int i = 1; i <= 3; i++) {
        for (int j = 1; j <= 3; j++) {
            printf("i=%d, j=%d\n", i, j);
        }
    }
    
    return 0;
}
```

**Output:**
```
i=1, j=1
i=1, j=2
i=1, j=3
i=2, j=1
i=2, j=2
i=2, j=3
i=3, j=1
i=3, j=2
i=3, j=3
```

### Example: Print Pattern

```c
#include <stdio.h>

int main() {
    int rows = 5;
    
    for (int i = 1; i <= rows; i++) {
        for (int j = 1; j <= i; j++) {
            printf("* ");
        }
        printf("\n");
    }
    
    return 0;
}
```

**Output:**
```
* 
* * 
* * * 
* * * * 
* * * * * 
```

### Example: Multiplication Table (Grid)

```c
#include <stdio.h>

int main() {
    printf("   ");
    for (int i = 1; i <= 10; i++) {
        printf("%4d", i);
    }
    printf("\n");
    printf("   ");
    for (int i = 1; i <= 10; i++) {
        printf("----");
    }
    printf("\n");
    
    for (int i = 1; i <= 10; i++) {
        printf("%2d |", i);
        for (int j = 1; j <= 10; j++) {
            printf("%4d", i * j);
        }
        printf("\n");
    }
    
    return 0;
}
```

### Example: 2D Array Iteration

```c
#include <stdio.h>

int main() {
    int matrix[3][3] = {
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9}
    };
    
    for (int i = 0; i < 3; i++) {
        for (int j = 0; j < 3; j++) {
            printf("%d ", matrix[i][j]);
        }
        printf("\n");
    }
    
    return 0;
}
```

---

## Loop Control Statements

### break - Exit the Loop

```c
#include <stdio.h>

int main() {
    for (int i = 1; i <= 10; i++) {
        if (i == 5) {
            break;  // Exit loop when i is 5
        }
        printf("%d ", i);
    }
    printf("\n");
    
    return 0;
}
```

**Output:**
```
1 2 3 4 
```

### continue - Skip Current Iteration

```c
#include <stdio.h>

int main() {
    for (int i = 1; i <= 10; i++) {
        if (i % 2 == 0) {
            continue;  // Skip even numbers
        }
        printf("%d ", i);
    }
    printf("\n");
    
    return 0;
}
```

**Output:**
```
1 3 5 7 9 
```

### Example: Search with break

```c
#include <stdio.h>

int main() {
    int numbers[] = {10, 20, 30, 40, 50};
    int search = 30;
    int found = 0;
    
    for (int i = 0; i < 5; i++) {
        if (numbers[i] == search) {
            printf("Found %d at index %d\n", search, i);
            found = 1;
            break;  // Stop searching once found
        }
    }
    
    if (!found) {
        printf("%d not found\n", search);
    }
    
    return 0;
}
```

### Example: Skip Negative Numbers

```c
#include <stdio.h>

int main() {
    int numbers[] = {5, -3, 8, -1, 12, -7, 15};
    int sum = 0;
    
    for (int i = 0; i < 7; i++) {
        if (numbers[i] < 0) {
            continue;  // Skip negative numbers
        }
        sum += numbers[i];
    }
    
    printf("Sum of positive numbers: %d\n", sum);  // 40
    
    return 0;
}
```

### break in Nested Loops

```c
#include <stdio.h>

int main() {
    for (int i = 1; i <= 3; i++) {
        for (int j = 1; j <= 3; j++) {
            if (j == 2) {
                break;  // Only breaks inner loop
            }
            printf("i=%d, j=%d\n", i, j);
        }
    }
    
    return 0;
}
```

**Output:**
```
i=1, j=1
i=2, j=1
i=3, j=1
```

---

## Loop Comparison

| Feature | for | while | do-while |
|---------|-----|-------|----------|
| **Use Case** | Known iterations | Unknown iterations | At least one iteration |
| **When to Test** | Before each iteration | Before each iteration | After each iteration |
| **Minimum Executions** | 0 | 0 | 1 |
| **Syntax Complexity** | More compact | Simple | Simple |
| **Best For** | Counters, arrays | Input validation | Menu systems |

---

## Common Loop Patterns

### 1. Counting Pattern

```c
// Count up
for (int i = 0; i < n; i++) { }

// Count down
for (int i = n; i > 0; i--) { }

// Count by steps
for (int i = 0; i < n; i += 2) { }
```

### 2. Array Iteration

```c
int arr[5] = {10, 20, 30, 40, 50};

for (int i = 0; i < 5; i++) {
    printf("%d ", arr[i]);
}
```

### 3. Accumulator Pattern

```c
int sum = 0;
for (int i = 1; i <= n; i++) {
    sum += i;
}
```

### 4. Flag Pattern

```c
int found = 0;
for (int i = 0; i < n && !found; i++) {
    if (arr[i] == target) {
        found = 1;
    }
}
```

### 5. Sentinel Pattern

```c
int value;
while (1) {
    scanf("%d", &value);
    if (value == -1) break;  // Sentinel value
    // Process value
}
```

---

## Practical Examples

### Example 1: Factorial

```c
#include <stdio.h>

int main() {
    int n;
    unsigned long long factorial = 1;
    
    printf("Enter a number: ");
    scanf("%d", &n);
    
    for (int i = 1; i <= n; i++) {
        factorial *= i;
    }
    
    printf("Factorial of %d = %llu\n", n, factorial);
    
    return 0;
}
```

### Example 2: Fibonacci Series

```c
#include <stdio.h>

int main() {
    int n, first = 0, second = 1, next;
    
    printf("Enter number of terms: ");
    scanf("%d", &n);
    
    printf("Fibonacci Series: ");
    
    for (int i = 0; i < n; i++) {
        if (i <= 1) {
            next = i;
        } else {
            next = first + second;
            first = second;
            second = next;
        }
        printf("%d ", next);
    }
    printf("\n");
    
    return 0;
}
```

### Example 3: Prime Number Checker

```c
#include <stdio.h>

int main() {
    int n, isPrime = 1;
    
    printf("Enter a number: ");
    scanf("%d", &n);
    
    if (n <= 1) {
        isPrime = 0;
    } else {
        for (int i = 2; i * i <= n; i++) {
            if (n % i == 0) {
                isPrime = 0;
                break;
            }
        }
    }
    
    if (isPrime) {
        printf("%d is prime\n", n);
    } else {
        printf("%d is not prime\n", n);
    }
    
    return 0;
}
```

### Example 4: Reverse a Number

```c
#include <stdio.h>

int main() {
    int num, reversed = 0, remainder;
    
    printf("Enter a number: ");
    scanf("%d", &num);
    
    int original = num;
    
    while (num != 0) {
        remainder = num % 10;
        reversed = reversed * 10 + remainder;
        num /= 10;
    }
    
    printf("Original: %d\n", original);
    printf("Reversed: %d\n", reversed);
    
    return 0;
}
```

### Example 5: GCD (Greatest Common Divisor)

```c
#include <stdio.h>

int main() {
    int a, b;
    
    printf("Enter two numbers: ");
    scanf("%d %d", &a, &b);
    
    // Euclidean algorithm
    while (b != 0) {
        int temp = b;
        b = a % b;
        a = temp;
    }
    
    printf("GCD: %d\n", a);
    
    return 0;
}
```

### Example 6: Print Diamond Pattern

```c
#include <stdio.h>

int main() {
    int n = 5;
    
    // Upper half
    for (int i = 1; i <= n; i++) {
        // Print spaces
        for (int j = 1; j <= n - i; j++) {
            printf(" ");
        }
        // Print stars
        for (int j = 1; j <= 2 * i - 1; j++) {
            printf("*");
        }
        printf("\n");
    }
    
    // Lower half
    for (int i = n - 1; i >= 1; i--) {
        // Print spaces
        for (int j = 1; j <= n - i; j++) {
            printf(" ");
        }
        // Print stars
        for (int j = 1; j <= 2 * i - 1; j++) {
            printf("*");
        }
        printf("\n");
    }
    
    return 0;
}
```

**Output:**
```
    *
   ***
  *****
 *******
*********
 *******
  *****
   ***
    *
```

---

## Common Pitfalls

### 1. Off-by-One Error

```c
// ❌ BAD: Misses last element
for (int i = 0; i < 5 - 1; i++) { }

// ✅ GOOD
for (int i = 0; i < 5; i++) { }
```

### 2. Infinite Loop

```c
// ❌ BAD: Never increments
int i = 0;
while (i < 10) {
    printf("%d\n", i);
    // Missing: i++;
}

// ✅ GOOD
int i = 0;
while (i < 10) {
    printf("%d\n", i);
    i++;
}
```

### 3. Modifying Loop Variable

```c
// ❌ BAD: Unpredictable behavior
for (int i = 0; i < 10; i++) {
    i += 2;  // Don't modify loop variable inside
}

// ✅ GOOD: Use appropriate increment
for (int i = 0; i < 10; i += 3) { }
```

### 4. Floating Point Loop Counter

```c
// ❌ BAD: Precision issues
for (float f = 0.0; f < 1.0; f += 0.1) {
    printf("%.1f ", f);
}

// ✅ GOOD: Use integer counter
for (int i = 0; i < 10; i++) {
    float f = i * 0.1;
    printf("%.1f ", f);
}
```

### 5. Unintended Semicolon

```c
// ❌ BAD: Empty loop body
for (int i = 0; i < 10; i++);
    printf("Hello\n");  // Only prints once!

// ✅ GOOD
for (int i = 0; i < 10; i++) {
    printf("Hello\n");
}
```

---

## Performance Tips

### 1. Cache Loop Limit

```c
// ❌ SLOW: Recalculates length every iteration
for (int i = 0; i < strlen(str); i++) { }

// ✅ FAST: Calculate once
int len = strlen(str);
for (int i = 0; i < len; i++) { }
```

### 2. Use Appropriate Loop Type

```c
// Less efficient: while loop for counting
int i = 0;
while (i < 100) {
    // code
    i++;
}

// More efficient: for loop for counting
for (int i = 0; i < 100; i++) {
    // code
}
```

---

## Best Practices

1. **Choose the right loop type** for your use case
2. **Initialize loop variables** properly
3. **Avoid infinite loops** - ensure termination condition
4. **Use meaningful variable names** (not just `i`, `j`, `k`)
5. **Keep loop body simple** - extract complex code to functions
6. **Limit nesting depth** - max 2-3 levels
7. **Use break and continue judiciously** - don't overuse
8. **Comment complex loop logic**
9. **Watch for off-by-one errors**
10. **Test edge cases** (empty arrays, single element, etc.)

```c
// ✅ GOOD: Clear and readable
for (int studentIndex = 0; studentIndex < studentCount; studentIndex++) {
    processStudent(students[studentIndex]);
}

// ❌ BAD: Unclear and complex
for (int i = 0; i < n; i++) {
    for (int j = 0; j < m; j++) {
        for (int k = 0; k < p; k++) {
            // Complex nested logic...
        }
    }
}
```

---

## Quick Reference

```c
// for loop
for (int i = 0; i < n; i++) {
    // code
}

// while loop
while (condition) {
    // code
}

// do-while loop
do {
    // code
} while (condition);

// break - exit loop
break;

// continue - skip to next iteration
continue;

// Nested loop
for (int i = 0; i < rows; i++) {
    for (int j = 0; j < cols; j++) {
        // code
    }
}
```

Loops are essential for repetitive tasks in C - master them to write efficient, elegant code!