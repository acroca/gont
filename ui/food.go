package ui

import (
	"encoding/binary"
	"unsafe"

	"github.com/acroca/gont/sim"
	"github.com/go-gl/gl/v3.3-core/gl"
)

var (
	foodProgram  uint32
	foodVao      uint32
	foodVbo      uint32
	foodPoints   []foodPoint
	foodPointVar foodPoint
)

type foodPoint struct {
	position [2]float32
}

func initFoodProgram(food *sim.Food) {
	buildFoodPoints(food)

	gl.GenVertexArrays(1, &foodVao)
	gl.BindVertexArray(foodVao)

	gl.GenBuffers(1, &foodVbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, foodVbo)

	gl.BufferData(
		gl.ARRAY_BUFFER,
		binary.Size(foodPointVar)*cap(foodPoints),
		gl.Ptr(foodPoints),
		gl.STATIC_DRAW)

	vShader := makeShader(gl.VERTEX_SHADER, foodV)
	defer gl.DeleteShader(vShader)
	gShader := makeShader(gl.GEOMETRY_SHADER, foodG)
	defer gl.DeleteShader(gShader)
	fShader := makeShader(gl.FRAGMENT_SHADER, foodF)
	defer gl.DeleteShader(fShader)

	foodProgram = gl.CreateProgram()
	gl.AttachShader(foodProgram, vShader)
	gl.AttachShader(foodProgram, gShader)
	gl.AttachShader(foodProgram, fShader)
	gl.LinkProgram(foodProgram)
	gl.ValidateProgram(foodProgram)

	gl.UseProgram(foodProgram)

	positionAttrib := uint32(gl.GetAttribLocation(foodProgram, gl.Str("position\x00")))
	gl.EnableVertexAttribArray(positionAttrib)
	gl.VertexAttribPointer(
		positionAttrib,
		2, gl.FLOAT,
		false,
		int32(binary.Size(foodPointVar)),
		gl.PtrOffset(int(unsafe.Offsetof(foodPointVar.position))))

}

func renderFood() {
	gl.UseProgram(foodProgram)
	gl.BindVertexArray(foodVao)
	gl.BindBuffer(foodVbo, gl.ARRAY_BUFFER)

	gl.DrawArrays(gl.POINTS, 0, int32(len(foodPoints)))
}

func buildFoodPoints(food *sim.Food) {
	foodPoints = make([]foodPoint, 1)
	pointToScreen(food.Position, &foodPoints[0].position)
}
