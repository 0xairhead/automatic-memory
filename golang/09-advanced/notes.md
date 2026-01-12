# Phase 9: Advanced Go

## 1. Generics (Go 1.18+)
*   `func Fn[T any](input T) T`.
*   Enables writing algorithms (Map, Filter, Stack) once for all types.
*   **Constraints**: `comparable`, `int | float64`.

## 2. Reflection
*   `reflect` package allows inspection of types/values at runtime.
*   Used heavily by JSON libraries and ORMs.
*   **Performance cost**: Slower than static code. Use sparingly.

## 3. Unsafe
*   `unsafe.Pointer`: Bypasses Go's type system.
*   Allows reading raw memory bits or casting structs arbitrarily.
*   **Danger**: Can causing segfaults or data corruption if Garbage Collector moves memory.
