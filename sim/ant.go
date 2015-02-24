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

	timeSinceLastPheromone time.Duration
}

// NewAnt builds and returns a new ant
func NewAnt(position *util.Point) *Ant {
	ant := &Ant{
		Position:  position,
		Direction: util.RandomDirection(),
	}
	return ant
}

// Move implements the ant movement
func (ant *Ant) Move(d time.Duration) {
	ant.Position.X += math.Cos(ant.Direction.Angle) * d.Seconds() * antMovementPerSecond
	if ant.Position.X > 1 {
		ant.Position.X = 2 - ant.Position.X
		ant.Direction.MirrorY()
	}
	if ant.Position.X < 0 {
		ant.Position.X *= -1
		ant.Direction.MirrorY()
	}
	ant.Position.Y += math.Sin(ant.Direction.Angle) * d.Seconds() * antMovementPerSecond
	if ant.Position.Y > 1 {
		ant.Position.Y = 2 - ant.Position.Y
		ant.Direction.MirrorX()
	}
	if ant.Position.Y < 0 {
		ant.Position.Y *= -1
		ant.Direction.MirrorX()
	}
	ant.Direction.Angle += ((2 * rand.Float64()) - 1) * d.Seconds() * antMaxRotationPerSecond
}

// DropPheromone implements a pheromone if the ant drops it
func (ant *Ant) DropPheromone(d time.Duration) *Pheromone {
	ant.timeSinceLastPheromone += d
	if ant.timeSinceLastPheromone > antPheromoneFrequency {
		ant.timeSinceLastPheromone -= antPheromoneFrequency
		return NewPheromone(ant.Position.Clone())
	}
	return nil
}
