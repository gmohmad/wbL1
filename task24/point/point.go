package point

import "math"

// Define Point struct
type Point struct {
	x, y float64 // x and y are incapsulated
}

// NewPoint returns a pointer to a new Point object
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// Distance returns the distance between p and other Points
func (p *Point) Distance(other *Point) float64 {
	return math.Sqrt(math.Pow(p.x - other.x, 2) + math.Pow(p.y - other.y, 2))
}
