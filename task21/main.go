package main

import "fmt"

// Target is the interface clients expect to work with
type Target interface {
	Method()
}

// Current is the interface that is currently used in our system
type Current interface {
	AdapteeMethod()
}

// Adaptee is the existing struct that is going to be adapted
type Adaptee struct{}

// The Adaptee method that needs to be adapted
func (a *Adaptee) AdapteeMethod() {
	fmt.Println("adaptee method")
}

// The Adapter struct that adapts Adaptee to Target interface
type Adapter struct {
	adaptee Current
}

// NewAdapter takes a type that implements Current interface and returns a pointer to a new adapter
func NewAdapter(c Current) *Adapter {
	return &Adapter{c}
}

// Adapt 'AdapteeMethod' to Target interface by wrapping it in Adapter's 'Method'
func (a *Adapter) Method() {
	a.adaptee.AdapteeMethod()
}

func main() {
	adaptee := &Adaptee{}          // New Adaptee
	adapter := NewAdapter(adaptee) // New Adapter that takes in adaptee

	adapter.Method() // Call .Method() on adapter which will call the adapted 'AdapteeMethod'
}
