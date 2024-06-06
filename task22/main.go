package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Initialize two big.Int's - num1 and num2, with values > 2^20
	num1 := big.NewInt(2 << 24)
	num2 := big.NewInt(2 << 23)

	// This is the big.Int we're going to do */-+ operations on
	res := new(big.Int)

	// Set the product of num1 and num2 as res's value and print it
	res.Mul(num1, num2)
	fmt.Printf("Product: %s\n", res.String())

	// Set the division of num1 and num2 as res's value and print it
	res.Div(num1, num2)
	fmt.Printf("Division: %s\n", res.String())

	// Set the substraction of num1 and num2 as res's value and print it
	res.Sub(num1, num2)
	fmt.Printf("Substraction: %s\n", res.String())

	// Set the sum of num1 and num2 as res's value and print it
	res.Add(num1, num2)
	fmt.Printf("Addition: %s\n", res.String())
}
