# Structures in C

## Table of Contents

- [What are Structures?](#what-are-structures)
- [Why Use Structures?](#why-use-structures)
- [Basic Structure Syntax](#basic-structure-syntax)
- [Declaring Structure Variables](#declaring-structure-variables)
- [Accessing Structure Members](#accessing-structure-members)
- [Initializing Structures](#initializing-structures)
- [typedef with Structures](#typedef-with-structures)
- [Nested Structures](#nested-structures)
- [Arrays of Structures](#arrays-of-structures)
- [Structures and Pointers](#structures-and-pointers)
- [Structures and Functions](#structures-and-functions)
- [Structure Padding and Alignment](#structure-padding-and-alignment)
- [Practical Examples](#practical-examples)
- [Common Pitfalls](#common-pitfalls)
- [Best Practices](#best-practices)
- [Quick Reference](#quick-reference)

---

## What are Structures?

**Structures** (or structs) are user-defined data types that allow you to group related variables of different types together under a single name.

```
Without structure:
char name[50];
int age;
float salary;
int employeeID;

With structure:
struct Employee {
    char name[50];
    int age;
    float salary;
    int employeeID;
};
```

Think of a structure as a **container** that holds related information together, like a record in a database or an object in real life.

---

## Why Use Structures?

1. **Organization** - Group related data together
2. **Readability** - Code is easier to understand
3. **Maintainability** - Changes are easier to manage
4. **Reusability** - Define once, use multiple times
5. **Real-world modeling** - Represent real entities (Person, Car, Book)

---

## Basic Structure Syntax

### Defining a Structure

```c
struct StructureName {
    datatype member1;
    datatype member2;
    datatype member3;
    // ... more members
};
```

### Example: Student Structure

```c
#include <stdio.h>

struct Student {
    char name[50];
    int rollNumber;
    float marks;
    char grade;
};

int main() {
    struct Student student1;
    
    return 0;
}
```

**Key Points:**
- Structure definition ends with a semicolon
- Members can be of any data type
- Structure is just a template - no memory is allocated until you create a variable

---

## Declaring Structure Variables

### Method 1: After Structure Definition

```c
struct Student {
    char name[50];
    int rollNumber;
    float marks;
};

int main() {
    struct Student student1;      // Single variable
    struct Student s1, s2, s3;    // Multiple variables
    
    return 0;
}
```

### Method 2: During Structure Definition

```c
struct Student {
    char name[50];
    int rollNumber;
    float marks;
} student1, student2;  // Declare while defining

int main() {
    // student1 and student2 are already declared
    return 0;
}
```

### Method 3: Anonymous Structure

```c
struct {
    char title[100];
    int pages;
} book1, book2;  // Can't create more variables of this type later
```

---

## Accessing Structure Members

Use the **dot operator** (`.`) to access structure members.

### Syntax

```c
structureVariable.memberName
```

### Basic Example

```c
#include <stdio.h>
#include <string.h>

struct Person {
    char name[50];
    int age;
    float height;
};

int main() {
    struct Person person1;
    
    // Assign values
    strcpy(person1.name, "Alice");
    person1.age = 25;
    person1.height = 5.6;
    
    // Access and print values
    printf("Name: %s\n", person1.name);
    printf("Age: %d\n", person1.age);
    printf("Height: %.1f feet\n", person1.height);
    
    return 0;
}
```

**Output:**
```
Name: Alice
Age: 25
Height: 5.6 feet
```

---

## Initializing Structures

### Method 1: Member by Member

```c
struct Point {
    int x;
    int y;
};

int main() {
    struct Point p1;
    p1.x = 10;
    p1.y = 20;
    
    return 0;
}
```

### Method 2: At Declaration (Order Matters)

```c
struct Point {
    int x;
    int y;
};

int main() {
    struct Point p1 = {10, 20};  // x=10, y=20
    
    return 0;
}
```

### Method 3: Designated Initializers (C99+)

```c
struct Point {
    int x;
    int y;
};

int main() {
    struct Point p1 = {.x = 10, .y = 20};
    struct Point p2 = {.y = 30, .x = 15};  // Order doesn't matter
    
    return 0;
}
```

### Partial Initialization

```c
struct Point {
    int x;
    int y;
    int z;
};

int main() {
    struct Point p1 = {10, 20};  // z is automatically 0
    struct Point p2 = {0};       // All members set to 0
    
    return 0;
}
```

### Complete Example

```c
#include <stdio.h>
#include <string.h>

struct Book {
    char title[100];
    char author[50];
    int pages;
    float price;
};

int main() {
    // Method 1: Initialize at declaration
    struct Book book1 = {"C Programming", "Dennis Ritchie", 228, 29.99};
    
    // Method 2: Designated initializers
    struct Book book2 = {
        .title = "The C Programming Language",
        .author = "Kernighan & Ritchie",
        .pages = 272,
        .price = 39.99
    };
    
    // Method 3: Assign after declaration
    struct Book book3;
    strcpy(book3.title, "Learn C");
    strcpy(book3.author, "John Doe");
    book3.pages = 350;
    book3.price = 24.99;
    
    printf("Book 1: %s by %s\n", book1.title, book1.author);
    printf("Book 2: %s by %s\n", book2.title, book2.author);
    printf("Book 3: %s by %s\n", book3.title, book3.author);
    
    return 0;
}
```

---

## typedef with Structures

`typedef` creates an **alias** for a data type, eliminating the need to write `struct` keyword repeatedly.

### Without typedef

```c
struct Student {
    char name[50];
    int age;
};

int main() {
    struct Student s1;  // Must use 'struct' keyword
    struct Student s2;
    
    return 0;
}
```

### With typedef

```c
typedef struct {
    char name[50];
    int age;
} Student;

int main() {
    Student s1;  // No 'struct' keyword needed!
    Student s2;
    
    return 0;
}
```

### Alternative typedef Syntax

```c
// Define first, then typedef
struct Student {
    char name[50];
    int age;
};

typedef struct Student Student;

int main() {
    Student s1;
    
    return 0;
}
```

### Complete Example

```c
#include <stdio.h>
#include <string.h>

typedef struct {
    char brand[30];
    char model[30];
    int year;
    float price;
} Car;

int main() {
    Car car1 = {"Toyota", "Camry", 2023, 25000.00};
    Car car2;
    
    strcpy(car2.brand, "Honda");
    strcpy(car2.model, "Civic");
    car2.year = 2024;
    car2.price = 23000.00;
    
    printf("Car 1: %d %s %s - $%.2f\n", 
           car1.year, car1.brand, car1.model, car1.price);
    printf("Car 2: %d %s %s - $%.2f\n", 
           car2.year, car2.brand, car2.model, car2.price);
    
    return 0;
}
```

**Output:**
```
Car 1: 2023 Toyota Camry - $25000.00
Car 2: 2024 Honda Civic - $23000.00
```

---

## Nested Structures

Structures can contain other structures as members.

### Basic Nested Structure

```c
#include <stdio.h>

struct Date {
    int day;
    int month;
    int year;
};

struct Employee {
    char name[50];
    int id;
    struct Date joinDate;  // Nested structure
};

int main() {
    struct Employee emp1;
    
    strcpy(emp1.name, "John Smith");
    emp1.id = 1001;
    emp1.joinDate.day = 15;
    emp1.joinDate.month = 6;
    emp1.joinDate.year = 2023;
    
    printf("Employee: %s\n", emp1.name);
    printf("ID: %d\n", emp1.id);
    printf("Join Date: %d/%d/%d\n", 
           emp1.joinDate.day, 
           emp1.joinDate.month, 
           emp1.joinDate.year);
    
    return 0;
}
```

**Output:**
```
Employee: John Smith
ID: 1001
Join Date: 15/6/2023
```

### Example: Address in Person

```c
#include <stdio.h>
#include <string.h>

typedef struct {
    char street[100];
    char city[50];
    char state[30];
    int zipCode;
} Address;

typedef struct {
    char name[50];
    int age;
    Address address;  // Nested structure
} Person;

int main() {
    Person person1 = {
        .name = "Alice Johnson",
        .age = 30,
        .address = {
            .street = "123 Main St",
            .city = "New York",
            .state = "NY",
            .zipCode = 10001
        }
    };
    
    printf("Name: %s\n", person1.name);
    printf("Age: %d\n", person1.age);
    printf("Address: %s, %s, %s %d\n",
           person1.address.street,
           person1.address.city,
           person1.address.state,
           person1.address.zipCode);
    
    return 0;
}
```

### Multiple Levels of Nesting

```c
#include <stdio.h>

typedef struct {
    int day;
    int month;
    int year;
} Date;

typedef struct {
    char eventName[50];
    Date eventDate;
} Event;

typedef struct {
    char personName[50];
    Event birthday;
} Person;

int main() {
    Person person = {
        .personName = "Bob",
        .birthday = {
            .eventName = "Birthday Party",
            .eventDate = {.day = 20, .month = 5, .year = 1995}
        }
    };
    
    printf("%s's %s is on %d/%d/%d\n",
           person.personName,
           person.birthday.eventName,
           person.birthday.eventDate.day,
           person.birthday.eventDate.month,
           person.birthday.eventDate.year);
    
    return 0;
}
```

---

## Arrays of Structures

You can create arrays of structures to store multiple records.

### Basic Array of Structures

```c
#include <stdio.h>
#include <string.h>

struct Student {
    char name[50];
    int rollNo;
    float marks;
};

int main() {
    struct Student students[3];
    
    // Assign values
    strcpy(students[0].name, "Alice");
    students[0].rollNo = 1;
    students[0].marks = 85.5;
    
    strcpy(students[1].name, "Bob");
    students[1].rollNo = 2;
    students[1].marks = 90.0;
    
    strcpy(students[2].name, "Charlie");
    students[2].rollNo = 3;
    students[2].marks = 78.5;
    
    // Display
    for (int i = 0; i < 3; i++) {
        printf("Student %d: %s (Roll No: %d) - Marks: %.1f\n",
               i + 1,
               students[i].name,
               students[i].rollNo,
               students[i].marks);
    }
    
    return 0;
}
```

**Output:**
```
Student 1: Alice (Roll No: 1) - Marks: 85.5
Student 2: Bob (Roll No: 2) - Marks: 90.0
Student 3: Charlie (Roll No: 3) - Marks: 78.5
```

### Initialize Array at Declaration

```c
#include <stdio.h>

typedef struct {
    char name[30];
    int age;
} Person;

int main() {
    Person people[3] = {
        {"Alice", 25},
        {"Bob", 30},
        {"Charlie", 28}
    };
    
    for (int i = 0; i < 3; i++) {
        printf("%s is %d years old\n", people[i].name, people[i].age);
    }
    
    return 0;
}
```

### Example: Student Grade System

```c
#include <stdio.h>
#include <string.h>

typedef struct {
    char name[50];
    int rollNo;
    float marks;
    char grade;
} Student;

char calculateGrade(float marks) {
    if (marks >= 90) return 'A';
    else if (marks >= 80) return 'B';
    else if (marks >= 70) return 'C';
    else if (marks >= 60) return 'D';
    else return 'F';
}

int main() {
    Student students[5];
    int n = 5;
    
    // Input student data
    for (int i = 0; i < n; i++) {
        printf("\nEnter details for student %d:\n", i + 1);
        printf("Name: ");
        scanf("%s", students[i].name);
        printf("Roll Number: ");
        scanf("%d", &students[i].rollNo);
        printf("Marks: ");
        scanf("%f", &students[i].marks);
        
        students[i].grade = calculateGrade(students[i].marks);
    }
    
    // Display results
    printf("\n=== Student Grade Report ===\n");
    printf("%-20s %-10s %-10s %-10s\n", "Name", "Roll No", "Marks", "Grade");
    printf("---------------------------------------------------\n");
    
    for (int i = 0; i < n; i++) {
        printf("%-20s %-10d %-10.1f %-10c\n",
               students[i].name,
               students[i].rollNo,
               students[i].marks,
               students[i].grade);
    }
    
    return 0;
}
```

---

## Structures and Pointers

Pointers to structures allow efficient passing of structures to functions and dynamic memory allocation.

### Declaring Structure Pointers

```c
struct Student {
    char name[50];
    int age;
};

int main() {
    struct Student s1 = {"Alice", 20};
    struct Student *ptr;
    
    ptr = &s1;  // Pointer points to s1
    
    return 0;
}
```

### Accessing Members via Pointers

**Method 1: Arrow Operator (`->`)**

```c
#include <stdio.h>

struct Point {
    int x;
    int y;
};

int main() {
    struct Point p1 = {10, 20};
    struct Point *ptr = &p1;
    
    // Access using arrow operator (preferred)
    printf("x = %d\n", ptr->x);
    printf("y = %d\n", ptr->y);
    
    // Modify values
    ptr->x = 30;
    ptr->y = 40;
    
    printf("New x = %d\n", ptr->x);
    printf("New y = %d\n", ptr->y);
    
    return 0;
}
```

**Method 2: Dereference Operator**

```c
#include <stdio.h>

struct Point {
    int x;
    int y;
};

int main() {
    struct Point p1 = {10, 20};
    struct Point *ptr = &p1;
    
    // Access using dereference (less common)
    printf("x = %d\n", (*ptr).x);
    printf("y = %d\n", (*ptr).y);
    
    return 0;
}
```

### Complete Example

```c
#include <stdio.h>
#include <string.h>

typedef struct {
    char title[100];
    char author[50];
    float price;
} Book;

void displayBook(Book *b) {
    printf("Title: %s\n", b->title);
    printf("Author: %s\n", b->author);
    printf("Price: $%.2f\n", b->price);
}

void updatePrice(Book *b, float newPrice) {
    b->price = newPrice;
}

int main() {
    Book book1 = {"C Programming", "Dennis Ritchie", 29.99};
    
    printf("Original book details:\n");
    displayBook(&book1);
    
    updatePrice(&book1, 34.99);
    
    printf("\nAfter price update:\n");
    displayBook(&book1);
    
    return 0;
}
```

**Output:**
```
Original book details:
Title: C Programming
Author: Dennis Ritchie
Price: $29.99

After price update:
Title: C Programming
Author: Dennis Ritchie
Price: $34.99
```

---

## Structures and Functions

Structures can be passed to functions in three ways:

### 1. Pass by Value (Copy)

The entire structure is copied - changes don't affect the original.

```c
#include <stdio.h>

typedef struct {
    int x;
    int y;
} Point;

void displayPoint(Point p) {
    printf("Point: (%d, %d)\n", p.x, p.y);
}

void modifyPoint(Point p) {
    p.x = 100;  // Changes only local copy
    p.y = 200;
}

int main() {
    Point p1 = {10, 20};
    
    displayPoint(p1);
    modifyPoint(p1);
    displayPoint(p1);  // Still (10, 20) - unchanged
    
    return 0;
}
```

**Output:**
```
Point: (10, 20)
Point: (10, 20)
```

### 2. Pass by Reference (Pointer)

Only the address is passed - more efficient, changes affect the original.

```c
#include <stdio.h>

typedef struct {
    int x;
    int y;
} Point;

void displayPoint(Point *p) {
    printf("Point: (%d, %d)\n", p->x, p->y);
}

void modifyPoint(Point *p) {
    p->x = 100;  // Changes original
    p->y = 200;
}

int main() {
    Point p1 = {10, 20};
    
    displayPoint(&p1);
    modifyPoint(&p1);
    displayPoint(&p1);  // Now (100, 200) - changed!
    
    return 0;
}
```

**Output:**
```
Point: (10, 20)
Point: (100, 200)
```

### 3. Return Structure from Function

```c
#include <stdio.h>
#include <string.h>

typedef struct {
    char name[50];
    int age;
    float salary;
} Employee;

Employee createEmployee(char name[], int age, float salary) {
    Employee emp;
    strcpy(emp.name, name);
    emp.age = age;
    emp.salary = salary;
    return emp;
}

void displayEmployee(Employee emp) {
    printf("Name: %s\n", emp.name);
    printf("Age: %d\n", emp.age);
    printf("Salary: $%.2f\n", emp.salary);
}

int main() {
    Employee emp1 = createEmployee("John Doe", 30, 50000.00);
    displayEmployee(emp1);
    
    return 0;
}
```

### Example: Rectangle Operations

```c
#include <stdio.h>

typedef struct {
    float length;
    float width;
} Rectangle;

Rectangle createRectangle(float l, float w) {
    Rectangle r = {l, w};
    return r;
}

float calculateArea(Rectangle *r) {
    return r->length * r->width;
}

float calculatePerimeter(Rectangle *r) {
    return 2 * (r->length + r->width);
}

void displayRectangle(Rectangle *r) {
    printf("Rectangle: %.2f x %.2f\n", r->length, r->width);
    printf("Area: %.2f\n", calculateArea(r));
    printf("Perimeter: %.2f\n", calculatePerimeter(r));
}

int main() {
    Rectangle rect1 = createRectangle(5.0, 3.0);
    displayRectangle(&rect1);
    
    return 0;
}
```

**Output:**
```
Rectangle: 5.00 x 3.00
Area: 15.00
Perimeter: 16.00
```

---

## Structure Padding and Alignment

Compilers add padding bytes between structure members for memory alignment optimization.

### Memory Layout Example

```c
#include <stdio.h>

struct Example1 {
    char a;    // 1 byte
    int b;     // 4 bytes
    char c;    // 1 byte
};

struct Example2 {
    char a;    // 1 byte
    char c;    // 1 byte
    int b;     // 4 bytes
};

int main() {
    printf("Size of Example1: %zu bytes\n", sizeof(struct Example1));
    printf("Size of Example2: %zu bytes\n", sizeof(struct Example2));
    
    return 0;
}
```

**Output (typical):**
```
Size of Example1: 12 bytes
Size of Example2: 8 bytes
```

### Understanding Padding

```
Example1 (12 bytes):
char a    [1 byte]
padding   [3 bytes]  <- Padding added
int b     [4 bytes]
char c    [1 byte]
padding   [3 bytes]  <- Padding added

Example2 (8 bytes):
char a    [1 byte]
char c    [1 byte]
padding   [2 bytes]  <- Padding added
int b     [4 bytes]
```

### Tips to Reduce Padding

1. **Order members by size** (largest to smallest)
2. **Group similar types together**

```c
// ❌ BAD: More padding
struct Bad {
    char a;
    double b;
    char c;
    int d;
};

// ✅ GOOD: Less padding
struct Good {
    double b;  // 8 bytes
    int d;     // 4 bytes
    char a;    // 1 byte
    char c;    // 1 byte
};
```

---

## Practical Examples

### Example 1: Library Management System

```c
#include <stdio.h>
#include <string.h>

typedef struct {
    int bookID;
    char title[100];
    char author[50];
    int isIssued;  // 0 = available, 1 = issued
} Book;

void addBook(Book *b, int id, char title[], char author[]) {
    b->bookID = id;
    strcpy(b->title, title);
    strcpy(b->author, author);
    b->isIssued = 0;
}

void issueBook(Book *b) {
    if (b->isIssued) {
        printf("Book already issued!\n");
    } else {
        b->isIssued = 1;
        printf("Book '%s' has been issued.\n", b->title);
    }
}

void returnBook(Book *b) {
    if (!b->isIssued) {
        printf("Book is already available!\n");
    } else {
        b->isIssued = 0;
        printf("Book '%s' has been returned.\n", b->title);
    }
}

void displayBook(Book *b) {
    printf("\n--- Book Details ---\n");
    printf("ID: %d\n", b->bookID);
    printf("Title: %s\n", b->title);
    printf("Author: %s\n", b->author);
    printf("Status: %s\n", b->isIssued ? "Issued" : "Available");
}

int main() {
    Book library[3];
    
    addBook(&library[0], 101, "C Programming", "Dennis Ritchie");
    addBook(&library[1], 102, "Data Structures", "Tanenbaum");
    addBook(&library[2], 103, "Algorithms", "Cormen");
    
    displayBook(&library[0]);
    issueBook(&library[0]);
    displayBook(&library[0]);
    returnBook(&library[0]);
    displayBook(&library[0]);
    
    return 0;
}
```

### Example 2: Student Database

```c
#include <stdio.h>
#include <string.h>

typedef struct {
    int id;
    char name[50];
    float marks[3];  // 3 subjects
    float average;
    char grade;
} Student;

void calculateAverage(Student *s) {
    float sum = 0;
    for (int i = 0; i < 3; i++) {
        sum += s->marks[i];
    }
    s->average = sum / 3.0;
}

void assignGrade(Student *s) {
    if (s->average >= 90) s->grade = 'A';
    else if (s->average >= 80) s->grade = 'B';
    else if (s->average >= 70) s->grade = 'C';
    else if (s->average >= 60) s->grade = 'D';
    else s->grade = 'F';
}

void displayStudent(Student *s) {
    printf("\n--- Student Report ---\n");
    printf("ID: %d\n", s->id);
    printf("Name: %s\n", s->name);
    printf("Marks: %.1f, %.1f, %.1f\n", 
           s->marks[0], s->marks[1], s->marks[2]);
    printf("Average: %.2f\n", s->average);
    printf("Grade: %c\n", s->grade);
}

int main() {
    Student student1;
    
    student1.id = 1001;
    strcpy(student1.name, "Alice Johnson");
    student1.marks[0] = 85.0;
    student1.marks[1] = 90.0;
    student1.marks[2] = 88.0;
    
    calculateAverage(&student1);
    assignGrade(&student1);
    displayStudent(&student1);
    
    return 0;
}
```

### Example 3: Bank Account System

```c
#include <stdio.h>
#include <string.h>

typedef struct {
    int accountNumber;
    char accountHolder[50];
    float balance;
} BankAccount;

void createAccount(BankAccount *acc, int accNo, char name[], float initialBalance) {
    acc->accountNumber = accNo;
    strcpy(acc->accountHolder, name);
    acc->balance = initialBalance;
    printf("Account created for %s with balance $%.2f\n", name, initialBalance);
}

void deposit(BankAccount *acc, float amount) {
    if (amount > 0) {
        acc->balance += amount;
        printf("Deposited $%.2f. New balance: $%.2f\n", amount, acc->balance);
    } else {
        printf("Invalid deposit amount!\n");
    }
}

void withdraw(BankAccount *acc, float amount) {
    if (amount > 0 && amount <= acc->balance) {
        acc->balance -= amount;
        printf("Withdrew $%.2f. New balance: $%.2f\n", amount, acc->balance);
    } else {
        printf("Invalid withdrawal or insufficient funds!\n");
    }
}

void checkBalance(BankAccount *acc) {
    printf("Account Holder: %s\n", acc->accountHolder);
    printf("Account Number: %d\n", acc->accountNumber);
    printf("Current Balance: $%.2f\n", acc->balance);
}

int main() {
    BankAccount account1;
    
    createAccount(&account1, 123456, "John Doe", 1000.00);
    deposit(&account1, 500.00);
    withdraw(&account1, 200.00);
    checkBalance(&account1);
    
    return 0;
}
```

**Output:**
```
Account created for John Doe with balance $1000.00
Deposited $500.00. New balance: $1500.00
Withdrew $200.00. New balance: $1300.00
Account Holder: John Doe
Account Number: 123456
Current Balance: $1300.00
```

### Example 4: Complex Number Operations

```c
#include <stdio.h>

typedef struct {
    float real;
    float imag;
} Complex;

Complex addComplex(Complex c1, Complex c2) {
    Complex result;
    result.real = c1.real + c2.real;
    result.imag = c1.imag + c2.imag;
    return result;
}

Complex multiplyComplex(Complex c1, Complex c2) {
    Complex result;
    result.real = c1.real * c2.real - c1.imag * c2.imag;
    result.imag = c1.real * c2.imag + c1.imag * c2.real;
    return result;
}

void displayComplex(Complex c) {
    if (c.imag >= 0) {
        printf("%.2f + %.2fi\n", c.real, c.imag);
    } else {
        printf("%.2f - %.2fi\n", c.real, -c.imag);
    }
}

int main() {
    Complex c1 = {3.0, 4.0};
    Complex c2 = {1.0, 2.0};
    
    printf("c1 = ");
    displayComplex(c1);
    
    printf("c2 = ");
    displayComplex(c2);
    
    Complex sum = addComplex(c1, c2);
    printf("c1 + c2 = ");
    displayComplex(sum);
    
    Complex product = multiplyComplex(c1, c2);
    printf("c1 * c2 = ");
    displayComplex(product);
    
    return 0;
}
```

### Example 5: Employee Management

```c
#include <stdio.h>
#include <string.h>

typedef struct {
    int day;
    int month;
    int year;
} Date;

typedef struct {
    int empID;
    char name[50];
    char department[30];
    float salary;
    Date joinDate;
} Employee;

void displayEmployee(Employee *emp) {
    printf("\n=== Employee Details ===\n");
    printf("ID: %d\n", emp->empID);
    printf("Name: %s\n", emp->name);
    printf("Department: %s\n", emp->department);
    printf("Salary: $%.2f\n", emp->salary);
    printf("Join Date: %02d/%02d/%d\n", 
           emp->joinDate.day, 
           emp->joinDate.month, 
           emp->joinDate.year);
}

void giveRaise(Employee *emp, float percentage) {
    float raise = emp->salary * (percentage / 100.0);
    emp->salary += raise;
    printf("%s received a %.1f%% raise of $%.2f\n", 
           emp->name, percentage, raise);
}

int main() {
    Employee employees[3] = {
        {101, "Alice Smith", "Engineering", 75000, {15, 3, 2020}},
        {102, "Bob Johnson", "Marketing", 65000, {20, 6, 2021}},
        {103, "Charlie Brown", "Sales", 70000, {10, 1, 2022}}
    };
    
    for (int i = 0; i < 3; i++) {
        displayEmployee(&employees[i]);
    }
    
    printf("\n--- Giving Raises ---\n");
    giveRaise(&employees[0], 10.0);
    giveRaise(&employees[1], 8.0);
    
    printf("\n--- After Raises ---\n");
    displayEmployee(&employees[0]);
    displayEmployee(&employees[1]);
    
    return 0;
}
```

---

## Common Pitfalls

### 1. Forgetting Semicolon After Structure Definition

```c
// ❌ BAD: Missing semicolon
struct Student {
    char name[50];
    int age;
}  // ERROR: Missing semicolon

// ✅ GOOD
struct Student {
    char name[50];
    int age;
};
```

### 2. Not Using strcpy for Strings

```c
#include <string.h>

struct Person {
    char name[50];
};

int main() {
    struct Person p1;
    
    // ❌ BAD: Can't assign strings directly
    p1.name = "Alice";  // ERROR!
    
    // ✅ GOOD: Use strcpy
    strcpy(p1.name, "Alice");
    
    return 0;
}
```

### 3. Comparing Structures Directly

```c
// ❌ BAD: Can't compare structures directly
if (struct1 == struct2) { }  // ERROR!

// ✅ GOOD: Compare members individually
if (strcmp(struct1.name, struct2.name) == 0 && 
    struct1.age == struct2.age) { }
```

### 4. Not Initializing Structure Members

```c
struct Point {
    int x;
    int y;
};

int main() {
    // ❌ BAD: Uninitialized (garbage values)
    struct Point p1;
    printf("%d, %d\n", p1.x, p1.y);  // Unpredictable output
    
    // ✅ GOOD: Initialize
    struct Point p2 = {0, 0};
    // Or
    struct Point p3 = {0};  // All members set to 0
    
    return 0;
}
```

### 5. Incorrect Pointer Syntax

```c
struct Point {
    int x;
    int y;
};

int main() {
    struct Point p1 = {10, 20};
    struct Point *ptr = &p1;
    
    // ❌ BAD: Using dot with pointer
    printf("%d\n", ptr.x);  // ERROR!
    
    // ✅ GOOD: Use arrow operator
    printf("%d\n", ptr->x);
    
    // ✅ ALSO GOOD: Dereference first
    printf("%d\n", (*ptr).x);
    
    return 0;
}
```

### 6. Returning Local Structure Address

```c
// ❌ BAD: Returning address of local variable
struct Point* createPoint() {
    struct Point p = {10, 20};
    return &p;  // DANGER: p is destroyed after function returns
}

// ✅ GOOD: Return by value
struct Point createPoint() {
    struct Point p = {10, 20};
    return p;
}

// ✅ ALSO GOOD: Use dynamic allocation
struct Point* createPoint() {
    struct Point *p = malloc(sizeof(struct Point));
    p->x = 10;
    p->y = 20;
    return p;  // Remember to free later!
}
```

---

## Best Practices

1. **Use typedef** to avoid writing `struct` keyword repeatedly
2. **Use meaningful structure names** (PascalCase is common)
3. **Group related data** into structures
4. **Pass structures by pointer** to functions for efficiency
5. **Initialize all structure members** to avoid garbage values
6. **Order structure members** by size to reduce padding
7. **Use const** for read-only structure parameters
8. **Comment complex structures** to explain their purpose
9. **Keep structures focused** - don't make them too large
10. **Consider alignment** for performance-critical code

```c
// ✅ GOOD: Well-designed structure
typedef struct {
    int id;              // Unique identifier
    char name[50];       // Student name
    float gpa;           // Grade point average
    int graduationYear;  // Year of graduation
} Student;

// ✅ GOOD: Pass by const pointer when not modifying
void displayStudent(const Student *s) {
    printf("Name: %s, GPA: %.2f\n", s->name, s->gpa);
}

// ✅ GOOD: Clear function purpose
Student createStudent(int id, char name[], float gpa) {
    Student s;
    s.id = id;
    strcpy(s.name, name);
    s.gpa = gpa;
    return s;
}
```

---

## Quick Reference

```c
// Define structure
struct StructName {
    datatype member1;
    datatype member2;
};

// With typedef
typedef struct {
    datatype member1;
    datatype member2;
} StructName;

// Declare variable
struct StructName var1;
StructName var2;  // If using typedef

// Access members
var1.member1 = value;

// Pointer to structure
struct StructName *ptr = &var1;
ptr->member1 = value;        // Arrow operator
(*ptr).member1 = value;      // Dereference

// Initialize
struct StructName var = {val1, val2};
struct StructName var = {.member1 = val1, .member2 = val2};

// Array of structures
struct StructName arr[10];

// Nested structures
struct Outer {
    struct Inner nested;
};

// Pass to function
void func(struct StructName s);        // By value
void func(struct StructName *s);       // By pointer (preferred)

// Return from function
struct StructName func() {
    struct StructName s;
    return s;
}
```

Structures are the foundation of organizing complex data in C - master them to write clean, maintainable code!
