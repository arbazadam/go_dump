package main

import (
	"fmt"
)

/*
Perfect! Let’s design the **Simple Chat Server** in detail. I’ll outline **requirements, features, and behavior** so you can turn this into working Go code step by step.

---

# 📝 Project: Simple Chat Server (Go)

## 1️⃣ Purpose

Create a TCP-based chat server where multiple clients can connect concurrently, send messages, and see messages from others. Messages are **buffered in memory** before being written to a log file. This will cover:

* Sockets/networking
* Concurrency (goroutines, channels)
* Buffers for message handling
* File I/O
* Basic logging

---

## 2️⃣ Core Components

### a) Server

* Listens on a TCP port (e.g., `localhost:9000`)
* Accepts multiple client connections simultaneously
* Maintains a list of **connected clients**
* Receives messages from clients and broadcasts to all others

---

### b) Client Handler

* Each connected client runs in its **own goroutine**
* Reads messages from the client
* Forwards messages to the **message broadcasting system**

---

### c) Message Buffer

* Use a `bytes.Buffer` (or a custom struct) to **store incoming messages temporarily**
* Periodically flush buffered messages to a **log file** (e.g., every N seconds or after N messages)
* Optionally, implement a **mutex** to make buffer access **thread-safe**

---

### d) Broadcasting System

* Receives messages from clients via a **channel**
* Sends messages to **all connected clients**
* Also writes messages to the **buffer** for later logging

---

### e) Logging

* Buffered messages are written to a file using `fmt.Fprintf`
* Log format example:

```
[2025-12-01 14:35:22] Alice: Hello, world!
```

* Ensure **concurrent writes are safe** (use mutexes or channels)

---

## 3️⃣ Features / Requirements

1. **Multiple clients** can connect without blocking the server
2. Messages are **broadcast** to all connected clients
3. Messages are **buffered in memory** and written to a log file periodically
4. Message format includes **timestamp + client name + message**
5. Graceful handling of **client disconnects**
6. Thread-safe buffer access
7. Server should **continue running** even if one client disconnects or sends bad data
8. Optional: implement a **simple command** like `/quit` to disconnect

---

## 4️⃣ Suggested Data Structures

* `type Client struct`

  * Holds connection (`net.Conn`), name/ID, outgoing channel
* `var clients map[string]Client` or `[]Client`
* `messagesChannel chan string` for broadcasting
* `var buf bytes.Buffer` for message storage
* `sync.Mutex` for buffer access

---

## 5️⃣ Flow of Execution

1. Server starts and listens on TCP port
2. Server accepts a new connection → spawns **client goroutine**
3. Client sends a message → goroutine reads it → sends it to **broadcast channel**
4. Broadcast goroutine:

   * Sends message to all connected clients
   * Appends message to **buffer**
5. Periodically (or when buffer reaches threshold):

   * Flush buffer to **log file**
6. Server handles client disconnects gracefully

---

## 6️⃣ Optional Enhancements (for practice)

* Private messages (e.g., `/msg Bob Hello`)
* Message history on new client connection
* Limit buffer size to prevent memory overflow
* Timestamps in **different timezones** or **formatted nicely**
* Support for **UTF-8 emojis and special characters**

---

If you want, I can also create a **step-by-step implementation roadmap** breaking the above into **coding milestones**, so you can build it incrementally and test each part.

Do you want me to do that next?

*/

func main() {
	intChan := make(chan int)
	stringChan := make(chan string)

	go func() {
		for {
			select {
			case p := <-intChan:
				fmt.Println(p)
			case g := <-stringChan:
				fmt.Println(g)
			}
		}
	}()

	intChan <- 20
	stringChan <- "arbaz"
}
