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
			time.Sleep(25 * time.Millisecond)
		}
	}()
	return ant
}

// Move implements the ant movement
func (ant *Ant) Move() {
	ant.Position.X += math.Cos(ant.Direction.Angle) / 1000
	if ant.Position.X > 1 {
		ant.Position.X = 2 - ant.Position.X
		ant.Direction.MirrorY()
	}
	if ant.Position.X < 0 {
		ant.Position.X *= -1
		ant.Direction.MirrorY()
	}
	ant.Position.Y += math.Sin(ant.Direction.Angle) / 1000
	if ant.Position.Y > 1 {
		ant.Position.Y = 2 - ant.Position.Y
		ant.Direction.MirrorX()
	}
	if ant.Position.Y < 0 {
		ant.Position.Y *= -1
		ant.Direction.MirrorX()
	}
	ant.Direction.Angle += (rand.Float64() - 0.5) / 4
}
