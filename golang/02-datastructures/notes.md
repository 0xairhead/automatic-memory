# Phase 2: Core Data Structures

## 1. Arrays vs Slices
*   **Array**: Fixed size (`[5]int`). Value semantics (copy on pass).
*   **Slice**: Dynamic view (`[]int`). Pointer semantics (cheap to pass).
    *   `make([]int, len, cap)`: Allocates underlying array.
    *   `append(slice, val)`: Adds elements, may trigger reallocation if capacity is exceeded.

## 2. Maps
*   Hash tables `map[keyType]valType`.
*   Iterating: `range` returns keys in random order.
*   Check existence: `val, ok := m["key"]`. If `!ok`, key is missing.

## 3. Structs
*   Custom data types grouping fields.
*   **Tags**: Metadata like `` `json:"name"` `` used by reflection libraries.
*   **Embedding**: Go's pseudo-inheritance. `type Admin struct { User }` allows Admin to access User fields directly.

## 4. Range
The universal iterator:
*   Slice: `for index, value := range slice`
*   Map: `for key, value := range map`
