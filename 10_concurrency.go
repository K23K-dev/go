//go:build ignore

// Lesson 10: Concurrency.
//
// Run it with:
//
//	go run 10_concurrency.go
package main

import (
	"fmt"
	"slices"
	"sync"
	"time"
)

func main() {
	goroutinesAndWaitGroups()
	channels()
	bufferedChannels()
	selectWithTimeout()
	workerPool()
	sharedCounter()
}

// A goroutine is a function running concurrently with the rest of the
// program; you start one by writing "go" before a function call. Goroutines
// are much cheaper than OS threads (each starts with a few KB of stack), so
// programs commonly run thousands of them. Go's runtime spreads them across
// a small number of OS threads.
func goroutinesAndWaitGroups() {
	fmt.Println("--- Goroutines and WaitGroups (the order of lines varies) ---")

	// A WaitGroup waits for a group of goroutines to finish. wg.Go starts a
	// function in a new goroutine and keeps count; wg.Wait blocks until all
	// of them have returned.
	var wg sync.WaitGroup
	for i := range 3 {
		wg.Go(func() {
			fmt.Println("goroutine", i, "running")
		})
	}
	wg.Wait()
	fmt.Println("all goroutines finished")

	// Code written before Go 1.25 does the same thing like this, and you'll
	// see it a lot:
	//
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		...
	//	}()
}

// A channel is a typed pipe that goroutines use to pass values to each
// other. ch <- v sends and <-ch receives. On an unbuffered channel, a send
// waits until another goroutine receives, and vice versa, so the channel
// also synchronizes them.
func channels() {
	fmt.Println("--- Channels ---")
	results := make(chan string)
	go func() {
		time.Sleep(10 * time.Millisecond) // pretend to do some work
		results <- "work done"
	}()
	fmt.Println(<-results) // waits here until the goroutine sends

	// The sender can close a channel to say there are no more values.
	// A for-range loop over a channel receives until it's closed.
	squares := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			squares <- i * i
		}
		close(squares)
	}()
	for sq := range squares {
		fmt.Print(sq, " ")
	}
	fmt.Println()
}

// A buffered channel holds up to a fixed number of values. Sends only wait
// when the buffer is full, and receives only wait when it's empty.
func bufferedChannels() {
	fmt.Println("--- Buffered channels ---")
	queue := make(chan string, 2)
	queue <- "first" // doesn't wait: there's room in the buffer
	queue <- "second"
	// A third send would wait forever, since nothing else is running to
	// receive. Go would detect that and crash with
	// "fatal error: all goroutines are asleep - deadlock!".
	fmt.Println(<-queue, <-queue)
}

// select waits on several channel operations at once and runs the first one
// that's ready. Combined with time.After, it gives you a timeout.
func selectWithTimeout() {
	fmt.Println("--- select and timeouts ---")
	fast := slowOperation(10 * time.Millisecond)
	slow := slowOperation(500 * time.Millisecond)

	for _, ch := range []<-chan string{fast, slow} {
		select {
		case result := <-ch:
			fmt.Println("got:", result)
		case <-time.After(100 * time.Millisecond):
			fmt.Println("gave up after 100ms")
		}
	}
}

// slowOperation returns a channel that receives a message after d.
// The channel has room for one value, so the goroutine can always send its
// result and exit, even if nobody is waiting anymore. With an unbuffered
// channel, the goroutine behind a timed-out operation would wait forever:
// a goroutine leak.
func slowOperation(d time.Duration) <-chan string {
	ch := make(chan string, 1)
	go func() {
		time.Sleep(d)
		ch <- fmt.Sprint("finished after ", d)
	}()
	return ch
}

// A worker pool: a few goroutines take jobs from one channel and send
// results to another. Here three workers square the numbers 1 to 9.
func workerPool() {
	fmt.Println("--- Worker pool ---")
	jobs := make(chan int)
	results := make(chan int)

	var wg sync.WaitGroup
	for range 3 {
		wg.Go(func() {
			for n := range jobs { // each worker takes jobs until jobs is closed
				results <- n * n
			}
		})
	}

	// Send the jobs from another goroutine, so this one is free to collect
	// results at the same time.
	go func() {
		for n := 1; n <= 9; n++ {
			jobs <- n
		}
		close(jobs)
	}()

	// Close results once every worker has finished. That ends the range
	// loop below.
	go func() {
		wg.Wait()
		close(results)
	}()

	var squares []int
	for sq := range results {
		squares = append(squares, sq)
	}
	// Results arrive in whatever order the workers finish, so sort them.
	slices.Sort(squares)
	fmt.Println(squares)
}

// When goroutines share a variable, access to it must be synchronized.
// Otherwise you have a data race, and updates get lost.
func sharedCounter() {
	fmt.Println("--- Shared state and mutexes ---")
	const workers, increments = 100, 10000

	// Without a lock: count++ is really "read, add 1, write back". Two
	// goroutines can read the same value, both add 1, and one update is lost.
	unsafeCount := 0
	var wg sync.WaitGroup
	for range workers {
		wg.Go(func() {
			for range increments {
				unsafeCount++
			}
		})
	}
	wg.Wait()
	fmt.Println("without a lock:", unsafeCount, "(should be 1000000; usually less, and it varies)")

	// A sync.Mutex lets only one goroutine at a time run the code between
	// Lock and Unlock.
	var mu sync.Mutex
	safeCount := 0
	for range workers {
		wg.Go(func() {
			for range increments {
				mu.Lock()
				safeCount++
				mu.Unlock()
			}
		})
	}
	wg.Wait()
	fmt.Println("with a mutex:  ", safeCount)

	// For a plain counter like this, the sync/atomic package (atomic.Int64)
	// is another option. More generally, Go's rule of thumb is "don't
	// communicate by sharing memory; share memory by communicating": pass
	// data between goroutines through channels where you can, and use a
	// mutex for simple shared state.
	//
	// Go has a race detector (go run -race) that catches bugs like the first
	// counter, but on Windows it needs a C compiler installed.
}
