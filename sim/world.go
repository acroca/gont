package sim

import (
	"sync"
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
	Mutex      sync.Mutex

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
		world.Ants[i] = NewAnt(util.RandomPoint())
		if i == 0 {
			world.Ants[i].IsRed = true
		}
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
	}

}

// Step runs a simulation step
func (world *World) Step(elapsed time.Duration) {
	world.Mutex.Lock()
	defer world.Mutex.Unlock()
	var pheromone *Pheromone
	for e := world.Pheromones.Front(); e != nil; {
		pheromone = e.Value.(*PheromoneStorageItem).Pheromone
		e = e.Next()
		pheromone.Decay(elapsed)
		if pheromone.Intensity < 0 {
			world.Pheromones.Remove(pheromone)
		}
	}
	for _, ant := range world.Ants {
		ant.Move(elapsed, world.Pheromones.Partition(ant.Position, 0, 0))
		pheromone = ant.DropPheromone(elapsed)
		if pheromone != nil && world.Pheromones.Len() <= world.MaxPheromones {
			world.Pheromones.Add(pheromone)
		}
	}

	util.Stats.Mutex.Lock()
	util.Stats.Steps++
	util.Stats.Mutex.Unlock()
}
