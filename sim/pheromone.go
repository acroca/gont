package sim

import (
	"time"

	"github.com/go-gl/mathgl/mgl32"
)

// Pheromone represents a pheromone
type Pheromone struct {
	Intensity            float32
	Position             mgl32.Vec2
	PheromoneStorageItem *PheromoneStorageItem
}

// NewPheromone builds and returns a new pheromone
func NewPheromone(position mgl32.Vec2) *Pheromone {
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
