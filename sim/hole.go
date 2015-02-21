package sim

import "github.com/acroca/gont/util"

// Hole implements a holw
type Hole struct {
	Position *util.Point
}

// NewHole builds and returns a new hole
func NewHole(position *util.Point) *Hole {
	return &Hole{
		Position: position,
	}
}
