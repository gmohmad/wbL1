package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const randSource = 1000 // Source for rand.Intn function

// Here we define a function that will infinitely print the value from mainStream channel and sleep for a second
func worker(mainStream <-chan int) {
	for {
		fmt.Println(<-mainStream)
		time.Sleep(time.Second)
	}
}

func main() {
	// We use flag package to let user define the number of workers using -w flag
	workersCount := flag.Int("w", 5, "Number of workers")
	flag.Parse() // Parse the flags

	mainStream := make(chan int) // Initialize the mainStream channel

	// Launch workersCount workers
	for range *workersCount {
		go worker(mainStream)
	}

	// infinitely write random numbers in range [0, randSource) to mainStream channel
	go func() {
		for {
			mainStream <- rand.Intn(randSource)
		}
	}()

	// Here we initialize signals channel that will listen on syscall.SIGINT, and print 'Program stopped' after user presses <C-c>
	signals := make(chan os.Signal)
	signal.Notify(signals, syscall.SIGINT)
	<-signals
	fmt.Println("Program stopped")
}
