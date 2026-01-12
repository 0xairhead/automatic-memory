# Phase 11: Cloud-Native Go

## 1. 12-Factor Apps
*   Go is ideal for containers (single binary, fast startup).
*   **Signals**: Handling `SIGTERM` is mandatory for Kubernetes rolling updates.

## 2. Health Checks
*   `/healthz` endpoint.
*   **Liveness**: "Am I running?" (If no, restart me).
*   **Readiness**: "Can I take traffic?" (If no, remove from Load Balancer).

## 3. Observability
*   **Structured Logging**: Text logs (`printf`) are hard to parse. JSON logs (`slog`) are machine-queryable.
*   **Metrics**: Prometheus exposition (counters, gauges).

## 4. CLI Tools (Cobra)
*   **Command Structure**: `app [command] [subcommand] --flag`.
*   Cobra provides help generation, flag parsing, and command routing.
