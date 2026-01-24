# concurrency correctness, patterns, and trade-offs.

## The Mechanics of Concurrency

**Goal:** Understand how the runtime handles goroutines and data flow (Channels).

#### **1. Goroutine Lifecycle & Scheduling**

* **The M:N Scheduler:** Go multiplexes **M** goroutines onto **N** OS threads.
* **Key Insight:** Goroutines are cheap (~2KB stack) compared to threads (~1MB). Context switching is user-space and fast.
* **The "Main" Trap:** If `main()` returns, the program exits immediately. All other goroutines are killed instantly without cleanup. You *must* block `main` (using `sync.WaitGroup` or a channel) to wait for work to finish.



#### **2. Channels: The Conduits**

* **Unbuffered (`make(chan T)`):**
* **Behavior:** A sender blocks until a receiver is ready. A receiver blocks until a sender is ready.
* **Use Case:** strict synchronization; "handing off" data guarantees the other side received it.


* **Buffered (`make(chan T, size)`):**
* **Behavior:** A sender only blocks if the buffer is **full**. A receiver blocks if the buffer is **empty**.
* **Use Case:** Decoupling producers from consumers; handling bursty traffic.


* **The "Closed" State:**
* **Send:** Panic.
* **Receive:** Returns zero-value immediately. Use `val, ok := <-ch` to detect closure (`!ok` means closed).
* **Close:** Only the sender should close. Closing a closed channel panics.



#### **3. `select` Statement**

The "switch" for channels.

* **Randomness:** If multiple cases are ready, `select` picks one **randomly**. This prevents starvation of one channel.
* **Default:** Adding a `default` case makes the `select` **non-blocking**.
* *Tip:* Use `select` with `default` to implement "try-send" or "try-receive" logic.



---

## Synchronization & Safety

**Goal:** Prevent data races and choose the right tool (Mutex vs. Channel).

#### **1. Race Conditions**

A race occurs when two goroutines access the same variable concurrently, and at least one access is a write.

* **Detection:** Always run tests with `go test -race`.
* **Memory Model:** Without explicit synchronization, CPU caches may not sync. Goroutine A might see "stale" data from Goroutine B, or worse, a partially written struct.

#### **2. Mutex (`sync.Mutex`) vs. Channels**

* **Use Mutex When:** You are protecting **state** (e.g., a map, a counter, a struct field).
* *Why:* It is faster (CPU-level spinlock/futex) and semantically clearer for "locking data."


* **Use Channels When:** You are **passing ownership** of data or coordinating workflow.
* *Why:* "Share memory by communicating." It decouples the sender from the receiver.



#### **3. `sync.WaitGroup**`

* **Purpose:** Waiting for a set of concurrent operations to complete.
* **Pitfall:** **Never copy a WaitGroup.** Always pass it by pointer (`*sync.WaitGroup`) to functions. Copying it copies the internal counter state, leading to deadlocks or panics.

---

## Patterns (Fan-out, Worker Pools)

**Goal:** Implement standard system design patterns in Go.

#### **1. Fan-Out / Fan-In**

* **Fan-Out:** Spawning multiple goroutines to handle tasks in parallel.
* **Fan-In:** merging results from multiple channels into a single channel.

#### **2. Worker Pool (The "Throttle")**

Launching 1 million goroutines is possible but bad (exhausts file descriptors, DB connections). A worker pool limits concurrency.

**Code Exercise: Worker Pool**
*Requirement:* Process jobs with a fixed number of workers.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs { // Loop runs until 'jobs' is closed
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(100 * time.Millisecond) // Simulate work
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	var wg sync.WaitGroup

	// 1. Start fixed number of workers (e.g., 3)
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// 2. Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs) // Signal "no more jobs"

	// 3. Wait and close results in a separate goroutine
	go func() {
		wg.Wait()
		close(results)
	}()

	// 4. Collect results
	for res := range results {
		fmt.Println("Result:", res)
	}
}

```


## Advanced Control & Trade-offs

**Goal:** Handle timeouts and justify decisions.

#### **1. Context & Timeouts**

Never let a goroutine run forever. Use `context.WithTimeout`.

**Code Exercise: Concurrent Fetch with Timeout**
*Requirement:* Fetch URLs, max 2 concurrently, global timeout 500ms.

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func fetch(ctx context.Context, url string, results chan<- string) {
	select {
	case <-time.After(200 * time.Millisecond): // Simulate work
		results <- fmt.Sprintf("Success: %s", url)
	case <-ctx.Done(): // Cancelled
		results <- fmt.Sprintf("Timeout: %s", url)
	}
}

func main() {
	urls := []string{"http://a.com", "http://b.com", "http://c.com"}
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	results := make(chan string)
	
	// Semaphore pattern: buffered channel of size 2
	sem := make(chan struct{}, 2) 

	for _, url := range urls {
		go func(u string) {
			// Acquire token
			sem <- struct{}{} 
			defer func() { <-sem }() // Release token

			// Check context before working
			if ctx.Err() == nil {
				fetch(ctx, u, results)
			}
		}(url)
	}

	// Read results (simplified for brevity)
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-results)
	}
}

```

### Interview Questions

**Q1: Mutex vs Channel for a Metrics Counter?**

* **Answer:** **Mutex (or `sync/atomic`).**
* **Reasoning:** A counter is simple shared state. Using a channel involves overhead (locking + context switch) to send a value to a receiver. An `atomic.AddInt64` is a single CPU instruction, and a Mutex is very fast. Channels are for flow, not just safe access.

**Q2: How do you debug a deadlock?**

* **Answer:** Use `Ctrl+\` (SIGQUIT) to dump the stack trace of all goroutines. Look for goroutines stuck in `semacquire` (Mutex wait) or `chan send/receive`. Also, use `pprof` to visualize blocking profiles.

**Q3: What happens if you add to a WaitGroup *inside* the goroutine?**

* **Answer:** It's a race condition. If the scheduler runs the "waiter" before the "worker" starts, `wg.Wait()` might return immediately because the counter is still 0. Always `wg.Add` *before* launching the goroutine.