package main

import (
	"fmt"
	"sync"
)

// Define ThreadSafeMap struct with read-write mutex for thread safe operations
type ThreadSafeMap struct {
	mu sync.RWMutex
	m  map[int]int
}

// A function to return a new ThreadSafeMap
func NewThreadSafeMap() *ThreadSafeMap {
	return &ThreadSafeMap{
		mu: sync.RWMutex{},
		m:  make(map[int]int),
	}
}

// Put method-function locks the RWMutex on ThreadSafeMap, puts the key-value pair into the map, and ulocks the RWMutex
func (t *ThreadSafeMap) Put(key, value int) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.m[key] = value
}

func main() {
	var wg sync.WaitGroup   // WaitGroup for goroutine synchronization
	t := NewThreadSafeMap() // Initialize new ThreadSafeMap

	wg.Add(10) // Add 10 to WaitGroup
	for i := range 10 {
		go func() {
			defer wg.Done() // Indicate that the goroutine is done
			t.Put(i, i+17)  // Call Put on t
		}()
	}
	wg.Wait()        // Wait for all goroutines to finish
	fmt.Println(t.m) // Print the map from t
}
