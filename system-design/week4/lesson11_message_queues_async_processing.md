# Week 4, Lesson 11: Message Queues & Async Processing

## Table of Contents
- [The Synchronous Problem](#the-synchronous-problem)
- [Synchronous vs Asynchronous](#synchronous-vs-asynchronous)
  - [Synchronous Communication](#synchronous-communication)
  - [Asynchronous Communication](#asynchronous-communication)
- [What is a Message Queue?](#what-is-a-message-queue)
  - [Core Components](#core-components)
  - [How It Works](#how-it-works)
- [Message Queue Patterns](#message-queue-patterns)
  - [Pattern 1: Point-to-Point (Queue)](#pattern-1-point-to-point-queue)
  - [Pattern 2: Publish-Subscribe (Pub/Sub)](#pattern-2-publish-subscribe-pubsub)
  - [Pattern 3: Request-Reply](#pattern-3-request-reply)
  - [Pattern 4: Fan-Out](#pattern-4-fan-out)
  - [Pattern 5: Work Queue (Competing Consumers)](#pattern-5-work-queue-competing-consumers)
- [Message Delivery Guarantees](#message-delivery-guarantees)
  - [At-Most-Once](#at-most-once)
  - [At-Least-Once](#at-least-once)
  - [Exactly-Once](#exactly-once)
- [Popular Message Queue Systems](#popular-message-queue-systems)
  - [RabbitMQ](#rabbitmq)
  - [Apache Kafka](#apache-kafka)
  - [Amazon SQS](#amazon-sqs)
  - [Redis (as a Queue)](#redis-as-a-queue)
  - [Comparison Table](#comparison-table)
- [Real-World Use Cases](#real-world-use-cases)
  - [Email Sending](#email-sending)
  - [Order Processing](#order-processing)
  - [Image/Video Processing](#imagevideo-processing)
  - [Notifications](#notifications)
  - [Data Pipelines](#data-pipelines)
- [Designing with Message Queues](#designing-with-message-queues)
  - [When to Use Message Queues](#when-to-use-message-queues)
  - [When NOT to Use Message Queues](#when-not-to-use-message-queues)
- [Handling Failures](#handling-failures)
  - [Dead Letter Queues](#dead-letter-queues)
  - [Retry Strategies](#retry-strategies)
  - [Idempotency](#idempotency)
- [Scaling Message Queues](#scaling-message-queues)
- [Common Mistakes](#common-mistakes)
- [Key Concepts to Remember](#key-concepts-to-remember)
- [Practice Questions](#practice-questions)
- [Next Up](#next-up)

---

Welcome to Lesson 11! We're continuing Week 4's distribution fundamentals by exploring **message queues** - the backbone of asynchronous, decoupled systems. Understanding queues is essential for building scalable, resilient architectures.

---

## The Synchronous Problem

Imagine an e-commerce checkout flow:

```
User clicks "Buy Now"
         │
         ▼
    ┌─────────────────────────────────────────────────────┐
    │              Synchronous Processing                  │
    │                                                     │
    │  1. Validate payment ──────────────────────► 500ms  │
    │  2. Update inventory ──────────────────────► 200ms  │
    │  3. Send confirmation email ───────────────► 800ms  │
    │  4. Update analytics ──────────────────────► 300ms  │
    │  5. Notify warehouse ──────────────────────► 400ms  │
    │  6. Update recommendation engine ──────────► 600ms  │
    │                                                     │
    │  Total: 2,800ms (2.8 seconds!)                      │
    └─────────────────────────────────────────────────────┘
         │
         ▼
    User sees "Order confirmed"
```

**Problems:**
- User waits 2.8 seconds (terrible UX!)
- If email service is slow → entire checkout is slow
- If analytics is down → checkout fails completely
- All services must be up simultaneously

---

## Synchronous vs Asynchronous

### Synchronous Communication

```
Client ──request──► Server ──wait──► Client ──response──► Done
         │                                      │
         └──────────── Blocked! ───────────────┘

The client WAITS until the server responds.
Like a phone call - you wait for the other person to respond.
```

**Characteristics:**
- Request-response in real-time
- Caller is blocked until response
- Tight coupling between services
- Easier to reason about
- Failures propagate immediately

**Example:**
```python
# Synchronous API call
def checkout(order):
    payment = process_payment(order)      # Wait...
    inventory = update_inventory(order)    # Wait...
    email = send_confirmation(order)       # Wait...
    return {"status": "complete"}          # Finally!
```

### Asynchronous Communication

```
Client ──request──► Queue ──► "Got it!" ──► Client continues
                      │
                      ▼
              Worker processes later

The client does NOT wait for processing to complete.
Like sending an email - you don't wait for a reply.
```

**Characteristics:**
- Fire and forget (or fire and callback)
- Caller continues immediately
- Loose coupling between services
- Better fault isolation
- More complex to reason about

**Example:**
```python
# Asynchronous with message queue
def checkout(order):
    payment = process_payment(order)       # Still sync (critical!)

    # Queue these for async processing
    queue.publish("inventory", order)      # Returns immediately
    queue.publish("email", order)          # Returns immediately
    queue.publish("analytics", order)      # Returns immediately

    return {"status": "complete"}          # User sees this fast!
```

**Result:**
```
Synchronous: 2,800ms
Asynchronous: ~600ms (just payment + queue publish)

4.7x faster response time!
```

---

## What is a Message Queue?

A message queue is a form of **asynchronous service-to-service communication**.

```
┌──────────┐         ┌─────────────┐         ┌──────────┐
│ Producer │────────►│   Message   │────────►│ Consumer │
│          │ publish │    Queue    │ consume │          │
└──────────┘         └─────────────┘         └──────────┘

Producer: Creates and sends messages
Queue: Stores messages until consumed
Consumer: Retrieves and processes messages
```

### Core Components

```
┌─────────────────────────────────────────────────────────────┐
│                      MESSAGE QUEUE                          │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│   ┌─────┐  ┌─────┐  ┌─────┐  ┌─────┐  ┌─────┐             │
│   │ Msg │  │ Msg │  │ Msg │  │ Msg │  │ Msg │             │
│   │  1  │◄─│  2  │◄─│  3  │◄─│  4  │◄─│  5  │◄── New      │
│   └─────┘  └─────┘  └─────┘  └─────┘  └─────┘             │
│      │                                                      │
│      ▼                                                      │
│   Consumer reads                                            │
│                                                             │
│   Properties:                                               │
│   - FIFO (usually): First In, First Out                    │
│   - Persistent: Messages survive restarts                  │
│   - Acknowledgment: Confirm processing complete            │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### How It Works

```
Step 1: Producer sends message
┌──────────┐                    ┌─────────────┐
│ Order    │───"Process order"─►│    Queue    │
│ Service  │                    │ [─────────] │
└──────────┘                    └─────────────┘

Step 2: Queue stores message
┌─────────────┐
│    Queue    │
│ [Order#123] │  ← Message persisted
└─────────────┘

Step 3: Consumer pulls message
┌─────────────┐                    ┌──────────┐
│    Queue    │───"Order#123"─────►│ Inventory│
│ [─────────] │                    │ Service  │
└─────────────┘                    └──────────┘

Step 4: Consumer acknowledges
┌──────────┐                    ┌─────────────┐
│ Inventory│───────"ACK"───────►│    Queue    │
│ Service  │                    │  (removes)  │
└──────────┘                    └─────────────┘
```

---

## Message Queue Patterns

### Pattern 1: Point-to-Point (Queue)

**One producer, one consumer per message**

```
┌──────────┐         ┌─────────┐         ┌──────────┐
│ Producer │────────►│  Queue  │────────►│ Consumer │
└──────────┘         └─────────┘         └──────────┘

Each message is processed by exactly ONE consumer.
```

**Use case:** Task processing, job queues

```
Example: Email sending

[Web Server] ──► [Email Queue] ──► [Email Worker]
                 │ send to: alice │
                 │ send to: bob   │
                 │ send to: carol │

Each email sent by exactly one worker.
```

### Pattern 2: Publish-Subscribe (Pub/Sub)

**One publisher, multiple subscribers**

```
                              ┌──────────────┐
                         ┌───►│ Subscriber A │
                         │    └──────────────┘
┌───────────┐   ┌──────┐ │    ┌──────────────┐
│ Publisher │──►│ Topic │─┼───►│ Subscriber B │
└───────────┘   └──────┘ │    └──────────────┘
                         │    ┌──────────────┐
                         └───►│ Subscriber C │
                              └──────────────┘

Each message is delivered to ALL subscribers.
```

**Use case:** Event broadcasting, notifications

```
Example: Order placed event

[Order Service] ──"ORDER_PLACED"──► [Orders Topic]
                                         │
                    ┌────────────────────┼────────────────────┐
                    ▼                    ▼                    ▼
             [Inventory]           [Email Service]      [Analytics]
             Reduce stock          Send confirmation    Track sale

All services receive the same event!
```

### Pattern 3: Request-Reply

**Async request with eventual response**

```
┌──────────┐  request   ┌─────────┐  request  ┌──────────┐
│ Client   │───────────►│ Request │──────────►│ Server   │
│          │            │  Queue  │           │          │
└──────────┘            └─────────┘           └──────────┘
     ▲                                              │
     │                  ┌─────────┐                 │
     └──────────────────│  Reply  │◄───────────────┘
            response    │  Queue  │    response
                        └─────────┘

Client sends request to one queue, waits for reply on another.
```

**Use case:** Async API calls, long-running operations

```
Example: PDF generation

Client ──"Generate report"──► [Request Queue]
   │                                │
   │ (waits on reply queue)         ▼
   │                         [PDF Generator]
   │                         (takes 30 seconds)
   │                                │
   ▼                                ▼
[Reply Queue] ◄────────────"PDF ready at /files/123.pdf"
```

### Pattern 4: Fan-Out

**One message triggers multiple parallel processes**

```
                    ┌─────────┐     ┌─────────┐     ┌──────────┐
               ┌───►│ Queue A │────►│Worker A1│────►│ Result A │
               │    └─────────┘     └─────────┘     └──────────┘
┌───────────┐  │    ┌─────────┐     ┌─────────┐     ┌──────────┐
│  Message  │──┼───►│ Queue B │────►│Worker B1│────►│ Result B │
└───────────┘  │    └─────────┘     └─────────┘     └──────────┘
               │    ┌─────────┐     ┌─────────┐     ┌──────────┐
               └───►│ Queue C │────►│Worker C1│────►│ Result C │
                    └─────────┘     └─────────┘     └──────────┘

One event fans out to multiple processing pipelines.
```

**Use case:** Multi-step processing, parallel execution

```
Example: Video upload

"Video uploaded" ──► Fan-out Exchange
                         │
         ┌───────────────┼───────────────┐
         ▼               ▼               ▼
   [Transcode Queue] [Thumbnail Q] [Metadata Q]
         │               │               │
         ▼               ▼               ▼
    720p, 1080p,    Generate 5      Extract title,
    4K versions     thumbnails      duration, etc.
```

### Pattern 5: Work Queue (Competing Consumers)

**Multiple consumers share work from one queue**

```
                              ┌────────────┐
                         ┌───►│ Consumer 1 │
                         │    └────────────┘
┌──────────┐   ┌───────┐ │    ┌────────────┐
│ Producer │──►│ Queue │─┼───►│ Consumer 2 │
└──────────┘   └───────┘ │    └────────────┘
                         │    ┌────────────┐
                         └───►│ Consumer 3 │
                              └────────────┘

Messages are distributed among consumers (load balanced).
Each message goes to ONLY ONE consumer.
```

**Use case:** Parallel processing, scaling workers

```
Example: Image processing

[Upload Service] ──► [Image Queue: 1000 images]
                           │
              ┌────────────┼────────────┐
              ▼            ▼            ▼
         [Worker 1]   [Worker 2]   [Worker 3]
         333 images   333 images   334 images

Work is split among workers!
```

---

## Message Delivery Guarantees

### At-Most-Once

```
Producer ──message──► Queue ──message──► Consumer
                              │
                              └── No acknowledgment
                                  Message may be lost

Delivery: 0 or 1 times
```

**Characteristics:**
- Fastest (no ack overhead)
- Messages may be lost
- No duplicates

**Use case:** Metrics, logs where losing some data is acceptable

### At-Least-Once

```
Producer ──message──► Queue ──message──► Consumer
                              │              │
                              │◄────ACK──────┘
                              │
                        (if no ACK, redeliver)

Delivery: 1 or more times
```

**Characteristics:**
- Messages won't be lost
- May have duplicates (if ACK fails after processing)
- Consumer must be idempotent

**Use case:** Most applications - email, orders, notifications

```
Scenario causing duplicates:

1. Consumer receives message
2. Consumer processes message successfully
3. Consumer sends ACK
4. Network drops ACK ← Problem!
5. Queue assumes failure, redelivers
6. Consumer processes AGAIN (duplicate!)

Solution: Make consumer idempotent
```

### Exactly-Once

```
Producer ──message──► Queue ──message──► Consumer
                        │                   │
                        │◄─Transactional────┘
                        │   ACK
                        │
              (Complex coordination to prevent
               both loss AND duplicates)

Delivery: Exactly 1 time
```

**Characteristics:**
- Most complex to implement
- Highest overhead
- Requires distributed transactions or deduplication

**Use case:** Financial transactions, critical data

```
Implementation approaches:

1. Transactional Outbox:
   - Write message + business data in same DB transaction
   - Separate process reads outbox, publishes to queue

2. Idempotency Keys:
   - Each message has unique ID
   - Consumer tracks processed IDs
   - Duplicates are ignored

3. Kafka's Exactly-Once:
   - Idempotent producer
   - Transactional consumer
   - Coordinated through Kafka's transaction coordinator
```

---

## Popular Message Queue Systems

### RabbitMQ

```
┌─────────────────────────────────────────────────────────────┐
│                        RabbitMQ                              │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  Type: Traditional Message Broker                           │
│  Protocol: AMQP (Advanced Message Queuing Protocol)         │
│                                                             │
│  ┌──────────┐    ┌──────────┐    ┌─────────┐    ┌────────┐ │
│  │ Producer │───►│ Exchange │───►│  Queue  │───►│Consumer│ │
│  └──────────┘    └──────────┘    └─────────┘    └────────┘ │
│                       │                                     │
│            Routing logic here                               │
│            (direct, topic, fanout)                          │
│                                                             │
│  Best for:                                                  │
│  - Complex routing logic                                    │
│  - Request-reply patterns                                   │
│  - Traditional message queuing                              │
│  - When message order per consumer matters                  │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

**Example:**
```python
# RabbitMQ producer
import pika

connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
channel = connection.channel()
channel.queue_declare(queue='tasks')

channel.basic_publish(
    exchange='',
    routing_key='tasks',
    body='Process order #123'
)
```

### Apache Kafka

```
┌─────────────────────────────────────────────────────────────┐
│                        Apache Kafka                          │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  Type: Distributed Event Streaming Platform                 │
│  Protocol: Custom binary protocol                           │
│                                                             │
│  ┌──────────┐    ┌─────────────────────────┐    ┌────────┐ │
│  │ Producer │───►│         Topic           │───►│Consumer│ │
│  └──────────┘    │ ┌─────┬─────┬─────┐     │    │ Group  │ │
│                  │ │ P0  │ P1  │ P2  │     │    └────────┘ │
│                  │ └─────┴─────┴─────┘     │               │
│                  │     Partitions          │               │
│                  └─────────────────────────┘               │
│                                                             │
│  Key features:                                              │
│  - Persistent message log (retains all messages)           │
│  - High throughput (millions/sec)                          │
│  - Message replay (re-read old messages)                   │
│  - Partitioned for parallelism                             │
│                                                             │
│  Best for:                                                  │
│  - Event streaming                                          │
│  - Log aggregation                                          │
│  - High-throughput applications                             │
│  - Event sourcing                                           │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

**Example:**
```python
# Kafka producer
from kafka import KafkaProducer

producer = KafkaProducer(bootstrap_servers='localhost:9092')

producer.send(
    'orders',
    key=b'user_123',
    value=b'{"order_id": 456, "amount": 99.99}'
)
```

### Amazon SQS

```
┌─────────────────────────────────────────────────────────────┐
│                        Amazon SQS                            │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  Type: Managed Queue Service                                │
│  Protocol: HTTP/HTTPS API                                   │
│                                                             │
│  Two queue types:                                           │
│                                                             │
│  Standard Queue:               FIFO Queue:                  │
│  ┌─────────────┐              ┌─────────────┐              │
│  │ ≈ ordered   │              │ Strictly    │              │
│  │ At-least-   │              │ ordered     │              │
│  │ once        │              │ Exactly-    │              │
│  │ Unlimited   │              │ once        │              │
│  │ throughput  │              │ 300 msg/sec │              │
│  └─────────────┘              └─────────────┘              │
│                                                             │
│  Best for:                                                  │
│  - AWS-native applications                                  │
│  - Simple queue needs                                       │
│  - Serverless (Lambda triggers)                             │
│  - No infrastructure management                             │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

**Example:**
```python
# SQS producer
import boto3

sqs = boto3.client('sqs')

sqs.send_message(
    QueueUrl='https://sqs.us-east-1.amazonaws.com/123456/my-queue',
    MessageBody='Process order #123',
    MessageAttributes={
        'OrderType': {'StringValue': 'standard', 'DataType': 'String'}
    }
)
```

### Redis (as a Queue)

```
┌─────────────────────────────────────────────────────────────┐
│                     Redis as Queue                           │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  Type: In-memory data structure (List or Stream)            │
│                                                             │
│  Using Lists (simple):                                      │
│  LPUSH myqueue "message"    # Add to queue                 │
│  BRPOP myqueue 0            # Blocking pop                 │
│                                                             │
│  Using Streams (better):                                    │
│  XADD mystream * field value    # Add message              │
│  XREAD STREAMS mystream >       # Read new messages        │
│  XACK mystream group msgid      # Acknowledge              │
│                                                             │
│  Best for:                                                  │
│  - Simple queuing needs                                     │
│  - When you already use Redis                               │
│  - Low latency requirements                                 │
│  - Temporary/ephemeral messages                             │
│                                                             │
│  Limitations:                                               │
│  - Memory-bound                                             │
│  - Less durable by default                                  │
│  - Simpler features than dedicated queues                   │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### Comparison Table

| Feature | RabbitMQ | Kafka | SQS | Redis |
|---------|----------|-------|-----|-------|
| **Type** | Broker | Log | Managed | In-memory |
| **Throughput** | ~50K/sec | Millions/sec | ~3K/sec | ~100K/sec |
| **Latency** | Low | Low | Medium | Very Low |
| **Ordering** | Per queue | Per partition | FIFO option | FIFO |
| **Replay** | No | Yes | No | Limited |
| **Persistence** | Yes | Yes | Yes | Optional |
| **Complexity** | Medium | High | Low | Low |
| **Best For** | Routing | Streaming | AWS apps | Simple cases |

---

## Real-World Use Cases

### Email Sending

```
Problem: Sending emails is slow (500ms-2s per email)

Synchronous:
User signs up → Send welcome email → Return response
               ↑
         2 seconds waiting!

Asynchronous:
User signs up → Queue email → Return response immediately
                    │
                    ▼ (background)
              Email worker sends email

┌────────────┐     ┌─────────────┐     ┌────────────┐
│ Web Server │────►│ Email Queue │────►│ Email      │
│            │     │             │     │ Worker     │
│ "Signup    │     │ to: alice   │     │            │
│  complete" │     │ to: bob     │     │ Sends at   │
│  (instant) │     │ to: carol   │     │ own pace   │
└────────────┘     └─────────────┘     └────────────┘
```

### Order Processing

```
┌─────────────────────────────────────────────────────────────┐
│                    Order Processing Pipeline                 │
└─────────────────────────────────────────────────────────────┘

[User] ──► [Order Service] ──► [Order Created Event]
                                      │
           ┌──────────────────────────┼──────────────────────┐
           ▼                          ▼                      ▼
    [Payment Queue]           [Inventory Queue]      [Notification Q]
           │                          │                      │
           ▼                          ▼                      ▼
    [Payment Worker]          [Inventory Worker]     [Email Worker]
    Process payment           Reserve items          Send confirmation
           │                          │
           ▼                          ▼
    [Payment Complete]        [Items Reserved]
           │                          │
           └──────────┬───────────────┘
                      ▼
              [Shipping Queue]
                      │
                      ▼
              [Shipping Worker]
              Schedule pickup
```

### Image/Video Processing

```
User uploads image
        │
        ▼
┌─────────────────┐
│  Upload Service │
│                 │
│ 1. Store raw    │
│ 2. Queue jobs   │
│ 3. Return URL   │
└────────┬────────┘
         │
    ┌────┴────┐
    ▼         ▼
[Resize Q] [Thumbnail Q]
    │         │
    ▼         ▼
┌────────┐ ┌────────┐
│Resize  │ │Thumb   │
│Worker  │ │Worker  │
│        │ │        │
│ 100x100│ │ Create │
│ 500x500│ │ preview│
│1000x...│ │        │
└────────┘ └────────┘
    │         │
    └────┬────┘
         ▼
   [Storage Service]
   Save processed images
```

### Notifications

```
Event: "User liked your post"
              │
              ▼
    [Notification Service]
              │
    ┌─────────┼─────────┐
    ▼         ▼         ▼
[Push Q]  [Email Q]  [SMS Q]
    │         │         │
    ▼         ▼         ▼
[Push    [Email    [SMS
Worker]  Worker]   Worker]
    │         │         │
    ▼         ▼         ▼
 Mobile    Inbox     Phone
```

### Data Pipelines

```
┌─────────────────────────────────────────────────────────────┐
│                  Real-time Analytics Pipeline                │
└─────────────────────────────────────────────────────────────┘

[Web Servers]──┐
[Mobile Apps]──┼──► [Kafka: Events Topic] ──► [Stream Processor]
[IoT Devices]──┘           │                        │
                           │                        ▼
                           │                [Aggregations]
                           │                        │
                           ▼                        ▼
                    [Raw Storage]           [Real-time Dashboard]
                    (Data Lake)             (Grafana/Kibana)
```

---

## Designing with Message Queues

### When to Use Message Queues

```
✅ Long-running tasks
   - Video encoding
   - Report generation
   - Data exports

✅ Spiky workloads
   - Flash sales
   - Viral content
   - Marketing campaigns

✅ Decoupling services
   - Independent deployments
   - Different scaling needs
   - Fault isolation

✅ Cross-service communication
   - Microservices events
   - System integrations
   - Third-party webhooks

✅ Rate limiting downstream services
   - Database writes
   - External API calls
   - Email sending

✅ Guaranteed delivery
   - Financial transactions
   - Order processing
   - Critical notifications
```

### When NOT to Use Message Queues

```
❌ Real-time responses needed
   - Login authentication
   - Search queries
   - Page rendering

❌ Simple request-response
   - REST API calls
   - Database queries
   - Cache lookups

❌ Very low latency required
   - Gaming
   - Trading systems
   - Real-time bidding

❌ Small scale / simple systems
   - Single-server apps
   - MVP / prototypes
   - Low traffic sites

❌ Tight consistency requirements
   - Banking balance checks
   - Inventory availability display
   - Real-time collaboration
```

---

## Handling Failures

### Dead Letter Queues

```
What happens when a message fails repeatedly?

Main Queue                     Dead Letter Queue
┌─────────────────┐           ┌─────────────────┐
│ msg1 ✓          │           │                 │
│ msg2 ✓          │           │                 │
│ msg3 ✗ (fail 1) │           │                 │
│ msg3 ✗ (fail 2) │           │                 │
│ msg3 ✗ (fail 3) │──────────►│ msg3 (dead)     │
│                 │ 3 retries │                 │
│ msg4 ✓          │  = move   │                 │
└─────────────────┘           └─────────────────┘
                                     │
                                     ▼
                              [Alert + Manual Review]
```

**DLQ Configuration:**
```python
# RabbitMQ DLQ setup
channel.queue_declare(
    queue='orders',
    arguments={
        'x-dead-letter-exchange': '',
        'x-dead-letter-routing-key': 'orders-dlq',
        'x-message-ttl': 86400000  # 24 hours
    }
)

channel.queue_declare(queue='orders-dlq')
```

### Retry Strategies

```
Immediate Retry (Bad):
Fail → Retry → Fail → Retry → Fail → ...
       0ms      0ms      0ms
Overwhelms the system!

Exponential Backoff (Good):
Fail → Wait 1s → Retry → Wait 2s → Retry → Wait 4s → ...

Exponential Backoff with Jitter (Best):
Fail → Wait 1s ± random → Retry → Wait 2s ± random → ...
Prevents thundering herd
```

**Implementation:**
```python
def process_with_retry(message, max_retries=5):
    for attempt in range(max_retries):
        try:
            return process(message)
        except Exception as e:
            if attempt == max_retries - 1:
                send_to_dlq(message)
                raise

            # Exponential backoff with jitter
            delay = (2 ** attempt) + random.uniform(0, 1)
            time.sleep(delay)
```

### Idempotency

```
Problem: Message delivered twice = processed twice!

Order: {"id": 123, "action": "charge $100"}

Without idempotency:
Process → Charge $100
Process again (retry) → Charge $100 AGAIN!
Customer charged $200! 😱

With idempotency:
Process → Check: "Did I process 123 before?" No → Charge $100 → Mark 123 done
Process again → Check: "Did I process 123 before?" YES → Skip
Customer charged $100 ✓
```

**Implementation:**
```python
def process_order(message):
    order_id = message['order_id']

    # Check if already processed
    if redis.get(f"processed:{order_id}"):
        logger.info(f"Order {order_id} already processed, skipping")
        return

    # Process the order
    charge_customer(message)
    update_inventory(message)

    # Mark as processed (with TTL)
    redis.setex(f"processed:{order_id}", 86400, "done")
```

---

## Scaling Message Queues

```
Scaling Producers:
┌──────────┐
│Producer 1│──┐
└──────────┘  │    ┌─────────┐
┌──────────┐  ├───►│  Queue  │
│Producer 2│──┤    └─────────┘
└──────────┘  │
┌──────────┐  │
│Producer 3│──┘
└──────────┘
Easy! Just add more producers.


Scaling Consumers:
              ┌─────────┐    ┌──────────┐
              │ Queue   │───►│Consumer 1│
              │         │    └──────────┘
              │ ─────── │    ┌──────────┐
              │ ─────── │───►│Consumer 2│
              │ ─────── │    └──────────┘
              │ ─────── │    ┌──────────┐
              └─────────┘───►│Consumer 3│
                             └──────────┘
Add consumers to process faster.
Each message goes to ONE consumer.


Scaling the Queue (Kafka):
┌─────────────────────────────────────────┐
│                Topic                     │
│  ┌────────────┐┌────────────┐┌────────┐ │
│  │Partition 0 ││Partition 1 ││Part. 2 │ │
│  │ msg1, msg4 ││ msg2, msg5 ││msg3,msg6│ │
│  └────────────┘└────────────┘└────────┘ │
└─────────────────────────────────────────┘
        │              │            │
        ▼              ▼            ▼
   Consumer 0    Consumer 1   Consumer 2

Partitions allow parallel processing.
Each partition → one consumer (per group).
More partitions → more parallelism.
```

---

## Common Mistakes

### Mistake 1: Not Handling Duplicates

```
❌ Assuming exactly-once delivery
   message arrives → process blindly

✅ Always implement idempotency
   message arrives → check if processed → process or skip
```

### Mistake 2: No Dead Letter Queue

```
❌ Message fails → retry forever
   Queue backs up, system degrades

✅ Message fails N times → move to DLQ
   Investigate manually, fix issues
```

### Mistake 3: Huge Messages

```
❌ Putting 10MB image in message
   Slow, memory-intensive, queue backs up

✅ Put reference in message
   {
     "image_id": "abc123",
     "storage_url": "s3://bucket/abc123.jpg"
   }
   Worker fetches image separately
```

### Mistake 4: No Monitoring

```
❌ Ignoring queue metrics

✅ Monitor:
   - Queue depth (messages waiting)
   - Processing rate (messages/sec)
   - Error rate
   - Consumer lag (Kafka)
   - DLQ size

Alert when queue depth grows faster than consumption!
```

### Mistake 5: Wrong Queue for the Job

```
❌ Using Kafka for 10 messages/day
   Overkill! Complex setup for simple needs.

❌ Using Redis Lists for critical financial events
   May lose messages if Redis crashes.

✅ Match tool to requirements:
   - Simple tasks → Redis or SQS
   - Complex routing → RabbitMQ
   - High throughput streaming → Kafka
   - AWS native → SQS/SNS
```

### Mistake 6: Blocking in Consumers

```
❌ Consumer blocks on slow operations
   def process(message):
       response = slow_api_call()  # 30 second timeout
       # Queue message invisible, blocking other work

✅ Set appropriate timeouts, use async where possible
   def process(message):
       response = api_call(timeout=5)  # Fail fast
       if timeout:
           requeue with backoff
```

---

## Key Concepts to Remember

1. **Async communication** decouples services and improves response times
2. **Message queues** store messages until consumers process them
3. **Point-to-point** = one consumer per message; **Pub/Sub** = all subscribers get every message
4. **At-least-once** is most common; requires **idempotent consumers**
5. **Dead Letter Queues** catch messages that fail repeatedly
6. **Exponential backoff with jitter** prevents thundering herd on retries
7. **Kafka** = high throughput, message replay, event streaming
8. **RabbitMQ** = complex routing, traditional queuing
9. **SQS** = simple, managed, AWS-native
10. **Don't put large payloads in messages** - use references to external storage

---

## Practice Questions

**Q1:** An e-commerce site sends confirmation emails synchronously during checkout, causing 2-second delays. How would you redesign this using a message queue? What components would you need?

<details>
<summary>View Answer</summary>

**Current Problem:**
```
User checkout → Process payment → Send email (2s) → Return response
                                        ↑
                              User waiting here!
```

**Redesigned Architecture:**

```
┌────────────────────────────────────────────────────────────┐
│                     Checkout Flow                           │
└────────────────────────────────────────────────────────────┘

[User] ──► [Checkout Service]
                  │
                  ├── 1. Validate order
                  ├── 2. Process payment (sync - critical!)
                  ├── 3. Save order to DB
                  ├── 4. Publish to Email Queue ◄── Async!
                  │         │
                  │         ▼
                  │   ┌─────────────┐
                  │   │ Email Queue │
                  │   │ order: 123  │
                  │   │ order: 124  │
                  │   └─────────────┘
                  │         │
                  └── 5. Return "Order confirmed!"
                          │
                          ▼ (background)
                    ┌─────────────┐
                    │Email Worker │
                    │             │
                    │ Fetch order │
                    │ Build email │
                    │ Send via    │
                    │ SMTP/SES    │
                    └─────────────┘
```

**Components Needed:**

1. **Message Queue** (RabbitMQ, SQS, or Redis)
   ```python
   # Queue configuration
   queue_name = "email-notifications"
   ```

2. **Producer** (Checkout Service)
   ```python
   def checkout(order):
       payment = process_payment(order)  # Still sync
       save_order(order)

       # Publish to queue (fast, ~1ms)
       queue.publish("email-notifications", {
           "type": "order_confirmation",
           "order_id": order.id,
           "user_email": order.user.email
       })

       return {"status": "confirmed", "order_id": order.id}
   ```

3. **Consumer** (Email Worker)
   ```python
   def process_email(message):
       order = db.get_order(message["order_id"])
       email_content = render_template("order_confirm", order)

       try:
           send_email(message["user_email"], email_content)
           acknowledge(message)
       except Exception:
           # Will retry automatically
           raise
   ```

4. **Dead Letter Queue** for failed emails
   ```
   email-notifications-dlq
   - Stores emails that failed 3+ times
   - Alert ops team to investigate
   ```

5. **Monitoring**
   - Queue depth
   - Processing rate
   - Error rate

**Result:**
- Checkout response: ~500ms (down from 2.5s)
- Emails still sent reliably (within seconds)
- Email service issues don't affect checkout
- Can scale email workers independently

</details>

**Q2:** You're using Kafka with 4 partitions for an order processing topic. You have 6 consumers in a consumer group. How many consumers are actually processing messages? How would you fix this?

<details>
<summary>View Answer</summary>

**Answer: Only 4 consumers are processing messages!**

**Why:**

```
Kafka partition assignment rule:
- Each partition → exactly ONE consumer (per consumer group)
- Each consumer → zero or more partitions

With 4 partitions and 6 consumers:

Topic: orders
┌────────────┬────────────┬────────────┬────────────┐
│Partition 0 │Partition 1 │Partition 2 │Partition 3 │
└────────────┴────────────┴────────────┴────────────┘
      │            │            │            │
      ▼            ▼            ▼            ▼
 Consumer 0   Consumer 1   Consumer 2   Consumer 3

Consumer 4: IDLE (no partition assigned)
Consumer 5: IDLE (no partition assigned)

2 consumers doing nothing!
```

**How to Fix:**

**Option 1: Increase Partitions (Recommended)**
```
Add more partitions to match consumers:

kafka-topics --alter --topic orders \
  --partitions 6 \
  --bootstrap-server localhost:9092

Now:
┌────────┬────────┬────────┬────────┬────────┬────────┐
│ Part 0 │ Part 1 │ Part 2 │ Part 3 │ Part 4 │ Part 5 │
└────────┴────────┴────────┴────────┴────────┴────────┘
    │        │        │        │        │        │
    ▼        ▼        ▼        ▼        ▼        ▼
  Con 0    Con 1    Con 2    Con 3    Con 4    Con 5

All 6 consumers active!
```

**Option 2: Reduce Consumers**
```
If 4 partitions provide enough throughput:
- Scale down to 4 consumers
- Save resources
- Each consumer handles 1 partition
```

**Option 3: Multiple Consumer Groups**
```
If you need to process same messages differently:

Consumer Group A (4 consumers): Order fulfillment
Consumer Group B (4 consumers): Analytics

Each group gets ALL messages independently.
```

**Best Practices:**
```
1. Plan partitions based on expected consumers
   - Rule of thumb: partitions = 2x to 3x expected consumers
   - Allows room to scale

2. Partitions can be increased but not decreased
   - Start with more than you need
   - Typical: 6-12 partitions for moderate traffic

3. Monitor consumer lag
   - If lag grows, add partitions + consumers
```

</details>

**Q3:** Your message consumer processes messages but occasionally fails. Without proper handling, what are two things that could go wrong? How do you prevent each?

<details>
<summary>View Answer</summary>

**Problem 1: Message Loss**

```
Scenario:
1. Consumer receives message
2. Consumer ACKs immediately (before processing)
3. Processing fails/crashes
4. Message is gone forever!

Timeline:
Queue: [msg1]
Consumer: receive(msg1) → ACK → process() → CRASH!
Queue: [] (empty - thinks msg1 was processed)
Result: msg1 LOST! ❌
```

**Prevention: ACK After Processing**
```python
# ❌ Wrong - ACK before processing
def consume(message):
    channel.basic_ack(message.delivery_tag)  # Too early!
    process(message)  # If this fails, message lost

# ✅ Correct - ACK after processing
def consume(message):
    try:
        process(message)  # Process first
        channel.basic_ack(message.delivery_tag)  # Then ACK
    except Exception:
        channel.basic_nack(message.delivery_tag, requeue=True)
```

---

**Problem 2: Infinite Retry Loop (Poison Message)**

```
Scenario:
1. Message has invalid data
2. Consumer tries to process → fails
3. Message requeued
4. Consumer tries again → fails
5. Repeat forever!

Timeline:
Queue: [bad_msg]
Consumer: process(bad_msg) → ERROR → requeue
Queue: [bad_msg]
Consumer: process(bad_msg) → ERROR → requeue
... (forever)

Result: Consumer stuck, queue blocked! ❌
```

**Prevention: Dead Letter Queue + Max Retries**
```python
MAX_RETRIES = 3

def consume(message):
    retry_count = message.headers.get('x-retry-count', 0)

    try:
        process(message)
        channel.basic_ack(message.delivery_tag)
    except Exception as e:
        if retry_count >= MAX_RETRIES:
            # Send to DLQ instead of requeuing
            send_to_dlq(message, error=str(e))
            channel.basic_ack(message.delivery_tag)
            alert_ops(f"Message {message.id} moved to DLQ")
        else:
            # Requeue with incremented retry count
            channel.basic_nack(message.delivery_tag, requeue=False)
            republish_with_delay(message, retry_count + 1)
```

**DLQ Setup:**
```
Main Queue ──────────────► Consumer
                               │
                          fails 3x
                               │
                               ▼
Dead Letter Queue ──────► Manual Review
```

---

**Summary of Both Problems:**

| Problem | Cause | Prevention |
|---------|-------|------------|
| Message Loss | ACK before processing | ACK only after success |
| Infinite Loop | No retry limit | DLQ after N retries |

</details>

**Q4:** Compare these two approaches for a notification system that sends push notifications, emails, and SMS:

**Approach A:** Three separate queues (push-queue, email-queue, sms-queue)
**Approach B:** One queue with message type field

Which is better and why?

<details>
<summary>View Answer</summary>

**Approach A: Separate Queues (Recommended)**

```
┌─────────────────────────────────────────────────────────────┐
│                    Separate Queues                          │
└─────────────────────────────────────────────────────────────┘

[Notification Service]
         │
         ├──► [Push Queue] ──► [Push Workers × 3]
         │                      Fast, high volume
         │
         ├──► [Email Queue] ──► [Email Workers × 5]
         │                      Medium speed
         │
         └──► [SMS Queue] ──► [SMS Workers × 2]
                               Slow, rate-limited
```

**Approach B: Single Queue**

```
┌─────────────────────────────────────────────────────────────┐
│                    Single Queue                              │
└─────────────────────────────────────────────────────────────┘

[Notification Service]
         │
         ▼
  [Notifications Queue]
  {type: "push", ...}
  {type: "email", ...}
  {type: "sms", ...}
         │
         ▼
  [Generic Workers × 10]
  if type == "push": send_push()
  if type == "email": send_email()
  if type == "sms": send_sms()
```

---

**Why Approach A is Better:**

**1. Independent Scaling**
```
Approach A:
- Push backed up? Add push workers only
- Email slow? Add email workers only
- SMS rate limited? Keep 2 workers (more won't help)

Approach B:
- SMS rate-limited at provider level
- Workers waiting on SMS block other message types
- Can't scale email without also scaling SMS workers
```

**2. Fault Isolation**
```
Approach A:
- Email service down → only email queue backs up
- Push and SMS continue working!

Approach B:
- Email service down → workers stuck on email messages
- Push and SMS delayed behind backed-up emails
- Or need complex skip/retry logic
```

**3. Different Processing Characteristics**
```
Push notifications:
- Very fast (10ms)
- High volume
- Fire and forget

Emails:
- Medium speed (500ms)
- Template rendering
- Attachment handling

SMS:
- Slow (1-2s)
- Strict rate limits (100/sec from provider)
- Most expensive

Separate queues let each optimize independently.
```

**4. Simpler Workers**
```
Approach A workers:
- Single responsibility
- Simple code
- Easy to test
- Clear metrics per channel

Approach B workers:
- Multiple code paths
- Complex routing logic
- Harder to debug
- Mixed metrics
```

**5. Easier Monitoring**
```
Approach A:
- "Email queue depth: 5000" → email issue
- "Push queue depth: 100" → push fine

Approach B:
- "Queue depth: 5100" → which channel?
- Need to parse messages to understand backlog
```

---

**When Approach B Might Work:**

```
✓ Very low volume (< 100 messages/day)
✓ All channels have similar characteristics
✓ Simple MVP / prototype
✓ Same team owns all channels
```

---

**Recommended Architecture:**

```
[Event: User signed up]
        │
        ▼
[Notification Router Service]
        │
        ├── Determine which channels user prefers
        ├── Determine which channels apply to event
        │
        ├──► [Push Queue]
        ├──► [Email Queue]
        └──► [SMS Queue]

Each queue:
- Own workers
- Own scaling rules
- Own DLQ
- Own monitoring
```

</details>

**Q5:** Design a message queue architecture for a video processing pipeline. When a user uploads a video, it needs to: (1) transcode to multiple resolutions, (2) generate thumbnails, (3) extract metadata, (4) scan for inappropriate content. Some of these can run in parallel, some must wait for others.

<details>
<summary>View Answer</summary>

**Analysis of Dependencies:**

```
Can run in parallel (no dependencies):
- Thumbnails ─────────────┐
- Metadata extraction ────┤── All just need raw video
- Content scan ───────────┘

Must wait:
- Transcode → needs raw video first
- Publish → needs transcode + content scan complete
```

**Architecture:**

```
┌─────────────────────────────────────────────────────────────┐
│              Video Processing Pipeline                       │
└─────────────────────────────────────────────────────────────┘

[Upload Service]
      │
      │ video_id: abc123
      │ s3_path: /raw/abc123.mp4
      ▼
[Video Uploaded Event] ──► [Video Events Topic (Kafka)]
                                    │
        ┌───────────────────────────┼───────────────────────┐
        │                           │                       │
        ▼                           ▼                       ▼
[Thumbnail Queue]          [Metadata Queue]         [Content Scan Queue]
        │                           │                       │
        ▼                           ▼                       ▼
[Thumbnail Worker]         [Metadata Worker]        [Content Scan Worker]
  - Extract frames           - Duration              - AI moderation
  - Generate 5 thumbs        - Resolution            - Flag if unsafe
  - Store in S3              - Codec info            - Store result
        │                           │                       │
        ▼                           ▼                       ▼
[Thumbnails Ready]         [Metadata Ready]         [Content Scanned]
        │                           │                       │
        └───────────────────────────┼───────────────────────┘
                                    │
                                    ▼
                        [Orchestrator Service]
                        Waits for all 3 events
                        per video_id
                                    │
                         When all complete:
                                    │
                    ┌───────────────┴───────────────┐
                    │                               │
                    ▼                               ▼
            [Transcode Queue]               [If content unsafe]
                    │                               │
                    ▼                               ▼
            [Transcode Worker]              [Rejection Queue]
              - 360p                                │
              - 720p                                ▼
              - 1080p                       [Notify Uploader]
              - 4K                          "Video rejected"
                    │
                    ▼
            [Transcoding Complete]
                    │
                    ▼
            [Publish Queue]
                    │
                    ▼
            [Publish Worker]
              - Update DB status
              - Enable playback
              - Notify uploader
                    │
                    ▼
            [Video Live!]
```

**Message Schemas:**

```json
// Video Uploaded Event
{
  "event": "video.uploaded",
  "video_id": "abc123",
  "user_id": "user_456",
  "s3_path": "s3://uploads/raw/abc123.mp4",
  "timestamp": "2024-01-15T10:30:00Z"
}

// Thumbnail Message
{
  "video_id": "abc123",
  "s3_path": "s3://uploads/raw/abc123.mp4",
  "output_path": "s3://thumbnails/abc123/"
}

// Processing Complete Event
{
  "event": "thumbnail.complete",
  "video_id": "abc123",
  "thumbnails": [
    "s3://thumbnails/abc123/thumb_0.jpg",
    "s3://thumbnails/abc123/thumb_1.jpg"
  ]
}
```

**Orchestrator Implementation:**

```python
# Using Redis to track completion status
class VideoOrchestrator:
    def __init__(self):
        self.redis = Redis()
        self.required_steps = ['thumbnail', 'metadata', 'content_scan']

    def on_step_complete(self, video_id, step, result):
        # Store result
        self.redis.hset(f"video:{video_id}", step, json.dumps(result))

        # Check if all steps complete
        status = self.redis.hgetall(f"video:{video_id}")

        if all(step in status for step in self.required_steps):
            # Check content scan result
            scan_result = json.loads(status['content_scan'])

            if scan_result['safe']:
                # All good - proceed to transcode
                queue.publish("transcode", {
                    "video_id": video_id,
                    "metadata": json.loads(status['metadata'])
                })
            else:
                # Content rejected
                queue.publish("rejection", {
                    "video_id": video_id,
                    "reason": scan_result['reason']
                })
```

**Queue Configuration:**

| Queue | Workers | Notes |
|-------|---------|-------|
| thumbnail | 10 | Fast, CPU-bound |
| metadata | 5 | Very fast |
| content_scan | 20 | GPU workers, slower |
| transcode | 15 | Slow, CPU/GPU heavy |
| publish | 5 | Fast DB updates |

**Monitoring:**

```
Key Metrics:
- Videos pending per stage
- Average processing time per stage
- Content rejection rate
- End-to-end latency (upload → live)

Alerts:
- Transcode queue > 1000 (need more workers)
- Content scan queue > 500 (GPU capacity)
- Any DLQ has messages (failures need attention)
```

</details>

**Q6:** Your at-least-once message queue occasionally delivers duplicates. Design an idempotent consumer for a payment processing system that charges customers.

<details>
<summary>View Answer</summary>

**The Problem:**

```
Message: {"payment_id": "pay_123", "amount": 100, "user": "alice"}

Without idempotency:
1. Consumer receives message
2. Charges customer $100
3. ACK fails (network issue)
4. Queue redelivers message
5. Consumer charges customer $100 AGAIN!
6. Alice charged $200 😱
```

**Solution: Idempotent Payment Consumer**

```python
import redis
import psycopg2
from contextlib import contextmanager

class IdempotentPaymentConsumer:
    def __init__(self):
        self.redis = redis.Redis()
        self.db = psycopg2.connect(...)

        # Idempotency key TTL (7 days)
        self.idempotency_ttl = 7 * 24 * 60 * 60

    def process_payment(self, message):
        payment_id = message['payment_id']
        idempotency_key = f"payment:processed:{payment_id}"

        # Step 1: Check if already processed (fast Redis check)
        if self.redis.exists(idempotency_key):
            logger.info(f"Payment {payment_id} already processed, skipping")
            return {"status": "duplicate", "payment_id": payment_id}

        # Step 2: Try to acquire processing lock
        lock_key = f"payment:lock:{payment_id}"
        lock_acquired = self.redis.set(
            lock_key,
            "processing",
            nx=True,  # Only if not exists
            ex=300    # 5 minute lock timeout
        )

        if not lock_acquired:
            logger.info(f"Payment {payment_id} being processed by another worker")
            raise RetryLater()  # Let queue redeliver later

        try:
            # Step 3: Check database (source of truth)
            existing = self.db.execute(
                "SELECT status FROM payments WHERE payment_id = %s",
                (payment_id,)
            ).fetchone()

            if existing:
                logger.info(f"Payment {payment_id} exists in DB: {existing.status}")
                self.redis.setex(idempotency_key, self.idempotency_ttl, "done")
                return {"status": "duplicate", "payment_id": payment_id}

            # Step 4: Process payment atomically
            with self.db_transaction() as txn:
                # Insert payment record FIRST (before charging)
                txn.execute("""
                    INSERT INTO payments (payment_id, user_id, amount, status)
                    VALUES (%s, %s, %s, 'pending')
                """, (payment_id, message['user_id'], message['amount']))

                # Charge payment provider
                charge_result = self.charge_provider(message)

                if charge_result['success']:
                    txn.execute("""
                        UPDATE payments
                        SET status = 'completed',
                            provider_ref = %s
                        WHERE payment_id = %s
                    """, (charge_result['reference'], payment_id))
                else:
                    txn.execute("""
                        UPDATE payments
                        SET status = 'failed',
                            error = %s
                        WHERE payment_id = %s
                    """, (charge_result['error'], payment_id))
                    raise PaymentFailed(charge_result['error'])

            # Step 5: Mark as processed in Redis
            self.redis.setex(idempotency_key, self.idempotency_ttl, "done")

            return {"status": "success", "payment_id": payment_id}

        finally:
            # Always release lock
            self.redis.delete(lock_key)

    def charge_provider(self, message):
        """
        Call payment provider with THEIR idempotency key
        Most providers (Stripe, etc.) support this
        """
        return stripe.PaymentIntent.create(
            amount=message['amount'],
            currency='usd',
            customer=message['user_id'],
            idempotency_key=message['payment_id']  # Provider-level protection!
        )

    @contextmanager
    def db_transaction(self):
        cursor = self.db.cursor()
        try:
            yield cursor
            self.db.commit()
        except Exception:
            self.db.rollback()
            raise
        finally:
            cursor.close()
```

**Multiple Layers of Protection:**

```
Layer 1: Redis check (fast, handles 99% of duplicates)
         ↓ not found
Layer 2: Distributed lock (prevents parallel processing)
         ↓ acquired
Layer 3: Database check (source of truth)
         ↓ not exists
Layer 4: Database INSERT before charging (crash recovery)
         ↓ inserted
Layer 5: Payment provider idempotency key (external protection)
         ↓ charged
Layer 6: Mark complete in Redis (future duplicate prevention)
```

**Handling Edge Cases:**

```
Crash after INSERT, before charge:
→ Payment record exists with status='pending'
→ Reconciliation job finds pending payments
→ Either completes or refunds

Crash after charge, before UPDATE:
→ Provider charged the money
→ Payment record exists with status='pending'
→ Reconciliation checks provider, updates status

Network timeout from provider:
→ Don't know if charged or not!
→ Payment record exists with status='pending'
→ Query provider using payment_id
→ Update accordingly
```

**Database Schema:**

```sql
CREATE TABLE payments (
    payment_id VARCHAR(255) PRIMARY KEY,  -- Idempotency key
    user_id VARCHAR(255) NOT NULL,
    amount INTEGER NOT NULL,
    status VARCHAR(50) NOT NULL,  -- pending, completed, failed
    provider_ref VARCHAR(255),
    error TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),

    -- Prevent double-charging at DB level
    CONSTRAINT unique_payment UNIQUE (payment_id)
);

CREATE INDEX idx_payments_status ON payments(status);
CREATE INDEX idx_payments_user ON payments(user_id);
```

</details>

**Q7:** RabbitMQ vs Kafka: A startup is building a real-time analytics platform that processes 100,000 events per second. They also need to replay events from the past week for debugging. Which would you recommend and why?

<details>
<summary>View Answer</summary>

**Recommendation: Apache Kafka**

**Why Kafka:**

```
Requirement 1: 100,000 events/second
─────────────────────────────────────
RabbitMQ:
- Typical throughput: 20,000-50,000 msg/sec
- Can scale but requires more complex clustering
- Not designed for this volume

Kafka:
- Designed for millions of messages/sec
- 100K/sec is moderate load ✓
- Horizontal scaling via partitions

Winner: Kafka ✓
```

```
Requirement 2: Replay events from past week
─────────────────────────────────────────────
RabbitMQ:
- Messages deleted after consumption
- No built-in replay capability
- Would need to archive to separate storage
- Then replay = re-publish everything

Kafka:
- Persistent log by design
- Configurable retention (set to 7+ days) ✓
- Consumer can "seek" to any offset
- Replay = reset consumer offset, re-read

Example replay:
kafka-consumer-groups --reset-offsets \
  --group analytics-processor \
  --topic events \
  --to-datetime 2024-01-08T00:00:00.000 \
  --execute

Winner: Kafka ✓
```

**Kafka Architecture for This Use Case:**

```
┌─────────────────────────────────────────────────────────────┐
│                Real-time Analytics Platform                  │
└─────────────────────────────────────────────────────────────┘

Event Sources ──┬──► [Kafka Cluster]
(100K events/s) │    │
                │    │ Topic: events
                │    │ Partitions: 20
                │    │ Retention: 7 days
                │    │ Replication: 3
                │    │
                │    ├──► [Consumer Group: real-time]
                │    │    20 consumers (1 per partition)
                │    │    Process events → Dashboard
                │    │
                │    ├──► [Consumer Group: storage]
                │    │    Write to Data Lake
                │    │
                │    └──► [Consumer Group: debug] (on-demand)
                │         Reset offset → replay events
                │
                └──► [Schema Registry]
                     Ensure event format consistency
```

**Configuration:**

```properties
# Topic configuration
num.partitions=20
retention.ms=604800000  # 7 days
replication.factor=3

# Producer configuration
acks=all                # Durability
batch.size=32768        # Batch for throughput
linger.ms=5             # Wait for batches

# Consumer configuration
max.poll.records=1000   # Process in batches
enable.auto.commit=false # Manual commit for reliability
```

**Capacity Planning:**

```
100,000 events/sec × 1KB average = 100 MB/sec

Storage for 7 days:
100 MB/sec × 86400 sec × 7 days = ~60 TB

With replication factor 3:
60 TB × 3 = 180 TB total storage

Kafka cluster:
- 6 brokers
- 30 TB SSD each
- 32 GB RAM each
- 10 Gbps network
```

**When RabbitMQ Would Be Better:**

```
✓ Lower throughput (< 50K/sec)
✓ Complex routing requirements
✓ Request-reply patterns
✓ Don't need replay
✓ Priority queues
✓ Simpler operations

Example: Order processing system
- 1,000 orders/minute
- Complex routing (by region, type)
- Don't need to replay orders
→ RabbitMQ is fine
```

**Summary:**

| Requirement | RabbitMQ | Kafka |
|-------------|----------|-------|
| 100K events/sec | Challenging | Easy ✓ |
| 7-day replay | Not supported | Built-in ✓ |
| Real-time processing | Yes | Yes |
| Operational complexity | Lower | Higher |

**Final Answer:** Kafka is the clear choice for high-throughput analytics with replay requirements.

</details>

---

## Next Up

In Lesson 12, we'll explore the **CAP Theorem** - the fundamental trade-off in distributed systems between Consistency, Availability, and Partition Tolerance. Understanding CAP is essential for making informed architectural decisions!
