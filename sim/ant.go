package sim

import (
	"math"
	"math/rand"
	"time"

	"github.com/acroca/gont/util"
	"github.com/go-gl/mathgl/mgl32"
)

// Ant represents an ant
type Ant struct {
	Position  mgl32.Vec2
	Direction *util.Direction

	timeSinceLastPheromone time.Duration
}

// NewAnt builds and returns a new ant
func NewAnt(position mgl32.Vec2) *Ant {
	ant := &Ant{
		Position:  position,
		Direction: util.RandomDirection(),
	}
	return ant
}

// Move implements the ant movement
func (ant *Ant) Move(d time.Duration) {
	ant.Position = ant.Position.Add(mgl32.Vec2{
		float32(math.Cos(float64(ant.Direction.Angle)) * d.Seconds() * antMovementPerSecond),
		float32(math.Sin(float64(ant.Direction.Angle)) * d.Seconds() * antMovementPerSecond),
	})
	if ant.Position.X() > 1 {
		ant.Position = ant.Position.Add(mgl32.Vec2{1 - ant.Position.X(), 0})
		ant.Direction.MirrorY()
	}
	if ant.Position.X() < 0 {
		ant.Position = ant.Position.Add(mgl32.Vec2{-ant.Position.X(), 0})
		ant.Direction.MirrorY()
	}
	if ant.Position.Y() > 1 {
		ant.Position = ant.Position.Add(mgl32.Vec2{0, 1 - ant.Position.Y()})
		ant.Direction.MirrorX()
	}
	if ant.Position.Y() < 0 {
		ant.Position = ant.Position.Add(mgl32.Vec2{0, -ant.Position.Y()})
		ant.Direction.MirrorX()
	}
	ant.Direction.Angle += float32(((2 * rand.Float64()) - 1) * d.Seconds() * antMaxRotationPerSecond)
}

// DropPheromone implements a pheromone if the ant drops it
func (ant *Ant) DropPheromone(d time.Duration) *Pheromone {
	ant.timeSinceLastPheromone += d
	if ant.timeSinceLastPheromone > antPheromoneFrequency {
		ant.timeSinceLastPheromone -= antPheromoneFrequency
		return NewPheromone(ant.Position)
	}
	return nil
}
