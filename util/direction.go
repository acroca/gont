package util

import (
	"math"
	"math/rand"
)

// Direction implements a direction
type Direction struct {
	Angle float32
}

// Clone clones a direction object
func (direction *Direction) Clone() *Direction {
	return &Direction{
		Angle: direction.Angle,
	}
}

// MirrorX clones a direction object
func (direction *Direction) MirrorX() {
	direction.Angle = (2 * math.Pi) - direction.Angle
}

// MirrorY clones a direction object
func (direction *Direction) MirrorY() {
	direction.Angle = math.Pi - direction.Angle
}

// RandomDirection returns a random direction
func RandomDirection() *Direction {
	return &Direction{
		Angle: rand.Float32() * 2 * math.Pi,
	}
}
