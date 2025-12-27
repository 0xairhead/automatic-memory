# Week 1, Lesson 1: Client-Server Architecture & HTTP Basics

## Table of Contents
- [Media Resources](#media-resources)
- [What is Client-Server Architecture?](#what-is-client-server-architecture)
  - [The Client](#the-client)
  - [The Server](#the-server)
  - [Simple Example](#simple-example)
- [HTTP: The Language They Speak](#http-the-language-they-speak)
  - [HTTP Request - What the Client Sends](#http-request---what-the-client-sends)
  - [HTTP Response - What the Server Sends Back](#http-response---what-the-server-sends-back)
- [Real-World Example: Loading Instagram](#real-world-example-loading-instagram)
- [HTTPS: The Secure Version](#https-the-secure-version)
- [Key Concepts to Remember](#key-concepts-to-remember)
- [Practice Questions](#practice-questions)
- [Next Up](#next-up)

---

Welcome to your first lesson! Let's build your foundation.

## Media Resources

**Visual Guide:**
![The Web's Restaurant: A Guide to Client-Server & HTTP](./assets/client_server_http_infographic.png)

**Audio Lesson:**
[HTTP, Client-Server and Status Codes (Audio)](./assets/HTTP_Client_Server_and_Status_Codes.m4a)

---

## What is Client-Server Architecture?

Think of a restaurant:
- **You (the customer)** = Client
- **The kitchen** = Server
- **The waiter** = Network/Protocol

You don't cook your own food. You make a request (order), the waiter carries it to the kitchen, the kitchen prepares it, and the waiter brings back your food. This is exactly how the internet works!

### The Client

The client is anything that **requests** information or services:
- Your web browser (Chrome, Safari)
- Your mobile app (Instagram, Gmail)
- Even other servers can be clients!

**What clients do:**
- Send requests ("Give me the homepage")
- Receive responses (the HTML/data)
- Display information to users
- Handle user interactions

### The Server

The server **provides** information or services:
- A computer that's always running and waiting for requests
- Stores data, runs business logic, manages databases
- Sends back responses to clients

**What servers do:**
- Listen for incoming requests
- Process the request (maybe query a database)
- Send back a response (data, webpage, error message)

### Simple Example

```
You type "twitter.com" in your browser:

1. Your browser (client) sends: "GET me the Twitter homepage"
2. Twitter's server receives it
3. Server thinks: "Who is this user? What should I show them?"
4. Server gathers the data (recent tweets, trends, etc.)
5. Server sends back: HTML, CSS, JavaScript, images
6. Your browser displays the beautiful Twitter feed
```

---

## HTTP: The Language They Speak

**HTTP = HyperText Transfer Protocol**

It's like the agreed-upon language between clients and servers. Just like you and the waiter both speak English (or Spanish, etc.), clients and servers speak HTTP.

### HTTP Request - What the Client Sends

An HTTP request has several parts:

**1. Method (What do you want to do?)**
- `GET` - "Give me some data" (reading)
- `POST` - "Here's new data, save it" (creating)
- `PUT` - "Update this data" (updating)
- `DELETE` - "Remove this data" (deleting)

**2. URL (Where?)**
- `https://twitter.com/home`
- `https://api.twitter.com/tweets/12345`

**3. Headers (Extra information)**
```
User-Agent: Chrome/120.0
Cookie: session_id=abc123
Content-Type: application/json
```

**4. Body (Optional - the actual data)**
```json
{
  "tweet": "Just learned about HTTP!",
  "user_id": 789
}
```

### HTTP Response - What the Server Sends Back

**1. Status Code (Did it work?)**
- `200 OK` - Success! Here's your data
- `201 Created` - Successfully created new resource
- `400 Bad Request` - Your request doesn't make sense
- `401 Unauthorized` - You need to log in
- `404 Not Found` - That page doesn't exist
- `500 Internal Server Error` - Oops, something broke on our end

**2. Headers**
```
Content-Type: application/json
Set-Cookie: session_id=xyz789
Cache-Control: max-age=3600
```

**3. Body (The actual data)**
```json
{
  "user": "Alice",
  "tweets": [...]
}
```

---

## Real-World Example: Loading Instagram

Let me walk you through what happens when you open Instagram:

```
Step 1: You open the app
└─> Client sends: GET https://instagram.com/api/feed
    Headers: Authorization: Bearer [your_token]

Step 2: Instagram server receives it
└─> Checks: Is this user logged in? (validates token)
    Queries database: Get recent posts from people Alice follows
    Processes: Ranks posts, applies algorithm

Step 3: Server responds
└─> Status: 200 OK
    Body: JSON with 30 posts, images, likes, comments

Step 4: Your app receives the data
└─> Displays the feed
    Loads images
    Shows like counts
```

But wait! Loading those images requires MORE requests:

```
For each image:
Client sends: GET https://instagram.com/images/post123.jpg
Server responds: The actual image file
```

One page load = dozens or hundreds of HTTP requests!

---

## HTTPS: The Secure Version

**HTTPS = HTTP + S (Secure)**

Regular HTTP sends everything in plain text. Anyone between you and the server can read it!

HTTPS encrypts everything:
- Your passwords can't be stolen
- Your credit card info stays safe
- Nobody can modify the data in transit

**How it works (simplified):**
1. Your browser and server perform a "handshake"
2. They agree on encryption keys
3. All data is encrypted before sending
4. Only the intended recipient can decrypt it

---

## Key Concepts to Remember

1. **Client-Server is a relationship** - One requests, one responds
2. **HTTP is stateless** - Each request is independent (server doesn't "remember" you between requests)
3. **HTTPS encrypts** - Always use HTTPS for sensitive data
4. **Multiple requests** - One webpage = many client-server interactions

---

## Practice Questions

Think about these, and when you're ready, share your answers with me:

**Q1:** You're building a todo app. What HTTP method would you use for:
- Getting all todos?
- Creating a new todo?
- Marking a todo as complete?
- Deleting a todo?

**Q2:** If HTTP is stateless, how does Twitter "remember" you're logged in when you navigate from your feed to your profile page?

**Q3:** You visit an e-commerce site and get a "504 Gateway Timeout" error. Based on what you learned, what does this likely mean? (Hint: Is it a 4xx client error or 5xx server error?)

**Q4:** Why do you think modern websites make so many HTTP requests? (Think about images, styles, scripts, data...)

---

## Next Up

In Lesson 2, we'll explore **Databases: SQL vs NoSQL** - understanding when to use relational databases and when NoSQL is the better choice!
