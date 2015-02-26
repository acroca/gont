package ui

import (
	"encoding/binary"
	"unsafe"

	"github.com/acroca/gont/sim"
	"github.com/go-gl/gl/v3.3-core/gl"
)

type pheromonePoint struct {
	intensity float32
	position  [2]float32
}

// Pheromones models a ui element for a pheromone
type Pheromones struct {
	World    *sim.World
	Program  *Program
	PointVar pheromonePoint
	Points   []pheromonePoint
}

// NewPheromones Initializes a UI pheromone
func NewPheromones(world *sim.World) *Pheromones {
	pheromones := &Pheromones{
		World:  world,
		Points: make([]pheromonePoint, world.MaxPheromones),
	}

	attributes := []*Attribute{
		{
			Type:          AttributeFloat,
			AttributeName: "position",
			Amount:        2,
			Stride:        int32(binary.Size(pheromones.PointVar)),
			Offset:        int(unsafe.Offsetof(pheromones.PointVar.position)),
		},
		{
			Type:          AttributeFloat,
			AttributeName: "intensity",
			Amount:        1,
			Stride:        int32(binary.Size(pheromones.PointVar)),
			Offset:        int(unsafe.Offsetof(pheromones.PointVar.intensity)),
		},
	}
	pheromones.Program = NewProgram(
		binary.Size(pheromones.PointVar)*len(pheromones.Points),
		gl.Ptr(pheromones.Points),
		pheromoneShaders,
		attributes)

	return pheromones
}

// Render renders the pheromones
func (pheromones *Pheromones) Render() {
	pheromones.Program.Use()

	idx := 0
	var simPheromone *sim.Pheromone
	for e := pheromones.World.Pheromones.Front(); e != nil; e = e.Next() {
		simPheromone = e.Value.(*sim.PheromoneStorageItem).Pheromone
		pointToScreen(simPheromone.Position, &pheromones.Points[idx].position)
		pheromones.Points[idx].intensity = simPheromone.Intensity
		idx++
	}

	gl.BufferSubData(gl.ARRAY_BUFFER, 0, binary.Size(pheromones.PointVar)*pheromones.World.Pheromones.Len(), gl.Ptr(pheromones.Points))
	gl.DrawArrays(gl.POINTS, 0, int32(pheromones.World.Pheromones.Len()))
}
