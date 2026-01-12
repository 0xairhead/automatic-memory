# Phase 7: Testing & Quality

## 1. The `testing` Package
*   Files ending in `_test.go` are compiled only during `go test`.
*   Functions `TestXxx(t *testing.T)` verify logic. `t.Error` logs failure but keeps going; `t.Fatal` stops immediately.

## 2. Table-Driven Tests
*   Go Pattern: Define a slice of structs `{input, expected}` and iterate.
*   Makes adding test cases trivial and keeps test logic DRY.
*   `t.Run(name, func)` creates subtests for better reporting.

## 3. Benchmarking
*   `BenchmarkXxx(b *testing.B)` measures performance.
*   Loop `for i := 0; i < b.N; i++` runs the code enough times to get statistical significance.
*   `go test -bench=. -benchmem` shows time/op and allocs/op.

## 4. Race Detection
*   `go test -race` instruments code to find concurrent memory access. Essential for verifying Goroutine safety.
