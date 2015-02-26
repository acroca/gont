package sim

import "github.com/go-gl/mathgl/mgl32"

// Food represents an ant
type Food struct {
	Position mgl32.Vec2
}

// NewFood builds and returns a new food
func NewFood(position mgl32.Vec2) *Food {
	return &Food{
		Position: position,
	}
}
