# SELinux Syllabus

A streamlined, practitioner-focused path to mastering SELinux.

## Phase 0: Environment Setup (CRITICAL for macOS Users)
**[00-environment-setup.md](00-environment-setup.md)**
*   **Required for Mac Users**: You cannot run these labs natively on macOS.
*   **Recommendation**: Use **Lima** to create a lightweight, text-based Fedora environment.
*   Instructions for installing Lima and mounting your workflow are in the guide.

## Phase 1: Prerequisites
**[01-prerequisites.md](01-prerequisites.md)**
*   Linux internals refresher: Permissions, PIDs, Systemd, filesystems.
*   Goal: Understand the layers *below* SELinux.

## Phase 2: Fundamentals
**[02-selinux-fundamentals.md](02-selinux-fundamentals.md)**
*   modes (`Enforcing`, `Permissive`), Policy Types (`Targeted`).
*   Core Commands: `sestatus`, `getenforce`, `setenforce`.

## Phase 3: Labels, Contexts & Types
**[03-labels-contexts-types.md](03-labels-contexts-types.md)**
*   The Security Context: `user:role:type:level`.
*   Type Enforcement (TE) basics.
*   The `mv` vs `cp` labeling problem.

## Phase 4: Booleans
**[04-selinux-booleans.md](04-selinux-booleans.md)**
*   Runtime policy toggles (Safe Flexibility).
*   Commands: `getsebool`, `setsebool -P`.

## Phase 5: Troubleshooting & Auditing
**[05-troubleshooting-auditing.md](05-troubleshooting-auditing.md)**
*   Detecting denials: `ausearch`, `sealert`.
*   The persistence workflow: Break -> Detect -> Analyze -> Fix.

## Phase 6: Managing Contexts Properly
**[06-managing-contexts.md](06-managing-contexts.md)**
*   Permanent fixes vs temporary hacks.
*   `semanage fcontext` (Persistent) vs `chcon` (Temporary).
*   `restorecon` usage.

## Phase 7: Custom Policy
**[07-custom-policy.md](07-custom-policy.md)**
*   Creating custom modules with `audit2allow`.
*   Compiling and installing `.pp` files.
*   Dangers of blind policy generation.

## Phase 8: Containers & Cloud
**[08-containers-cloud.md](08-containers-cloud.md)**
*   SELinux in Docker/Podman (`container_t`).
*   Volume labelling with `:Z`.
*   `udica` for custom container policies.

## Phase 9: Advanced & Offensive
**[09-advanced-offensive.md](09-advanced-offensive.md)**
*   Domain transitions.
*   Exploit mitigation mechanics.
*   Offensive perspective (Red Teaming against SELinux).

## Practical Labs
**[10-practical-labs.md](10-practical-labs.md)**
*   10 Scenario-based exercises ranging from Basic to Expert.
*   Includes SSH ports, Custom Webroots, Container Volumes, and Manual Policy writing.
