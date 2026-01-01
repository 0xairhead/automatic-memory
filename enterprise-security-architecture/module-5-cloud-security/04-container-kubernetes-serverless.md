# Module 5, Lesson 4: Container, Kubernetes & Serverless Security

## Table of Contents
- [Media Resources](#media-resources)
- [Container Security Fundamentals](#container-security-fundamentals)
  - [Container Architecture & Attack Surface](#container-architecture--attack-surface)
  - [Image Security](#image-security)
  - [Registry Security](#registry-security)
  - [Runtime Security](#runtime-security)
- [Kubernetes Security](#kubernetes-security)
  - [Kubernetes Attack Surface](#kubernetes-attack-surface)
  - [Authentication & Authorization](#authentication--authorization)
  - [Pod Security Standards](#pod-security-standards)
  - [Network Policies](#network-policies)
  - [Secrets Management](#secrets-management)
  - [Cluster Hardening](#cluster-hardening)
- [Serverless Security](#serverless-security)
  - [The Serverless Security Model](#the-serverless-security-model)
  - [Function-Level Security](#function-level-security)
  - [Event Injection Attacks](#event-injection-attacks)
- [Key Concepts to Remember](#key-concepts-to-remember)
- [Practice Questions](#practice-questions)
- [Next Up](#next-up)

---

Containers and serverless have transformed how we deploy applications. They also transformed how attackers think about cloud environments.

---

## Media Resources

**Visual Guide:**

![Container and Kubernetes Security Architecture](./assets/04-container-kubernetes-security.png)

**Audio Lecture:**

🎧 [Securing Containers, Kubernetes & Serverless (Audio)](./assets/04-container-kubernetes-audio.m4a)

---

## Container Security Fundamentals

---

### Container Architecture & Attack Surface

Containers share the host kernel—this is both their efficiency advantage and their security challenge.

```
Virtual Machines:                    Containers:
┌─────────┐ ┌─────────┐              ┌─────────┐ ┌─────────┐
│  App A  │ │  App B  │              │  App A  │ │  App B  │
├─────────┤ ├─────────┤              ├─────────┤ ├─────────┤
│ Guest OS│ │ Guest OS│              │ Bins/   │ │ Bins/   │
│ (Full)  │ │ (Full)  │              │ Libs    │ │ Libs    │
├─────────┴─┴─────────┤              ├─────────┴─┴─────────┤
│     Hypervisor      │              │   Container Runtime │
├─────────────────────┤              ├─────────────────────┤
│      Host OS        │              │      Host OS        │
├─────────────────────┤              │   (SHARED KERNEL)   │
│     Hardware        │              ├─────────────────────┤
└─────────────────────┘              │     Hardware        │
                                     └─────────────────────┘
Strong isolation boundary            Weaker isolation boundary
```

**Container attack surface:**
```
┌─────────────────────────────────────────────────────────────────────┐
│ Attack Vectors                                                      │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  1. Malicious/Vulnerable Images                                     │
│     └── Backdoors, malware, vulnerable packages                     │
│                                                                     │
│  2. Registry Attacks                                                │
│     └── Image tampering, typosquatting, poisoned base images        │
│                                                                     │
│  3. Container Escape                                                │
│     └── Kernel exploits, misconfigurations, privileged containers   │
│                                                                     │
│  4. Runtime Attacks                                                 │
│     └── Cryptomining, reverse shells, data exfiltration             │
│                                                                     │
│  5. Orchestrator Attacks                                            │
│     └── Kubernetes API abuse, RBAC bypass, etcd exposure            │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

---

### Image Security

**The supply chain problem:**
```
Your Dockerfile:

FROM python:3.11-slim          ← Do you trust this image?
                                  (Who built it? What's inside?)
RUN pip install flask==2.3.0   ← Do you trust PyPI?
                                  (Dependency vulnerabilities?)
COPY ./app /app                ← Is YOUR code secure?
                                  (SAST/DAST checked?)
```

**Image scanning workflow:**
```
┌──────────┐    ┌──────────┐    ┌──────────┐    ┌──────────┐
│ Build    │───▶│ Scan     │───▶│ Sign     │───▶│ Deploy   │
│ Image    │    │ (Trivy,  │    │ (Cosign, │    │ (Verify  │
│          │    │ Snyk)    │    │ Notary)  │    │ Signature)│
└──────────┘    └──────────┘    └──────────┘    └──────────┘
                     │
                     ▼
            ┌──────────────┐
            │ Fail if:     │
            │ - Critical   │
            │   CVEs found │
            │ - Secrets    │
            │   detected   │
            │ - Policy     │
            │   violations │
            └──────────────┘
```

**Best practices for secure images:**

1. **Use minimal base images:**
   ```dockerfile
   # Bad - Full OS, large attack surface
   FROM ubuntu:22.04

   # Better - Minimal image
   FROM python:3.11-slim

   # Best - Distroless (no shell, no package manager)
   FROM gcr.io/distroless/python3
   ```

2. **Don't run as root:**
   ```dockerfile
   # Create non-root user
   RUN useradd -r -u 1001 appuser
   USER appuser
   ```

3. **Multi-stage builds:**
   ```dockerfile
   # Build stage
   FROM golang:1.21 AS builder
   COPY . .
   RUN go build -o /app

   # Runtime stage (minimal)
   FROM gcr.io/distroless/static
   COPY --from=builder /app /app
   ENTRYPOINT ["/app"]
   ```

4. **Pin versions explicitly:**
   ```dockerfile
   # Bad - mutable tag
   FROM python:latest

   # Good - immutable digest
   FROM python@sha256:abc123...
   ```

---

### Registry Security

Your container registry is a crown jewel—compromise it, and attackers can inject malware into every deployment.

**Registry security controls:**

```
┌─────────────────────────────────────────────────────────────────────┐
│ Secure Registry Architecture                                       │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  ┌─────────────────┐                                                │
│  │ CI/CD Pipeline  │                                                │
│  │                 │                                                │
│  │ 1. Build image  │                                                │
│  │ 2. Scan image   │                                                │
│  │ 3. Sign image   │────────────────┐                               │
│  │    (Cosign)     │                │                               │
│  └─────────────────┘                ▼                               │
│                           ┌─────────────────┐                       │
│                           │ Private Registry│                       │
│                           │ (ECR, ACR, GCR) │                       │
│                           │                 │                       │
│                           │ • Encryption    │                       │
│                           │ • IAM access    │                       │
│                           │ • Vuln scanning │                       │
│                           │ • Image signing │                       │
│                           └────────┬────────┘                       │
│                                    │                                │
│                                    ▼                                │
│  ┌─────────────────────────────────────────────────────────────┐   │
│  │ Kubernetes Cluster                                          │   │
│  │                                                             │   │
│  │ Admission Controller:                                       │   │
│  │ • Verify signature (Cosign/Kyverno)                        │   │
│  │ • Check vulnerability scan passed                          │   │
│  │ • Enforce registry allowlist                               │   │
│  └─────────────────────────────────────────────────────────────┘   │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

**Image signing with Cosign:**
```bash
# Sign image
cosign sign --key cosign.key myregistry.com/myapp:v1.2.3

# Verify signature before deployment
cosign verify --key cosign.pub myregistry.com/myapp:v1.2.3
```

---

### Runtime Security

Once containers are running, you need visibility into their behavior.

**Falco - Runtime threat detection:**
```yaml
# Falco rule: Detect shell spawned in container
- rule: Terminal shell in container
  desc: Detect shell being spawned in a container
  condition: >
    spawned_process and container and
    shell_procs and
    not known_shell_spawn_cmdlines
  output: >
    Shell spawned in container
    (user=%user.name container=%container.name
     shell=%proc.name parent=%proc.pname)
  priority: WARNING
```

**Runtime protection layers:**
```
┌─────────────────────────────────────────────────────────────────────┐
│ Container Runtime Security Stack                                    │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  Application Layer                                                  │
│  └── WAF, Input validation, Auth                                   │
│                                                                     │
│  Container Layer                                                    │
│  └── Seccomp profiles (restrict syscalls)                          │
│  └── AppArmor/SELinux (mandatory access control)                   │
│  └── Read-only filesystem                                          │
│  └── Dropped capabilities                                          │
│                                                                     │
│  Runtime Layer                                                      │
│  └── Falco (behavioral detection)                                  │
│  └── eBPF-based monitoring (Cilium, Tetragon)                      │
│                                                                     │
│  Host Layer                                                         │
│  └── Minimal host OS (Bottlerocket, Flatcar)                       │
│  └── Host hardening (CIS Benchmark)                                │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

---

## Kubernetes Security

---

### Kubernetes Attack Surface

Kubernetes adds significant complexity—and attack surface.

```
┌─────────────────────────────────────────────────────────────────────┐
│ Kubernetes Components & Attack Vectors                              │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  Control Plane:                                                     │
│  ┌─────────────────┐  ┌─────────────────┐  ┌─────────────────┐     │
│  │ API Server      │  │ etcd            │  │ Controller Mgr  │     │
│  │                 │  │                 │  │                 │     │
│  │ Attack: Unauth  │  │ Attack: Direct  │  │ Attack: SSRF    │     │
│  │ access, RBAC    │  │ access to       │  │ to internal     │     │
│  │ bypass          │  │ cluster secrets │  │ services        │     │
│  └─────────────────┘  └─────────────────┘  └─────────────────┘     │
│                                                                     │
│  Worker Nodes:                                                      │
│  ┌─────────────────┐  ┌─────────────────┐  ┌─────────────────┐     │
│  │ Kubelet         │  │ Container       │  │ Pods            │     │
│  │                 │  │ Runtime         │  │                 │     │
│  │ Attack: Exposed │  │ Attack: Escape  │  │ Attack: Lateral │     │
│  │ API without     │  │ via CVE or      │  │ movement, priv  │     │
│  │ auth            │  │ misconfig       │  │ escalation      │     │
│  └─────────────────┘  └─────────────────┘  └─────────────────┘     │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

---

### Authentication & Authorization

**Authentication (Who are you?):**

```
┌─────────────────────────────────────────────────────────────────────┐
│ Kubernetes Authentication Methods                                   │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  Humans (kubectl):                                                  │
│  ├── OIDC (recommended) → Okta, Azure AD, Google                   │
│  ├── Client certificates                                           │
│  └── Webhook token authentication                                  │
│                                                                     │
│  Service Accounts (pods):                                          │
│  ├── Auto-mounted tokens (be careful!)                             │
│  └── Projected tokens (bound, time-limited - recommended)          │
│                                                                     │
│  Cloud Provider:                                                    │
│  ├── AWS: IAM Roles for Service Accounts (IRSA)                    │
│  ├── GCP: Workload Identity                                        │
│  └── Azure: Azure AD Workload Identity                             │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

**Authorization (What can you do?) - RBAC:**

```yaml
# Role: defines what actions are allowed
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  namespace: production
  name: pod-reader
rules:
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["get", "list", "watch"]  # Read-only, no create/delete

---
# RoleBinding: assigns role to user/service account
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  namespace: production
  name: read-pods
subjects:
- kind: ServiceAccount
  name: monitoring-sa
  namespace: production
roleRef:
  kind: Role
  name: pod-reader
  apiGroup: rbac.authorization.k8s.io
```

**RBAC security principles:**
- **Least privilege:** Only grant needed permissions
- **Namespace isolation:** Use Roles (not ClusterRoles) when possible
- **Avoid wildcards:** `resources: ["*"]` is dangerous
- **Review regularly:** Audit RBAC with tools like `kubectl-who-can`

---

### Pod Security Standards

Pod Security Standards (PSS) replaced the deprecated PodSecurityPolicy:

```yaml
# Namespace labels to enforce security standards
apiVersion: v1
kind: Namespace
metadata:
  name: production
  labels:
    # Enforce: violations are rejected
    pod-security.kubernetes.io/enforce: restricted
    # Warn: violations show warning but are allowed
    pod-security.kubernetes.io/warn: restricted
    # Audit: violations are logged
    pod-security.kubernetes.io/audit: restricted
```

**Three security profiles:**

| Profile | Description | Use Case |
|---------|-------------|----------|
| **Privileged** | Unrestricted, full capabilities | System components, legacy apps |
| **Baseline** | Blocks known privilege escalations | General workloads |
| **Restricted** | Heavily hardened, best practices | Security-sensitive workloads |

**Restricted profile requirements:**
```yaml
apiVersion: v1
kind: Pod
spec:
  securityContext:
    runAsNonRoot: true              # Must run as non-root
    seccompProfile:
      type: RuntimeDefault          # Seccomp required
  containers:
  - name: app
    securityContext:
      allowPrivilegeEscalation: false  # No privilege escalation
      capabilities:
        drop: ["ALL"]               # Drop all capabilities
      readOnlyRootFilesystem: true  # Read-only root
```

---

### Network Policies

By default, all pods can communicate with all other pods. Network Policies add segmentation.

```yaml
# Allow web pods to receive traffic only from ingress
# Allow web pods to connect only to api pods
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: web-policy
  namespace: production
spec:
  podSelector:
    matchLabels:
      app: web
  policyTypes:
  - Ingress
  - Egress
  ingress:
  - from:
    - namespaceSelector:
        matchLabels:
          name: ingress-nginx
    ports:
    - protocol: TCP
      port: 8080
  egress:
  - to:
    - podSelector:
        matchLabels:
          app: api
    ports:
    - protocol: TCP
      port: 3000
  - to:  # Allow DNS
    - namespaceSelector: {}
      podSelector:
        matchLabels:
          k8s-app: kube-dns
    ports:
    - protocol: UDP
      port: 53
```

**Zero-trust network pattern:**
```yaml
# Default deny all ingress and egress
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: default-deny-all
spec:
  podSelector: {}  # Applies to all pods
  policyTypes:
  - Ingress
  - Egress
```

---

### Secrets Management

Kubernetes Secrets are base64 encoded, not encrypted. Anyone with API access can read them.

**The problem:**
```bash
# Secrets are just base64
kubectl get secret db-creds -o yaml
# apiVersion: v1
# data:
#   password: cGFzc3dvcmQxMjM=  ← base64 decode = "password123"
```

**Solutions:**

**1. Encrypt secrets at rest:**
```yaml
# EncryptionConfiguration
apiVersion: apiserver.config.k8s.io/v1
kind: EncryptionConfiguration
resources:
  - resources:
    - secrets
    providers:
    - aescbc:
        keys:
        - name: key1
          secret: <base64-encoded-key>
    - identity: {}
```

**2. External secrets management:**
```
┌─────────────────┐      ┌─────────────────┐      ┌─────────────────┐
│ Vault / AWS     │◀────▶│ External Secrets│◀────▶│ Kubernetes      │
│ Secrets Manager │      │ Operator        │      │ Secrets         │
│                 │      │                 │      │ (synced)        │
└─────────────────┘      └─────────────────┘      └─────────────────┘
```

**3. Sealed Secrets (GitOps-friendly):**
```bash
# Encrypt secret with cluster's public key
kubeseal --format yaml < secret.yaml > sealed-secret.yaml

# Only the cluster can decrypt
# Safe to commit to Git
```

---

### Cluster Hardening

**CIS Kubernetes Benchmark essentials:**

```
1. API Server:
   ✓ Enable audit logging
   ✓ Use RBAC (--authorization-mode=RBAC)
   ✓ Disable anonymous auth (--anonymous-auth=false)
   ✓ Enable admission controllers

2. etcd:
   ✓ Encrypt at rest
   ✓ TLS for client connections
   ✓ Restrict access to API server only

3. Kubelet:
   ✓ Disable anonymous auth (--anonymous-auth=false)
   ✓ Use webhook authorization
   ✓ Rotate certificates

4. Network:
   ✓ Use Network Policies
   ✓ Encrypt pod-to-pod traffic (service mesh)

5. General:
   ✓ Keep Kubernetes updated
   ✓ Use namespaces for isolation
   ✓ Enable Pod Security Standards
```

---

## Serverless Security

---

### The Serverless Security Model

Serverless shifts responsibility—you manage less, but you still have security obligations.

```
┌─────────────────────────────────────────────────────────────────────┐
│ Serverless Shared Responsibility                                    │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  Provider Manages:              You Manage:                         │
│  ├── Physical infrastructure    ├── Function code security          │
│  ├── Host operating system      ├── Dependencies/libraries          │
│  ├── Container runtime          ├── IAM permissions                 │
│  ├── Scaling                    ├── Input validation                │
│  ├── Patching runtime           ├── Secrets handling                │
│  └── Network infrastructure     ├── Event data validation           │
│                                 └── Application logic               │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

---

### Function-Level Security

**Least privilege for functions:**
```yaml
# Bad - too permissive
Resources:
  MyFunction:
    Type: AWS::Lambda::Function
    Properties:
      Role: arn:aws:iam::123456789:role/admin  # Full admin!

# Good - scoped permissions
Resources:
  MyFunctionRole:
    Type: AWS::IAM::Role
    Properties:
      Policies:
        - PolicyDocument:
            Statement:
              - Effect: Allow
                Action:
                  - dynamodb:GetItem
                  - dynamodb:PutItem
                Resource: arn:aws:dynamodb:*:*:table/MyTable
```

**Dependency vulnerabilities:**
```
Your function:
├── handler.py (your code)
└── requirements.txt
    ├── requests==2.28.0 (you chose this)
    ├── urllib3==1.26.0  (transitive dependency - vulnerable!)
    └── certifi==2022.9.24 (transitive)

Attack: Exploit vulnerable urllib3 through requests library
```

**Mitigation:**
- Scan dependencies in CI/CD (Snyk, Dependabot)
- Pin all versions, including transitive
- Regular dependency updates
- Minimal dependencies

---

### Event Injection Attacks

Serverless functions are triggered by events. Those events can be malicious.

**Event sources and risks:**
```
┌─────────────────────────────────────────────────────────────────────┐
│ Event Source                    │ Injection Risk                    │
├─────────────────────────────────┼───────────────────────────────────┤
│ API Gateway (HTTP)              │ SQL injection, XSS, command inj   │
│ S3 (object upload)              │ Malicious filenames, content      │
│ SNS/SQS (messages)              │ Crafted message payloads          │
│ DynamoDB Streams                │ Poisoned data from compromised DB │
│ CloudWatch Events               │ Less risky (AWS-generated)        │
└─────────────────────────────────┴───────────────────────────────────┘
```

**Example - S3 event injection:**
```python
# Vulnerable Lambda function
def handler(event, context):
    bucket = event['Records'][0]['s3']['bucket']['name']
    key = event['Records'][0]['s3']['object']['key']

    # DANGER: Key could be: "file; rm -rf /"
    os.system(f"process_file {key}")  # Command injection!

# Secure version
def handler(event, context):
    bucket = event['Records'][0]['s3']['bucket']['name']
    key = event['Records'][0]['s3']['object']['key']

    # Validate and sanitize
    if not re.match(r'^[a-zA-Z0-9._-]+$', key):
        raise ValueError("Invalid key format")

    # Use subprocess with array (no shell)
    subprocess.run(["process_file", key], check=True)
```

**Cold start security implications:**
```
Cold Start:
┌────────────────────────────────────────┐
│ 1. Provision container                 │ ← New isolation boundary
│ 2. Download function code              │ ← Integrity check?
│ 3. Initialize runtime                  │ ← Init code runs
│ 4. Execute handler                     │
└────────────────────────────────────────┘

Warm Invocation:
┌────────────────────────────────────────┐
│ 4. Execute handler                     │ ← Same container, state persists!
└────────────────────────────────────────┘

Security concern: Data from previous invocation may persist in memory
```

---

## Key Concepts to Remember

1. **Secure the image supply chain** - Scan, sign, verify at every stage
2. **Containers share the kernel** - Escape risks are real; minimize capabilities
3. **Kubernetes RBAC is essential** - Least privilege, namespace isolation
4. **Pod Security Standards replace PSP** - Enforce at namespace level
5. **Network Policies create zero-trust** - Default deny, explicit allow
6. **Secrets need external management** - Vault, Secrets Manager, Sealed Secrets
7. **Serverless isn't "no security"** - Code, dependencies, and IAM are yours
8. **Validate all event inputs** - Every event source is an attack vector

---

## Practice Questions

**Q1:** Your organization is running a Kubernetes cluster where developers have `kubectl` access to deploy pods. A security scan reveals that several pods are running as root with host network access. What controls would you implement to prevent this?

<details>
<summary>View Answer</summary>

**Implement multiple layers of control:**

1. **Pod Security Standards (Immediate):**
   ```yaml
   apiVersion: v1
   kind: Namespace
   metadata:
     name: production
     labels:
       pod-security.kubernetes.io/enforce: restricted
       pod-security.kubernetes.io/warn: restricted
   ```
   This blocks pods running as root or with host network.

2. **Admission Controller (Kyverno or OPA Gatekeeper):**
   ```yaml
   # Kyverno policy
   apiVersion: kyverno.io/v1
   kind: ClusterPolicy
   metadata:
     name: disallow-host-network
   spec:
     validationFailureAction: enforce
     rules:
     - name: deny-host-network
       match:
         resources:
           kinds:
           - Pod
       validate:
         message: "Host network is not allowed"
         pattern:
           spec:
             hostNetwork: false
   ```

3. **RBAC restrictions:**
   - Remove ability to create pods with elevated privileges
   - Developers use CI/CD pipeline (not direct kubectl) for production
   - Service accounts with minimal permissions

4. **Remediation for existing pods:**
   - Identify violating pods: `kubectl get pods -A -o json | jq '.items[] | select(.spec.hostNetwork==true)'`
   - Work with teams to update deployments
   - Set deadline for compliance

5. **Audit and alerting:**
   - Enable Kubernetes audit logging
   - Alert on policy violations
   - Regular compliance reports

</details>

**Q2:** An attacker compromises a pod in your Kubernetes cluster. Describe the attack path they might take to escalate privileges and move laterally, and how you would prevent each step.

<details>
<summary>View Answer</summary>

**Attack path and defenses:**

```
Step 1: Enumerate environment
─────────────────────────────
Attack: Read mounted service account token
        cat /var/run/secrets/kubernetes.io/serviceaccount/token

Defense:
- Disable auto-mounting: automountServiceAccountToken: false
- Use projected tokens with short TTL
- Minimal RBAC for service accounts


Step 2: Query Kubernetes API
────────────────────────────
Attack: Use token to list secrets, pods, etc.
        kubectl --token=<token> get secrets -A

Defense:
- Strict RBAC - service accounts can only access needed resources
- Network Policies blocking pod-to-API-server (if not needed)


Step 3: Access cloud metadata
─────────────────────────────
Attack: curl http://169.254.169.254/latest/meta-data/iam/...

Defense:
- Network Policy blocking metadata service
- IRSA/Workload Identity (no instance metadata access)
- IMDSv2 enforcement


Step 4: Lateral movement to other pods
──────────────────────────────────────
Attack: Scan internal network, access other services

Defense:
- Network Policies (default deny, explicit allow)
- Service mesh with mTLS (Istio, Linkerd)
- Pod-to-pod encryption


Step 5: Container escape attempt
────────────────────────────────
Attack: Exploit kernel vulnerability, abuse privileges

Defense:
- Pod Security Standards (restricted)
- Drop all capabilities
- Read-only root filesystem
- Seccomp profiles
- Minimal/hardened host OS (Bottlerocket)


Step 6: Persistence
───────────────────
Attack: Create malicious CronJob, modify existing deployments

Defense:
- RBAC prevents resource creation
- Admission webhooks block unauthorized changes
- Audit logging + alerting on resource changes
```

**Detection:**
- Falco rules for suspicious behavior
- Network flow analysis
- API audit log monitoring

</details>

**Q3:** You're architecting a serverless application on AWS Lambda that processes credit card payments. What security controls would you implement across the function lifecycle?

<details>
<summary>View Answer</summary>

**Security controls across the lifecycle:**

**1. Development:**
```
├── Secure coding standards (OWASP)
├── SAST scanning (Semgrep, SonarQube)
├── Dependency scanning (Snyk, Dependabot)
├── Secrets in code detection (TruffleHog, git-secrets)
└── Code review requirements
```

**2. Build/Deploy (CI/CD):**
```
├── Dependency pinning (lock files)
├── SCA scanning with vulnerability threshold
├── IAM policy review (cfn-nag, checkov)
├── Infrastructure as Code (SAM, CDK, Terraform)
├── Signed artifacts
└── Separate accounts (dev/staging/prod)
```

**3. Runtime Configuration:**
```yaml
# SAM template
Resources:
  PaymentFunction:
    Type: AWS::Serverless::Function
    Properties:
      Runtime: python3.11
      MemorySize: 256
      Timeout: 30  # Short timeout
      ReservedConcurrentExecutions: 100  # DoS protection

      # Minimal IAM
      Policies:
        - DynamoDBCrudPolicy:
            TableName: !Ref TransactionsTable
        - KMSDecryptPolicy:
            KeyId: !Ref PaymentKey

      # VPC for network isolation
      VpcConfig:
        SecurityGroupIds: [!Ref LambdaSG]
        SubnetIds: [!Ref PrivateSubnet]

      # Environment
      Environment:
        Variables:
          LOG_LEVEL: INFO
          # NO secrets here - use Secrets Manager
```

**4. Data Protection:**
```
├── Encryption at rest (DynamoDB, S3)
├── Encryption in transit (TLS 1.2+)
├── Tokenization for card data
├── Don't log sensitive data
├── KMS for encryption keys
└── Secrets Manager for credentials
```

**5. Input Validation:**
```python
def handler(event, context):
    # Validate input structure
    body = json.loads(event.get('body', '{}'))

    # Validate card number format
    card_number = body.get('card_number', '')
    if not re.match(r'^\d{13,19}$', card_number):
        return {'statusCode': 400, 'body': 'Invalid card format'}

    # Use payment processor's tokenization
    # Never store raw card numbers
```

**6. Monitoring & Compliance:**
```
├── CloudWatch Logs (no sensitive data)
├── X-Ray tracing
├── CloudTrail for API calls
├── GuardDuty for threat detection
├── PCI DSS audit logging requirements
├── Retention policies
└── Alerting on anomalies
```

</details>

**Q4:** Explain the security implications of container image layers and how an attacker could exploit them. How would you mitigate these risks?

<details>
<summary>View Answer</summary>

**Container image layer security:**

**How layers work:**
```
Dockerfile:                    Resulting layers:
FROM ubuntu:22.04       →      Layer 1: Base OS (ubuntu)
RUN apt-get update      →      Layer 2: Package lists
RUN apt-get install -y  →      Layer 3: Installed packages
    nginx
COPY nginx.conf /etc/   →      Layer 4: Config file
COPY --chown=root       →      Layer 5: App with secrets!
    secrets.txt /app/
RUN rm /app/secrets.txt →      Layer 6: Deletion (file still in Layer 5!)
```

**Attack vector - Layer archaeology:**
```bash
# Pull image
docker pull myapp:latest

# Save and extract layers
docker save myapp:latest -o myapp.tar
tar -xf myapp.tar

# Each layer is a tar archive
# Examine each layer for secrets
for layer in */layer.tar; do
  tar -tf $layer | grep -E '(secret|password|key|token)'
done

# Even "deleted" files exist in earlier layers!
```

**Real-world examples:**
- AWS credentials in early layer, removed in later layer
- Private SSH keys copied then deleted
- Database passwords in environment setup

**Mitigations:**

1. **Multi-stage builds (primary defense):**
   ```dockerfile
   # Build stage - can have secrets for build
   FROM golang:1.21 AS builder
   COPY . .
   # Secrets used here don't end up in final image
   RUN go build -o /app

   # Runtime stage - clean image
   FROM gcr.io/distroless/static
   COPY --from=builder /app /app
   # Only this layer ships
   ```

2. **BuildKit secrets (build-time only):**
   ```dockerfile
   # syntax=docker/dockerfile:1.4
   FROM alpine
   RUN --mount=type=secret,id=api_key \
       API_KEY=$(cat /run/secrets/api_key) && \
       ./configure --api-key=$API_KEY
   # Secret never written to layer
   ```

3. **Image scanning for secrets:**
   ```bash
   # Trivy scans all layers
   trivy image --scanners secret myapp:latest

   # Trufflehog for git and images
   trufflehog docker --image myapp:latest
   ```

4. **Squash layers (use carefully):**
   ```bash
   docker build --squash -t myapp .
   # Combines all layers - deleted files truly gone
   # But loses layer caching benefits
   ```

5. **Use .dockerignore:**
   ```
   # .dockerignore
   .env
   *.pem
   *.key
   credentials/
   .git/
   ```

6. **External secrets at runtime:**
   - Mount secrets from Kubernetes Secrets / Vault
   - Use cloud provider secrets (Parameter Store, Secrets Manager)
   - Never bake secrets into images

</details>

---

## Next Up

In Lesson 5, we'll explore **Cloud Data Protection** — encryption strategies, key management, data loss prevention, and keeping your data safe across cloud services!
