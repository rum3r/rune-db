# RuneDB Roadmap

RuneDB is a database built from first principles in Go.

The goal of this project is not to build the fastest or most feature-rich database, but to understand how modern storage engines and distributed databases work by implementing them from scratch.

Each milestone introduces a new concept commonly found in production databases while keeping the implementation as simple as possible.

---

## Phase 1 — Core Storage

- [ ] 0001 - In-Memory Key-Value Store
- [ ] 0002 - Persistent Storage
- [ ] 0003 - Write-Ahead Log (WAL)
- [ ] 0004 - Recovery from WAL

**Goal**

Build a simple key-value database that survives process restarts.

---

## Phase 2 — Storage Engine

- [ ] 0005 - Memtable
- [ ] 0006 - SSTables
- [ ] 0007 - Compaction
- [ ] 0008 - Bloom Filters

**Goal**

Build an LSM Tree based storage engine capable of efficient reads and writes.

---

## Phase 3 — Query Engine

- [ ] 0009 - Secondary Indexes
- [ ] 0010 - Range Scans
- [ ] 0011 - Simple Query Language

**Goal**

Provide richer ways to retrieve and query stored data.

---

## Phase 4 — Transactions

- [ ] 0012 - Transactions
- [ ] 0013 - MVCC
- [ ] 0014 - Snapshot Isolation

**Goal**

Support atomic and isolated operations while maintaining correctness under concurrency.

---

## Phase 5 — Concurrency

- [ ] 0015 - Reader/Writer Locks
- [ ] 0016 - Concurrent Reads and Writes
- [ ] 0017 - Background Compaction

**Goal**

Allow multiple clients to interact with RuneDB safely and efficiently.

---

## Phase 6 — Networking

- [ ] 0018 - TCP Server
- [ ] 0019 - Client Library
- [ ] 0020 - Wire Protocol

**Goal**

Allow external applications to communicate with RuneDB over the network.

---

## Phase 7 — Replication

- [ ] 0021 - Leader/Follower Replication
- [ ] 0022 - Log Replication
- [ ] 0023 - Failure Recovery

**Goal**

Replicate data across multiple nodes for higher availability.

---

## Phase 8 — Distributed Database

- [ ] 0024 - Partitioning (Sharding)
- [ ] 0025 - Consistent Hashing
- [ ] 0026 - Cluster Membership
- [ ] 0027 - Raft Consensus

**Goal**

Transform RuneDB from a single-node database into a distributed database.

---

## Future Ideas

- SQL Parser
- Query Optimizer
- B+ Tree Storage Engine
- Columnar Storage
- Compression
- Encryption at Rest
- Time-to-Live (TTL)
- Metrics & Observability
- Backup & Restore
- Web UI

---

## Reading Plan

### Designing Data-Intensive Applications

- Reliable, Scalable and Maintainable Applications
- Storage and Retrieval
- Encoding and Evolution
- Replication
- Partitioning
- Transactions
- Distributed Systems

### Operating Systems: Three Easy Pieces

- Processes
- Threads
- Concurrency
- Locks
- Condition Variables
- Virtual Memory
- File Systems
- Crash Consistency

---

## Principles

- Keep every implementation as simple as possible.
- Understand the "why" before writing code.
- Prefer correctness over optimization.
- Benchmark before optimizing.
- Write design documents before major features.
- Learn by building, not by copying.