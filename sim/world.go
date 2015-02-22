package sim

import "github.com/acroca/gont/util"

// World implements a holw
type World struct {
	Ants []*Ant
	Food *Food
	Hole *Hole
}

// NewWorld builds and returns a new world
func NewWorld(maxAnts int) *World {
	world := &World{
		Ants: make([]*Ant, maxAnts),
		Food: NewFood(&util.Point{0.2, 0.2}),
		Hole: NewHole(&util.Point{0.8, 0.8}),
	}
	for i := 0; i < maxAnts; i++ {
		world.Ants[i] = NewAnt(world.Hole.Position.Clone())
	}
	return world
}
