package main

import (
	"flag"
	"fmt"
)

func main() {
	// Use flags to get the values
	num := flag.Int64("n", 100, "The number to perform bit swap on")
	position := flag.Int64("p", 5, "The position of the bit to be swapped from the right")
	value := flag.Uint("v", 1, "The value that will be placed instead of the bit on the given position. 0 or 1")
	flag.Parse() // Parse the flags

	mask := int64(1) << int64(*position) // We craete a mask, where 1 will be shifted by position zeroes to the right

	if *value == 1 {
		fmt.Println(*num | mask) // If value is 1, we simply OR the num and the mask
	} else if *value == 0 {
		fmt.Println(*num &^ mask) // If value is 0, we have to do AND the num and inverse of mask
	} else {
		fmt.Println("Value must be 0 or 1") // If value not 1 or 0, we print an error message
	}
}
