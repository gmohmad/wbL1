package main

import (
	"fmt"
	"sync"
)

// Define ThreadSafeCounter struct with a mutex for thread safe write operations
type ThreadSafeCounter struct {
	mu      sync.Mutex
	counter int
}

// Inc funcion that locks the mutex on t, add delta to t.counter, and in defer unlocks the mutex
func (t *ThreadSafeCounter) Inc(delta int) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.counter += delta
}

// A funcion that returns a new ThreadSafeCounter
func NewThreadSafeCounter() *ThreadSafeCounter {
	return &ThreadSafeCounter{
		mu:      sync.Mutex{},
		counter: 0,
	}
}

const numWorkers = 1000 // Number of goroutines to be launched

func main() {
	var wg sync.WaitGroup       // A WaitGroup for goroutine synchronization
	t := NewThreadSafeCounter() // Initialize a new ThreadSafeCounter

	wg.Add(numWorkers) // Add numWorkers to WaitGroup

	// Here we launch numWorkers goroutines that will call wg.Done() and increment the t.counter by 1
	for range numWorkers {
		go func() {
			defer wg.Done()
			t.Inc(1)
		}()
	}

	wg.Wait()                                          // Wait for the goroutines
	fmt.Printf("The value of couter: %d\n", t.counter) // Print the value of t.counter
}
