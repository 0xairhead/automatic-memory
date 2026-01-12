# Phase 10: Security-Focused Go

## 1. Safe Coding
*   **Path Traversal**: Never trust input in file paths. Use `filepath.Clean` and check `strings.HasPrefix`.
*   **Input Validation**: Always validate data at the edge.

## 2. Hashing (bcrypt)
*   Do NOT save plain text passwords.
*   Do NOT plain SHA256 passwords (susceptible to Rainbow Tables).
*   Use **bcrypt** (or Argon2). It includes Salting + Slow Hashing automatically.

## 3. JWT (JSON Web Tokens)
*   Stateless authentication.
*   **Structure**: Header.Payload.Signature.
*   **Signing**: Server keeps a Secret Key. If signature matches, the server knows *it* issued the token and data is untampered.

## 4. Static Analysis
*   `go/ast` allows parsing Go source code programmatically.
*   Used to build custom linters (e.g., "Ban usage of 'unsafe' package in this repo").
