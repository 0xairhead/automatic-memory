# System Design Mastery
## 12-Week Course Curriculum

### **Phase 1: Fundamentals (Weeks 1-2)**

#### **Week 1: Building Blocks**
- **Lesson 1**: Client-Server Architecture & HTTP Basics
  - Understanding the request-response model
  - HTTP methods, status codes, and headers
  - RESTful principles
  
- **Lesson 2**: Databases 101 - SQL vs NoSQL
  - Relational databases and ACID properties
  - NoSQL databases and when to use them
  - Trade-offs and selection criteria
  
- **Lesson 3**: APIs - REST, Design Principles, Pagination
  - RESTful API design best practices
  - API versioning strategies
  - Pagination, filtering, and sorting
  - Rate limiting basics

#### **Week 2: Core Infrastructure**
- **Lesson 4**: Networking Basics - DNS, Load Balancers, CDNs
  - How DNS resolution works
  - Load balancing strategies
  - Content delivery networks
  - Geographic distribution
  
- **Lesson 5**: Caching Fundamentals - Why, Where, and How
  - Cache invalidation strategies
  - Cache eviction policies (LRU, LFU)
  - Different caching layers
  - Cache-aside vs write-through patterns
  
- **Lesson 6**: Storage Systems - Files, Blocks, Objects
  - File storage vs block storage vs object storage
  - When to use each type
  - Distributed file systems
  - Blob storage concepts

---

### **Phase 2: Scalability (Weeks 3-4)**

#### **Week 3: Growing Your System**
- **Lesson 7**: Vertical vs Horizontal Scaling
  - Scale-up vs scale-out strategies
  - Trade-offs and limitations
  - When to choose each approach
  
- **Lesson 8**: Database Replication & Sharding
  - Master-slave replication
  - Multi-master replication
  - Horizontal partitioning (sharding)
  - Sharding strategies and challenges
  
- **Lesson 9**: Stateless vs Stateful Architecture
  - Benefits of stateless design
  - Managing session state
  - Sticky sessions vs session stores

#### **Week 4: Distribution Basics**
- **Lesson 10**: Consistent Hashing
  - The problem with traditional hashing
  - How consistent hashing works
  - Virtual nodes and load distribution
  
- **Lesson 11**: Message Queues & Async Processing
  - Synchronous vs asynchronous processing
  - Message queue patterns
  - Popular queue systems
  - Use cases and benefits
  
- **Lesson 12**: CAP Theorem Explained Simply
  - Consistency, Availability, Partition Tolerance
  - Real-world trade-offs
  - CP vs AP systems

---

### **Phase 3: System Design Patterns (Weeks 5-6)**

#### **Week 5: Common Patterns**
- **Lesson 13**: Microservices vs Monoliths
  - Monolithic architecture pros and cons
  - Microservices benefits and challenges
  - Service boundaries and communication
  - When to choose each approach
  
- **Lesson 14**: Event-Driven Architecture
  - Event sourcing concepts
  - Publish-subscribe patterns
  - Event streaming platforms
  - Benefits and use cases
  
- **Lesson 15**: Rate Limiting & Throttling
  - Why rate limiting matters
  - Token bucket algorithm
  - Leaky bucket algorithm
  - Distributed rate limiting

#### **Week 6: Reliability**
- **Lesson 16**: Fault Tolerance & Circuit Breakers
  - Designing for failure
  - Circuit breaker pattern
  - Retry strategies and backoff
  - Bulkhead pattern
  
- **Lesson 17**: Monitoring & Observability
  - Metrics, logs, and traces
  - Key performance indicators
  - Alerting strategies
  - Debugging distributed systems
  
- **Lesson 18**: Backup & Disaster Recovery
  - Backup strategies (full, incremental, differential)
  - Recovery time objective (RTO)
  - Recovery point objective (RPO)
  - Multi-region deployments

---

### **Phase 4: Real System Designs (Weeks 7-10)**

Each week, design 2-3 complete systems together with step-by-step guidance:

#### **Week 7: Basic Systems**
- URL Shortener (e.g., bit.ly)
- Pastebin
- Web Crawler

#### **Week 8: Social Media**
- Twitter-like service
- Instagram photo sharing
- News feed system

#### **Week 9: Media & Streaming**
- YouTube video platform
- Spotify music streaming
- Netflix recommendation system

#### **Week 10: Real-time & Location**
- Uber/Lyft ride-sharing
- WhatsApp messaging
- Real-time collaboration (Google Docs)

**For each system, you will:**
1. Gather requirements and constraints
2. Estimate capacity and scale
3. Design APIs and data models
4. Create high-level architecture
5. Address bottlenecks and trade-offs
6. Discuss monitoring and maintenance

---

### **Phase 5: Advanced Topics (Weeks 11-12)**

#### **Week 11: Advanced Concepts**
- **Lesson 19**: Distributed Consensus
  - Paxos and Raft algorithms
  - Leader election
  - Distributed locks
  
- **Lesson 20**: Search Systems
  - Inverted indexes
  - Search ranking algorithms
  - Elasticsearch and search architecture
  - Auto-complete systems
  
- **Lesson 21**: Real-time Systems
  - WebSocket connections
  - Server-sent events
  - Long polling vs streaming
  - Real-time data pipelines

#### **Week 12: Production Readiness**
- **Lesson 22**: Security in System Design
  - Authentication vs authorization
  - API security best practices
  - DDoS protection
  - Data encryption (at rest and in transit)
  - HTTPS/TLS essentials


