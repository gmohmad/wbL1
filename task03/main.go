package main

import (
	"fmt"
	"sync"
)

var (
	squareSum int            // Declare a variable we'll store the result in
	m         sync.Mutex     // We'll use sync.Mutex to prevent data race when adding to squareSum
	wg        sync.WaitGroup // And a WaitGroup for goroutine synchronization
)

func main() {
	nums := [5]int{2, 4, 6, 8, 10} // Initialize nums array

	wg.Add(len(nums)) // Add length of nums to WaitGroup

	// Iterate over nums array
	for _, num := range nums {
		// Launch a goroutine that will lock the mutex, add num*num to squareSum, unlock the mutex and call wg.Done()
		go func() {
			defer wg.Done()

			m.Lock()
			defer m.Unlock()

			squareSum += num * num
		}()
	}
	// Wait for the goroutines to finish
	wg.Wait()

	fmt.Println(squareSum) // Print the squareSum
}
