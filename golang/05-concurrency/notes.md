# Phase 5: Concurrency

## 1. Goroutines
*   Lightweight threads managed by the Go Runtime (not OS threads).
*   `go fn()` starts execution immediately.
*   **WaitGroup**: Used to wait for a collection of goroutines to finish. `wg.Add(1)`, `wg.Done()`, `wg.Wait()`.

## 2. Channels
*   Typed conduits for passing data between goroutines.
*   **Unbuffered**: Blocks sender until receiver is ready (synchronization).
*   **Buffered**: Blocks sender only when full.
*   `close(ch)`: Signals no more values. Receiver sees zero value + `false` ok flag.

## 3. Select Statement
*   Let's a goroutine wait on multiple communication operations.
*   Used for timeouts (`time.After`) and cancellation (`ctx.Done`).

## 4. Sync & Context
*   **Mutex**: `mu.Lock()` protects shared memory (Maps are not thread-safe!).
*   **Context**: Propagates deadlines, cancellation signals, and request-scoped values across API boundaries.
