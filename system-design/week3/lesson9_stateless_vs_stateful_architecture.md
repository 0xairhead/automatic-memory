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

**Q2:** Your API uses JWT for authentication. A user reports their account was compromised. How do you immediately revoke their access? What are the trade-offs of your solution?

**Q3:** Compare these session strategies for a chat application with 1 million concurrent users:
- Sticky sessions
- Redis session store
- JWT tokens

Which would you choose and why?

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

**Q5:** Design a stateless file upload system that:
- Accepts files up to 100MB
- Allows users to list their uploads
- Works across 10 API servers

**Q6:** Your microservices architecture has 5 services. Users authenticate with Service A, then call Services B, C, D, E. How do you propagate authentication without each service hitting a session store?

**Q7:** A gaming company has stateful game servers (player positions in memory). How would you handle:
- Server crashes (player progress lost?)
- Scaling during peak hours
- Matching players to the right server

---

## Next Up

Congratulations on completing Week 3: Growing Your System! You now understand vertical vs horizontal scaling, database replication & sharding, and stateless vs stateful architecture.

In Week 4, we'll start with **Consistent Hashing** - the elegant algorithm behind distributed caches, databases, and load balancers!
