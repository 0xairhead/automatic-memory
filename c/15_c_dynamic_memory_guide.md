# Dynamic Memory Allocation in C

## Table of Contents

- [What is Dynamic Memory Allocation?](#what-is-dynamic-memory-allocation)
- [Static vs Dynamic Memory](#static-vs-dynamic-memory)
- [Why Use Dynamic Memory?](#why-use-dynamic-memory)
- [Memory Layout in C Programs](#memory-layout-in-c-programs)
- [malloc - Memory Allocation](#malloc---memory-allocation)
- [calloc - Contiguous Allocation](#calloc---contiguous-allocation)
- [realloc - Resize Allocation](#realloc---resize-allocation)
- [free - Deallocate Memory](#free---deallocate-memory)
- [Memory Allocation Best Practices](#memory-allocation-best-practices)
- [Dynamic Arrays](#dynamic-arrays)
- [Dynamic Strings](#dynamic-strings)
- [Dynamic Structures](#dynamic-structures)
- [Dynamic 2D Arrays](#dynamic-2d-arrays)
- [Memory Alignment](#memory-alignment)
- [Memory Leaks](#memory-leaks)
- [Detecting Memory Leaks](#detecting-memory-leaks)
- [Common Memory Errors](#common-memory-errors)
- [Advanced Techniques](#advanced-techniques)
- [Practical Examples](#practical-examples)
- [Common Pitfalls](#common-pitfalls)
- [Best Practices](#best-practices)
- [Quick Reference](#quick-reference)

---

## What is Dynamic Memory Allocation?

**Dynamic Memory Allocation** is the process of allocating memory at runtime (while the program is executing) rather than at compile time. This memory comes from the **heap** and must be manually managed.

```
Static Allocation (Compile-time):
int arr[100];  // Size fixed at compile time

Dynamic Allocation (Runtime):
int *arr = malloc(size * sizeof(int));  // Size determined at runtime
```

**Key Point:** With dynamic allocation, you control when memory is allocated and when it's freed.

---

## Static vs Dynamic Memory

| Feature | Static Memory | Dynamic Memory |
|---------|--------------|----------------|
| **Allocated** | Compile time | Runtime |
| **Location** | Stack | Heap |
| **Size** | Fixed | Flexible |
| **Lifetime** | Automatic (scope-based) | Manual (until freed) |
| **Speed** | Faster | Slower |
| **Management** | Automatic | Manual |
| **Max Size** | Limited (stack size ~1-8MB) | Large (available RAM) |

### Example Comparison

```c
#include <stdio.h>
#include <stdlib.h>

int main() {
    // Static allocation
    int staticArr[5] = {1, 2, 3, 4, 5};
    
    // Dynamic allocation
    int *dynamicArr = (int*)malloc(5 * sizeof(int));
    
    for (int i = 0; i < 5; i++) {
        dynamicArr[i] = i + 1;
    }
    
    printf("Static: ");
    for (int i = 0; i < 5; i++) {
        printf("%d ", staticArr[i]);
    }
    printf("\n");
    
    printf("Dynamic: ");
    for (int i = 0; i < 5; i++) {
        printf("%d ", dynamicArr[i]);
    }
    printf("\n");
    
    // Must free dynamic memory
    free(dynamicArr);
    // Static memory freed automatically
    
    return 0;
}
```

**Output:**
```
Static: 1 2 3 4 5
Dynamic: 1 2 3 4 5
```

---

## Why Use Dynamic Memory?

1. **Unknown Size at Compile Time** - Size determined by user input
2. **Variable Memory Needs** - Grow or shrink as needed
3. **Large Data Structures** - Exceed stack size limits
4. **Flexible Lifetime** - Control when memory exists
5. **Data Structures** - Linked lists, trees, graphs
6. **Memory Efficiency** - Allocate only what you need

### Example: Size Unknown at Compile Time

```c
#include <stdio.h>
#include <stdlib.h>

int main() {
    int n;
    
    printf("How many numbers? ");
    scanf("%d", &n);
    
    // Can't do: int arr[n]; in older C standards
    // Dynamic allocation solves this
    int *arr = (int*)malloc(n * sizeof(int));
    
    if (arr == NULL) {
        printf("Memory allocation failed\n");
        return 1;
    }
    
    printf("Enter %d numbers:\n", n);
    for (int i = 0; i < n; i++) {
        scanf("%d", &arr[i]);
    }
    
    printf("You entered: ");
    for (int i = 0; i < n; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");
    
    free(arr);
    
    return 0;
}
```

---

## Memory Layout in C Programs

```
High Memory Addresses
    ↑
    |
    |  Command Line Arguments & Environment Variables
    |
    |  ═══════════════════════════════════════
    |  STACK (grows downward)
    |  - Local variables
    |  - Function parameters
    |  - Return addresses
    |  - Automatic deallocation
    ↓
    
    ↑
    |  ═══════════════════════════════════════
    |  HEAP (grows upward)
    |  - Dynamic allocation (malloc, calloc, realloc)
    |  - Manual management (free)
    |  - Large memory available
    ↓
    |
    |  ═══════════════════════════════════════
    |  BSS (Uninitialized Data)
    |  - Uninitialized global/static variables
    |  - Initialized to 0 automatically
    |
    |  ═══════════════════════════════════════
    |  DATA (Initialized Data)
    |  - Initialized global/static variables
    |  - Constants
    |
    |  ═══════════════════════════════════════
    |  TEXT (Code)
    |  - Program instructions
    |  - Read-only
    ↓
Low Memory Addresses
```

---

## malloc - Memory Allocation

**`malloc`** (Memory Allocation) allocates a block of memory of specified size and returns a pointer to the beginning.

### Syntax

```c
void* malloc(size_t size);
```

### Basic Usage

```c
#include <stdio.h>
#include <stdlib.h>

int main() {
    int *ptr;
    
    // Allocate memory for one integer
    ptr = (int*)malloc(sizeof(int));
    
    if (ptr == NULL) {
        printf("Memory allocation failed\n");
        return 1;
    }
    
    *ptr = 42;
    printf("Value: %d\n", *ptr);
    
    free(ptr);
    
    return 0;
}
```

**Output:**
```
Value: 42
```

### Allocating Arrays

```c
#include <stdio.h>
#include <stdlib.h>

int main() {
    int n = 5;
    int *arr;
    
    // Allocate memory for 5 integers
    arr = (int*)malloc(n * sizeof(int));
    
    if (arr == NULL) {
        printf("Memory allocation failed\n");
        return 1;
    }
    
    // Initialize
    for (int i = 0; i < n; i++) {
        arr[i] = (i + 1) * 10;
    }
    
    // Print
    printf("Array: ");
    for (int i = 0; i < n; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");
    
    free(arr);
    
    return 0;
}
```

**Output:**
```
Array: 10 20 30 40 50
```

### Important Notes About malloc

**⚠️ malloc does NOT initialize memory** - Contains garbage values

```c
#include <stdio.h>
#include <stdlib.h>

int main() {
    int *ptr = (int*)malloc(5 * sizeof(int));
    
    printf("Uninitialized values:\n");
    for (int i = 0; i < 5; i++) {
        printf("%d ", ptr[i]);  // Unpredictable values!
    }
    printf("\n");
    
    free(ptr);
    
    return 0;
}
```

**Output (unpredictable):**
```
Uninitialized values:
0 1249832 32767 0 1249616
```

### Return Value

- **Success:** Returns pointer to allocated memory
- **Failure:** Returns `NULL` (no memory available)

**✅ Always check for NULL:**

```c
int *ptr = (int*)malloc(size * sizeof(int));
if (ptr == NULL) {
    fprintf(stderr, "Memory allocation failed\n");
    exit(1);
}
```

---

## calloc - Contiguous Allocation

**`calloc`** (Contiguous Allocation) allocates memory for an array and **initializes all bytes to zero**.

### Syntax

```c
void* calloc(size_t num_elements, size_t element_size);
```

### Basic Usage

```c
#include <stdio.h>
#include <stdlib.h>

int main() {
    int n = 5;
    int *arr;
    
    // Allocate and initialize to 0
    arr = (int*)calloc(n, sizeof(int));
    
    if (arr == NULL) {
        printf("Memory allocation failed\n");
        return 1;
    }
    
    printf("Initialized values:\n");
    for (int i = 0; i < n; i++) {
        printf("%d ", arr[i]);  // All zeros
    }
    printf("\n");
    
    free(arr);
    
    return 0;
}
```

**Output:**
```
Initialized values:
0 0 0 0 0
```

### malloc vs calloc

```c
#include <stdio.h>
#include <stdlib.h>

int main() {
    int n = 5;
    
    // malloc - uninitialized
    int *arr1 = (int*)malloc(n * sizeof(int));
    printf("malloc: ");
    for (int i = 0; i < n; i++) {
        printf("%d ", arr1[i]);  // Garbage
    }
    printf("\n");
    
    // calloc - initialized to 0
    int *arr2 = (int*)calloc(n, sizeof(int));
    printf("calloc: ");
    for (int i = 0; i < n; i++) {
        printf("%d ", arr2[i]);  // All zeros
    }
    printf("\n");
    
    free(arr1);
    free(arr2);
    
    return 0;
}
```

**Output:**
```
malloc: 1249832 32767 0 1249616 0
calloc: 0 0 0 0 0
```

### When to Use calloc

✅ **Use calloc when:**
- You need zero-initialized memory
- Working with arrays of structures
- Clearing memory is important for security

✅ **Use malloc when:**
- Initialization overhead is unnecessary
- You'll immediately overwrite all values
- Performance is critical

---

## realloc - Resize Allocation

**`realloc`** (Reallocation) changes the size of previously allocated memory block.

### Syntax

```c
void* realloc(void* ptr, size_t new_size);
```

### Basic Usage

```c
#include <stdio.h>
#include <stdlib.h>

int main() {
    int *arr;
    int size = 3;
    
    // Initial allocation
    arr = (int*)malloc(size * sizeof(int));
    
    for (int i = 0; i < size; i++) {
        arr[i] = i + 1;
    }
    
    printf("Original: ");
    for (int i = 0; i < size; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");
    
    // Resize to 6 elements
    size = 6;
    arr = (int*)realloc(arr, size * sizeof(int));
    
    if (arr == NULL) {
        printf("Reallocation failed\n");
        return 1;
    }
    
    // Initialize new elements
    for (int i = 3; i < size; i++) {
        arr[i] = i + 1;
    }
    
    printf("After realloc: ");
    for (int i = 0; i < size; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");
    
    free(arr);
    
    return 0;
}
```

**Output:**
```
Original: 1 2 3
After realloc: 1 2 3 4 5 6
```

### How realloc Works

```
Case 1: Enough space at current location
[Old Data] [Free Space] → [Old Data][New Space]
(Extends in place)

Case 2: Not enough space at current location
[Old Data] [Used]
              ↓ Copy
      [Old Data][New Space] (New location)
      
Case 3: Shrinking
[Old Data][Extra] → [Old Data] (Truncates)
```

### Important realloc Notes

**⚠️ Always use temporary pointer:**

```c
// ❌ BAD: Lose original pointer if realloc fails
arr = realloc(arr, new_size);

// ✅ GOOD: Keep original pointer
int *temp = realloc(arr, new_size);
if (temp == NULL) {
    // arr still valid
    free(arr);
    return 1;
}
arr = temp;
```

**Special Cases:**

```c
// realloc with NULL is same as malloc
ptr = realloc(NULL, size);  // Same as malloc(size)

// realloc with size 0 is same as free
realloc(ptr, 0);  // Same as free(ptr)
```

### Dynamic Resizing Example

```c
#include <stdio.h>
#include <stdlib.h>

int main() {
    int *arr = NULL;
    int capacity = 2;
    int size = 0;
    
    arr = (int*)malloc(capacity * sizeof(int));
    
    // Add elements, resize when needed
    for (int i = 1; i <= 10; i++) {
        if (size >= capacity) {
            capacity *= 2;
            printf("Resizing to capacity %d\n", capacity);
            
            int *temp = (int*)realloc(arr, capacity * sizeof(int));
            if (temp == NULL) {
                free(arr);
                return 1;
            }
            arr = temp;
        }
        
        arr[size++] = i;
    }
    
    printf("Final array: ");
    for (int i = 0; i < size; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");
    
    free(arr);
    
    return 0;
}
```

**Output:**
```
Resizing to capacity 4
Resizing to capacity 8
Resizing to capacity 16
Final array: 1 2 3 4 5 6 7 8 9 10
```

---

## free - Deallocate Memory

**`free`** releases dynamically allocated memory back to the system.

### Syntax

```c
void free(void* ptr);
```

### Basic Usage

```c
#include <stdio.h>
#include <stdlib.h>

int main() {
    int *ptr = (int*)malloc(sizeof(int));
    
    if (ptr == NULL) {
        return 1;
    }
    
    *ptr = 42;
    printf("Value: %d\n", *ptr);
    
    // Free the memory
    free(ptr);
    
    // Good practice: set to NULL
    ptr = NULL;
    
    return 0;
}
```

### Important free Rules

**✅ DO:**
- Free every malloc/calloc/realloc
- Set pointer to NULL after freeing
- Free in reverse order of dependencies

**❌ DON'T:**
- Free the same pointer twice (double free)
- Use pointer after freeing (dangling pointer)
- Free stack-allocated memory
- Free NULL (safe but unnecessary)

### Free Examples

```c
#include <stdio.h>
#include <stdlib.h>

int main() {
    // ✅ GOOD: Free allocated memory
    int *ptr1 = (int*)malloc(sizeof(int));
    free(ptr1);
    ptr1 = NULL;
    
    // ✅ GOOD: Safe to free NULL
    int *ptr2 = NULL;
    free(ptr2);  // No operation
    
    // ❌ BAD: Double free
    // int *ptr3 = (int*)malloc(sizeof(int));
    // free(ptr3);
    // free(ptr3);  // CRASH!
    
    // ❌ BAD: Use after free
    // int *ptr4 = (int*)malloc(sizeof(int));
    // free(ptr4);
    // *ptr4 = 10;  // UNDEFINED BEHAVIOR!
    
    // ❌ BAD: Free stack memory
    // int x = 10;
    // free(&x);  // CRASH!
    
    return 0;
}
```

### Memory Leak Prevention

```c
#include <stdio.h>
#include <stdlib.h>

void leakyFunction() {
    int *ptr = (int*)malloc(sizeof(int));
    *ptr = 42;
    // ❌ Forgot to free - MEMORY LEAK!
}

void goodFunction() {
    int *ptr = (int*)malloc(sizeof(int));
    if (ptr == NULL) {
        return;
    }
    
    *ptr = 42;
    // Use ptr...
    
    free(ptr);  // ✅ No leak
}

int main() {
    // Call leakyFunction 1000 times = 1000 leaks!
    for (int i = 0; i < 1000; i++) {
        leakyFunction();
    }
    
    // Call goodFunction 1000 times = 0 leaks!
    for (int i = 0; i < 1000; i++) {
        goodFunction();
    }
    
    return 0;
}
```

---

## Memory Allocation Best Practices

### 1. Always Check Return Value

```c
// ❌ BAD: No check
int *ptr = (int*)malloc(size * sizeof(int));
*ptr = 10;  // May crash if malloc failed!

// ✅ GOOD: Check before use
int *ptr = (int*)malloc(size * sizeof(int));
if (ptr == NULL) {
    fprintf(stderr, "Memory allocation failed\n");
    return 1;
}
*ptr = 10;
```

### 2. Use sizeof for Portability

```c
// ❌ BAD: Assumes int is 4 bytes
ptr = (int*)malloc(n * 4);

// ✅ GOOD: Portable
ptr = (int*)malloc(n * sizeof(int));

// ✅ BETTER: Use variable type
int *ptr = (int*)malloc(n * sizeof(*ptr));
```

### 3. Initialize After Allocation

```c
// ✅ Option 1: Use calloc
int *arr = (int*)calloc(n, sizeof(int));

// ✅ Option 2: Manual initialization
int *arr = (int*)malloc(n * sizeof(int));
if (arr != NULL) {
    for (int i = 0; i < n; i++) {
        arr[i] = 0;
    }
}

// ✅ Option 3: memset
int *arr = (int*)malloc(n * sizeof(int));
if (arr != NULL) {
    memset(arr, 0, n * sizeof(int));
}
```

### 4. Set to NULL After Freeing

```c
// ✅ GOOD: Prevents use-after-free
free(ptr);
ptr = NULL;

// Now safe to check
if (ptr != NULL) {
    *ptr = 10;
}
```

### 5. Match Allocations with Frees

```c
// ✅ GOOD: Every allocation has a free
void processData() {
    int *data = (int*)malloc(100 * sizeof(int));
    if (data == NULL) {
        return;
    }
    
    // ... process data ...
    
    free(data);  // Always free before return
}
```

---

## Dynamic Arrays

### Resizable Array Implementation

```c
#include <stdio.h>
#include <stdlib.h>

typedef struct {
    int *data;
    int size;
    int capacity;
} DynamicArray;

DynamicArray* createArray(int initial_capacity) {
    DynamicArray *arr = (DynamicArray*)malloc(sizeof(DynamicArray));
    if (arr == NULL) return NULL;
    
    arr->data = (int*)malloc(initial_capacity * sizeof(int));
    if (arr->data == NULL) {
        free(arr);
        return NULL;
    }
    
    arr->size = 0;
    arr->capacity = initial_capacity;
    return arr;
}

int push(DynamicArray *arr, int value) {
    if (arr->size >= arr->capacity) {
        int new_capacity = arr->capacity * 2;
        int *temp = (int*)realloc(arr->data, new_capacity * sizeof(int));
        
        if (temp == NULL) {
            return 0;  // Resize failed
        }
        
        arr->data = temp;
        arr->capacity = new_capacity;
    }
    
    arr->data[arr->size++] = value;
    return 1;
}

void printArray(DynamicArray *arr) {
    printf("Array [size=%d, capacity=%d]: ", arr->size, arr->capacity);
    for (int i = 0; i < arr->size; i++) {
        printf("%d ", arr->data[i]);
    }
    printf("\n");
}

void freeArray(DynamicArray *arr) {
    if (arr != NULL) {
        free(arr->data);
        free(arr);
    }
}

int main() {
    DynamicArray *arr = createArray(2);
    
    for (int i = 1; i <= 10; i++) {
        push(arr, i);
        printArray(arr);
    }
    
    freeArray(arr);
    
    return 0;
}
```

**Output:**
```
Array [size=1, capacity=2]: 1 
Array [size=2, capacity=2]: 1 2 
Array [size=3, capacity=4]: 1 2 3 
Array [size=4, capacity=4]: 1 2 3 4 
Array [size=5, capacity=8]: 1 2 3 4 5 
Array [size=6, capacity=8]: 1 2 3 4 5 6 
Array [size=7, capacity=8]: 1 2 3 4 5 6 7 
Array [size=8, capacity=8]: 1 2 3 4 5 6 7 8 
Array [size=9, capacity=16]: 1 2 3 4 5 6 7 8 9 
Array [size=10, capacity=16]: 1 2 3 4 5 6 7 8 9 10
```

---

## Dynamic Strings

### String Manipulation with Dynamic Memory

```c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* createString(const char *str) {
    if (str == NULL) return NULL;
    
    char *newStr = (char*)malloc(strlen(str) + 1);
    if (newStr != NULL) {
        strcpy(newStr, str);
    }
    return newStr;
}

char* concatenate(const char *s1, const char *s2) {
    if (s1 == NULL || s2 == NULL) return NULL;
    
    size_t len = strlen(s1) + strlen(s2) + 1;
    char *result = (char*)malloc(len);
    
    if (result != NULL) {
        strcpy(result, s1);
        strcat(result, s2);
    }
    
    return result;
}

int main() {
    char *str1 = createString("Hello");
    char *str2 = createString(" World");
    char *combined = concatenate(str1, str2);
    
    if (combined != NULL) {
        printf("Combined: %s\n", combined);
    }
    
    free(str1);
    free(str2);
    free(combined);
    
    return 0;
}
```

**Output:**
```
Combined: Hello World
```

### Reading String of Unknown Length

```c
#include <stdio.h>
#include <stdlib.h>

char* readLine() {
    int capacity = 16;
    int size = 0;
    char *line = (char*)malloc(capacity);
    
    if (line == NULL) return NULL;
    
    int c;
    while ((c = getchar()) != '\n' && c != EOF) {
        if (size >= capacity - 1) {
            capacity *= 2;
            char *temp = (char*)realloc(line, capacity);
            if (temp == NULL) {
                free(line);
                return NULL;
            }
            line = temp;
        }
        line[size++] = c;
    }
    
    line[size] = '\0';
    return line;
}

int main() {
    printf("Enter a line: ");
    char *input = readLine();
    
    if (input != NULL) {
        printf("You entered: %s\n", input);
        printf("Length: %zu\n", strlen(input));
        free(input);
    }
    
    return 0;
}
```

---

## Dynamic Structures

### Dynamic Structure Allocation

```c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
    int id;
    char *name;
    float salary;
} Employee;

Employee* createEmployee(int id, const char *name, float salary) {
    Employee *emp = (Employee*)malloc(sizeof(Employee));
    if (emp == NULL) return NULL;
    
    emp->id = id;
    emp->salary = salary;
    
    emp->name = (char*)malloc(strlen(name) + 1);
    if (emp->name == NULL) {
        free(emp);
        return NULL;
    }
    strcpy(emp->name, name);
    
    return emp;
}

void printEmployee(Employee *emp) {
    if (emp != NULL) {
        printf("ID: %d, Name: %s, Salary: $%.2f\n",
               emp->id, emp->name, emp->salary);
    }
}

void freeEmployee(Employee *emp) {
    if (emp != NULL) {
        free(emp->name);  // Free string first
        free(emp);         // Then free structure
    }
}

int main() {
    Employee *emp1 = createEmployee(101, "John Doe", 50000.0);
    Employee *emp2 = createEmployee(102, "Jane Smith", 60000.0);
    
    printEmployee(emp1);
    printEmployee(emp2);
    
    freeEmployee(emp1);
    freeEmployee(emp2);
    
    return 0;
}
```

**Output:**
```
ID: 101, Name: John Doe, Salary: $50000.00
ID: 102, Name: Jane Smith, Salary: $60000.00
```

### Array of Dynamic Structures

```c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
    char *title;
    int year;
} Movie;

int main() {
    int n = 3;
    Movie *movies = (Movie*)calloc(n, sizeof(Movie));
    
    if (movies == NULL) {
        return 1;
    }
    
    // Initialize movies
    movies[0].title = strdup("The Matrix");
    movies[0].year = 1999;
    
    movies[1].title = strdup("Inception");
    movies[1].year = 2010;
    
    movies[2].title = strdup("Interstellar");
    movies[2].year = 2014;
    
    // Print
    for (int i = 0; i < n; i++) {
        printf("%s (%d)\n", movies[i].title, movies[i].year);
    }
    
    // Free
    for (int i = 0; i < n; i++) {
        free(movies[i].title);
    }
    free(movies);
    
    return 0;
}
```

**Output:**
```
The Matrix (1999)
Inception (2010)
Interstellar (2014)
```

---

## Dynamic 2D Arrays

### Method 1: Array of Pointers

```c
#include <stdio.h>
#include <stdlib.h>

int main() {
    int rows = 3, cols = 4;
    int **matrix;
    
    // Allocate array of row pointers
    matrix = (int**)malloc(rows * sizeof(int*));
    
    // Allocate each row
    for (int i = 0; i < rows; i++) {
        matrix[i] = (int*)malloc(cols * sizeof(int));
    }
    
    // Initialize
    int value = 1;
    for (int i = 0; i < rows; i++) {
        for (int j = 0; j < cols; j++) {
            matrix[i][j] = value++;
        }
    }
    
    // Print
    for (int i = 0; i < rows; i++) {
        for (int j = 0; j < cols; j++) {
            printf("%3d ", matrix[i][j]);
        }
        printf("\n");
    }
    
    // Free
    for (int i = 0; i < rows; i++) {
        free(matrix[i]);
    }
    free(matrix);
    
    return 0;
}
```

**Output:**
```
  1   2   3   4 
  5   6   7   8 
  9  10  11  12
```

### Method 2: Single Contiguous Block

```c
#include <stdio.h>
#include <stdlib.h>

int main() {
    int rows = 3, cols = 4;
    
    // Allocate as single block
    int *matrix = (int*)malloc(rows * cols * sizeof(int));
    
    if (matrix == NULL) {
        return 1;
    }
    
    // Initialize
    int value = 1;
    for (int i = 0; i < rows; i++) {
        for (int j = 0; j < cols; j++) {
            matrix[i * cols + j] = value++;
        }
    }
    
    // Print
    for (int i = 0; i < rows; i++) {
        for (int j = 0; j < cols; j++) {
            printf("%3d ", matrix[i * cols + j]);
        }
        printf("\n");
    }
    
    free(matrix);
    
    return 0;
}
```

### Method 3: Helper Pointer Array

```c
#include <stdio.h>
#include <stdlib.h>

int main() {
    int rows = 3, cols = 4;
    
    // Allocate data in contiguous block
    int *data = (int*)malloc(rows * cols * sizeof(int));
    
    // Create helper pointers for 2D access
    int **matrix = (int**)malloc(rows * sizeof(int*));
    for (int i = 0; i < rows; i++) {
        matrix[i] = data + i * cols;
    }
    
    // Initialize using 2D notation
    int value = 1;
    for (int i = 0; i < rows; i++) {
        for (int j = 0; j < cols; j++) {
            matrix[i][j] = value++;
        }
    }
    
    // Print
    for (int i = 0; i < rows; i++) {
        for (int j = 0; j < cols; j++) {
            printf("%3d ", matrix[i][j]);
        }
        printf("\n");
    }
    
    // Free
    free(data);
    free(matrix);
    
    return 0;
}
```

---

## Memory Alignment

Memory alignment affects performance and correctness.

### Understanding Alignment

```c
#include <stdio.h>
#include <stdlib.h>

int main() {
    printf("Alignment requirements:\n");
    printf("char: %zu byte\n", _Alignof(char));
    printf("short: %zu bytes\n", _Alignof(short));
    printf("int: %zu bytes\n", _Alignof(int));
    printf("long: %zu bytes\n", _Alignof(long));
    printf("float: %zu bytes\n", _Alignof(float));
    printf("double: %zu bytes\n", _Alignof(double));
    printf("pointer: %zu bytes\n", _Alignof(void*));
    
    return 0;
}
```

**Output (typical):**
```
Alignment requirements:
char: 1 byte
short: 2 bytes
int: 4 bytes
long: 8 bytes
float: 4 bytes
double: 8 bytes
pointer: 8 bytes
```

### Aligned Allocation (C11)

```c
#include <stdio.h>
#include <stdlib.h>

int main() {
    // Allocate with specific alignment
    void *ptr = aligned_alloc(64, 1024);  // 64-byte aligned
    
    if (ptr != NULL) {
        printf("Address: %p\n", ptr);
        printf("Aligned: %s\n", 
               ((uintptr_t)ptr % 64 == 0) ? "Yes" : "No");
        
        free(ptr);
    }
    
    return 0;
}
```

---

## Memory Leaks

A **memory leak** occurs when allocated memory is not freed, causing the program to consume more and more memory.

### Common Leak Scenarios

```c
#include <stdio.h>
#include <stdlib.h>

// ❌ Leak 1: Forgot to free
void leak1() {
    int *ptr = (int*)malloc(sizeof(int));
    *ptr = 42;
    // Forgot free(ptr);
}

// ❌ Leak 2: Lost pointer
void leak2() {
    int *ptr = (int*)malloc(sizeof(int));
    ptr = NULL;  // Lost reference to allocated memory!
}

// ❌ Leak 3: Overwriting pointer
void leak3() {
    int *ptr = (int*)malloc(sizeof(int));
    ptr = (int*)malloc(sizeof(int));  // Lost first allocation!
    free(ptr);
}

// ❌ Leak 4: Early return
void leak4(int condition) {
    int *ptr = (int*)malloc(sizeof(int));
    
    if (condition) {
        return;  // Forgot to free!
    }
    
    free(ptr);
}

// ✅ Fixed: Always free
void good() {
    int *ptr = (int*)malloc(sizeof(int));
    if (ptr == NULL) {
        return;
    }
    
    *ptr = 42;
    // ... use ptr ...
    
    free(ptr);
}

int main() {
    // Don't do this!
    // for (int i = 0; i < 1000000; i++) {
    //     leak1();  // Leaks memory 1 million times!
    // }
    
    return 0;
}
```

### Memory Leak Detection Tools

**Valgrind (Linux/Mac):**
```bash
gcc -g program.c -o program
valgrind --leak-check=full ./program
```

**Dr. Memory (Windows):**
```bash
drmemory -- program.exe
```

**AddressSanitizer (Compiler built-in):**
```bash
gcc -fsanitize=address -g program.c -o program
./program
```

---

## Detecting Memory Leaks

### Simple Memory Tracker

```c
#include <stdio.h>
#include <stdlib.h>

static int allocations = 0;
static int deallocations = 0;

void* tracked_malloc(size_t size) {
    void *ptr = malloc(size);
    if (ptr != NULL) {
        allocations++;
        printf("[ALLOC] %p (%zu bytes) - Total: %d\n", 
               ptr, size, allocations - deallocations);
    }
    return ptr;
}

void tracked_free(void *ptr) {
    if (ptr != NULL) {
        deallocations++;
        printf("[FREE] %p - Total: %d\n", 
               ptr, allocations - deallocations);
        free(ptr);
    }
}

void report() {
    printf("\n=== Memory Report ===\n");
    printf("Allocations: %d\n", allocations);
    printf("Deallocations: %d\n", deallocations);
    printf("Leaks: %d\n", allocations - deallocations);
}

int main() {
    int *a = (int*)tracked_malloc(sizeof(int));
    int *b = (int*)tracked_malloc(sizeof(int));
    int *c = (int*)tracked_malloc(sizeof(int));
    
    tracked_free(a);
    tracked_free(b);
    // Forgot to free c!
    
    report();
    
    return 0;
}
```

**Output:**
```
[ALLOC] 0x55555555a2a0 (4 bytes) - Total: 1
[ALLOC] 0x55555555a2c0 (4 bytes) - Total: 2
[ALLOC] 0x55555555a2e0 (4 bytes) - Total: 3
[FREE] 0x55555555a2a0 - Total: 2
[FREE] 0x55555555a2c0 - Total: 1

=== Memory Report ===
Allocations: 3
Deallocations: 2
Leaks: 1
```

---

## Common Memory Errors

### 1. Buffer Overflow

```c
// ❌ BAD: Writing beyond allocated memory
int *arr = (int*)malloc(5 * sizeof(int));
for (int i = 0; i <= 10; i++) {  // Oops! Should be < 5
    arr[i] = i;  // BUFFER OVERFLOW!
}

// ✅ GOOD: Stay within bounds
int *arr = (int*)malloc(5 * sizeof(int));
for (int i = 0; i < 5; i++) {
    arr[i] = i;
}
free(arr);
```

### 2. Use After Free

```c
// ❌ BAD: Using freed memory
int *ptr = (int*)malloc(sizeof(int));
*ptr = 42;
free(ptr);
printf("%d\n", *ptr);  // UNDEFINED BEHAVIOR!

// ✅ GOOD: Don't use after freeing
int *ptr = (int*)malloc(sizeof(int));
*ptr = 42;
printf("%d\n", *ptr);
free(ptr);
ptr = NULL;
```

### 3. Double Free

```c
// ❌ BAD: Freeing twice
int *ptr = (int*)malloc(sizeof(int));
free(ptr);
free(ptr);  // CRASH!

// ✅ GOOD: Set to NULL after freeing
int *ptr = (int*)malloc(sizeof(int));
free(ptr);
ptr = NULL;
free(ptr);  // Safe (no-op)
```

### 4. Uninitialized Memory

```c
// ❌ BAD: Using uninitialized memory
int *arr = (int*)malloc(5 * sizeof(int));
printf("%d\n", arr[0]);  // Garbage value!

// ✅ GOOD: Initialize before use
int *arr = (int*)calloc(5, sizeof(int));  // All zeros
// Or
int *arr = (int*)malloc(5 * sizeof(int));
for (int i = 0; i < 5; i++) {
    arr[i] = 0;
}
```

### 5. Memory Fragmentation

```c
// ❌ BAD: Many small allocations
for (int i = 0; i < 1000; i++) {
    int *ptr = (int*)malloc(sizeof(int));
    // ... use ptr ...
    free(ptr);
}

// ✅ BETTER: Single large allocation
int *arr = (int*)malloc(1000 * sizeof(int));
// ... use arr ...
free(arr);
```

---

## Advanced Techniques

### Custom Memory Pool

```c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define POOL_SIZE 1024

typedef struct {
    char memory[POOL_SIZE];
    size_t used;
} MemoryPool;

MemoryPool* createPool() {
    MemoryPool *pool = (MemoryPool*)malloc(sizeof(MemoryPool));
    if (pool != NULL) {
        pool->used = 0;
    }
    return pool;
}

void* poolAlloc(MemoryPool *pool, size_t size) {
    if (pool == NULL || pool->used + size > POOL_SIZE) {
        return NULL;
    }
    
    void *ptr = pool->memory + pool->used;
    pool->used += size;
    
    return ptr;
}

void resetPool(MemoryPool *pool) {
    if (pool != NULL) {
        pool->used = 0;
    }
}

void destroyPool(MemoryPool *pool) {
    free(pool);
}

int main() {
    MemoryPool *pool = createPool();
    
    int *a = (int*)poolAlloc(pool, sizeof(int));
    int *b = (int*)poolAlloc(pool, sizeof(int));
    
    *a = 10;
    *b = 20;
    
    printf("a = %d, b = %d\n", *a, *b);
    printf("Pool used: %zu/%d bytes\n", pool->used, POOL_SIZE);
    
    resetPool(pool);  // No individual frees needed!
    destroyPool(pool);
    
    return 0;
}
```

### Memory Allocator Wrapper

```c
#include <stdio.h>
#include <stdlib.h>

void* safe_malloc(size_t size, const char *file, int line) {
    void *ptr = malloc(size);
    if (ptr == NULL) {
        fprintf(stderr, "malloc failed at %s:%d\n", file, line);
        exit(1);
    }
    return ptr;
}

#define MALLOC(size) safe_malloc(size, __FILE__, __LINE__)

int main() {
    int *arr = (int*)MALLOC(5 * sizeof(int));
    
    for (int i = 0; i < 5; i++) {
        arr[i] = i;
    }
    
    free(arr);
    
    return 0;
}
```

---

## Practical Examples

### Example 1: Dynamic Stack

```c
#include <stdio.h>
#include <stdlib.h>

typedef struct {
    int *data;
    int top;
    int capacity;
} Stack;

Stack* createStack(int capacity) {
    Stack *stack = (Stack*)malloc(sizeof(Stack));
    if (stack == NULL) return NULL;
    
    stack->data = (int*)malloc(capacity * sizeof(int));
    if (stack->data == NULL) {
        free(stack);
        return NULL;
    }
    
    stack->top = -1;
    stack->capacity = capacity;
    return stack;
}

int push(Stack *s, int value) {
    if (s->top >= s->capacity - 1) {
        int new_cap = s->capacity * 2;
        int *temp = (int*)realloc(s->data, new_cap * sizeof(int));
        if (temp == NULL) return 0;
        
        s->data = temp;
        s->capacity = new_cap;
    }
    
    s->data[++s->top] = value;
    return 1;
}

int pop(Stack *s, int *value) {
    if (s->top < 0) return 0;
    
    *value = s->data[s->top--];
    return 1;
}

void freeStack(Stack *s) {
    if (s != NULL) {
        free(s->data);
        free(s);
    }
}

int main() {
    Stack *stack = createStack(2);
    
    push(stack, 10);
    push(stack, 20);
    push(stack, 30);
    
    int value;
    while (pop(stack, &value)) {
        printf("%d ", value);
    }
    printf("\n");
    
    freeStack(stack);
    
    return 0;
}
```

**Output:**
```
30 20 10
```

### Example 2: Dynamic Linked List

```c
#include <stdio.h>
#include <stdlib.h>

typedef struct Node {
    int data;
    struct Node *next;
} Node;

Node* createNode(int data) {
    Node *node = (Node*)malloc(sizeof(Node));
    if (node != NULL) {
        node->data = data;
        node->next = NULL;
    }
    return node;
}

void append(Node **head, int data) {
    Node *newNode = createNode(data);
    if (newNode == NULL) return;
    
    if (*head == NULL) {
        *head = newNode;
        return;
    }
    
    Node *current = *head;
    while (current->next != NULL) {
        current = current->next;
    }
    current->next = newNode;
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
    Node *list = NULL;
    
    append(&list, 10);
    append(&list, 20);
    append(&list, 30);
    
    printList(list);
    
    freeList(list);
    
    return 0;
}
```

**Output:**
```
10 -> 20 -> 30 -> NULL
```

### Example 3: Hash Table

```c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define TABLE_SIZE 10

typedef struct Entry {
    char *key;
    int value;
    struct Entry *next;
} Entry;

typedef struct {
    Entry *buckets[TABLE_SIZE];
} HashTable;

unsigned int hash(const char *key) {
    unsigned int hash = 0;
    while (*key) {
        hash = (hash << 5) + *key++;
    }
    return hash % TABLE_SIZE;
}

HashTable* createTable() {
    HashTable *table = (HashTable*)calloc(1, sizeof(HashTable));
    return table;
}

void insert(HashTable *table, const char *key, int value) {
    unsigned int index = hash(key);
    
    Entry *entry = (Entry*)malloc(sizeof(Entry));
    entry->key = strdup(key);
    entry->value = value;
    entry->next = table->buckets[index];
    table->buckets[index] = entry;
}

int search(HashTable *table, const char *key, int *value) {
    unsigned int index = hash(key);
    Entry *entry = table->buckets[index];
    
    while (entry != NULL) {
        if (strcmp(entry->key, key) == 0) {
            *value = entry->value;
            return 1;
        }
        entry = entry->next;
    }
    
    return 0;
}

void freeTable(HashTable *table) {
    for (int i = 0; i < TABLE_SIZE; i++) {
        Entry *entry = table->buckets[i];
        while (entry != NULL) {
            Entry *temp = entry;
            entry = entry->next;
            free(temp->key);
            free(temp);
        }
    }
    free(table);
}

int main() {
    HashTable *table = createTable();
    
    insert(table, "apple", 5);
    insert(table, "banana", 3);
    insert(table, "orange", 7);
    
    int value;
    if (search(table, "banana", &value)) {
        printf("banana: %d\n", value);
    }
    
    freeTable(table);
    
    return 0;
}
```

**Output:**
```
banana: 3
```

---

## Common Pitfalls

### 1. Not Checking malloc Return

```c
// ❌ BAD
int *ptr = (int*)malloc(sizeof(int));
*ptr = 10;  // May crash if malloc failed

// ✅ GOOD
int *ptr = (int*)malloc(sizeof(int));
if (ptr == NULL) {
    fprintf(stderr, "Allocation failed\n");
    return 1;
}
*ptr = 10;
```

### 2. Assuming malloc Initializes

```c
// ❌ BAD: Using uninitialized memory
int *arr = (int*)malloc(5 * sizeof(int));
printf("%d\n", arr[0]);  // Garbage!

// ✅ GOOD: Use calloc or initialize
int *arr = (int*)calloc(5, sizeof(int));
```

### 3. Integer Overflow in Allocation

```c
// ❌ BAD: Potential overflow
size_t n = 1000000000;
int *arr = (int*)malloc(n * sizeof(int));  // May overflow!

// ✅ GOOD: Check for overflow
size_t n = 1000000000;
if (n > SIZE_MAX / sizeof(int)) {
    fprintf(stderr, "Size too large\n");
    return 1;
}
int *arr = (int*)malloc(n * sizeof(int));
```

### 4. Freeing Non-Heap Memory

```c
// ❌ BAD
int x = 10;
free(&x);  // CRASH! x is on stack

// ✅ GOOD: Only free heap memory
int *ptr = (int*)malloc(sizeof(int));
free(ptr);
```

### 5. Partial Structure Cleanup

```c
typedef struct {
    char *name;
    int *data;
} Record;

// ❌ BAD: Memory leak
void cleanup(Record *r) {
    free(r);  // Forgot to free r->name and r->data!
}

// ✅ GOOD: Free all members
void cleanup(Record *r) {
    if (r != NULL) {
        free(r->name);
        free(r->data);
        free(r);
    }
}
```

---

## Best Practices

1. **Always check malloc/calloc/realloc return values**
2. **Free every allocation** - one malloc = one free
3. **Set pointers to NULL after freeing**
4. **Use sizeof for portable code**
5. **Initialize allocated memory**
6. **Free in reverse order of dependencies**
7. **Use valgrind or similar tools** to detect leaks
8. **Prefer calloc when zero-initialization is needed**
9. **Keep allocation and deallocation close together**
10. **Document ownership** - who is responsible for freeing?

```c
// ✅ COMPREHENSIVE EXAMPLE
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
    int *data;
    size_t size;
    size_t capacity;
} DynamicArray;

DynamicArray* create(size_t initial_capacity) {
    // Validate input
    if (initial_capacity == 0) {
        initial_capacity = 10;
    }
    
    // Allocate structure
    DynamicArray *arr = (DynamicArray*)malloc(sizeof(DynamicArray));
    if (arr == NULL) {
        return NULL;
    }
    
    // Allocate data
    arr->data = (int*)calloc(initial_capacity, sizeof(int));
    if (arr->data == NULL) {
        free(arr);  // Clean up partial allocation
        return NULL;
    }
    
    arr->size = 0;
    arr->capacity = initial_capacity;
    
    return arr;
}

int resize(DynamicArray *arr, size_t new_capacity) {
    if (arr == NULL || new_capacity < arr->size) {
        return 0;
    }
    
    int *temp = (int*)realloc(arr->data, new_capacity * sizeof(int));
    if (temp == NULL) {
        return 0;  // Original array still valid
    }
    
    arr->data = temp;
    arr->capacity = new_capacity;
    
    return 1;
}

void destroy(DynamicArray *arr) {
    if (arr != NULL) {
        free(arr->data);  // Free member first
        free(arr);         // Then structure
    }
}

int main() {
    DynamicArray *arr = create(5);
    
    if (arr == NULL) {
        fprintf(stderr, "Creation failed\n");
        return 1;
    }
    
    // Use array...
    
    destroy(arr);  // Always cleanup
    arr = NULL;     // Prevent use-after-free
    
    return 0;
}
```

---

## Quick Reference

```c
// Allocation functions
void* malloc(size_t size);              // Allocate uninitialized
void* calloc(size_t n, size_t size);    // Allocate and zero
void* realloc(void *ptr, size_t size);  // Resize allocation
void free(void *ptr);                   // Deallocate

// Common patterns
int *ptr = (int*)malloc(n * sizeof(int));
int *ptr = (int*)calloc(n, sizeof(int));
ptr = (int*)realloc(ptr, new_size * sizeof(int));
free(ptr);
ptr = NULL;

// Always check
if (ptr == NULL) {
    // Handle error
}

// Dynamic array
int *arr = (int*)malloc(n * sizeof(int));
// Use arr[i]
free(arr);

// Dynamic structure
MyStruct *s = (MyStruct*)malloc(sizeof(MyStruct));
// Use s->member
free(s);

// 2D array
int **matrix = (int**)malloc(rows * sizeof(int*));
for (int i = 0; i < rows; i++) {
    matrix[i] = (int*)malloc(cols * sizeof(int));
}
// Use matrix[i][j]
for (int i = 0; i < rows; i++) {
    free(matrix[i]);
}
free(matrix);

// String
char *str = (char*)malloc(len + 1);
strcpy(str, "text");
free(str);

// Key rules
// 1. Check every allocation
// 2. Match every malloc with free
// 3. Set to NULL after freeing
// 4. Don't use after freeing
// 5. Don't free twice
// 6. Free in reverse dependency order
```

Dynamic memory allocation is powerful but requires discipline - always pair allocations with deallocations to avoid leaks!
