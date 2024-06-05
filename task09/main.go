package main

import "fmt"

func main() {
	arr := [5]int{2, 4, 6, 8, 10} // Define the array

	// ints is a stage of a conveyor that returns intStream with values from the array
	ints := func(arr [5]int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream) // We have to close the channel, if we don't - deadlock
			for _, v := range arr {
				intStream <- v
			}
		}()
		return intStream
	}

	// squares is a stage of a conveyor that returns squaresStreamw with values from the given channel
	squares := func(intStream <-chan int) <-chan int {
		squaresStream := make(chan int) // Same as in ints function, we have to close to prevent a deadlock
		go func() {
			defer close(squaresStream)
			for v := range intStream {
				squaresStream <- v * v
			}
		}()
		return squaresStream
	}

	// Here we pass ints() to squares(), which will return intStream, then iterate over the squaresStream returned by squares()
	for v := range squares(ints(arr)) {
		fmt.Println(v)
	}
}
