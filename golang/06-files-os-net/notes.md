# Phase 6: Files, OS & Networking

## 1. File Handling
*   `os.ReadFile` / `os.WriteFile`: Quick, all-at-once I/O.
*   `os.Open` + `bufio.Scanner`: Efficient, line-by-line reading for large files.
*   **Permissions**: `0644` (rw-r--r--) is standard for text files.

## 2. OS Signals
*   Programs don't just stop; they receive signals like `SIGINT` (Ctrl+C).
*   `signal.Notify(chan, os.Interrupt)` lets you capture this to close database connections or save state before exiting.

## 3. Networking
*   **net/http**: Production-grade Client and Server.
*   **Server**: `http.HandleFunc`. Runs each request in a separate goroutine automatically.
*   **Client**: `http.Get`. IMPORTANT: Always `defer resp.Body.Close()` to prevent file descriptor leaks.
