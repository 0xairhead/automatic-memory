# Phase 12: Capstone Projects

## 1. XDR Agent Architecture
*   **Agent**:
    *   **Producers**: Concurrent monitors (File, Process) generating alerts.
    *   **Aggregation**: A shared `chan Alert` buffer.
    *   **Consumers**: A Worker Pool of HTTP clients sending alerts to the server.
    *   **Graceful Shutdown**: Essential for agents to finish sending buffer before exiting during updates.

## 2. Server
*   **Ingestion**: High-throughput HTTP endpoint.
*   **Logging**: JSON structured logging for SIEM integration.

## 3. Key Takeaway
This project combines Concurrency (Channels/Workers), Networking (HTTP), OS (Signals), and Architecture (Producers/Consumers) into a single cohesive system.
