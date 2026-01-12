# Phase 3: Pointers, Memory & Errors

## 1. Pointers
*   `&x` gets the address. `*p` reads the value at address.
*   **Value vs Reference**: Go defaults to pass-by-value. Use pointers to modify variables across function boundaries or avoid copying large structs.
*   **Nil**: The zero value for pointers. Always check `if p != nil` before dereferencing if unsure.

## 2. Error Handling
*   **No Exceptions**: Go treats errors as values.
*   **Idion**: `if err != nil { return err }`.
*   **Wrapping**: `fmt.Errorf("context: %w", err)` preserves the original error for `errors.Is` and `errors.As`.

## 3. Panic & Recover
*   **Panic**: Crashes the program. Use only for unrecoverable state (e.g., config missing on startup).
*   **Recover**: A `defer` function can "catch" a panic preventing a crash. Used in servers to keep running even if one handler crashes.
