# Week 2, Lesson 4: Networking Basics - DNS, Load Balancers, CDNs

## Table of Contents
- [Media Resources](#media-resources)
- [Part 1: DNS (Domain Name System)](#part-1-dns-domain-name-system)
  - [What is DNS?](#what-is-dns)
  - [How DNS Works: The Journey](#how-dns-works-the-journey)
  - [DNS Record Types](#dns-record-types)
  - [DNS in System Design](#dns-in-system-design)
  - [DNS Propagation](#dns-propagation)
- [Part 2: Load Balancers](#part-2-load-balancers)
  - [What is a Load Balancer?](#what-is-a-load-balancer)
  - [Benefits of Load Balancing](#benefits-of-load-balancing)
  - [Load Balancer Algorithms](#load-balancer-algorithms)
  - [Load Balancer Types](#load-balancer-types)
  - [Health Checks](#health-checks)
  - [Real-World Load Balancer Architecture](#real-world-load-balancer-architecture)
  - [Popular Load Balancers](#popular-load-balancers)
- [Part 3: CDN (Content Delivery Network)](#part-3-cdn-content-delivery-network)
  - [What is a CDN?](#what-is-a-cdn)
  - [How CDN Works](#how-cdn-works)
  - [CDN Architecture](#cdn-architecture)
  - [What CDNs Cache](#what-cdns-cache)
  - [Cache Control Headers](#cache-control-headers)
  - [CDN Cache Invalidation](#cdn-cache-invalidation)
  - [CDN Benefits](#cdn-benefits)
  - [Popular CDNs](#popular-cdns)
  - [CDN Use Cases](#cdn-use-cases)
- [Putting It All Together: Complete Request Flow](#putting-it-all-together-complete-request-flow)
- [Real-World Example: Netflix Architecture](#real-world-example-netflix-architecture)
- [Key Concepts to Remember](#key-concepts-to-remember)
- [System Design Impact](#system-design-impact)
- [Practice Questions](#practice-questions)
- [Next Up](#next-up)

---

Welcome to Week 2! Now we're getting into the infrastructure that makes the internet work at scale. Let's understand the invisible magic that happens when you type "google.com" into your browser.

## Media Resources

**Visual Guide:**
![The Pillars of Modern Networking: DNS, Load Balancers & CDNs](./assets/dns_loadbalancers_cdn_infographic.png)

**Audio Lesson:**
<audio controls src="./assets/dns_loadbalancers_cdn.m4a"></audio>

---

## Part 1: DNS (Domain Name System)

### What is DNS?

**DNS = The Internet's Phone Book**

Computers communicate using IP addresses (like `142.250.185.46`), but humans remember names (like `google.com`). DNS translates between them.

**The Problem:**
```
You type: www.instagram.com
Computer needs: 157.240.22.174
```

**The Solution: DNS**

### How DNS Works: The Journey

When you type `www.instagram.com`:

```
Step 1: Check Browser Cache
└─> "Have I visited instagram.com recently?"
    └─> If yes: Use cached IP (142.250.185.46)
    └─> If no: Continue...

Step 2: Check OS Cache
└─> Operating system checks its DNS cache
    └─> If found: Use it
    └─> If no: Continue...

Step 3: Query DNS Resolver (Your ISP)
└─> Your computer asks: "What's instagram.com's IP?"
    └─> Resolver checks its cache
    └─> If no: Start DNS recursion...

Step 4: Root DNS Server
└─> Resolver asks root server: "Where's .com?"
    └─> Root: "Ask the .com TLD server at 192.5.6.30"

Step 5: TLD (Top-Level Domain) Server
└─> Resolver asks TLD: "Where's instagram.com?"
    └─> TLD: "Ask Instagram's nameserver at 157.240.1.1"

Step 6: Authoritative Nameserver
└─> Resolver asks Instagram's nameserver: "What's www.instagram.com?"
    └─> Nameserver: "It's 157.240.22.174"

Step 7: Return to You
└─> Resolver caches the answer and returns it
    └─> Your browser connects to 157.240.22.174
```

**Typical Timeline:**
- Cache hit: 0-5ms
- Full DNS resolution: 20-120ms
- That's why caching is crucial!

### DNS Record Types

```
A Record (Address)
instagram.com → 157.240.22.174 (IPv4)

AAAA Record
instagram.com → 2a03:2880:f12f:83:face:b00c:0:25de (IPv6)

CNAME Record (Alias)
www.instagram.com → instagram.com
blog.instagram.com → instagram.com

MX Record (Mail Server)
instagram.com → mail.instagram.com (priority: 10)

TXT Record (Text/Verification)
instagram.com → "google-site-verification=abc123"

NS Record (Nameserver)
instagram.com → ns1.instagram.com, ns2.instagram.com
```

### DNS in System Design

**Key Concepts:**

**1. TTL (Time To Live)**
```
instagram.com    A    157.240.22.174    TTL: 300 seconds

After 300 seconds, cached record expires and needs refresh
```

- Short TTL (60s): Fast updates, more DNS queries
- Long TTL (86400s/1 day): Fewer queries, slow updates

**2. DNS Load Balancing**
```
Request 1: google.com → 142.250.185.46
Request 2: google.com → 172.217.14.206
Request 3: google.com → 142.250.186.78

One domain, multiple IPs! DNS rotates responses.
```

**3. Geographic DNS (Geo-DNS)**
```
User in USA requests netflix.com
└─> DNS returns: 54.88.208.4 (US server)

User in Europe requests netflix.com
└─> DNS returns: 52.48.122.11 (EU server)

Route users to nearest server!
```

**4. DNS Failover**
```
Primary server: 192.168.1.10 (healthy)
└─> Return this

Primary server: 192.168.1.10 (down!)
└─> Return backup: 192.168.1.11
```

### DNS Propagation

When you update DNS records:
```
1. Update nameserver: Instant
2. ISP caches expire: Minutes to hours (depends on TTL)
3. All users see new IP: 24-48 hours (worst case)
```

**Pro tip:** Lower TTL before making changes!

---

## Part 2: Load Balancers

### What is a Load Balancer?

**Load Balancer = Traffic Director**

Instead of one server handling all requests, distribute traffic across many servers.

```
Without Load Balancer:
User 1 ─┐
User 2 ─┼─> Single Server (overwhelmed!) 😵
User 3 ─┘

With Load Balancer:
User 1 ─┐           ┌─> Server 1 ✅
User 2 ─┼─> LB ────┼─> Server 2 ✅
User 3 ─┘           └─> Server 3 ✅
```

### Benefits of Load Balancing

**1. Scalability**
- Add more servers to handle more traffic
- Remove servers when traffic decreases

**2. High Availability**
- If Server 1 crashes, LB routes to Server 2 & 3
- No downtime!

**3. Performance**
- Distribute load evenly
- No single server is overwhelmed

**4. Maintenance**
- Take servers offline for updates without downtime

### Load Balancer Algorithms

#### 1. Round Robin (Most Common)
```
Request 1 → Server 1
Request 2 → Server 2
Request 3 → Server 3
Request 4 → Server 1 (cycle repeats)
```
**Pros:** Simple, fair
**Cons:** Doesn't consider server load

#### 2. Least Connections
```
Server 1: 10 active connections
Server 2: 5 active connections  ← Send here!
Server 3: 12 active connections

Route to server with fewest connections
```
**Pros:** Considers actual load
**Cons:** Slightly more complex

#### 3. Weighted Round Robin
```
Server 1 (powerful): Weight 3
Server 2 (medium): Weight 2
Server 3 (weak): Weight 1

Pattern: S1, S1, S1, S2, S2, S3, repeat...
```
**Pros:** Handles servers of different capacities
**Cons:** Need to configure weights

#### 4. IP Hash
```
User IP: 192.168.1.50
Hash(192.168.1.50) % 3 = 2 → Server 2

Same user always goes to same server!
```
**Pros:** Session affinity (user stays on same server)
**Cons:** Uneven distribution possible

#### 5. Least Response Time
```
Server 1: Avg response 50ms
Server 2: Avg response 30ms ← Send here!
Server 3: Avg response 80ms
```
**Pros:** Best user experience
**Cons:** Requires monitoring

### Load Balancer Types

#### Layer 4 (Transport Layer) - Fast
- Routes based on IP and TCP/UDP port
- Doesn't look at actual content
- Very fast but "dumb"

```
Any HTTP request to port 80 → Web servers
Any request to port 443 (HTTPS) → Web servers
Any request to port 3306 (MySQL) → Database servers
```

#### Layer 7 (Application Layer) - Smart
- Looks at HTTP headers, URLs, cookies
- Can make intelligent routing decisions
- Slightly slower but powerful

```
/api/* → API servers
/static/* → Static file servers
/admin/* → Admin servers

www.example.com → Web servers
api.example.com → API servers
```

### Health Checks

Load balancers constantly check if servers are healthy:

```
Every 5 seconds:
LB → Server 1: "GET /health"
Server 1 → LB: "200 OK" ✅

LB → Server 2: "GET /health"
Server 2 → LB: No response ❌ (mark as down)

LB → Server 3: "GET /health"
Server 3 → LB: "500 Error" ❌ (mark as down)

Only route traffic to Server 1!
```

### Real-World Load Balancer Architecture

```
                    Internet
                       │
                       ↓
              [Global Load Balancer]
                       │
        ┌──────────────┼──────────────┐
        ↓              ↓              ↓
  [US Data Center] [EU Data Center] [Asia Data Center]
        │              │              │
        ↓              ↓              ↓
   [Regional LB]   [Regional LB]   [Regional LB]
        │              │              │
    ┌───┴───┐      ┌───┴───┐      ┌───┴───┐
    ↓       ↓      ↓       ↓      ↓       ↓
  [Web1] [Web2]  [Web1] [Web2]  [Web1] [Web2]
```

### Popular Load Balancers

**Hardware:**
- F5 BIG-IP
- Citrix ADC

**Software:**
- Nginx (most popular)
- HAProxy
- AWS ELB/ALB
- Google Cloud Load Balancer

---

## Part 3: CDN (Content Delivery Network)

### What is a CDN?

**CDN = Global Cache of Your Content**

Instead of everyone downloading from your origin server, copy content to servers around the world.

```
Without CDN:
User in Tokyo → Requests image from USA server → 200ms latency 😢

With CDN:
User in Tokyo → Requests image from Tokyo CDN → 10ms latency 😊
```

### How CDN Works

```
Step 1: User requests www.example.com/logo.png

Step 2: DNS returns CDN IP (nearest edge server)

Step 3: CDN checks: "Do I have logo.png cached?"

If YES (Cache HIT):
└─> Return cached file (super fast!)

If NO (Cache MISS):
└─> CDN fetches from origin server
└─> Stores in cache
└─> Returns to user
└─> Next user gets cache HIT!
```

### CDN Architecture

```
                    [Origin Server]
                    (Your server)
                          │
                          │ Pull content
                          ↓
        ┌─────────────────┼─────────────────┐
        │                 │                 │
        ↓                 ↓                 ↓
  [CDN - USA]      [CDN - Europe]    [CDN - Asia]
   (Edge Server)    (Edge Server)     (Edge Server)
        │                 │                 │
    ┌───┴───┐         ┌───┴───┐         ┌───┴───┐
    ↓       ↓         ↓       ↓         ↓       ↓
  Users   Users     Users   Users     Users   Users
```

### What CDNs Cache

**Static Content (Perfect for CDN):**
- Images, CSS, JavaScript
- Videos, audio files
- PDFs, documents
- Fonts

**Dynamic Content (Harder to cache):**
- Personalized feeds
- Real-time data
- User-specific content

**Modern CDNs can even cache dynamic content with smart strategies!**

### Cache Control Headers

Servers tell CDNs how to cache:

```http
Response Headers:

Cache-Control: public, max-age=86400
└─> Cache for 24 hours, anyone can cache

Cache-Control: private, max-age=3600
└─> Cache for 1 hour, only user's browser can cache

Cache-Control: no-cache
└─> Must revalidate with server before using cache

Cache-Control: no-store
└─> Don't cache at all (sensitive data)

ETag: "abc123xyz789"
└─> Version identifier (for cache validation)
```

### CDN Cache Invalidation

**Problem:** You updated logo.png but CDN still serves old version!

**Solutions:**

**1. Wait for TTL to expire**
```
Cache-Control: max-age=3600
Wait 1 hour... 😴
```

**2. Cache Purge (Manual)**
```
Tell CDN: "Delete /logo.png from all servers"
Next request triggers cache MISS and refetches
```

**3. Versioned URLs (Best Practice)**
```
Old: /logo.png
New: /logo.png?v=2
Or:  /logo-v2.png

URL changed = CDN treats as new file!
```

### CDN Benefits

**1. Performance**
```
Without CDN:
London user → USA server → 150ms

With CDN:
London user → London CDN → 15ms (10x faster!)
```

**2. Reduced Origin Load**
```
1000 users request logo.png

Without CDN:
└─> 1000 requests to your server 😰

With CDN:
└─> 1 request to your server (first cache MISS)
└─> 999 requests served from CDN 😊
```

**3. DDoS Protection**
```
Attack: 1 million requests/second

CDN absorbs and filters:
└─> Only legitimate traffic reaches your server
```

**4. High Availability**
```
Your origin server crashes!

CDN still serves cached content
└─> Users don't notice (for cached content)
```

### Popular CDNs

- **Cloudflare** - Free tier, DDoS protection
- **AWS CloudFront** - Integrates with AWS
- **Fastly** - Real-time purging
- **Akamai** - Enterprise, largest network
- **Google Cloud CDN**
- **Azure CDN**

### CDN Use Cases

**Perfect for:**
- E-commerce (product images)
- Media sites (videos, images)
- Software downloads
- Global applications
- High-traffic sites

**Not needed for:**
- Internal applications
- Very low traffic sites
- Mostly dynamic content
- Local-only services

---

## Putting It All Together: Complete Request Flow

Let's see how DNS, Load Balancers, and CDN work together:

```
User types: www.instagram.com

STEP 1: DNS Resolution
└─> DNS: "instagram.com = 151.101.1.15" (CDN IP)

STEP 2: Request hits CDN
└─> User → CDN edge server (London)
└─> CDN checks cache for /feed

STEP 3a: Cache HIT (Static content)
└─> Return cached HTML/CSS/JS/images
└─> Done! (Super fast!)

STEP 3b: Cache MISS (Dynamic content)
└─> CDN → Global Load Balancer
└─> Global LB → Europe Regional Load Balancer
└─> Regional LB → Web Server (least connections)

STEP 4: Web Server processes request
└─> Fetches from database
└─> Generates personalized feed
└─> Returns to CDN
└─> CDN caches (if appropriate) & returns to user

STEP 5: Subsequent requests
└─> Cached content served instantly from CDN!
```

---

## Real-World Example: Netflix Architecture

```
                    [User]
                       │
                   [DNS Query]
                       │
              Returns: CDN IP (nearest)
                       │
                       ↓
              [CDN (Cloudfront)]
           ┌───────────┴───────────┐
           │                       │
    [Cache HIT]              [Cache MISS]
     (video)                       │
           │                       ↓
           │              [Load Balancer]
           │                       │
           │            ┌──────────┼──────────┐
           │            ↓          ↓          ↓
           │        [Server1]  [Server2]  [Server3]
           │            │
           │            ↓
           │      [Database Query]
           │      (user preferences,
           │       video metadata)
           │
           └────> [Return Video Stream]
```

---

## Key Concepts to Remember

### DNS
1. Translates domain names to IP addresses
2. Hierarchical system (Root → TLD → Authoritative)
3. Caching at multiple levels dramatically improves speed
4. Can be used for load balancing and geographic routing
5. TTL controls how long records are cached

### Load Balancers
1. Distribute traffic across multiple servers
2. Provide high availability and scalability
3. Can be Layer 4 (fast) or Layer 7 (smart)
4. Use health checks to detect failed servers
5. Multiple algorithms for different use cases

### CDN
1. Cache content at edge servers globally
2. Dramatically reduces latency for users
3. Reduces load on origin servers
4. Provides DDoS protection
5. Best for static content, but can cache dynamic too
6. Cache invalidation is a key challenge

---

## System Design Impact

When designing systems, always consider:

**1. Global users?**
→ Use CDN + Geo-DNS

**2. High traffic?**
→ Use Load Balancer + multiple servers

**3. Need high availability?**
→ Use Load Balancer + health checks

**4. Static content?**
→ Always use CDN

**5. Complex routing?**
→ Use Layer 7 Load Balancer

---

## Practice Questions

**Q1:** A user in Australia visits your US-based website. Walk through the complete request flow involving DNS, Load Balancers, and CDN. Where are the main latency points?

<details>
<summary>View Answer</summary>

**Request Flow:**

1. **DNS Resolution (~50-200ms)**
   - Browser checks local cache → ISP DNS → Root DNS → .com TLD → Your authoritative DNS
   - GeoDNS returns IP of nearest CDN edge (Sydney)

2. **CDN Edge (Sydney) (~10-50ms)**
   - Request hits Sydney CDN edge
   - If cached: Return immediately (fast!)
   - If not cached: Forward to origin

3. **Load Balancer (~5-10ms if regional)**
   - Request routed to US load balancer
   - LB selects healthy backend server

4. **Origin Server (~100-300ms round trip)**
   - US server processes request
   - Response travels back to Australia

5. **Response cached at CDN edge**
   - Future Australian users get fast response

**Main Latency Points:**
- **DNS lookup:** First visit is slow (mitigate with DNS prefetch)
- **Origin fetch:** 300ms+ round trip to US (mitigate with CDN caching)
- **TLS handshake:** Adds 1-2 round trips (mitigate with session resumption)

</details>

**Q2:** Your company is launching a video streaming service globally. Design the infrastructure using DNS, Load Balancers, and CDN. What challenges might you face?

<details>
<summary>View Answer</summary>

**Infrastructure Design:**

```
                    [GeoDNS]
                       │
    ┌──────────────────┼──────────────────┐
    ↓                  ↓                  ↓
[CDN Edge US]    [CDN Edge EU]    [CDN Edge Asia]
    │                  │                  │
    ↓                  ↓                  ↓
[Regional LB]    [Regional LB]    [Regional LB]
    │                  │                  │
[Origin US]      [Origin EU]      [Origin Asia]
```

**Key Components:**
- **GeoDNS:** Route users to nearest region
- **CDN:** Cache video chunks at edge (huge bandwidth savings)
- **Regional origins:** Reduce latency for uploads and API calls
- **Adaptive bitrate:** Serve different quality based on connection

**Challenges:**

1. **Video file size:** Terabytes of content, expensive CDN costs
2. **Live streaming:** Can't cache live content, need low-latency protocols
3. **Cold start:** First viewer in a region gets slow experience
4. **Cache invalidation:** When videos are updated or removed
5. **DRM/geo-restrictions:** Different content rights per country
6. **Peak load:** Viral content causes traffic spikes
7. **Cost:** CDN bandwidth for video is expensive

</details>

**Q3:** You have these servers:
- Server A: 16GB RAM, handling 100 connections
- Server B: 32GB RAM, handling 50 connections
- Server C: 8GB RAM, handling 150 connections

Which load balancing algorithm would you use? Why?

<details>
<summary>View Answer</summary>

**Least Connections** is the best choice here.

**Why:**
- Servers have different capacities AND different current loads
- Round-robin would send equal traffic, ignoring current load
- Weighted would require manual configuration of weights

**Current state analysis:**
- Server A: 100 connections (moderate load)
- Server B: 50 connections (underutilized despite more RAM!)
- Server C: 150 connections (overloaded for 8GB)

**With Least Connections:**
- Next request → Server B (only 50 connections)
- This naturally balances based on actual load

**Even better: Weighted Least Connections**
- Weight by RAM: A=2, B=4, C=1
- Considers both capacity and current load
- Server B gets more traffic (more RAM, fewer connections)

</details>

**Q4:** Your website has:
- Homepage (changes hourly)
- Product images (rarely change)
- User profiles (personalized)
- Blog posts (change weekly)

What caching strategy (TTL values) would you use for each?

<details>
<summary>View Answer</summary>

| Content | TTL | Cache-Control Header | Reason |
|---------|-----|---------------------|--------|
| **Homepage** | 5-15 minutes | `public, max-age=900, stale-while-revalidate=300` | Changes hourly, short TTL with stale-while-revalidate for smooth updates |
| **Product images** | 1 year | `public, max-age=31536000, immutable` | Rarely change; use versioned URLs (/img/product.v2.jpg) |
| **User profiles** | 0 (no cache) | `private, no-store` | Personalized content should never be cached on CDN |
| **Blog posts** | 1 day | `public, max-age=86400, stale-while-revalidate=3600` | Change weekly, 1-day TTL is safe |

**Additional strategies:**
- Use `stale-while-revalidate` for better UX during revalidation
- Version static assets (CSS, JS, images) for instant updates
- Set `private` for any user-specific content
- Consider `s-maxage` for CDN-specific TTL different from browser

</details>

**Q5:** Your origin server goes down for 2 hours. You have a CDN with 24-hour cache. What happens to:
- Cached content?
- Non-cached content?
- Users who last visited 1 hour ago vs 25 hours ago?

<details>
<summary>View Answer</summary>

**Cached content:**
- ✅ Continues to be served from CDN edge
- Users see no disruption for cached pages
- CDN serves "stale" content (which is fine for 2 hours)

**Non-cached content:**
- ❌ CDN tries to fetch from origin, fails
- Users see error (502 Bad Gateway or 504 Timeout)
- Any uncached pages, API calls, or dynamic content fails

**Users who visited 1 hour ago:**
- ✅ Their content is still in CDN cache (TTL: 24h - 1h = 23h remaining)
- They can continue browsing cached content normally

**Users who visited 25 hours ago:**
- ❌ Their cached content expired (24h TTL exceeded)
- CDN needs to revalidate with origin
- Origin is down → Error or stale content (if `stale-if-error` is configured)

**Best practices for resilience:**
- Configure `stale-if-error` directive to serve expired cache when origin fails
- Set up health checks and automatic failover
- Use multiple origin servers
- Cache API responses where possible

</details>

**Q6:** Explain why "versioned URLs" (e.g., /style.css?v=2) is better than cache purging for CDN updates.

<details>
<summary>View Answer</summary>

**Versioned URLs advantages:**

1. **Instant update:** New URL = new cache entry, immediately fetched
   - Old: `/style.css?v=1` (still cached)
   - New: `/style.css?v=2` (fetched immediately)

2. **No purge propagation delay:** Cache purges can take minutes to propagate across all CDN edges globally

3. **Atomic deployment:** Users get all-new or all-old assets, never a mix

4. **Rollback is trivial:** Just change back to `?v=1`

5. **No purge API needed:** Works with any CDN, no special configuration

6. **Long cache TTL:** Can set 1-year cache since URL changes when content changes

**Cache purging problems:**

1. **Propagation delay:** 1-15 minutes for global purge
2. **Race conditions:** Some users get new HTML with old CSS
3. **Cost:** Many CDNs charge for purge API calls
4. **Complexity:** Need to track what to purge
5. **Inconsistency:** Some edges purged, others not

**Best practice:** Use content-hash in filename:
```
/style.a1b2c3d4.css  (hash of file content)
```
Build tools (Webpack, Vite) do this automatically!

</details>

---

## Next Up

In Lesson 5, we'll explore **Caching Fundamentals** - the secret weapon behind fast, scalable systems!
