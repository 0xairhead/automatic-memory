# Advanced Software Fuzzing: Scalability and Engineering

This syllabus outlines a comprehensive curriculum for advanced software fuzzing, with a specific focus on **distributed systems** and **production-scale engineering**. The curriculum transitions from fundamental security research concepts to the architectural challenges of running fuzzing campaigns across large-scale compute clusters.

---

## Phase 0: Prerequisites
**Objective:** Establish a baseline understanding of low-level systems and tooling.

*   **Languages:** C/C++, Python scripting
*   **Operating Systems:** Linux kernel internals, process management
*   **Tooling:** Debuggers (GDB/LLDB), build systems (Make/CMake)
*   **Security:** Common memory corruption vulnerability classes

---

## [Phase 1: Fuzzing Fundamentals](01-fuzzing-fundamentals.md)
**Objective:** Understand the limitations of single-node fuzzing and the necessity of distributed architectures.

*   Computational bounds of deep state expiration
*   Coverage plateaus and logarithmic returns
*   Corpus management challenges
*   Non-determinism and crash reproducibility at scale

---

## [Phase 2: Mutation-Based Fuzzing](02-mutation-fuzzing.md)
**Objective:** Master input generation strategies and their impact on throughput and coverage.

*   **Mutation Strategies:** Bit-flipping, arithmetic operations, glossary injection
*   **Seed Optimization:** Minimization techniques for high-throughput execution
*   **Determinism:** Pseudo-Random Number Generator (PRNG) usage for consistent reproduction
*   **Scheduling:** Prioritization of high-velocity inputs

---

## [Phase 3: Coverage-Guided Fuzzing](03-coverage-fuzzing.md)
**Objective:** Implement feedback-driven fuzzing loops using compile-time instrumentation.

*   **Mechanisms:** Instrumentation, shared memory bitmaps, and hit counts
*   **Optimization:** Preventing path explosion (bucketing)
*   **Saturation:** Metrics for determining campaign efficiency
*   **Synchronization:** Distributed syncing protocols (Main/Secondary nodes)
*   **Tools:** AFL++ architecture and deployment

---

## [Phase 4: Harness Engineering for Scale](04-harness-engineering.md)
**Objective:** Optimize target integration for maximum execution speed and stability.

*   **Interfaces:** `LLVMFuzzerTestOneInput` standard
*   **Performance:** Persistent mode vs. Fork server models
*   **Isolation:** Managing global state and side effects
*   **Mocking:** Abstracting network, disk, and high-latency subsystems

---

## [Phase 5: Sanitizers & Crash Triage](05-sanitizers-triage.md)
**Objective:** Automate the detection, deduplication, and prioritization of software faults.

*   **Detection:** AddressSanitizer (ASAN), MemorySanitizer (MSAN), UndefinedBehaviorSanitizer (UBSAN)
*   **Triage:** Automated stack trace hashing and signature generation
*   **Deduplication:** Handling crash storms and identifying unique root causes
*   **Minimization:** Automated reduction of crash inputs (`afl-tmin`)

---

## Phase 6: Advanced Fuzzing Techniques
**Objective:** Apply specialized techniques for complex targets.

*   Context-aware grammar fuzzing
*   Hybrid fuzzing (Concolic execution + Fuzzing)
*   Targeted fuzzing for specific code paths

---

## Phase 7: Fuzzing at Scale (Distributed Architecture)
**Objective:** Architect and deploy continuous fuzzing infrastructure.

### 7.1 System Architecture
*   Master/Worker design patterns
*   Centralized vs. Distributed corpus management
*   Scheduling algorithms (Push vs. Pull)

### 7.2 Cloud Infrastructure
*   Kubernetes deployments and containerization
*   Preemptible/Spot instance orchestration
*   Security isolation (gVisor, Kata Containers)

### 7.3 Continuous Integration
*   Regression fuzzing pipelines
*   Pull Request integration (CI/CD blocking)
*   Corpus management strategies (Backup, Sync, Pruning)

### 7.4 Observability
*   Prometheus/Grafana metrics for fuzzing campaigns
*   Alerting thresholds (Stalls, Coverage regressions)

---

## Phase 8: Kernel Fuzzing
**Objective:** Fuzzing operating system kernels.

*   Syzkaller architecture
*   Virtual Machine snapshotting and restore
*   System call fuzzing strategies

---

## Phase 9: Network & Protocol Fuzzing
**Objective:** Fuzzing stateful network services.

*   Stateful protocol modeling
*   Snapshot-based fuzzing
*   TLS/Encryption bypass strategies

---

## Phase 10: Lifecycle Management
**Objective:** Managing the software security lifecycle.

*   Vulnerability disclosure processes
*   SLA management for fuzzing findings
*   Risk assessment and prioritization

---

# Specialization Track: Browser Fuzzing
**Objective:** A deep-dive into the complex world of browser security, from DOM rendering to JIT compilers.

## Module 1: Introduction to Browser Fuzzing
*   Modern Browser Architecture (Multiprocess, Sandbox)
*   Setting up a compiled debug environment (Chromium/Firefox)
*   Fuzzing Workflow Automation

## Module 2: Fuzzing DOM & Rendering Engines
*   **Targets:** Blink (Chrome), Gecko (Firefox), WebKit (Safari)
*   HTML/CSS/XML Parsing logic
*   DOM rendering & layout engine fuzzing
*   Grammar-based fuzzing for DOM structures
*   Analysis of previous RCEs in rendering engines

## Module 3: Fuzzing JavaScript Engines & JIT
*   **Targets:** V8 (Chrome), SpiderMonkey (Firefox), JavaScriptCore (Safari)
*   JS Engine Internals (Parser, Interpreter, Compiler)
*   **JIT Fuzzing:** TurboFan and IonMonkey
*   Memory management & Garbage Collection bugs
*   Fuzzing APIs and optimizers

## Module 4: Fuzzing WebAssembly (Wasm)
*   Wasm VM Architecture & Implementation
*   Fuzzing Wasm-JS Interface APIs
*   WebAssembly Compiler internals
*   In-process fuzzing of Wasm runtimes

## Module 5: Fuzzing IPC & Components
*   **IPC:** Chrome Mojo / Legacy IPC fuzzing
*   Inter-Process Communication internals and race conditions
*   Fuzzing Media pipelines (Audio/Video decoders)
*   Networking & Data Persistence APIs

---

# Workshop Schedule: Windows & Applied Fuzzing
**Objective:** A hands-on, target-rich workshop focusing on Windows environments and complex file formats.

## Module 1: Fuzzing Essentials with winAFL
**Targets:** LibArchive, WinRAR
*   **Focus:** Foundational fuzzing techniques on Windows.
*   **Topics:**
    *   winAFL internals and DynamoRIO instrumentation
    *   Creating harness for Windows DLLs
    *   Corpus generation for archive formats (zip, rar, 7z)
    *   **Lab:** Fuzzing WinRAR content parsing logic.

## Module 2: Vulnerability Discovery & Coverage Analysis
**Target:** IrfanView
**Tools:** Jackalope, Lighthouse
*   **Focus:** Advanced analysis and visual coverage tracking.
*   **Topics:**
    *   Jackalope architecture (blackbox binary fuzzing)
    *   Using Lighthouse to visualize code coverage in IDA Pro / Binary Ninja
    *   Triage analysis of crashes (exploitable vs nuisance)
    *   **Lab:** Rediscovering RCE in IrfanView PSP parser.

## Module 3: Structural Fuzzing & Symbol-less Reversing
**Targets:** IrfanView PDF Plugin, PDF-XChange
*   **Focus:** Grammar-based fuzzing without source code.
*   **Topics:**
    *   Defining grammars for complex formats (PDF structure)
    *   Symbol-less binary analysis key strategies
    *   Reversing proprietary file parsers
    *   **Lab:** Writing a custom grammar for PDF objects.

## Module 4: Snapshot Fuzzing
**Target:** Assault Cube (FPS Game)
**Framework:** Wtf (What The Fuzz)
*   **Focus:** Fuzzing stateful, GUI-heavy applications using memory snapshots.
*   **Topics:**
    *   Concept of Snapshot Fuzzing (saving RAM state vs restarting process)
    *   The `wtf` fuzzing framework architecture
    *   Fuzzing network packets in multiplayer games
    *   **Lab:** Fuzzing map parsing logic in Assault Cube.

---

# Specialization Track: Rust Security & Fuzzing
**Objective:** Mastering memory safety auditing and fuzzing ecosystem for Rust applications.

## Module 1: Rust Audit & Code Review
*   **Core Concepts:** Ownership, Borrowing, and Lifetimes from a security perspective.
*   **Vulnerability Classes:**
    *   `unsafe` block auditing & invariants
    *   Panic-induced DoS (unwrapping, index out of bounds)
    *   FFI (Foreign Function Interface) boundary issues
    *   Logic bugs & Cryptographic misuse
*   **Tooling:** `cargo-audit`, `clippy` for security, and manual review patterns.

## Module 2: Rust Fuzzing & Crash Analysis
*   **Workflow:** Rust-native fuzzing corpus selection and setup.
*   **Toolchain:**
    *   `cargo-fuzz` (libFuzzer wrapper)
    *   `afl.rs` (AFL++ bindings)
    *   `honggfuzz-rs`
*   **Advanced Techniques:**
    *   Structure-aware fuzzing with `arbitrary` trait
    *   Differential fuzzing (Rust implementation vs C implementation)
    *   Writing custom Rust fuzzers from scratch
