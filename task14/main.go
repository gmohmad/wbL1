package main

import (
	"fmt"
	"reflect"
)

func main() {
	var I interface{} // Declare I variable which is an empty interface

	I = 23 // Assign I to int type

	fmt.Println(reflect.TypeOf(I)) // Print the type of I using reflect.TypeOf

	I = "asdf" // Assign I to string type

	fmt.Printf("%T\n", I) // Print the type of I using %T, which gets the type of the given value
}
