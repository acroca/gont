package sim

import (
	"time"

	"github.com/acroca/gont/util"
)

// Pheromone represents a pheromone
type Pheromone struct {
	Intensity float32
	Position  *util.Point
}

// NewPheromone builds and returns a new pheromone
func NewPheromone(position *util.Point) *Pheromone {
	pheromone := &Pheromone{
		Position:  position,
		Intensity: 1.0,
	}
	return pheromone
}

// Decay decrements the pheromone intensity for the given time
func (pheromone *Pheromone) Decay(d time.Duration) {
	pheromone.Intensity -= float32(d) / float32(pheromoneDuration)
}
