package main

import "fmt"

// Since there are no sets in Go at the moment, in this solution by 'set' i mean a map that stores
// empty structs as values (because they take 0 bytes).

// Define setIntersection function
func setIntersection(set1, set2 map[int]struct{}) map[int]struct{} {
	intersectionSet := make(map[int]struct{}) // New set to store keys that exist in both set1 and set2
	for k := range set1 {
		// If k exists in set2, we add it to intersectedSet
		if v, ok := set2[k]; ok {
			intersectionSet[k] = v
		}
	}
	return intersectionSet // return the intersectedSet
}

func main() {
	// Initialize two sets
	set1 := map[int]struct{}{1: {}, 3: {}, 7: {}, 4: {}}
	set2 := map[int]struct{}{2: {}, 3: {}, 7: {}, 11: {}}

	// Print the result of calling setIntersection with set1 and set2 as arguments
	fmt.Println(setIntersection(set1, set2))
}
