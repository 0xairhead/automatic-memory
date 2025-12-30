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

---

## Next Up

In Lesson 3, we'll explore **Cloud Security Posture Management (CSPM) & Cloud Workload Protection (CWPP)** — automated tools that find and fix misconfigurations before attackers do!
