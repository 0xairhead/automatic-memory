# Strings in C

## Table of Contents

- [What is a String in C?](#what-is-a-string-in-c)
- [String Declaration and Initialization](#string-declaration-and-initialization)
- [String Input and Output](#string-input-and-output)
- [String Length](#string-length)
- [String Library Functions (string.h)](#string-library-functions-stringh)
- [Common String Operations](#common-string-operations)
- [String Conversion Functions](#string-conversion-functions)
- [Array of Strings (2D Character Array)](#array-of-strings-2d-character-array)
- [Dynamic String Allocation](#dynamic-string-allocation)
- [Common String Pitfalls](#common-string-pitfalls)
- [Practical Examples](#practical-examples)
- [String Functions Summary](#string-functions-summary)
- [Best Practices](#best-practices)

---

## What is a String in C?

A **string** in C is an array of characters terminated by a **null character** `'\0'`.

```
String: "Hello"
Memory: ['H']['e']['l']['l']['o']['\0']
Index:   0    1    2    3    4    5
```

**Important:** C does NOT have a built-in string data type like other languages. Strings are character arrays.

---

## String Declaration and Initialization

### Method 1: Character Array

```c
char str[6] = {'H', 'e', 'l', 'l', 'o', '\0'};
```

### Method 2: String Literal

```c
char str[] = "Hello";  // Automatically adds '\0'
// Size is 6 (5 characters + '\0')
```

### Method 3: Fixed Size Array

```c
char str[20] = "Hello";  // Size 20, only first 6 used
```

### Method 4: Pointer to String Literal

```c
char *str = "Hello";  // Points to read-only memory
// ⚠️ Cannot modify this string!
```

### Method 5: Without Initialization

```c
char str[50];  // Uninitialized, contains garbage
```

---

## String Input and Output

### Output: printf()

```c
#include <stdio.h>

int main() {
    char name[] = "Alice";
    
    printf("%s\n", name);           // Alice
    printf("Hello, %s!\n", name);   // Hello, Alice!
    printf("Name: %10s\n", name);   // Right-aligned in 10 chars
    printf("Name: %-10s|\n", name); // Left-aligned in 10 chars
    
    return 0;
}
```

### Output: puts()

```c
#include <stdio.h>

int main() {
    char message[] = "Hello, World!";
    
    puts(message);  // Prints string + newline
    // Same as: printf("%s\n", message);
    
    return 0;
}
```

### Input: scanf()

```c
#include <stdio.h>

int main() {
    char name[50];
    
    printf("Enter your name: ");
    scanf("%s", name);  // Reads until whitespace
    
    printf("Hello, %s!\n", name);
    
    return 0;
}
```

**⚠️ Problem:** `scanf()` stops at whitespace!

```
Input: "John Doe"
Reads: "John"  (stops at space)
```

### Input: gets() ❌ NEVER USE

```c
// ❌ DANGEROUS: Buffer overflow risk
gets(name);  // Deprecated and unsafe!
```

### Input: fgets() ✅ RECOMMENDED

```c
#include <stdio.h>

int main() {
    char name[50];
    
    printf("Enter your full name: ");
    fgets(name, sizeof(name), stdin);  // Safe, reads entire line
    
    printf("Hello, %s", name);  // fgets includes '\n'
    
    return 0;
}
```

**Remove newline from fgets:**

```c
#include <stdio.h>
#include <string.h>

int main() {
    char name[50];
    
    printf("Enter your name: ");
    fgets(name, sizeof(name), stdin);
    
    // Remove trailing newline
    name[strcspn(name, "\n")] = '\0';
    
    printf("Hello, %s!\n", name);
    
    return 0;
}
```

---

## String Length

### Manual Calculation

```c
#include <stdio.h>

int main() {
    char str[] = "Hello";
    int length = 0;
    
    while (str[length] != '\0') {
        length++;
    }
    
    printf("Length: %d\n", length);  // 5
    
    return 0;
}
```

### Using strlen()

```c
#include <stdio.h>
#include <string.h>

int main() {
    char str[] = "Hello, World!";
    
    printf("Length: %lu\n", strlen(str));  // 13
    
    // Note: strlen() doesn't count '\0'
    printf("Array size: %lu\n", sizeof(str));  // 14
    
    return 0;
}
```

---

## String Library Functions (string.h)

### 1. strcpy() - String Copy

```c
#include <stdio.h>
#include <string.h>

int main() {
    char source[] = "Hello";
    char destination[20];
    
    strcpy(destination, source);
    
    printf("Source: %s\n", source);           // Hello
    printf("Destination: %s\n", destination); // Hello
    
    return 0;
}
```

**Safe version: strncpy()**

```c
char dest[10];
strncpy(dest, source, sizeof(dest) - 1);
dest[sizeof(dest) - 1] = '\0';  // Ensure null termination
```

### 2. strcat() - String Concatenation

```c
#include <stdio.h>
#include <string.h>

int main() {
    char str1[50] = "Hello";
    char str2[] = " World";
    
    strcat(str1, str2);  // Appends str2 to str1
    
    printf("%s\n", str1);  // Hello World
    
    return 0;
}
```

**Safe version: strncat()**

```c
strncat(str1, str2, sizeof(str1) - strlen(str1) - 1);
```

### 3. strcmp() - String Comparison

```c
#include <stdio.h>
#include <string.h>

int main() {
    char str1[] = "Apple";
    char str2[] = "Banana";
    char str3[] = "Apple";
    
    int result1 = strcmp(str1, str2);  // Negative (Apple < Banana)
    int result2 = strcmp(str1, str3);  // 0 (Equal)
    int result3 = strcmp(str2, str1);  // Positive (Banana > Apple)
    
    if (strcmp(str1, str3) == 0) {
        printf("Strings are equal\n");
    }
    
    return 0;
}
```

**Returns:**
- `0` if strings are equal
- `< 0` if str1 < str2 (lexicographically)
- `> 0` if str1 > str2

### 4. strchr() - Find Character

```c
#include <stdio.h>
#include <string.h>

int main() {
    char str[] = "Hello, World!";
    char *ptr;
    
    ptr = strchr(str, 'W');  // Find first 'W'
    
    if (ptr != NULL) {
        printf("Found at position: %ld\n", ptr - str);  // 7
        printf("Remaining string: %s\n", ptr);  // World!
    }
    
    return 0;
}
```

### 5. strstr() - Find Substring

```c
#include <stdio.h>
#include <string.h>

int main() {
    char str[] = "Hello, World!";
    char *ptr;
    
    ptr = strstr(str, "World");  // Find "World"
    
    if (ptr != NULL) {
        printf("Substring found at position: %ld\n", ptr - str);  // 7
        printf("Remaining: %s\n", ptr);  // World!
    } else {
        printf("Substring not found\n");
    }
    
    return 0;
}
```

### 6. strtok() - String Tokenization

```c
#include <stdio.h>
#include <string.h>

int main() {
    char str[] = "apple,banana,cherry";
    char *token;
    
    token = strtok(str, ",");  // First token
    
    while (token != NULL) {
        printf("%s\n", token);
        token = strtok(NULL, ",");  // Next token
    }
    
    return 0;
}
```

**Output:**
```
apple
banana
cherry
```

### 7. Other Useful Functions

```c
#include <string.h>
#include <ctype.h>

// String length
size_t len = strlen(str);

// String duplicate (allocates memory)
char *dup = strdup(str);

// Case conversion (single character)
char upper = toupper('a');  // 'A'
char lower = tolower('A');  // 'a'

// Memory copy
memcpy(dest, src, n);

// Memory comparison
memcmp(ptr1, ptr2, n);

// Memory set
memset(str, 'A', n);
```

---

## Common String Operations

### 1. Convert String to Uppercase

```c
#include <stdio.h>
#include <ctype.h>

void toUpperCase(char str[]) {
    for (int i = 0; str[i] != '\0'; i++) {
        str[i] = toupper(str[i]);
    }
}

int main() {
    char str[] = "Hello, World!";
    
    toUpperCase(str);
    printf("%s\n", str);  // HELLO, WORLD!
    
    return 0;
}
```

### 2. Convert String to Lowercase

```c
#include <stdio.h>
#include <ctype.h>

void toLowerCase(char str[]) {
    for (int i = 0; str[i] != '\0'; i++) {
        str[i] = tolower(str[i]);
    }
}

int main() {
    char str[] = "Hello, World!";
    
    toLowerCase(str);
    printf("%s\n", str);  // hello, world!
    
    return 0;
}
```

### 3. Reverse a String

```c
#include <stdio.h>
#include <string.h>

void reverseString(char str[]) {
    int len = strlen(str);
    
    for (int i = 0; i < len / 2; i++) {
        char temp = str[i];
        str[i] = str[len - 1 - i];
        str[len - 1 - i] = temp;
    }
}

int main() {
    char str[] = "Hello";
    
    reverseString(str);
    printf("%s\n", str);  // olleH
    
    return 0;
}
```

### 4. Count Words in String

```c
#include <stdio.h>
#include <ctype.h>

int countWords(char str[]) {
    int count = 0;
    int inWord = 0;
    
    for (int i = 0; str[i] != '\0'; i++) {
        if (isspace(str[i])) {
            inWord = 0;
        } else if (!inWord) {
            inWord = 1;
            count++;
        }
    }
    
    return count;
}

int main() {
    char str[] = "Hello World from C";
    
    printf("Word count: %d\n", countWords(str));  // 4
    
    return 0;
}
```

### 5. Check if String is Palindrome

```c
#include <stdio.h>
#include <string.h>
#include <ctype.h>

int isPalindrome(char str[]) {
    int left = 0;
    int right = strlen(str) - 1;
    
    while (left < right) {
        if (tolower(str[left]) != tolower(str[right])) {
            return 0;  // Not palindrome
        }
        left++;
        right--;
    }
    
    return 1;  // Is palindrome
}

int main() {
    char str1[] = "radar";
    char str2[] = "hello";
    
    if (isPalindrome(str1)) {
        printf("%s is a palindrome\n", str1);
    }
    
    if (!isPalindrome(str2)) {
        printf("%s is not a palindrome\n", str2);
    }
    
    return 0;
}
```

### 6. Remove Spaces from String

```c
#include <stdio.h>
#include <ctype.h>

void removeSpaces(char str[]) {
    int i = 0, j = 0;
    
    while (str[i]) {
        if (!isspace(str[i])) {
            str[j++] = str[i];
        }
        i++;
    }
    str[j] = '\0';
}

int main() {
    char str[] = "Hello World from C";
    
    removeSpaces(str);
    printf("%s\n", str);  // HelloWorldfromC
    
    return 0;
}
```

---

## String Conversion Functions

### String to Integer: atoi()

```c
#include <stdio.h>
#include <stdlib.h>

int main() {
    char str[] = "12345";
    
    int num = atoi(str);
    printf("Number: %d\n", num);  // 12345
    
    return 0;
}
```

### String to Float: atof()

```c
#include <stdio.h>
#include <stdlib.h>

int main() {
    char str[] = "3.14159";
    
    double num = atof(str);
    printf("Number: %.5f\n", num);  // 3.14159
    
    return 0;
}
```

### String to Long: strtol()

```c
#include <stdio.h>
#include <stdlib.h>

int main() {
    char str[] = "12345";
    char *endptr;
    
    long num = strtol(str, &endptr, 10);  // Base 10
    printf("Number: %ld\n", num);  // 12345
    
    return 0;
}
```

### Integer to String: sprintf()

```c
#include <stdio.h>

int main() {
    int num = 12345;
    char str[20];
    
    sprintf(str, "%d", num);
    printf("String: %s\n", str);  // "12345"
    
    return 0;
}
```

---

## Array of Strings (2D Character Array)

### Method 1: 2D Character Array

```c
#include <stdio.h>

int main() {
    char names[3][20] = {
        "Alice",
        "Bob",
        "Charlie"
    };
    
    for (int i = 0; i < 3; i++) {
        printf("%s\n", names[i]);
    }
    
    return 0;
}
```

### Method 2: Array of Pointers

```c
#include <stdio.h>

int main() {
    char *names[] = {
        "Alice",
        "Bob",
        "Charlie"
    };
    
    int count = sizeof(names) / sizeof(names[0]);
    
    for (int i = 0; i < count; i++) {
        printf("%s\n", names[i]);
    }
    
    return 0;
}
```

---

## Dynamic String Allocation

### Using malloc()

```c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main() {
    char *str;
    
    // Allocate memory
    str = (char*)malloc(50 * sizeof(char));
    
    if (str == NULL) {
        printf("Memory allocation failed\n");
        return 1;
    }
    
    // Use the string
    strcpy(str, "Hello, World!");
    printf("%s\n", str);
    
    // Free memory
    free(str);
    
    return 0;
}
```

### Dynamic Input

```c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main() {
    char buffer[100];
    char *str;
    
    printf("Enter a string: ");
    fgets(buffer, sizeof(buffer), stdin);
    buffer[strcspn(buffer, "\n")] = '\0';
    
    // Allocate exact size needed
    str = (char*)malloc(strlen(buffer) + 1);
    strcpy(str, buffer);
    
    printf("You entered: %s\n", str);
    
    free(str);
    
    return 0;
}
```

---

## Common String Pitfalls

### 1. Missing Null Terminator

```c
// ❌ BAD: No null terminator
char str[5] = {'H', 'e', 'l', 'l', 'o'};
printf("%s\n", str);  // Undefined behavior!

// ✅ GOOD
char str[6] = {'H', 'e', 'l', 'l', 'o', '\0'};
// Or simply
char str[] = "Hello";
```

### 2. Buffer Overflow

```c
// ❌ BAD: Buffer too small
char str[5];
strcpy(str, "Hello, World!");  // Overflow!

// ✅ GOOD
char str[20];
strncpy(str, "Hello, World!", sizeof(str) - 1);
str[sizeof(str) - 1] = '\0';
```

### 3. Modifying String Literals

```c
// ❌ BAD: String literals are read-only
char *str = "Hello";
str[0] = 'h';  // Segmentation fault!

// ✅ GOOD: Use array
char str[] = "Hello";
str[0] = 'h';  // OK
```

### 4. Not Checking NULL

```c
// ❌ BAD
char *ptr = strchr(str, 'x');
printf("%c\n", *ptr);  // Crash if not found!

// ✅ GOOD
char *ptr = strchr(str, 'x');
if (ptr != NULL) {
    printf("%c\n", *ptr);
}
```

### 5. Comparing Strings with ==

```c
char str1[] = "Hello";
char str2[] = "Hello";

// ❌ BAD: Compares addresses, not content
if (str1 == str2) {  // Always false!
    printf("Equal\n");
}

// ✅ GOOD: Use strcmp()
if (strcmp(str1, str2) == 0) {
    printf("Equal\n");
}
```

---

## Practical Examples

### Example 1: Email Validator

```c
#include <stdio.h>
#include <string.h>

int isValidEmail(char email[]) {
    char *at = strchr(email, '@');
    char *dot = strrchr(email, '.');
    
    if (at == NULL || dot == NULL) {
        return 0;
    }
    
    if (at > dot) {
        return 0;
    }
    
    if (at == email || dot == email + strlen(email) - 1) {
        return 0;
    }
    
    return 1;
}

int main() {
    char email[100];
    
    printf("Enter email: ");
    scanf("%s", email);
    
    if (isValidEmail(email)) {
        printf("Valid email\n");
    } else {
        printf("Invalid email\n");
    }
    
    return 0;
}
```

### Example 2: Word Counter

```c
#include <stdio.h>
#include <string.h>

int main() {
    char text[200];
    char *word;
    int count = 0;
    
    printf("Enter text: ");
    fgets(text, sizeof(text), stdin);
    
    word = strtok(text, " \t\n");
    
    while (word != NULL) {
        count++;
        printf("Word %d: %s\n", count, word);
        word = strtok(NULL, " \t\n");
    }
    
    printf("\nTotal words: %d\n", count);
    
    return 0;
}
```

### Example 3: Caesar Cipher

```c
#include <stdio.h>
#include <ctype.h>

void caesarCipher(char str[], int shift) {
    for (int i = 0; str[i] != '\0'; i++) {
        if (isalpha(str[i])) {
            char base = isupper(str[i]) ? 'A' : 'a';
            str[i] = (str[i] - base + shift) % 26 + base;
        }
    }
}

int main() {
    char message[100];
    int shift;
    
    printf("Enter message: ");
    fgets(message, sizeof(message), stdin);
    message[strcspn(message, "\n")] = '\0';
    
    printf("Enter shift: ");
    scanf("%d", &shift);
    
    caesarCipher(message, shift);
    
    printf("Encrypted: %s\n", message);
    
    return 0;
}
```

---

## String Functions Summary

| Function | Purpose | Example |
|----------|---------|---------|
| `strlen()` | Get string length | `strlen("Hello")` → 5 |
| `strcpy()` | Copy string | `strcpy(dest, src)` |
| `strncpy()` | Copy n characters | `strncpy(dest, src, n)` |
| `strcat()` | Concatenate | `strcat(str1, str2)` |
| `strncat()` | Concatenate n chars | `strncat(str1, str2, n)` |
| `strcmp()` | Compare strings | `strcmp(s1, s2)` → 0 if equal |
| `strncmp()` | Compare n chars | `strncmp(s1, s2, n)` |
| `strchr()` | Find character | `strchr(str, 'c')` |
| `strrchr()` | Find last char | `strrchr(str, 'c')` |
| `strstr()` | Find substring | `strstr(str, "sub")` |
| `strtok()` | Tokenize | `strtok(str, delim)` |
| `atoi()` | String to int | `atoi("123")` → 123 |
| `atof()` | String to float | `atof("3.14")` → 3.14 |
| `sprintf()` | Format to string | `sprintf(str, "%d", num)` |

---

## Best Practices

1. **Always null-terminate strings** manually when needed
2. **Use `fgets()` instead of `gets()`** for input
3. **Check buffer sizes** to prevent overflow
4. **Use `strncpy()`, `strncat()`** for safer operations
5. **Check for NULL** after functions like `strchr()`, `strstr()`
6. **Use `strcmp()` for comparison**, not `==`
7. **Free dynamically allocated strings** with `free()`
8. **Validate input** before processing
9. **Use `const` for read-only strings** in functions
10. **Initialize strings properly** to avoid garbage values

```c
// ✅ GOOD PRACTICES
char str[50] = {0};  // Initialize to zeros
fgets(str, sizeof(str), stdin);  // Safe input
str[strcspn(str, "\n")] = '\0';  // Remove newline
if (strcmp(str1, str2) == 0) { }  // Proper comparison
void func(const char *str);  // Read-only parameter
```

Strings in C require careful handling, but mastering them is essential for C programming!