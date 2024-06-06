package main

import (
	"errors"
	"fmt"
)

// Define a function 'deleteFromSlice' that deletes a value from 'sl' at specified 'idx', or returns an error
func deleteFromSlice(sl []int, idx int) ([]int, error) {
	// If index is less than 0 or greater than/equal to the length of sl, we return an error
	if idx < 0 || idx >= len(sl) {
		return nil, errors.New("Index out of range")
	}

	// Else, using append, we kinda 'cover' the element at the 'idx' by elements from sl[idx+1:]
	sl = append(sl[:idx], sl[idx+1:]...)

	return sl, nil // And return the new slice
}

func main() {
	sl := []int{1, 2, 3, 4, 5} // Initialize our slice

	sl, err := deleteFromSlice(sl, 3) // Call 'deleteFromSlice' on sl with index 3

	// If error occurred we just print it and stop the program
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(sl) // Else we print the slice
}
