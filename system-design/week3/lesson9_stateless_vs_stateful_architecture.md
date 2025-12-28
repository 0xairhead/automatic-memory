# Week 3, Lesson 9: Stateless vs Stateful Architecture

## Table of Contents
- [What is State?](#what-is-state)
- [The Core Difference](#the-core-difference)
  - [Stateful Server](#stateful-server)
  - [Stateless Server](#stateless-server)
- [Why Stateless is Better for Scaling](#why-stateless-is-better-for-scaling)
  - [Stateful Scaling Problem](#stateful-scaling-problem)
  - [Stateless Scaling Freedom](#stateless-scaling-freedom)
- [Types of State](#types-of-state)
  - [1. Session State](#1-session-state)
  - [2. Application State](#2-application-state)
  - [3. Resource State](#3-resource-state)
- [Making Services Stateless](#making-services-stateless)
  - [Strategy 1: Client-Side State](#strategy-1-client-side-state)
  - [Strategy 2: External Session Store](#strategy-2-external-session-store)
  - [Strategy 3: Token-Based Authentication (JWT)](#strategy-3-token-based-authentication-jwt)
- [When You Need Statefulness](#when-you-need-statefulness)
  - [WebSocket Connections](#websocket-connections)
  - [Long-Running Operations](#long-running-operations)
  - [Caching User Data](#caching-user-data)
- [Session Management Strategies](#session-management-strategies)
  - [Strategy 1: Sticky Sessions](#strategy-1-sticky-sessions)
  - [Strategy 2: Session Replication](#strategy-2-session-replication)
  - [Strategy 3: Centralized Session Store](#strategy-3-centralized-session-store)
  - [Strategy 4: Client-Side Sessions (JWT)](#strategy-4-client-side-sessions-jwt)
- [Comparison Table](#comparison-table)
- [Real-World Architecture Patterns](#real-world-architecture-patterns)
  - [Pattern 1: Stateless API + External Everything](#pattern-1-stateless-api--external-everything)
  - [Pattern 2: Hybrid (Stateless HTTP + Stateful WebSocket)](#pattern-2-hybrid-stateless-http--stateful-websocket)
  - [Pattern 3: Microservices with Shared Auth](#pattern-3-microservices-with-shared-auth)
- [Converting Stateful to Stateless](#converting-stateful-to-stateless)
- [Performance Considerations](#performance-considerations)
- [Common Mistakes](#common-mistakes)
- [Key Concepts to Remember](#key-concepts-to-remember)
- [Practice Questions](#practice-questions)
- [Next Up](#next-up)

---

This is the final lesson of Week 3! We've covered scaling strategies and database techniques. Now we'll tackle a fundamental architectural decision that determines how easily your system can scale: **stateless vs stateful design**.

---

## What is State?

**State = Data that persists between requests**

```
Stateful Example:
Request 1: User logs in
Server: "I'll remember this user is logged in"
         Stores: session['user_123'] = {logged_in: true}

Request 2: User views dashboard
Server: "Ah, I remember you're logged in!"
         Reads: session['user_123']
         Returns dashboard

The server REMEMBERS things between requests.
That memory is STATE.
```

```
Stateless Example:
Request 1: User logs in
Server: "Here's a token proving you're logged in"
         Returns: JWT token

Request 2: User views dashboard (sends token)
Server: "Token is valid, here's your dashboard"
         Validates token, returns dashboard

The server DOESN'T remember anything.
Each request contains everything needed.
```

---

## The Core Difference

### Stateful Server

```
┌─────────────────────────────────────────┐
│              Stateful Server            │
├─────────────────────────────────────────┤
│  Memory:                                │
│  ┌─────────────────────────────────┐    │
│  │ session_abc: {user: 'alice'}   │    │
│  │ session_xyz: {user: 'bob'}     │    │
│  │ cart_abc: [{item: 'book'}]     │    │
│  │ cart_xyz: [{item: 'pen'}]      │    │
│  └─────────────────────────────────┘    │
│                                         │
│  Server holds user-specific data!       │
└─────────────────────────────────────────┘

Request must go to THIS server
because only THIS server has the data!
```

### Stateless Server

```
┌─────────────────────────────────────────┐
│             Stateless Server            │
├─────────────────────────────────────────┤
│  Memory:                                │
│  ┌─────────────────────────────────┐    │
│  │  (application code only)       │    │
│  │  (no user-specific data)       │    │
│  └─────────────────────────────────┘    │
│                                         │
│  Server holds NO user data!             │
│  Everything comes with the request.     │
└─────────────────────────────────────────┘

Request can go to ANY server
because ALL servers are identical!
```

---

## Why Stateless is Better for Scaling

### Stateful Scaling Problem

```
3 Stateful Servers:

[Server 1]          [Server 2]          [Server 3]
session: alice      session: bob        session: carol
cart: alice's       cart: bob's         cart: carol's

Alice's requests MUST go to Server 1!
Bob's requests MUST go to Server 2!

What happens when:

1. Server 1 crashes?
   → Alice loses her session! 😱
   → Alice loses her cart! 😱
   → Alice must log in again!

2. Need to add Server 4?
   → New users can go there
   → But existing users stuck on old servers
   → Uneven load distribution!

3. Need to remove Server 2?
   → Must migrate Bob's session somehow
   → Or Bob loses everything!
```

### Stateless Scaling Freedom

```
3 Stateless Servers:

[Server 1]          [Server 2]          [Server 3]
(no user data)      (no user data)      (no user data)

Alice's request → ANY server works!
Bob's request → ANY server works!

What happens when:

1. Server 1 crashes?
   → Next request goes to Server 2 or 3
   → User doesn't notice! ✅
   → No data lost (it's in external store)

2. Need to add Server 4?
   → Load balancer sends traffic there
   → Immediately useful! ✅

3. Need to remove Server 2?
   → Stop sending traffic
   → Shut down
   → Zero migration needed! ✅
```

---

## Types of State

### 1. Session State

**User's authentication and identity**

```
Examples:
- Is user logged in?
- What's their user ID?
- What permissions do they have?
- When did they log in?

Stateful approach:
Server memory: sessions['abc123'] = {
    userId: 42,
    loggedIn: true,
    role: 'admin'
}

Stateless approach:
JWT token contains all info:
{
    userId: 42,
    loggedIn: true,
    role: 'admin',
    exp: 1699999999
}
```

### 2. Application State

**Temporary data during user's session**

```
Examples:
- Shopping cart contents
- Multi-step form progress
- Search filters
- UI preferences (dark mode)

Stateful approach:
Server memory: carts['user_42'] = [
    {productId: 1, qty: 2},
    {productId: 5, qty: 1}
]

Stateless approach:
External store (Redis): carts:user_42 = [...]
Or client-side: localStorage, cookies
```

### 3. Resource State

**Business data in databases**

```
Examples:
- User profiles
- Orders
- Products
- Messages

This is ALWAYS external (database)
Not what we mean by "stateful server"
All servers access same database
```

---

## Making Services Stateless

### Strategy 1: Client-Side State

**Store state on the client**

```
Browser stores:
┌─────────────────────────────────────┐
│  localStorage:                      │
│    cart: [{id: 1, qty: 2}]          │
│    darkMode: true                   │
│    recentSearches: ['shoes']        │
│                                     │
│  Cookies:                           │
│    session_token: 'abc123...'       │
│                                     │
│  JWT in memory/storage:             │
│    eyJhbGciOiJIUzI1NiIs...          │
└─────────────────────────────────────┘

Server is completely stateless!
Client sends everything with each request.
```

**Implementation:**
```javascript
// Client sends cart with every request
POST /api/checkout
{
    "cart": [
        {"productId": 1, "quantity": 2},
        {"productId": 5, "quantity": 1}
    ]
}

// Server processes and responds
// Server remembers NOTHING
```

**Pros:**
- ✅ Server is perfectly stateless
- ✅ Scales infinitely
- ✅ No session storage costs

**Cons:**
- ❌ Limited storage (cookies: 4KB, localStorage: 5-10MB)
- ❌ Security concerns (client can tamper)
- ❌ Can't share across devices
- ❌ Lost if user clears browser data

**Good for:**
- UI preferences
- Non-sensitive temporary data
- Small amounts of data

### Strategy 2: External Session Store

**Move state to shared storage**

```
Before (Stateful):
[Server 1] ← session data
[Server 2] ← session data
[Server 3] ← session data

Each server has different sessions!

After (Stateless with shared store):
[Server 1] ─┐
[Server 2] ─┼─→ [Redis] ← ALL sessions here!
[Server 3] ─┘

All servers access same session store!
Any server can handle any request!
```

**Implementation:**
```javascript
// Server 1 handles login
app.post('/login', async (req, res) => {
    const user = await authenticate(req.body)
    const sessionId = generateId()

    // Store in Redis (external!)
    await redis.set(`session:${sessionId}`, JSON.stringify({
        userId: user.id,
        role: user.role,
        createdAt: Date.now()
    }), 'EX', 3600)  // 1 hour expiry

    res.cookie('sessionId', sessionId)
    res.json({success: true})
})

// Server 2 handles subsequent request
app.get('/dashboard', async (req, res) => {
    const sessionId = req.cookies.sessionId

    // Read from Redis (same data Server 1 wrote!)
    const session = await redis.get(`session:${sessionId}`)
    if (!session) return res.status(401).send('Not logged in')

    const userData = JSON.parse(session)
    // ... return dashboard
})
```

**Architecture:**
```
              [Load Balancer]
                     │
       ┌─────────────┼─────────────┐
       ↓             ↓             ↓
  [Server 1]    [Server 2]    [Server 3]
       │             │             │
       └─────────────┼─────────────┘
                     ↓
               ┌───────────┐
               │   Redis   │
               │  Cluster  │
               └───────────┘
                     │
               (session data)
```

**Pros:**
- ✅ Servers are stateless
- ✅ Sessions survive server restarts
- ✅ Easy horizontal scaling
- ✅ Can share sessions across services

**Cons:**
- ❌ Extra infrastructure (Redis)
- ❌ Extra latency (network call)
- ❌ Session store becomes critical
- ❌ Costs for session storage

**Good for:**
- Most web applications
- Complex session data
- Multi-device sessions

### Strategy 3: Token-Based Authentication (JWT)

**Encode state in signed tokens**

```
Traditional Session:
┌─────────────┐        ┌─────────────┐
│   Client    │        │   Server    │
├─────────────┤        ├─────────────┤
│ session_id  │───────→│ Look up ID  │
│ = abc123    │        │ in store    │
└─────────────┘        └─────────────┘
                              ↓
                       ┌─────────────┐
                       │Session Store│
                       └─────────────┘

JWT (JSON Web Token):
┌─────────────────────────────────────┐
│   Client                            │
├─────────────────────────────────────┤
│ JWT = eyJhbGciOiJIUzI1NiIs...       │
│                                     │
│ Contains (encoded):                 │
│ {                                   │
│   userId: 42,                       │
│   role: 'admin',                    │
│   exp: 1699999999                   │
│ }                                   │
└─────────────────────────────────────┘
         │
         ↓
┌─────────────────────────────────────┐
│   Server                            │
├─────────────────────────────────────┤
│ 1. Verify signature (cryptographic) │
│ 2. Decode payload                   │
│ 3. No database/store lookup!        │
└─────────────────────────────────────┘
```

**How JWT Works:**
```
JWT Structure:
header.payload.signature

eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.    ← Header (algorithm)
eyJ1c2VySWQiOjQyLCJyb2xlIjoiYWRtaW4ifQ.  ← Payload (data)
SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV...    ← Signature (proof)

Signature = HMAC(header + payload, SECRET_KEY)

Server can verify token wasn't tampered with!
No database lookup needed!
```

**Implementation:**
```javascript
// Login - Create JWT
app.post('/login', async (req, res) => {
    const user = await authenticate(req.body)

    const token = jwt.sign({
        userId: user.id,
        role: user.role,
        email: user.email
    }, SECRET_KEY, {expiresIn: '24h'})

    res.json({token})
})

// Any request - Verify JWT
app.get('/dashboard', (req, res) => {
    const token = req.headers.authorization?.split(' ')[1]

    try {
        const decoded = jwt.verify(token, SECRET_KEY)
        // decoded = {userId: 42, role: 'admin', email: '...'}
        // No database lookup! Just cryptographic verification!

        return res.json({dashboard: '...'})
    } catch (err) {
        return res.status(401).json({error: 'Invalid token'})
    }
})
```

**Pros:**
- ✅ Truly stateless (no session store!)
- ✅ Scalable across services/regions
- ✅ No database lookup per request
- ✅ Works for microservices

**Cons:**
- ❌ Can't invalidate individual tokens (until expiry)
- ❌ Token size grows with claims
- ❌ Sensitive data in token (must be encrypted)
- ❌ Token theft = full access until expiry

**The Revocation Problem:**
```
User changes password at 10:00 AM
Old JWT (stolen) expires at 11:00 AM

Attacker can use stolen JWT until 11:00 AM!
Server has no way to know password changed!

Solutions:
1. Short expiry (15 min) + refresh tokens
2. Token blacklist (adds state back!)
3. Token versioning in database
```

---

## When You Need Statefulness

### WebSocket Connections

```
WebSockets are inherently stateful!

[Client] ←──────→ [Server]
         Persistent connection
         Server remembers the client

Can't just switch servers mid-connection!
```

**Solutions:**
```
1. Sticky Sessions
   Same client → same server (for WebSocket)

2. External Pub/Sub
   [Client 1] ↔ [Server 1] ─┐
                            ├─→ [Redis Pub/Sub]
   [Client 2] ↔ [Server 2] ─┘

   Servers publish/subscribe through Redis
   Messages reach correct client

3. Connection Registry
   Track which server holds which connection
   Route messages accordingly
```

### Long-Running Operations

```
User starts video encoding:
Request 1: "Start encoding video_123"
Server: "OK, processing..."

Request 2: "What's the status?"
Server: "What video? I don't remember!" 😕
```

**Solutions:**
```
1. Job Queue (Recommended)
   Store job in external queue (Redis, RabbitMQ)
   Any worker can process
   Status tracked externally

2. Database State
   Store operation status in database
   Any server can read status

3. Async with Callbacks
   Client provides callback URL
   Server notifies when done
   No need to remember!
```

### Caching User Data

```
Optimization: Cache user profile in server memory

Request 1 → Server 1: Load profile, cache it
Request 2 → Server 2: Cache miss! Load again 😰

Caching creates implicit state!
```

**Solutions:**
```
1. External Cache (Redis)
   All servers share same cache
   Still "stateless" servers

2. CDN/Edge Caching
   Cache at network edge
   Server doesn't hold state

3. Accept Cache Misses
   If cache hit ratio is high enough
   Misses are acceptable overhead
```

---

## Session Management Strategies

### Strategy 1: Sticky Sessions

**Route user to same server**

```
[Load Balancer]
      │
      │ Cookie: server_id=2
      │
      └────────────→ [Server 2]

User always goes to Server 2
Server 2 has their session in memory
```

**Implementation (NGINX):**
```nginx
upstream backend {
    ip_hash;  # Same IP → same server
    server backend1.example.com;
    server backend2.example.com;
    server backend3.example.com;
}

# Or cookie-based
upstream backend {
    server backend1.example.com;
    server backend2.example.com;

    sticky cookie srv_id expires=1h;
}
```

**Pros:**
- ✅ Simple to implement
- ✅ No external session store needed
- ✅ Fast (memory access)

**Cons:**
- ❌ Server failure = lost sessions
- ❌ Uneven load distribution
- ❌ Can't easily scale down
- ❌ Deployment requires draining

**When to use:**
- Legacy applications
- Quick fix for stateful apps
- Development/testing

### Strategy 2: Session Replication

**Copy sessions to all servers**

```
[Server 1] ←──→ [Server 2] ←──→ [Server 3]
     ↑              ↑               ↑
     └──────────────┴───────────────┘
           Session Replication

All servers have ALL sessions!
Any server can handle any request!
```

**Pros:**
- ✅ Any server can handle request
- ✅ Server failure = no data loss
- ✅ Fast reads (local memory)

**Cons:**
- ❌ Memory: N servers × M sessions
- ❌ Write overhead (replicate to all)
- ❌ Consistency delays
- ❌ Doesn't scale well (> 10 servers)

**When to use:**
- Small clusters (3-5 servers)
- High session read rate
- Low session write rate

### Strategy 3: Centralized Session Store

**External store for all sessions**

```
[Server 1] ─┐
[Server 2] ─┼──→ [Redis Cluster]
[Server 3] ─┘    (all sessions)
```

Already covered above. This is the **recommended approach** for most applications.

### Strategy 4: Client-Side Sessions (JWT)

**No server-side session storage**

Already covered above. Best for **APIs and microservices**.

---

## Comparison Table

| Aspect | Sticky Sessions | Session Replication | Centralized Store | JWT |
|--------|----------------|---------------------|-------------------|-----|
| **Scalability** | Poor | Limited | Excellent | Excellent |
| **Availability** | Poor | Good | Good | Excellent |
| **Complexity** | Low | Medium | Medium | Low |
| **Latency** | Best | Good | +1-5ms | Best |
| **Server Memory** | High | Very High | Low | None |
| **Revocation** | Instant | Instant | Instant | Hard |
| **Cross-Service** | No | No | Yes | Yes |

---

## Real-World Architecture Patterns

### Pattern 1: Stateless API + External Everything

```
┌─────────────────────────────────────────────────┐
│                   CLIENT                         │
│  ┌─────────────────────────────────────────┐    │
│  │ JWT Token (auth)                        │    │
│  │ localStorage (preferences)              │    │
│  └─────────────────────────────────────────┘    │
└─────────────────────────────────────────────────┘
                        │
                        ↓
┌─────────────────────────────────────────────────┐
│              LOAD BALANCER                       │
│         (any server, any request)               │
└─────────────────────────────────────────────────┘
                        │
        ┌───────────────┼───────────────┐
        ↓               ↓               ↓
┌─────────────┐ ┌─────────────┐ ┌─────────────┐
│  API Pod 1  │ │  API Pod 2  │ │  API Pod 3  │
│ (stateless) │ │ (stateless) │ │ (stateless) │
└─────────────┘ └─────────────┘ └─────────────┘
        │               │               │
        └───────────────┼───────────────┘
                        ↓
        ┌───────────────┼───────────────┐
        ↓               ↓               ↓
┌─────────────┐ ┌─────────────┐ ┌─────────────┐
│    Redis    │ │  PostgreSQL │ │     S3      │
│   (cache)   │ │    (data)   │ │   (files)   │
└─────────────┘ └─────────────┘ └─────────────┘

API servers hold ZERO state
All state in managed services
Scale API pods infinitely!
```

### Pattern 2: Hybrid (Stateless HTTP + Stateful WebSocket)

```
HTTP Requests (Stateless):
┌────────┐    ┌─────────────┐    ┌─────────────┐
│ Client │───→│ Any API Pod │───→│  Database   │
└────────┘    └─────────────┘    └─────────────┘

WebSocket (Stateful with Sticky):
┌────────┐    ┌─────────────┐    ┌─────────────┐
│ Client │←──→│ Specific WS │←──→│ Redis       │
│        │    │ Pod (sticky)│    │ Pub/Sub     │
└────────┘    └─────────────┘    └─────────────┘

HTTP: Fully stateless, any pod
WebSocket: Sticky to specific pod, Redis for cross-pod messaging
```

### Pattern 3: Microservices with Shared Auth

```
┌─────────────────────────────────────────────────┐
│                   API Gateway                    │
│              (validates JWT once)               │
└─────────────────────────────────────────────────┘
                        │
        ┌───────────────┼───────────────┐
        ↓               ↓               ↓
┌─────────────┐ ┌─────────────┐ ┌─────────────┐
│   Users     │ │   Orders    │ │   Products  │
│   Service   │ │   Service   │ │   Service   │
│ (stateless) │ │ (stateless) │ │ (stateless) │
└─────────────┘ └─────────────┘ └─────────────┘
        │               │               │
        ↓               ↓               ↓
┌─────────────┐ ┌─────────────┐ ┌─────────────┐
│  Users DB   │ │  Orders DB  │ │ Products DB │
└─────────────┘ └─────────────┘ └─────────────┘

JWT validated at gateway
Services trust gateway's validation
Each service is independently stateless
Each service has own database
```

---

## Converting Stateful to Stateless

### Step 1: Identify State

```
Audit your application:

□ Where are sessions stored?
□ What's in server memory?
□ What happens on server restart?
□ Can requests go to any server?

Common state locations:
- In-memory sessions
- In-memory caches
- File uploads in /tmp
- WebSocket connections
- Background job status
```

### Step 2: Externalize State

```
For each piece of state:

Session data → Redis
File uploads → S3/Cloud Storage
Cache data → Redis/Memcached
Job status → Database/Redis
WebSocket → Add Pub/Sub layer
```

### Step 3: Update Application

```javascript
// Before: In-memory session
const sessions = {}  // BAD!

app.post('/login', (req, res) => {
    sessions[sessionId] = userData  // Stored in memory!
})

// After: External session store
const redis = new Redis()

app.post('/login', async (req, res) => {
    await redis.set(`session:${sessionId}`, JSON.stringify(userData))
})
```

### Step 4: Verify Statelessness

```
Test: Can you restart any server without affecting users?

1. User logs in through Server 1
2. Restart Server 1
3. User's next request goes to Server 2
4. Does it work? ✅ = Stateless!
```

---

## Performance Considerations

### Latency Impact

```
Stateful (in-memory):
Request → Memory lookup → Response
Latency: ~0.1ms

Stateless (Redis):
Request → Network → Redis → Network → Response
Latency: ~1-5ms

Stateless (JWT):
Request → CPU (verify signature) → Response
Latency: ~0.5ms
```

### Optimizations

**1. Connection Pooling**
```javascript
// Reuse Redis connections
const redis = new Redis({
    maxRetriesPerRequest: 3,
    enableReadyCheck: true,
    connectionPool: {
        min: 5,
        max: 20
    }
})
```

**2. Local Caching of Sessions**
```javascript
// Cache validated sessions briefly
const sessionCache = new LRU({max: 1000, ttl: 60000})  // 1 min

async function getSession(sessionId) {
    // Check local cache first
    let session = sessionCache.get(sessionId)
    if (session) return session

    // Fall back to Redis
    session = await redis.get(`session:${sessionId}`)
    if (session) {
        sessionCache.set(sessionId, session)
    }
    return session
}
```

**3. Batch Operations**
```javascript
// Instead of multiple Redis calls
const user = await redis.get('user:1')
const cart = await redis.get('cart:1')
const prefs = await redis.get('prefs:1')

// Use pipeline
const [user, cart, prefs] = await redis.pipeline()
    .get('user:1')
    .get('cart:1')
    .get('prefs:1')
    .exec()
```

---

## Common Mistakes

### Mistake 1: Hidden State

```javascript
// Looks stateless, but isn't!
let requestCount = 0  // Module-level variable!

app.get('/api/data', (req, res) => {
    requestCount++  // Different value on each server!
    res.json({count: requestCount})
})

// Server 1: count = 100
// Server 2: count = 50
// User sees inconsistent values!
```

### Mistake 2: File System State

```javascript
// Uploads stored locally
app.post('/upload', (req, res) => {
    const path = `/tmp/uploads/${filename}`
    fs.writeFileSync(path, file)  // Only on THIS server!
    res.json({path})
})

app.get('/download/:file', (req, res) => {
    const path = `/tmp/uploads/${req.params.file}`
    res.sendFile(path)  // Might hit different server! ❌
})

// Fix: Use S3/Cloud Storage
```

### Mistake 3: Assuming Request Order

```javascript
// Dangerous assumption
app.post('/step1', (req, res) => {
    // Assume step2 comes next to same server
    serverMemory.step1Data = req.body
})

app.post('/step2', (req, res) => {
    // Might hit different server!
    const step1 = serverMemory.step1Data  // undefined! ❌
})

// Fix: Pass data through client or external store
```

### Mistake 4: Long-Lived Process State

```javascript
// Background process holds state
let processingQueue = []

app.post('/process', (req, res) => {
    processingQueue.push(req.body.jobId)
    res.json({queued: true})
})

// Only THIS server processes these jobs!
// If server dies, jobs are lost!

// Fix: Use Redis/RabbitMQ for job queue
```

---

## Key Concepts to Remember

1. **Stateless** = Server remembers nothing between requests
2. **Stateful** = Server stores user-specific data in memory
3. **Stateless enables horizontal scaling** without coordination
4. **Three ways to handle state**: Client-side, external store, or tokens
5. **JWT** = Truly stateless auth (but can't revoke easily)
6. **Redis/Memcached** = Popular external session stores
7. **Sticky sessions** = Quick fix but limits scalability
8. **WebSockets are inherently stateful** - need special handling
9. **Design for statelessness from the start** - retrofitting is hard
10. **Test by restarting servers** - users shouldn't notice

---

## Practice Questions

**Q1:** A web application stores shopping carts in server memory. During a deployment (rolling restart), some users lose their carts. Explain why and provide two solutions.

<details>
<summary>View Answer</summary>

**Why Users Lose Carts:**

```
Before deployment:
[Server 1] ← Cart for User A: [{item: 'book', qty: 2}]
[Server 2] ← Cart for User B: [{item: 'pen', qty: 5}]

During rolling restart:
[Server 1] 🔄 Restarting...
    └─> User A's cart stored in memory = LOST!

[Server 2] (still running) ← User B's cart still exists

After Server 1 restarts:
[Server 1] ← Memory empty, no carts!

User A's next request:
"Where's my cart?!" 😱
```

The cart data only exists in server memory. When the server restarts, all in-memory data is cleared.

**Solution 1: External Session Store (Redis)**

```
Architecture:
[Server 1] ─┐
[Server 2] ─┼──→ [Redis] ← All carts stored here!
[Server 3] ─┘

Implementation:
// Save cart to Redis
app.post('/cart/add', async (req, res) => {
    const cartKey = `cart:${req.user.id}`
    await redis.rpush(cartKey, JSON.stringify(req.body.item))
    await redis.expire(cartKey, 86400) // 24 hour TTL
    res.json({success: true})
})

// Load cart from Redis
app.get('/cart', async (req, res) => {
    const cartKey = `cart:${req.user.id}`
    const items = await redis.lrange(cartKey, 0, -1)
    res.json(items.map(JSON.parse))
})

Benefits:
- Servers can restart without data loss
- Any server can access any cart
- Cart survives deployments
```

**Solution 2: Database-Backed Carts**

```
Store carts in PostgreSQL:

CREATE TABLE carts (
    user_id INT PRIMARY KEY,
    items JSONB,
    updated_at TIMESTAMP DEFAULT NOW()
);

Implementation:
app.post('/cart/add', async (req, res) => {
    await db.query(`
        INSERT INTO carts (user_id, items)
        VALUES ($1, $2::jsonb)
        ON CONFLICT (user_id)
        DO UPDATE SET items = carts.items || $2::jsonb,
                      updated_at = NOW()
    `, [req.user.id, JSON.stringify([req.body.item])])
})

Benefits:
- Persistent across restarts
- Carts survive even Redis failures
- Can query/analyze cart data

Drawbacks:
- Slower than Redis
- More database load
```

**Comparison:**

| Solution | Speed | Durability | Cost |
|----------|-------|------------|------|
| Redis | Fast (~1ms) | Good (with persistence) | Medium |
| Database | Slower (~5-10ms) | Excellent | Lower |

**Recommendation:** Use Redis for active carts (fast access), periodically sync to database for persistence.

</details>

**Q2:** Your API uses JWT for authentication. A user reports their account was compromised. How do you immediately revoke their access? What are the trade-offs of your solution?

<details>
<summary>View Answer</summary>

**The Challenge:**

JWTs are stateless - the server doesn't track which tokens are valid. Once issued, a JWT is valid until it expires.

```
Compromised scenario:
- Token issued at 9:00 AM
- Token expires at 9:00 PM (12-hour expiry)
- Account compromised at 10:00 AM
- Attacker has valid token for 11 more hours!
```

**Solution Options:**

**Option 1: Token Blacklist (Recommended for immediate revocation)**

```
Maintain list of revoked tokens in Redis:

// On compromise report
app.post('/revoke-user/:userId', async (req, res) => {
    // Option A: Blacklist specific token
    const token = getUserCurrentToken(req.params.userId)
    await redis.set(`blacklist:${token}`, 'revoked', 'EX', 43200) // Until expiry

    // Option B: Revoke all user tokens by version
    await redis.incr(`user:${req.params.userId}:tokenVersion`)

    res.json({revoked: true})
})

// On every request
function validateToken(req, res, next) {
    const token = req.headers.authorization
    const decoded = jwt.verify(token, SECRET)

    // Check blacklist
    const isBlacklisted = await redis.get(`blacklist:${token}`)
    if (isBlacklisted) return res.status(401).send('Token revoked')

    // Or check token version
    const currentVersion = await redis.get(`user:${decoded.userId}:tokenVersion`)
    if (decoded.tokenVersion < currentVersion) {
        return res.status(401).send('Token revoked')
    }

    next()
}

Trade-offs:
✅ Immediate revocation
❌ Adds Redis lookup to every request (loses stateless benefit)
❌ Must maintain blacklist until token expiry
❌ Adds infrastructure dependency
```

**Option 2: Short-Lived Tokens + Refresh Tokens**

```
Architecture:
- Access token: 15-minute expiry
- Refresh token: 7-day expiry, stored in database

On compromise:
1. Delete refresh token from database
2. Attacker's access token expires in max 15 minutes
3. Can't get new access token without valid refresh token

// Login
app.post('/login', async (req, res) => {
    const accessToken = jwt.sign({userId}, SECRET, {expiresIn: '15m'})
    const refreshToken = generateSecureToken()

    await db.query('INSERT INTO refresh_tokens VALUES ($1, $2, $3)',
        [req.user.id, refreshToken, Date.now() + 7*24*60*60*1000])

    res.json({accessToken, refreshToken})
})

// Revoke
app.post('/revoke/:userId', async (req, res) => {
    await db.query('DELETE FROM refresh_tokens WHERE user_id = $1',
        [req.params.userId])
    // Attacker locked out within 15 minutes!
})

Trade-offs:
✅ Mostly stateless (only refresh requires DB)
✅ Limited exposure window (15 min max)
❌ Not instant (up to 15 min delay)
❌ More complex token management
❌ Requires client to handle token refresh
```

**Option 3: Reduce Token Lifetime Drastically**

```
Ultra-short tokens (5 minutes):
- Compromise window: max 5 minutes
- On password change: just wait 5 minutes

Trade-offs:
✅ Simple implementation
✅ Stays stateless
❌ Frequent token refresh (every 5 min)
❌ Not instant revocation
❌ Poor UX (more refreshes)
```

**Recommended Approach:**

```
Combine strategies:

1. Normal operation:
   - 15-min access tokens
   - Refresh tokens in database

2. On compromise:
   - Delete all refresh tokens for user
   - Add user ID to short-term blacklist (15 min)
   - Force password reset

This gives:
- Mostly stateless operation
- Instant revocation when needed
- Minimal performance impact
```

</details>

**Q3:** Compare these session strategies for a chat application with 1 million concurrent users:
- Sticky sessions
- Redis session store
- JWT tokens

Which would you choose and why?

<details>
<summary>View Answer</summary>

**Requirements Analysis:**

```
Chat application specifics:
- 1 million concurrent users
- Real-time messaging (likely WebSockets)
- Need to track: user identity, online status, current rooms
- High read rate (message delivery)
- Medium write rate (status changes, room joins)
```

**Option Analysis:**

**Sticky Sessions:**
```
Architecture:
[Load Balancer]
├─→ Server 1: Users A-G (142K users)
├─→ Server 2: Users H-N (142K users)
...
└─→ Server 7: Users U-Z (142K users)

Problems:
❌ 1M users ÷ 7 servers = 142K per server
   - Memory: 142K × 5KB session = 710MB per server
   - Too much memory!

❌ Server failure loses 142K users' sessions
   - Massive reconnection storm

❌ Uneven distribution
   - Some letters more common than others

❌ Can't easily add/remove servers
   - Must re-sticky all users

Verdict: NOT SUITABLE for this scale
```

**Redis Session Store:**
```
Architecture:
[WebSocket Servers (stateless)]
├─→ Server 1
├─→ Server 2
...
└─→ Server N
        │
        └──→ [Redis Cluster]
             - session:user123 → {rooms: [...], status: 'online'}
             - Total: 1M sessions × 5KB = 5GB (across cluster)

Advantages:
✅ Servers are stateless - scale easily
✅ Sessions survive server failures
✅ Any server can handle any user
✅ Redis handles 100K+ ops/sec

Implementation:
// Connection handling
async function onConnect(userId, socket) {
    await redis.hset(`session:${userId}`, {
        status: 'online',
        server: SERVER_ID,
        socketId: socket.id
    })
}

// Message routing
async function sendToUser(userId, message) {
    const session = await redis.hgetall(`session:${userId}`)
    if (session.server === SERVER_ID) {
        // User on this server
        sockets[session.socketId].send(message)
    } else {
        // Publish to user's server
        redis.publish(`server:${session.server}`, JSON.stringify({
            userId, message
        }))
    }
}

Verdict: GOOD CHOICE - handles scale well
```

**JWT Tokens:**
```
JWT alone doesn't work well for chat because:

❌ Can't track online status
   - JWT is stateless, can't know if user is connected

❌ Can't route messages
   - Don't know which server has user's WebSocket

❌ WebSockets are inherently stateful
   - Must track connection somewhere

JWT role in chat:
✅ Authentication (validate who user is)
❌ Session management (where user is connected)

Verdict: USE FOR AUTH ONLY, not session management
```

**Recommended Solution: Hybrid Approach**

```
┌────────────────────────────────────────────────────────┐
│                    Architecture                         │
├────────────────────────────────────────────────────────┤
│                                                        │
│  [Client]                                              │
│     │                                                  │
│     ├─── JWT (auth) ───→ Verify identity              │
│     │                                                  │
│     └─── WebSocket ───→ [WS Server Pool]              │
│                              │                         │
│                              ↓                         │
│                    [Redis Cluster]                     │
│                    ├─ Sessions (where is user?)        │
│                    ├─ Pub/Sub (cross-server messages)  │
│                    └─ Presence (online status)         │
│                                                        │
└────────────────────────────────────────────────────────┘

JWT for: Authentication (validate on connect)
Redis for: Session state (which server, which rooms)
Pub/Sub for: Cross-server message routing

This handles 1M users because:
- JWT auth: 0 state per user
- Redis sessions: distributed across cluster
- WebSocket servers: stateless, scale horizontally
- Pub/Sub: efficient cross-server communication
```

**Final Answer:** Redis session store + JWT authentication

</details>

**Q4:** An application has this code:
```javascript
const cache = new Map()
app.get('/user/:id', async (req, res) => {
    if (cache.has(req.params.id)) {
        return res.json(cache.get(req.params.id))
    }
    const user = await db.getUser(req.params.id)
    cache.set(req.params.id, user)
    res.json(user)
})
```
Is this stateful or stateless? What problems could occur with multiple servers?

<details>
<summary>View Answer</summary>

**This is STATEFUL**

The server holds user-specific data (the cache) in memory that persists between requests.

**Problems with Multiple Servers:**

**Problem 1: Cache Inconsistency**
```
Timeline:
1. Request for User 123 → Server 1
   Server 1: Cache miss, loads from DB, caches {name: "Alice"}

2. User 123 updates profile → Server 2
   Server 2: Updates DB to {name: "Alicia"}
   Server 2's cache: empty (never loaded)

3. Request for User 123 → Server 1
   Server 1: Cache HIT! Returns {name: "Alice"} ❌ STALE!

4. Request for User 123 → Server 2
   Server 2: Cache miss, loads from DB
   Returns {name: "Alicia"} ✅ Correct!

User sees different data depending on which server handles request!
```

**Problem 2: Memory Waste / Duplication**
```
With 10 servers and 100K users:

Server 1 cache: 50K users × 5KB = 250MB
Server 2 cache: 60K users × 5KB = 300MB
...
Server 10 cache: 45K users × 5KB = 225MB

Total memory: ~2.5GB
But unique users only need: 500MB in shared cache!

5x memory waste due to duplication!
```

**Problem 3: Cold Start After Restart**
```
Before restart:
Server 1 cache: 50K users (warm, fast)

After restart:
Server 1 cache: empty (cold!)

All requests hit database until cache warms up
Potential database overload during deployments!
```

**Problem 4: Unbounded Memory Growth**
```
const cache = new Map()  // No size limit!

Over time:
- cache.size = 10,000 users ✅
- cache.size = 100,000 users ⚠️
- cache.size = 1,000,000 users 💥 Out of memory!

No eviction policy = memory leak
```

**Solutions:**

**Solution 1: External Cache (Redis)**
```javascript
const redis = require('redis').createClient()

app.get('/user/:id', async (req, res) => {
    // Check Redis (shared across all servers)
    const cached = await redis.get(`user:${req.params.id}`)
    if (cached) {
        return res.json(JSON.parse(cached))
    }

    const user = await db.getUser(req.params.id)

    // Cache in Redis with TTL
    await redis.set(`user:${req.params.id}`, JSON.stringify(user), 'EX', 300)

    res.json(user)
})

Benefits:
✅ Consistent across all servers
✅ Survives server restarts
✅ Single cache, no duplication
✅ Built-in TTL for expiration
```

**Solution 2: Cache with TTL and Size Limit**
```javascript
const LRU = require('lru-cache')
const cache = new LRU({
    max: 10000,        // Max 10K entries
    ttl: 1000 * 60 * 5 // 5 minute TTL
})

// Still has inconsistency issues but won't OOM
// Use when stale data is acceptable
```

**Solution 3: Cache-Aside with Invalidation**
```javascript
// On user update, invalidate across all servers
app.put('/user/:id', async (req, res) => {
    await db.updateUser(req.params.id, req.body)

    // Publish invalidation event
    await redis.publish('cache-invalidation', JSON.stringify({
        type: 'user',
        id: req.params.id
    }))
})

// All servers subscribe
redis.subscribe('cache-invalidation', (message) => {
    const {type, id} = JSON.parse(message)
    cache.delete(`${type}:${id}`)
})
```

**Recommendation:** Use Redis for shared caching in multi-server deployments.

</details>

**Q5:** Design a stateless file upload system that:
- Accepts files up to 100MB
- Allows users to list their uploads
- Works across 10 API servers

<details>
<summary>View Answer</summary>

**Architecture:**

```
                        [Load Balancer]
                              │
              ┌───────────────┼───────────────┐
              ↓               ↓               ↓
        [API Server]    [API Server]    [API Server]
        (stateless)     (stateless)     (stateless)
              │               │               │
              └───────────────┼───────────────┘
                              │
              ┌───────────────┼───────────────┐
              ↓               ↓               ↓
         [S3/Cloud      [PostgreSQL]     [Redis]
          Storage]       (metadata)      (upload
         (files)                         progress)
```

**Upload Flow:**

```
Step 1: Get Upload URL (Presigned URL)
┌────────┐         ┌────────────┐         ┌─────┐
│ Client │──GET───→│ API Server │──────→ │ S3  │
│        │←────────│ (any)      │←──────  │     │
│        │ presigned│            │presigned│     │
│        │   URL   │            │  URL   │     │
└────────┘         └────────────┘         └─────┘

Step 2: Direct Upload to S3
┌────────┐                               ┌─────┐
│ Client │──────────PUT file────────────→│ S3  │
│        │←─────────200 OK──────────────│     │
└────────┘                               └─────┘

Step 3: Confirm Upload
┌────────┐         ┌────────────┐         ┌──────┐
│ Client │──POST──→│ API Server │──────→ │ DB   │
│        │  /done  │ (any)      │metadata │      │
└────────┘         └────────────┘         └──────┘
```

**Implementation:**

```javascript
// Step 1: Generate presigned upload URL
app.post('/uploads/start', async (req, res) => {
    const { filename, contentType } = req.body
    const userId = req.user.id

    // Generate unique key
    const key = `uploads/${userId}/${Date.now()}-${filename}`

    // Create presigned URL (valid 15 min)
    const uploadUrl = await s3.getSignedUrl('putObject', {
        Bucket: 'user-uploads',
        Key: key,
        ContentType: contentType,
        Expires: 900
    })

    // Track pending upload in Redis
    await redis.setex(`pending:${key}`, 3600, JSON.stringify({
        userId,
        filename,
        contentType,
        startedAt: Date.now()
    }))

    res.json({ uploadUrl, key })
})

// Step 2: Client uploads directly to S3 (no server involvement)

// Step 3: Confirm upload complete
app.post('/uploads/complete', async (req, res) => {
    const { key } = req.body
    const userId = req.user.id

    // Verify file exists in S3
    const metadata = await s3.headObject({
        Bucket: 'user-uploads',
        Key: key
    }).promise()

    // Save to database
    await db.query(`
        INSERT INTO uploads (user_id, s3_key, filename, size, content_type, uploaded_at)
        VALUES ($1, $2, $3, $4, $5, NOW())
    `, [userId, key, req.body.filename, metadata.ContentLength, metadata.ContentType])

    // Clean up Redis
    await redis.del(`pending:${key}`)

    res.json({ success: true, size: metadata.ContentLength })
})

// List user's uploads
app.get('/uploads', async (req, res) => {
    const userId = req.user.id

    const uploads = await db.query(`
        SELECT id, filename, size, content_type, uploaded_at
        FROM uploads
        WHERE user_id = $1
        ORDER BY uploaded_at DESC
        LIMIT 100
    `, [userId])

    res.json(uploads.rows)
})

// Get download URL
app.get('/uploads/:id/download', async (req, res) => {
    const upload = await db.query(`
        SELECT s3_key FROM uploads WHERE id = $1 AND user_id = $2
    `, [req.params.id, req.user.id])

    if (!upload.rows[0]) return res.status(404).send('Not found')

    const downloadUrl = await s3.getSignedUrl('getObject', {
        Bucket: 'user-uploads',
        Key: upload.rows[0].s3_key,
        Expires: 3600
    })

    res.json({ downloadUrl })
})
```

**Why This is Stateless:**

```
✅ Files stored in S3 (not server filesystem)
✅ Metadata in PostgreSQL (shared across servers)
✅ Upload progress in Redis (optional, shared)
✅ Any server can handle any request
✅ No server-local state

Stateless benefits:
- Server failure doesn't lose uploads
- Scale to any number of servers
- Rolling deployments don't interrupt uploads
```

**100MB File Handling:**

```
Direct upload to S3 advantages:
- No server memory/disk usage
- No 100MB through API server
- S3 handles large files efficiently
- Multipart upload for files > 100MB

Client-side:
const response = await fetch(presignedUrl, {
    method: 'PUT',
    body: file,
    headers: { 'Content-Type': file.type }
})
```

</details>

**Q6:** Your microservices architecture has 5 services. Users authenticate with Service A, then call Services B, C, D, E. How do you propagate authentication without each service hitting a session store?

<details>
<summary>View Answer</summary>

**Solution: JWT Propagation**

```
┌─────────────────────────────────────────────────────────┐
│                    API Gateway                           │
│              (validates JWT once)                        │
└─────────────────────────────────────────────────────────┘
                          │
                          │ JWT in header
                          ↓
┌─────────────────────────────────────────────────────────┐
│                    Service A (Auth)                      │
│                 Issues JWT on login                      │
└─────────────────────────────────────────────────────────┘
                          │
         JWT passed in Authorization header
                          │
        ┌─────────────────┼─────────────────┐
        ↓                 ↓                 ↓
   [Service B]       [Service C]       [Service D]
   Verifies JWT      Verifies JWT      Verifies JWT
   locally           locally           locally
        │                 │                 │
        └─────────────────┼─────────────────┘
                          ↓
                    [Service E]
                    Verifies JWT locally
```

**Implementation:**

**Service A (Authentication Service):**
```javascript
// Login endpoint - issues JWT
app.post('/auth/login', async (req, res) => {
    const user = await validateCredentials(req.body)

    const token = jwt.sign({
        userId: user.id,
        email: user.email,
        roles: user.roles,
        // Include claims needed by other services
        tenantId: user.tenantId,
        permissions: user.permissions
    }, PRIVATE_KEY, {
        algorithm: 'RS256',  // Asymmetric for security
        expiresIn: '1h',
        issuer: 'auth-service'
    })

    res.json({ token })
})

// Public key endpoint (other services fetch this)
app.get('/auth/.well-known/jwks.json', (req, res) => {
    res.json({
        keys: [{ /* public key in JWK format */ }]
    })
})
```

**Services B, C, D, E (Resource Services):**
```javascript
// Shared middleware for all services
const jwksClient = require('jwks-rsa')

const client = jwksClient({
    jwksUri: 'http://auth-service/auth/.well-known/jwks.json',
    cache: true,            // Cache public keys
    cacheMaxAge: 86400000   // 24 hours
})

function getKey(header, callback) {
    client.getSigningKey(header.kid, (err, key) => {
        callback(null, key.publicKey || key.rsaPublicKey)
    })
}

const authMiddleware = (req, res, next) => {
    const token = req.headers.authorization?.split(' ')[1]
    if (!token) return res.status(401).send('No token')

    jwt.verify(token, getKey, {
        algorithms: ['RS256'],
        issuer: 'auth-service'
    }, (err, decoded) => {
        if (err) return res.status(401).send('Invalid token')

        // Attach user info to request
        req.user = decoded
        next()
    })
}

// Use in routes
app.get('/orders', authMiddleware, (req, res) => {
    // req.user contains userId, roles, etc.
    // No session store lookup!
    const orders = await getOrdersForUser(req.user.userId)
    res.json(orders)
})
```

**Service-to-Service Calls:**
```javascript
// Service B calling Service C
async function callServiceC(req) {
    // Forward the same JWT
    const response = await fetch('http://service-c/data', {
        headers: {
            'Authorization': req.headers.authorization  // Pass JWT through
        }
    })
    return response.json()
}
```

**Why No Session Store Needed:**

```
Traditional (with session store):
Request → Service → Redis lookup → User info
Every service hits Redis!

JWT approach:
Request → Service → Verify signature (CPU only) → User info
No external calls!

Benefits:
✅ No network hop to session store
✅ Services remain stateless
✅ Works across regions (no shared Redis needed)
✅ Scales infinitely
✅ Services can be in different clouds/networks
```

**Security Considerations:**

```
1. Use RS256 (asymmetric) not HS256:
   - Auth service: has private key (signs)
   - Other services: only have public key (verify)
   - Compromised service can't forge tokens

2. Include appropriate claims:
   {
     userId: 123,
     roles: ['user', 'admin'],
     tenantId: 'acme',       // For multi-tenant
     permissions: ['read:orders', 'write:orders']
   }

3. Set reasonable expiry:
   - Access token: 15 min - 1 hour
   - Refresh token: handled by auth service only

4. Validate all claims:
   - Check issuer
   - Check expiry
   - Check audience (if used)
```

</details>

**Q7:** A gaming company has stateful game servers (player positions in memory). How would you handle:
- Server crashes (player progress lost?)
- Scaling during peak hours
- Matching players to the right server

<details>
<summary>View Answer</summary>

**The Challenge:**

Game servers are inherently stateful - they must track real-time game state including player positions, health, inventory, etc. You can't make them fully stateless, but you can make them resilient.

**Architecture:**

```
                    [Matchmaking Service]
                    (stateless, assigns players to servers)
                              │
                    [Game Server Registry]
                    (Redis - which server has which game)
                              │
        ┌─────────────────────┼─────────────────────┐
        ↓                     ↓                     ↓
   [Game Server 1]      [Game Server 2]      [Game Server N]
   Games A, B, C        Games D, E, F        (auto-scaled)
        │                     │                     │
        └─────────────────────┼─────────────────────┘
                              │
                    [State Persistence Layer]
                    (Redis + PostgreSQL)
```

**Problem 1: Server Crashes**

```
Solution: Periodic State Snapshots + Event Sourcing

Game Server writes state to Redis every 1-2 seconds:
┌──────────────────────────────────────────────────────────┐
│                    Game Server                            │
│                                                          │
│   In-Memory State:                                       │
│   players: {                                             │
│     player_1: {x: 100, y: 200, health: 75, inventory: []}│
│     player_2: {x: 150, y: 180, health: 100}              │
│   }                                                      │
│                                                          │
│   Every 1 second:                                        │
│   └──→ Snapshot to Redis: game:123:state = {...}         │
│                                                          │
│   Every action:                                          │
│   └──→ Event to queue: {action: 'move', player: 1, ...}  │
│                                                          │
└──────────────────────────────────────────────────────────┘

On Crash Recovery:
1. Load last snapshot from Redis
2. Replay events since snapshot
3. Resume game with minimal data loss (1-2 seconds max)

Implementation:
async function snapshotState(gameId, state) {
    // Atomic snapshot
    await redis.set(`game:${gameId}:state`, JSON.stringify(state))
    await redis.set(`game:${gameId}:snapshot_time`, Date.now())
}

async function recoverGame(gameId) {
    // Load snapshot
    const snapshot = JSON.parse(await redis.get(`game:${gameId}:state`))
    const snapshotTime = await redis.get(`game:${gameId}:snapshot_time`)

    // Get events since snapshot
    const events = await redis.lrange(`game:${gameId}:events`, 0, -1)
    const recentEvents = events.filter(e => e.timestamp > snapshotTime)

    // Replay events
    let state = snapshot
    for (const event of recentEvents) {
        state = applyEvent(state, event)
    }

    return state
}
```

**Problem 2: Scaling During Peak Hours**

```
Solution: Auto-Scaling with Game Server Pool

Architecture:
┌─────────────────────────────────────────────────────────┐
│                 Game Server Orchestrator                 │
│                                                         │
│  Monitors:                                              │
│  - Active games per server                              │
│  - CPU/Memory usage                                     │
│  - Queue of waiting players                             │
│                                                         │
│  Scaling rules:                                         │
│  - If avg games/server > 10: scale up                   │
│  - If queue wait > 30sec: scale up                      │
│  - If avg games/server < 3: scale down                  │
│                                                         │
└─────────────────────────────────────────────────────────┘

Key insight: Don't scale servers with active games!

Scale-Down Process:
1. Mark server as "draining" (no new games)
2. Wait for current games to end (or migrate)
3. Terminate server

Game Migration (for long games):
1. Snapshot game state
2. Notify players: "Migrating to new server..."
3. Start game on new server with snapshot
4. Redirect player connections
5. Old server releases game

Implementation:
async function migrateGame(gameId, fromServer, toServer) {
    // 1. Pause game briefly
    await fromServer.pauseGame(gameId)

    // 2. Snapshot state
    const state = await fromServer.getState(gameId)

    // 3. Transfer to new server
    await toServer.loadGame(gameId, state)

    // 4. Update registry
    await redis.hset('game-servers', gameId, toServer.id)

    // 5. Notify players to reconnect
    await notifyPlayers(gameId, {
        type: 'server-change',
        newServer: toServer.address
    })

    // 6. Resume game
    await toServer.resumeGame(gameId)
}
```

**Problem 3: Matching Players to Right Server**

```
Solution: Registry + Smart Matchmaking

┌─────────────────────────────────────────────────────────┐
│                   Game Registry (Redis)                  │
│                                                         │
│  game:123 → server:game-server-5.example.com            │
│  game:456 → server:game-server-2.example.com            │
│  game:789 → server:game-server-5.example.com            │
│                                                         │
│  server:game-server-5:games → [123, 789]                │
│  server:game-server-5:capacity → 8/20 games             │
│  server:game-server-5:region → us-east                  │
│                                                         │
└─────────────────────────────────────────────────────────┘

Player Connection Flow:

1. Player wants to join game 123:
   Client → Matchmaker: "Join game 123"

2. Matchmaker looks up server:
   server = redis.hget('games', '123')
   // Returns: game-server-5.example.com

3. Return server to client:
   Matchmaker → Client: "Connect to game-server-5:7777"

4. Client connects directly to game server:
   Client ←→ Game Server 5 (WebSocket)

New Game Creation:
1. Matchmaker selects server with capacity + good latency
2. Creates game on server
3. Registers in Redis
4. Returns server address to players

Implementation:
async function joinGame(playerId, gameId) {
    // Check if game exists
    const serverAddress = await redis.hget('games', gameId)

    if (serverAddress) {
        // Existing game - return server address
        return { server: serverAddress, gameId }
    }

    // New game - find best server
    const bestServer = await findBestServer(playerId)

    // Create game on server
    await bestServer.createGame(gameId)

    // Register in Redis
    await redis.hset('games', gameId, bestServer.address)
    await redis.sadd(`server:${bestServer.id}:games`, gameId)

    return { server: bestServer.address, gameId }
}

async function findBestServer(playerId) {
    const playerRegion = await getPlayerRegion(playerId)
    const servers = await getServersInRegion(playerRegion)

    // Find server with:
    // 1. Available capacity
    // 2. Lowest latency to player
    // 3. Balanced load
    return servers
        .filter(s => s.capacity < s.maxCapacity)
        .sort((a, b) => a.currentGames - b.currentGames)[0]
}
```

**Complete Solution Summary:**

| Problem | Solution |
|---------|----------|
| Server crashes | Periodic snapshots to Redis + event replay |
| Lost progress | Max 1-2 seconds of data loss |
| Peak scaling | Auto-scale based on capacity/queue |
| Scale down | Drain mode + game migration |
| Find server | Game registry in Redis |
| Latency | Region-aware matchmaking |

</details>

---

## Next Up

Congratulations on completing Week 3: Growing Your System! You now understand vertical vs horizontal scaling, database replication & sharding, and stateless vs stateful architecture.

In Week 4, we'll start with **Consistent Hashing** - the elegant algorithm behind distributed caches, databases, and load balancers!
