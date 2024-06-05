package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

const randSource = 1000 // Source for rand.Intn function

func main() {
	// We use flag package to let user define the number of seconds using -t flag
	secondsCount := flag.Int("t", 5, "Seconds count")
	flag.Parse() // Parse the flags

	mainStream := make(chan int) // Initialize the mainStream channel

	// Launch a goroutine that will write a random value to mainStream and sleep for half a second
	go func() {
		for {
			r := rand.Intn(randSource)
			mainStream <- r
			fmt.Printf("wrote to mainStream: %d\n", r)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	// Launch a goroutine that will read from mainStream and print it to stdout
	go func() {
		for {
			fmt.Printf("read from mainStream: %d\n", <-mainStream)
		}
	}()

	// Sleep for secondsCount seconds
	time.Sleep(time.Second * time.Duration(*secondsCount))

	// Print 'Done' message
	fmt.Println("Done")
}
