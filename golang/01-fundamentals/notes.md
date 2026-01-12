# Phase 1: Go Fundamentals

## 1. Syntax Basics
*   **Package `main`**: The entry point of every executable program.
*   **Imports**: How we include other code. `import "fmt"` is standard.
*   **Variables**:
    *   `var x int = 10` (Explicit)
    *   `y := 20` (Short declaration, type inferred)

## 2. Types
Go is statically typed. Common types:
*   `int`, `float64` (default float size)
*   `string` (immutable)
*   `bool` (true/false)

## 3. Control Flow
*   **For Loops**: The *only* loop in Go.
    *   `for i := 0; i < 10; i++ {}` (Standard)
    *   `for condition {}` (Like a `while`)
    *   `for {}` (Infinite loop)
*   **If/Else**: Standard, but can include a short statement: `if x := fn(); x < 0 { ... }`
*   **Switch**: No `break` needed (implicit). Use `fallthrough` if you really need it.

## 4. Functions
*   **Multiple Returns**: `func swap(a, b int) (int, int)` is idiomatic.
*   **Named Returns**: `func split(sum int) (x, y int)` - defines variables at top.
