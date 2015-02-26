package sim

import (
	"time"

	"github.com/acroca/gont/util"
	"github.com/go-gl/mathgl/mgl32"
)

// World implements a holw
type World struct {
	Ants       []*Ant
	Pheromones *PheromoneStorage
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
		Pheromones:    NewPheromoneStorage(maxPheromones),
		Food:          NewFood(mgl32.Vec2{0.2, 0.2}),
		Hole:          NewHole(mgl32.Vec2{0.8, 0.8}),
		MaxPheromones: maxPheromones,
	}
	for i := 0; i < maxAnts; i++ {
		world.Ants[i] = NewAnt(world.Hole.Position)
	}
	return world
}

// Start starts the world simulation
func (world *World) Start() {
	t := time.Now()
	var now time.Time
	var elapsed time.Duration

	for !world.Stop {
		time.Sleep(1 * time.Millisecond)
		now = time.Now()
		elapsed = now.Sub(t)
		t = now
		world.Step(elapsed)
		util.Stats.Steps++
	}

}

// Step runs a simulation step
func (world *World) Step(elapsed time.Duration) {
	var pheromone *Pheromone
	for _, ant := range world.Ants {
		ant.Move(elapsed)
		pheromone = ant.DropPheromone(elapsed)
		if pheromone != nil && world.Pheromones.Len() <= world.MaxPheromones {
			world.Pheromones.Add(pheromone)
		}
	}

	for e := world.Pheromones.Front(); e != nil; e = e.Next() {
		pheromone = e.Value.(*PheromoneStorageItem).Pheromone
		pheromone.Decay(elapsed)
		if pheromone.Intensity < 0 {
			world.Pheromones.Remove(pheromone)
		}
	}
}
