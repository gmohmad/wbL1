package main

import (
	"fmt"
)

// Define 'reverseString' that will return the inverse of the given string
func reverseString(s string) string {
	reversed := []rune(s) // Initialize reversed, which is a slice of runes from string 's'

	l := len(reversed) // length of reversed, (we cant take the length of s, because it will get the amount of bytes in it)

	left, right := 0, l-1 // Initialize left and right, the two pointers

	// Loop while left is less than right, swapping the values from the left of 'reversed' with their opposites from the right
	for left < right {
		reversed[left], reversed[right] = reversed[right], reversed[left]
		// Increment left and decrement right
		left++
		right--
	}

	return string(reversed) // Make 'reversed' a string and return it
}

func main() {
	fmt.Println(reverseString("«главрыба")) // Print the result of calling 'reverseString' on '«главрыба'
}
