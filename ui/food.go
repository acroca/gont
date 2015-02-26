package ui

import (
	"encoding/binary"
	"unsafe"

	"github.com/acroca/gont/sim"
	"github.com/go-gl/gl/v3.3-core/gl"
)

type foodPoint struct {
	position [2]float32
}

// Food models a ui element for a food
type Food struct {
	World    *sim.World
	Program  *Program
	PointVar foodPoint
	Points   []foodPoint
}

// NewFood Initializes a UI food
func NewFood(world *sim.World) *Food {
	food := &Food{
		World:  world,
		Points: make([]foodPoint, 1),
	}
	pointToScreen(world.Food.Position, &(food.Points[0].position))

	attributes := []*Attribute{
		{
			Type:          AttributeFloat,
			AttributeName: "position",
			Amount:        2,
			Stride:        int32(binary.Size(food.PointVar)),
			Offset:        int(unsafe.Offsetof(food.PointVar.position)),
		},
	}
	food.Program = NewProgram(
		binary.Size(food.PointVar)*len(food.Points),
		gl.Ptr(food.Points),
		foodShaders,
		attributes)
	return food
}

// Render renders the food
func (food *Food) Render() {
	food.Program.Use()
	gl.DrawArrays(gl.POINTS, 0, int32(len(food.Points)))
}
