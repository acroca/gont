package sim

import "github.com/acroca/gont/util"

// Food represents an ant
type Food struct {
	Position *util.Point
}

// NewFood builds and returns a new food
func NewFood(position *util.Point) *Food {
	return &Food{
		Position: position,
	}
}
