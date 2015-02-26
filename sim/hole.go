package sim

import "github.com/go-gl/mathgl/mgl32"

// Hole implements a holw
type Hole struct {
	Position mgl32.Vec2
}

// NewHole builds and returns a new hole
func NewHole(position mgl32.Vec2) *Hole {
	return &Hole{
		Position: position,
	}
}
