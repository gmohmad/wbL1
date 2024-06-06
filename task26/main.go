package main

import (
	"fmt"
	"strings"
)

// Define 'allUnique' that will return true if all characters in 's' are unique and false otherwise
func allUnique(s string) bool {
	s = strings.ToLower(s)         // Convert all characters from 's' to lowercase
	set := make(map[rune]struct{}) // Initialize a set (map of empty structs)

	// Loop over 's'
	for _, v := range s {
		// We use the comma ok pattern here to check if 'v' is already in the set
		if _, ok := set[v]; ok {
			return false // if it is we return false
		}
		set[v] = struct{}{} // else we add 'v' into the set
	}

	return true // Return true if all unique
}

func main() {
	// Test 'allUnique'
	fmt.Println(allUnique("abcd"))
	fmt.Println(allUnique("abCdefAaf"))
	fmt.Println(allUnique("aabcd"))
}
