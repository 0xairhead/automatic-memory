# Phase 7: SELinux for Containers & Cloud (Week 7)

## Concepts
SELinux is a core defense layer in modern container platforms (Podman, Docker, Kubernetes/OpenShift). It prevents "container escapes"—where a process inside a container attacks the host.

### sVirt (Secure Virtualization)
**sVirt** is the technology that uses SELinux to protect VMs and Containers. It dynamically assigns a unique label to each individual container process.

*   **The Problem**: If two containers both run as `container_t`, what stops them from attacking each other?
*   **The Solution**: sVirt adds a unique **MCS (Multi-Category Security)** label (e.g., `c1,c2`) to the process.
*   **The Result**: Container A (`container_t:s0:c1,c2`) cannot access the files of Container B (`container_t:s0:c3,c4`), even though they share the same base type.

### Container Labels
By default, containers run with the label `container_t`.
Files on the host intended for containers must be labeled `container_file_t`.

### The `:Z` Flag
When maximizing Docker/Podman volumes, you often hit "Permission Denied".
*   **Problem**: You mount `/host/data` into a container. `/host/data` is `admin_home_t` (or similar). Container (`container_t`) cannot read it.
*   **Solution**: The `:z` or `:Z` suffixes.
    *   `:z`: Relabels the bind mount content to be shared among multiple containers.
    *   `:Z`: Relabels the bind mount content to be **private and unshared** (only this container).

## Hands-On Exercises

### 1. The Volume Mount Issue
**Step 1: Create a directory**
```bash
mkdir data
echo "Hello" > data/file.txt
```

**Step 2: Run a container without Z (Fail)**
```bash
# This often fails depending on distro defaults
podman run --rm -v $(pwd)/data:/data alpine cat /data/file.txt
```
*If it fails, check audit logs!*

**Step 3: Run with Z (Success)**
Podman/Docker automatically runs `chcon` (or equivalent) for you.
```bash
podman run --rm -v $(pwd)/data:/data:Z alpine cat /data/file.txt
```

**Step 4: Check the label after**
```bash
ls -Z data/file.txt
# Notice the type changed to container_file_t
```

### 2. UDICA (Advanced)
**UDICA** helps you write custom SELinux policies for containers that need *more* access than a standard web server, but *less* than full `--privileged`.

**The Scenario**: You have a monitoring container that needs to read access to `/var/log` on the host. Standard `container_t` blocks this.
**The Fix**: Instead of disabling SELinux for the container (bad!), `udica` inspects your container definition and generates a custom policy that allows *exactly* what you need.

#### Workflow
1.  **Start the container** (it will fail or run with errors due to denials).
2.  **Export the JSON definition**:
    ```bash
    podman inspect my_container > container.json
    ```
3.  **Generate the Policy**:
    `udica` reads the JSON (ports, mounts, capabilities) and writes a `.cil` (Common Intermediate Language) policy.
    ```bash
    udica my_custom_policy container.json > my_custom_policy.cil
    ```
4.  **Load the Policy**:
    ```bash
    sudo semodule -i my_custom_policy.cil
    ```
5.  **Restart the Container with the New Label**:
    Now you tell Podman/Docker to use your new custom type instead of the generic `container_t`.
    ```bash
    podman run --security-opt label=type:my_custom_policy.process ...
    ```

## Cloud Context: OpenShift & Kubernetes
*   **OpenShift**: heavily relies on SELinux. Every namespace gets a unique MCS category (e.g., `s0:c1,c2`), ensuring Container A cannot accesses Container B's files even if they run as the same UID.
*   **Security Context Constraints (SCC)**: The K8s/OpenShift object that controls what SELinux labels a pod can run with.

## Key Takeaway
> In a container world, **SELinux is your safety net against breakouts.**
> Use the `:Z` flag for volume mounts to let the runtime handle labeling for you.
