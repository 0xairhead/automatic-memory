# Module 5, Lesson 2: Cloud Network Security

## Table of Contents
- [Media Resources](#media-resources)
- [Virtual Private Cloud (VPC) Fundamentals](#virtual-private-cloud-vpc-fundamentals)
  - [VPC Design Patterns](#vpc-design-patterns)
  - [Subnets: Public vs Private](#subnets-public-vs-private)
- [Network Security Controls](#network-security-controls)
  - [Security Groups (Stateful Firewalls)](#security-groups-stateful-firewalls)
  - [Network ACLs (Stateless Firewalls)](#network-acls-stateless-firewalls)
  - [Security Groups vs NACLs](#security-groups-vs-nacls)
- [Advanced Network Architecture](#advanced-network-architecture)
  - [Private Endpoints & Service Endpoints](#private-endpoints--service-endpoints)
  - [Transit Gateway & Hub-Spoke](#transit-gateway--hub-spoke)
  - [Cloud-Native DDoS Protection](#cloud-native-ddos-protection)
  - [Deep Dive: Cloud Networking Internals](#deep-dive-cloud-networking-internals)
- [Network Visibility & Forensics](#network-visibility--forensics)
- [Zero Trust Network Access (ZTNA)](#zero-trust-network-access-ztna)
- [DNS Security](#dns-security)
- [Defense in Depth: Layered Network Security](#defense-in-depth-layered-network-security)
- [Key Concepts to Remember](#key-concepts-to-remember)
- [Practice Questions](#practice-questions)
- [Next Up](#next-up)

---

Network security is where traditional security skills meet cloud-native concepts. This lesson bridges that gap.

---

## Media Resources

**Visual Guide:**

![Cloud Network Architecture](./assets/02-cloud-network-architecture.png)

**Audio Lecture:**

🎧 [VPC, Security Groups & Transit Gateways (Audio)](./assets/02-cloud-network-audio.m4a)

---

## Virtual Private Cloud (VPC) Fundamentals

A VPC is your own isolated network within the cloud provider's infrastructure. Think of it as your private data center, but software-defined.

```
┌─────────────────────────────────────────────────────────────────────┐
│ AWS Cloud / Azure / GCP                                             │
│  ┌───────────────────────────────────────────────────────────────┐  │
│  │ Your VPC (10.0.0.0/16)                                        │  │
│  │                                                               │  │
│  │  ┌─────────────────┐      ┌─────────────────┐                 │  │
│  │  │ Public Subnet   │      │ Private Subnet  │                 │  │
│  │  │ 10.0.1.0/24     │      │ 10.0.2.0/24     │                 │  │
│  │  │                 │      │                 │                 │  │
│  │  │  ┌───────────┐  │      │  ┌───────────┐  │                 │  │
│  │  │  │ Web Server│  │      │  │ Database  │  │                 │  │
│  │  │  └───────────┘  │      │  └───────────┘  │                 │  │
│  │  └─────────────────┘      └─────────────────┘                 │  │
│  │          │                        ▲                           │  │
│  │          │                        │ (internal only)           │  │
│  │          ▼                        │                           │  │
│  │    ┌─────────────┐                │                           │  │
│  │    │ Internet GW │────────────────┘                           │  │
│  │    └─────────────┘                                            │  │
│  └───────────────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────────────┘
```

---

### VPC Design Patterns

**Single VPC (Simple):**
```
Best for: Startups, small applications, dev/test
┌──────────────────────────────────┐
│ VPC                              │
│  ├── Public Subnet (web tier)    │
│  ├── Private Subnet (app tier)   │
│  └── Private Subnet (data tier)  │
└──────────────────────────────────┘
```

**Multi-VPC (Account Segmentation):**
```
Best for: Enterprise, compliance requirements, team isolation

┌─────────────┐  ┌─────────────┐  ┌─────────────┐
│ Prod VPC    │  │ Dev VPC     │  │ Shared Svcs │
│ Account     │  │ Account     │  │ VPC         │
└──────┬──────┘  └──────┬──────┘  └──────┬──────┘
       │                │                │
       └────────────────┼────────────────┘
                        │
                ┌───────┴───────┐
                │ Transit GW /  │
                │ VPC Peering   │
                └───────────────┘
```

**Hub-and-Spoke (Enterprise Standard):**
```
Best for: Large enterprises, centralized security, shared services

                    ┌───────────────┐
                    │   Hub VPC     │
                    │ (Shared Svcs) │
                    │  - DNS        │
                    │  - Security   │
                    │  - Logging    │
                    └───────┬───────┘
           ┌────────────────┼────────────────┐
           │                │                │
    ┌──────┴──────┐  ┌──────┴──────┐  ┌──────┴──────┐
    │ Spoke: Prod │  │ Spoke: Dev  │  │ Spoke: Test │
    └─────────────┘  └─────────────┘  └─────────────┘
```

---

### Subnets: Public vs Private

**Public Subnet:**
- Has a route to an Internet Gateway
- Resources CAN have public IP addresses
- Used for: Load balancers, bastion hosts, NAT gateways

**Private Subnet:**
- NO direct route to the internet
- Resources cannot be directly accessed from internet
- Used for: Application servers, databases, internal services
- Can reach internet via NAT Gateway (outbound only)

**Architecture principle:** Minimize your public subnet footprint. Only what MUST be public should be public.

```
Internet Traffic Flow:

Inbound (to your app):
Internet → IGW → ALB (public) → App Server (private)

Outbound (from your app):
App Server (private) → NAT GW (public) → IGW → Internet
```

---

## Network Security Controls

---

### Security Groups (Stateful Firewalls)

Security groups are virtual firewalls attached to resources (EC2, RDS, etc.).

**Key characteristics:**
- **Stateful:** If you allow inbound traffic, the response is automatically allowed
- **Allow rules only:** No explicit deny (implicit deny all)
- **Attached to resources:** Not subnets

**Example - Web Server Security Group:**
```
Inbound Rules:
┌────────────┬──────────┬─────────────────────────────┐
│ Protocol   │ Port     │ Source                      │
├────────────┼──────────┼─────────────────────────────┤
│ TCP        │ 443      │ 0.0.0.0/0 (anywhere)        │
│ TCP        │ 80       │ 0.0.0.0/0 (redirect to 443) │
│ TCP        │ 22       │ sg-bastion (bastion SG)     │
└────────────┴──────────┴─────────────────────────────┘

Outbound Rules:
┌────────────┬──────────┬─────────────────────────────┐
│ Protocol   │ Port     │ Destination                 │
├────────────┼──────────┼─────────────────────────────┤
│ TCP        │ 443      │ 0.0.0.0/0 (API calls)       │
│ TCP        │ 5432     │ sg-database (DB SG)         │
└────────────┴──────────┴─────────────────────────────┘
```

**Security Group chaining (best practice):**
```
┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│ sg-alb      │───▶│ sg-app      │───▶│ sg-database │
│             │    │             │    │             │
│ Allow: 443  │    │ Allow: 8080 │    │ Allow: 5432 │
│ from: ANY   │    │ from: sg-alb│    │ from: sg-app│
└─────────────┘    └─────────────┘    └─────────────┘

Database only accepts connections from app servers.
App servers only accept connections from ALB.
This is better than using IP ranges!
```

---

### Network ACLs (Stateless Firewalls)

NACLs operate at the subnet level and are stateless.

**Key characteristics:**
- **Stateless:** Must explicitly allow both inbound AND outbound
- **Allow AND Deny rules:** Can explicitly block IPs
- **Rule numbers:** Evaluated in order (lowest first)
- **Applied to subnets:** All traffic in/out of subnet

**Example - Private Subnet NACL:**
```
Inbound Rules:
┌────────┬────────────┬──────────┬─────────────────┬────────┐
│ Rule # │ Protocol   │ Port     │ Source          │ Action │
├────────┼────────────┼──────────┼─────────────────┼────────┤
│ 100    │ TCP        │ 8080     │ 10.0.1.0/24     │ ALLOW  │
│ 110    │ TCP        │ 1024-65535│ 0.0.0.0/0      │ ALLOW  │
│ *      │ ALL        │ ALL      │ 0.0.0.0/0       │ DENY   │
└────────┴────────────┴──────────┴─────────────────┴────────┘

Outbound Rules:
┌────────┬────────────┬──────────┬─────────────────┬────────┐
│ Rule # │ Protocol   │ Port     │ Destination     │ Action │
├────────┼────────────┼──────────┼─────────────────┼────────┤
│ 100    │ TCP        │ 443      │ 0.0.0.0/0       │ ALLOW  │
│ 110    │ TCP        │ 1024-65535│ 10.0.1.0/24    │ ALLOW  │
│ *      │ ALL        │ ALL      │ 0.0.0.0/0       │ DENY   │
└────────┴────────────┴──────────┴─────────────────┴────────┘
```

**Why ephemeral ports (1024-65535)?** When a client connects to port 443, the response goes to a random high port on the client. Stateless firewalls must allow these.

---

### Security Groups vs NACLs

| Feature | Security Groups | Network ACLs |
|---------|----------------|--------------|
| Level | Resource (ENI) | Subnet |
| State | Stateful | Stateless |
| Rules | Allow only | Allow and Deny |
| Evaluation | All rules | In order by number |
| Default | Deny all in, allow all out | Allow all |
| Use case | Primary control | Subnet-level blocklist |

**Best practice:** Use Security Groups as your primary control. Use NACLs for:
- Blocking known malicious IPs at the subnet edge
- Compliance requirements for explicit subnet controls
- Defense in depth (belt AND suspenders)

---

## Advanced Network Architecture

As you scale beyond a single VPC, you need robust patterns for connectivity and traffic management. This section covers enterprise-grade architecture components.

---

### Private Endpoints & Service Endpoints

**The problem:** Your application in a private subnet needs to call S3, DynamoDB, or other cloud services. Traditional approach sends traffic through the internet.

```
OLD (Internet path):
App (private) → NAT GW → IGW → Internet → S3

ISSUES:
- Data traverses public internet
- NAT Gateway costs ($$$ per GB)
- Higher latency
- Larger attack surface
```

**Private Endpoints (AWS VPC Endpoints, Azure Private Link, GCP Private Service Connect):**

```
NEW (Private path):
App (private) → VPC Endpoint → S3

BENEFITS:
- Traffic stays within cloud backbone
- No NAT costs
- Lower latency
- No internet exposure
```

**Types of endpoints:**

**Gateway Endpoints (AWS - S3 and DynamoDB only):**
- Free
- Route table entry points to the service
- Same region only

**Interface Endpoints (AWS PrivateLink):**
- Creates an ENI in your VPC with private IP
- Works for 100+ AWS services
- Works across regions and accounts
- Costs per hour + per GB

**Gateway Load Balancer (GWLB) & GENEVE:**
*   **Purpose:** To transparently insert third-party firewalls (like Palo Alto, Fortinet) into the traffic flow.
*   **Protocol:** Uses **GENEVE** encapsulation to send traffic to the appliance and receive it back unchanged (bump-in-the-wire).
*   **Why:** Traditional load balancers change the Source IP (SNAT). GWLB preserves the original packet header so your firewall sees the real source IP.

**Example architecture:**
```
┌─────────────────────────────────────────────────────────────┐
│ VPC                                                         │
│  ┌─────────────────┐     ┌─────────────────────────────┐    │
│  │ Private Subnet  │     │ Interface Endpoints         │    │
│  │                 │     │  ┌────────────────────────┐ │    │
│  │  ┌───────────┐  │     │  │ vpce-ssm  (10.0.3.5)   │ │    │
│  │  │ EC2       │──┼────▶│  │ vpce-s3   (10.0.3.6)   │ │    │
│  │  └───────────┘  │     │  │ vpce-sqs  (10.0.3.7)   │ │    │
│  │                 │     │  └────────────────────────┘ │    │
│  └─────────────────┘     └─────────────────────────────┘    │
│                                      │                      │
└──────────────────────────────────────┼──────────────────────┘
                                       │ AWS Backbone (private)
                                       ▼
                              ┌─────────────────┐
                              │ AWS Services    │
                              │ (S3, SSM, SQS)  │
                              └─────────────────┘
```

---

### Transit Gateway & Hub-Spoke

When you have many VPCs, VPC peering becomes a mesh nightmare:

```
VPC Peering (N VPCs = N*(N-1)/2 connections):

    VPC-A ─────── VPC-B
      │  \       /  │
      │   \     /   │
      │    \   /    │
      │     \ /     │
      │      X      │
      │     / \     │
      │    /   \    │
      │   /     \   │
      │  /       \  │
    VPC-C ─────── VPC-D

4 VPCs = 6 peering connections
10 VPCs = 45 peering connections
```

**Transit Gateway:**
```
                    ┌─────────────────┐
                    │ Transit Gateway │
                    └────────┬────────┘
          ┌──────────────────┼──────────────────┐
          │                  │                  │
      ┌───┴───┐          ┌───┴───┐          ┌───┴───┐
      │ VPC-A │          │ VPC-B │          │ VPC-C │
      └───────┘          └───────┘          └───────┘
                              │
                    ┌─────────┴─────────┐
                    │                   │
              ┌─────┴─────┐       ┌─────┴─────┐
              │ VPN to    │       │ Direct    │
              │ On-Prem   │       │ Connect   │
              └───────────┘       └───────────┘

10 VPCs = 10 attachments (not 45 peerings)
```

**Security features of Transit Gateway:**
- Route table segmentation (prod can't talk to dev)
- Centralized firewall inspection
- Centralized egress with NAT
- Logging of inter-VPC traffic

---

### Cloud-Native DDoS Protection

Every cloud provider includes baseline DDoS protection:

**AWS Shield:**
- **Standard (free):** Automatic protection against L3/L4 attacks
- **Advanced ($3K/month):** L7 protection, 24/7 DDoS Response Team, cost protection

**Azure DDoS Protection:**
- **Basic (free):** Always-on traffic monitoring
- **Standard:** Adaptive tuning, attack analytics, rapid response

**GCP Cloud Armor:**
- WAF + DDoS protection
- Pre-configured WAF rules (OWASP, SQLi, XSS)
- Adaptive protection using ML

**Architecture for DDoS resilience:**
```
                            ┌─────────────────┐
Internet ──▶ CloudFront ───▶│ AWS Shield      │
             (CDN)          │ (DDoS filtering)│
                            └────────┬────────┘
                                     │
                            ┌────────┴────────┐
                            │ AWS WAF         │
                            │ (L7 filtering)  │
                            └────────┬────────┘
                                     │
                            ┌────────┴────────┐
                            │ ALB             │
                            │ (Load balancer) │
                            └────────┬────────┘
                                     │
                              Your Application
```

---

### Deep Dive: Cloud Networking Internals

Abstraction is great, but knowing what happens *under the hood* distinguishes a senior engineer from a junior one.

#### 1. AWS Hyperplane (The Magic Behind NAT/NLB/TGW)
Ever wonder how a NAT Gateway handles 100Gbps+ without crashing? It's not a single EC2 instance. It uses **Hyperplane**, AWS's internal state management system.
*   **Packet Sharding:** Hyperplane shards connections across thousands of underlying hosts.
*   **Result:** A "Network Load Balancer" or "Transit Gateway" isn't a device; it's a massive distributed fleet. This is why you get valid static IPs for NLBs but not for ALBs (which run on standard EC2 fleets).

#### 2. Azure VNet Injection (PaaS Inside Your Network)
In Azure, PaaS services (like Databricks or App Service) normally live in public Azure space. To secure them, you "inject" them into your VNet.
*   **Mechanism:** Azure mounts distinct NICs representing the PaaS service directly into your subnet.
*   **Constraint:** This "delegates" the subnet to that service, meaning you often can't mix other resources (like standard VMs) in that same subnet. This is why you need so many "dedicated subnets" in Azure.

#### 3. GCP Andromeda and Jupiter (Software Defined Everything)
GCP's network is essentially one giant global switch.
*   **Andromeda:** The virtualization stack. It allows GCP to offer "Global VPCs" (a single VPC spanning Asia, US, and Europe) out of the box—something harder to achieve in AWS/Azure without peering/TGWs.
*   **Jupiter:** The data center fabric. It separates control plane from data plane so cleanly that you can live-migrate VMs across hosts while they are maintaining active TCP connections.

---

## Network Visibility & Forensics

You can't secure what you can't see.

### VPC Flow Logs
Flow Logs capture **metadata** about the IP traffic going to and from your network interfaces.
*   **Key limitation:** They do **NOT** Capture packet contents (payload). You can see *that* IP X talked to IP Y on Port 80, but you cannot see *what* they said (e.g., you can't see the SQL query or the exfiltrated file).
*   **Use case:** Troubleshooting connectivity, detecting scanning attempts, audit trails.

**Sample Log:**
`2 123456789 eni-abc 10.0.1.2 198.51.100.1 443 49152 6 25 20000 1627233333 1627233393 ACCEPT OK`
(Account, Interface, SrcIP, DstIP, SrcPort, DstPort, Protocol, Packets, Bytes, Start, End, Action, Status)

### Traffic Mirroring
For Deep Packet Inspection (DPI):
*   **AWS:** Traffic Mirroring
*   **Azure:** vTAP (Virtual Network TAP)
*   **GCP:** Packet Mirroring
These services verify the actual payload (e.g., "Is this packet containing malware signature X?").

---

## Zero Trust Network Access (ZTNA)

The traditional model (VPN) is "crunchy on the outside, soft and chewy on the center." Once you VPN in, you often have broad network access.

**The ZTNA Shift:**
*   **Old Way (VPN):** Connect to network → Get IP → Access resource.
*   **New Way (ZTNA):** Authenticate Request → Proxy checks policy → Access Resource.

**Key Difference:** ZTNA authenticates *every single request* based on identity + context (device health, IP, time), regardless of network location. It removes the need for public Bastion hosts.

**Cloud Tools:**
*   **AWS:** Verified Access
*   **Azure:** Private Access / App Proxy
*   **GCP:** Identity-Aware Proxy (IAP) - The pioneer in this space (BeyondCorp).

---

## DNS Security

DNS is often the "forgotten" vector. Attackers love it because firewalls rarely block UDP 53.

### The Problem: DNS Tunneling / Exfiltration
Bad actors encode stolen data into DNS queries.
*   *Example:* `base64_stolen_data.attacker.com`
*   Your server queries this. The query bypasses your firewall (allowed outbound DNS). The attacker's authoritative name server logs the query and decodes the data.

### The Defense: DNS Firewall
*   **Mechanism:** Inspects every DNS query *before* it leaves your VPC.
*   **Tools:** Route 53 Resolver DNS Firewall, Azure DNS Private Resolver.
*   **Action:** Block queries to known bad domains (Command & Control) or look for high-entropy subdomains (tunneling signatures).

---

## Defense in Depth: Layered Network Security

Never rely on a single control. Layer them:

```
┌─────────────────────────────────────────────────────────────────────┐
│ Layer 1: Edge                                                       │
│ - CDN with WAF (CloudFront, Cloudflare, Akamai)                     │
│ - DDoS protection                                                   │
│ - Geographic blocking                                               │
├─────────────────────────────────────────────────────────────────────┤
│ Layer 2: Perimeter                                                  │
│ - Web Application Firewall (AWS WAF, Azure WAF)                     │
│ - API Gateway with throttling                                       │
│ - Bot detection                                                     │
├─────────────────────────────────────────────────────────────────────┤
│ Layer 3: VPC/Network                                                │
│ - Network ACLs (subnet level)                                       │
│ - VPC Flow Logs (visibility)                                        │
│ - Network segmentation                                              │
├─────────────────────────────────────────────────────────────────────┤
│ Layer 4: Resource                                                   │
│ - Security Groups (per-resource)                                    │
│ - Instance-level firewall (iptables/Windows Firewall)               │
├─────────────────────────────────────────────────────────────────────┤
│ Layer 5: Application                                                │
│ - TLS everywhere (even internal)                                    │
│ - Input validation                                                  │
│ - Authentication/Authorization                                      │
└─────────────────────────────────────────────────────────────────────┘
```

---

## Key Concepts to Remember

1. **Private by default** - Put resources in private subnets unless they MUST be public
2. **Security Groups are your primary control** - Use SG chaining, reference other SGs
3. **NACLs for blocklisting** - Block known bad IPs at subnet edge
4. **Use private endpoints** - Keep traffic off the internet, save money, reduce risk
5. **Transit Gateway for scale** - Hub-spoke beats VPC peering mesh
6. **Layer your controls** - Edge → Perimeter → Network → Resource → Application

---

## Practice Questions

**Q1:** Your security team discovers that a database in a private subnet has been sending data to an external IP address. The database has no public IP and is in a private subnet with no internet gateway route. How is this possible?

<details>
<summary>View Answer</summary>

**The database is likely using the NAT Gateway for outbound internet access.**

Even though the database is in a private subnet without a direct internet route, if the route table has a default route (0.0.0.0/0) pointing to a NAT Gateway in a public subnet, the database can initiate outbound connections.

**Attack scenario:**
1. Attacker gains access to the database (SQL injection, stolen credentials)
2. Uses built-in functions (e.g., `xp_cmdshell` in SQL Server) to make HTTP calls
3. Traffic flows: DB → NAT Gateway → Internet Gateway → Attacker's server

**Mitigations:**
- Remove NAT Gateway access for databases that don't need internet
- Use VPC endpoints for AWS service access instead
- Implement VPC Flow Logs and alert on unexpected outbound connections
- Use Security Groups to restrict outbound to only required destinations
- Network ACL deny rules for non-approved external IPs

</details>

**Q2:** You're designing a multi-tier application with web, app, and database tiers. The security team requires that the database tier cannot be accessed from the web tier under any circumstances—only the app tier should reach the database. However, a developer argues they need database access for debugging. How do you architect this?

<details>
<summary>View Answer</summary>

**Architecture:**

```
┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│ Web Tier        │────▶│ App Tier        │────▶│ DB Tier         │
│ sg-web          │     │ sg-app          │     │ sg-db           │
│                 │     │                 │     │                 │
│ Inbound: 443    │     │ Inbound: 8080   │     │ Inbound: 5432   │
│ from: 0.0.0.0/0 │     │ from: sg-web    │     │ from: sg-app    │
│                 │     │                 │     │ from: sg-bastion│
│ Outbound: 8080  │     │ Outbound: 5432  │     │                 │
│ to: sg-app      │     │ to: sg-db       │     │ Outbound: none  │
└─────────────────┘     └─────────────────┘     └─────────────────┘
```

**Key security controls:**

1. **Security Group sg-db** only allows port 5432 from sg-app and sg-bastion
2. **No rule** allows web tier to reach database tier
3. **NACL on DB subnet** explicitly denies traffic from web tier CIDR (defense in depth)

**For developer debugging:**
- Create a bastion host in a separate management subnet
- Bastion SG allows SSH from approved IPs only (or use SSM Session Manager—no bastion needed)
- Database SG allows connections from bastion SG
- All bastion sessions are logged and audited
- Implement JIT access—bastion only available when requested

**No standing access:** Developer must request access, which is time-limited and logged.

</details>

**Q3:** Your organization is planning to connect 25 VPCs across 3 AWS regions, plus VPN connections to 2 on-premises data centers. What architecture would you recommend, and what are the security considerations?

<details>
<summary>View Answer</summary>

**Recommended Architecture: Multi-Region Transit Gateway with Hub-Spoke**

```
Region 1 (Primary)              Region 2                Region 3
┌─────────────────┐         ┌─────────────────┐    ┌─────────────────┐
│ Transit GW 1    │◀───────▶│ Transit GW 2    │◀──▶│ Transit GW 3    │
│  │              │ Peering │  │              │    │  │              │
│  ├── Prod VPCs  │         │  ├── Prod VPCs  │    │  ├── Prod VPCs  │
│  ├── Dev VPCs   │         │  ├── Dev VPCs   │    │  └── Dev VPCs   │
│  ├── Shared Svc │         │  └── Shared Svc │    │                 │
│  │              │         │                 │    │                 │
│  ├── VPN to DC1 │         └─────────────────┘    └─────────────────┘
│  └── VPN to DC2 │
└─────────────────┘
```

**Security considerations:**

1. **Route table segmentation:**
   - Separate route tables for Prod, Dev, Shared Services
   - Prod VPCs cannot route to Dev VPCs
   - All traffic to on-prem goes through security inspection

2. **Centralized inspection:**
   - Deploy AWS Network Firewall in inspection VPC
   - All inter-VPC and outbound traffic routed through firewall
   - Use Gateway Load Balancer for third-party firewall appliances

3. **Traffic encryption:**
   - Inter-region TGW peering is encrypted by default
   - VPN connections use IPsec
   - Consider PrivateLink for sensitive service-to-service communication

4. **Logging and visibility:**
   - VPC Flow Logs on all VPCs
   - Transit Gateway Flow Logs
   - Centralized log aggregation to security SIEM

5. **Access control:**
   - Resource Access Manager to share TGW with specific accounts only
   - Attachment approval workflow
   - Terraform/CloudFormation for infrastructure as code

</details>

**Q4:** An auditor asks you to explain the difference between VPC peering and AWS PrivateLink. When would you use each?

<details>
<summary>View Answer</summary>

**VPC Peering:**
- Connects two VPCs at the network layer
- Allows any resource in VPC-A to potentially reach any resource in VPC-B
- Bidirectional by default
- No transitive routing (A-B-C doesn't mean A can reach C)
- Free (no per-hour or per-GB charge)

**AWS PrivateLink:**
- Connects a specific service/endpoint
- Consumer VPC gets an ENI with private IP pointing to provider's service
- Unidirectional (consumer initiates to provider)
- Provider controls exactly what's exposed
- Charged per hour + per GB

**When to use VPC Peering:**
- Full network connectivity between VPCs is needed
- Both VPCs are in the same organization with similar trust levels
- You need bidirectional communication
- Cost is a concern (it's free)

**When to use PrivateLink:**
- Exposing a specific service to customers/partners (SaaS)
- You want to avoid exposing your entire VPC
- Consumer and provider are different organizations
- You need to cross account/org boundaries with minimal exposure
- Overlapping CIDR ranges (PrivateLink works, peering doesn't)

**Example scenarios:**

| Scenario | Solution |
|----------|----------|
| Connect prod and shared-services VPCs in same org | VPC Peering |
| Expose your SaaS API to customer VPCs | PrivateLink (provider) |
| Access partner's service without internet | PrivateLink (consumer) |
| Full mesh of 5 internal VPCs | Transit Gateway |
| Connect to AWS services privately | VPC Endpoints (PrivateLink) |

</details>

**Q5: The Valid Static IP Mystery (AWS)**
**Scenario:** A client requires your application to have a static inbound IP address for their corporate firewall allowlist. You are using an Application Load Balancer (ALB). You check the documentation and realize ALBs don't support static IPs, but Network Load Balancers (NLBs) do.
**Why does this difference exist at the architectural level?**

<details>
<summary>View Answer</summary>

*   **Architecture Difference:** **Hyperplane.**
*   **Explanation:** ALB runs on a standard fleet of EC2 instances that scale in and out, changing IPs constantly. NLB runs on **AWS Hyperplane**, which manages connections using a massive distributed state across the zone. Hyperplane exposes a single, stable IP address per availability zone that routes traffic to the underlying fleet, allowing for static IPs.

</details>

**Q6: The "Subnet Delegation" Error (Azure)**
**Scenario:** You have a subnet `10.0.1.0/24` where you are running several Virtual Machines. You attempt to deploy an Azure Databricks workspace into this same subnet, but the deployment fails with a "Subnet Delegation" error.
**What is happening?**

<details>
<summary>View Answer</summary>

*   **The Issue:** **VNet Injection constraints.**
*   **Explanation:** To secure the PaaS service (Databricks) inside your private network, Azure uses VNet Injection. This often requires **delegating** the entire subnet to that specific service, meaning Azure takes control of the subnet's configuration. You cannot mix general-purpose VMs and delegated PaaS resources in the same subnet.

</details>

**Q7: The Global VPC (GCP vs AWS)**
**Scenario:** You need to deploy a database in Tokyo (`asia-northeast1`) and a web server in Iowa (`us-central1`). They need to talk to each other over private IP addresses.
**How does the setup differ between GCP and AWS?**

<details>
<summary>View Answer</summary>

*   **GCP:** You can put both resources in the **same VPC** (Global VPC). GCP's Andromeda virtualization allows subnets in different regions to coexist in the same logic network. Routing is automatic.
*   **AWS:** You must create **two separate VPCs** (one in Tokyo, one in Iowa) and connect them via **VPC Peering** or **Transit Gateway**. A single VPC cannot span regions.

</details>

**Q8: The Silent Exfiltration (Forensics)**
**Scenario:** You suspect an EC2 instance has been compromised and is exfiltrating sensitive credit card data to an external IP. You review the **VPC Flow Logs**.
**Can the Flow Logs prove that credit card data was stolen?**

<details>
<summary>View Answer</summary>

*   **Answer:** **No.**
*   **Why:** VPC Flow Logs are **metadata only**. They will prove that a connection occurred, how long it lasted, and how many bytes were transferred (Volume).
*   **Analysis:** If you see 5GB of data sent to a suspicious IP, you have *strong circumstantial evidence* of exfiltration, but you do not have *proof* of what the data was. To prove it was credit card data, you would have needed **Traffic Mirroring** (full packet capture) enabled *during* the event.

</details>

**Q9: The Mysterious CPU Spike (DDoS)**
**Scenario:** Your web application is experiencing extreme slowness. Your monitoring dashboard shows CPU usage at 100%, but your network bandwidth metrics are surprisingly low (normal levels).
**What kind of attack is this, and why didn't AWS Shield Standard block it?**

<details>
<summary>View Answer</summary>

*   **Attack Type:** **Layer 7 (Application) Attack / HTTP Flood.**
*   **Explanation:** The attacker is sending complex requests (like expensive SQL queries or login attempts) that consume CPU, rather than flooding the pipe with junk traffic (volumetric).
*   **Why Shield missed it:** AWS Shield **Standard** (free) focuses on Layer 3/4 volumetric attacks (SYN floods, UDP reflection). To block Layer 7 attacks, you need **AWS WAF** (Web Application Firewall) with Rate Limiting rules to block IPs sending too many requests.

</details>

**Q10: The Transit Gateway Leak**
**Scenario:** You have a "Prod" VPC and a "Dev" VPC attached to the same Transit Gateway. You have configured strict Security Groups on your Prod servers to only allow traffic from the Bastion host. However, a penetration tester running a port scan from a Dev server is successfully pinging your Prod database.
**How is this possible, and where is the misconfiguration?**

<details>
<summary>View Answer</summary>

*   **The Issue:** **Shared TGW Route Tables.**
*   **Explanation:** Even if Security Groups are tight, ICMP (ping) might be allowed by default or accident. The root cause is that both Prod and Dev VPCs are likely associated with the **same** Transit Gateway Route Table, which propagates routes for both. This means the network path exists.
*   **The Fix:** Create separate TGW Route Tables (e.g., "Prod Table" and "Dev Table"). Do **not** propagate Dev routes into the Prod table. This ensures complete network isolation at the routing layer, so packets cannot even attempt to cross boundaries.

</details>

**Q11: The Death of the Bastion Host**
**Scenario:** You are performing a security audit. You see 20 different Bastion Hosts (EC2 instances with public IPs and open SSH ports) scattered across various VPCs. You want to modernize this to align with **Zero Trust** principles.
**What architecture change do you recommend?**

<details>
<summary>View Answer</summary>

*   **Recommendation:** **Remove Bastion Hosts** and replace them with **Session Manager (AWS Systems Manager)** or **Identity-Aware Proxy (GCP IAP)**.
*   **Why:**
    1.  **Attack Surface:** Bastions listen on open ports (22) and have public IPs. Scanners hit them constantly.
    2.  **Identity:** SSH keys are hard to manage and rotate.
    3.  **Zero Trust:** Modern tools like Session Manager usually use an agent (SSM Agent) to open an outbound channel. There are **no inbound ports opened**, and access is controlled via IAM policy (Identity), not network reachability.

</details>

**Q12: The Data in the Domain Name (DNS Security)**
**Scenario:** Your firewall logs look clean (no connections to known malicious IPs). However, you see a strange pattern in your DNS logs. An internal server is making thousands of queries to subdomains like `user=admin&pass=123.badsite.com` and `file=confidential.pdf.badsite.com`.
**What is happening, and how do you stop it?**

<details>
<summary>View Answer</summary>

*   **The Attack:** **DNS Tunneling / Exfiltration.**
*   **Explanation:** The attacker is encoding stolen data into the *subdomain* portion of a DNS query. Since firewalls usually allow outbound UDP port 53 (DNS) to resolve legitimate names, this traffic bypasses standard IP blocking.
*   **The Defense:** Implement a **DNS Firewall** (like Route 53 Resolver DNS Firewall). It can block queries to known malicious domains or detect high-entropy domains (random strings) indicative of tunneling.

</details>

**Q13: The One-Way Street (Stateful vs Stateless)**
**Scenario:** You are troubleshooting connection issues to a web server in a public subnet. You start by opening the **Network ACL** completely (Allow All Inbound, Allow All Outbound) and everything works. Then, you lock it down: You allow Inbound Port 80, and you Deny everything else. Suddenly, users can't connect, even though the Inbound rule allows it.
**What did you forget?**

<details>
<summary>View Answer</summary>

*   **The Mistake:** Forgetting **Ephemeral Ports (Return Traffic).**
*   **Explanation:** NACLs are **stateless**. Just because you allowed the *request* to come IN on port 80 doesn't mean the *response* is allowed to go OUT. The response will leave on a high-numbered ephemeral port (1024-65535). You must explicitly create an **Outbound Rule** allowing traffic to destination ports 1024-65535.

</details>

**Q14: The Billion Dollar Bill (Endpoints)**
**Scenario:** Your application in a private subnet processes terabytes of data daily and stores it in S3. To keep traffic secure, you configure **Interface VPC Endpoints (PrivateLink)** for S3. At the end of the month, you receive a massive bill for "VPC Endpoint Processing Bytes."
**How could you have achieved the same security for free?**

<details>
<summary>View Answer</summary>

*   **Optimization:** Use a **Gateway Endpoint** for S3 instead.
*   **Explanation:** AWS offers two types of endpoints for S3.
    1.  **Interface Endpoints:** Cost money per hour + per GB processed.
    2.  **Gateway Endpoints:** **Free.** They work by adding a route entry to your VPC Route Table.
*   **Best Practice:** Always use Gateway Endpoints for S3 and DynamoDB (the only two services that support them) to save costs, unless you have a specific requirement like accessing S3 from on-premises via VPN/DX (which requires Interface Endpoints).

</details>

---

## Next Up

In Lesson 3, we'll explore **Cloud Security Posture Management (CSPM) & Cloud Workload Protection (CWPP)** — automated tools that find and fix misconfigurations before attackers do!
