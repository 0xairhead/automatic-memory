# 🧠 eBPF Syllabus

## Table of Contents
* [Phase 0: Prerequisites](./00-prerequisites.md)
* [Phase 1: Foundations](./01-foundations.md)
* [Phase 2: First Programs](./02-first-programs.md)
* [Phase 3: Maps](./03-maps.md)
* [Phase 4: Observability](./04-observability.md)
* [Phase 5: Security](./05-security.md)
* [Phase 6: DevOps & Cloud](./06-devops-cloud.md)
* [Phase 7: Advanced](./07-advanced.md)
* [Phase 8: Projects](./08-projects.md)

---

## [Phase 0: Prerequisites](./00-prerequisites.md)
* Linux process model (PID, threads), Syscalls, Virtual memory, `/proc`, `/sys`.
* **C Programming:** Pointers, Structs, Inline functions, Endianness.

## [Phase 1: Foundations](./01-foundations.md)
* **Kernel Internals:** User/Kernel space, Hooks (kprobes, tracepoints), BTF.
* **eBPF Concepts:** JIT compilation, Verifier safety model, Program lifecycle.

## [Phase 2: First Programs](./02-first-programs.md)
* **Tooling:** `clang`, `llvm`, `bpftool`, `libbpf`.
* **Program Types:** kprobe, tracepoint, socket filter.
* **Practice:** Tracing `execve`, logging file opens.

## [Phase 3: Maps](./03-maps.md)
* **Map Types:** Hash, Array, LRU, Per-CPU, Ring Buffer.
* **Concepts:** Data sharing, Key/value design, Memory limits.

## [Phase 4: Observability](./04-observability.md)
* **Metrics:** Histograms, Latency, CPU usage.
* **Tracing:** Function entry/exit, Stack traces, Kernel scheduling.
* **Networking:** TCP latency, Packet drops, HTTP tracing.
* **Tools:** `bpftrace`, Pixie, Cilium Hubble.

## [Phase 5: Security](./05-security.md)
* **Runtime Security:** Execution monitoring, File access auditing.
* **Kubernetes:** Container-aware tracing, cgroups.
* **Frameworks:** Falco, Tracee, Tetragon.

## [Phase 6: DevOps & Cloud](./06-devops-cloud.md)
* **Networking:** XDP vs TC, Traffic shaping, Load balancing.
* **Kubernetes:** CNI (Cilium), Network policies, Sidecarless Service Mesh.
* **Ops:** Health checks, Production debugging.

## [Phase 7: Advanced](./07-advanced.md)
* **Verifier:** Instruction limits, Stack rules, Loop restrictions.
* **Performance:** Per-CPU data, Map contention, Tail calls.
* **Advanced:** CO-RE, BTF, eBPF LSM.

## [Phase 8: Projects](./08-projects.md)
1. eBPF Runtime Threat Detector
2. Kubernetes syscall profiler
3. HTTP latency tracer
4. Cloud workload anomaly detector
5. Custom Falco-like rule engine

---

## Recommended Resources
* **Books:** *Learning eBPF* (Liz Rice), *BPF Performance Tools* (Brendan Gregg).
* **Repos:** `libbpf/libbpf-bootstrap`, `cilium/cilium`, `aquasecurity/tracee`.
