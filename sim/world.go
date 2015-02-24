package sim

import (
	"container/list"
	"time"

	"github.com/acroca/gont/util"
)

// World implements a holw
type World struct {
	Ants       []*Ant
	Pheromones *list.List
	Food       *Food
	Hole       *Hole
	Stop       bool

	MaxPheromones int
}

// NewWorld builds and returns a new world
func NewWorld(maxAnts int) *World {
	maxPheromones := maxAnts * int(pheromoneDuration/antPheromoneFrequency) * 2

	world := &World{
		Ants:          make([]*Ant, maxAnts),
		Pheromones:    list.New(),
		Food:          NewFood(&util.Point{X: 0.2, Y: 0.2}),
		Hole:          NewHole(&util.Point{X: 0.8, Y: 0.8}),
		MaxPheromones: maxPheromones,
	}
	for i := 0; i < maxAnts; i++ {
		world.Ants[i] = NewAnt(world.Hole.Position.Clone())
	}
	return world
}

// Start starts the world simulation
func (world *World) Start() {
	t := time.Now()
	var now time.Time
	var elapsed time.Duration

	for !world.Stop {
		now = time.Now()
		elapsed = now.Sub(t)
		t = now
		world.Step(elapsed)
	}
}

// Step runs a simulation step
func (world *World) Step(elapsed time.Duration) {
	for _, ant := range world.Ants {
		ant.Move(elapsed)
		pheromone := ant.DropPheromone(elapsed)
		if pheromone != nil && world.Pheromones.Len() <= world.MaxPheromones {
			world.Pheromones.PushBack(pheromone)
		}
	}

	var pheromone *Pheromone
	for e := world.Pheromones.Front(); e != nil; e = e.Next() {
		pheromone = e.Value.(*Pheromone)
		pheromone.Decay(elapsed)
		if pheromone.Intensity < 0 {
			world.Pheromones.Remove(e)
		}
	}

}
