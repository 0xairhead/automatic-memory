# Week 3, Lesson 8: Database Replication & Sharding

## Table of Contents
- [Why Database Scaling is Different](#why-database-scaling-is-different)
- [Part 1: Database Replication](#part-1-database-replication)
  - [What is Replication?](#what-is-replication)
  - [Why Replicate?](#why-replicate)
- [Replication Strategy 1: Master-Slave (Primary-Replica)](#replication-strategy-1-master-slave-primary-replica)
  - [How It Works](#how-it-works)
  - [The Replication Process](#the-replication-process)
  - [Synchronous vs Asynchronous Replication](#synchronous-vs-asynchronous-replication)
  - [Handling Primary Failure (Failover)](#handling-primary-failure-failover)
  - [Reading from Replicas: Consistency Concerns](#reading-from-replicas-consistency-concerns)
- [Replication Strategy 2: Multi-Master (Active-Active)](#replication-strategy-2-multi-master-active-active)
  - [Why Multi-Master?](#why-multi-master)
  - [The Conflict Problem](#the-conflict-problem)
  - [Conflict Resolution Strategies](#conflict-resolution-strategies)
  - [When to Use Multi-Master](#when-to-use-multi-master)
- [Part 2: Database Sharding](#part-2-database-sharding)
  - [What is Sharding?](#what-is-sharding)
  - [Why Sharding?](#why-sharding)
- [Sharding Strategy 1: Range-Based Sharding](#sharding-strategy-1-range-based-sharding)
- [Sharding Strategy 2: Hash-Based Sharding](#sharding-strategy-2-hash-based-sharding)
- [Sharding Strategy 3: Consistent Hashing](#sharding-strategy-3-consistent-hashing)
- [Sharding Strategy 4: Directory-Based Sharding](#sharding-strategy-4-directory-based-sharding)
- [Sharding Challenges](#sharding-challenges)
  - [Challenge 1: Cross-Shard Queries](#challenge-1-cross-shard-queries)
  - [Challenge 2: Cross-Shard Joins](#challenge-2-cross-shard-joins)
  - [Challenge 3: Cross-Shard Transactions](#challenge-3-cross-shard-transactions)
  - [Challenge 4: Rebalancing](#challenge-4-rebalancing)
  - [Challenge 5: Schema Changes](#challenge-5-schema-changes)
- [Combining Replication and Sharding](#combining-replication-and-sharding)
- [Real-World Examples](#real-world-examples)
- [Decision Framework](#decision-framework)
- [Common Mistakes](#common-mistakes)
- [Key Concepts to Remember](#key-concepts-to-remember)
- [Practice Questions](#practice-questions)
- [Next Up](#next-up)

---

In the previous lesson, we learned about scaling applications horizontally. But what about databases? This lesson covers the two most important techniques for scaling databases: **replication** and **sharding**.

---

## Why Database Scaling is Different

**Applications are (mostly) stateless. Databases are stateful.**

```
Application Scaling (Easy):
[Server 1] ─┐
[Server 2] ─┼─> Any server can handle any request
[Server 3] ─┘    No coordination needed!

Database Scaling (Hard):
[DB 1] ─┐
[DB 2] ─┼─> Data must be consistent!
[DB 3] ─┘    Where does data live?
             What happens on writes?
```

**The Challenge:**
- Data must not be lost
- Data must be consistent
- Queries must find the right data
- Transactions must work correctly

---

## Part 1: Database Replication

### What is Replication?

**Keeping copies of the same data on multiple machines**

```
Primary Database (Source of Truth)
        │
        │ (replication)
        ↓
┌───────┴───────┐
↓               ↓
[Replica 1]  [Replica 2]

All three have the SAME data!
```

### Why Replicate?

**1. High Availability**
```
Primary dies?
        ↓
Replica becomes new primary!
        ↓
Zero data loss, minimal downtime
```

**2. Scale Reads**
```
Before (1 server):
[Primary] ← 10,000 reads/sec 😰

After (3 servers):
[Primary]  ← 3,333 reads/sec ✅
[Replica 1] ← 3,333 reads/sec ✅
[Replica 2] ← 3,333 reads/sec ✅
```

**3. Geographic Distribution**
```
US Users ──→ [US Replica]    (20ms latency)
EU Users ──→ [EU Replica]    (20ms latency)
Asia Users → [Asia Replica]  (20ms latency)

vs. Single server in US:
EU Users ──→ [US Server]     (150ms latency!) 😰
```

---

## Replication Strategy 1: Master-Slave (Primary-Replica)

### How It Works

```
                    WRITES
                      │
                      ↓
              ┌──────────────┐
              │   PRIMARY    │
              │   (Master)   │
              └──────────────┘
                      │
           ┌──────────┼──────────┐
           │    Replication      │
           ↓          ↓          ↓
      ┌─────────┐ ┌─────────┐ ┌─────────┐
      │ Replica │ │ Replica │ │ Replica │
      │    1    │ │    2    │ │    3    │
      └─────────┘ └─────────┘ └─────────┘
           ↑          ↑          ↑
           └──────────┼──────────┘
                   READS

All writes → Primary
Reads → Any replica (or primary)
```

### The Replication Process

```
Step 1: Client writes to Primary
┌─────────────────────────────────────┐
│ INSERT INTO users (name) VALUES     │
│ ('Alice')                           │
└─────────────────────────────────────┘
                │
                ↓
Step 2: Primary writes to its log
┌─────────────────────────────────────┐
│ Binary Log / Write-Ahead Log (WAL)  │
│ ─────────────────────────────────── │
│ Position 1001: INSERT users Alice   │
└─────────────────────────────────────┘
                │
                ↓
Step 3: Replicas pull log and apply
┌─────────────────────────────────────┐
│ Replica reads log position 1001     │
│ Applies: INSERT users Alice         │
│ Now replica has same data!          │
└─────────────────────────────────────┘
```

### Synchronous vs Asynchronous Replication

#### Asynchronous Replication (Most Common)

```
Client          Primary         Replica
  │                │               │
  │── INSERT ─────→│               │
  │                │── (later) ───→│
  │←── OK ─────────│               │
  │                │               │

Primary responds BEFORE replica confirms
Fast, but replica might lag behind!
```

**Pros:**
- ✅ Low latency (don't wait for replicas)
- ✅ Primary not affected by slow replicas
- ✅ Works across geographic regions

**Cons:**
- ❌ Replica might have stale data
- ❌ Data loss if primary fails before replication

**Replication Lag:**
```
Timeline:
─────────────────────────────────────────→
Primary:  [Write A] [Write B] [Write C]
Replica:  [Write A] [Write B] . . . waiting . . .
                              ↑
                         Replication Lag
                         (usually ms to seconds)
```

#### Synchronous Replication

```
Client          Primary         Replica
  │                │               │
  │── INSERT ─────→│               │
  │                │── Replicate ─→│
  │                │←── ACK ───────│
  │←── OK ─────────│               │
  │                │               │

Primary waits for replica confirmation
Slower, but guaranteed consistency!
```

**Pros:**
- ✅ No data loss on primary failure
- ✅ Replicas always up-to-date

**Cons:**
- ❌ Higher latency (wait for slowest replica)
- ❌ Primary blocked if replica is slow/down
- ❌ Doesn't work well across regions

#### Semi-Synchronous (Hybrid)

```
Primary waits for AT LEAST ONE replica

[Primary] ──→ [Replica 1] ✅ (fast, nearby)
          └─→ [Replica 2] ... (slow, far away)

Primary responds after Replica 1 confirms
Replica 2 catches up asynchronously

Balance of speed and safety!
```

### Handling Primary Failure (Failover)

```
Normal Operation:
[Primary] ←── Writes
    │
    └──→ [Replica 1] [Replica 2]

Primary Crashes! 💥:
[Primary] ❌
    │
    └──→ [Replica 1] [Replica 2]

Failover:
[Primary] ❌ (removed)

[Replica 1] ←── Promoted to Primary!
    │
    └──→ [Replica 2] (now replicates from Replica 1)
```

**Failover Steps:**
1. Detect primary failure (health checks)
2. Choose a replica to promote
3. Promote replica to primary
4. Reconfigure other replicas
5. Update application to use new primary

**Automatic vs Manual Failover:**
```
Manual:
- Operator decides when to failover
- Safer (no false positives)
- Slower (depends on human response)

Automatic:
- System detects and failovers automatically
- Faster recovery
- Risk of false positives (split-brain!)
```

### Reading from Replicas: Consistency Concerns

```
Timeline:
───────────────────────────────────────────────→

User writes profile update (Primary):
    "name: Alice → Alicia"
              │
              ↓
User immediately reads profile (Replica):
    "name: Alice"  ← STALE! Replication not complete!

This is "read-your-writes" inconsistency
```

**Solutions:**

**1. Read from Primary after Write**
```javascript
// After updating profile
await db.primary.update(user)

// Read from primary for 5 seconds
if (recentlyUpdated) {
    return db.primary.read(userId)
} else {
    return db.replica.read(userId)
}
```

**2. Sticky Reads**
```
User session always reads from same replica
That replica has all their recent writes
```

**3. Causal Consistency**
```
Track replication position:
- Write returns position: 1001
- Read includes position: "at least 1001"
- Replica waits until caught up to 1001
```

---

## Replication Strategy 2: Multi-Master (Active-Active)

### How It Works

```
┌──────────────┐         ┌──────────────┐
│   Master 1   │←───────→│   Master 2   │
│  (US-East)   │         │  (US-West)   │
└──────────────┘         └──────────────┘
       ↑                        ↑
       │                        │
    Writes                   Writes
    Reads                    Reads

Both can accept writes!
Both replicate to each other!
```

### Why Multi-Master?

**1. Higher Write Availability**
```
Single Master:
Master down → NO WRITES! ❌

Multi-Master:
Master 1 down → Write to Master 2 ✅
```

**2. Geographic Write Performance**
```
Single Master (in US-East):
EU User write → US-East → 150ms latency 😰

Multi-Master:
EU User write → EU Master → 20ms latency ✅
```

### The Conflict Problem

```
Same record, different masters, same time:

Master 1 (US-East):
UPDATE users SET name = 'Alice' WHERE id = 1

Master 2 (US-West):
UPDATE users SET name = 'Alicia' WHERE id = 1

Both succeed locally... but which one wins?

CONFLICT! 💥
```

### Conflict Resolution Strategies

**1. Last Write Wins (LWW)**
```
Use timestamps to pick winner:

Master 1: name = 'Alice'  @ 10:00:00.001
Master 2: name = 'Alicia' @ 10:00:00.002  ← WINS!

Simple but dangerous:
- Clock skew can cause wrong winner
- Data can be silently lost
```

**2. Application-Level Resolution**
```
Store both versions:
{
  id: 1,
  name: ['Alice', 'Alicia'],  // Conflict!
  conflict: true
}

Application or user resolves:
- Show user both options
- Merge logic specific to domain
```

**3. Conflict-Free Replicated Data Types (CRDTs)**
```
Special data structures that merge automatically:

Counter CRDT:
Master 1: counter + 5
Master 2: counter + 3
Merged:   counter + 8 (both apply!)

Set CRDT:
Master 1: add 'apple'
Master 2: add 'banana'
Merged:   {'apple', 'banana'}

No conflicts possible by design!
```

**4. Conflict Avoidance**
```
Partition writes by key:

User ID 1-1000:    → Master 1 only
User ID 1001-2000: → Master 2 only

No conflicts because different masters own different data!
(This is essentially sharding)
```

### When to Use Multi-Master

**Good for:**
- ✅ Geo-distributed writes (users in multiple regions)
- ✅ High write availability requirements
- ✅ Systems with mostly non-overlapping writes
- ✅ Offline-capable applications

**Avoid for:**
- ❌ High-conflict workloads
- ❌ Strong consistency requirements
- ❌ Complex transaction requirements
- ❌ Simple applications (overkill)

---

## Part 2: Database Sharding

### What is Sharding?

**Splitting data across multiple databases**

```
Before (Single Database):
┌─────────────────────────────────┐
│         All Users (10M)         │
│  ID 1-10,000,000 in one place!  │
└─────────────────────────────────┘
         Getting slow... 😰

After (Sharded):
┌───────────┐ ┌───────────┐ ┌───────────┐
│  Shard 1  │ │  Shard 2  │ │  Shard 3  │
│  Users    │ │  Users    │ │  Users    │
│  1-3.3M   │ │ 3.3M-6.6M │ │ 6.6M-10M  │
└───────────┘ └───────────┘ └───────────┘
Each shard handles 1/3 of queries!
```

### Why Sharding?

**1. Scale Beyond Single Machine**
```
Single server limits:
- Storage: 10TB max
- Connections: 10,000 max
- CPU: 128 cores max

Need more? → Shard!
```

**2. Scale Writes**
```
Replication only scales reads!
All writes still go to primary.

Sharding scales BOTH:
Write to Shard 1: User 1 update
Write to Shard 2: User 1M update
Write to Shard 3: User 5M update

Parallel writes to different shards!
```

**3. Reduce Query Scope**
```
Query all 10M users:
└─> Scan 10M rows 😰

Query shard with 3M users:
└─> Scan 3M rows ✅

Each shard is smaller → faster queries
```

---

## Sharding Strategy 1: Range-Based Sharding

### How It Works

```
Divide data by ranges of a key:

User ID 1-1,000,000:      → Shard 1
User ID 1,000,001-2,000,000: → Shard 2
User ID 2,000,001-3,000,000: → Shard 3
```

### Implementation

```
Shard Mapping Table:
┌─────────────────┬─────────┐
│    Range        │  Shard  │
├─────────────────┼─────────┤
│ 1 - 1,000,000   │ Shard 1 │
│ 1M - 2,000,000  │ Shard 2 │
│ 2M - 3,000,000  │ Shard 3 │
└─────────────────┴─────────┘

Query routing:
SELECT * FROM users WHERE id = 1,500,000
→ Check mapping: 1.5M is in 1M-2M range
→ Route to Shard 2
```

### Advantages

```
✅ Simple to understand and implement
✅ Range queries are efficient
   "Get users 100-200" → Single shard!
✅ Easy to add new shards
   "Add Shard 4 for users 3M-4M"
```

### Disadvantages: Hot Spots!

```
Problem: Uneven distribution

New users get sequential IDs:
- User 2,999,001 → Shard 3
- User 2,999,002 → Shard 3
- User 2,999,003 → Shard 3

ALL new users go to Shard 3! 🔥

Shard 3: 100% of new writes
Shard 1: 0% of new writes (old users)
Shard 2: 0% of new writes (old users)
```

**Solutions:**
- Use random/UUID keys (no sequential hotspots)
- Use hash-based sharding instead
- Regularly rebalance shards

---

## Sharding Strategy 2: Hash-Based Sharding

### How It Works

```
hash(key) % number_of_shards = shard

Example:
hash("user_123") = 7823456
7823456 % 3 = 1
→ User goes to Shard 1

hash("user_456") = 9823411
9823411 % 3 = 2
→ User goes to Shard 2
```

### Implementation

```python
def get_shard(user_id, num_shards):
    hash_value = hash(user_id)
    shard_num = hash_value % num_shards
    return shards[shard_num]

# Example
user_id = "user_12345"
shard = get_shard(user_id, 3)  # Returns Shard 0, 1, or 2
```

### Advantages

```
✅ Even distribution
   Hash function spreads data uniformly
   No hot spots!

✅ Any key type works
   Strings, UUIDs, composite keys
   All get hashed the same way
```

### Disadvantages

```
❌ Range queries are expensive!
   "Get users 100-200"
   → Could be on ANY shard
   → Must query ALL shards! 😰

❌ Resharding is painful!
   Adding Shard 4 changes the math:

   Before: hash % 3 = shard
   After:  hash % 4 = shard

   Most data needs to move!
```

**The Resharding Problem:**
```
Before (3 shards):
hash("user_1") % 3 = 1 → Shard 1
hash("user_2") % 3 = 2 → Shard 2
hash("user_3") % 3 = 0 → Shard 0

After (4 shards):
hash("user_1") % 4 = 1 → Shard 1 ✅ (same)
hash("user_2") % 4 = 2 → Shard 2 ✅ (same)
hash("user_3") % 4 = 3 → Shard 3 ❌ (MOVED!)

~75% of data moves when adding 1 shard!
```

---

## Sharding Strategy 3: Consistent Hashing

### The Problem with Regular Hashing

```
Add/remove shard = massive data movement

3 → 4 shards: ~75% data moves
4 → 5 shards: ~80% data moves

This is terrible for production systems!
```

### How Consistent Hashing Works

```
Imagine a ring (0 to 2^32):

                    0
                    │
           Shard 1 ●│
                   ╱│╲
                  ╱ │ ╲
         Shard 2●   │   ●Shard 3
                 ╲  │  ╱
                  ╲ │ ╱
                   ╲│╱
                    │
              2^32 (wraps to 0)

Each shard is placed on the ring
Data is placed on the ring (by hash)
Data goes to the NEXT shard clockwise
```

### Consistent Hashing in Action

```
Placing data:
hash("user_1") = position on ring
Walk clockwise → hit Shard 2
user_1 → Shard 2

hash("user_2") = different position
Walk clockwise → hit Shard 3
user_2 → Shard 3
```

### Adding a Shard (Minimal Disruption!)

```
Before:
        Shard 1
       ╱       ╲
      ╱         ╲
Shard 2─────────Shard 3

user_1 → Shard 2
user_2 → Shard 3
user_3 → Shard 1

After (add Shard 4 between 2 and 3):
        Shard 1
       ╱       ╲
      ╱         ╲
Shard 2──Shard 4──Shard 3

user_1 → Shard 4 (moved!)
user_2 → Shard 3 (same)
user_3 → Shard 1 (same)

Only ~25% of data moves (1/N where N = new shard count)
```

### Virtual Nodes

```
Problem: Uneven distribution with few shards

Shard 1 ●───────────────────────●
        ↑                       ↑
   Very far apart = too much data!

Solution: Virtual nodes
Each shard has multiple positions on ring:

Shard 1: V1.1, V1.2, V1.3, V1.4
Shard 2: V2.1, V2.2, V2.3, V2.4
Shard 3: V3.1, V3.2, V3.3, V3.4

      V1.1    V2.2    V3.1
        ●       ●       ●
       ╱                 ╲
   V3.2●                   ●V1.3
       ╲                 ╱
        ●       ●       ●
      V2.1    V1.2    V3.3

More even distribution!
Better load balancing!
```

---

## Sharding Strategy 4: Directory-Based Sharding

### How It Works

```
Lookup service tells you where data lives:

┌─────────────────────────────┐
│     Shard Directory         │
├─────────────────────────────┤
│ User 1     → Shard A        │
│ User 2     → Shard B        │
│ User 3     → Shard A        │
│ User 1000  → Shard C        │
└─────────────────────────────┘

Query: "Where is User 2?"
Directory: "Shard B"
→ Route query to Shard B
```

### Advantages

```
✅ Complete flexibility
   Any key can go anywhere

✅ Easy rebalancing
   Move data, update directory

✅ Custom placement
   "VIP users on fast shard"
   "Same-company users together"
```

### Disadvantages

```
❌ Directory is a single point of failure
   Directory down = can't route queries!

❌ Extra hop for every query
   Query directory first, then shard

❌ Directory can become bottleneck
   All queries hit directory

Solution: Cache directory heavily!
```

---

## Sharding Challenges

### Challenge 1: Cross-Shard Queries

```
Query: "Find all orders for user 123"

If orders are sharded by order_id:
- User 123's orders could be on ANY shard
- Must query ALL shards!
- Slow and expensive 😰

                    [Application]
                          │
        ┌─────────────────┼─────────────────┐
        ↓                 ↓                 ↓
   [Shard 1]         [Shard 2]         [Shard 3]
   Orders 1-1M       Orders 1M-2M      Orders 2M-3M
   (some user 123)   (some user 123)   (some user 123)
```

**Solution: Co-locate Related Data**
```
Shard by user_id instead:

Shard 1: Users 1-1M     + All their orders
Shard 2: Users 1M-2M    + All their orders
Shard 3: Users 2M-3M    + All their orders

"Find all orders for user 123" → Single shard!
```

### Challenge 2: Cross-Shard Joins

```sql
-- This is hard across shards!
SELECT u.name, o.total
FROM users u
JOIN orders o ON u.id = o.user_id
WHERE o.date > '2024-01-01'

If users and orders are on different shards:
→ Can't do a simple JOIN
→ Must fetch from both, join in application
```

**Solutions:**
1. **Denormalize:** Store user info in orders table
2. **Co-locate:** Same shard key for related tables
3. **Application-side joins:** Fetch separately, join in code
4. **Avoid joins:** Design for single-shard queries

### Challenge 3: Cross-Shard Transactions

```
Transaction: Transfer $100 from User A to User B

BEGIN TRANSACTION;
UPDATE accounts SET balance = balance - 100 WHERE user = 'A';
UPDATE accounts SET balance = balance + 100 WHERE user = 'B';
COMMIT;

If User A on Shard 1, User B on Shard 2:
→ Can't do atomic transaction across shards!
→ What if Shard 2 fails after Shard 1 commits?
→ Money disappears! 💸😱
```

**Solutions:**

**1. Two-Phase Commit (2PC)**
```
Phase 1 - Prepare:
Coordinator: "Can you commit?"
Shard 1: "Yes, prepared"
Shard 2: "Yes, prepared"

Phase 2 - Commit:
Coordinator: "Commit!"
Shard 1: "Committed"
Shard 2: "Committed"

Guarantees atomicity but:
- Slow (multiple round trips)
- Blocking (locks held during prepare)
- Coordinator failure = stuck transactions
```

**2. Saga Pattern**
```
Execute steps with compensating actions:

Step 1: Debit User A (-$100)
Step 2: Credit User B (+$100)

If Step 2 fails:
→ Compensate Step 1: Credit User A (+$100)

Eventually consistent, not ACID
But works across shards!
```

**3. Avoid Cross-Shard Transactions**
```
Design so transactions stay on one shard:

Transfer between accounts of SAME user?
→ Same shard, normal transaction!

Transfer between DIFFERENT users?
→ Use eventual consistency
→ Or use a central ledger service
```

### Challenge 4: Rebalancing

```
Over time, shards become uneven:

Shard 1: 5M users (overloaded! 🔥)
Shard 2: 2M users
Shard 3: 3M users

Need to move data from Shard 1 to others
Without downtime!
```

**Rebalancing Steps:**
```
1. Copy data from Shard 1 to Shard 2
   (Shard 1 still serving traffic)

2. Update routing to point to Shard 2
   (New writes go to Shard 2)

3. Verify Shard 2 has all data

4. Delete data from Shard 1

5. Repeat for other overloaded ranges
```

### Challenge 5: Schema Changes

```
ALTER TABLE across 100 shards?

Bad approach:
for shard in shards:
    shard.execute("ALTER TABLE...")
→ Takes hours
→ Shards out of sync during migration
→ Failures leave inconsistent state

Better approach:
1. Deploy code that handles both schemas
2. Migrate shards one by one
3. Monitor for issues
4. Remove old schema code
```

---

## Combining Replication and Sharding

```
Real production setup:

                [Application]
                      │
              [Router/Proxy]
                      │
    ┌─────────────────┼─────────────────┐
    ↓                 ↓                 ↓
[Shard 1]        [Shard 2]         [Shard 3]
    │                 │                 │
┌───┴───┐        ┌───┴───┐        ┌───┴───┐
↓       ↓        ↓       ↓        ↓       ↓
[P][R][R]        [P][R][R]        [P][R][R]

Each shard has:
- 1 Primary (writes)
- 2 Replicas (reads + failover)

Sharding for scale
Replication for availability
```

---

## Real-World Examples

### Example 1: Instagram

```
Challenge: 1 billion users, photos, likes, comments

Solution: Sharding by user_id

Shard assignment:
user_id → shard_id via consistent hashing

Co-located data per shard:
- User profile
- User's photos
- User's followers/following
- User's likes

"Show user's feed" = mostly single shard!
```

### Example 2: Uber

```
Challenge: Real-time location of millions of drivers

Solution: Geographic sharding

┌─────────────────────────────────┐
│           World Map             │
├────────┬────────┬───────────────┤
│ Shard 1│ Shard 2│    Shard 3    │
│  NYC   │   LA   │   Chicago     │
├────────┼────────┼───────────────┤
│ Shard 4│ Shard 5│    Shard 6    │
│ Miami  │ Seattle│   Denver      │
└────────┴────────┴───────────────┘

"Find drivers near me" → Single shard!
Riders and drivers in same city → Same shard
```

### Example 3: Discord

```
Challenge: Billions of messages across millions of servers

Solution: Shard by Discord server (guild)

Shard 1: Guild IDs ending in 0
Shard 2: Guild IDs ending in 1
...
Shard 9: Guild IDs ending in 9

All messages for a guild → Same shard
"Load channel history" → Single shard query!

But: Large servers (1M+ members) need special handling
→ Further partition by channel
```

---

## Decision Framework

### When to Use Replication Only

```
✅ Read-heavy workload (90%+ reads)
✅ Dataset fits on single machine
✅ Need high availability
✅ Need geographic read distribution
✅ Strong consistency requirements

Examples:
- Content management systems
- Product catalogs
- Configuration services
```

### When to Add Sharding

```
✅ Dataset too large for one machine
✅ Write-heavy workload
✅ Single primary can't handle write load
✅ Need to scale beyond single machine limits

Examples:
- Social media (billions of posts)
- E-commerce (millions of orders)
- IoT (billions of sensor readings)
```

### Comparison

| Aspect | Replication Only | Sharding |
|--------|------------------|----------|
| **Read Scale** | Excellent | Excellent |
| **Write Scale** | Limited | Excellent |
| **Data Size** | Limited | Unlimited |
| **Complexity** | Low | High |
| **Consistency** | Easy | Challenging |
| **Transactions** | Normal | Complex |
| **Query Flexibility** | Full | Limited |

---

## Common Mistakes

### Mistake 1: Sharding Too Early
```
❌ "We might have 1M users someday, let's shard now!"

Start simple:
1. Single database
2. Add read replicas
3. Optimize queries
4. Vertical scale database
5. Shard only when necessary

Sharding adds massive complexity!
```

### Mistake 2: Wrong Shard Key
```
❌ Shard by created_date
   → All new data hits same shard!

❌ Shard by random field
   → Related data scattered everywhere!

✅ Shard by access pattern
   → Data accessed together stays together
```

### Mistake 3: Ignoring Cross-Shard Operations
```
❌ "We'll figure out joins later"
   → Design forces expensive cross-shard queries

✅ Design schema around shard boundaries
   → Most queries hit single shard
```

### Mistake 4: No Rebalancing Plan
```
❌ "Shards will stay balanced"
   → 1 year later: Shard 1 = 80%, Shard 2 = 10%, Shard 3 = 10%

✅ Monitor shard sizes
   → Automate or plan for rebalancing
```

---

## Key Concepts to Remember

1. **Replication** = Same data on multiple machines (availability + read scale)
2. **Sharding** = Different data on different machines (write scale + capacity)
3. **Master-Slave** = Single writer, multiple readers
4. **Multi-Master** = Multiple writers (beware conflicts!)
5. **Range sharding** = Simple but can have hot spots
6. **Hash sharding** = Even distribution but hard to rebalance
7. **Consistent hashing** = Minimal data movement on changes
8. **Co-locate related data** on same shard to avoid cross-shard queries
9. **Cross-shard transactions** are hard - design to avoid them
10. **Start simple** - replication before sharding!

---

## Practice Questions

**Q1:** You have a database with 80% reads and 20% writes. Currently using a single PostgreSQL server that's reaching capacity. What's your first step? Justify your answer.

**Q2:** An e-commerce platform shards their orders table by order_id using hash sharding. Users complain that "View My Orders" is slow. Why? How would you fix it?

**Q3:** You're designing a chat application like Slack. Messages need to be queried by:
- Channel (most common)
- User (less common)
- Date range (analytics)

What's your sharding strategy? What trade-offs are you making?

**Q4:** Your multi-master MySQL setup has conflicts occurring 100 times per day. 90% are on the "page_views" counter column. How do you eliminate these conflicts?

**Q5:** Explain why consistent hashing is better than regular hash sharding when you need to add or remove shards frequently.

**Q6:** Design the replication and sharding strategy for a social media app with:
- 100M users
- Users post ~2 times per day
- Users read feeds ~50 times per day
- Users are globally distributed

**Q7:** A team wants to shard by user_id but needs to run this query efficiently:
```sql
SELECT COUNT(*) FROM orders
WHERE created_at > '2024-01-01'
GROUP BY product_id
```
What's the problem? Propose solutions.

---

## Next Up

In Lesson 9, we'll explore **Stateless vs Stateful Architecture** - the key to making your applications truly scalable!
