package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s) // Split words in 's' by spaces, and assing the result slice to 'words'

	l := len(words) // length of reversed, (we cant take the length of s, because it will get the amount of bytes in it)

	left, right := 0, l-1 // Initialize left and right, the two pointers

	// Loop while left is less than right, swapping the values from the left of 'words' with their opposites from the right
	for left < right {
		words[left], words[right] = words[right], words[left]
		// Increment left and decrement right
		left++
		right--
	}

	return strings.Join(words, " ") // Join the 'words' slice by spaces and return it
}

func main() {
	s := "snow dog sun" // Initialize s string

	fmt.Println(reverseWords(s)) // Print the result of calling reverseWords on 's'
}
