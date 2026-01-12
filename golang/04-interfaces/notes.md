# Phase 4: Interfaces & Design Patterns

## 1. Implicit Interfaces
*   **Duck Typing**: If a type implements the methods, it implements the interface. No `implements` keyword.
*   **Decoupling**: Code against behavior (`io.Reader`), not concrete types (`os.File`). This makes testing easy.

## 2. Empty Interface (`interface{}` / `any`)
*   Holds values of any type.
*   **Type Assertion**: `s, ok := val.(string)`. Checks if the interface holds a specific type.
*   **Type Switch**: `switch v := val.(type)` to handle multiple types safely.

## 3. Dependency Injection
*   Pass interfaces (dependencies) into struct constructors or methods.
*   Allows swapping real implementations (e.g., Database) with Mocks during tests.
