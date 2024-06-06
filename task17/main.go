package main

import (
	"fmt"
)

func binarySearch(nums []int, target int) (int, int) {
	l, r := 0, len(nums)-1 // Initialize l and r with 0 and lenght of nums - 1

	tries := 0 // Initialize the tries counter

	// Here we loop infinitely until l becomes greater than r
	for l <= r {
		// On each iteration we increment the tries counter, and get a new guess, which is l + r divied by 2
		tries++
		guess := (l + r) / 2

		switch {
		// If nums[guess] equals to target, we return the guess and tries it took to find it
		case nums[guess] == target:
			return guess, tries
		// if nums[guess] is less than targe, we make l equal to guess + 1, effectively cutting off the half of the
		// nums slice because we know that the target is not there
		case nums[guess] < target:
			l = guess + 1
		// if nums[guess] is greater than targe, we make r equal to guess - 1, effectively cutting off the half of the
		// nums slice because we know that the target is not there
		case nums[guess] > target:
			r = guess - 1
		}
	}

	return -1, tries // If the target was not found, we return -1 and tries
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // Initialize the nums slice
	res, tries := binarySearch(nums, 11)         // Assign res and tries to the result of calling binarySearch function

	fmt.Printf("Result: %d, tries: %d\n", res, tries) // Print the result and tries
}
