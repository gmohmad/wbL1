package main

import (
	"fmt"
	"time"
)

// Define function 'Sleep' that will read from time.Tick channel, which will be blocked for 'd' seconds
func Sleep(d int) {
	<-time.Tick(time.Second * time.Duration(d))
}

func main() {
	fmt.Println("falling asleep...") // Print a message before sleeping
	Sleep(3)                         // Call 'Sleep' function for 3 seconds
	fmt.Println("woke up!")          // Print a message after sleeping
}
