package main

import "fmt"

// Define Human struct with fields Name and Age
type Human struct {
	Name string
	Age  uint8
}

// Define sayHi method with *Human receiver
func (h *Human) sayHi() {
	fmt.Println("Hi")
}

// Define Action struct and embed Human struct into it
type Action struct {
	Human
}

// Define sayHi method with *Action receiver
func (a *Action) learnGo() {
	fmt.Println("*learning go*")
}

func main() {
	// Create a new Action object and initialize the embedded Human struct (notice we can't just define Name and Age as if they were Action fields)
	action := Action{
		Human: Human{
			Name: "Rob",
			Age:  40,
		},
	}

	// But we can access the embedded Human struct's methods and fields as if they were defined on Action
	// Call sayHi on Action
	action.sayHi()

	// Call learnGo on Action
	action.learnGo()
}
