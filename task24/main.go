package main

import (
	"fmt"

	"github.com/gmohmad/wbL1/task24/point"
)


func main() {
	p := point.NewPoint(1, 2) // Create a new Point

	fmt.Println(p.Distance(point.NewPoint(3, 2))) // print the distance between p and a new point
}
