package util

import (
	"math"
	"math/rand"
)

// Point implements a point
type Point struct {
	X float64
	Y float64
}

// Clone clones the point
func (point *Point) Clone() *Point {
	return &Point{
		X: point.X,
		Y: point.Y,
	}
}

// Distance returns the distance between two points
func Distance(p1 *Point, p2 *Point) float64 {
	return math.Sqrt(math.Pow(p1.X-p2.X, 2.0) + math.Pow(p1.Y-p2.Y, 2.0))
}

// RandomPoint returns a random point
func RandomPoint() *Point {
	return &Point{
		X: rand.Float64(),
		Y: rand.Float64(),
	}
}
