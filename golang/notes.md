# Go Course Module Guide

This document explains the contents and purpose of each module in the repository.

## Phase 0: Setup (`00-setup/`)
*   **Purpose**: Initial environment configuration and understanding Go tooling.
*   **Contents**:
    *   `main.go`: A "Hello World" sanity check.
    *   `notes.md`: Basics of Go CLI (`run`, `build`, `fmt`) and module system.

## Phase 1: Fundamentals (`01-fundamentals/`)
*   **Purpose**: Core syntax and control structures.
*   **Contents**:
    *   `syntax/`: Lessons on Variables, Control Flow (`if`/`for`), and Functions.
    *   `hands-on/calculator`: CLI tool demonstrating `switch` and args.
    *   `hands-on/linecounter`: File processing example.
    *   `hands-on/password`: String manipulation and `unicode` package usage.

## Phase 2: Data Structures (`02-datastructures/`)
*   **Purpose**: Mastering Go's built-in collections.
*   **Contents**:
    *   `basics/`: Lessons on Slices, Maps, and Structs.
    *   `hands-on/json_demo`: Parsing JSON data.
    *   `hands-on/config_loader`: Loading configuration from files.
    *   `hands-on/kv_store`: An interactive Key-Value store CLI.

## Phase 3: Pointers & Errors (`03-pointers-errors/`)
*   **Purpose**: Memory management and robust error handling.
*   **Contents**:
    *   `basics/`: Concept of Pointers, generic `error` interface, and `panic`/`recover`.
    *   `hands-on/custom_error`: Implementing custom error types w/ context.
    *   `hands-on/safe_runner`: Using `recover` to prevent crashes.

## Phase 4: Interfaces (`04-interfaces/`)
*   **Purpose**: Polymorphism and clean architecture.
*   **Contents**:
    *   `basics/`: Implicit interface satisfaction and Type Assertions.
    *   `hands-on/payment_system`: Strategy Pattern implementation (CreditCard vs PayPal).
    *   `hands-on/pluggable_logger`: Dependency Injection example.

## Phase 5: Concurrency (`05-concurrency/`)
*   **Purpose**: Go's killer feature: Goroutines and Channels.
*   **Contents**:
    *   `basics/`: Using `go` keyword, `chan`, `sync.WaitGroup`, `sync.Mutex`, and `context`.
    *   `hands-on/worker_pool`: Parallel job processing pattern.
    *   `hands-on/web_scraper`: Concurrent HTTP fetching with aggregation.

## Phase 6: OS & Networking (`06-files-os-net/`)
*   **Purpose**: Interacting with the system and building network services.
*   **Contents**:
    *   `basics/`: File I/O, OS Signals (Ctrl+C), and HTTP Client/Server.
    *   `hands-on/log_analyzer`: File parsing tool.
    *   `hands-on/api_server`: A REST API implementation.

## Phase 7: Testing (`07-testing/`)
*   **Purpose**: Ensuring code quality and performance.
*   **Contents**:
    *   `basics/mathutils`: Unit testing and Table-Driven Tests.
    *   `basics/benchmarking`: Performance comparison (`strings.Builder` vs `+`).
    *   `hands-on/cache`: A Thread-Safe Cache with Race Condition testing.

## Phase 8: Modules (`08-modules/`)
*   **Purpose**: Dependency management.
*   **Contents**:
    *   `app/`: Application importing external and local modules.
    *   Demonstrated `go get`, `replace` directives, and `vendor` folders.

## Phase 9: Advanced Go (`09-advanced/`)
*   **Purpose**: Capabilities beyond standard static typing.
*   **Contents**:
    *   `basics/generics`: Generic functions and types (`[T any]`).
    *   `basics/reflection`: inspecting Struct Tags at runtime.
    *   `basics/unsafe`: Direct memory manipulation.
    *   `hands-on/validator`: A custom struct validation library using reflection.

## Phase 10: Security (`10-security/`)
*   **Purpose**: Writing secure code and security tooling.
*   **Contents**:
    *   `basics/safe_input`: Path Traversal prevention.
    *   `hands-on/password_tool`: Hashing with `bcrypt`.
    *   `hands-on/jwt_auth`: JWT issuance and validation.
    *   `hands-on/static_scanner`: AST analysis to detect dangerous code imports.

## Phase 11: Cloud-Native (`11-cloud-native/`)
*   **Purpose**: Production-ready patterns for Cloud/K8s.
*   **Contents**:
    *   `basics/health`: Health checks and Graceful Shutdown.
    *   `hands-on/k8s-cli`: CLI tool built with `Cobra`.
    *   `hands-on/microservice_logs`: Service using structured logging (`slog`).

## Phase 12: Capstones (`12-capstones/`)
*   **Purpose**: Final integration project.
*   **Contents**:
    *   `xdr-agent/`: An Extended Detection and Response prototype.
        *   `agent/`: Concurrent host monitor (File + Process).
        *   `server/`: Central log aggregation server.
