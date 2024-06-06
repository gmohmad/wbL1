package main

import "fmt"

func main() {
	words := [5]string{"cat", "cat", "dog", "cat", "tree"} // Initialize words array

	set := make(map[string]struct{}) // Initialize the set

	// Loop over the words array and add values from it to set on each iteration
	for _, v := range words {
		set[v] = struct{}{}
	}

	fmt.Println(set) // Print the set
}
