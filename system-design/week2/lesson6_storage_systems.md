# Week 2, Lesson 6: Storage Systems - Files, Blocks, and Objects

## Table of Contents
- [Media Resources](#media-resources)
- [The Storage Hierarchy](#the-storage-hierarchy)
- [The Three Main Storage Types](#the-three-main-storage-types)
- [Part 1: File Storage](#part-1-file-storage)
  - [What is File Storage?](#what-is-file-storage)
  - [How File Storage Works](#how-file-storage-works)
  - [File Storage Protocols](#file-storage-protocols)
  - [File Storage Use Cases](#file-storage-use-cases)
  - [File Storage in the Cloud](#file-storage-in-the-cloud)
- [Part 2: Block Storage](#part-2-block-storage)
  - [What is Block Storage?](#what-is-block-storage)
  - [How Block Storage Works](#how-block-storage-works)
  - [Block Storage Types](#block-storage-types)
  - [Block Storage Features](#block-storage-features)
  - [Block Storage Use Cases](#block-storage-use-cases)
- [Part 3: Object Storage](#part-3-object-storage)
  - [What is Object Storage?](#what-is-object-storage)
  - [Object Storage Structure](#object-storage-structure)
  - [How Object Storage Works](#how-object-storage-works)
  - [Object Storage Architecture](#object-storage-architecture)
  - [Object Storage Features](#object-storage-features)
  - [Object Storage Use Cases](#object-storage-use-cases)
  - [Popular Object Storage Services](#popular-object-storage-services)
- [Comparing Storage Types](#comparing-storage-types)
- [Real-World Architecture Examples](#real-world-architecture-examples)
- [Storage Security](#storage-security)
- [Storage Best Practices](#storage-best-practices)
- [Hybrid Storage Solutions](#hybrid-storage-solutions)
- [Key Concepts to Remember](#key-concepts-to-remember)
- [Practice Questions](#practice-questions)
- [Next Up](#next-up)

---

Welcome to the final lesson of Week 2! You've learned about networking and caching—now let's talk about where and how we actually store data long-term.

## Media Resources

**Visual Guide:**
![Choosing Your Storage: A Guide to File, Block, and Object](./assets/storage_systems_infographic.png)

**Audio Lesson:**
[File, Block and Object Storage Explained (Audio)](./assets/storage_systems.m4a)

---

## The Storage Hierarchy

Before diving into storage types, understand the hierarchy:

```
Speed ←→ Cost ←→ Capacity

Fastest & Most Expensive:
├─ CPU Registers (nanoseconds, bytes)
├─ L1/L2/L3 Cache (nanoseconds, KB-MB)
├─ RAM (100ns, GB)
├─ SSD (microseconds, TB)
├─ HDD (milliseconds, TB)
├─ Network Storage (milliseconds, PB)
└─ Tape/Archive (seconds-minutes, PB)
Slowest & Cheapest
```

This lesson focuses on **persistent storage** (SSD, HDD, network storage).

---

## The Three Main Storage Types

| Type | Analogy | Best For |
|------|---------|----------|
| **File Storage** | Filing cabinet with folders | Traditional files, shared drives |
| **Block Storage** | Raw notebook pages | Databases, VMs, high performance |
| **Object Storage** | Warehouse with labeled boxes | Cloud files, backups, media |

Let's explore each in detail!

---

## Part 1: File Storage

### What is File Storage?

**File storage organizes data in hierarchical directories (folders).**

Think: Your computer's file system!

```
/
├── home/
│   ├── alice/
│   │   ├── documents/
│   │   │   └── report.pdf
│   │   └── photos/
│   │       └── vacation.jpg
│   └── bob/
│       └── projects/
│           └── code.py
└── var/
    └── logs/
        └── app.log
```

### How File Storage Works

```
1. Client makes request: "Read /home/alice/documents/report.pdf"
2. File system looks up path
3. Finds file metadata (location, permissions, size)
4. Reads file from disk
5. Returns file content
```

### File Storage Protocols

#### NFS (Network File System)
```
Linux/Unix shared file system

Server exports directory:
/mnt/shared

Clients mount it:
mount server:/mnt/shared /local/shared

Now clients can access files as if local!
```

**Pros:**
- ✅ POSIX compliant (works like local filesystem)
- ✅ Good for Unix/Linux

**Cons:**
- ❌ Not great for Windows
- ❌ Performance degrades with many small files

#### SMB/CIFS (Server Message Block)
```
Windows file sharing protocol

Share on Windows:
\\server\shared\documents\

Mac/Linux can also connect
```

**Pros:**
- ✅ Native Windows support
- ✅ Good for mixed environments

**Cons:**
- ❌ Complex protocol
- ❌ Security concerns (older versions)

### File Storage Use Cases

**Perfect for:**
- ✅ Shared team folders
- ✅ Home directories
- ✅ Content management systems
- ✅ Development environments
- ✅ Traditional applications expecting filesystem

**Not ideal for:**
- ❌ Massive scale (billions of files)
- ❌ High-performance databases
- ❌ Unstructured data at scale
- ❌ Global distribution

### File Storage in the Cloud

**AWS EFS (Elastic File System)**
```
Fully managed NFS file system
Auto-scaling (no capacity planning)
Multi-AZ availability
Expensive but convenient
```

**Azure Files**
```
SMB file shares in cloud
Integrates with Active Directory
Good for lift-and-shift scenarios
```

**Google Filestore**
```
Managed NFS service
High performance
Good for GKE (Kubernetes)
```

---

## Part 2: Block Storage

### What is Block Storage?

**Block storage provides raw storage volumes (like virtual hard drives).**

Think: A blank hard drive you can format however you want!

```
Block Storage Volume = Array of fixed-size blocks

[Block 0] [Block 1] [Block 2] [Block 3] ... [Block N]
  4KB       4KB       4KB       4KB           4KB

You decide:
- File system (ext4, NTFS, XFS)
- Partition scheme
- Formatting
```

### How Block Storage Works

```
Application → OS → File System → Block Device → Physical Storage

Example: Database writes data
1. Database: "Write 8KB to address 0x1000"
2. OS: Translates to blocks 4 and 5
3. Block device: Writes to physical disk
4. Returns success

Block storage doesn't know about files!
It just reads/writes fixed-size blocks.
```

### Block Storage Types

#### Direct Attached Storage (DAS)
```
Disk directly connected to server

Server
  ├─ SSD 1 (500GB)
  ├─ SSD 2 (1TB)
  └─ HDD (4TB)

Pros: Fast, low latency
Cons: Not shareable, not resilient
```

#### SAN (Storage Area Network)
```
High-speed network of storage devices

[Server 1]  ┐
[Server 2]  ├─> [Fiber Channel Network] ─> [Storage Array]
[Server 3]  ┘                               ├─ Disk 1
                                            ├─ Disk 2
                                            └─ Disk N

Multiple servers access shared block storage
```

**Pros:**
- ✅ High performance
- ✅ Centralized management
- ✅ Shared storage pool

**Cons:**
- ❌ Expensive
- ❌ Complex setup
- ❌ Requires specialized hardware

#### Cloud Block Storage
```
Virtual block devices in cloud

AWS EBS (Elastic Block Store)
Azure Managed Disks
Google Persistent Disks

Attach to VMs like local disks!
```

### Block Storage Features

#### 1. IOPS (Input/Output Operations Per Second)
```
Measure of performance

Standard SSD: 3,000 IOPS
High-performance SSD: 16,000+ IOPS
NVMe SSD: 100,000+ IOPS

Higher IOPS = faster database performance!
```

#### 2. Snapshots
```
Point-in-time backups

Original Volume:
[Data at 10:00 AM]

Snapshot 1 (10:00 AM):
[Copy of all blocks]

Snapshot 2 (11:00 AM):
[Only changed blocks since 10:00]

Incremental snapshots save space!
```

#### 3. Replication
```
Multi-AZ replication:

Primary Volume (AZ-1)
└─> Synchronously replicates to
    Secondary Volume (AZ-2)

If AZ-1 fails, use AZ-2!
```

### Block Storage Use Cases

**Perfect for:**
- ✅ Databases (MySQL, PostgreSQL, MongoDB)
- ✅ Virtual machines (boot drives)
- ✅ High-performance computing
- ✅ Enterprise applications
- ✅ Anything needing low latency

**Not ideal for:**
- ❌ File sharing across many servers
- ❌ Unstructured data (images, videos)
- ❌ Object metadata and versioning
- ❌ HTTP/REST access

### Real Example: Database on Block Storage

```
[Application Server]
        ↓
[PostgreSQL Database]
        ↓
[EBS Volume - 1TB, 10,000 IOPS]
        ↓ (replication)
[EBS Snapshot - S3]

Why block storage?
- Low latency (<1ms)
- High IOPS for transactions
- Consistent performance
- Point-in-time snapshots
```

---

## Part 3: Object Storage

### What is Object Storage?

**Object storage stores data as objects (files + metadata) in a flat namespace.**

Think: Amazon S3, Google Cloud Storage!

```
Not hierarchical like files:
❌ /photos/2024/vacation/img1.jpg

Flat with unique keys:
✅ bucket-name/photos-2024-vacation-img1.jpg

Each object has:
- Data (the file contents)
- Metadata (content-type, custom tags)
- Unique identifier (key)
```

### Object Storage Structure

```
Object = Data + Metadata + Unique Key

Example object:
{
  "bucket": "my-photos",
  "key": "vacation/beach.jpg",
  "data": [binary image data],
  "metadata": {
    "content-type": "image/jpeg",
    "size": 2048576,
    "last-modified": "2024-12-26T10:30:00Z",
    "custom-tag": "vacation-2024",
    "owner": "alice"
  }
}
```

### How Object Storage Works

```
HTTP REST API (not file system!)

Upload:
PUT /bucket/my-file.jpg
Content-Type: image/jpeg
[binary data]

Download:
GET /bucket/my-file.jpg

List:
GET /bucket?prefix=photos/

Delete:
DELETE /bucket/my-file.jpg
```

### Object Storage Architecture

```
                [Load Balancer]
                       │
        ┌──────────────┼──────────────┐
        ↓              ↓              ↓
   [API Server]   [API Server]   [API Server]
        │              │              │
        └──────────────┼──────────────┘
                       ↓
              [Metadata Store]
            (which node has what)
                       ↓
        ┌──────────────┼──────────────┐
        ↓              ↓              ↓
   [Storage Node] [Storage Node] [Storage Node]
   [Replication]  [Replication]  [Replication]

Objects distributed across many nodes
Automatically replicated (3+ copies)
Highly scalable!
```

### Object Storage Features

#### 1. Unlimited Scale
```
Store billions of objects
Petabytes or exabytes of data
No capacity planning needed
Pay only for what you use
```

#### 2. Durability (99.999999999% - "11 nines")
```
S3 automatically replicates objects across:
- Multiple devices
- Multiple facilities
- Multiple availability zones

Probability of losing an object: 0.000000001%
```

#### 3. Versioning
```
my-document.pdf (version 1)
my-document.pdf (version 2)
my-document.pdf (version 3)

Delete = just marks as deleted
Can restore previous versions!
```

#### 4. Access Control
```
Bucket policies:
- Public read
- Private (owner only)
- Specific users/roles
- IP-based restrictions
- Time-based URLs (pre-signed URLs)
```

#### 5. Lifecycle Policies
```
Automatic data management:

Rule 1: After 30 days → Move to cheaper storage
Rule 2: After 90 days → Archive to Glacier
Rule 3: After 365 days → Delete

Saves money automatically!
```

#### 6. Event Notifications
```
When object uploaded:
└─> Trigger Lambda function
    └─> Process image (resize, compress)
    └─> Update database
    └─> Send notification

Event-driven workflows!
```

### Object Storage Use Cases

**Perfect for:**
- ✅ Static website hosting
- ✅ Media files (images, videos, audio)
- ✅ Backups and archives
- ✅ Big data analytics
- ✅ Data lakes
- ✅ Content distribution (with CDN)
- ✅ Application data storage
- ✅ Log storage
- ✅ Machine learning datasets

**Not ideal for:**
- ❌ Databases (too slow)
- ❌ File systems (no POSIX)
- ❌ Frequent small updates
- ❌ Low-latency requirements

### Popular Object Storage Services

#### AWS S3 (Simple Storage Service)
```
Most popular object storage

Storage classes:
- S3 Standard: Frequent access
- S3 IA: Infrequent access (cheaper)
- S3 Glacier: Archive (very cheap)
- S3 Intelligent-Tiering: Automatic

Features:
- Versioning
- Encryption
- Access logging
- Cross-region replication
- Event notifications
```

#### Google Cloud Storage
```
Similar to S3

Storage classes:
- Standard
- Nearline (30-day access)
- Coldline (90-day access)
- Archive (365-day access)

Integration with BigQuery, ML
```

#### Azure Blob Storage
```
Microsoft's object storage

Blob types:
- Block blobs (general purpose)
- Append blobs (logs)
- Page blobs (VM disks)

Good Windows integration
```

#### MinIO
```
Open-source, S3-compatible

Self-hosted option
Run on-premises or any cloud
Kubernetes-native
```

---

## Comparing Storage Types

### Performance Comparison

```
Operation          File Storage    Block Storage    Object Storage
─────────────────────────────────────────────────────────────────
Small random reads      Good          Excellent          Poor
Large sequential reads  Good          Excellent          Good
Metadata operations     Good          Poor               Excellent
Concurrent access       Good          Poor (locking)     Excellent
Latency                 ~10ms         ~1ms               ~50-100ms
```

### Cost Comparison (rough estimates)

```
Block Storage: $0.10/GB/month (high performance)
File Storage:  $0.08/GB/month (managed NFS)
Object Storage: $0.02/GB/month (standard)
Object Archive: $0.004/GB/month (glacier)

Object storage is ~5x cheaper!
```

### Scalability Comparison

```
File Storage:
└─> Limited (millions of files)
└─> Complex to scale horizontally

Block Storage:
└─> Limited to server capacity
└─> Need SAN for sharing

Object Storage:
└─> Virtually unlimited
└─> Built for horizontal scale
└─> Billions of objects, no problem
```

### Use Case Matrix

| Requirement | File | Block | Object |
|-------------|------|-------|--------|
| Database | ❌ | ✅ | ❌ |
| VM boot disk | ❌ | ✅ | ❌ |
| Shared documents | ✅ | ❌ | ⚠️ |
| Media files | ⚠️ | ❌ | ✅ |
| Backups | ⚠️ | ⚠️ | ✅ |
| Big data | ❌ | ❌ | ✅ |
| Website assets | ❌ | ❌ | ✅ |
| Low latency | ⚠️ | ✅ | ❌ |
| Massive scale | ❌ | ❌ | ✅ |

---

## Real-World Architecture Examples

### Example 1: Instagram

```
User uploads photo:

1. [App] → Upload to [S3 - Object Storage]
   └─> Original: photos/original/abc123.jpg
   
2. [Lambda] Triggered by S3 event
   └─> Resize to multiple sizes
   └─> Store: photos/thumb/abc123.jpg
   └─> Store: photos/medium/abc123.jpg
   └─> Store: photos/large/abc123.jpg
   
3. [CloudFront CDN] Caches from S3
   
4. [Database - Block Storage]
   └─> Store metadata: photo_id, user_id, S3_urls
   
5. User views photo:
   └─> CDN serves from cache (fast!)
```

**Why this design?**
- Object storage: Cheap, scalable for billions of photos
- Block storage: Fast database queries
- CDN: Fast global delivery

### Example 2: Netflix

```
Video Streaming Architecture:

[Content Creation]
└─> [S3 - Master Videos]
    └─> Trigger encoding pipeline
    └─> Multiple formats (1080p, 720p, 4K)
    └─> Store in [S3 - Encoded Videos]
    └─> Distribute to [CDN]

[User Watches]
└─> CDN serves video chunks
└─> Buffering from edge servers
└─> No S3 access during playback!

[Metadata/User Data]
└─> [Cassandra - Block Storage]
└─> Watch history, preferences
```

### Example 3: Dropbox

```
File Sync Architecture:

[User uploads file]
└─> Break into 4MB blocks
└─> Hash each block (SHA-256)
└─> Check if block exists (deduplication)
    ├─> Exists: Reference existing block
    └─> New: Upload to [S3]

[Metadata - Block Storage Database]
└─> File structure, permissions, versions
└─> Which blocks belong to which files

[Sync to other devices]
└─> Download only changed blocks
└─> Efficient sync!

Storage breakdown:
- File blocks: S3 (object storage)
- Metadata: PostgreSQL on EBS (block storage)
- Cache: Redis on EBS (block storage)
```

---

## Storage Security

### Encryption

#### 1. Encryption at Rest
```
Encrypt data on disk

Block Storage:
- Full disk encryption (LUKS, BitLocker)
- Volume encryption (AWS KMS)

Object Storage:
- Server-side encryption (SSE-S3, SSE-KMS)
- Client-side encryption (encrypt before upload)
```

#### 2. Encryption in Transit
```
TLS/SSL for data transfer

HTTPS for object storage
Encrypted block device protocols
VPN for file storage
```

### Access Control

```
Block Storage:
└─> OS-level permissions
└─> Volume-level access control

File Storage:
└─> POSIX permissions (chmod, chown)
└─> ACLs (Access Control Lists)
└─> Active Directory integration

Object Storage:
└─> IAM policies
└─> Bucket policies
└─> Pre-signed URLs (temporary access)
└─> Access logging
```

---

## Storage Best Practices

### 1. Choose the Right Storage Type

```
Decision tree:

Need low latency database?
└─> Block Storage (EBS, Persistent Disk)

Need shared file system?
└─> File Storage (EFS, Azure Files)

Need to store millions of files?
└─> Object Storage (S3, GCS)

Need all three?
└─> Use all three! (common in practice)
```

### 2. Plan for Redundancy

```
Block Storage:
- Multi-AZ replication
- Regular snapshots
- Test restores!

Object Storage:
- Already replicated (11 nines)
- Enable versioning
- Lifecycle policies

File Storage:
- Backup to object storage
- Multi-AZ deployment
```

### 3. Monitor Storage Metrics

```
Key metrics:

Block Storage:
- IOPS utilization
- Throughput
- Latency
- Snapshot age

Object Storage:
- Request rates
- Bandwidth
- Error rates
- Storage costs

File Storage:
- Throughput
- Client connections
- File count
```

### 4. Optimize Costs

```
Object Storage strategies:
- Use lifecycle policies
- Delete old versions
- Archive cold data
- Use intelligent tiering

Block Storage strategies:
- Right-size volumes
- Delete unused snapshots
- Use appropriate IOPS tier

File Storage strategies:
- Use appropriate tier
- Enable compression
- Archive old files
```

---

## Hybrid Storage Solutions

Modern systems often combine storage types:

```
Typical Web Application:

[Application Servers]
├─> [Block Storage] - OS, application code
├─> [File Storage] - Shared logs, uploads
└─> [Object Storage] - User uploads, backups

[Database Servers]
└─> [Block Storage] - Database files, high IOPS

[Static Assets]
└─> [Object Storage + CDN] - Images, CSS, JS

[Backups]
└─> [Object Storage - Glacier] - Long-term archives
```

---

## Key Concepts to Remember

1. **Three storage types**: File (hierarchical), Block (raw volumes), Object (flat namespace)
2. **File storage**: Like your computer's filesystem, good for sharing
3. **Block storage**: Raw disk, best for databases and VMs, high performance
4. **Object storage**: Unlimited scale, cheap, best for unstructured data
5. **Object storage durability**: 11 nines (99.999999999%)
6. **Block storage = fast but limited**; Object storage = slow but unlimited
7. **Most systems use multiple storage types** for different needs
8. **Object storage accessed via HTTP/REST**, not filesystem
9. **Security**: Always encrypt at rest and in transit
10. **Cost optimization**: Use lifecycle policies and right storage tier

---

## Practice Questions

**Q1:** You're building a video streaming platform like YouTube. Design the storage architecture:
- Where do you store uploaded videos?
- Where do you store video metadata (title, views, etc.)?
- Where do you store thumbnails?
- How do you serve videos globally?

Explain your choice for each component.

<details>
<summary>View Answer</summary>

| Component | Storage | Reasoning |
|-----------|---------|-----------|
| **Raw uploaded videos** | S3 (object storage) | Massive files, write-once, cost-effective |
| **Transcoded videos** | S3 + CDN | Multiple resolutions, served via CloudFront/CDN |
| **Video metadata** | PostgreSQL/MySQL (block storage) | Structured data, complex queries, ACID |
| **Thumbnails** | S3 + CDN | Small images, millions of files, HTTP access |
| **View counts** | Redis + async DB sync | High-frequency updates, eventual consistency OK |

**Architecture:**
```
Upload: User → S3 → Lambda trigger → Transcoding queue → S3 (multiple resolutions)
Serve: User → CDN → S3 (cache miss) or Edge (cache hit)
Metadata: App → PostgreSQL
```

**Global serving:** CDN with edge locations worldwide. Videos cached at edges closest to users.

</details>

**Q2:** Compare storage solutions for these scenarios:

a) A database handling 10,000 transactions/second
b) A backup system for 100TB of files
c) A shared development environment for 50 developers

Which storage type (file, block, or object) for each? Why?

<details>
<summary>View Answer</summary>

| Scenario | Storage Type | Solution | Reasoning |
|----------|--------------|----------|-----------|
| **a) 10K TPS database** | Block | AWS EBS io2 or local NVMe | Low latency, random I/O, IOPS provisioning |
| **b) 100TB backup** | Object | S3 Glacier | Cheapest for archival, durability, no IOPS needed |
| **c) 50 dev shared env** | File | AWS EFS or FSx | Multiple concurrent access, POSIX compatibility |

**Detailed reasoning:**

a) **Block storage** because databases need:
- Random read/write at block level
- Consistent low latency (<1ms)
- High IOPS (can provision 64,000+ IOPS)

b) **Object storage** because backups need:
- Massive capacity at low cost ($0.004/GB/month for Glacier)
- 11 nines durability
- No random access needed

c) **File storage** because developers need:
- Shared access to same files
- POSIX semantics (locks, permissions)
- Mount as network drive

</details>

**Q3:** Your S3 bill is $10,000/month for 500TB of data. 80% of files haven't been accessed in 90 days. How can you reduce costs?

<details>
<summary>View Answer</summary>

**Current cost:** ~$0.023/GB × 500,000 GB = ~$11,500 (S3 Standard)

**Strategy: Implement S3 Lifecycle policies**

```json
{
  "Rules": [
    {
      "ID": "ArchiveOldFiles",
      "Filter": {},
      "Status": "Enabled",
      "Transitions": [
        {"Days": 30, "StorageClass": "STANDARD_IA"},
        {"Days": 90, "StorageClass": "GLACIER"},
        {"Days": 365, "StorageClass": "DEEP_ARCHIVE"}
      ]
    }
  ]
}
```

**New cost calculation:**
- 100TB (20%) in Standard: $2,300
- 200TB in Glacier: $800 ($0.004/GB)
- 200TB in Deep Archive: $400 ($0.002/GB)
- **Total: ~$3,500/month (65% savings!)**

**Additional optimizations:**
1. Delete truly unused data
2. Enable S3 Intelligent-Tiering for unpredictable access
3. Compress files before storage
4. Review if all data needs to be kept

</details>

**Q4:** You need to store user profile pictures for a social network with 100 million users. Should you:
a) Store as files in EFS (file storage)?
b) Store as BLOBs in PostgreSQL on EBS (block storage)?
c) Store in S3 (object storage)?

Justify your answer with scalability, cost, and performance considerations.

<details>
<summary>View Answer</summary>

**Answer: c) S3 (object storage)**

| Option | Scalability | Cost | Performance | Verdict |
|--------|-------------|------|-------------|---------|
| **EFS** | Good | $300/TB/mo | Good | ❌ Too expensive |
| **PostgreSQL BLOBs** | Poor | Medium | Terrible | ❌ Database not for images |
| **S3** | Unlimited | $23/TB/mo | Excellent via CDN | ✅ Best choice |

**Why S3 wins:**

1. **Scalability:** Unlimited storage, no provisioning
2. **Cost:** 100M users × 1MB avg = 100TB = $2,300/month
3. **Performance:** Serve via CloudFront CDN, sub-50ms globally
4. **Simplicity:** Direct URL access, no app server needed for serving

**Why NOT the others:**

- **EFS:** 13x more expensive ($30,000/month for 100TB), designed for shared file systems
- **PostgreSQL:** Database bloat, backup nightmares, terrible query performance with BLOBs, no CDN integration

**Architecture:**
```
Upload: App → S3 (generate unique key)
Serve: Browser → CloudFront CDN → S3
Database: Only store S3 key (e.g., "users/123/profile.jpg")
```

</details>

**Q5:** Explain why object storage (S3) is terrible for databases, but databases on block storage (EBS) are great. What's the fundamental difference?

<details>
<summary>View Answer</summary>

**The fundamental difference: Access patterns**

| Aspect | Block Storage | Object Storage |
|--------|---------------|----------------|
| **Access unit** | 4KB-64KB blocks | Entire objects (KB-TB) |
| **Random I/O** | Excellent | Not supported |
| **Latency** | <1ms | 50-200ms |
| **Update** | Modify any block | Replace entire object |
| **POSIX** | Yes (mount as disk) | No (HTTP API only) |

**Why databases need block storage:**

```
Database operation: UPDATE users SET name='Alice' WHERE id=123

What happens on disk:
1. Read 8KB page containing row
2. Modify a few bytes
3. Write 8KB page back
4. Update index (another random write)
5. Write to WAL (sequential write)

This requires: Random access, small writes, low latency
```

**Why S3 can't do this:**

```
S3 operation flow:
1. GET entire object (even if TB)
2. Modify in memory
3. PUT entire object back
4. No random access within object
5. No concurrent writes

Result: Impossibly slow for database workloads
```

**S3 latency:** 50-200ms per operation
**EBS latency:** 0.1-1ms per operation

For 10,000 TPS database:
- EBS: 10,000 × 1ms = 10 seconds of I/O time (parallelized)
- S3: 10,000 × 100ms = 1,000 seconds = impossible

</details>

**Q6:** Design a backup strategy that uses all three storage types:
- What goes in block storage?
- What goes in file storage?
- What goes in object storage?

<details>
<summary>View Answer</summary>

**Comprehensive backup architecture:**

```
[Production Systems]
        │
        ├─→ [Block: EBS Snapshots] ──→ [Object: S3 Archive]
        │         (hot backup)              (cold archive)
        │
        ├─→ [File: EFS/NFS] ──────────→ [Object: S3]
        │    (shared backups)              (offsite copy)
        │
        └─→ [Object: S3]
             (direct backup)
```

| Storage Type | Use For | Retention | Example |
|--------------|---------|-----------|---------|
| **Block (EBS Snapshots)** | Database point-in-time recovery | 7-30 days | PostgreSQL snapshots every 6 hours |
| **File (EFS)** | Shared backup staging, incremental backups | 1-7 days | rsync destination for multiple servers |
| **Object (S3 Standard)** | Recent backups needing quick restore | 30-90 days | Application data, configs |
| **Object (S3 Glacier)** | Long-term archive | 1-7 years | Compliance, legal holds |
| **Object (Deep Archive)** | Rarely accessed archives | 7+ years | Historical records |

**3-2-1 Backup Rule:**
- **3** copies of data
- **2** different storage types
- **1** offsite (different region)

</details>

**Q7:** A company is migrating from on-premises to cloud. They have:
- 50TB of shared files (documents, spreadsheets)
- SQL Server database (2TB, heavily used)
- Media archive (500TB, rarely accessed)

Recommend a storage solution for each and explain the migration approach.

<details>
<summary>View Answer</summary>

| Asset | Solution | Migration Approach |
|-------|----------|-------------------|
| **50TB shared files** | AWS FSx for Windows / EFS | AWS DataSync over Direct Connect |
| **SQL Server 2TB** | RDS SQL Server on EBS io2 | AWS DMS with minimal downtime |
| **500TB media archive** | S3 Glacier Deep Archive | AWS Snowball devices |

**Detailed migration:**

**1. Shared Files (50TB) → FSx/EFS**
```
Week 1-2: Set up Direct Connect (faster than internet)
Week 3-4: Initial sync with AWS DataSync
Week 5: Final sync + cutover
Tool: AWS DataSync (handles incremental, ~125 MB/s)
Downtime: Minimal (final sync during maintenance window)
```

**2. SQL Server (2TB) → RDS**
```
Option A: AWS DMS (Database Migration Service)
- Continuous replication until cutover
- Minimal downtime (<1 hour)

Option B: Native backup/restore to S3
- Full backup → S3 → Restore to RDS
- Longer downtime (hours)

Recommendation: DMS for minimal downtime
```

**3. Media Archive (500TB) → Glacier Deep Archive**
```
Problem: 500TB over internet = months
Solution: AWS Snowball Edge devices

Process:
1. Order 5-6 Snowball Edge devices (100TB each)
2. Copy data to devices on-premises
3. Ship to AWS
4. AWS imports to S3 Glacier Deep Archive

Timeline: 2-3 weeks
Cost: ~$300 per device + $0.002/GB storage
```

**Total migration timeline:** 6-8 weeks with parallel tracks

</details>

---

## Next Up

Congratulations on completing Week 2! You now understand the infrastructure fundamentals: networking, caching, and storage.

In Week 3, we'll start with **Vertical vs Horizontal Scaling** - learning how to grow your system from hundreds to millions of users!
