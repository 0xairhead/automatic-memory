# Week 4, Lesson 10: Consistent Hashing

## Table of Contents
- [The Distribution Problem](#the-distribution-problem)
- [Traditional Hashing: The Naive Approach](#traditional-hashing-the-naive-approach)
  - [How It Works](#how-it-works)
  - [The Problem: Adding or Removing Servers](#the-problem-adding-or-removing-servers)
  - [Why This Is Catastrophic](#why-this-is-catastrophic)
- [Consistent Hashing: The Solution](#consistent-hashing-the-solution)
  - [The Hash Ring Concept](#the-hash-ring-concept)
  - [Placing Servers on the Ring](#placing-servers-on-the-ring)
  - [Placing Data on the Ring](#placing-data-on-the-ring)
  - [Finding the Right Server](#finding-the-right-server)
- [Adding and Removing Servers](#adding-and-removing-servers)
  - [Adding a Server](#adding-a-server)
  - [Removing a Server](#removing-a-server)
  - [The Math: How Much Data Moves?](#the-math-how-much-data-moves)
- [The Uneven Distribution Problem](#the-uneven-distribution-problem)
- [Virtual Nodes: The Complete Solution](#virtual-nodes-the-complete-solution)
  - [What Are Virtual Nodes?](#what-are-virtual-nodes)
  - [How Virtual Nodes Work](#how-virtual-nodes-work)
  - [Benefits of Virtual Nodes](#benefits-of-virtual-nodes)
  - [Choosing the Number of Virtual Nodes](#choosing-the-number-of-virtual-nodes)
- [Implementing Consistent Hashing](#implementing-consistent-hashing)
  - [Basic Implementation](#basic-implementation)
  - [Production Considerations](#production-considerations)
- [Real-World Applications](#real-world-applications)
  - [Distributed Caches (Memcached, Redis Cluster)](#distributed-caches-memcached-redis-cluster)
  - [Content Delivery Networks (CDNs)](#content-delivery-networks-cdns)
  - [Distributed Databases (Cassandra, DynamoDB)](#distributed-databases-cassandra-dynamodb)
  - [Load Balancers](#load-balancers)
- [Consistent Hashing Variants](#consistent-hashing-variants)
  - [Jump Consistent Hashing](#jump-consistent-hashing)
  - [Rendezvous Hashing (HRW)](#rendezvous-hashing-hrw)
  - [Maglev Hashing](#maglev-hashing)
- [Common Mistakes](#common-mistakes)
- [Key Concepts to Remember](#key-concepts-to-remember)
- [Practice Questions](#practice-questions)

---

Welcome to Week 4! We're now diving into distribution fundamentals. This lesson covers one of the most elegant algorithms in distributed systems: **consistent hashing**. It's the secret behind how systems like Cassandra, DynamoDB, and Memcached distribute data efficiently.

---

## The Distribution Problem

When you have multiple servers, you need to decide: **which server handles which data?**

```
You have 3 cache servers:
[Server 0] [Server 1] [Server 2]

User requests data for key "user_123"
Which server should store/retrieve it?

You need:
1. Deterministic: Same key → same server (always!)
2. Balanced: Data spread evenly across servers
3. Minimal disruption: Adding/removing servers shouldn't reshuffle everything
```

---

## Traditional Hashing: The Naive Approach

### How It Works

The simplest approach: **hash the key, mod by server count**

```
server = hash(key) % number_of_servers

Example with 3 servers:
hash("user_1") = 7    → 7 % 3 = 1    → Server 1
hash("user_2") = 15   → 15 % 3 = 0   → Server 0
hash("user_3") = 23   → 23 % 3 = 2   → Server 2
hash("user_4") = 42   → 42 % 3 = 0   → Server 0
hash("user_5") = 58   → 58 % 3 = 1   → Server 1
```

```
Distribution:
[Server 0]: user_2, user_4
[Server 1]: user_1, user_5
[Server 2]: user_3

Looks good! Evenly distributed!
```

### The Problem: Adding or Removing Servers

**What happens when we add a 4th server?**

```
BEFORE (3 servers):
hash("user_1") = 7    → 7 % 3 = 1    → Server 1
hash("user_2") = 15   → 15 % 3 = 0   → Server 0
hash("user_3") = 23   → 23 % 3 = 2   → Server 2
hash("user_4") = 42   → 42 % 3 = 0   → Server 0
hash("user_5") = 58   → 58 % 3 = 1   → Server 1

AFTER (4 servers):
hash("user_1") = 7    → 7 % 4 = 3    → Server 3  ← MOVED!
hash("user_2") = 15   → 15 % 4 = 3   → Server 3  ← MOVED!
hash("user_3") = 23   → 23 % 4 = 3   → Server 3  ← MOVED!
hash("user_4") = 42   → 42 % 4 = 2   → Server 2  ← MOVED!
hash("user_5") = 58   → 58 % 4 = 2   → Server 2  ← MOVED!

ALL 5 keys moved to different servers! 😱
```

### Why This Is Catastrophic

```
Scenario: Cache cluster with 100 servers, 1 billion keys

One server fails (99 servers now):
- hash(key) % 100 → hash(key) % 99
- ~99% of keys now map to DIFFERENT servers!
- 990 million cache misses!
- Database gets hammered with 990 million requests!
- Potential cascading failure! 💥

This is called a "cache stampede" or "thundering herd"
```

**The Math:**

```
When changing from N to N±1 servers:
Keys that move ≈ (N-1)/N of all keys

3 → 4 servers: ~75% of keys move
10 → 11 servers: ~91% of keys move
100 → 101 servers: ~99% of keys move

The more servers, the WORSE it gets!
```

---

## Consistent Hashing: The Solution

**Key Insight:** Instead of `hash % N`, map both servers AND keys to a fixed circular space (a "ring").

### The Hash Ring Concept

```
Imagine a circle (ring) with positions 0 to 2^32-1:

                        0
                        │
                    ────┼────
                  ╱     │     ╲
                ╱       │       ╲
              ╱         │         ╲
            ╱           │           ╲
   2^32 * 3/4           │            2^32 * 1/4
           │            │            │
           │            │            │
            ╲           │           ╱
              ╲         │         ╱
                ╲       │       ╱
                  ╲     │     ╱
                    ────┼────
                        │
                    2^32 / 2

The ring wraps around: position 2^32 = position 0
```

### Placing Servers on the Ring

**Hash the server identifier to get its position on the ring**

```
hash("server_A") = 10,000      → Position 10,000
hash("server_B") = 1,000,000   → Position 1,000,000
hash("server_C") = 3,500,000   → Position 3,500,000

On the ring:
                    0
                    ●
                   ╱│╲
                  ╱ │ ╲
       Server_A ●  │  ╲
                │   │   ╲
                │   │    ● Server_B
                │   │   ╱
                │   │  ╱
                 ╲  │ ╱
                  ╲ │╱
                   ╲●
               Server_C
```

### Placing Data on the Ring

**Hash the key to get its position, then find the next server clockwise**

```
hash("user_1") = 500,000       → Position 500,000
hash("user_2") = 2,000,000     → Position 2,000,000
hash("user_3") = 9,000         → Position 9,000

                    0
                    │
                   ╱│╲
  user_3 (9K) ────x │ ╲
       Server_A ●  │  ╲
  user_1 (500K) ──x│   ╲
                │   │    ● Server_B
                │   │   ╱
  user_2 (2M) ────x│  ╱
                 ╲  │ ╱
                  ╲ │╱
                   ╲●
               Server_C
```

### Finding the Right Server

**Walk clockwise from the key's position to find the first server**

```
user_3 (position 9,000):
→ Walk clockwise
→ First server: Server_A (position 10,000)
→ user_3 → Server_A

user_1 (position 500,000):
→ Walk clockwise
→ First server: Server_B (position 1,000,000)
→ user_1 → Server_B

user_2 (position 2,000,000):
→ Walk clockwise
→ First server: Server_C (position 3,500,000)
→ user_2 → Server_C
```

```
Final assignment:
Server_A: user_3
Server_B: user_1
Server_C: user_2
```

---

## Adding and Removing Servers

### Adding a Server

```
Add Server_D at position 400,000:

BEFORE:
                    0
       Server_A ●  (10K)
                │
    user_1 (500K)─x
                │
          Server_B ● (1M)

AFTER:
                    0
       Server_A ●  (10K)
                │
       Server_D ●  (400K)  ← NEW!
                │
    user_1 (500K)─x
                │
          Server_B ● (1M)

user_1 now maps to Server_D (not Server_B)!
But user_3 still maps to Server_A (unchanged!)
```

**Only keys between Server_D and the previous server (Server_A) are affected!**

### Removing a Server

```
Remove Server_B:

BEFORE:
       Server_A ●  (10K)
                │
    user_1 (500K)─x
                │
          Server_B ● (1M)
                │
    user_2 (2M) ─x
                │
          Server_C ● (3.5M)

AFTER:
       Server_A ●  (10K)
                │
    user_1 (500K)─x───────┐
                │         │
          (Server_B gone) │
                │         │
    user_2 (2M) ─x        │
                │         ↓
          Server_C ● (3.5M)

user_1 now maps to Server_C
user_2 still maps to Server_C (unchanged!)
```

**Only keys that were on the removed server need to move!**

### The Math: How Much Data Moves?

```
Traditional hashing:
Adding 1 server to N servers → ~(N-1)/N keys move
3 → 4 servers: 75% move
100 → 101 servers: 99% move

Consistent hashing:
Adding 1 server to N servers → ~1/N keys move
3 → 4 servers: 25% move
100 → 101 servers: 1% move

That's a MASSIVE improvement!
```

```
Example with 100 servers, 1 billion keys:

Traditional: 990 million keys move (99%)
Consistent:  10 million keys move (1%)

99x less disruption!
```

---

## The Uneven Distribution Problem

**With few servers, distribution can be very uneven**

```
3 servers placed "randomly" on the ring:

                    0
                    │
                    │
       Server_A ●───┤ (position 100)
                    │
                    │
                    │
       Server_B ●───┤ (position 500)
                    │
                    │
                    │
                    │
                    │
                    │
       Server_C ●───┤ (position 3,000,000,000)
                    │
                    │

Server_A handles: 100 to 500 (range: 400)
Server_B handles: 500 to 3B (range: ~3 billion!)
Server_C handles: 3B to 100 (range: ~1 billion)

Server_B has 75% of the key space!
Terrible distribution! 😰
```

**As you add more servers, it gets better, but with just a few servers, the variance is high.**

---

## Virtual Nodes: The Complete Solution

### What Are Virtual Nodes?

**Instead of 1 position per server, give each server multiple positions**

```
Without virtual nodes:
Server_A → 1 position on ring
Server_B → 1 position on ring
Server_C → 1 position on ring

With virtual nodes (100 per server):
Server_A → 100 positions on ring (A-0, A-1, ... A-99)
Server_B → 100 positions on ring (B-0, B-1, ... B-99)
Server_C → 100 positions on ring (C-0, C-1, ... C-99)
```

### How Virtual Nodes Work

```
Create virtual nodes by hashing server + index:

hash("Server_A_0") = position 1,000
hash("Server_A_1") = position 500,000
hash("Server_A_2") = position 2,100,000
hash("Server_A_3") = position 3,800,000
...

hash("Server_B_0") = position 250,000
hash("Server_B_1") = position 1,500,000
hash("Server_B_2") = position 2,800,000
...
```

```
The ring now looks like:

                    0
                    │
           A-0  ●───┤
                    │
           B-0  ●───┤
                    │
           A-1  ●───┤
                    │
           B-1  ●───┤
                    │
           A-2  ●───┤
                    │
           C-0  ●───┤
                    │
           B-2  ●───┤
                    │
           ...      │

Servers are spread across the ring!
Much better distribution!
```

### Benefits of Virtual Nodes

**1. Even Distribution**
```
With enough virtual nodes, each server handles roughly 1/N of the keys

Mathematical property:
- With V virtual nodes per server
- Standard deviation of load ≈ 1/√V
- 100 virtual nodes → ~10% variance
- 1000 virtual nodes → ~3% variance
```

**2. Heterogeneous Servers**
```
Powerful server? Give it more virtual nodes!

Server_A (8 cores):   100 virtual nodes
Server_B (8 cores):   100 virtual nodes
Server_C (16 cores):  200 virtual nodes  ← Handles 2x the load!

Natural load balancing based on capacity!
```

**3. Smoother Rebalancing**
```
Adding a server with 100 virtual nodes:
- Affects 100 small segments of the ring
- Each existing server loses a little data
- More gradual than taking one big chunk

Removing a server:
- Its 100 virtual nodes disappear
- Load distributed across all remaining servers
- No single server gets overwhelmed
```

### Choosing the Number of Virtual Nodes

```
Trade-offs:

Few virtual nodes (10-50):
+ Less memory for ring structure
+ Faster lookups
- More uneven distribution
- More variance when adding/removing servers

Many virtual nodes (100-500):
+ Very even distribution
+ Smooth rebalancing
- More memory
- Slightly slower lookups (larger ring)

Recommendation:
- Start with 100-150 virtual nodes per physical server
- Adjust based on your needs
- Most systems use 100-256
```

---

## Implementing Consistent Hashing

### Basic Implementation

```python
import hashlib
import bisect

class ConsistentHash:
    def __init__(self, nodes=None, virtual_nodes=100):
        self.virtual_nodes = virtual_nodes
        self.ring = {}           # hash -> node
        self.sorted_keys = []    # sorted hash values

        if nodes:
            for node in nodes:
                self.add_node(node)

    def _hash(self, key):
        """Generate hash value for a key"""
        return int(hashlib.md5(key.encode()).hexdigest(), 16)

    def add_node(self, node):
        """Add a node with virtual nodes to the ring"""
        for i in range(self.virtual_nodes):
            virtual_key = f"{node}_vn{i}"
            hash_value = self._hash(virtual_key)
            self.ring[hash_value] = node
            bisect.insort(self.sorted_keys, hash_value)

    def remove_node(self, node):
        """Remove a node and its virtual nodes from the ring"""
        for i in range(self.virtual_nodes):
            virtual_key = f"{node}_vn{i}"
            hash_value = self._hash(virtual_key)
            if hash_value in self.ring:
                del self.ring[hash_value]
                self.sorted_keys.remove(hash_value)

    def get_node(self, key):
        """Find the node responsible for a key"""
        if not self.ring:
            return None

        hash_value = self._hash(key)

        # Find first node with hash >= key's hash
        idx = bisect.bisect_left(self.sorted_keys, hash_value)

        # Wrap around to first node if past the end
        if idx == len(self.sorted_keys):
            idx = 0

        return self.ring[self.sorted_keys[idx]]


# Usage
ch = ConsistentHash(["server_A", "server_B", "server_C"])

print(ch.get_node("user_1"))    # server_B
print(ch.get_node("user_2"))    # server_A
print(ch.get_node("user_3"))    # server_C

# Add a server - minimal redistribution
ch.add_node("server_D")

# Remove a server - minimal redistribution
ch.remove_node("server_A")
```

### Production Considerations

**1. Replication**
```
Get not just one node, but N nodes for replication:

def get_nodes(self, key, count=3):
    """Get multiple nodes for replication"""
    if not self.ring:
        return []

    hash_value = self._hash(key)
    idx = bisect.bisect_left(self.sorted_keys, hash_value)

    nodes = []
    seen = set()

    while len(nodes) < count and len(seen) < len(set(self.ring.values())):
        if idx >= len(self.sorted_keys):
            idx = 0

        node = self.ring[self.sorted_keys[idx]]
        if node not in seen:
            nodes.append(node)
            seen.add(node)
        idx += 1

    return nodes

# Get 3 nodes for "user_1" (for 3x replication)
nodes = ch.get_nodes("user_1", 3)
# Returns: ["server_B", "server_C", "server_A"]
```

**2. Weighted Nodes**
```python
def add_node(self, node, weight=1):
    """Add node with weight (more virtual nodes)"""
    num_virtual = self.virtual_nodes * weight
    for i in range(num_virtual):
        virtual_key = f"{node}_vn{i}"
        hash_value = self._hash(virtual_key)
        self.ring[hash_value] = node
        bisect.insort(self.sorted_keys, hash_value)

# Powerful server gets 2x the load
ch.add_node("big_server", weight=2)
ch.add_node("small_server", weight=1)
```

**3. Prefer Consistent Hash Libraries**
```
Production systems should use battle-tested libraries:

Python: uhashring, hash_ring
Java: ketama
Go: consistent (by stathat)
Node.js: hashring

These handle edge cases and are optimized for performance.
```

---

## Real-World Applications

### Distributed Caches (Memcached, Redis Cluster)

```
Memcached cluster with consistent hashing:

Client wants to cache "user_123":
1. hash("user_123") → find server on ring
2. Store/retrieve from that specific server
3. If server fails, only its keys need re-caching

┌────────────┐
│   Client   │
└─────┬──────┘
      │ hash("user_123") → Server 2
      ↓
[Server 1] [Server 2] [Server 3]
              ↑
         Cache here!

Benefits:
- Cache hits stay consistent during scaling
- Server failure only affects 1/N of cache
- Easy to add capacity without full cache rebuild
```

### Content Delivery Networks (CDNs)

```
CDN edge server selection:

User requests: video_xyz.mp4
1. hash("video_xyz.mp4") → find edge server
2. If edge server has it, serve directly
3. If not, fetch from origin, cache, then serve

Why consistent hashing?
- Same content → same edge server
- Better cache hit ratio
- When edge servers fail, minimal recaching

    User in NYC
         │
         ↓
    [CDN Router]
         │
         │ hash("video_xyz") → Edge-East
         ↓
[Edge-West] [Edge-East] [Edge-Central]
                ↑
           Serve from here
```

### Distributed Databases (Cassandra, DynamoDB)

```
Cassandra's token ring:

Each node owns a range of the hash ring:

Node A: tokens 0 - 1000
Node B: tokens 1001 - 2000
Node C: tokens 2001 - 3000

INSERT INTO users (id, name) VALUES ('user_123', 'Alice')
1. hash("user_123") = 1500
2. 1500 falls in Node B's range
3. Write to Node B (and replicas)

Node B fails:
- Node A and C take over Node B's range
- Only 1/3 of data needs rebalancing
```

### Load Balancers

```
Session-sticky load balancing:

hash(client_ip) → backend server

Client 1.2.3.4 → always goes to Backend_A
Client 5.6.7.8 → always goes to Backend_B

Benefits:
- Session affinity without cookies
- If Backend_A fails, only its clients move
- New backend gets proportional share of traffic

     [Load Balancer]
      │    │    │
      │    │    │  hash(client_ip)
      ↓    ↓    ↓
   [BE_A][BE_B][BE_C]
```

---

## Consistent Hashing Variants

### Jump Consistent Hashing

```
Google's algorithm (2014):
- No ring, no virtual nodes
- O(1) memory
- Very fast: O(log n) time
- Perfect distribution

int32_t JumpConsistentHash(uint64_t key, int32_t num_buckets) {
    int64_t b = -1, j = 0;
    while (j < num_buckets) {
        b = j;
        key = key * 2862933555777941757ULL + 1;
        j = (b + 1) * (double(1LL << 31) / double((key >> 33) + 1));
    }
    return b;
}

Limitations:
- Buckets must be numbered 0 to N-1
- Can only add/remove from the end
- No weighted nodes
```

### Rendezvous Hashing (HRW)

```
Highest Random Weight hashing:

For each key, compute score for all servers:
score(key, server) = hash(key + server)

Pick server with highest score!

get_server("user_123"):
  score("user_123", "A") = 50000
  score("user_123", "B") = 80000  ← highest
  score("user_123", "C") = 30000

  → Server B

Benefits:
- Simple concept
- Good distribution
- Easy to add weights

Drawback:
- O(n) time to find server (must check all)
```

### Maglev Hashing

```
Google's load balancer hashing (2016):

- Builds a lookup table for O(1) lookups
- Minimal disruption when backends change
- Used in Google's network load balancer

More complex but very efficient for
high-throughput load balancing.
```

---

## Common Mistakes

### Mistake 1: Using Modulo Hashing for Distributed Systems
```
❌ server = hash(key) % num_servers

This causes massive reshuffling when servers change!

✅ Use consistent hashing from the start
   Even with 3 servers, you'll thank yourself later
```

### Mistake 2: Too Few Virtual Nodes
```
❌ 1 virtual node per server
   → Uneven distribution
   → One server might get 50% of traffic!

✅ Use 100+ virtual nodes per server
   → Even distribution
   → ~1% variance
```

### Mistake 3: Ignoring Hotspots
```
❌ Assuming consistent hashing solves all distribution problems

Hot key "celebrity_post" gets 1M requests:
→ All 1M go to ONE server (the one on the ring)
→ That server is overwhelmed!

✅ Solutions for hotspots:
   - Local caching at the application layer
   - Key replication (store hot keys on multiple servers)
   - Key splitting ("celebrity_post_1", "celebrity_post_2", ...)
```

### Mistake 4: Not Handling Server Health
```
❌ Routing to dead servers

Consistent hashing tells you WHERE, not IF server is alive

✅ Combine with health checks:
   - Remove failed servers from ring
   - Route to next healthy server
   - Re-add servers when recovered
```

### Mistake 5: Hash Function Collisions
```
❌ Using weak hash functions

If two keys hash to same value:
→ Always go to same server
→ If many collisions, uneven distribution

✅ Use strong hash functions:
   - MD5 (good for distribution, not security)
   - MurmurHash3
   - xxHash

Don't use: simple modulo, CRC32
```

---

## Key Concepts to Remember

1. **Traditional hash(key) % N** fails when N changes - massive data movement
2. **Consistent hashing** uses a ring where both servers and keys are mapped
3. **Walk clockwise** from key position to find the responsible server
4. **Adding/removing servers** only affects 1/N of keys (not N-1/N)
5. **Virtual nodes** ensure even distribution and smooth rebalancing
6. **100-150 virtual nodes** per server is a good starting point
7. **Replication** = get multiple nodes walking clockwise
8. **Weighted nodes** = more virtual nodes for bigger servers
9. **Consistent hashing doesn't solve hotspots** - hot keys still overwhelm one server
10. **Used everywhere**: Cassandra, DynamoDB, Memcached, CDNs, load balancers

---

## Practice Questions

**Q1:** A cache cluster uses `hash(key) % 10` for 10 servers. One server fails. What percentage of cache misses would you expect? What if they used consistent hashing?

**Q2:** You have 4 servers with different capacities:
- Server A: 8GB RAM
- Server B: 8GB RAM
- Server C: 16GB RAM
- Server D: 32GB RAM

How would you configure virtual nodes to distribute data proportionally to capacity?

**Q3:** Explain why consistent hashing with 1 virtual node per server (no virtual nodes) might result in one server handling 70% of the requests while another handles only 5%.

**Q4:** A social media platform has a "viral tweet" problem. When a celebrity posts, millions of requests try to fetch that tweet. The tweet is on Server 3 per consistent hashing. How would you handle this hotspot?

**Q5:** Compare these approaches for a 100-server cluster:
- Traditional modulo hashing
- Consistent hashing with 10 virtual nodes each
- Consistent hashing with 200 virtual nodes each

Consider: distribution evenness, memory usage, lookup time, and data movement when adding 1 server.

**Q6:** Implement a `get_nodes(key, n)` method that returns n different physical servers for replication, handling the case where virtual nodes of the same physical server are adjacent on the ring.

**Q7:** You're designing a distributed key-value store. Should you use consistent hashing or jump consistent hashing? What factors would influence your decision?

---

## Next Up

In Lesson 11, we'll explore **Message Queues & Async Processing** - essential patterns for building resilient, scalable systems!
