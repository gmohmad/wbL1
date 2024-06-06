package main

import "fmt"

// Define 'quickSort' function that recursively sorts the given slice in O(log(n)) time complexity and O(1) memory usage
func quickSort(nums []int) {
	// This is the base case, if lenght of 'nums' is less than 2, it's already sorted
	if len(nums) < 2 {
		return
	}

	// We choose the pivot as the last element of the slice
	pivot := nums[len(nums)-1]
	// 'i' will be the index where the pivot will end up, and 'j' will be index of the current value
	i, j := -1, 0

	// Here we loop until 'j' is equal to length of nums slice - 1.
	for j < len(nums)-1 {
		// On each iteration, if nums[j] is less than pivot, we increment 'i' and swap nums[i] with nums[j]
		if nums[j] < pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
		// And increment 'j' anyways
		j++
	}
	// After j reaches length nums - 1, we increment 'i' and swap nums[i] and nums[j] again, which will put pivot at index 'i'
	i++
	nums[i], nums[j] = nums[j], nums[i]

	// And here we recursively call quickSort on left and right sides of the pivot
	quickSort(nums[:i])
	quickSort(nums[i+1:])
}

func main() {
	nums := []int{1, 2, 10, -1, 3, 6, 4, 6, 5, 9} // Initialize nums slice

	quickSort(nums) // Call 'quickSort' on nums

	fmt.Println(nums) // Print sorted 'nums' slice
}
