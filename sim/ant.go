package sim

import (
	"math"
	"math/rand"
	"time"

	"github.com/acroca/gont/util"
)

// Ant represents an ant
type Ant struct {
	Position  *util.Point
	Direction *util.Direction
}

// NewAnt builds and returns a new ant
func NewAnt(position *util.Point) *Ant {
	ant := &Ant{
		Position:  position,
		Direction: util.RandomDirection(),
	}
	go func() {
		for {
			ant.Move()
			time.Sleep(50 * time.Millisecond)
		}
	}()
	return ant
}

// Move implements the ant movement
func (ant *Ant) Move() {
	ant.Position.X += math.Cos(ant.Direction.Angle) / 1000
	ant.Position.Y += math.Sin(ant.Direction.Angle) / 1000
	ant.Direction.Angle += (rand.Float64() - 0.5) / 4
}
