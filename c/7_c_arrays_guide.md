# Arrays in C

## Table of Contents

- [What is an Array?](#what-is-an-array)
- [Array Declaration and Initialization](#array-declaration-and-initialization)
- [Accessing Array Elements](#accessing-array-elements)
- [Array Size and Length](#array-size-and-length)
- [Iterating Through Arrays](#iterating-through-arrays)
- [Common Array Operations](#common-array-operations)
- [Multi-Dimensional Arrays](#multi-dimensional-arrays)
- [Character Arrays (Strings)](#character-arrays-strings)
- [Passing Arrays to Functions](#passing-arrays-to-functions)
- [Arrays and Pointers](#arrays-and-pointers)
- [Dynamic Arrays (Using malloc)](#dynamic-arrays-using-malloc)
- [Common Array Pitfalls](#common-array-pitfalls)
- [Practical Examples](#practical-examples)
- [Summary](#summary)
- [Best Practices](#best-practices)

---

## What is an Array?

An **array** is a collection of elements of the **same data type** stored in **contiguous memory locations**.

```
Array: [10] [20] [30] [40] [50]
Index:  0    1    2    3    4
```

---

## Array Declaration and Initialization

### Declaration

```c
// Syntax: datatype arrayName[size];
int numbers[5];           // Array of 5 integers
float prices[10];         // Array of 10 floats
char letters[26];         // Array of 26 characters
```

### Initialization

#### Method 1: Initialize All Elements

```c
int numbers[5] = {10, 20, 30, 40, 50};
float prices[3] = {19.99, 29.99, 39.99};
char vowels[5] = {'a', 'e', 'i', 'o', 'u'};
```

#### Method 2: Partial Initialization

```c
int numbers[5] = {10, 20};  // {10, 20, 0, 0, 0}
// Remaining elements are initialized to 0
```

#### Method 3: Initialize All to Zero

```c
int numbers[5] = {0};  // All elements = 0
int zeros[100] = {0};  // All 100 elements = 0
```

#### Method 4: Omit Size (Compiler Determines)

```c
int numbers[] = {10, 20, 30, 40, 50};  // Size = 5
char name[] = {'J', 'o', 'h', 'n', '\0'};  // Size = 5
```

#### Method 5: Designated Initializers (C99)

```c
int numbers[10] = {[0] = 1, [4] = 5, [9] = 10};
// Result: {1, 0, 0, 0, 5, 0, 0, 0, 0, 10}
```

---

## Accessing Array Elements

Arrays use **zero-based indexing** (first element is at index 0).

```c
#include <stdio.h>

int main() {
    int numbers[5] = {10, 20, 30, 40, 50};
    
    // Access elements
    printf("%d\n", numbers[0]);  // 10 (first element)
    printf("%d\n", numbers[2]);  // 30
    printf("%d\n", numbers[4]);  // 50 (last element)
    
    // Modify elements
    numbers[1] = 25;
    printf("%d\n", numbers[1]);  // 25
    
    return 0;
}
```

---

## Array Size and Length

### Get Array Size

```c
#include <stdio.h>

int main() {
    int numbers[] = {10, 20, 30, 40, 50};
    
    // Calculate number of elements
    int length = sizeof(numbers) / sizeof(numbers[0]);
    
    printf("Array size in bytes: %zu\n", sizeof(numbers));  // 20
    printf("Element size in bytes: %zu\n", sizeof(numbers[0]));  // 4
    printf("Number of elements: %d\n", length);  // 5
    
    return 0;
}
```

---

## Iterating Through Arrays

### Using for Loop

```c
#include <stdio.h>

int main() {
    int numbers[5] = {10, 20, 30, 40, 50};
    int length = sizeof(numbers) / sizeof(numbers[0]);
    
    // Print all elements
    for (int i = 0; i < length; i++) {
        printf("numbers[%d] = %d\n", i, numbers[i]);
    }
    
    return 0;
}
```

**Output:**
```
numbers[0] = 10
numbers[1] = 20
numbers[2] = 30
numbers[3] = 40
numbers[4] = 50
```

### Using while Loop

```c
int i = 0;
while (i < length) {
    printf("%d ", numbers[i]);
    i++;
}
```

---

## Common Array Operations

### 1. Sum of Array Elements

```c
#include <stdio.h>

int main() {
    int numbers[] = {10, 20, 30, 40, 50};
    int length = sizeof(numbers) / sizeof(numbers[0]);
    int sum = 0;
    
    for (int i = 0; i < length; i++) {
        sum += numbers[i];
    }
    
    printf("Sum: %d\n", sum);  // 150
    
    return 0;
}
```

### 2. Find Maximum Element

```c
#include <stdio.h>

int main() {
    int numbers[] = {45, 23, 78, 12, 67};
    int length = sizeof(numbers) / sizeof(numbers[0]);
    int max = numbers[0];
    
    for (int i = 1; i < length; i++) {
        if (numbers[i] > max) {
            max = numbers[i];
        }
    }
    
    printf("Maximum: %d\n", max);  // 78
    
    return 0;
}
```

### 3. Find Minimum Element

```c
#include <stdio.h>

int main() {
    int numbers[] = {45, 23, 78, 12, 67};
    int length = sizeof(numbers) / sizeof(numbers[0]);
    int min = numbers[0];
    
    for (int i = 1; i < length; i++) {
        if (numbers[i] < min) {
            min = numbers[i];
        }
    }
    
    printf("Minimum: %d\n", min);  // 12
    
    return 0;
}
```

### 4. Calculate Average

```c
#include <stdio.h>

int main() {
    int numbers[] = {10, 20, 30, 40, 50};
    int length = sizeof(numbers) / sizeof(numbers[0]);
    int sum = 0;
    
    for (int i = 0; i < length; i++) {
        sum += numbers[i];
    }
    
    float average = (float)sum / length;
    printf("Average: %.2f\n", average);  // 30.00
    
    return 0;
}
```

### 5. Search for Element (Linear Search)

```c
#include <stdio.h>

int main() {
    int numbers[] = {10, 20, 30, 40, 50};
    int length = sizeof(numbers) / sizeof(numbers[0]);
    int search = 30;
    int found = -1;
    
    for (int i = 0; i < length; i++) {
        if (numbers[i] == search) {
            found = i;
            break;
        }
    }
    
    if (found != -1) {
        printf("%d found at index %d\n", search, found);
    } else {
        printf("%d not found\n", search);
    }
    
    return 0;
}
```

### 6. Reverse an Array

```c
#include <stdio.h>

int main() {
    int numbers[] = {10, 20, 30, 40, 50};
    int length = sizeof(numbers) / sizeof(numbers[0]);
    
    // Reverse the array
    for (int i = 0; i < length / 2; i++) {
        int temp = numbers[i];
        numbers[i] = numbers[length - 1 - i];
        numbers[length - 1 - i] = temp;
    }
    
    // Print reversed array
    for (int i = 0; i < length; i++) {
        printf("%d ", numbers[i]);
    }
    // Output: 50 40 30 20 10
    
    return 0;
}
```

### 7. Copy an Array

```c
#include <stdio.h>

int main() {
    int source[] = {10, 20, 30, 40, 50};
    int length = sizeof(source) / sizeof(source[0]);
    int destination[5];
    
    // Copy elements
    for (int i = 0; i < length; i++) {
        destination[i] = source[i];
    }
    
    // Print destination
    for (int i = 0; i < length; i++) {
        printf("%d ", destination[i]);
    }
    // Output: 10 20 30 40 50
    
    return 0;
}
```

---

## Multi-Dimensional Arrays

### 2D Arrays (Matrices)

```c
#include <stdio.h>

int main() {
    // Declaration and initialization
    int matrix[3][4] = {
        {1, 2, 3, 4},
        {5, 6, 7, 8},
        {9, 10, 11, 12}
    };
    
    // Access elements
    printf("%d\n", matrix[0][0]);  // 1
    printf("%d\n", matrix[1][2]);  // 7
    printf("%d\n", matrix[2][3]);  // 12
    
    // Print entire matrix
    for (int i = 0; i < 3; i++) {
        for (int j = 0; j < 4; j++) {
            printf("%3d ", matrix[i][j]);
        }
        printf("\n");
    }
    
    return 0;
}
```

**Output:**
```
  1   2   3   4 
  5   6   7   8 
  9  10  11  12 
```

### 3D Arrays

```c
int cube[2][3][4];  // 2 layers, 3 rows, 4 columns

// Initialize
int cube[2][2][2] = {
    {{1, 2}, {3, 4}},
    {{5, 6}, {7, 8}}
};

// Access
printf("%d\n", cube[1][1][0]);  // 7
```

---

## Character Arrays (Strings)

### String Declaration and Initialization

```c
#include <stdio.h>
#include <string.h>

int main() {
    // Method 1: Character array
    char name1[6] = {'H', 'e', 'l', 'l', 'o', '\0'};
    
    // Method 2: String literal
    char name2[] = "Hello";  // Automatically adds '\0'
    
    // Method 3: Fixed size
    char name3[20] = "Hello";  // Remaining chars are '\0'
    
    printf("%s\n", name1);  // Hello
    printf("%s\n", name2);  // Hello
    printf("%s\n", name3);  // Hello
    
    return 0;
}
```

### String Operations

```c
#include <stdio.h>
#include <string.h>

int main() {
    char str1[20] = "Hello";
    char str2[20] = "World";
    char str3[20];
    
    // String length
    printf("Length: %lu\n", strlen(str1));  // 5
    
    // String copy
    strcpy(str3, str1);
    printf("Copy: %s\n", str3);  // Hello
    
    // String concatenation
    strcat(str1, " ");
    strcat(str1, str2);
    printf("Concat: %s\n", str1);  // Hello World
    
    // String comparison
    if (strcmp(str2, "World") == 0) {
        printf("Strings are equal\n");
    }
    
    return 0;
}
```

---

## Passing Arrays to Functions

### 1D Array as Parameter

```c
#include <stdio.h>

void printArray(int arr[], int size) {
    for (int i = 0; i < size; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");
}

int sumArray(int arr[], int size) {
    int sum = 0;
    for (int i = 0; i < size; i++) {
        sum += arr[i];
    }
    return sum;
}

void modifyArray(int arr[], int size) {
    for (int i = 0; i < size; i++) {
        arr[i] *= 2;  // Double each element
    }
}

int main() {
    int numbers[] = {10, 20, 30, 40, 50};
    int size = sizeof(numbers) / sizeof(numbers[0]);
    
    printArray(numbers, size);  // 10 20 30 40 50
    
    int total = sumArray(numbers, size);
    printf("Sum: %d\n", total);  // 150
    
    modifyArray(numbers, size);
    printArray(numbers, size);  // 20 40 60 80 100
    
    return 0;
}
```

### 2D Array as Parameter

```c
#include <stdio.h>

void print2DArray(int arr[][4], int rows) {
    for (int i = 0; i < rows; i++) {
        for (int j = 0; j < 4; j++) {
            printf("%3d ", arr[i][j]);
        }
        printf("\n");
    }
}

int main() {
    int matrix[3][4] = {
        {1, 2, 3, 4},
        {5, 6, 7, 8},
        {9, 10, 11, 12}
    };
    
    print2DArray(matrix, 3);
    
    return 0;
}
```

---

## Arrays and Pointers

Arrays and pointers are closely related in C.

```c
#include <stdio.h>

int main() {
    int numbers[] = {10, 20, 30, 40, 50};
    
    // Array name is a pointer to first element
    printf("numbers = %p\n", numbers);
    printf("&numbers[0] = %p\n", &numbers[0]);
    
    // Access using pointer
    printf("*numbers = %d\n", *numbers);  // 10
    printf("*(numbers + 1) = %d\n", *(numbers + 1));  // 20
    printf("*(numbers + 2) = %d\n", *(numbers + 2));  // 30
    
    // Pointer arithmetic
    int *ptr = numbers;
    for (int i = 0; i < 5; i++) {
        printf("%d ", *(ptr + i));
    }
    printf("\n");
    
    return 0;
}
```

---

## Dynamic Arrays (Using malloc)

```c
#include <stdio.h>
#include <stdlib.h>

int main() {
    int n;
    printf("Enter array size: ");
    scanf("%d", &n);
    
    // Allocate memory dynamically
    int *arr = (int*)malloc(n * sizeof(int));
    
    if (arr == NULL) {
        printf("Memory allocation failed\n");
        return 1;
    }
    
    // Input elements
    printf("Enter %d elements:\n", n);
    for (int i = 0; i < n; i++) {
        scanf("%d", &arr[i]);
    }
    
    // Print elements
    printf("Array elements: ");
    for (int i = 0; i < n; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");
    
    // Free memory
    free(arr);
    
    return 0;
}
```

---

## Common Array Pitfalls

### 1. Array Index Out of Bounds

```c
int numbers[5] = {10, 20, 30, 40, 50};

// ❌ ERROR: Undefined behavior
numbers[5] = 60;  // Valid indices are 0-4
numbers[10] = 100;  // Way out of bounds!

// ✅ CORRECT
numbers[4] = 60;  // Last valid index
```

### 2. Uninitialized Arrays

```c
int numbers[5];  // ❌ Contains garbage values

// ✅ CORRECT: Initialize
int numbers[5] = {0};  // All zeros
```

### 3. Array Size in Functions

```c
void func(int arr[]) {
    // ❌ This won't work as expected
    int size = sizeof(arr) / sizeof(arr[0]);  // Always returns 1 or 2
}

// ✅ CORRECT: Pass size as parameter
void func(int arr[], int size) {
    // Use size parameter
}
```

### 4. Returning Arrays from Functions

```c
// ❌ WRONG: Returns pointer to local array
int* getArray() {
    int arr[5] = {1, 2, 3, 4, 5};
    return arr;  // Undefined behavior!
}

// ✅ CORRECT: Use static or dynamic allocation
int* getArray() {
    static int arr[5] = {1, 2, 3, 4, 5};
    return arr;
}

// Or use dynamic allocation
int* getArray() {
    int *arr = (int*)malloc(5 * sizeof(int));
    arr[0] = 1; arr[1] = 2; // ... initialize
    return arr;  // Caller must free()
}
```

---

## Practical Examples

### Example 1: Grade Calculator

```c
#include <stdio.h>

int main() {
    float grades[5];
    float sum = 0, average;
    
    // Input grades
    printf("Enter 5 grades:\n");
    for (int i = 0; i < 5; i++) {
        printf("Grade %d: ", i + 1);
        scanf("%f", &grades[i]);
        sum += grades[i];
    }
    
    // Calculate average
    average = sum / 5;
    
    // Display results
    printf("\nGrades: ");
    for (int i = 0; i < 5; i++) {
        printf("%.2f ", grades[i]);
    }
    printf("\nAverage: %.2f\n", average);
    
    // Determine letter grade
    if (average >= 90) printf("Grade: A\n");
    else if (average >= 80) printf("Grade: B\n");
    else if (average >= 70) printf("Grade: C\n");
    else if (average >= 60) printf("Grade: D\n");
    else printf("Grade: F\n");
    
    return 0;
}
```

### Example 2: Bubble Sort

```c
#include <stdio.h>

void bubbleSort(int arr[], int n) {
    for (int i = 0; i < n - 1; i++) {
        for (int j = 0; j < n - i - 1; j++) {
            if (arr[j] > arr[j + 1]) {
                // Swap
                int temp = arr[j];
                arr[j] = arr[j + 1];
                arr[j + 1] = temp;
            }
        }
    }
}

int main() {
    int numbers[] = {64, 34, 25, 12, 22, 11, 90};
    int n = sizeof(numbers) / sizeof(numbers[0]);
    
    printf("Original array: ");
    for (int i = 0; i < n; i++) {
        printf("%d ", numbers[i]);
    }
    
    bubbleSort(numbers, n);
    
    printf("\nSorted array: ");
    for (int i = 0; i < n; i++) {
        printf("%d ", numbers[i]);
    }
    printf("\n");
    
    return 0;
}
```

### Example 3: Matrix Addition

```c
#include <stdio.h>

int main() {
    int a[2][2] = {{1, 2}, {3, 4}};
    int b[2][2] = {{5, 6}, {7, 8}};
    int sum[2][2];
    
    // Add matrices
    for (int i = 0; i < 2; i++) {
        for (int j = 0; j < 2; j++) {
            sum[i][j] = a[i][j] + b[i][j];
        }
    }
    
    // Display result
    printf("Matrix A:\n");
    for (int i = 0; i < 2; i++) {
        for (int j = 0; j < 2; j++) {
            printf("%d ", a[i][j]);
        }
        printf("\n");
    }
    
    printf("\nMatrix B:\n");
    for (int i = 0; i < 2; i++) {
        for (int j = 0; j < 2; j++) {
            printf("%d ", b[i][j]);
        }
        printf("\n");
    }
    
    printf("\nSum:\n");
    for (int i = 0; i < 2; i++) {
        for (int j = 0; j < 2; j++) {
            printf("%d ", sum[i][j]);
        }
        printf("\n");
    }
    
    return 0;
}
```

---

## Summary

| Concept | Syntax | Example |
|---------|--------|---------|
| **Declaration** | `type name[size];` | `int arr[5];` |
| **Initialization** | `type name[] = {...};` | `int arr[] = {1,2,3};` |
| **Access** | `name[index]` | `arr[0]` |
| **Length** | `sizeof(arr)/sizeof(arr[0])` | `5` |
| **2D Array** | `type name[rows][cols];` | `int mat[3][4];` |
| **String** | `char name[];` | `char str[] = "Hi";` |
| **Pass to Function** | `void func(int arr[], int size)` | Function parameter |
| **Dynamic** | `malloc()` | `int *arr = malloc(...)` |

---

## Best Practices

1. **Always initialize arrays** to avoid garbage values
2. **Check array bounds** before accessing
3. **Pass array size** to functions as a parameter
4. **Use meaningful names** for arrays
5. **Free dynamically allocated arrays** with `free()`
6. **Use `const`** for read-only arrays in functions
7. **Prefer stack arrays** for small, fixed-size data
8. **Use dynamic arrays** for large or variable-size data

```c
// ✅ GOOD
int grades[5] = {0};  // Initialized
void printArray(const int arr[], int size);  // const for read-only

// ❌ BAD
int grades[5];  // Uninitialized
void printArray(int arr[]);  // No size parameter
```

Arrays are fundamental to C programming - master them to handle collections of data efficiently!