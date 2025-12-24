# Conditional Statements in C

## Table of Contents

- [What are Conditional Statements?](#what-are-conditional-statements)
- [Types of Conditional Statements](#types-of-conditional-statements)
- [1. if Statement](#1-if-statement)
- [2. if-else Statement](#2-if-else-statement)
- [3. if-else if-else Ladder](#3-if-else-if-else-ladder)
- [4. Nested if Statements](#4-nested-if-statements)
- [5. switch Statement](#5-switch-statement)
- [6. Ternary Operator (? :)](#6-ternary-operator--)
- [Comparison Operators](#comparison-operators)
- [Logical Operators](#logical-operators)
- [Truth Values in C](#truth-values-in-c)
- [Common Pitfalls](#common-pitfalls)
- [Practical Examples](#practical-examples)
- [if vs switch: When to Use](#if-vs-switch-when-to-use)
- [Summary Table](#summary-table)
- [Best Practices](#best-practices)

---

## What are Conditional Statements?

**Conditional statements** allow your program to make decisions and execute different code based on conditions.

```
if (condition is true) {
    execute this code
} else {
    execute this code
}
```

---

## Types of Conditional Statements

1. **if statement**
2. **if-else statement**
3. **if-else if-else ladder**
4. **Nested if statements**
5. **switch statement**
6. **Ternary operator (? :)**

---

## 1. if Statement

Executes code only if the condition is **true**.

### Syntax

```c
if (condition) {
    // code to execute if condition is true
}
```

### Example

```c
#include <stdio.h>

int main() {
    int age = 20;
    
    if (age >= 18) {
        printf("You are an adult\n");
    }
    
    printf("Program continues...\n");
    
    return 0;
}
```

**Output:**
```
You are an adult
Program continues...
```

### Single Statement (No Braces)

```c
int x = 10;

if (x > 5)
    printf("x is greater than 5\n");  // Single statement

// ⚠️ Best practice: Always use braces for clarity
if (x > 5) {
    printf("x is greater than 5\n");
}
```

---

## 2. if-else Statement

Executes one block if condition is **true**, another if **false**.

### Syntax

```c
if (condition) {
    // code if condition is true
} else {
    // code if condition is false
}
```

### Example

```c
#include <stdio.h>

int main() {
    int number = 7;
    
    if (number % 2 == 0) {
        printf("%d is even\n", number);
    } else {
        printf("%d is odd\n", number);
    }
    
    return 0;
}
```

**Output:**
```
7 is odd
```

### Example: Login System

```c
#include <stdio.h>
#include <string.h>

int main() {
    char password[20];
    
    printf("Enter password: ");
    scanf("%s", password);
    
    if (strcmp(password, "secret123") == 0) {
        printf("Access granted!\n");
    } else {
        printf("Access denied!\n");
    }
    
    return 0;
}
```

---

## 3. if-else if-else Ladder

Tests multiple conditions sequentially.

### Syntax

```c
if (condition1) {
    // code if condition1 is true
} else if (condition2) {
    // code if condition2 is true
} else if (condition3) {
    // code if condition3 is true
} else {
    // code if all conditions are false
}
```

### Example: Grade Calculator

```c
#include <stdio.h>

int main() {
    int score;
    
    printf("Enter your score (0-100): ");
    scanf("%d", &score);
    
    if (score >= 90) {
        printf("Grade: A (Excellent!)\n");
    } else if (score >= 80) {
        printf("Grade: B (Good)\n");
    } else if (score >= 70) {
        printf("Grade: C (Average)\n");
    } else if (score >= 60) {
        printf("Grade: D (Below Average)\n");
    } else {
        printf("Grade: F (Fail)\n");
    }
    
    return 0;
}
```

### Example: Temperature Classifier

```c
#include <stdio.h>

int main() {
    float temp;
    
    printf("Enter temperature (°C): ");
    scanf("%f", &temp);
    
    if (temp < 0) {
        printf("Freezing\n");
    } else if (temp >= 0 && temp < 10) {
        printf("Very Cold\n");
    } else if (temp >= 10 && temp < 20) {
        printf("Cold\n");
    } else if (temp >= 20 && temp < 30) {
        printf("Warm\n");
    } else {
        printf("Hot\n");
    }
    
    return 0;
}
```

---

## 4. Nested if Statements

**if statements** inside other **if statements**.

### Syntax

```c
if (condition1) {
    if (condition2) {
        // code if both conditions are true
    }
}
```

### Example: Age and License Check

```c
#include <stdio.h>

int main() {
    int age;
    char hasLicense;
    
    printf("Enter your age: ");
    scanf("%d", &age);
    
    printf("Do you have a license? (y/n): ");
    scanf(" %c", &hasLicense);
    
    if (age >= 18) {
        if (hasLicense == 'y' || hasLicense == 'Y') {
            printf("You can drive!\n");
        } else {
            printf("You need a license to drive\n");
        }
    } else {
        printf("You are too young to drive\n");
    }
    
    return 0;
}
```

### Example: Number Classification

```c
#include <stdio.h>

int main() {
    int num;
    
    printf("Enter a number: ");
    scanf("%d", &num);
    
    if (num > 0) {
        if (num % 2 == 0) {
            printf("Positive Even Number\n");
        } else {
            printf("Positive Odd Number\n");
        }
    } else if (num < 0) {
        if (num % 2 == 0) {
            printf("Negative Even Number\n");
        } else {
            printf("Negative Odd Number\n");
        }
    } else {
        printf("Number is Zero\n");
    }
    
    return 0;
}
```

---

## 5. switch Statement

Tests a variable against multiple **constant** values.

### Syntax

```c
switch (expression) {
    case constant1:
        // code
        break;
    case constant2:
        // code
        break;
    default:
        // code if no case matches
}
```

### Example: Menu System

```c
#include <stdio.h>

int main() {
    int choice;
    
    printf("=== Menu ===\n");
    printf("1. Add\n");
    printf("2. Subtract\n");
    printf("3. Multiply\n");
    printf("4. Divide\n");
    printf("Enter choice: ");
    scanf("%d", &choice);
    
    switch (choice) {
        case 1:
            printf("You selected Addition\n");
            break;
        case 2:
            printf("You selected Subtraction\n");
            break;
        case 3:
            printf("You selected Multiplication\n");
            break;
        case 4:
            printf("You selected Division\n");
            break;
        default:
            printf("Invalid choice\n");
    }
    
    return 0;
}
```

### Example: Day of Week

```c
#include <stdio.h>

int main() {
    int day;
    
    printf("Enter day number (1-7): ");
    scanf("%d", &day);
    
    switch (day) {
        case 1:
            printf("Monday\n");
            break;
        case 2:
            printf("Tuesday\n");
            break;
        case 3:
            printf("Wednesday\n");
            break;
        case 4:
            printf("Thursday\n");
            break;
        case 5:
            printf("Friday\n");
            break;
        case 6:
            printf("Saturday\n");
            break;
        case 7:
            printf("Sunday\n");
            break;
        default:
            printf("Invalid day number\n");
    }
    
    return 0;
}
```

### Fall-Through Behavior

Without `break`, execution continues to the next case:

```c
#include <stdio.h>

int main() {
    char grade;
    
    printf("Enter grade (A-F): ");
    scanf(" %c", &grade);
    
    switch (grade) {
        case 'A':
        case 'a':
            printf("Excellent! (90-100)\n");
            break;
        case 'B':
        case 'b':
            printf("Good! (80-89)\n");
            break;
        case 'C':
        case 'c':
            printf("Average (70-79)\n");
            break;
        case 'D':
        case 'd':
            printf("Below Average (60-69)\n");
            break;
        case 'F':
        case 'f':
            printf("Fail (0-59)\n");
            break;
        default:
            printf("Invalid grade\n");
    }
    
    return 0;
}
```

### Example: Vowel or Consonant

```c
#include <stdio.h>

int main() {
    char ch;
    
    printf("Enter a character: ");
    scanf(" %c", &ch);
    
    switch (ch) {
        case 'a':
        case 'e':
        case 'i':
        case 'o':
        case 'u':
        case 'A':
        case 'E':
        case 'I':
        case 'O':
        case 'U':
            printf("%c is a vowel\n", ch);
            break;
        default:
            printf("%c is a consonant\n", ch);
    }
    
    return 0;
}
```

---

## 6. Ternary Operator (? :)

Compact way to write simple if-else statements.

### Syntax

```c
condition ? expression_if_true : expression_if_false;
```

### Example: Find Maximum

```c
#include <stdio.h>

int main() {
    int a = 10, b = 20;
    int max;
    
    max = (a > b) ? a : b;
    
    printf("Maximum: %d\n", max);  // 20
    
    return 0;
}
```

### Example: Even or Odd

```c
#include <stdio.h>

int main() {
    int num = 7;
    
    printf("%d is %s\n", num, (num % 2 == 0) ? "even" : "odd");
    
    return 0;
}
```

### Example: Absolute Value

```c
#include <stdio.h>

int main() {
    int num = -15;
    int absolute;
    
    absolute = (num < 0) ? -num : num;
    
    printf("Absolute value: %d\n", absolute);  // 15
    
    return 0;
}
```

### Nested Ternary (Use Sparingly)

```c
#include <stdio.h>

int main() {
    int num = 0;
    
    char *result = (num > 0) ? "Positive" : 
                   (num < 0) ? "Negative" : 
                   "Zero";
    
    printf("Number is: %s\n", result);
    
    return 0;
}
```

---

## Comparison Operators

Used in conditional expressions:

| Operator | Meaning | Example | Result |
|----------|---------|---------|--------|
| `==` | Equal to | `5 == 5` | true |
| `!=` | Not equal to | `5 != 3` | true |
| `>` | Greater than | `5 > 3` | true |
| `<` | Less than | `3 < 5` | true |
| `>=` | Greater than or equal | `5 >= 5` | true |
| `<=` | Less than or equal | `3 <= 5` | true |

```c
int x = 10, y = 20;

if (x == y) { }   // false
if (x != y) { }   // true
if (x > y) { }    // false
if (x < y) { }    // true
if (x >= 10) { }  // true
if (x <= 5) { }   // false
```

---

## Logical Operators

Combine multiple conditions:

| Operator | Meaning | Example | Result |
|----------|---------|---------|--------|
| `&&` | AND | `(5 > 3) && (2 < 4)` | true |
| `\|\|` | OR | `(5 > 3) \|\| (2 > 4)` | true |
| `!` | NOT | `!(5 > 3)` | false |

### AND (&&) - Both conditions must be true

```c
#include <stdio.h>

int main() {
    int age = 25;
    int income = 50000;
    
    if (age >= 18 && income >= 30000) {
        printf("Loan approved\n");
    } else {
        printf("Loan denied\n");
    }
    
    return 0;
}
```

### OR (||) - At least one condition must be true

```c
#include <stdio.h>

int main() {
    char day[10];
    
    printf("Enter day: ");
    scanf("%s", day);
    
    if (strcmp(day, "Saturday") == 0 || strcmp(day, "Sunday") == 0) {
        printf("It's a weekend!\n");
    } else {
        printf("It's a weekday\n");
    }
    
    return 0;
}
```

### NOT (!) - Negates a condition

```c
#include <stdio.h>

int main() {
    int logged_in = 0;  // false
    
    if (!logged_in) {
        printf("Please log in\n");
    } else {
        printf("Welcome!\n");
    }
    
    return 0;
}
```

### Combining Logical Operators

```c
#include <stdio.h>

int main() {
    int age = 25;
    int hasTicket = 1;
    int hasID = 1;
    
    if ((age >= 18 && hasTicket) && hasID) {
        printf("You can enter the club\n");
    } else {
        printf("Entry denied\n");
    }
    
    return 0;
}
```

---

## Truth Values in C

In C, any **non-zero** value is considered **true**, and **zero** is **false**.

```c
#include <stdio.h>

int main() {
    if (1) printf("True\n");        // Prints
    if (0) printf("False\n");       // Doesn't print
    if (42) printf("True\n");       // Prints
    if (-1) printf("True\n");       // Prints
    
    int x = 10;
    if (x) printf("x is true\n");   // Prints (10 is non-zero)
    
    x = 0;
    if (x) printf("x is true\n");   // Doesn't print
    
    return 0;
}
```

---

## Common Pitfalls

### 1. Assignment vs Comparison

```c
int x = 10;

// ❌ BAD: Assignment, always true
if (x = 5) {
    printf("This always executes\n");
}

// ✅ GOOD: Comparison
if (x == 5) {
    printf("x equals 5\n");
}
```

### 2. Floating Point Comparison

```c
float a = 0.1 + 0.2;
float b = 0.3;

// ❌ BAD: May fail due to precision
if (a == b) { }

// ✅ GOOD: Use epsilon comparison
if (fabs(a - b) < 0.0001) { }
```

### 3. Semicolon After if

```c
// ❌ BAD: Empty if body
if (x > 5);
    printf("This always executes\n");

// ✅ GOOD
if (x > 5) {
    printf("x is greater than 5\n");
}
```

### 4. Forgetting break in switch

```c
// ❌ BAD: Fall-through unintended
switch (x) {
    case 1:
        printf("One\n");
        // Falls through to case 2!
    case 2:
        printf("Two\n");
        break;
}

// ✅ GOOD
switch (x) {
    case 1:
        printf("One\n");
        break;
    case 2:
        printf("Two\n");
        break;
}
```

### 5. Complex Nested Conditions

```c
// ❌ BAD: Hard to read
if (a) {
    if (b) {
        if (c) {
            if (d) {
                // code
            }
        }
    }
}

// ✅ GOOD: Early returns or logical operators
if (!a || !b || !c || !d) {
    return;
}
// code
```

---

## Practical Examples

### Example 1: Calculator

```c
#include <stdio.h>

int main() {
    float num1, num2, result;
    char operator;
    
    printf("Enter expression (e.g., 5 + 3): ");
    scanf("%f %c %f", &num1, &operator, &num2);
    
    switch (operator) {
        case '+':
            result = num1 + num2;
            printf("%.2f + %.2f = %.2f\n", num1, num2, result);
            break;
        case '-':
            result = num1 - num2;
            printf("%.2f - %.2f = %.2f\n", num1, num2, result);
            break;
        case '*':
            result = num1 * num2;
            printf("%.2f * %.2f = %.2f\n", num1, num2, result);
            break;
        case '/':
            if (num2 != 0) {
                result = num1 / num2;
                printf("%.2f / %.2f = %.2f\n", num1, num2, result);
            } else {
                printf("Error: Division by zero\n");
            }
            break;
        default:
            printf("Error: Invalid operator\n");
    }
    
    return 0;
}
```

### Example 2: Leap Year Checker

```c
#include <stdio.h>

int main() {
    int year;
    
    printf("Enter a year: ");
    scanf("%d", &year);
    
    if ((year % 4 == 0 && year % 100 != 0) || (year % 400 == 0)) {
        printf("%d is a leap year\n", year);
    } else {
        printf("%d is not a leap year\n", year);
    }
    
    return 0;
}
```

### Example 3: Triangle Validator

```c
#include <stdio.h>

int main() {
    float a, b, c;
    
    printf("Enter three sides of triangle: ");
    scanf("%f %f %f", &a, &b, &c);
    
    if (a + b > c && b + c > a && a + c > b) {
        if (a == b && b == c) {
            printf("Equilateral triangle\n");
        } else if (a == b || b == c || a == c) {
            printf("Isosceles triangle\n");
        } else {
            printf("Scalene triangle\n");
        }
    } else {
        printf("Invalid triangle\n");
    }
    
    return 0;
}
```

### Example 4: BMI Calculator

```c
#include <stdio.h>

int main() {
    float weight, height, bmi;
    
    printf("Enter weight (kg): ");
    scanf("%f", &weight);
    
    printf("Enter height (m): ");
    scanf("%f", &height);
    
    bmi = weight / (height * height);
    
    printf("Your BMI: %.2f\n", bmi);
    
    if (bmi < 18.5) {
        printf("Category: Underweight\n");
    } else if (bmi >= 18.5 && bmi < 25) {
        printf("Category: Normal weight\n");
    } else if (bmi >= 25 && bmi < 30) {
        printf("Category: Overweight\n");
    } else {
        printf("Category: Obese\n");
    }
    
    return 0;
}
```

---

## if vs switch: When to Use

### Use if-else when:
- Testing ranges of values
- Complex conditions with logical operators
- Comparing non-constant values
- Different types of conditions

```c
if (age >= 18 && age <= 65) { }
if (x > y && z < 10) { }
if (strcmp(str1, str2) == 0) { }
```

### Use switch when:
- Testing a single variable against multiple constant values
- Menu-driven programs
- Simple equality checks
- Better readability for many cases

```c
switch (menuChoice) {
    case 1: // ...
    case 2: // ...
    case 3: // ...
}
```

---

## Summary Table

| Statement | Use Case | Example |
|-----------|----------|---------|
| **if** | Single condition | `if (x > 5) { }` |
| **if-else** | Two alternatives | `if (x > 5) { } else { }` |
| **if-else if-else** | Multiple conditions | `if (x>5) { } else if (x<0) { }` |
| **Nested if** | Conditions within conditions | `if (a) { if (b) { } }` |
| **switch** | Multiple constant values | `switch(x) { case 1: ... }` |
| **Ternary** | Simple if-else in one line | `max = (a>b) ? a : b;` |

---

## Best Practices

1. **Always use braces `{}`** even for single statements
2. **Avoid deep nesting** - use early returns or extract functions
3. **Use meaningful condition names** or add comments
4. **Don't compare floats with `==`** - use epsilon comparison
5. **Watch for assignment in conditions** (`=` vs `==`)
6. **Use switch for multiple constant comparisons**
7. **Always include `default` in switch statements**
8. **Don't forget `break` in switch cases**
9. **Keep conditions simple and readable**
10. **Use logical operators to combine related conditions**

```c
// ✅ GOOD
if (age >= 18) {
    printf("Adult\n");
}

// Better with descriptive variables
int isAdult = (age >= 18);
if (isAdult) {
    printf("Adult\n");
}
```

Conditional statements are the foundation of decision-making in C - master them to write intelligent, responsive programs!