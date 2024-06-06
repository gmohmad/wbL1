package main

import "fmt"

func main() {
	a, b := 10, 100                      // Initialize a and b with values 10 and 100
	fmt.Printf("a - %d, b - %d\n", a, b) // Print the values of a and b

	a, b = b, a                          // Swap the values of a and b
	fmt.Printf("a - %d, b - %d\n", a, b) // Print the values of a and b
}
