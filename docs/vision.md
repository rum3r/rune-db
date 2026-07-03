# RuneDB Vision

## Why RuneDB?

As software engineers, we rely on databases every day, yet most of us treat them as black boxes. We know how to use them, but rarely understand how they actually work under the hood.

RuneDB is my attempt to change that.

This project is not about building the next PostgreSQL or RocksDB. It is about learning the engineering principles behind modern databases by implementing them from first principles in Go.

The goal is to understand not only *what* databases do, but *why* they are designed the way they are.

---

## What do I want to learn?

Through RuneDB, I want to gain a deep understanding of modern database internals, including:

- Storage engines and how data is organized on disk.
- Write-Ahead Logging (WAL) and crash recovery.
- LSM Trees and SSTables.
- Transactions and MVCC.
- Indexing and query execution.
- Concurrency control.
- Replication and distributed systems.
- The operating system concepts that make databases efficient, such as filesystems, virtual memory, page caching, and disk I/O.

I don't want to simply read about these concepts—I want to understand them by building each component myself.

---

## How will I achieve this?

RuneDB will be built incrementally.

Each milestone introduces one new concept found in production databases while keeping the implementation as simple as possible.

The implementation will be guided by:

- *Designing Data-Intensive Applications* (DDIA)
- *Operating Systems: Three Easy Pieces* (OSTEP)
- Research papers
- Database source code and engineering blogs

The objective is not to copy an existing database, but to understand the trade-offs behind every design decision and implement my own solutions.

---

## What RuneDB will **NOT** do

RuneDB is **not** intended to be production-ready.

It is **not** designed to compete with PostgreSQL, RocksDB, SQLite, or any other production database.

The project will prioritize:

- Simplicity over optimization.
- Correctness over performance.
- Learning over feature completeness.

Whenever a design decision becomes too complex, I will first implement the simplest correct solution and improve it incrementally as my understanding grows.

---

## Philosophy

> Build. Break. Learn. Repeat.

Every feature should answer one engineering question.

Every implementation should teach one new concept.

Every design decision should be intentional.