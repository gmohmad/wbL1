package main

import "fmt"

// Define the groupTemperatures function
func groupTemperatures(temps []float64) map[int][]float64 {
	m := make(map[int][]float64) // Initialize the map

	for _, v := range temps {
		// We round down the numbers from temps slice to nearest number that is divided by 10 with no remainder
		key := int(v/10) * 10
		m[key] = append(m[key], v) // Append the value to the key we calculated
	}

	return m // return the map
}

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5} // Initialize the temps slice

	fmt.Println(groupTemperatures(temps)) // Print the result of calling groupTemperatures on temps
}
