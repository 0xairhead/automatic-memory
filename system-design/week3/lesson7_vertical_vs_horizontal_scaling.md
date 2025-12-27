# Week 3, Lesson 7: Vertical vs Horizontal Scaling

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

**Q2:** Your application stores user sessions in server memory. You want to scale horizontally by adding 3 more servers. What problems will you face? Propose 3 solutions.

**Q3:** Compare these two approaches for a news website expecting a traffic spike (Super Bowl night):

**Approach A:** One huge server (64 cores, 256GB RAM) running all the time
**Approach B:** Auto-scaling from 4 to 20 small servers (4 cores, 8GB RAM each)

Calculate approximate costs and discuss trade-offs.

**Q4:** A database handles 90% reads and 10% writes. Currently at capacity (1 server). What scaling strategy would you use? Why?

**Q5:** Explain why Netflix can scale horizontally easily, but a banking system often cannot. What fundamental differences in requirements drive this?

**Q6:** Your horizontally scaled application (5 servers) experiences this pattern:
- User logs in → Server 2
- User adds item to cart → Server 4 (cart is empty! 😱)

What went wrong? How do you fix it?

**Q7:** Design a scaling strategy for an e-commerce site with these traffic patterns:
- Normal: 1,000 req/sec
- Black Friday: 50,000 req/sec
- Duration of spike: 24 hours

Consider costs, user experience, and engineering effort.

---

Excellent work! You now understand one of the most fundamental concepts in system design. Next, we'll dive into database-specific scaling techniques! 🚀
