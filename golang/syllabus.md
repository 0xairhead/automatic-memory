# Go Syllabus

## 📊 Course Progress at a Glance

| Phase | Description | Key Artifacts |
| :--- | :--- | :--- |
| **0** | **Setup & Mindset** | `00-setup/` |
| **1** | **Fundamentals** | `01-fundamentals/syntax/` |
| **2** | **Data Structures** | `02-datastructures/basics/` |
| **3** | **Pointers & Errors** | `03-pointers-errors/basics/` |
| **4** | **Interfaces** | `04-interfaces/basics/` |
| **5** | **Concurrency** | `05-concurrency/basics/` |
| **6** | **OS & Networking** | `06-files-os-net/basics/` |
| **7** | **Testing & Quality** | `07-testing/basics/` |
| **8** | **Modules & Deps** | `08-modules/` |
| **9** | **Advanced Go** | `09-advanced/` |
| **10** | **Security Go** | `10-security/` |
| **11** | **Cloud-Native Go** | `11-cloud-native/` |
| **12** | **Capstones** | `12-capstones/` |

---

## 🏗️ Detailed Module Breakdown

### [0–4] Foundation & Idiomatic Go (Core Blocks)
*   **Outcome:** Ability to write clean, idiomatic Go as a single-threaded application.
*   **Key Skills:** Type safety, multiple returns, interface-based design, error wrapping.

### [5-6] Systems Programming (The Go Superpowers)
*   **Outcome:** Mastering concurrency and OS interactions.
*   **Key Projects:**
    *   `worker_pool`: Distributed job processing.
    *   `api_server`: High-performance HTTP routing and JSON handling.

### [7-9] Production Grade (Scale & Performance)
*   **Outcome:** Moving from "scripting" to "engineering".
*   **Key Skills:** Table-driven testing, Benchmarking (fighting allocations), Generics, and Reflection for dynamic code.

### [10] Security Engineering
*   **Outcome:** Building tools that resist attack and can analyze other code.
*   **Key Projects:**
    *   `jwt_auth`: Industry-standard authentication.
    *   `static_scanner`: AST parsing to find "unsafe" code.

---

## 🚀 The Path Forward: Phase 11 & 12

### Phase 11: Cloud-Native Architecture
**Objective:** Scale Go to the cloud using industry-standard protocols and observability.
*   **CLI Engineering**: Building "CLI first" tools using [Cobra](https://github.com/spf13/cobra).
*   **Communication**: Transitioning from REST to **gRPC** and **Protobuf**.
*   **Observability**: Integrated Structured Logging (slog), Prometheus metrics, and OpenTelemetry.
*   **K8s Deep Dive**: Using `client-go` to interact with Kubernetes clusters.

### Phase 12: Professional Capstone Projects
**Objective:** Apply 100% of the knowledge to a portfolio-grade project.
*   **Option A: Cloud Security Scanner** (IAM & S3 Auditor).
*   **Option B: Endpoint Agent (XDR)** (Process monitoring & File integrity).
*   **Option C: High-Performance WAF** (HTTP Middleware for exploit detection).
