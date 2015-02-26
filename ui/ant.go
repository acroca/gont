package ui

import (
	"encoding/binary"
	"unsafe"

	"github.com/acroca/gont/sim"
	"github.com/go-gl/gl/v3.3-core/gl"
)

type antPoint struct {
	position  [2]float32
	direction float32
}

// Ants models a ui element for a ant
type Ants struct {
	World    *sim.World
	Program  *Program
	PointVar antPoint
	Points   []antPoint
}

// NewAnts Initializes a UI ants
func NewAnts(world *sim.World) *Ants {
	ants := &Ants{
		World:  world,
		Points: make([]antPoint, len(world.Ants)),
	}

	attributes := []*Attribute{
		{
			Type:          AttributeFloat,
			AttributeName: "position",
			Amount:        2,
			Stride:        int32(binary.Size(ants.PointVar)),
			Offset:        int(unsafe.Offsetof(ants.PointVar.position)),
		},
		{
			Type:          AttributeFloat,
			AttributeName: "direction",
			Amount:        1,
			Stride:        int32(binary.Size(ants.PointVar)),
			Offset:        int(unsafe.Offsetof(ants.PointVar.direction)),
		},
	}
	ants.Program = NewProgram(
		binary.Size(ants.PointVar)*len(ants.Points),
		gl.Ptr(ants.Points),
		antShaders,
		attributes)

	ants.Program.LoadTexture("tex", antTex)

	return ants
}

// Render renders the ants
func (ants *Ants) Render() {
	ants.Program.Use()

	for idx := range ants.World.Ants {
		pointToScreen(ants.World.Ants[idx].Position, &ants.Points[idx].position)
		ants.Points[idx].direction = ants.World.Ants[idx].Direction.Angle
	}

	gl.BufferSubData(gl.ARRAY_BUFFER, 0, binary.Size(ants.PointVar)*len(ants.Points), gl.Ptr(ants.Points))
	gl.DrawArrays(gl.POINTS, 0, int32(len(ants.Points)))
}
