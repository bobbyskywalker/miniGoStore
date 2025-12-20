# 🦫 What is miniGoStore?

**miniGoStore** is a tiny in-memory caching service written in Go.

It provides a lightweight key–value store with a custom, simple TCP-based protocol.

## What's inside?
**miniGoStore** offers a simple protocol for storing, fetching, and querying TTLs of currently stored data.

### 🔗 Commands:
- The protocol is line-based (`\n` terminated).
- Commands and options are case-insensitive.
- Each command returns exactly one response.
- One TCP connection represents one client session.

**`PASS`** <password>

  Authenticate the client.

**`GET`** <key>

  Return the value stored at <key>, or (nil) if the key does not exist.

**`SET`** <key> <value> [NX|XX] [EX|PX <seconds|milliseconds>] [GET]

  Store <value> at <key>
  * NX  -   Set only if the key does not exist.
  * XX  -   Set only if the key already exists.
  * EX  -   Set expiration time in seconds.
  * PX  -   Set expiration time in milliseconds.
  * GET -   Return the previous value stored at the key.

**`GETEX`** <key> [EX|PX <seconds|milliseconds>] [PERSIST]

  Return the value of <key> and optionally update its TTL.
  * EX/PX   Set a new expiration.
  * PERSIST Remove any existing expiration.

**`TTL`** <key>
  Return the remaining time-to-live of <key>.
  -1 if the key exists but has no expiration.
  -2 if the key does not exist.

**`PING`**

  Check server liveness. Get back PONG.

**`DEL`** <key>

  Delete <key> and its associated value.

**`QUIT`**

  Close the client connection.

### 🧹 Cleanup routine
A background cleanup routine continuously scans and removes TTL-expired keys and values.

### 🧪 Unit tests
The ***store*** package contains unit tests for command behavior.

### ⚡ Benchmark

Alongside the server executable, a benchmark script is provided:
***./cmd/benchmark/main.go***

It can be used to stress-test concurrency, throughput, and locking behavior.

### 🔐 Authentication
The server requires authentication using the `PASS` command.

The provided password is:
- hashed at startup
- never stored in plaintext
- explicitly erased from memory after hashing (best-effort)

### 🔧 Build & Run miniGoStore

```bash
go build ./cmd/server/main.go
./main <password> [port]
```
The port is set to 8080 by default if none is provided.

## Differences from Redis

miniGoStore is a learning project inspired by Redis but intentionally simpler:
- No persistence (RDB/AOF)
- No pub/sub
- No transactions
- Minimal command set

Think twice before trying to use it in your app :)

## Author

Built by https://github.com/bobbyskywalker  🦫

## Contributions
Fork & post PRs freely if you are interested in this project :)

## License

MIT License. See `LICENSE` tab.
