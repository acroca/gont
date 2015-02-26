package ui

import (
	"encoding/binary"
	"unsafe"

	"github.com/acroca/gont/sim"
	"github.com/go-gl/gl/v3.3-core/gl"
)

type holePoint struct {
	position [2]float32
}

// Hole models a ui element for a hole
type Hole struct {
	World    *sim.World
	Program  *Program
	PointVar holePoint
	Points   []holePoint
}

// NewHole Initializes a UI hole
func NewHole(world *sim.World) *Hole {
	hole := &Hole{
		World:  world,
		Points: make([]holePoint, 1),
	}
	pointToScreen(world.Hole.Position, &(hole.Points[0].position))

	attributes := []*Attribute{
		{
			Type:          AttributeFloat,
			AttributeName: "position",
			Amount:        2,
			Stride:        int32(binary.Size(hole.PointVar)),
			Offset:        int(unsafe.Offsetof(hole.PointVar.position)),
		},
	}
	hole.Program = NewProgram(
		binary.Size(hole.PointVar)*len(hole.Points),
		gl.Ptr(hole.Points),
		holeShaders,
		attributes)
	return hole
}

// Render renders the hole
func (hole *Hole) Render() {
	hole.Program.Use()
	gl.DrawArrays(gl.POINTS, 0, int32(len(hole.Points)))
}
