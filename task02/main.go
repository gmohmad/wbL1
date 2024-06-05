package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup // Declare a WaitGroup for goroutine synchronization

	nums := [5]int{2, 4, 6, 8, 10} // Initialize nums array
	// Add length of the nums array to WaitGroup, we'll launch as many goroutines as there are values in our array
	wg.Add(len(nums))

	// In the latest versions of Go a new variable gets created on each iteration of the loop, so we dont have to worry about closures
	for _, num := range nums {
		go func() {
			defer wg.Done()        // We have to call Done on the WaitGroup to indicate that a goroutine is done, if we dont, we're gonna get a deadlock
			fmt.Println(num * num) // Println the square of the num
		}() // Call anonoumus fucntion in a goroutine
	}
	wg.Wait() // The 'Join point' of our program. Here we wait till all the goroutines call Done on the WaitGroup

	fmt.Println("done") // Print 'done' message
}
