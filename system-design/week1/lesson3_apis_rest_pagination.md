# Week 1, Lesson 3: APIs - REST, Design Principles, and Pagination

## Table of Contents
- [Media Resources](#media-resources)
- [What is an API?](#what-is-an-api)
- [What is REST?](#what-is-rest)
  - [The 6 REST Principles](#the-6-rest-principles)
- [RESTful API Design: Best Practices](#restful-api-design-best-practices)
- [Pagination: Handling Large Datasets](#pagination-handling-large-datasets)
  - [Method 1: Offset-Based Pagination](#method-1-offset-based-pagination)
  - [Method 2: Cursor-Based Pagination](#method-2-cursor-based-pagination-better-for-feeds)
  - [Method 3: Keyset Pagination](#method-3-keyset-pagination-most-efficient)
- [Real-World Example: Twitter API](#real-world-example-twitter-api)
- [API Design Best Practices Summary](#api-design-best-practices-summary)
- [Rate Limiting](#rate-limiting)
- [API Authentication](#api-authentication)
- [Key Concepts to Remember](#key-concepts-to-remember)
- [Practice Questions](#practice-questions)
- [Next Up](#next-up)

---

## Media Resources

**Visual Guide:**
![A Developer's Guide to REST APIs](./assets/rest_apis_guide_infographic.png)

**Audio Lesson:**
[Designing REST APIs for Massive Scale (Audio)](./assets/designing_rest_apis_for_scale.m4a)

---

Welcome to Lesson 3! You've learned about client-server communication and databases. Now let's talk about APIs—the contracts that define how clients and servers talk to each other.

## What is an API?

**API = Application Programming Interface**

Think of an API as a **menu at a restaurant**:
- The menu tells you what you can order (available endpoints)
- You don't need to know how the kitchen works (internal implementation)
- You just need to know how to order (request format)
- The waiter brings your food (response)

**Real example:**
When you use the Instagram app, it doesn't directly access Instagram's database. Instead, it uses Instagram's API:
- `GET /api/users/alice` → Get Alice's profile
- `POST /api/posts` → Create a new post
- `GET /api/feed` → Get your personalized feed

---

## What is REST?

**REST = Representational State Transfer**

REST is a set of architectural principles for designing APIs. It's not a protocol or standard—it's a style.

### The 6 REST Principles

#### 1. Client-Server Architecture
- Client and server are separate
- They can evolve independently
- Client doesn't need to know database details

#### 2. Stateless
- Each request contains ALL information needed
- Server doesn't remember previous requests
- No session stored on server

**Example:**
```
❌ BAD (Stateful):
Request 1: "Login as Alice"
Request 2: "Get my posts" (server remembers you're Alice)

✅ GOOD (Stateless):
Request 1: "Login as Alice" → Get token: abc123
Request 2: "Get posts for token abc123" (includes who you are)
```

#### 3. Cacheable
- Responses should say if they can be cached
- Reduces server load
- Improves performance

#### 4. Uniform Interface
- Use standard HTTP methods
- Use standard status codes
- Use consistent URL structure

#### 5. Layered System
- Client doesn't know if connected directly to server or through proxy/load balancer
- Improves scalability and security

#### 6. Code on Demand (Optional)
- Server can send executable code (like JavaScript)
- Rarely used in practice

---

## RESTful API Design: Best Practices

### 1. Use Nouns for Resources, Not Verbs

**Resources** are things (users, posts, comments). Use HTTP methods to describe actions.

```
❌ BAD:
POST /createUser
GET /getUsers
POST /deleteUser/123
GET /getUserPosts/123

✅ GOOD:
POST /users           (create user)
GET /users            (get all users)
DELETE /users/123     (delete user)
GET /users/123/posts  (get posts by user 123)
```

### 2. Use HTTP Methods Correctly

| Method | Purpose | Idempotent? | Safe? |
|--------|---------|-------------|-------|
| **GET** | Retrieve data | Yes | Yes |
| **POST** | Create new resource | No | No |
| **PUT** | Update entire resource | Yes | No |
| **PATCH** | Partial update | No* | No |
| **DELETE** | Remove resource | Yes | No |

**Idempotent** = Calling it multiple times has the same effect as calling once
**Safe** = Doesn't modify data

**Examples:**
```
GET /users/123
→ Safe: Just reading, not changing anything
→ Idempotent: Calling it 10 times gives same result

POST /users
→ Not Safe: Creates new user
→ Not Idempotent: Calling it 10 times creates 10 users!

DELETE /users/123
→ Not Safe: Deletes user
→ Idempotent: Calling it 10 times = user still deleted (same state)
```

### 3. Use Proper HTTP Status Codes

**2xx - Success**
- `200 OK` - Request succeeded
- `201 Created` - New resource created
- `204 No Content` - Success, but no data to return

**3xx - Redirection**
- `301 Moved Permanently` - Resource has new URL
- `304 Not Modified` - Use cached version

**4xx - Client Errors**
- `400 Bad Request` - Invalid request format
- `401 Unauthorized` - Need to log in
- `403 Forbidden` - Logged in but not allowed
- `404 Not Found` - Resource doesn't exist
- `429 Too Many Requests` - Rate limit exceeded

**5xx - Server Errors**
- `500 Internal Server Error` - Something broke on server
- `502 Bad Gateway` - Server got invalid response from upstream
- `503 Service Unavailable` - Server temporarily down
- `504 Gateway Timeout` - Server didn't respond in time

### 4. Use Consistent Naming Conventions

```
✅ GOOD:
GET /users              (plural, lowercase)
GET /users/123
GET /users/123/posts
GET /users/123/posts/456/comments

❌ BAD:
GET /Users              (inconsistent capitalization)
GET /user/123           (inconsistent plural/singular)
GET /users/123/Post     (inconsistent capitalization)
```

### 5. Version Your API

APIs evolve. Don't break existing clients!

```
✅ GOOD:
/v1/users
/v2/users
/v1/posts

Or:
Header: Accept: application/vnd.myapp.v1+json
```

### 6. Use Query Parameters for Filtering, Sorting, Searching

```
GET /users?role=admin
GET /users?sort=created_at&order=desc
GET /users?search=john
GET /products?category=electronics&price_min=100&price_max=500
GET /posts?author=123&published=true&sort=likes
```

---

## Pagination: Handling Large Datasets

**The Problem:**
Imagine Instagram tried to send ALL posts in the world when you open the app. Your phone would explode! 💥

**The Solution: Pagination**
Send data in small "pages" (chunks).

### Method 1: Offset-Based Pagination

**How it works:**
```
GET /posts?limit=20&offset=0    → Posts 1-20 (page 1)
GET /posts?limit=20&offset=20   → Posts 21-40 (page 2)
GET /posts?limit=20&offset=40   → Posts 41-60 (page 3)
```

**Pros:**
- Simple to implement
- Easy to jump to any page
- Works with SQL: `SELECT * FROM posts LIMIT 20 OFFSET 40`

**Cons:**
- ❌ **Inconsistent with new data**: If 5 new posts are added while you're on page 2, you might see duplicates or miss posts
- ❌ **Slow for large offsets**: `OFFSET 1000000` makes database skip 1M rows
- ❌ **Not real-time friendly**

**Response format:**
```json
{
  "data": [
    {"id": 1, "title": "Post 1"},
    {"id": 2, "title": "Post 2"}
  ],
  "pagination": {
    "total": 100,
    "page": 1,
    "per_page": 20,
    "total_pages": 5
  }
}
```

### Method 2: Cursor-Based Pagination (Better for Feeds)

**How it works:**
Use the ID of the last item as the "cursor" for the next page.

```
GET /posts?limit=20                    → Posts 1-20, cursor=20
GET /posts?limit=20&cursor=20          → Posts 21-40, cursor=40
GET /posts?limit=20&cursor=40          → Posts 41-60, cursor=60
```

**Pros:**
- ✅ **Consistent results**: New posts don't affect your pagination
- ✅ **Fast**: Database uses index on ID
- ✅ **Perfect for infinite scroll**

**Cons:**
- Can't jump to specific page
- Slightly more complex

**Response format:**
```json
{
  "data": [
    {"id": 21, "title": "Post 21"},
    {"id": 22, "title": "Post 22"}
  ],
  "pagination": {
    "next_cursor": "22",
    "has_more": true
  }
}
```

**SQL implementation:**
```sql
-- First page
SELECT * FROM posts ORDER BY id DESC LIMIT 20;

-- Next pages
SELECT * FROM posts 
WHERE id < 20 
ORDER BY id DESC 
LIMIT 20;
```

### Method 3: Keyset Pagination (Most Efficient)

Like cursor pagination but works with any sortable field:

```
GET /posts?limit=20&sort=created_at
GET /posts?limit=20&sort=created_at&before=2024-01-15T10:30:00Z
```

**Perfect for:**
- Time-based feeds (social media)
- Sorted lists
- Real-time data

---

## Real-World Example: Twitter API

Let's design a simplified Twitter API:

### Endpoints

```
# Users
GET    /v1/users/:id              Get user profile
PUT    /v1/users/:id              Update user profile
GET    /v1/users/:id/followers    Get user's followers
POST   /v1/users/:id/follow       Follow user

# Tweets
GET    /v1/tweets                 Get feed (requires auth)
POST   /v1/tweets                 Create tweet
GET    /v1/tweets/:id             Get specific tweet
DELETE /v1/tweets/:id             Delete tweet
POST   /v1/tweets/:id/like        Like tweet
GET    /v1/tweets/:id/likes       Get who liked tweet

# Timeline
GET    /v1/timeline/home          Get home timeline (paginated)
GET    /v1/timeline/user/:id      Get user's tweets (paginated)
```

### Example Request/Response

**Request:**
```http
GET /v1/timeline/home?limit=20&cursor=abc123
Authorization: Bearer eyJhbGciOiJIUzI1NiIs...
```

**Response:**
```json
{
  "data": [
    {
      "id": "tweet_789",
      "user": {
        "id": "user_123",
        "username": "alice",
        "display_name": "Alice Smith"
      },
      "text": "Just learned about REST APIs!",
      "created_at": "2024-12-26T10:30:00Z",
      "likes_count": 42,
      "retweets_count": 7,
      "replies_count": 3
    }
  ],
  "pagination": {
    "next_cursor": "xyz789",
    "has_more": true
  }
}
```

---

## API Design Best Practices Summary

### 1. Be Consistent
- Use the same patterns across all endpoints
- Consistent naming, response formats, error handling

### 2. Be Predictable
- Follow REST conventions
- Use standard HTTP status codes
- Clear, logical URL structure

### 3. Think About Developers
- Good documentation is crucial
- Clear error messages
- Helpful validation feedback

**Good error response:**
```json
{
  "error": {
    "code": "INVALID_EMAIL",
    "message": "Email format is invalid",
    "field": "email",
    "suggestion": "Use format: user@example.com"
  }
}
```

### 4. Design for Scale
- Use pagination
- Allow filtering and sorting
- Support caching with proper headers

### 5. Secure by Default
- Always use HTTPS
- Require authentication
- Rate limit to prevent abuse
- Validate all input

### 6. Version Early
- Start with /v1/
- Never break existing clients
- Deprecate old versions gradually

---

## Rate Limiting

Prevent abuse and ensure fair usage:

```
Response Headers:
X-RateLimit-Limit: 1000          (requests per hour)
X-RateLimit-Remaining: 987       (requests left)
X-RateLimit-Reset: 1640000000    (when limit resets)

If exceeded:
Status: 429 Too Many Requests
Retry-After: 3600                (seconds to wait)
```

---

## API Authentication

Quick overview (we'll dive deeper later):

### 1. API Keys
```
GET /v1/users
X-API-Key: abc123xyz789
```
- Simple
- Used for server-to-server
- Can't identify individual users

### 2. OAuth 2.0 / JWT Tokens
```
GET /v1/users
Authorization: Bearer eyJhbGciOiJIUzI1NiIs...
```
- Industry standard
- Secure
- Can identify users
- Tokens expire

---

## Key Concepts to Remember

1. **REST is an architectural style, not a strict protocol**
2. **Use HTTP methods semantically** (GET for read, POST for create, etc.)
3. **Resources are nouns, actions are verbs (HTTP methods)**
4. **Always use proper status codes**
5. **Pagination is essential for scalability**
6. **Cursor-based pagination > Offset pagination for feeds**
7. **Version your API from day one**
8. **Design for developers—they're your users**

---

## Practice Questions

**Q1:** Design a RESTful API for a blog platform. List at least 8 endpoints covering:
- Posts (create, read, update, delete)
- Comments
- Users
- Tags/Categories

<details>
<summary>View Answer</summary>

```
Posts:
GET    /posts              - List all posts (with pagination)
POST   /posts              - Create a new post
GET    /posts/:id          - Get a specific post
PUT    /posts/:id          - Update a post
DELETE /posts/:id          - Delete a post

Comments:
GET    /posts/:id/comments - Get comments for a post
POST   /posts/:id/comments - Add a comment to a post
DELETE /comments/:id       - Delete a comment

Users:
GET    /users/:id          - Get user profile
PUT    /users/:id          - Update user profile
GET    /users/:id/posts    - Get all posts by a user

Tags/Categories:
GET    /tags               - List all tags
GET    /posts?tag=tech     - Filter posts by tag
GET    /categories         - List all categories
GET    /categories/:id/posts - Get posts in a category
```

</details>

**Q2:** What's wrong with these endpoints? How would you fix them?
```
GET /getUserById/123
POST /posts/delete/456
GET /posts/getByCategory?category=tech
PUT /updateUserEmail
```

<details>
<summary>View Answer</summary>

| Wrong | Problem | Fixed |
|-------|---------|-------|
| `GET /getUserById/123` | Verb in URL, not RESTful | `GET /users/123` |
| `POST /posts/delete/456` | Using POST for delete, verb in URL | `DELETE /posts/456` |
| `GET /posts/getByCategory?category=tech` | Redundant "getByCategory" | `GET /posts?category=tech` |
| `PUT /updateUserEmail` | Missing resource ID, verb in URL | `PATCH /users/123` with email in body |

**Key principles:**
- URLs should be nouns (resources), not verbs
- HTTP methods express the action
- Resource IDs go in the URL path
- Filters go in query parameters

</details>

**Q3:** You're building Instagram's feed API. Would you use offset-based or cursor-based pagination? Why?

<details>
<summary>View Answer</summary>

**Cursor-based pagination** is the right choice for Instagram's feed:

1. **Feed constantly changes:** New posts appear frequently. With offset-based, you'd see duplicates or miss posts as the list shifts
2. **Infinite scroll UX:** Cursor-based is perfect for "load more" functionality
3. **Performance:** Cursor-based is O(1) with an index, while offset-based becomes slower as offset grows (OFFSET 10000 must skip 10000 rows)
4. **Real-time data:** Cursor marks your exact position regardless of new insertions

**Example:**
```
GET /feed?cursor=eyJpZCI6MTIzNH0=&limit=20

Response includes:
{
  "posts": [...],
  "next_cursor": "eyJpZCI6MTI1NH0="
}
```

</details>

**Q4:** A client makes this request:
```
DELETE /users/123
```
The user doesn't exist. What HTTP status code should you return?
- 200 OK?
- 404 Not Found?
- 204 No Content?
Why?

<details>
<summary>View Answer</summary>

This is debatable, but **404 Not Found** is the most common and recommended choice:

- **404 Not Found:** The resource doesn't exist. Client should know they tried to delete something that wasn't there. Most RESTful approach.

- **204 No Content:** Some argue this is fine because the end state is achieved (user doesn't exist). This follows "idempotent" principles.

- **200 OK:** Misleading because no actual deletion happened.

**Best practice:** Return **404** with a clear message:
```json
{
  "error": "User not found",
  "code": "USER_NOT_FOUND"
}
```

This helps clients distinguish between "successfully deleted" vs "never existed" which is important for debugging and logging.

</details>

**Q5:** Design a pagination response format for an e-commerce product listing API. Include:
- Product data
- Total count
- Current page info
- Links to next/previous pages

<details>
<summary>View Answer</summary>

```json
{
  "data": [
    {"id": 1, "name": "iPhone 15", "price": 999},
    {"id": 2, "name": "Galaxy S24", "price": 899},
    {"id": 3, "name": "Pixel 8", "price": 699}
  ],
  "pagination": {
    "total_items": 1250,
    "total_pages": 125,
    "current_page": 3,
    "per_page": 10,
    "has_next": true,
    "has_previous": true
  },
  "links": {
    "self": "/products?page=3&per_page=10",
    "first": "/products?page=1&per_page=10",
    "last": "/products?page=125&per_page=10",
    "next": "/products?page=4&per_page=10",
    "previous": "/products?page=2&per_page=10"
  }
}
```

**Key elements:**
- `data`: The actual products
- `pagination`: Metadata about the current position
- `links`: HATEOAS-style navigation links (clients don't need to construct URLs)

</details>

**Q6:** Why is this a problem, and how would you fix it?
```
GET /api/users/123/posts    (returns 10,000 posts with no pagination)
```

<details>
<summary>View Answer</summary>

**Problems:**

1. **Performance:** Fetching 10,000 records is slow and memory-intensive
2. **Bandwidth:** Huge payload wastes bandwidth, especially on mobile
3. **Timeout risk:** Request might timeout before completing
4. **Poor UX:** User waits forever, browser might freeze
5. **Server strain:** One request consumes excessive resources
6. **No one needs 10K at once:** Users can only view ~10-50 at a time

**Fixes:**

```
# Add pagination (required)
GET /api/users/123/posts?page=1&limit=20

# Add cursor-based pagination (better for feeds)
GET /api/users/123/posts?cursor=abc123&limit=20

# Set a maximum limit server-side
limit = min(request.limit, 100)  # Cap at 100

# Add filtering options
GET /api/users/123/posts?limit=20&year=2024&sort=newest
```

**Best practice:** Always enforce a default limit (e.g., 20) and maximum limit (e.g., 100) on list endpoints.

</details>

---

## Next Up

Congratulations on completing Week 1! In Week 2, we'll start with **Networking Basics: DNS, Load Balancers, and CDNs** - the infrastructure that makes the internet fast and reliable!
