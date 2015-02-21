package util

import (
	"math"
	"math/rand"
)

// Direction implements a direction
type Direction struct {
	Angle float64
}

// Clone clones a direction object
func (direction *Direction) Clone() *Direction {
	return &Direction{
		Angle: direction.Angle,
	}
}

// RandomDirection returns a random direction
func RandomDirection() *Direction {
	return &Direction{
		Angle: rand.Float64() * 2 * math.Pi,
	}
}
