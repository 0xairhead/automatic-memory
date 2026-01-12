# Phase 8: Modules & Dependency Management

## 1. go.mod & go.sum
*   **go.mod**: The manifest. Defines usage of semantic versioning (`v1.2.3`).
*   **go.sum**: Checksums. Ensures everyone gets the exact same bits (security).
*   `go get <pkg>`: Adds dependency.
*   `go mod tidy`: Prunes unused dependencies.

## 2. Local Replacements
*   `replace example.com/pkg => ../local/path`
*   Allowed working on a local fork or multi-module repo without pushing to Git.

## 3. Vendoring
*   `go mod vendor`: Copies all dependency code into a `vendor/` folder.
*   Protcts against "LeftPad" incidents (dependencies disappearing) and allows offline builds.
