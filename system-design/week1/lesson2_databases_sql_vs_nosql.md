# Week 1, Lesson 2: Databases 101 - SQL vs NoSQL

## Table of Contents
- [Media Resources](#media-resources)
- [What is a Database?](#what-is-a-database)
- [The Two Main Types: SQL vs NoSQL](#the-two-main-types-sql-vs-nosql)
  - [SQL Databases (Relational Databases)](#sql-databases-relational-databases)
  - [NoSQL Databases (Non-Relational)](#nosql-databases-non-relational)
- [SQL vs NoSQL: The Big Comparison](#sql-vs-nosql-the-big-comparison)
- [Real-World Decision Making](#real-world-decision-making)
- [Understanding "Eventual Consistency"](#understanding-eventual-consistency)
- [When to Choose What?](#when-to-choose-what)
- [Hybrid Approach: The Modern Way](#hybrid-approach-the-modern-way)
- [Key Concepts to Remember](#key-concepts-to-remember)
- [Practice Questions](#practice-questions)
- [Next Up](#next-up)

---

## Media Resources

**Visual Guide:**
![SQL vs NoSQL: Choosing the Right Database](./assets/sql_vs_nosql_infographic.png)

**Audio Lesson:**
[SQL, NoSQL, Consistency and Scale (Audio)](./assets/sql_nosql_consistency_and_scale.m4a)

---

Welcome to Lesson 2! Now that you understand how clients and servers communicate, let's talk about where servers store all that data.

## What is a Database?

A database is an organized collection of data that can be easily accessed, managed, and updated.

Think of it like this:
- **Excel spreadsheet** = Simple database
- **Library catalog system** = More complex database
- **Amazon's product catalog** = Massive, sophisticated database

Without databases, every time you turned off Instagram, all your photos would be gone! Databases persist (save) data permanently.

---

## The Two Main Types: SQL vs NoSQL

### SQL Databases (Relational Databases)

**SQL = Structured Query Language**

Think of SQL databases like **Excel spreadsheets with multiple related sheets**.

**Key characteristics:**
- Data is stored in **tables** (rows and columns)
- Tables have relationships with each other
- **Schema** = predefined structure (you must define columns before adding data)
- ACID compliant (we'll explain this)

**Popular SQL databases:**
- PostgreSQL
- MySQL
- Oracle
- Microsoft SQL Server
- SQLite

#### Example: A Social Media App

**Users Table:**
```
| user_id | username | email              | created_at |
|---------|----------|--------------------|------------|
| 1       | alice    | alice@email.com    | 2024-01-15 |
| 2       | bob      | bob@email.com      | 2024-02-20 |
| 3       | charlie  | charlie@email.com  | 2024-03-10 |
```

**Posts Table:**
```
| post_id | user_id | content                | likes | created_at |
|---------|---------|------------------------|-------|------------|
| 101     | 1       | "Hello world!"         | 25    | 2024-01-16 |
| 102     | 2       | "Learning databases!"  | 10    | 2024-02-21 |
| 103     | 1       | "SQL is cool"          | 40    | 2024-01-18 |
```

**Comments Table:**
```
| comment_id | post_id | user_id | comment          |
|------------|---------|---------|------------------|
| 501        | 101     | 2       | "Welcome!"       |
| 502        | 101     | 3       | "Nice post!"     |
| 503        | 102     | 1       | "Keep learning!" |
```

Notice how `user_id` in the Posts table **relates** to `user_id` in the Users table? That's why they're called **relational** databases!

#### How You Query SQL Databases

```sql
-- Get all posts by alice
SELECT * FROM posts 
WHERE user_id = 1;

-- Get all comments on a specific post with user info
SELECT comments.comment, users.username 
FROM comments 
JOIN users ON comments.user_id = users.user_id
WHERE comments.post_id = 101;

-- Count total posts per user
SELECT users.username, COUNT(posts.post_id) as post_count
FROM users
LEFT JOIN posts ON users.user_id = posts.user_id
GROUP BY users.username;
```

#### ACID Properties (Why SQL is Reliable)

**A - Atomicity:** All or nothing. If you're transferring $100 from Account A to Account B, either both operations succeed (deduct from A, add to B) or both fail. Never just one.

**C - Consistency:** Data follows all rules. If you defined that email must be unique, the database enforces this.

**I - Isolation:** Multiple operations don't interfere with each other. If two people try to buy the last concert ticket simultaneously, only one succeeds.

**D - Durability:** Once data is saved, it's permanent (even if server crashes).

**When to use SQL:**
- You need complex relationships (users → posts → comments → likes)
- Data structure is predictable and stable
- You need strong consistency (banking, e-commerce orders)
- You need complex queries and reporting
- Examples: Banking systems, e-commerce, HR systems

---

### NoSQL Databases (Non-Relational)

**NoSQL = "Not Only SQL"** or "Non-SQL"

Think of NoSQL databases like **flexible JSON documents or key-value pairs**.

**Key characteristics:**
- Flexible schema (can change structure anytime)
- Data stored in various formats (documents, key-value, graphs, columns)
- Built for scale and speed
- Eventually consistent (not always ACID)

**Types of NoSQL databases:**

#### 1. Document Databases (Most Popular)
**Examples:** MongoDB, CouchDB

Data stored as JSON-like documents:

```json
// User document
{
  "_id": "user_001",
  "username": "alice",
  "email": "alice@email.com",
  "profile": {
    "bio": "Love coding!",
    "location": "San Francisco",
    "interests": ["coding", "hiking", "photography"]
  },
  "posts": [
    {
      "post_id": "post_101",
      "content": "Hello world!",
      "likes": 25,
      "comments": [
        {
          "user": "bob",
          "text": "Welcome!",
          "timestamp": "2024-01-16T10:30:00Z"
        }
      ]
    }
  ],
  "created_at": "2024-01-15T09:00:00Z"
}
```

Notice: Everything about the user (including posts and comments) can be in ONE document! No joins needed.

#### 2. Key-Value Databases
**Examples:** Redis, DynamoDB

Like a giant hash map/dictionary:

```
Key: "user:1:session"          → Value: "abc123xyz789"
Key: "product:5:price"         → Value: "29.99"
Key: "cache:homepage:user:1"   → Value: "<html>...</html>"
```

Super fast lookups! Used heavily for caching.

#### 3. Column-Family Databases
**Examples:** Cassandra, HBase

Optimized for reading/writing columns instead of rows:

```
Row Key: user_001
  | Columns: username="alice", email="alice@email.com"

Row Key: user_002
  | Columns: username="bob", email="bob@email.com", age=25
```

Great for analytics and time-series data.

#### 4. Graph Databases
**Examples:** Neo4j, Amazon Neptune

Perfect for relationships and connections:

```
(Alice)-[:FOLLOWS]->(Bob)
(Alice)-[:LIKES]->(Post1)
(Bob)-[:COMMENTED_ON]->(Post1)
(Charlie)-[:FOLLOWS]->(Alice)
```

Perfect for social networks, recommendation engines.

---

## SQL vs NoSQL: The Big Comparison

| Feature | SQL | NoSQL |
|---------|-----|-------|
| **Schema** | Fixed (must define upfront) | Flexible (change anytime) |
| **Data Model** | Tables with relationships | Documents, key-value, graph, etc. |
| **Scalability** | Vertical (bigger server) | Horizontal (more servers) |
| **Consistency** | Strong (ACID) | Eventual (BASE) |
| **Query Language** | SQL (standardized) | Database-specific APIs |
| **Best For** | Complex queries, transactions | Fast reads/writes, flexibility |
| **Examples** | Banking, inventory, CRM | Social media, real-time analytics, IoT |

---

## Real-World Decision Making

### Scenario 1: E-commerce Platform

**Should you use SQL or NoSQL?**

**Answer: SQL (PostgreSQL or MySQL)**

**Why?**
- Need ACID transactions (when someone buys something, inventory must decrease atomically)
- Complex relationships: Users → Orders → Products → Reviews
- Need complex queries: "Show me all orders from California in December with products over $50"
- Data structure is stable and predictable

### Scenario 2: Social Media Feed (Like Twitter)

**Should you use SQL or NoSQL?**

**Answer: NoSQL (MongoDB or Cassandra)**

**Why?**
- Huge volume of data (millions of tweets)
- Need to scale horizontally across many servers
- Schema flexibility (different tweet types: text, images, videos, polls)
- Fast writes and reads are critical
- Can tolerate eventual consistency (if a like count is off by 1 for a second, it's okay)

### Scenario 3: Actually... Both!

Many big companies use **BOTH**:

**Example: Netflix**
- **SQL (MySQL)**: Billing, subscriptions, payment processing
- **NoSQL (Cassandra)**: Viewing history, recommendations, user preferences

This is called **Polyglot Persistence** - using the right database for each job!

---

## Understanding "Eventual Consistency"

**SQL = Strong Consistency:**
```
You update your profile picture.
IMMEDIATELY, everyone sees the new picture.
Guaranteed.
```

**NoSQL = Eventual Consistency:**
```
You update your profile picture.
Your friend in New York sees it immediately.
Your friend in Tokyo sees it 2 seconds later.
Your friend in London sees it 5 seconds later.
Eventually, everyone sees the same thing.
```

**Why eventual consistency?**
- Faster! Don't have to wait for all servers to sync
- More available! System keeps working even if some servers are down
- Scales better! Can spread across the globe

---

## When to Choose What?

### Choose SQL when:
✅ Data has clear structure and relationships  
✅ Need complex queries and joins  
✅ Need ACID guarantees (banking, inventory)  
✅ Data integrity is critical  
✅ Application is not write-heavy  

### Choose NoSQL when:
✅ Massive scale (millions of users)  
✅ Need fast reads/writes  
✅ Schema might change frequently  
✅ Can tolerate eventual consistency  
✅ Data is hierarchical or nested  
✅ Horizontal scaling is needed  

---

## Hybrid Approach: The Modern Way

Smart companies use both:

```
User Service (SQL) ──┐
                     ├─> API Gateway
Session Cache        │   (serves users)
(Redis/NoSQL) ──────┤
                     │
Product Catalog      │
(MongoDB/NoSQL) ────┤
                     │
Transaction DB      ─┘
(PostgreSQL/SQL)
```

---

## Key Concepts to Remember

1. **SQL = Structured, relational, ACID, complex queries**
2. **NoSQL = Flexible, scalable, fast, eventual consistency**
3. **There's no "best" database - only best for your use case**
4. **Many apps use multiple databases (polyglot persistence)**
5. **Trade-offs always exist: consistency vs availability vs performance**

---

## Practice Questions

**Q1:** You're building a banking app where users can transfer money between accounts. Should you use SQL or NoSQL? Why?

**Q2:** You're building a chat app like WhatsApp that needs to handle millions of messages per second. Should you use SQL or NoSQL? Why?

**Q3:** What problems might occur if you use a NoSQL database with eventual consistency for an e-commerce checkout system?

**Q4:** Look at this data structure. Would SQL or NoSQL be better? Why?
```
User:
  - name
  - email
  - orders: [
      {product: "iPhone", price: 999, date: "2024-01-15"},
      {product: "Case", price: 29, date: "2024-01-15"}
    ]
  - preferences: {theme: "dark", notifications: true}
```

**Q5:** A startup is building a new app. They're unsure if their data model will change. They expect 100 users initially but hope to scale to millions. What database approach would you recommend?

---

## Next Up

In Lesson 3, we'll dive into **APIs, REST, and Pagination** - learning how to design clean, scalable APIs that developers love to use!
