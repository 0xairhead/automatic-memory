# Unions in C

## Table of Contents

- [What are Unions?](#what-are-unions)
- [Union vs Structure](#union-vs-structure)
- [Why Use Unions?](#why-use-unions)
- [Basic Union Syntax](#basic-union-syntax)
- [Declaring Union Variables](#declaring-union-variables)
- [Accessing Union Members](#accessing-union-members)
- [Initializing Unions](#initializing-unions)
- [typedef with Unions](#typedef-with-unions)
- [Memory Layout of Unions](#memory-layout-of-unions)
- [Unions with Structures](#unions-with-structures)
- [Anonymous Unions](#anonymous-unions)
- [Unions and Pointers](#unions-and-pointers)
- [Practical Use Cases](#practical-use-cases)
- [Tagged Unions](#tagged-unions)
- [Bit Fields in Unions](#bit-fields-in-unions)
- [Common Pitfalls](#common-pitfalls)
- [Best Practices](#best-practices)
- [Quick Reference](#quick-reference)

---

## What are Unions?

**Unions** are user-defined data types that allow storing different data types in the same memory location. Unlike structures, unions can hold only **one member value at a time**.

```
Structure (separate memory for each member):
struct Example {
    int a;    // 4 bytes
    float b;  // 4 bytes
    char c;   // 1 byte
};
Total: ~12 bytes (with padding)

Union (shared memory for all members):
union Example {
    int a;    // 4 bytes
    float b;  // 4 bytes
    char c;   // 1 byte
};
Total: 4 bytes (size of largest member)
```

**Key Point:** All members of a union share the same memory location!

---

## Union vs Structure

| Feature | Structure | Union |
|---------|-----------|-------|
| **Memory Allocation** | Each member gets separate memory | All members share same memory |
| **Size** | Sum of all members (+ padding) | Size of largest member |
| **Access** | All members accessible simultaneously | Only one member valid at a time |
| **Use Case** | Store multiple values together | Store one of many possible types |
| **Keyword** | `struct` | `union` |

### Visual Comparison

```c
#include <stdio.h>

struct MyStruct {
    int a;
    float b;
    char c;
};

union MyUnion {
    int a;
    float b;
    char c;
};

int main() {
    printf("Size of struct: %zu bytes\n", sizeof(struct MyStruct));
    printf("Size of union: %zu bytes\n", sizeof(union MyUnion));
    
    return 0;
}
```

**Output:**
```
Size of struct: 12 bytes
Size of union: 4 bytes
```

---

## Why Use Unions?

1. **Memory Efficiency** - Save memory when only one value is needed at a time
2. **Type Conversion** - View same data in different ways
3. **Variant Data Types** - Store different types in same variable
4. **Hardware Programming** - Access individual bits/bytes of data
5. **Protocol Implementation** - Handle different message types

**Real-world analogy:** A parking space that can hold either a car OR a motorcycle OR a bicycle, but only one at a time.

---

## Basic Union Syntax

### Defining a Union

```c
union UnionName {
    datatype member1;
    datatype member2;
    datatype member3;
    // ... more members
};
```

### Basic Example

```c
#include <stdio.h>

union Data {
    int i;
    float f;
    char c;
};

int main() {
    union Data data;
    
    return 0;
}
```

---

## Declaring Union Variables

### Method 1: After Union Definition

```c
union Data {
    int i;
    float f;
    char str[20];
};

int main() {
    union Data data1;           // Single variable
    union Data d1, d2, d3;      // Multiple variables
    
    return 0;
}
```

### Method 2: During Union Definition

```c
union Data {
    int i;
    float f;
    char c;
} data1, data2;  // Declare while defining

int main() {
    // data1 and data2 are already declared
    return 0;
}
```

### Method 3: Anonymous Union

```c
union {
    int value;
    float decimal;
} data1, data2;  // Can't create more variables of this type later
```

---

## Accessing Union Members

Use the **dot operator** (`.`) to access union members, just like structures.

### Basic Example

```c
#include <stdio.h>

union Data {
    int i;
    float f;
    char c;
};

int main() {
    union Data data;
    
    data.i = 10;
    printf("data.i = %d\n", data.i);
    
    data.f = 3.14;
    printf("data.f = %.2f\n", data.f);
    
    data.c = 'A';
    printf("data.c = %c\n", data.c);
    
    return 0;
}
```

**Output:**
```
data.i = 10
data.f = 3.14
data.c = A
```

### ⚠️ Important: Only Last Assignment is Valid

```c
#include <stdio.h>

union Data {
    int i;
    float f;
    char c;
};

int main() {
    union Data data;
    
    data.i = 10;
    data.f = 3.14;     // This overwrites data.i
    data.c = 'A';      // This overwrites data.f
    
    printf("data.i = %d\n", data.i);     // Garbage value!
    printf("data.f = %.2f\n", data.f);   // Garbage value!
    printf("data.c = %c\n", data.c);     // Valid: 'A'
    
    return 0;
}
```

**Output (undefined behavior for i and f):**
```
data.i = 65
data.f = 0.00
data.c = A
```

### Demonstrating Shared Memory

```c
#include <stdio.h>

union Test {
    int i;
    char c[4];
};

int main() {
    union Test test;
    
    test.i = 0x41424344;  // Hexadecimal value
    
    printf("test.i = 0x%X\n", test.i);
    printf("test.c[0] = %c (0x%X)\n", test.c[0], test.c[0]);
    printf("test.c[1] = %c (0x%X)\n", test.c[1], test.c[1]);
    printf("test.c[2] = %c (0x%X)\n", test.c[2], test.c[2]);
    printf("test.c[3] = %c (0x%X)\n", test.c[3], test.c[3]);
    
    return 0;
}
```

**Output (on little-endian system):**
```
test.i = 0x41424344
test.c[0] = D (0x44)
test.c[1] = C (0x43)
test.c[2] = B (0x42)
test.c[3] = A (0x41)
```

---

## Initializing Unions

### Method 1: Initialize First Member

```c
union Data {
    int i;
    float f;
    char c;
};

int main() {
    union Data data = {10};  // Initializes first member (i)
    
    return 0;
}
```

### Method 2: Designated Initializers (C99+)

```c
union Data {
    int i;
    float f;
    char c;
};

int main() {
    union Data data1 = {.f = 3.14};   // Initialize specific member
    union Data data2 = {.c = 'A'};
    union Data data3 = {.i = 100};
    
    return 0;
}
```

### Complete Example

```c
#include <stdio.h>

union Number {
    int integer;
    float decimal;
};

int main() {
    // Initialize first member
    union Number num1 = {42};
    printf("num1.integer = %d\n", num1.integer);
    
    // Designated initializer
    union Number num2 = {.decimal = 3.14};
    printf("num2.decimal = %.2f\n", num2.decimal);
    
    // Assign after declaration
    union Number num3;
    num3.integer = 100;
    printf("num3.integer = %d\n", num3.integer);
    
    return 0;
}
```

**Output:**
```
num1.integer = 42
num2.decimal = 3.14
num3.integer = 100
```

---

## typedef with Unions

`typedef` creates an alias for a union type, just like with structures.

### Without typedef

```c
union Data {
    int i;
    float f;
};

int main() {
    union Data d1;  // Must use 'union' keyword
    union Data d2;
    
    return 0;
}
```

### With typedef

```c
typedef union {
    int i;
    float f;
} Data;

int main() {
    Data d1;  // No 'union' keyword needed!
    Data d2;
    
    return 0;
}
```

### Complete Example

```c
#include <stdio.h>

typedef union {
    int asInt;
    float asFloat;
    char asBytes[4];
} Value;

int main() {
    Value val;
    
    val.asInt = 1078523331;  // Binary representation of ~3.14
    printf("As int: %d\n", val.asInt);
    printf("As float: %.2f\n", val.asFloat);
    
    return 0;
}
```

---

## Memory Layout of Unions

All union members start at the same memory address and share the same memory space.

### Memory Visualization

```c
#include <stdio.h>

union Example {
    int i;      // 4 bytes
    float f;    // 4 bytes
    char c;     // 1 byte
};

int main() {
    union Example ex;
    
    printf("Size of union: %zu bytes\n", sizeof(ex));
    printf("Address of union: %p\n", (void*)&ex);
    printf("Address of ex.i: %p\n", (void*)&ex.i);
    printf("Address of ex.f: %p\n", (void*)&ex.f);
    printf("Address of ex.c: %p\n", (void*)&ex.c);
    
    return 0;
}
```

**Output:**
```
Size of union: 4 bytes
Address of union: 0x7fff5fbff8ac
Address of ex.i: 0x7fff5fbff8ac
Address of ex.f: 0x7fff5fbff8ac
Address of ex.c: 0x7fff5fbff8ac
```

### Size Calculation

```c
#include <stdio.h>

union Small {
    char c;     // 1 byte
    int i;      // 4 bytes
};

union Large {
    char c;     // 1 byte
    double d;   // 8 bytes
    int arr[5]; // 20 bytes
};

int main() {
    printf("Size of Small union: %zu bytes\n", sizeof(union Small));
    printf("Size of Large union: %zu bytes\n", sizeof(union Large));
    
    return 0;
}
```

**Output:**
```
Size of Small union: 4 bytes
Size of Large union: 20 bytes
```

**Rule:** Union size = size of its largest member (+ alignment padding if needed)

---

## Unions with Structures

Unions are often combined with structures to create flexible data types.

### Example 1: Variant Record

```c
#include <stdio.h>
#include <string.h>

typedef enum {
    TYPE_INT,
    TYPE_FLOAT,
    TYPE_STRING
} DataType;

typedef union {
    int i;
    float f;
    char str[50];
} Value;

typedef struct {
    DataType type;
    Value value;
} Data;

void printData(Data *d) {
    switch(d->type) {
        case TYPE_INT:
            printf("Integer: %d\n", d->value.i);
            break;
        case TYPE_FLOAT:
            printf("Float: %.2f\n", d->value.f);
            break;
        case TYPE_STRING:
            printf("String: %s\n", d->value.str);
            break;
    }
}

int main() {
    Data data1;
    data1.type = TYPE_INT;
    data1.value.i = 42;
    printData(&data1);
    
    Data data2;
    data2.type = TYPE_FLOAT;
    data2.value.f = 3.14;
    printData(&data2);
    
    Data data3;
    data3.type = TYPE_STRING;
    strcpy(data3.value.str, "Hello, World!");
    printData(&data3);
    
    return 0;
}
```

**Output:**
```
Integer: 42
Float: 3.14
String: Hello, World!
```

### Example 2: IP Address Representation

```c
#include <stdio.h>

typedef union {
    unsigned int address;
    unsigned char bytes[4];
} IPAddress;

typedef struct {
    char name[50];
    IPAddress ip;
} Server;

void printIP(IPAddress *ip) {
    printf("IP Address: %d.%d.%d.%d\n",
           ip->bytes[0], ip->bytes[1], ip->bytes[2], ip->bytes[3]);
}

int main() {
    Server server;
    strcpy(server.name, "Web Server");
    
    // Set IP as 192.168.1.1
    server.ip.bytes[0] = 192;
    server.ip.bytes[1] = 168;
    server.ip.bytes[2] = 1;
    server.ip.bytes[3] = 1;
    
    printf("Server: %s\n", server.name);
    printIP(&server.ip);
    printf("As integer: %u\n", server.ip.address);
    
    return 0;
}
```

---

## Anonymous Unions

Anonymous unions can be placed inside structures (C11 standard).

### Example

```c
#include <stdio.h>

typedef struct {
    int id;
    enum { INT_TYPE, FLOAT_TYPE } type;
    union {  // Anonymous union
        int iValue;
        float fValue;
    };
} Data;

int main() {
    Data data1;
    data1.id = 1;
    data1.type = INT_TYPE;
    data1.iValue = 100;  // Access without union name
    
    printf("ID: %d\n", data1.id);
    printf("Value: %d\n", data1.iValue);
    
    Data data2;
    data2.id = 2;
    data2.type = FLOAT_TYPE;
    data2.fValue = 3.14;  // Access without union name
    
    printf("ID: %d\n", data2.id);
    printf("Value: %.2f\n", data2.fValue);
    
    return 0;
}
```

**Output:**
```
ID: 1
Value: 100
ID: 2
Value: 3.14
```

---

## Unions and Pointers

Pointers to unions work just like pointers to structures.

### Basic Pointer Usage

```c
#include <stdio.h>

typedef union {
    int i;
    float f;
    char c;
} Data;

int main() {
    Data data;
    Data *ptr = &data;
    
    // Access using arrow operator
    ptr->i = 100;
    printf("ptr->i = %d\n", ptr->i);
    
    ptr->f = 3.14;
    printf("ptr->f = %.2f\n", ptr->f);
    
    ptr->c = 'X';
    printf("ptr->c = %c\n", ptr->c);
    
    return 0;
}
```

### Function with Union Pointer

```c
#include <stdio.h>

typedef union {
    int integer;
    float decimal;
} Number;

void setInteger(Number *num, int value) {
    num->integer = value;
}

void setDecimal(Number *num, float value) {
    num->decimal = value;
}

void printAsInteger(Number *num) {
    printf("As integer: %d\n", num->integer);
}

void printAsDecimal(Number *num) {
    printf("As decimal: %.2f\n", num->decimal);
}

int main() {
    Number num;
    
    setInteger(&num, 42);
    printAsInteger(&num);
    
    setDecimal(&num, 3.14);
    printAsDecimal(&num);
    
    return 0;
}
```

**Output:**
```
As integer: 42
As decimal: 3.14
```

---

## Practical Use Cases

### Use Case 1: Type Conversion / Type Punning

```c
#include <stdio.h>

typedef union {
    float f;
    unsigned int bits;
} FloatBits;

void printFloatBits(float value) {
    FloatBits fb;
    fb.f = value;
    
    printf("Float value: %.2f\n", fb.f);
    printf("Bit representation: 0x%08X\n", fb.bits);
    printf("Binary: ");
    
    for (int i = 31; i >= 0; i--) {
        printf("%d", (fb.bits >> i) & 1);
        if (i % 8 == 0) printf(" ");
    }
    printf("\n");
}

int main() {
    printFloatBits(3.14);
    printf("\n");
    printFloatBits(-1.5);
    
    return 0;
}
```

### Use Case 2: Byte Order (Endianness) Detection

```c
#include <stdio.h>

typedef union {
    unsigned int value;
    unsigned char bytes[4];
} EndianTest;

int main() {
    EndianTest test;
    test.value = 0x01020304;
    
    printf("Value: 0x%08X\n", test.value);
    printf("Byte order: ");
    
    for (int i = 0; i < 4; i++) {
        printf("0x%02X ", test.bytes[i]);
    }
    printf("\n");
    
    if (test.bytes[0] == 0x04) {
        printf("System is Little Endian\n");
    } else if (test.bytes[0] == 0x01) {
        printf("System is Big Endian\n");
    }
    
    return 0;
}
```

**Output (on little-endian system):**
```
Value: 0x01020304
Byte order: 0x04 0x03 0x02 0x01 
System is Little Endian
```

### Use Case 3: Memory-Efficient Storage

```c
#include <stdio.h>
#include <string.h>

typedef enum {
    MSG_TEXT,
    MSG_NUMBER,
    MSG_BOOL
} MessageType;

typedef union {
    char text[100];
    int number;
    int boolean;  // 0 or 1
} MessageData;

typedef struct {
    MessageType type;
    MessageData data;
} Message;

void printMessage(Message *msg) {
    switch(msg->type) {
        case MSG_TEXT:
            printf("Text message: %s\n", msg->data.text);
            break;
        case MSG_NUMBER:
            printf("Number message: %d\n", msg->data.number);
            break;
        case MSG_BOOL:
            printf("Boolean message: %s\n", msg->data.boolean ? "true" : "false");
            break;
    }
}

int main() {
    Message messages[3];
    
    messages[0].type = MSG_TEXT;
    strcpy(messages[0].data.text, "Hello, World!");
    
    messages[1].type = MSG_NUMBER;
    messages[1].data.number = 42;
    
    messages[2].type = MSG_BOOL;
    messages[2].data.boolean = 1;
    
    for (int i = 0; i < 3; i++) {
        printMessage(&messages[i]);
    }
    
    printf("\nSize of Message: %zu bytes\n", sizeof(Message));
    
    return 0;
}
```

**Output:**
```
Text message: Hello, World!
Number message: 42
Boolean message: true

Size of Message: 108 bytes
```

### Use Case 4: Color Representation

```c
#include <stdio.h>

typedef union {
    unsigned int color;  // 32-bit color
    struct {
        unsigned char b;  // Blue
        unsigned char g;  // Green
        unsigned char r;  // Red
        unsigned char a;  // Alpha
    } channels;
} Color;

void printColor(Color *c) {
    printf("Color: 0x%08X\n", c->color);
    printf("RGBA: (%d, %d, %d, %d)\n",
           c->channels.r, c->channels.g, c->channels.b, c->channels.a);
}

int main() {
    Color red;
    red.color = 0xFF0000FF;  // Red with full opacity
    
    printf("Red color:\n");
    printColor(&red);
    
    Color blue;
    blue.channels.r = 0;
    blue.channels.g = 0;
    blue.channels.b = 255;
    blue.channels.a = 255;
    
    printf("\nBlue color:\n");
    printColor(&blue);
    
    return 0;
}
```

### Use Case 5: Hardware Register Access

```c
#include <stdio.h>

// Simulating a hardware register
typedef union {
    unsigned int value;
    struct {
        unsigned int bit0 : 1;
        unsigned int bit1 : 1;
        unsigned int bit2 : 1;
        unsigned int bit3 : 1;
        unsigned int bit4 : 1;
        unsigned int bit5 : 1;
        unsigned int bit6 : 1;
        unsigned int bit7 : 1;
        unsigned int reserved : 24;
    } bits;
} Register;

int main() {
    Register reg;
    reg.value = 0;
    
    // Set individual bits
    reg.bits.bit0 = 1;
    reg.bits.bit2 = 1;
    reg.bits.bit7 = 1;
    
    printf("Register value: 0x%08X\n", reg.value);
    printf("Binary: ");
    
    for (int i = 7; i >= 0; i--) {
        printf("%d", (reg.value >> i) & 1);
    }
    printf("\n");
    
    return 0;
}
```

**Output:**
```
Register value: 0x00000085
Binary: 10000101
```

---

## Tagged Unions

A **tagged union** (also called discriminated union) is a union combined with a tag that identifies which member is currently valid.

### Basic Tagged Union

```c
#include <stdio.h>
#include <string.h>

typedef enum {
    TYPE_INT,
    TYPE_FLOAT,
    TYPE_STRING
} ValueType;

typedef union {
    int i;
    float f;
    char str[50];
} ValueData;

typedef struct {
    ValueType type;  // Tag
    ValueData data;  // Union
} TaggedValue;

void printValue(TaggedValue *tv) {
    switch(tv->type) {
        case TYPE_INT:
            printf("Integer: %d\n", tv->data.i);
            break;
        case TYPE_FLOAT:
            printf("Float: %.2f\n", tv->data.f);
            break;
        case TYPE_STRING:
            printf("String: %s\n", tv->data.str);
            break;
        default:
            printf("Unknown type\n");
    }
}

int main() {
    TaggedValue val1 = {TYPE_INT, {.i = 42}};
    TaggedValue val2 = {TYPE_FLOAT, {.f = 3.14}};
    TaggedValue val3;
    val3.type = TYPE_STRING;
    strcpy(val3.data.str, "Hello");
    
    printValue(&val1);
    printValue(&val2);
    printValue(&val3);
    
    return 0;
}
```

**Output:**
```
Integer: 42
Float: 3.14
String: Hello
```

### Advanced Tagged Union: Expression Evaluator

```c
#include <stdio.h>

typedef enum {
    EXPR_NUMBER,
    EXPR_ADD,
    EXPR_MULTIPLY
} ExprType;

typedef struct Expr Expr;

typedef union {
    int number;
    struct {
        Expr *left;
        Expr *right;
    } binary;
} ExprData;

struct Expr {
    ExprType type;
    ExprData data;
};

int evaluate(Expr *expr) {
    switch(expr->type) {
        case EXPR_NUMBER:
            return expr->data.number;
        case EXPR_ADD:
            return evaluate(expr->data.binary.left) + 
                   evaluate(expr->data.binary.right);
        case EXPR_MULTIPLY:
            return evaluate(expr->data.binary.left) * 
                   evaluate(expr->data.binary.right);
    }
    return 0;
}

int main() {
    // Create expression: (5 + 3) * 2
    Expr num5 = {EXPR_NUMBER, {.number = 5}};
    Expr num3 = {EXPR_NUMBER, {.number = 3}};
    Expr num2 = {EXPR_NUMBER, {.number = 2}};
    
    Expr add = {EXPR_ADD, {.binary = {&num5, &num3}}};
    Expr multiply = {EXPR_MULTIPLY, {.binary = {&add, &num2}}};
    
    printf("Result: %d\n", evaluate(&multiply));  // (5 + 3) * 2 = 16
    
    return 0;
}
```

**Output:**
```
Result: 16
```

---

## Bit Fields in Unions

Bit fields can be used within unions for fine-grained control over memory.

### Example: Flags and Status Register

```c
#include <stdio.h>

typedef union {
    unsigned char value;
    struct {
        unsigned char ready : 1;
        unsigned char error : 1;
        unsigned char busy : 1;
        unsigned char ack : 1;
        unsigned char reserved : 4;
    } flags;
} StatusRegister;

void printStatus(StatusRegister *sr) {
    printf("Status: 0x%02X\n", sr->value);
    printf("  Ready: %d\n", sr->flags.ready);
    printf("  Error: %d\n", sr->flags.error);
    printf("  Busy: %d\n", sr->flags.busy);
    printf("  Ack: %d\n", sr->flags.ack);
}

int main() {
    StatusRegister status;
    status.value = 0;
    
    status.flags.ready = 1;
    status.flags.busy = 1;
    
    printf("Initial status:\n");
    printStatus(&status);
    
    status.flags.ready = 0;
    status.flags.busy = 0;
    status.flags.error = 1;
    
    printf("\nAfter error:\n");
    printStatus(&status);
    
    return 0;
}
```

**Output:**
```
Initial status:
Status: 0x05
  Ready: 1
  Error: 0
  Busy: 1
  Ack: 0

After error:
Status: 0x02
  Ready: 0
  Error: 1
  Busy: 0
  Ack: 0
```

### Example: Packed Data Structure

```c
#include <stdio.h>

typedef union {
    unsigned short value;
    struct {
        unsigned short day : 5;     // 1-31
        unsigned short month : 4;   // 1-12
        unsigned short year : 7;    // 0-127 (offset from 2000)
    } date;
} CompactDate;

void printDate(CompactDate *cd) {
    printf("Date: %02d/%02d/%04d\n", 
           cd->date.day, 
           cd->date.month, 
           cd->date.year + 2000);
    printf("Packed value: 0x%04X\n", cd->value);
}

int main() {
    CompactDate cd;
    cd.date.day = 25;
    cd.date.month = 12;
    cd.date.year = 24;  // 2024
    
    printDate(&cd);
    printf("Size: %zu bytes\n", sizeof(CompactDate));
    
    return 0;
}
```

**Output:**
```
Date: 25/12/2024
Packed value: 0x318C
Size: 2 bytes
```

---

## Common Pitfalls

### 1. Reading Wrong Member

```c
#include <stdio.h>

union Data {
    int i;
    float f;
};

int main() {
    union Data data;
    
    data.i = 42;
    
    // ❌ BAD: Reading wrong member
    printf("As float: %.2f\n", data.f);  // Undefined behavior!
    
    // ✅ GOOD: Read the member you wrote to
    printf("As int: %d\n", data.i);
    
    return 0;
}
```

### 2. Forgetting Which Member is Active

```c
// ❌ BAD: No way to know which member is valid
union Value {
    int i;
    float f;
    char str[20];
};

// ✅ GOOD: Use tagged union
typedef enum { INT, FLOAT, STRING } Type;

typedef struct {
    Type type;  // Tag to track active member
    union {
        int i;
        float f;
        char str[20];
    } value;
} TaggedValue;
```

### 3. Assuming Member Order

```c
#include <stdio.h>

union Example {
    int a;
    float b;
    char c;
};

int main() {
    // ❌ BAD: Order of members doesn't matter in unions
    union Example ex = {10, 3.14, 'A'};  // ERROR: Can't initialize multiple members
    
    // ✅ GOOD: Initialize one member
    union Example ex2 = {10};  // OK: Initializes first member
    union Example ex3 = {.b = 3.14};  // OK: Designated initializer
    
    return 0;
}
```

### 4. Using sizeof on Union Members

```c
#include <stdio.h>

union Data {
    char c;
    int i;
    double d;
};

int main() {
    union Data data;
    
    // ❌ CONFUSING: sizeof on member
    printf("%zu\n", sizeof(data.c));  // 1 byte (size of char)
    
    // ✅ CLEAR: sizeof on union
    printf("%zu\n", sizeof(data));     // 8 bytes (size of largest member)
    
    return 0;
}
```

### 5. Comparing Unions

```c
// ❌ BAD: Can't compare unions directly
union Data d1, d2;
if (d1 == d2) { }  // ERROR!

// ✅ GOOD: Compare members individually
if (d1.i == d2.i) { }  // If you know 'i' is active
```

### 6. Portability Issues

```c
#include <stdio.h>

union Test {
    int i;
    char bytes[4];
};

int main() {
    union Test t;
    t.i = 0x01020304;
    
    // ⚠️ WARNING: Results depend on endianness!
    printf("%02X %02X %02X %02X\n",
           t.bytes[0], t.bytes[1], t.bytes[2], t.bytes[3]);
    // Little-endian: 04 03 02 01
    // Big-endian: 01 02 03 04
    
    return 0;
}
```

---

## Best Practices

1. **Always use tagged unions** to track which member is active
2. **Only read the member you last wrote to** to avoid undefined behavior
3. **Document which member should be used when**
4. **Use typedef** to simplify union declarations
5. **Prefer structures when you need all values simultaneously**
6. **Use unions for memory efficiency** or type conversion
7. **Be aware of endianness** when accessing bytes
8. **Don't assume padding or alignment** - use sizeof
9. **Initialize unions explicitly** to avoid confusion
10. **Consider using enums for tags** in tagged unions

```c
// ✅ GOOD: Well-designed tagged union
typedef enum {
    DATA_TYPE_INT,
    DATA_TYPE_FLOAT,
    DATA_TYPE_STRING
} DataType;

typedef struct {
    DataType type;  // Always know which member is valid
    union {
        int intValue;
        float floatValue;
        char stringValue[100];
    } data;
} Data;

// ✅ GOOD: Safe access function
void printData(const Data *d) {
    switch(d->type) {
        case DATA_TYPE_INT:
            printf("Int: %d\n", d->data.intValue);
            break;
        case DATA_TYPE_FLOAT:
            printf("Float: %.2f\n", d->data.floatValue);
            break;
        case DATA_TYPE_STRING:
            printf("String: %s\n", d->data.stringValue);
            break;
        default:
            printf("Unknown type\n");
    }
}
```

---

## Quick Reference

```c
// Define union
union UnionName {
    datatype member1;
    datatype member2;
};

// With typedef
typedef union {
    datatype member1;
    datatype member2;
} UnionName;

// Declare variable
union UnionName var1;
UnionName var2;  // If using typedef

// Access members
var1.member1 = value;

// Initialize
union UnionName var = {value};              // First member
union UnionName var = {.member2 = value};   // Designated initializer

// Pointer to union
union UnionName *ptr = &var1;
ptr->member1 = value;

// Tagged union pattern
typedef enum { TYPE_A, TYPE_B } Tag;
typedef struct {
    Tag type;
    union {
        TypeA a;
        TypeB b;
    } data;
} TaggedUnion;

// Size of union
sizeof(union UnionName)  // Size of largest member

// Key rules
// 1. All members share same memory
// 2. Only one member valid at a time
// 3. Size = size of largest member
// 4. Use tags to track active member
```

**Key Difference:**
```c
struct MyStruct {
    int a;
    float b;
};
// Size ≈ sizeof(int) + sizeof(float) + padding
// Can access both a and b simultaneously

union MyUnion {
    int a;
    float b;
};
// Size = max(sizeof(int), sizeof(float))
// Can only use a OR b, not both
```

Unions are powerful for memory efficiency and low-level programming - use them wisely with proper tagging!
