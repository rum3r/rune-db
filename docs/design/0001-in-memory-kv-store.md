# 0001 In-Memory KV Store

## Goal
To make a smallest in memory key-value store.

## Problem
I need a place to store key-value pairs.

## Requirements
- Simple
- Fast
- No persistence
- Easy to extend later


## Design
Use Go's in-memory map to store key-value pairs.

## Trade-offs
- No persistence
- No locking
- No replication

## Alternatives Considered
A simple file based store but hard to support updates.

## Future Improvements
- Add persistence
- Add locking
- Add replication


## References
- [Go's in-memory map](https://golang.org/pkg/container/list/)