# Week 3, Lesson 7: Vertical vs Horizontal Scaling

## Table of Contents
- [What is Scalability?](#what-is-scalability)
- [The Two Scaling Approaches](#the-two-scaling-approaches)
- [Part 1: Vertical Scaling (Scale Up)](#part-1-vertical-scaling-scale-up)
  - [What is Vertical Scaling?](#what-is-vertical-scaling)
  - [Advantages of Vertical Scaling](#advantages-of-vertical-scaling)
  - [Disadvantages of Vertical Scaling](#disadvantages-of-vertical-scaling)
  - [When to Use Vertical Scaling](#when-to-use-vertical-scaling)
- [Part 2: Horizontal Scaling (Scale Out)](#part-2-horizontal-scaling-scale-out)
  - [What is Horizontal Scaling?](#what-is-horizontal-scaling)
  - [Architecture with Horizontal Scaling](#architecture-with-horizontal-scaling)
  - [Advantages of Horizontal Scaling](#advantages-of-horizontal-scaling)
  - [Disadvantages of Horizontal Scaling](#disadvantages-of-horizontal-scaling)
  - [When to Use Horizontal Scaling](#when-to-use-horizontal-scaling)
- [Stateful vs Stateless: Key to Horizontal Scaling](#stateful-vs-stateless-key-to-horizontal-scaling)
  - [Stateless Services (Easy to Scale)](#stateless-services-easy-to-scale)
  - [Stateful Services (Hard to Scale)](#stateful-services-hard-to-scale)
  - [Making Stateful Services Stateless](#making-stateful-services-stateless)
- [Scaling Databases](#scaling-databases)
  - [Vertical Scaling Databases](#vertical-scaling-databases)
  - [Horizontal Scaling Databases](#horizontal-scaling-databases)
- [Real-World Examples](#real-world-examples)
- [Auto-Scaling: Smart Horizontal Scaling](#auto-scaling-smart-horizontal-scaling)
  - [What is Auto-Scaling?](#what-is-auto-scaling)
  - [Auto-Scaling Metrics](#auto-scaling-metrics)
  - [Auto-Scaling Strategies](#auto-scaling-strategies)
- [Decision Framework: Vertical vs Horizontal](#decision-framework-vertical-vs-horizontal)
- [Common Mistakes](#common-mistakes)
- [Key Concepts to Remember](#key-concepts-to-remember)
- [Practice Questions](#practice-questions)
- [Next Up](#next-up)

---

Welcome to Week 3! Now we're entering the world of scalability—how to handle growth from hundreds to millions of users. This lesson covers one of the most fundamental decisions in system design.

---

## What is Scalability?

**Scalability = The ability to handle increased load**

Load could be:
- More users (10 → 10,000 → 10,000,000)
- More requests (100/sec → 100,000/sec)
- More data (1GB → 1TB → 1PB)
- More complexity (simple queries → complex analytics)

**The Goal:** System continues to work well as load increases

---

## The Two Scaling Approaches

```
Your server is overwhelmed! What do you do?

Option 1: VERTICAL SCALING (Scale Up)
└─> Buy a bigger server!

Option 2: HORIZONTAL SCALING (Scale Out)
└─> Buy more servers!
```

---

## Part 1: Vertical Scaling (Scale Up)

### What is Vertical Scaling?

**Add more resources to a single machine**

```
Current Server:
- 4 CPU cores
- 8GB RAM
- 500GB SSD

Scaled Up Server:
- 16 CPU cores ↑
- 64GB RAM ↑
- 2TB SSD ↑

Same machine, just BIGGER!
```

### Real Example

```
Monday:
[Server: 4 cores, 8GB RAM]
└─> Handling 1,000 users
└─> Response time: 100ms ✅

Friday (viral growth):
[Server: 4 cores, 8GB RAM]
└─> Trying to handle 10,000 users
└─> Response time: 5,000ms 😱
└─> CPU: 98%, RAM: 95%
└─> SERVER IS DYING!

Solution - Vertical Scale:
[New Server: 16 cores, 64GB RAM]
└─> Handling 10,000 users
└─> Response time: 120ms ✅
```

### Advantages of Vertical Scaling

**1. Simplicity**
```
No code changes needed!
Just shut down, upgrade hardware, restart.

Application still thinks it's one server.
No distributed system complexity.
```

**2. No Application Complexity**
```
Don't need to handle:
- Load balancing
- Data partitioning
- Distributed transactions
- Network latency between servers

Everything is local!
```

**3. Easier Data Consistency**
```
Single database on one machine
└─> ACID transactions work perfectly
└─> No distributed consensus needed
└─> No data synchronization issues
```

**4. Lower Latency**
```
All data in same machine
└─> No network calls
└─> RAM access: 100ns
└─> Local disk: microseconds

vs. Network call to another server: milliseconds
```

### Disadvantages of Vertical Scaling

**1. Hard Limits**
```
You can't scale forever!

Biggest AWS instance (as of 2024):
- 448 vCPUs
- 24TB RAM
- Cost: ~$100,000/month

What if you need more? You can't! 🚫
```

**2. Single Point of Failure**
```
One server = one point of failure

Server dies?
└─> Entire application DOWN! 😱
└─> No redundancy
└─> No failover
```

**3. Downtime for Upgrades**
```
To upgrade:
1. Shut down server
2. Add RAM/CPU
3. Restart
4. App is DOWN during this time! ⏱️

High availability? Not with one server!
```

**4. Expensive**
```
Cost doesn't scale linearly!

 4 cores,  8GB RAM = $100/month
16 cores, 64GB RAM = $800/month (8x cost for 4x capacity!)

Large machines are VERY expensive per unit of capacity
```

**5. Over-Provisioning**
```
Must buy capacity for peak load

Peak: 10,000 users
Average: 2,000 users

You're paying for:
└─> Capacity you only need 10% of the time!
└─> Wasted money 80% of the time
```

### When to Use Vertical Scaling

**Good for:**
- ✅ Small to medium applications
- ✅ Databases that need ACID transactions
- ✅ Applications not designed for distribution
- ✅ When you're just starting out
- ✅ Rapid prototyping

**Examples:**
- Early-stage startup (< 100,000 users)
- Single PostgreSQL database
- Traditional monolithic applications
- Development/test environments

---

## Part 2: Horizontal Scaling (Scale Out)

### What is Horizontal Scaling?

**Add more machines to share the load**

```
One Server:
[Server: 4 cores, 8GB RAM]
└─> Handling 1,000 users

Three Servers:
[Server 1: 4 cores, 8GB RAM] ┐
[Server 2: 4 cores, 8GB RAM] ├─> Handling 3,000 users
[Server 3: 4 cores, 8GB RAM] ┘

Each server: ~1,000 users
Total capacity: 3x!
```

### Architecture with Horizontal Scaling

```
                [Load Balancer]
                       │
        ┌──────────────┼──────────────┐
        ↓              ↓              ↓
   [Server 1]     [Server 2]     [Server 3]
        │              │              │
        └──────────────┼──────────────┘
                       ↓
                  [Database]

Load balancer distributes requests across servers
```

### Advantages of Horizontal Scaling

**1. No Hard Limits**
```
Need more capacity?
└─> Just add more servers!

1 server → 10 servers → 100 servers → 1,000 servers

Google, Facebook have millions of servers!
Practically unlimited scaling
```

**2. Better Fault Tolerance**
```
Server 2 crashes?

[Server 1] ✅
[Server 2] ❌ DEAD
[Server 3] ✅

Load balancer stops sending traffic to Server 2
Remaining servers handle the load
Application stays UP! 🎉
```

**3. No Downtime for Scaling**
```
Adding capacity:
1. Spin up new server
2. Load balancer detects it
3. Start sending traffic
4. Zero downtime! ✅

Removing capacity:
1. Stop sending traffic to server
2. Wait for active requests to finish
3. Shut down server
4. Zero downtime! ✅
```

**4. Cost Efficient**
```
Pay for what you need!

Low traffic (2AM): 2 servers running
High traffic (8PM): 10 servers running

Auto-scaling saves money!
Only pay for capacity you're using
```

**5. Geographic Distribution**
```
Servers in multiple regions:

US-East:    [Server 1, 2, 3]
US-West:    [Server 4, 5]
Europe:     [Server 6, 7]
Asia:       [Server 8, 9]

Users connect to nearest server
Lower latency globally!
```

### Disadvantages of Horizontal Scaling

**1. Application Complexity**
```
Must design for distribution!

Challenges:
- Session management
- Data consistency
- Transaction handling
- State management

Not all applications can easily scale horizontally
```

**2. Data Consistency Challenges**
```
Same user, different servers:

Request 1 → Server 1: Updates profile
Request 2 → Server 2: Reads profile (old data?) 🤔

Need strategies:
- Sticky sessions
- Distributed caching
- Database synchronization
```

**3. Network Overhead**
```
Servers communicate over network

Server-to-server calls:
└─> Add latency (5-50ms each)
└─> Can fail (network issues)
└─> Need retry logic

vs. Vertical: everything local (microseconds)
```

**4. Operational Complexity**
```
More servers = more problems!

Must manage:
- Deployment across servers
- Configuration consistency
- Monitoring all servers
- Debugging distributed issues
- Network infrastructure
```

**5. Data Partitioning Required**
```
Can't store all data on every server!

Must partition data (sharding):
- User data split across servers
- Complex query routing
- Join queries become difficult

Adds significant complexity
```

### When to Use Horizontal Scaling

**Good for:**
- ✅ Web applications
- ✅ APIs and microservices
- ✅ Stateless services
- ✅ Need high availability
- ✅ Unpredictable load patterns
- ✅ Global user base

**Examples:**
- Social media platforms
- E-commerce sites
- SaaS applications
- Streaming services

---

## Stateful vs Stateless: Key to Horizontal Scaling

### Stateless Services (Easy to Scale)

**No user-specific data stored on server**

```
Request 1 (User A) → Server 1
Request 2 (User A) → Server 2  ← Different server, NO PROBLEM!

Each request has everything needed:
- JWT token (who the user is)
- Request parameters

Server processes and responds
No need to remember anything
```

**Example - Stateless API:**
```javascript
// Every request is independent
GET /api/users/123
Authorization: Bearer eyJhbGciOiJIUzI1NiIs...

Server:
1. Validates token (stateless!)
2. Fetches user from database
3. Returns data
4. Forgets everything

Can go to ANY server!
```

**Benefits:**
- ✅ Any server can handle any request
- ✅ Easy to scale horizontally
- ✅ Load balancer can use any algorithm
- ✅ Server crashes don't affect users

### Stateful Services (Hard to Scale)

**Server stores user-specific data**

```
Request 1 (User A) → Server 1 (stores session)
Request 2 (User A) → Server 2 (no session!) ❌

Session data on Server 1!
Server 2 doesn't know User A!
```

**Example - Stateful:**
```javascript
// Login (creates session on server)
POST /login
└─> Server 1 stores: session[abc123] = {user: "alice"}

// Later request MUST go to Server 1
GET /dashboard
Cookie: session_id=abc123
└─> If goes to Server 2: NO SESSION DATA! ❌
```

**Problems:**
- ❌ Must route user to same server (sticky sessions)
- ❌ Server dies = user loses session
- ❌ Harder to load balance
- ❌ Can't easily add/remove servers

### Making Stateful Services Stateless

**Solution 1: Shared Session Store**
```
                [Load Balancer]
                       │
        ┌──────────────┼──────────────┐
        ↓              ↓              ↓
   [Server 1]     [Server 2]     [Server 3]
        │              │              │
        └──────────────┼──────────────┘
                       ↓
                  [Redis Cache]
                (shared sessions)

All servers read/write sessions from Redis!
Any server can handle any request!
```

**Solution 2: Client-Side Sessions (JWT)**
```
Instead of storing session on server:
└─> Store session in encrypted JWT token
└─> Send JWT with every request
└─> Server decodes JWT (no database lookup!)

Completely stateless!
```

---

## Scaling Databases

### Vertical Scaling Databases

```
Single PostgreSQL Server:
8 cores, 16GB RAM → 64 cores, 512GB RAM

Works well until:
- Too many connections
- Too much data
- Too many queries
- Hit hardware limits
```

**When vertical scaling works:**
- Read-heavy workloads (with read replicas)
- Small to medium datasets (< 1TB)
- Strong consistency requirements

### Horizontal Scaling Databases

More complex! Two main approaches:

#### 1. Read Replicas (Scale Reads)

```
        [Write] → [Primary DB]
                       │
                  (replication)
                       │
        ┌──────────────┼──────────────┐
        ↓              ↓              ↓
   [Replica 1]    [Replica 2]    [Replica 3]
        ↑              ↑              ↑
      [Read]        [Read]         [Read]

Writes: Go to primary
Reads: Distributed across replicas

Can scale reads infinitely!
```

**Limitations:**
- Only scales reads, not writes
- Replication lag (eventual consistency)
- Primary still bottleneck for writes

#### 2. Sharding (Scale Reads AND Writes)

```
Partition data across multiple databases:

Users 0-1M:     [Shard 1]
Users 1M-2M:    [Shard 2]
Users 2M-3M:    [Shard 3]

Application routes queries to correct shard
Each shard is independent

More in Lesson 8!
```

---

## Real-World Examples

### Example 1: Instagram (Horizontal Scaling)

```
Early Days (2010):
[1 Server] - 25,000 users

Growth (2011):
[Load Balancer]
├─> [Web Server 1]
├─> [Web Server 2]
└─> [Web Server 3]
└─> [Database] (vertical scaling)

Today:
Thousands of servers worldwide
Microservices architecture
Multiple data centers
Automatic scaling
```

### Example 2: Traditional Bank (Vertical Scaling)

```
Core Banking System:

[Mainframe]
- 256 cores
- 2TB RAM
- $5M+ cost

Why not horizontal?
- Legacy application (can't rewrite)
- ACID transactions critical
- Regulatory compliance
- Proven reliability
- 99.999% uptime

Sometimes vertical is the right choice!
```

### Example 3: Netflix (Hybrid Approach)

```
Stateless Components (Horizontal):
- API servers: Thousands
- Streaming servers: Auto-scale based on demand

Stateful Components (Vertical + Horizontal):
- Cassandra database: Horizontally scaled, partitioned
- User data: Replicated globally

Compute:
- Video encoding: Horizontal (batch jobs)
- Recommendation engine: Horizontal (microservices)
```

---

## Auto-Scaling: Smart Horizontal Scaling

### What is Auto-Scaling?

**Automatically add/remove servers based on load**

```
Traffic Pattern:
             ┌─High─┐
         ┌───┘      └───┐
         │              │
    ┌────┘              └────┐
────┘                        └────
 Low                             Low

Auto-scaling adds servers during high traffic
Removes servers during low traffic
Pay only for what you need!
```

### Auto-Scaling Metrics

**Scale based on:**

1. **CPU Utilization**
```
If average CPU > 70%:
└─> Add 2 servers

If average CPU < 30%:
└─> Remove 1 server
```

2. **Request Rate**
```
If requests/sec > 10,000:
└─> Add servers

If requests/sec < 2,000:
└─> Remove servers
```

3. **Response Time**
```
If P95 latency > 500ms:
└─> Add servers

If P95 latency < 100ms:
└─> Remove servers
```

4. **Custom Metrics**
```
Queue length, memory usage, database connections, etc.
```

### Auto-Scaling Strategies

#### 1. Reactive Scaling
```
Wait for metrics to breach threshold
Then add/remove servers

Pros: Cost-efficient
Cons: Slow to respond (5-10 min to add server)
```

#### 2. Predictive Scaling
```
Based on historical patterns:
"Traffic spikes every Friday 8 PM"
Pre-scale before the spike!

Pros: Ready before surge
Cons: More complex, might over-provision
```

#### 3. Scheduled Scaling
```
Known patterns:
"Scale up weekdays 9 AM - 5 PM"
"Scale down nights and weekends"

Pros: Simple, predictable
Cons: Can't handle unexpected spikes
```

---

## Decision Framework: Vertical vs Horizontal

### Start Vertical, Then Horizontal

**Phase 1: MVP (0-10K users)**
```
[Single Server - Vertical Scaling]

Why:
- Simple to build and deploy
- Fast development
- Low complexity
- Cost-effective at small scale

When to move on:
- Server maxing out
- Need high availability
- Planning for significant growth
```

**Phase 2: Growth (10K-100K users)**
```
[Load Balancer]
├─> [App Server 1]
├─> [App Server 2]
└─> [App Server 3]
└─> [Database] (still vertical)

Why:
- Application horizontally scaled
- Database still vertical (simpler)
- Good availability
- Reasonable cost

When to move on:
- Database becoming bottleneck
- Need global presence
```

**Phase 3: Scale (100K-1M+ users)**
```
[Global Load Balancer]
├─> [US Region]
│   ├─> [App Servers] (auto-scaling)
│   └─> [DB Shards] (horizontal)
├─> [EU Region]
└─> [Asia Region]

Why:
- Full horizontal scaling
- Multi-region
- High availability
- Unlimited growth potential
```

### Decision Matrix

| Factor | Vertical | Horizontal |
|--------|----------|------------|
| **Initial Complexity** | Low | High |
| **Max Scale** | Limited | Unlimited |
| **Availability** | Low | High |
| **Cost (small scale)** | Lower | Higher |
| **Cost (large scale)** | Much higher | Lower |
| **Data Consistency** | Easy | Challenging |
| **Downtime** | Required | Zero |
| **Development Speed** | Fast | Slower |

---

## Common Mistakes

### Mistake 1: Premature Horizontal Scaling
```
❌ Day 1: Build complex distributed system for 100 users

✅ Day 1: Single server
   Day 100: Scale when needed

Don't over-engineer early!
```

### Mistake 2: Ignoring Database
```
❌ Scale application servers, forget database

Bottleneck shifts to database!

✅ Scale application AND database together
```

### Mistake 3: Not Designing for Statelessness
```
❌ Store sessions on server
    Hard to scale later!

✅ Use Redis or JWT tokens
    Easy to scale anytime!
```

### Mistake 4: Forgetting About Data
```
❌ More servers, but data isn't partitioned
    All servers hit same database!

✅ Partition data appropriately
    True horizontal scaling
```

---

## Key Concepts to Remember

1. **Vertical Scaling** = Bigger machine (simple but limited)
2. **Horizontal Scaling** = More machines (complex but unlimited)
3. **Start vertical, scale horizontal** when needed
4. **Stateless services** scale easily horizontally
5. **Stateful services** need shared storage or sticky sessions
6. **Database scaling** is different from application scaling
7. **Auto-scaling** enables cost-efficient horizontal scaling
8. **No silver bullet** - most systems use both approaches
9. **Design for horizontal scaling** from the start (stateless)
10. **Monitor and scale proactively** before users notice

---

## Practice Questions

**Q1:** A startup has a single server with 8 cores and 16GB RAM serving 5,000 users. Traffic is growing 50% per month. Should they:
a) Upgrade to 32 cores and 128GB RAM?
b) Add 2 more servers behind a load balancer?

Justify your answer considering: current scale, growth rate, cost, complexity, and future needs.

<details>
<summary>View Answer</summary>

**Recommendation: Option B - Horizontal scaling**

**Growth projection:**
- Month 0: 5,000 users
- Month 3: ~17,000 users
- Month 6: ~57,000 users
- Month 12: ~650,000 users!

**Why horizontal scaling wins here:**

1. **Growth rate is aggressive:** 50% monthly growth means vertical scaling will hit limits within 6-12 months anyway

2. **Cost efficiency:**
   - 32 cores/128GB server: ~$800-1,200/month
   - 3 small servers (8 cores/16GB each): ~$300-400/month total
   - Horizontal is cheaper initially AND at scale

3. **Future-proofing:** Setting up load balancing now is easier than migrating later under pressure

4. **Availability:** 3 servers means the app stays up if one fails

5. **Learning opportunity:** Better to learn distributed systems patterns at 5K users than at 100K

**Vertical might be acceptable if:**
- Team has zero experience with distributed systems
- Application is deeply stateful and would require significant rewrite
- Growth plateaus unexpectedly

**Action plan:** Add load balancer + 2 servers now, ensure application is stateless, prepare for continued horizontal growth.

</details>

**Q2:** Your application stores user sessions in server memory. You want to scale horizontally by adding 3 more servers. What problems will you face? Propose 3 solutions.

<details>
<summary>View Answer</summary>

**The Problem:**
When sessions are stored in server memory, a user's session only exists on the server that created it. With multiple servers behind a load balancer, subsequent requests may route to different servers that don't have the session data.

**Example scenario:**
```
Request 1: Login → Server 1 (creates session: {user: "alice"})
Request 2: View profile → Server 3 (no session! User appears logged out)
```

**Three Solutions:**

**Solution 1: Sticky Sessions (Session Affinity)**
```
Load balancer always routes same user to same server
Based on: IP address, cookie, or session ID

Pros: Simple, no code changes
Cons: Uneven load distribution, server failure loses sessions
```

**Solution 2: Shared Session Store (Redis/Memcached)**
```
All servers read/write sessions from central Redis cluster

Pros:
- Any server can handle any request
- Sessions survive server failures
- Easy to scale

Cons:
- Additional infrastructure
- Network latency for session reads
- Redis becomes dependency
```

**Solution 3: Client-Side Sessions (JWT)**
```
Session data encoded in encrypted JWT token
Token sent with every request
Server validates and decodes token (no storage lookup)

Pros:
- Completely stateless servers
- No session storage needed
- Scales infinitely

Cons:
- Token size (can't store large data)
- Can't invalidate tokens easily
- Must handle token refresh
```

**Best practice:** Solution 2 (Redis) for most web apps, Solution 3 (JWT) for APIs and microservices.

</details>

**Q3:** Compare these two approaches for a news website expecting a traffic spike (Super Bowl night):

**Approach A:** One huge server (64 cores, 256GB RAM) running all the time
**Approach B:** Auto-scaling from 4 to 20 small servers (4 cores, 8GB RAM each)

Calculate approximate costs and discuss trade-offs.

<details>
<summary>View Answer</summary>

**Cost Analysis (approximate AWS pricing):**

**Approach A: Single Large Server**
```
Instance: ~m6i.16xlarge (64 vCPU, 256GB RAM)
Cost: ~$2.50/hour × 24 hours × 30 days = ~$1,800/month

Running 24/7 regardless of traffic
Annual cost: ~$21,600
```

**Approach B: Auto-Scaling**
```
Base: 4 servers × $0.17/hour = $0.68/hour
Peak: 20 servers × $0.17/hour = $3.40/hour

Normal days (29 days): 4 servers × 24h × $0.17 = ~$490
Super Bowl (1 day): Average 12 servers × 24h × $0.17 = ~$49

Monthly cost: ~$540
Annual cost: ~$6,500
```

**Cost winner: Approach B saves ~$15,000/year**

**Trade-off Analysis:**

| Factor | Approach A (Big Server) | Approach B (Auto-Scale) |
|--------|------------------------|-------------------------|
| **Cost** | $21,600/year | $6,500/year |
| **Complexity** | Simple | More complex |
| **Fault tolerance** | Single point of failure | High availability |
| **Scale ceiling** | Fixed at 64 cores | Practically unlimited |
| **Spike handling** | May not handle 50x spike | Scales to demand |
| **Setup time** | Quick | Requires infrastructure |

**Recommendation: Approach B**

Even accounting for engineering time to set up auto-scaling, the cost savings and improved reliability make horizontal scaling the clear winner for predictable traffic spikes.

</details>

**Q4:** A database handles 90% reads and 10% writes. Currently at capacity (1 server). What scaling strategy would you use? Why?

<details>
<summary>View Answer</summary>

**Best Strategy: Read Replicas**

**Why this is ideal for 90% reads:**

```
Architecture:
                    [Application Servers]
                           │
              ┌────────────┴────────────┐
              │                         │
         [Writes 10%]              [Reads 90%]
              │                         │
              ▼                         ▼
        [Primary DB] ──replication──> [Replica 1]
                     ──replication──> [Replica 2]
                     ──replication──> [Replica 3]
```

**How it works:**
- All writes go to primary database
- Reads distributed across 3+ read replicas
- Replication keeps replicas in sync (slight lag)

**Capacity calculation:**
```
Before: 1 server at 100% capacity
- 90% reads = 90 units of read work
- 10% writes = 10 units of write work

After (1 primary + 3 replicas):
- Primary handles: 10 units writes + some reads = ~30% capacity
- Each replica handles: 30 units reads = ~33% capacity

Result: 4x read capacity with minimal write overhead!
```

**Why NOT other strategies:**

- **Vertical scaling:** Works short-term but hits limits; doesn't add redundancy
- **Sharding:** Overkill for read-heavy workload; adds complexity for partitioning writes that aren't the bottleneck
- **Caching:** Good complement but doesn't solve all read patterns

**Additional optimization:** Add Redis cache in front of replicas for frequently accessed data (hot data caching).

</details>

**Q5:** Explain why Netflix can scale horizontally easily, but a banking system often cannot. What fundamental differences in requirements drive this?

<details>
<summary>View Answer</summary>

**Key Differences:**

| Requirement | Netflix | Banking |
|-------------|---------|---------|
| **Consistency** | Eventual (OK if recommendation is stale) | Strong (balance must be accurate) |
| **Transactions** | Simple (play video) | Complex (multi-step transfers) |
| **Failure impact** | Minor (retry, buffer) | Severe (lost money, legal issues) |
| **Data relationships** | Simple (user → videos) | Complex (accounts, transactions, audit) |
| **Regulatory** | Minimal | Heavy (SOX, PCI-DSS, banking laws) |

**Why Netflix Scales Easily:**

1. **Eventual consistency is acceptable:**
   ```
   User A sees 4.5 star rating
   User B sees 4.6 star rating (2 seconds later)

   Nobody cares! It's just a rating.
   ```

2. **Stateless operations:**
   ```
   Request: "Give me video chunk 47 of movie X"
   Any server can serve this! No state needed.
   ```

3. **Idempotent operations:**
   ```
   Playing video twice? No problem.
   Showing same recommendation twice? No problem.
   ```

4. **Failure tolerance:**
   ```
   Video buffers for 2 seconds? Users wait.
   Recommendation slightly outdated? Users don't notice.
   ```

**Why Banking Struggles:**

1. **ACID transactions are mandatory:**
   ```
   Transfer $100: A → B

   MUST be atomic:
   - Debit A: -$100
   - Credit B: +$100

   Both succeed or both fail. No exceptions!
   Distributed transactions are HARD.
   ```

2. **Strong consistency required:**
   ```
   Balance shows $500
   User withdraws $400
   Balance MUST show $100 immediately

   "Eventually $100" is unacceptable!
   ```

3. **Audit requirements:**
   ```
   Every transaction must be:
   - Logged permanently
   - Traceable
   - Consistent across all systems

   Distributed logging is complex.
   ```

4. **Regulatory compliance:**
   ```
   Must prove data integrity
   Must demonstrate controls
   Distributed systems harder to audit
   ```

**Summary:** Netflix optimizes for availability and partition tolerance (AP in CAP theorem). Banking optimizes for consistency and partition tolerance (CP), which is inherently harder to scale horizontally.

</details>

**Q6:** Your horizontally scaled application (5 servers) experiences this pattern:
- User logs in → Server 2
- User adds item to cart → Server 4 (cart is empty! 😱)

What went wrong? How do you fix it?

<details>
<summary>View Answer</summary>

**What Went Wrong:**

The application is **stateful** - cart data is stored in server memory. When the load balancer routes the second request to a different server, that server has no knowledge of the user's cart.

```
Request 1: POST /login → Server 2
           Server 2 memory: {user: "alice", cart: []}

Request 2: POST /cart/add → Server 4
           Server 4 memory: {} (empty! doesn't know about alice)
           Result: "Cart is empty" or "Please login"
```

**Fixes (in order of preference):**

**Fix 1: Externalize State to Redis (Best)**
```
All servers use shared Redis for cart data:

POST /cart/add (any server):
1. Get user ID from JWT/session
2. redis.LPUSH("cart:alice", item)
3. Return success

GET /cart (any server):
1. redis.LRANGE("cart:alice", 0, -1)
2. Return cart items

Any server can handle any request!
```

**Fix 2: Store Cart in Database**
```
Cart table in PostgreSQL/MySQL:
| user_id | product_id | quantity | added_at |

Pros: Persistent, survives restarts
Cons: Slower than Redis, more DB load
```

**Fix 3: Sticky Sessions (Quick Fix)**
```
Configure load balancer:
- Track user by cookie/IP
- Always route same user to same server

Pros: No code changes
Cons:
- Uneven load distribution
- Server death = lost carts
- Not truly scalable
```

**Fix 4: Client-Side Cart**
```
Store cart in browser localStorage or cookie
Send cart contents with each request

Pros: Zero server state
Cons:
- Limited size
- Lost if user clears browser
- Security concerns (price tampering)
```

**Recommended approach:** Redis for session/cart data + JWT for authentication. This gives you truly stateless servers that can scale infinitely.

</details>

**Q7:** Design a scaling strategy for an e-commerce site with these traffic patterns:
- Normal: 1,000 req/sec
- Black Friday: 50,000 req/sec
- Duration of spike: 24 hours

Consider costs, user experience, and engineering effort.

<details>
<summary>View Answer</summary>

**Scaling Strategy Design:**

**1. Architecture Overview**
```
                    [CloudFlare CDN]
                          │
                    [Global Load Balancer]
                          │
            ┌─────────────┼─────────────┐
            │             │             │
     [Web Tier]    [API Tier]    [Worker Tier]
     (Static)      (Dynamic)     (Background)
     Auto-scale    Auto-scale    Auto-scale
     2-20 servers  5-100 servers 2-20 servers
            │             │             │
            └─────────────┼─────────────┘
                          │
              ┌───────────┼───────────┐
              │           │           │
         [Redis]    [Primary DB]  [Read Replicas]
         (Cache)    + 3 replicas    (5 total)
```

**2. Auto-Scaling Configuration**
```yaml
Normal (1,000 req/sec):
  web_servers: 2
  api_servers: 5
  workers: 2

Black Friday (50,000 req/sec):
  web_servers: 20
  api_servers: 100
  workers: 20

Scaling triggers:
  scale_up: CPU > 60% OR response_time > 200ms
  scale_down: CPU < 30% AND response_time < 100ms
  cooldown: 5 minutes
```

**3. Pre-Black Friday Preparation**
```
1 week before:
- Pre-warm: Scale to 50% of expected peak
- Load test: Verify 50K req/sec capacity
- Cache warm-up: Pre-populate Redis with product data

1 hour before:
- Scale to 75% of expected peak
- Enable aggressive caching (longer TTLs)
- Disable non-critical features (recommendations)
```

**4. Cost Estimation**
```
Normal month (30 days at 1K req/sec):
- 5 API servers × $150 = $750
- 2 web servers × $100 = $200
- Database + Redis = $500
- Total: ~$1,450/month

Black Friday (1 day at 50K req/sec):
- 100 API servers × 24h × $0.20/h = $480
- 20 web servers × 24h × $0.13/h = $62
- Extra DB replicas = $100
- Total spike cost: ~$650

Annual cost:
- 11 normal months: $15,950
- Black Friday month: $2,100
- Total: ~$18,000/year
```

**5. User Experience Safeguards**
```
- Queue system for checkout (prevent overselling)
- Graceful degradation (disable reviews under load)
- Static fallback pages (if dynamic fails)
- Cart persistence (Redis with 24h TTL)
- Rate limiting (prevent abuse)
```

**6. Engineering Effort**
```
Initial setup: 2-3 weeks
- Configure auto-scaling groups
- Set up monitoring and alerts
- Implement feature flags for degradation
- Load testing infrastructure

Ongoing: 1-2 days/month
- Monitor and tune scaling rules
- Update capacity estimates
- Pre-Black Friday testing
```

**Key Success Metrics:**
- Response time < 500ms at peak
- Zero downtime during scaling events
- Checkout success rate > 99%
- Cost per transaction remains stable

</details>

---

## Next Up

In Lesson 8, we'll dive into **Database Replication & Sharding** - the techniques for scaling your database to handle massive amounts of data and traffic!
