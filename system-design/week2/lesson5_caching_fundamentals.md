# Week 2, Lesson 5: Caching Fundamentals - Why, Where, and How

## Table of Contents
- [Media Resources](#media-resources)
- [What is Caching?](#what-is-caching)
- [Why Cache? The Performance Impact](#why-cache-the-performance-impact)
- [Where to Cache? (Cache Layers)](#where-to-cache-cache-layers)
- [Cache Patterns (How to Use Cache)](#cache-patterns-how-to-use-cache)
  - [Pattern 1: Cache-Aside (Lazy Loading)](#pattern-1-cache-aside-lazy-loading)
  - [Pattern 2: Read-Through Cache](#pattern-2-read-through-cache)
  - [Pattern 3: Write-Through Cache](#pattern-3-write-through-cache)
  - [Pattern 4: Write-Behind (Write-Back) Cache](#pattern-4-write-behind-write-back-cache)
  - [Pattern 5: Write-Around Cache](#pattern-5-write-around-cache)
- [Cache Eviction Policies (When Cache is Full)](#cache-eviction-policies-when-cache-is-full)
- [Cache Invalidation: The Hard Problem](#cache-invalidation-the-hard-problem)
- [Distributed Caching](#distributed-caching)
- [Popular Caching Solutions](#popular-caching-solutions)
  - [Redis](#redis)
  - [Memcached](#memcached)
  - [Redis vs Memcached](#redis-vs-memcached)
- [Real-World Examples](#real-world-examples)
- [Cache Stampede Problem](#cache-stampede-problem)
- [Cache Monitoring & Metrics](#cache-monitoring--metrics)
- [Best Practices](#best-practices)
- [When NOT to Cache](#when-not-to-cache)
- [Key Concepts to Remember](#key-concepts-to-remember)
- [Practice Questions](#practice-questions)
- [Next Up](#next-up)

---

Welcome to one of the most important lessons in system design! Caching is the secret weapon behind fast, scalable systems. Master this, and you'll dramatically improve any system's performance.

## Media Resources

**Visual Guide:**
![Caching Fundamentals: A Guide to Faster Systems](./assets/caching_fundamentals_infographic.png)

**Audio Lesson:**
[The Secret Weapon of System Design: Caching (Audio)](./assets/caching_fundamentals.m4a)

---

## What is Caching?

**Cache = Temporary storage of frequently accessed data in a faster location**

Think of it like this:

```
Without Cache (Every time):
You need milk → Drive to store → Buy milk → Drive home → 30 minutes 😢

With Cache (Smart way):
You need milk → Check fridge (cache) → Milk is there! → 10 seconds 😊
If no milk in fridge → Drive to store → Buy extra → Store in fridge
```

**In computing:**
- Database query takes 100ms
- Cache lookup takes 1ms
- That's **100x faster!**

---

## Why Cache? The Performance Impact

### Real Numbers

```
Operation                          Latency
─────────────────────────────────────────────
L1 cache reference                 0.5 ns
L2 cache reference                 7 ns
RAM access                         100 ns
Redis/Memcached (network)          500 μs (0.5 ms)
SSD read                           16 μs
Database query (indexed)           10 ms
Database query (full scan)         100+ ms
```

**Key insight:** Caching can make your system **10-1000x faster!**

### The 80/20 Rule in Caching

**80% of requests access 20% of data**

Examples:
- Netflix: 80% of views are on 20% of content
- E-commerce: 80% of sales from 20% of products
- Social media: 80% of traffic to 20% of posts

**Solution: Cache that hot 20%!**

---

## Where to Cache? (Cache Layers)

### Layer 1: Browser Cache

```
User requests: /style.css

Browser checks:
└─> "Do I have style.css cached?"
    ├─> YES: Use cached version (0ms network!)
    └─> NO: Request from server
```

**Controlled by HTTP headers:**
```http
Cache-Control: public, max-age=86400
Expires: Thu, 27 Dec 2024 10:30:00 GMT
ETag: "abc123"
```

**Best for:** Static assets (CSS, JS, images)

### Layer 2: CDN Cache

```
Request: /logo.png

CDN Edge Server:
└─> "Do I have logo.png?"
    ├─> YES: Return (10-50ms)
    └─> NO: Fetch from origin, cache, return (100-200ms)
```

**Best for:** Static content served globally

### Layer 3: Application Cache (Redis/Memcached)

```
User requests: "Get user profile ID 123"

App checks Redis:
└─> GET user:123
    ├─> Cache HIT: Return data (1ms)
    └─> Cache MISS: Query database (100ms), store in Redis
```

**Best for:** Session data, frequently accessed data, computed results

### Layer 4: Database Cache

```
Query: SELECT * FROM users WHERE id = 123;

Database checks query cache:
└─> "Have I executed this exact query recently?"
    ├─> YES: Return cached result
    └─> NO: Execute query, cache result
```

**Best for:** Repeated identical queries

### Complete Cache Hierarchy

```
User Request
    ↓
[Browser Cache]
    ↓ (miss)
[CDN Cache]
    ↓ (miss)
[Load Balancer]
    ↓
[Application Server]
    ↓
[Redis/Memcached Cache]
    ↓ (miss)
[Database Query Cache]
    ↓ (miss)
[Database Disk]

Each layer = potential performance boost!
```

---

## Cache Patterns (How to Use Cache)

### Pattern 1: Cache-Aside (Lazy Loading)

**Most common pattern!**

```
1. Application requests data from cache
2. Cache MISS → Application queries database
3. Application stores data in cache
4. Next request → Cache HIT!

Code example:
───────────────
user = cache.get("user:123")

if user is None:  # Cache MISS
    user = database.query("SELECT * FROM users WHERE id=123")
    cache.set("user:123", user, ttl=3600)  # Cache for 1 hour

return user
```

**Pros:**
- ✅ Only cache what's needed
- ✅ Resilient (cache failure = slower, not broken)
- ✅ Simple to implement

**Cons:**
- ❌ First request is always slow (cold start)
- ❌ Cache miss penalty (3 round trips: cache → db → cache)

**Best for:** Read-heavy applications

### Pattern 2: Read-Through Cache

```
1. Application requests data from cache
2. Cache automatically fetches from database if miss
3. Cache handles all database interaction

Code example:
───────────────
user = cache.get("user:123")
# Cache automatically queries DB if needed
```

**Pros:**
- ✅ Application code simpler
- ✅ Automatic cache population

**Cons:**
- ❌ Cache has database dependency
- ❌ First request still slow

**Best for:** Read-heavy with cache abstraction layer

### Pattern 3: Write-Through Cache

```
1. Application writes to cache
2. Cache writes to database synchronously
3. Cache and database always in sync

Write flow:
───────────
cache.set("user:123", user)
└─> Cache updates itself
└─> Cache writes to database
└─> Returns success
```

**Pros:**
- ✅ Data consistency (cache = database)
- ✅ Never lose writes

**Cons:**
- ❌ Write latency (must wait for database)
- ❌ Writes lots of data that might never be read

**Best for:** Systems where consistency is critical

### Pattern 4: Write-Behind (Write-Back) Cache

```
1. Application writes to cache
2. Cache returns success immediately
3. Cache writes to database asynchronously (later)

Write flow:
───────────
cache.set("user:123", user)
└─> Cache updates itself
└─> Returns success immediately ✅
└─> Later: Cache syncs to database
```

**Pros:**
- ✅ Fast writes
- ✅ Can batch multiple writes

**Cons:**
- ❌ Risk of data loss (if cache fails before sync)
- ❌ Inconsistency window

**Best for:** Write-heavy applications (logging, analytics)

### Pattern 5: Write-Around Cache

```
1. Application writes directly to database
2. Cache is NOT updated
3. On next read, cache miss → fetch from database → cache it

Write flow:
───────────
database.write("user:123", user)
# Cache is not updated!

Read flow:
───────────
user = cache.get("user:123")  # MISS (stale data)
user = database.query(...)
cache.set("user:123", user)
```

**Pros:**
- ✅ Avoids cache pollution (don't cache data that won't be read)

**Cons:**
- ❌ Read after write = cache miss

**Best for:** Infrequently read data

---

## Cache Eviction Policies (When Cache is Full)

Cache has limited space. What do you remove when full?

### 1. LRU (Least Recently Used)

**Most popular!**

```
Cache (size=3):
[A, B, C]

Access B:
[A, C, B]  (B moved to front)

Access new item D:
[C, B, D]  (A removed - least recently used)
```

**Implementation:** Doubly linked list + hash map

**Best for:** Most applications (good default)

### 2. LFU (Least Frequently Used)

```
Cache:
Item A: accessed 5 times
Item B: accessed 10 times
Item C: accessed 2 times

New item arrives:
Remove C (least frequently used)
```

**Best for:** Items with consistent popularity

### 3. FIFO (First In, First Out)

```
Cache:
[A, B, C]  (A entered first)

New item D:
[B, C, D]  (A removed)
```

**Simple but not smart!**

### 4. Random Replacement

```
Cache full?
Remove random item.
```

**Surprisingly effective for large caches!**

### 5. TTL (Time To Live)

```
cache.set("user:123", user, ttl=3600)  # 1 hour

After 1 hour:
Entry automatically expires and is removed
```

**Best practice: Combine LRU + TTL**

---

## Cache Invalidation: The Hard Problem

> "There are only two hard things in Computer Science: cache invalidation and naming things."  
> — Phil Karlton

### The Problem

```
Database: User name = "Alice"
Cache: User name = "Alice"

User updates name to "Alice Smith"

Database: User name = "Alice Smith" ✅
Cache: User name = "Alice" ❌ STALE!
```

### Solution 1: TTL (Time To Live)

```
cache.set("user:123", user, ttl=300)  # 5 minutes

After 5 minutes:
Cache entry expires
Next read fetches fresh data from database
```

**Pros:** Simple, automatic
**Cons:** Data can be stale for up to TTL duration

### Solution 2: Explicit Invalidation

```
def update_user(user_id, new_data):
    database.update(user_id, new_data)
    cache.delete(f"user:{user_id}")  # Invalidate cache
```

**Pros:** Immediate consistency
**Cons:** Must remember to invalidate everywhere

### Solution 3: Write-Through

```
def update_user(user_id, new_data):
    cache.set(f"user:{user_id}", new_data)
    # Cache automatically updates database
```

**Pros:** Cache always consistent
**Cons:** Slower writes

### Solution 4: Event-Based Invalidation

```
def update_user(user_id, new_data):
    database.update(user_id, new_data)
    event_bus.publish("user.updated", user_id)

# Separate service listens:
def on_user_updated(user_id):
    cache.delete(f"user:{user_id}")
```

**Pros:** Decoupled, scalable
**Cons:** Complex architecture

---

## Distributed Caching

When one cache server isn't enough:

### Problem: Multiple Cache Servers

```
Cache Server 1: user:123 → "Alice"
Cache Server 2: user:123 → "Bob"

Which one is correct? 😱
```

### Solution: Consistent Hashing

```
hash(user:123) % num_servers = server_index

user:123 always goes to Server 2
user:456 always goes to Server 1

Same key = same server!
```

**More on this in distributed systems lesson!**

### Cache Replication

```
Master Cache ──┐
               ├─> Replica 1
               ├─> Replica 2
               └─> Replica 3

If master fails, promote replica!
```

---

## Popular Caching Solutions

### Redis

**In-memory data store with persistence**

```
Features:
- Multiple data types (strings, lists, sets, sorted sets, hashes)
- Pub/Sub messaging
- Transactions
- Lua scripting
- Persistence (AOF, RDB)
- Replication

Use cases:
- Session storage
- Real-time analytics
- Leaderboards
- Rate limiting
- Message queues
```

**Example:**
```redis
SET user:123 '{"name":"Alice","age":30}'
GET user:123

EXPIRE user:123 3600  # TTL 1 hour

INCR page:views
ZADD leaderboard 100 "player1"
```

### Memcached

**Simple, fast, in-memory key-value store**

```
Features:
- Simple key-value only
- No persistence
- Multi-threaded
- Simpler than Redis

Use cases:
- Simple caching
- Session storage
- Database query results
```

**Example:**
```
SET user:123 '{"name":"Alice"}' 3600
GET user:123
DELETE user:123
```

### Redis vs Memcached

| Feature | Redis | Memcached |
|---------|-------|-----------|
| Data types | Many (list, set, hash, etc.) | Key-value only |
| Persistence | Yes | No |
| Replication | Yes | No |
| Performance | Slightly slower | Slightly faster |
| Use case | Complex caching | Simple caching |

**Rule of thumb:** Use Redis unless you need Memcached's simplicity

---

## Real-World Examples

### Example 1: Facebook News Feed

```
User opens app:

1. Check Redis: GET feed:user:123
   ├─> Cache HIT: Return feed (1ms)
   └─> Cache MISS: Generate feed

2. Generate feed:
   ├─> Query database for friends' recent posts
   ├─> Rank posts by algorithm
   ├─> Store in Redis: SET feed:user:123 (TTL: 5 minutes)
   └─> Return feed (100ms)

3. Next request within 5 min: Cache HIT (1ms)
```

### Example 2: Amazon Product Page

```
Request: /product/12345

Caching strategy:
- Product details: Cache for 1 hour
- Product images: CDN cache for 1 day
- Price: Cache for 5 minutes (changes frequently)
- Inventory: Don't cache (real-time)
- Reviews: Cache for 15 minutes
- User session: Redis cache
```

### Example 3: Twitter Timeline

```
Write (user tweets):
1. Store tweet in database
2. Push to followers' timelines (fanout)
3. Cache in Redis: LIST timeline:user:123

Read (user views timeline):
1. GET LIST timeline:user:123 (1ms)
2. If cache miss: Build from database (100ms)

Result: 100ms → 1ms (100x faster!)
```

---

## Cache Stampede Problem

**The Problem:**

```
Popular cache entry expires

Suddenly, 1000 requests arrive:
Request 1: Cache MISS → Query database
Request 2: Cache MISS → Query database
...
Request 1000: Cache MISS → Query database

Database overwhelmed! 😱
```

**Solution 1: Locking**

```python
lock = cache.lock("user:123:lock")

if lock.acquire(timeout=5):
    user = database.query(...)
    cache.set("user:123", user)
    lock.release()
else:
    # Another request is fetching, wait and retry
    time.sleep(0.1)
    return cache.get("user:123")
```

**Solution 2: Probabilistic Early Expiration**

```python
# Instead of TTL=3600 (exact)
# Random TTL between 3600-3660
ttl = 3600 + random(0, 60)
cache.set("user:123", user, ttl=ttl)

# Entries expire at different times!
```

**Solution 3: Background Refresh**

```python
# Refresh cache before it expires
if ttl < 60:  # Less than 1 minute left
    async_task.refresh_cache("user:123")
```

---

## Cache Monitoring & Metrics

**Key Metrics:**

```
1. Hit Rate = Cache Hits / Total Requests
   └─> Goal: >80% for good cache effectiveness

2. Miss Rate = Cache Misses / Total Requests
   └─> Goal: <20%

3. Eviction Rate = Evictions / Time
   └─> High = cache too small or poor eviction policy

4. Latency
   └─> P50, P95, P99 response times

5. Memory Usage
   └─> How full is your cache?
```

**Example Dashboard:**
```
Cache Hit Rate: 87%
Cache Miss Rate: 13%
Avg Latency (HIT): 1.2ms
Avg Latency (MISS): 98ms
Memory Used: 8.2GB / 16GB (51%)
Evictions/hour: 12,400
```

---

## Best Practices

### 1. Choose the Right TTL

```
Frequently changing data: Short TTL (1-5 min)
├─> Stock prices, sports scores

Rarely changing data: Long TTL (1-24 hours)
├─> Product descriptions, blog posts

Static data: Very long TTL (7-30 days)
└─> User profile pictures, company logos
```

### 2. Cache Warm-Up

```
Don't wait for cache misses!

On deployment:
1. Preload popular items into cache
2. Gradually increase traffic
3. Monitor hit rates
```

### 3. Handle Cache Failures Gracefully

```python
try:
    user = cache.get("user:123")
except CacheConnectionError:
    # Fallback to database
    user = database.query(...)
    
# Never let cache failure break your app!
```

### 4. Monitor and Alert

```
Alert if:
- Hit rate drops below 70%
- Latency increases significantly
- Memory usage > 85%
- Eviction rate spikes
```

### 5. Use Namespacing

```
❌ BAD:
cache.set("123", user_data)  # Collision with order:123?

✅ GOOD:
cache.set("user:123", user_data)
cache.set("order:123", order_data)
```

---

## When NOT to Cache

**Don't cache if:**

1. **Data changes very frequently**
   - Real-time sensor data
   - Live auction bids

2. **Data is rarely accessed**
   - Cold data
   - Archive data

3. **Data is already fast to fetch**
   - Simple calculations
   - Data already in memory

4. **Consistency is critical**
   - Financial transactions
   - Inventory (unless using write-through)

5. **Cache overhead > benefit**
   - Serialization cost
   - Network latency to cache

---

## Key Concepts to Remember

1. **Caching = Speed**: Can improve performance 10-1000x
2. **Cache everywhere**: Browser, CDN, App, Database
3. **80/20 rule**: Cache the hot 20% that gets 80% of requests
4. **Cache-aside** is most common pattern
5. **LRU + TTL** is usually the best eviction strategy
6. **Cache invalidation is hard** - plan your strategy
7. **Always handle cache failures gracefully**
8. **Monitor cache hit rates** - should be >70%
9. **Redis > Memcached** for most use cases
10. **Distributed caching** requires consistent hashing

---

## Practice Questions

**Q1:** You're designing a blog platform. What caching strategy would you use for:
- Blog post content?
- Author profile?
- Comment count?
- User sessions?

Include TTL values and explain your reasoning.

<details>
<summary>View Answer</summary>

| Content | Pattern | TTL | Reasoning |
|---------|---------|-----|-----------|
| **Blog post content** | Cache-aside | 1 hour | Posts rarely change; invalidate on edit |
| **Author profile** | Cache-aside | 15 minutes | Updated occasionally; short TTL for freshness |
| **Comment count** | Write-through or short TTL | 30 seconds | Frequently changes; users expect near-real-time |
| **User sessions** | Write-through | 24 hours (sliding) | Must be consistent; extend on activity |

**Implementation details:**
- Blog posts: Cache full rendered HTML for maximum speed
- Author profile: Invalidate cache when user updates profile
- Comment count: Use Redis INCR for atomic updates, or accept slight staleness
- Sessions: Use Redis with `EXPIRE` command, refresh TTL on each request

</details>

**Q2:** Your cache has a 60% hit rate. Is this good or bad? What would you do to improve it?

<details>
<summary>View Answer</summary>

**60% is mediocre** - most well-optimized caches achieve 85-95%+.

**Investigation steps:**

1. **Analyze cache misses:**
   - Cold misses (first access)? → Pre-warm cache
   - Capacity misses (eviction)? → Increase cache size
   - Conflict misses (bad key distribution)? → Improve hashing

2. **Check TTL values:**
   - Too short? Data expires before reuse
   - Increase TTL for stable data

3. **Review access patterns:**
   - Long-tail distribution? Many unique keys accessed once
   - Consider caching only "hot" data

**Improvement strategies:**

- **Increase cache size:** More memory = fewer evictions
- **Extend TTL:** Keep data longer (if consistency allows)
- **Pre-warming:** Load popular data at startup
- **Better eviction policy:** LRU might not be optimal; try LFU
- **Cache more aggressively:** Cache computed results, not just DB queries
- **Multi-tier caching:** Local cache + distributed cache

**Target:** 85%+ for most applications, 95%+ for read-heavy workloads.

</details>

**Q3:** You're using cache-aside pattern. A user updates their profile. Walk through the exact steps (cache operations and database operations) to ensure consistency.

<details>
<summary>View Answer</summary>

**Correct sequence (Write-Invalidate):**

```
1. UPDATE database
   → UPDATE users SET name='New Name' WHERE id=123

2. DELETE from cache (invalidate)
   → DELETE cache key "user:123"

3. Next READ triggers cache refill
   → Cache miss → Read from DB → Write to cache
```

**Why this order?**

- **DB first:** If cache delete fails, data is still correct in DB; cache will eventually expire
- **Delete, don't update:** Avoids race conditions where stale data overwrites fresh data

**What NOT to do:**

```
❌ Wrong: Update cache, then database
   - If DB update fails, cache has wrong data

❌ Wrong: Update cache instead of delete
   - Race condition: another process might write stale data
```

**Race condition example (why delete is safer):**

```
Thread A: Update user name to "Alice"
Thread B: Read user (gets old name "Bob")

With UPDATE cache:
1. A: Update DB to "Alice"
2. B: Read DB (gets "Bob" - stale read)
3. A: Update cache to "Alice"
4. B: Update cache to "Bob" ← WRONG! Overwrites fresh data

With DELETE cache:
1. A: Update DB to "Alice"
2. A: Delete cache
3. B: Read → cache miss → reads "Alice" from DB ← CORRECT
```

</details>

**Q4:** Your Redis cache is at 95% memory capacity and evicting frequently. What are your options? Discuss trade-offs of each.

<details>
<summary>View Answer</summary>

| Option | Pros | Cons |
|--------|------|------|
| **Add more memory** | Simple, no code changes | Cost, vertical scaling limits |
| **Add more Redis nodes** | Horizontal scaling | Complexity, data distribution |
| **Reduce TTL** | Free up space faster | Lower hit rate |
| **Cache less data** | Reduces memory needs | Lower hit rate, code changes |
| **Better eviction policy** | Keeps hot data | Might not help if all data is hot |
| **Compress values** | 50-90% space savings | CPU overhead, complexity |
| **Use smaller data types** | Significant savings | Requires refactoring |

**Recommended approach:**

1. **First:** Analyze what's in the cache
   ```
   redis-cli --bigkeys
   redis-cli DEBUG OBJECT key
   ```

2. **Quick wins:**
   - Remove unused/stale keys
   - Reduce TTL on less critical data
   - Switch to `volatile-lfu` eviction (evict least frequently used)

3. **Medium-term:**
   - Add Redis node (cluster mode)
   - Implement compression for large values

4. **Review architecture:**
   - Are you caching things that shouldn't be cached?
   - Can you cache at a different layer (CDN, application)?

</details>

**Q5:** Compare these two approaches for a product inventory system:
- Approach A: Cache product inventory with 1-minute TTL
- Approach B: Write-through cache for inventory

Which is better and why?

<details>
<summary>View Answer</summary>

**Approach B (Write-through) is better for inventory.**

**Why:**

| Aspect | Approach A (TTL) | Approach B (Write-through) |
|--------|------------------|---------------------------|
| **Consistency** | Up to 1 min stale | Always current |
| **Overselling risk** | HIGH - stale count | LOW - accurate count |
| **Complexity** | Simple | More complex |
| **Write performance** | Fast | Slower (2 writes) |

**Inventory is critical:**
- Showing "5 in stock" when there's actually 0 = angry customers
- Overselling = refunds, bad reviews, lost trust
- 1-minute staleness is unacceptable for inventory

**Implementation with write-through:**

```python
def purchase_item(product_id):
    # Atomic operation
    with transaction:
        # 1. Update database
        db.execute("UPDATE products SET stock = stock - 1 WHERE id = ?", product_id)

        # 2. Update cache (write-through)
        new_stock = db.query("SELECT stock FROM products WHERE id = ?", product_id)
        cache.set(f"product:{product_id}:stock", new_stock)
```

**Even better: Use Redis for inventory directly**

```python
# Atomic decrement
remaining = redis.decr(f"inventory:{product_id}")
if remaining < 0:
    redis.incr(f"inventory:{product_id}")  # Rollback
    raise OutOfStockError()
# Then async sync to database
```

</details>

**Q6:** You have 3 cache servers. How would you ensure that the key "user:123" always goes to the same cache server, even if a server is added or removed?

<details>
<summary>View Answer</summary>

**Use Consistent Hashing!**

**Simple modulo (bad):**
```
server = hash("user:123") % 3  → Server 1

Add server 4:
server = hash("user:123") % 4  → Server 2 (DIFFERENT!)
```
~75% of keys move to different servers!

**Consistent hashing (good):**
```
1. Create a hash ring (0 to 2^32)
2. Place servers at positions on the ring
3. Place keys at positions on the ring
4. Key goes to next server clockwise

Add server 4:
- Only ~25% of keys move (1/N)
- "user:123" likely stays on same server
```

**Implementation options:**

1. **Client-side library:**
   ```python
   from uhashring import HashRing
   ring = HashRing(nodes=['server1', 'server2', 'server3'])
   server = ring.get_node('user:123')
   ```

2. **Redis Cluster:** Built-in consistent hashing with 16384 hash slots

3. **Memcached with ketama:** Consistent hashing algorithm

**This is covered in depth in Lesson 10!**

</details>

**Q7:** Design a caching strategy for an e-commerce flash sale where:
- 10,000 users trying to buy 100 products
- Inventory must be accurate
- System must stay fast

<details>
<summary>View Answer</summary>

**Architecture:**

```
[Users] → [CDN] → [Load Balancer] → [App Servers] → [Redis] → [Database]
                                          ↓
                                    [Queue] → [Order Processor]
```

**Strategy:**

1. **Pre-cache everything:**
   - Product details in Redis before sale starts
   - Pre-render product pages, cache in CDN

2. **Inventory in Redis (not DB):**
   ```python
   # Atomic check-and-decrement
   remaining = redis.eval("""
       local stock = redis.call('GET', KEYS[1])
       if tonumber(stock) > 0 then
           return redis.call('DECR', KEYS[1])
       end
       return -1
   """, 1, f"inventory:{product_id}")

   if remaining >= 0:
       queue.push(order_details)  # Process async
   else:
       return "Out of stock"
   ```

3. **Queue orders for processing:**
   - Don't hit database on hot path
   - Process orders asynchronously
   - Database can handle queue at its own pace

4. **CDN for static content:**
   - Product images, CSS, JS cached at edge
   - Only inventory check hits origin

5. **Rate limiting:**
   - Limit purchases per user
   - Prevent bots with CAPTCHA

6. **Graceful degradation:**
   - If Redis fails, show "high demand, try again"
   - Don't fall back to database (will crash)

**Key insight:** Keep the hot path (checking inventory) entirely in Redis. Database is only for durability, not real-time operations.

</details>

---

## Next Up

In Lesson 6, we'll dive into **Storage Systems: Files, Blocks, and Objects** - understanding where and how to store your data long-term!
