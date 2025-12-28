# Week 4, Lesson 12: CAP Theorem Explained Simply

## Table of Contents
- [The Distributed Systems Dilemma](#the-distributed-systems-dilemma)
- [What is the CAP Theorem?](#what-is-the-cap-theorem)
  - [C - Consistency](#c---consistency)
  - [A - Availability](#a---availability)
  - [P - Partition Tolerance](#p---partition-tolerance)
- [Why Can't We Have All Three?](#why-cant-we-have-all-three)
  - [The Network Partition Scenario](#the-network-partition-scenario)
  - [The Impossible Choice](#the-impossible-choice)
- [CP vs AP Systems](#cp-vs-ap-systems)
  - [CP Systems: Choose Consistency](#cp-systems-choose-consistency)
  - [AP Systems: Choose Availability](#ap-systems-choose-availability)
- [Real-World Examples](#real-world-examples)
  - [Banking Systems (CP)](#banking-systems-cp)
  - [Social Media (AP)](#social-media-ap)
  - [E-commerce (Mixed)](#e-commerce-mixed)
- [CAP in Popular Databases](#cap-in-popular-databases)
  - [CP Databases](#cp-databases)
  - [AP Databases](#ap-databases)
  - [Tunable Consistency](#tunable-consistency)
- [Beyond CAP: The PACELC Theorem](#beyond-cap-the-pacelc-theorem)
- [Consistency Models Spectrum](#consistency-models-spectrum)
  - [Strong Consistency](#strong-consistency)
  - [Eventual Consistency](#eventual-consistency)
  - [Causal Consistency](#causal-consistency)
  - [Read-Your-Writes Consistency](#read-your-writes-consistency)
- [Designing with CAP in Mind](#designing-with-cap-in-mind)
  - [Questions to Ask](#questions-to-ask)
  - [Hybrid Approaches](#hybrid-approaches)
- [Common Misconceptions](#common-misconceptions)
- [Key Concepts to Remember](#key-concepts-to-remember)
- [Practice Questions](#practice-questions)
- [Next Up](#next-up)

---

Welcome to the final lesson of Week 4! We're concluding our distribution fundamentals with the **CAP Theorem** - one of the most important concepts in distributed systems. Understanding CAP helps you make informed trade-offs when designing systems.

---

## The Distributed Systems Dilemma

When you have data on multiple servers, a fundamental problem arises:

```
Single Server (Easy):
┌─────────────┐
│   Server    │
│  balance:   │
│   $1000     │  ← One source of truth
└─────────────┘

Distributed System (Hard):
┌─────────────┐         ┌─────────────┐
│  Server A   │         │  Server B   │
│  balance:   │         │  balance:   │
│   $1000     │         │   $1000     │
└─────────────┘         └─────────────┘
       ↑                       ↑
       └───── Must stay in sync! ─────┘

What happens when:
- User withdraws $100 from Server A
- Network between A and B fails
- Another user checks balance on Server B

Does Server B show $1000 (stale) or refuse to answer?
```

This is the core dilemma that CAP theorem addresses.

---

## What is the CAP Theorem?

The CAP Theorem (proposed by Eric Brewer in 2000) states:

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│   A distributed system can provide at most TWO of these    │
│   three guarantees simultaneously:                          │
│                                                             │
│         ┌───────────────────┐                               │
│         │   Consistency     │                               │
│         │       (C)         │                               │
│         └─────────┬─────────┘                               │
│                  ╱ ╲                                        │
│                 ╱   ╲                                       │
│                ╱     ╲                                      │
│   ┌───────────┐       ┌───────────┐                        │
│   │Availability│       │ Partition │                        │
│   │    (A)    │       │ Tolerance │                        │
│   └───────────┘       │    (P)    │                        │
│                       └───────────┘                        │
│                                                             │
│   Pick TWO: CA, CP, or AP                                  │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### C - Consistency

**Every read receives the most recent write or an error.**

```
Consistent System:

Time 0: Balance = $1000 on all nodes
Time 1: User withdraws $100 from Node A
Time 2: Any node you ask gives $900

┌─────────┐     ┌─────────┐     ┌─────────┐
│ Node A  │     │ Node B  │     │ Node C  │
│  $900   │     │  $900   │     │  $900   │
└─────────┘     └─────────┘     └─────────┘
                    ↑
              Ask any node
              Same answer: $900
```

```
Inconsistent System:

Time 0: Balance = $1000 on all nodes
Time 1: User withdraws $100 from Node A
Time 2: Node A has $900, others still show $1000!

┌─────────┐     ┌─────────┐     ┌─────────┐
│ Node A  │     │ Node B  │     │ Node C  │
│  $900   │     │  $1000  │     │  $1000  │
└─────────┘     └─────────┘     └─────────┘
     ↑               ↑
  Updated        Stale data!
```

**Key point:** Consistency means all nodes see the same data at the same time.

### A - Availability

**Every request receives a response (not an error), without guarantee it's the most recent data.**

```
Available System:

User sends request → System responds (always)

┌─────────┐          ┌─────────┐
│  User   │ ──────►  │ System  │ ──────► Response!
└─────────┘          └─────────┘

Even if:
- Some nodes are down
- Network is slow
- Data might be stale

The system ALWAYS responds.
```

```
Unavailable System:

User sends request → Error or timeout

┌─────────┐          ┌─────────┐
│  User   │ ──────►  │ System  │ ──────► "Service Unavailable"
└─────────┘          └─────────┘
                           │
                     Can't guarantee
                     correct answer,
                     so refuses to answer
```

**Key point:** Availability means the system always responds, even during failures.

### P - Partition Tolerance

**The system continues to operate despite network partitions.**

```
Network Partition:

Normal:
┌─────────┐ ◄────────► ┌─────────┐
│ Node A  │   network  │ Node B  │
└─────────┘            └─────────┘

Partitioned (network failure):
┌─────────┐     ✗      ┌─────────┐
│ Node A  │ ◄──────►   │ Node B  │
└─────────┘   broken   └─────────┘

Nodes can't communicate!
But system must still work.
```

**Real partition causes:**
- Network cable cut
- Router failure
- Datacenter network issues
- Cloud provider problems
- Firewall misconfiguration

**Key point:** Partition tolerance means surviving network failures between nodes.

---

## Why Can't We Have All Three?

### The Network Partition Scenario

Let's walk through why you must choose:

```
Setup: Two nodes, both have balance = $1000

┌─────────────┐                    ┌─────────────┐
│   Node A    │ ◄───── Network ───►│   Node B    │
│ balance:    │                    │ balance:    │
│   $1000     │                    │   $1000     │
└─────────────┘                    └─────────────┘
```

**Step 1: Network partition occurs**

```
┌─────────────┐        ✗          ┌─────────────┐
│   Node A    │ ◄─────────────►   │   Node B    │
│   $1000     │    PARTITION!     │   $1000     │
└─────────────┘                   └─────────────┘

Nodes can no longer communicate!
```

**Step 2: Write request arrives at Node A**

```
User: "Withdraw $100 from my account"

┌─────────────┐        ✗          ┌─────────────┐
│   Node A    │ ◄─────────────►   │   Node B    │
│   $1000     │                   │   $1000     │
└─────────────┘                   └─────────────┘
      ▲
      │
    User: "Withdraw $100"
```

### The Impossible Choice

Now Node A has two options:

**Option 1: Be Consistent (CP) - Reject the write**

```
Node A: "I can't sync with Node B, so I'll refuse this request"

┌─────────────┐        ✗          ┌─────────────┐
│   Node A    │ ◄─────────────►   │   Node B    │
│   $1000     │                   │   $1000     │
│   LOCKED    │                   │             │
└─────────────┘                   └─────────────┘
      │
      ▼
   User: "Error: Service unavailable"

✅ Consistent: Both nodes agree ($1000)
❌ Not Available: User's request rejected
✅ Partition Tolerant: System handled the partition
```

**Option 2: Be Available (AP) - Accept the write**

```
Node A: "I'll process this even without syncing to B"

┌─────────────┐        ✗          ┌─────────────┐
│   Node A    │ ◄─────────────►   │   Node B    │
│   $900      │                   │   $1000     │
│  (updated)  │                   │   (stale)   │
└─────────────┘                   └─────────────┘
      │
      ▼
   User: "Success! Balance: $900"

❌ Not Consistent: Nodes disagree ($900 vs $1000)
✅ Available: User's request succeeded
✅ Partition Tolerant: System handled the partition
```

**Why not CA (Consistent + Available)?**

```
CA would mean:
- Always consistent (all nodes agree)
- Always available (always respond)
- BUT: Can't handle partitions

This only works if you guarantee NO network failures.

In distributed systems, network partitions WILL happen.
So you MUST be partition tolerant.

CA is essentially a single-node system!
```

**The real choice is: CP or AP?**

---

## CP vs AP Systems

### CP Systems: Choose Consistency

```
┌─────────────────────────────────────────────────────────────┐
│                    CP System Behavior                        │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  During normal operation:                                   │
│  ✅ Consistent                                              │
│  ✅ Available                                               │
│                                                             │
│  During network partition:                                  │
│  ✅ Consistent (refuse stale reads/writes)                  │
│  ❌ Available (some requests fail)                          │
│                                                             │
│  Philosophy: "Better to give no answer than wrong answer"   │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

**Use CP when:**
- Data accuracy is critical
- Wrong data causes serious problems
- Users prefer errors over incorrect info

**Examples:**
- Bank account balances
- Inventory counts (prevent overselling)
- Distributed locks
- Configuration management

**CP Databases:**
- MongoDB (with majority write concern)
- HBase
- Redis Cluster
- Zookeeper
- Consul

### AP Systems: Choose Availability

```
┌─────────────────────────────────────────────────────────────┐
│                    AP System Behavior                        │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  During normal operation:                                   │
│  ✅ Consistent (eventually)                                 │
│  ✅ Available                                               │
│                                                             │
│  During network partition:                                  │
│  ❌ Consistent (nodes may have different data)              │
│  ✅ Available (always respond)                              │
│                                                             │
│  Philosophy: "Better to give stale answer than no answer"   │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

**Use AP when:**
- Availability is critical
- Stale data is acceptable temporarily
- System must always respond

**Examples:**
- Social media feeds
- Product catalogs
- DNS
- Shopping carts
- Session stores

**AP Databases:**
- Cassandra
- DynamoDB
- CouchDB
- Riak

---

## Real-World Examples

### Banking Systems (CP)

```
Why banks choose Consistency over Availability:

Scenario: User has $100, tries to withdraw $100 twice

AP System (BAD for banks):
┌─────────┐     Partition     ┌─────────┐
│ ATM A   │ ◄───────✗───────► │ ATM B   │
│ $100    │                   │ $100    │
└─────────┘                   └─────────┘
     │                              │
  Withdraw                      Withdraw
   $100                          $100
     │                              │
     ▼                              ▼
   "Success"                    "Success"
   ($0 left)                    ($0 left)

Result: Bank gave out $200 from $100 account! 💸

CP System (GOOD for banks):
┌─────────┐     Partition     ┌─────────┐
│ ATM A   │ ◄───────✗───────► │ ATM B   │
│ $100    │                   │ $100    │
└─────────┘                   └─────────┘
     │                              │
  Withdraw                      Withdraw
   $100                          $100
     │                              │
     ▼                              ▼
   "Success"                   "Error:
   ($0 left)                    Cannot process,
                                try again later"

Result: Only $100 withdrawn, data consistent! ✅
```

### Social Media (AP)

```
Why Twitter/Facebook choose Availability:

Scenario: Celebrity posts tweet, network partition occurs

CP System (BAD for social media):
User in Europe: "Error: Cannot load feed"
User in Asia: "Error: Cannot load feed"

Millions of users see errors!
Terrible user experience.

AP System (GOOD for social media):
User in Europe: Sees feed (maybe missing latest tweet)
User in Asia: Sees feed (maybe missing latest tweet)

Everyone can use the app!
Tweet will appear once partition heals.

Trade-off:
- User might see 999 likes instead of 1000
- User might not see newest post for a few seconds
- Much better than showing error!
```

### E-commerce (Mixed)

```
Smart e-commerce uses BOTH:

┌─────────────────────────────────────────────────────────────┐
│                    E-commerce Architecture                   │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  Product Catalog: AP                                        │
│  - Always show products                                     │
│  - Price might be slightly stale                            │
│  - Better than empty page!                                  │
│                                                             │
│  Shopping Cart: AP                                          │
│  - Always let users add items                               │
│  - Sync when possible                                       │
│  - Don't lose sales!                                        │
│                                                             │
│  Checkout/Payment: CP                                       │
│  - Must have accurate inventory                             │
│  - Must process payment correctly                           │
│  - Okay to show "try again" on failure                      │
│                                                             │
│  Reviews: AP                                                │
│  - Show reviews even if slightly stale                      │
│  - New review might take seconds to appear                  │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

---

## CAP in Popular Databases

### CP Databases

| Database | How it achieves CP |
|----------|-------------------|
| **MongoDB** | Majority write concern requires acknowledgment from majority of nodes |
| **HBase** | Single master for writes, strong consistency |
| **Zookeeper** | Consensus protocol (ZAB), leader-based |
| **etcd** | Raft consensus, linearizable reads |
| **Redis Cluster** | Stops writes if can't reach majority |

```
MongoDB with Majority Write Concern:

Write "balance: $900"
        │
        ▼
┌─────────────┐
│   Primary   │ ──► Write to primary
└─────────────┘
        │
        ├──────────────┬──────────────┐
        ▼              ▼              ▼
┌─────────────┐ ┌─────────────┐ ┌─────────────┐
│ Secondary 1 │ │ Secondary 2 │ │ Secondary 3 │
└─────────────┘ └─────────────┘ └─────────────┘
        │              │
        └──────┬───────┘
               │
          2 of 3 ACK
               │
               ▼
        "Write successful"

If can't reach majority → Write fails (CP behavior)
```

### AP Databases

| Database | How it achieves AP |
|----------|-------------------|
| **Cassandra** | Tunable consistency, can read/write with single node |
| **DynamoDB** | Eventually consistent reads by default |
| **CouchDB** | Multi-master replication, conflict resolution |
| **Riak** | Vector clocks, sibling resolution |

```
Cassandra with Consistency Level ONE:

Write "balance: $900"
        │
        ▼
┌─────────────┐
│   Node 1    │ ──► Write successful!
│   (any)     │     (immediately returns)
└─────────────┘
        │
        │ (async replication)
        ├──────────────┬──────────────┐
        ▼              ▼              ▼
┌─────────────┐ ┌─────────────┐ ┌─────────────┐
│   Node 2    │ │   Node 3    │ │   Node 4    │
│ (eventually)│ │ (eventually)│ │ (eventually)│
└─────────────┘ └─────────────┘ └─────────────┘

Write succeeds even if other nodes unreachable (AP behavior)
```

### Tunable Consistency

Many modern databases let you choose per-operation:

```
Cassandra Consistency Levels:

┌─────────────────────────────────────────────────────────────┐
│  Level        │ Behavior                      │ CAP        │
├───────────────┼───────────────────────────────┼────────────┤
│ ONE           │ 1 node responds               │ AP         │
│ QUORUM        │ Majority responds             │ CP         │
│ ALL           │ All nodes respond             │ Strong CP  │
│ LOCAL_QUORUM  │ Majority in local datacenter  │ Regional   │
└─────────────────────────────────────────────────────────────┘

Example:
// Fast, eventually consistent read (AP)
SELECT * FROM users WHERE id = 123;
CONSISTENCY ONE;

// Strong consistent read (CP)
SELECT * FROM users WHERE id = 123;
CONSISTENCY QUORUM;
```

```
DynamoDB Consistency Options:

// Eventually consistent (default, faster, cheaper)
dynamodb.get_item(
    TableName='Users',
    Key={'user_id': '123'},
    ConsistentRead=False  # AP
)

// Strongly consistent (slower, costs more)
dynamodb.get_item(
    TableName='Users',
    Key={'user_id': '123'},
    ConsistentRead=True   # CP
)
```

---

## Beyond CAP: The PACELC Theorem

CAP only describes behavior during partitions. PACELC extends this:

```
┌─────────────────────────────────────────────────────────────┐
│                      PACELC Theorem                          │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  IF there's a Partition (P):                                │
│      Choose between Availability (A) and Consistency (C)    │
│                                                             │
│  ELSE (normal operation):                                   │
│      Choose between Latency (L) and Consistency (C)         │
│                                                             │
│  Full form: PAC / ELC                                       │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

**Why this matters:**

```
During partition: CP vs AP (as CAP says)

During NORMAL operation (no partition):
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  Strong Consistency:                                        │
│  Write → Sync to all nodes → Return                        │
│  Latency: 50-200ms (must wait for sync)                    │
│                                                             │
│  Eventual Consistency:                                      │
│  Write → Return immediately → Async sync                   │
│  Latency: 5-20ms (no waiting)                              │
│                                                             │
└─────────────────────────────────────────────────────────────┘

Even without partitions, there's a consistency/latency trade-off!
```

**Database classification with PACELC:**

| Database | P+A/C | E+L/C | Meaning |
|----------|-------|-------|---------|
| Cassandra | PA | EL | Favors availability and low latency |
| DynamoDB | PA | EL | Same as Cassandra |
| MongoDB | PC | EC | Favors consistency always |
| MySQL (async) | PA | EL | Availability and speed |
| MySQL (sync) | PC | EC | Consistency always |

---

## Consistency Models Spectrum

Consistency isn't binary. There's a spectrum:

```
Strong ◄──────────────────────────────────────────► Weak

Linearizable → Sequential → Causal → Read-your-writes → Eventual
     │              │           │            │              │
 Strictest      Ordered     Logical      Practical      Loosest
                           ordering       minimum
```

### Strong Consistency

```
Linearizable (Strictest):
- Operations appear instantaneous
- Global ordering of all operations
- If write completes, all subsequent reads see it

Timeline:
T1: Write X=1       ─────────────────────►
T2:                      Read X → Must return 1
T3:                                   Read X → Must return 1

Used in: Zookeeper, etcd, Spanner
Cost: High latency, lower availability
```

### Eventual Consistency

```
Eventually Consistent (Loosest):
- Updates propagate eventually
- No ordering guarantees
- Reads may return stale data

Timeline:
T1: Write X=1 to Node A   ─────────────────────►
T2:     Read X from Node B → May return old value!
T3:               Read X from Node B → Eventually returns 1

Used in: DNS, Cassandra (default), DynamoDB (default)
Cost: Low latency, high availability
```

### Causal Consistency

```
Causal Consistency (Middle ground):
- Causally related operations are ordered
- Unrelated operations may be seen in any order

Example:
Alice posts: "I got the job!" (Post A)
Alice comments: "So excited!" (Comment B, caused by A)

Causal consistency guarantees:
- If you see Comment B, you must have seen Post A
- You won't see "So excited!" without "I got the job!"

Used in: MongoDB (causal sessions), CockroachDB
```

### Read-Your-Writes Consistency

```
Read-Your-Writes:
- You always see your own writes
- Others may not see them yet

User Alice:
Write "profile_pic = cat.jpg"    ─────────────────────►
Read profile                      → Always sees cat.jpg ✓

User Bob (same time):
Read Alice's profile              → May still see old pic

Used in: Most web apps (via sticky sessions or read-after-write)
Practical minimum for good UX.
```

---

## Designing with CAP in Mind

### Questions to Ask

```
1. What happens if we show stale data?
   - Annoying but okay → Consider AP
   - Dangerous/costly → Must be CP

2. What happens if service is unavailable?
   - Users wait and retry → CP is okay
   - Users leave forever → Need AP

3. How often do partitions occur?
   - Rarely (single datacenter) → Less critical
   - Frequently (global) → Very important

4. Can we make it right later?
   - Yes (compensating transactions) → AP with reconciliation
   - No (can't un-send email) → Need CP for that operation

5. What's our SLA?
   - 99.99% uptime required → AP with eventual consistency
   - Correctness > availability → CP acceptable
```

### Hybrid Approaches

**Approach 1: Different consistency per operation**

```python
class UserService:
    def update_email(self, user_id, email):
        # CP: Email must be consistent (used for auth)
        self.db.write(
            consistency='QUORUM',
            data={'email': email}
        )

    def update_profile_bio(self, user_id, bio):
        # AP: Bio can be eventually consistent
        self.db.write(
            consistency='ONE',
            data={'bio': bio}
        )
```

**Approach 2: Write-ahead log for eventual consistency**

```
User places order → Write to local DB → Queue for sync

┌─────────────────┐        ┌─────────────────┐
│   Region A      │        │   Region B      │
│                 │        │                 │
│ ┌─────────────┐ │        │ ┌─────────────┐ │
│ │ Order: 123  │ │───────►│ │ Order: 123  │ │
│ │ (immediate) │ │ async  │ │ (eventual)  │ │
│ └─────────────┘ │        │ └─────────────┘ │
└─────────────────┘        └─────────────────┘

User sees order immediately (AP)
Globally consistent eventually
Reconciliation handles conflicts
```

**Approach 3: CRDTs for automatic conflict resolution**

```
CRDT = Conflict-free Replicated Data Type

Example: Counter that can increment on any node

Node A: counter = 5, increment → 6
Node B: counter = 5, increment → 6

Traditional: CONFLICT! 6 vs 6, which is right?

CRDT approach:
Node A: {A: 1} increments → {A: 2}
Node B: {B: 1} increments → {B: 2}
Merged: {A: 2, B: 2} → total = 4

No conflicts! Both increments preserved.

Used in: Riak, Redis CRDT, collaborative editing
```

---

## Common Misconceptions

### Misconception 1: "You must choose two, can't have all three"

```
Misleading because:
- During NORMAL operation, you can have all three!
- The choice only matters during PARTITIONS
- Partitions are temporary, not permanent

Reality:
┌─────────────────────────────────────────────────────────────┐
│  Normal: C + A + P (sort of)                                │
│  - Consistent reads                                         │
│  - Available                                                │
│  - No partition happening, so P is trivial                  │
│                                                             │
│  Partition: Must choose C or A                              │
│  - P is forced upon you                                     │
│  - Decide: errors (C) or stale data (A)?                    │
└─────────────────────────────────────────────────────────────┘
```

### Misconception 2: "CA systems exist"

```
NOT in distributed systems!

CA would mean:
- Consistent
- Available
- NOT partition tolerant

But in a distributed system, partitions WILL happen.
If you can't handle partitions, you're not distributed.

"CA" systems are really:
- Single-node databases
- Or systems that become unavailable during partitions
```

### Misconception 3: "AP means inconsistent forever"

```
AP systems are EVENTUALLY consistent, not NEVER consistent.

Timeline:
T0: Partition occurs
T1: Writes go to different nodes (inconsistent)
T2: Partition heals
T3: Nodes sync up
T4: Consistency restored!

AP doesn't mean "wrong data forever"
It means "temporarily stale, eventually correct"
```

### Misconception 4: "Consistency = ACID consistency"

```
CAP Consistency ≠ ACID Consistency

CAP Consistency:
- All nodes see same data
- About replication

ACID Consistency:
- Database enforces constraints
- About transactions

They're different concepts with the same word!
```

### Misconception 5: "Latency and availability are the same"

```
They're related but different:

High Latency, Available:
Request → Wait 30 seconds → Response
(System responded, just slowly)

Unavailable:
Request → Error/Timeout
(System didn't respond)

A slow response is still available!
But users might not see it that way...

This is why PACELC matters.
```

---

## Key Concepts to Remember

1. **CAP Theorem**: In a distributed system, during a partition, you must choose between Consistency and Availability

2. **Partition Tolerance is mandatory** for distributed systems - partitions will happen

3. **CP systems** reject requests during partitions to stay consistent

4. **AP systems** serve requests during partitions with potentially stale data

5. **The real choice** is what to do when the network fails

6. **Eventual consistency** means data will converge once the partition heals

7. **Most systems are hybrid** - different consistency for different operations

8. **PACELC** extends CAP: Even without partitions, there's a latency/consistency trade-off

9. **Consistency is a spectrum** from linearizable (strongest) to eventual (weakest)

10. **Design for your requirements** - not every operation needs the same consistency level

---

## Practice Questions

**Q1:** A bank is designing a new system for checking account balances. Should they use a CP or AP database? What about for viewing transaction history?

<details>
<summary>View Answer</summary>

**Account Balance: CP (Consistency)**

```
Why CP for balance:

Scenario: User has $100
- ATM A: Shows $100, user decides to withdraw $80
- ATM B (partition): Shows $100, user's spouse withdraws $80

With AP: Both withdrawals succeed → Overdraft! 💸
With CP: Second withdrawal fails → "Please try again"

User experience:
- Brief error is acceptable
- Incorrect balance is NOT acceptable (could cause overdrafts)
- Financial accuracy is legally required
```

**Transaction History: AP (Availability)**

```
Why AP for history:

- User wants to see past transactions
- Showing slightly stale history is fine
- Better to show "last 30 transactions" than error
- New transactions can appear with small delay

Trade-off acceptable:
✅ User can always see their history
⚠️ Newest transaction might take 2-3 seconds to appear
❌ "Service unavailable" error

This is READ-ONLY data - no risk of inconsistent actions
```

**Architecture:**

```
┌─────────────────────────────────────────────────────────────┐
│                    Banking System                            │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  Balance Service (CP)          History Service (AP)         │
│  ┌─────────────────────┐      ┌─────────────────────┐      │
│  │ PostgreSQL          │      │ Cassandra           │      │
│  │ Synchronous         │      │ Eventually          │      │
│  │ Replication         │      │ Consistent          │      │
│  │                     │      │                     │      │
│  │ Operations:         │      │ Operations:         │      │
│  │ - Check balance     │      │ - View history      │      │
│  │ - Withdraw          │      │ - Search trans.     │      │
│  │ - Transfer          │      │ - Download statement│      │
│  └─────────────────────┘      └─────────────────────┘      │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

</details>

**Q2:** During a network partition between two datacenters, your CP database stops accepting writes. Users are angry. Your boss asks "Why can't we just accept the writes and fix inconsistencies later?" What's your response?

<details>
<summary>View Answer</summary>

**Short Answer:**

"We could, but the inconsistencies might be unfixable, and the damage could be worse than temporary unavailability."

**Detailed Explanation:**

```
Scenario: E-commerce during partition

Datacenter A (New York)         Datacenter B (London)
┌─────────────────────┐         ┌─────────────────────┐
│ Inventory: 1 item   │    ✗    │ Inventory: 1 item   │
│                     │ PARTITION│                     │
│ Customer 1: "Buy"   │         │ Customer 2: "Buy"   │
│ → Inventory: 0      │         │ → Inventory: 0      │
│ → Order confirmed!  │         │ → Order confirmed!  │
└─────────────────────┘         └─────────────────────┘

Partition heals:
- Inventory is now: -1 (IMPOSSIBLE!)
- Two customers paid for ONE item
- Someone doesn't get their product
```

**The "Fix It Later" Problems:**

```
1. Some inconsistencies can't be fixed:
   - Both customers charged → Must refund one
   - But who? First customer? Random?
   - One customer gets angry email: "Sorry, we oversold"
   - Legal issues, reputation damage

2. Compensation is expensive:
   - Refund processing fees
   - Customer service time
   - Potential lawsuits
   - Lost customer trust

3. Some operations are irreversible:
   - Sent emails can't be unsent
   - Triggered webhooks to partners
   - Started physical fulfillment
   - Notified downstream systems

4. Reconciliation is complex:
   - Which write wins?
   - Need conflict resolution logic
   - Edge cases everywhere
   - Developer time to build and maintain
```

**When AP Actually Works:**

```
If your data CAN be reconciled:

✅ Like counts: Just add them up
✅ User preferences: Last write wins (usually fine)
✅ Shopping cart: Merge items from both sides
✅ Comments: Include all, sort by timestamp

But inventory/financial data:
❌ Can't create items out of thin air
❌ Can't double-spend money
❌ Business rules prevent merging
```

**Recommended Response to Boss:**

```
"We have two options:

Option A (Current - CP):
- Writes fail during partition
- ~5 minutes of errors per year
- Zero data inconsistencies
- No angry 'we oversold' emails

Option B (AP with reconciliation):
- Writes always succeed
- But we'll sometimes oversell
- Need refund process
- Customer service cost: $X per incident
- Reputation risk

For inventory/payments, Option A is safer.
For non-critical data (reviews, likes), we can switch to AP."
```

</details>

**Q3:** Explain how Cassandra can be configured as either CP or AP depending on consistency level settings. Give examples of when you'd use each.

<details>
<summary>View Answer</summary>

**Cassandra Consistency Levels:**

```
Cassandra has tunable consistency per query:

Replication Factor (RF) = 3 (data on 3 nodes)

Write Consistency Levels:
┌─────────────┬────────────────────────────────────────────┐
│ Level       │ Behavior                                   │
├─────────────┼────────────────────────────────────────────┤
│ ONE         │ Write to 1 node, return success           │
│ TWO         │ Write to 2 nodes, return success          │
│ QUORUM      │ Write to majority (2 of 3), return success│
│ ALL         │ Write to all 3 nodes, return success      │
└─────────────┴────────────────────────────────────────────┘

Read Consistency Levels:
┌─────────────┬────────────────────────────────────────────┐
│ Level       │ Behavior                                   │
├─────────────┼────────────────────────────────────────────┤
│ ONE         │ Read from 1 node                          │
│ QUORUM      │ Read from majority, return latest         │
│ ALL         │ Read from all nodes, return latest        │
└─────────────┴────────────────────────────────────────────┘
```

**AP Configuration (High Availability):**

```
Write: ONE
Read: ONE

┌─────────┐  Write  ┌─────────┐
│ Client  │────────►│ Node 1  │──► Success! (immediate)
└─────────┘         └─────────┘
                         │
                         │ (async replication)
                         ▼
                    Nodes 2, 3 (later)

Read:
┌─────────┐  Read   ┌─────────┐
│ Client  │────────►│ Node 2  │──► May return stale data!
└─────────┘         └─────────┘

AP because:
- Writes succeed if ANY node is up
- Reads succeed if ANY node is up
- Data might be inconsistent during partition
```

**When to use AP (ONE/ONE):**

```python
# Time-series metrics (can lose some points)
session.execute(
    "INSERT INTO metrics (ts, value) VALUES (?, ?)",
    consistency_level=ConsistencyLevel.ONE
)

# User activity logs (eventually consistent is fine)
session.execute(
    "INSERT INTO activity_log (user_id, action, ts) VALUES (?, ?, ?)",
    consistency_level=ConsistencyLevel.ONE
)

# Session storage (availability > consistency)
session.execute(
    "SELECT * FROM sessions WHERE session_id = ?",
    consistency_level=ConsistencyLevel.ONE
)
```

**CP Configuration (Strong Consistency):**

```
Write: QUORUM
Read: QUORUM

Formula: R + W > RF guarantees consistency
QUORUM + QUORUM = 2 + 2 = 4 > 3 ✓

┌─────────┐  Write  ┌─────────┐
│ Client  │────────►│ Node 1  │──┐
└─────────┘         └─────────┘  │
                                  │
                    ┌─────────┐  │
                    │ Node 2  │◄─┤ Must wait for 2 of 3
                    └─────────┘  │
                                  │
                    ┌─────────┐  │
                    │ Node 3  │  │ (optional)
                    └─────────┘  │
                                  │
                    Success! ◄───┘

Read:
┌─────────┐  Read   ┌─────────┐
│ Client  │────────►│ Node 1  │──┐
└─────────┘         └─────────┘  │
                                  │
                    ┌─────────┐  │
                    │ Node 2  │◄─┤ Read from 2, compare
                    └─────────┘  │ return latest!
                                  │
                    ──────────────┘

CP because:
- At least one node in read quorum saw the latest write
- Always returns most recent data
- Fails if majority unavailable (partition)
```

**When to use CP (QUORUM/QUORUM):**

```python
# User balance (must be accurate)
session.execute(
    "UPDATE users SET balance = ? WHERE user_id = ?",
    consistency_level=ConsistencyLevel.QUORUM
)

# Inventory count (prevent overselling)
session.execute(
    "UPDATE products SET stock = ? WHERE product_id = ?",
    consistency_level=ConsistencyLevel.QUORUM
)

# Read balance before allowing purchase
session.execute(
    "SELECT balance FROM users WHERE user_id = ?",
    consistency_level=ConsistencyLevel.QUORUM
)
```

**Mixed Configuration Example:**

```python
class ProductService:
    def update_stock(self, product_id, new_stock):
        # CP: Stock accuracy is critical
        self.session.execute(
            "UPDATE products SET stock = ? WHERE product_id = ?",
            (new_stock, product_id),
            consistency_level=ConsistencyLevel.QUORUM
        )

    def update_description(self, product_id, description):
        # AP: Description can be eventually consistent
        self.session.execute(
            "UPDATE products SET description = ? WHERE product_id = ?",
            (description, product_id),
            consistency_level=ConsistencyLevel.ONE
        )

    def get_product(self, product_id, need_accurate_stock=False):
        cl = ConsistencyLevel.QUORUM if need_accurate_stock else ConsistencyLevel.ONE
        return self.session.execute(
            "SELECT * FROM products WHERE product_id = ?",
            (product_id,),
            consistency_level=cl
        )
```

</details>

**Q4:** A global social media app has users in US, Europe, and Asia. They're deciding between:
- Single database in US (strong consistency)
- Replicas in each region (eventual consistency)

What are the trade-offs? What would you recommend?

<details>
<summary>View Answer</summary>

**Option A: Single Database in US**

```
┌─────────────────────────────────────────────────────────────┐
│                 Single US Database                           │
└─────────────────────────────────────────────────────────────┘

US Users ─────► [US Database] ◄───── Europe Users
    5ms              ▲                   150ms
                     │
               Asia Users
                  200ms

Latency:
- US users: ~5ms (local)
- Europe users: ~150ms (cross-Atlantic)
- Asia users: ~200ms (cross-Pacific)
```

**Pros:**
- Strong consistency (one source of truth)
- Simple architecture
- No conflict resolution needed
- Easy to reason about

**Cons:**
- High latency for non-US users
- Single point of failure
- All load on one region
- Poor UX for 60%+ of users

---

**Option B: Regional Replicas (Eventually Consistent)**

```
┌─────────────────────────────────────────────────────────────┐
│               Multi-Region Architecture                      │
└─────────────────────────────────────────────────────────────┘

     ┌─────────────┐       ┌─────────────┐       ┌─────────────┐
     │ US Replica  │◄─────►│Europe Replic│◄─────►│Asia Replica │
     │             │ async │             │ async │             │
     └─────────────┘       └─────────────┘       └─────────────┘
           ▲                     ▲                     ▲
           │ 5ms                 │ 5ms                 │ 5ms
           │                     │                     │
       US Users            Europe Users           Asia Users

Latency:
- US users: ~5ms
- Europe users: ~5ms
- Asia users: ~5ms
- All users get local performance!
```

**Pros:**
- Low latency for all users
- Regional fault tolerance
- Better user experience
- Load distributed globally

**Cons:**
- Eventually consistent
- Temporary inconsistencies across regions
- Need conflict resolution
- More complex operations

---

**Recommendation: Regional Replicas with Caveats**

For social media, eventual consistency is acceptable:

```
Why it works for social media:

1. Posts/Comments:
   - New post takes 2-3 seconds to appear globally
   - Users don't notice small delays
   - No financial risk

2. Likes/Reactions:
   - Count might be 999 instead of 1000 briefly
   - Converges quickly
   - Nobody cares about exact real-time count

3. Followers:
   - Eventually consistent follower count is fine
   - User won't notice if count updates slowly

4. DMs (special case):
   - Need stronger consistency
   - Use CP for message ordering
   - Or route both users to same region
```

**Hybrid Architecture:**

```
┌─────────────────────────────────────────────────────────────┐
│               Recommended Architecture                       │
└─────────────────────────────────────────────────────────────┘

Public Content (AP - Regional):
┌───────────┐     ┌───────────┐     ┌───────────┐
│ US Region │◄───►│ EU Region │◄───►│Asia Region│
│ Posts     │async│ Posts     │async│ Posts     │
│ Comments  │     │ Comments  │     │ Comments  │
│ Likes     │     │ Likes     │     │ Likes     │
└───────────┘     └───────────┘     └───────────┘

User Account Data (CP - Primary + Replicas):
┌───────────┐     ┌───────────┐     ┌───────────┐
│ US Primary│────►│ EU Replica│────►│Asia Replic│
│ Passwords │sync │ Read-only │sync │ Read-only │
│ Email     │     │           │     │           │
│ Settings  │     │           │     │           │
└───────────┘     └───────────┘     └───────────┘

Direct Messages (CP - Route to User's Region):
User A (US) messages User B (EU)
→ Route to EU region (User B's home)
→ Strong consistency for message ordering
→ User A sees slight latency (acceptable for DMs)
```

**Conflict Resolution Strategy:**

```python
# For posts/comments: Last-Write-Wins with timestamp
def resolve_conflict(version_a, version_b):
    return version_a if version_a.timestamp > version_b.timestamp else version_b

# For likes: CRDT Counter (merge by adding)
def resolve_like_count(region_counts):
    # Each region tracks its own likes
    # Total = sum of all regions
    return sum(region_counts.values())

# For follower lists: Set union
def resolve_followers(sets):
    return set.union(*sets)
```

**Summary:**

| Feature | Consistency | Approach |
|---------|-------------|----------|
| Posts/Feed | Eventual | AP, regional |
| Likes/Comments | Eventual | AP, CRDT |
| User Profile | Eventual | AP, LWW |
| Authentication | Strong | CP, single primary |
| Direct Messages | Strong | CP, routed |
| Payments | Strong | CP, single region |

</details>

**Q5:** Describe a scenario where choosing AP over CP actually results in a worse user experience, despite the system being "available."

<details>
<summary>View Answer</summary>

**Scenario: Airline Seat Selection with AP**

```
Setup:
- Flight has 1 window seat left (14A)
- Two users trying to book simultaneously
- AP system: Always responds, eventually consistent
```

**What Happens with AP:**

```
Time 0: Seat 14A available in both regions

US Region                        EU Region
┌──────────────────┐            ┌──────────────────┐
│ Seat 14A: FREE   │            │ Seat 14A: FREE   │
└──────────────────┘            └──────────────────┘
        │                               │
     User A                          User B
  "Select 14A"                    "Select 14A"
        │                               │
        ▼                               ▼

Time 1: Both succeed (AP - available!)

US Region                        EU Region
┌──────────────────┐            ┌──────────────────┐
│ Seat 14A: USER_A │            │ Seat 14A: USER_B │
│                  │            │                  │
│ "Confirmed! ✓"   │            │ "Confirmed! ✓"   │
└──────────────────┘            └──────────────────┘

Both users see: "Seat 14A confirmed!"
Both are happy... temporarily.
```

**The Horrible UX:**

```
Time 2: Systems sync, conflict detected

┌─────────────────────────────────────────────────────────────┐
│               Conflict Resolution Required!                  │
│                                                             │
│  Option 1: First-write-wins                                 │
│  → User B gets email: "Sorry, your seat was given away"     │
│  → User B now has NO seat (middle seat remaining)           │
│  → User B furious! "It said CONFIRMED!"                     │
│                                                             │
│  Option 2: Last-write-wins                                  │
│  → User A gets email: "Sorry, your seat was given away"     │
│  → Same problem, different victim                           │
│                                                             │
│  Option 3: Random winner                                    │
│  → Someone loses their "confirmed" seat                     │
│  → Completely arbitrary and unfair                          │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

**User B's Experience:**

```
Timeline:
10:00 - Selected seat 14A
10:00 - "Seat confirmed! ✓"
10:01 - Showed spouse: "Got the window seat!"
10:30 - Received email: "Booking update"
10:30 - "Your seat has been changed to 22B (middle)"
10:31 - RAGE 😤

User B's perspective:
"The system CONFIRMED my seat!"
"I planned my whole trip around that seat!"
"This is bait-and-switch!"
"I'm never flying this airline again!"
```

**Why CP Would Be Better Here:**

```
With CP:

Time 0: Seat 14A available

User A (first by 50ms)         User B (second by 50ms)
"Select 14A"                   "Select 14A"
     │                              │
     ▼                              ▼
Lock acquired ✓               Lock denied ✗
"Confirmed!"                  "Sorry, seat taken.
                               Please choose another."

User B's experience:
- Immediate feedback
- Can choose another seat
- No false confirmation
- Disappointed but not deceived
```

**Why AP Failed Here:**

```
AP is bad when:

1. Resources are scarce (one seat, one item)
2. "Confirmation" has meaning to users
3. Rollback causes worse UX than immediate failure
4. Users make decisions based on the response
   (User B told spouse, planned around window seat)

The "availability" was a lie:
- System was "available" to accept booking
- But booking wasn't actually guaranteed
- Delayed failure is worse than immediate failure
```

**Other Scenarios Where AP Hurts UX:**

```
1. Limited inventory flash sales:
   "Order confirmed!" → "Sorry, actually sold out"

2. Event ticket booking:
   "Tickets secured!" → "We oversold, you're waitlisted"

3. Hotel last-room booking:
   "Room confirmed!" → "Actually, you need to find another hotel"

4. Auction bidding:
   "You won!" → "Actually, someone outbid you"

5. Username registration:
   "Username reserved!" → "Actually, someone else got it"
```

**Lesson:**

```
AP is NOT always better user experience.

Consider:
- Is a delayed rejection worse than immediate rejection?
- Does "confirmation" matter to the user?
- Can the user take action based on the confirmation?
- Is there limited inventory/resources?

For scarce resources with meaningful confirmations,
CP provides better UX despite occasional errors.
```

</details>

**Q6:** Your team is building a collaborative document editor (like Google Docs). Analyze the CAP trade-offs for: (1) document content, (2) cursor positions, (3) document permissions.

<details>
<summary>View Answer</summary>

**Overview:**

```
Collaborative Editor Components:
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  1. Document Content                                        │
│     The actual text users are editing                       │
│                                                             │
│  2. Cursor Positions                                        │
│     Where each user's cursor is located                     │
│                                                             │
│  3. Document Permissions                                    │
│     Who can view/edit the document                          │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

---

**1. Document Content: AP with CRDTs**

```
Requirement Analysis:
- Users type continuously (high write frequency)
- Must never lose user's typing
- Offline editing should work
- Sync when connection restored
- Conflicts WILL happen (two users edit same line)

Why AP:
- Users expect typing to always work
- "Cannot save" error = unacceptable UX
- Offline mode is essential
- Temporary inconsistency is okay
```

**Implementation with CRDTs:**

```
CRDT: Conflict-free Replicated Data Type

User A types "Hello"        User B types "Hi"
at position 0               at position 0
        │                         │
        ▼                         ▼
   ┌─────────┐               ┌─────────┐
   │ Hello   │               │ Hi      │
   └─────────┘               └─────────┘
        │                         │
        └───────────┬─────────────┘
                    │ Merge
                    ▼
            ┌──────────────┐
            │ HelloHi      │  (or "HiHello")
            │              │
            │ Both edits   │
            │ preserved!   │
            └──────────────┘

CRDTs used:
- RGA (Replicated Growable Array) for text
- Unique IDs per character
- Deterministic merge rules
```

**CAP Choice: AP**
- Available: Always accept edits
- Eventually consistent: Merge on sync
- Partition tolerant: Works offline

---

**2. Cursor Positions: AP with Weak Consistency**

```
Requirement Analysis:
- Shows where other users are editing
- Updates very frequently (every keystroke)
- Stale cursor position is annoying but not critical
- Missing cursor update = minor UX issue
- High volume, low importance

Why AP:
- Not worth blocking for cursor sync
- Stale position is tolerable (user will update soon)
- Loss of cursor updates is okay
```

**Implementation:**

```
Cursor updates via WebSocket:

User A moves cursor
        │
        ▼
┌─────────────────┐
│ Broadcast to    │
│ other users     │
│                 │
│ Best-effort     │
│ No ACK required │
│ Fire and forget │
└─────────────────┘
        │
        ├─────────────────────────┐
        ▼                         ▼
   User B sees               User C sees
   cursor (maybe)            cursor (maybe)

If message lost:
- Next cursor movement will update
- No need to retry
- At-most-once delivery is fine
```

**CAP Choice: AP (weakest form)**
- Available: Always send updates
- Weakly consistent: May miss updates
- Best-effort delivery

---

**3. Document Permissions: CP**

```
Requirement Analysis:
- Controls who can access document
- Security-critical (wrong permissions = data leak)
- Changes rarely
- Must be accurate

Why CP:
- Incorrect permissions = security breach
- User removed from doc should NOT see content
- Error is better than wrong access
```

**Implementation:**

```
Permission check before any action:

User requests document
        │
        ▼
┌─────────────────────┐
│ Permission Service  │
│                     │
│ Check: Can user X   │
│ access document Y?  │
│                     │
│ MUST be consistent! │
└─────────────────────┘
        │
        ├── Yes ──► Serve document
        │
        └── No ──► Access denied
        │
        └── Error ──► "Please try again"
                      (NOT "here's the doc anyway")

Permission update:
Admin removes User B
        │
        ▼
┌─────────────────────┐
│ Sync to all nodes   │
│ BEFORE confirming   │
│                     │
│ Write concern: ALL  │
└─────────────────────┘
        │
        ▼
User B's next request → Denied immediately
(Not "denied in a few seconds")
```

**CAP Choice: CP**
- Consistent: All nodes agree on permissions
- Unavailable during partition: Deny access if unsure
- Partition tolerant: Handled by denying

---

**Combined Architecture:**

```
┌─────────────────────────────────────────────────────────────┐
│              Collaborative Editor Architecture               │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  ┌─────────────────────────────────────────────────────┐   │
│  │              Permission Service (CP)                 │   │
│  │  - PostgreSQL with sync replication                 │   │
│  │  - Check on every request                           │   │
│  │  - Deny if partition                                │   │
│  └─────────────────────────────────────────────────────┘   │
│                          │                                  │
│                          │ Auth check                       │
│                          ▼                                  │
│  ┌─────────────────────────────────────────────────────┐   │
│  │               Document Service (AP)                  │   │
│  │  - CRDT-based content storage                       │   │
│  │  - Accept writes even offline                       │   │
│  │  - Merge on reconnection                            │   │
│  │  - Yjs or Automerge library                         │   │
│  └─────────────────────────────────────────────────────┘   │
│                          │                                  │
│                          │ Real-time sync                   │
│                          ▼                                  │
│  ┌─────────────────────────────────────────────────────┐   │
│  │              Presence Service (AP-weak)              │   │
│  │  - WebSocket broadcast                              │   │
│  │  - Cursor positions                                 │   │
│  │  - User online status                               │   │
│  │  - Best-effort, no persistence                      │   │
│  └─────────────────────────────────────────────────────┘   │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

**Summary Table:**

| Component | CAP | Consistency | Reason |
|-----------|-----|-------------|--------|
| Permissions | CP | Strong | Security-critical |
| Content | AP | Eventual (CRDT) | Must always work, offline support |
| Cursors | AP | Weak | Best-effort, not critical |

</details>

**Q7:** A startup is building a ride-sharing app. During a network partition between their two datacenters, should they continue matching riders with drivers or stop the service? Justify your decision.

<details>
<summary>View Answer</summary>

**Recommendation: Continue service (AP) with safety measures**

**Why AP (Continue Matching):**

```
Business Reality:
- Each minute of downtime = lost revenue
- Drivers go offline if app "doesn't work"
- Riders switch to competitors immediately
- Reputation damage from outage

User Expectations:
- "I need a ride NOW"
- Users don't care about internal partitions
- Error message = open competitor app
- Availability is the core value proposition
```

**The Risk of CP (Stop Service):**

```
Partition occurs (5 minutes):

With CP:
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│   5:00 PM: Partition starts                                 │
│   5:00 PM: "Cannot request ride at this time"               │
│   5:01 PM: Users open Lyft                                  │
│   5:02 PM: Drivers: "App is broken, switching to Lyft"      │
│   5:05 PM: Partition heals                                  │
│   5:05 PM: "Service restored!" (but users are gone)         │
│                                                             │
│   Cost: Lost rides, lost drivers, reputation damage         │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

**How to Make AP Safe:**

```
Potential Problem: Double-matching
- Driver matched to Rider A in DC1
- Same driver matched to Rider B in DC2
- Partition heals: driver has 2 rides!

Solution: Regional affinity + conflict resolution

┌─────────────────────────────────────────────────────────────┐
│                 Ride-Sharing During Partition                │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│   DC1 (East Coast)              DC2 (West Coast)            │
│   ┌─────────────────┐          ┌─────────────────┐         │
│   │ Drivers: D1-D50 │          │ Drivers: D51-D100│         │
│   │ Match locally   │    ✗     │ Match locally    │         │
│   │                 │ PARTITION │                  │         │
│   └─────────────────┘          └─────────────────┘         │
│                                                             │
│   Rule: Each DC matches only LOCAL drivers                  │
│   - D1 (East) can only match in DC1                        │
│   - D51 (West) can only match in DC2                       │
│   - No double-matching possible!                            │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

**Conflict Resolution for Edge Cases:**

```python
class RideMatchingService:
    def match_rider(self, rider, available_drivers):
        # 1. Only consider drivers in our region
        local_drivers = [d for d in available_drivers
                        if d.home_region == self.region]

        # 2. Create match with unique ID
        match = Match(
            id=generate_uuid(),
            rider=rider,
            driver=best_driver(local_drivers),
            created_at=now(),
            region=self.region
        )

        # 3. Store locally (AP)
        self.local_db.save(match)

        # 4. Queue for sync (when partition heals)
        self.sync_queue.push(match)

        return match

    def resolve_conflicts_after_partition(self):
        # In rare case of double-match (driver changed regions):
        for driver in drivers_with_multiple_matches():
            matches = get_matches_for_driver(driver)

            # Keep earliest match, cancel others
            matches.sort(key=lambda m: m.created_at)
            keep = matches[0]

            for match in matches[1:]:
                cancel_match(match)
                notify_rider(match.rider,
                    "Your driver was reassigned. Finding new driver...")
                rematch_rider(match.rider)  # Priority re-matching
```

**Additional Safety Measures:**

```
1. Driver location updates:
   - Continue tracking driver GPS
   - Store locally during partition
   - Sync after partition heals
   - Rider sees accurate driver position

2. Payment processing:
   - Queue payment for after ride
   - Process when connectivity confirmed
   - Don't lose the charge!

3. Trip state machine:
   - REQUESTED → MATCHED → EN_ROUTE → IN_PROGRESS → COMPLETED
   - Each transition logged locally
   - Reconcile states after partition

4. Monitoring:
   - Alert on partition detection
   - Track match rate during partition
   - Measure conflict rate after heal
   - Learn and improve thresholds
```

**What About Payments? (The CP Part)**

```
Hybrid approach:

Matching & Trip: AP (must keep working)
├── Local matching
├── Local state management
└── Queue for sync

Payment Processing: CP (must be accurate)
├── Hold payment when matched (auth)
├── Charge only after trip confirmed
└── If conflict: refund the cancelled match

┌─────────────────────────────────────────────────────────────┐
│                    Trip State Machine                        │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  REQUESTED ──► MATCHED ──► EN_ROUTE ──► IN_PROGRESS ──► END │
│     │            │                           │              │
│   [Local]     [Local]                     [Local]           │
│     │            │                           │              │
│     └────────────┴───────────────────────────┘              │
│                          │                                  │
│                     After partition:                        │
│                     Sync & reconcile                        │
│                          │                                  │
│                    COMPLETED ──► PAYMENT (CP)               │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

**Summary:**

```
Decision: AP (Continue Service)

Justification:
1. Core business requires availability
2. Users have zero tolerance for downtime
3. Competitors are one tap away
4. Regional affinity prevents most conflicts
5. Conflict resolution handles edge cases
6. Payment can be CP (processed after trip)

Trade-off:
- ~0.1% of rides might need re-matching after partition
- Compensation: priority re-match + potential credit
- Better than 100% of users seeing errors!
```

</details>

---

## Next Up

Congratulations on completing Week 4: Distribution Basics! You now understand consistent hashing, message queues, and the fundamental CAP theorem trade-offs.

In Week 5, we'll dive into **System Design Patterns** starting with **Microservices vs Monoliths** - learning when to break apart your system and when to keep it together!
