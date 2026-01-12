# Phase 0: Setup & Go Mindset

## 1. Why Go?
*   **Simplicity**: Small language spec. Easy to read.
*   **Performance**: Compiled to machine code. Fast.
*   **Concurrency**: Goroutines and Channels make async work easy (Phase 5).
*   **Tooling**: The standard library and CLI tools (`go fmt`, `go test`) are world-class.

## 2. Workspace: Modules vs GOPATH
*   **Old Way (GOPATH)**: All code had to live in a single massive workspace directory (`~/go/src/github.com/user/project`).
*   **New Way (Go Modules)**: Create a folder anywhere. Run `go mod init <module-name>`.
    *   This creates a `go.mod` file which tracks dependencies (like `package.json` in Node or `requirements.txt` in Python).

## 3. Essential Commands
We will practice these in this folder.

*   `go run main.go`: Compiles and runs your code immediately. Does not create a binary file.
*   `go fmt`: Automatically formats your code (indentation, spacing). **Always run this**.
*   `go vet`: Checks for suspicious constructs (bugs that compile but might crash or be wrong).
*   `go build`: Compiles your code into an executable binary.
