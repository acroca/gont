package ui

import (
	"encoding/binary"
	"unsafe"

	"github.com/acroca/gont/sim"
	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
)

var (
	foodProgram  gl.Program
	foodVao      gl.VertexArray
	foodVbo      gl.Buffer
	foodPoints   []foodPoint
	foodPointVar foodPoint
)

type foodPoint struct {
	position [2]float32
}

func initFoodProgram(food *sim.Food) {
	buildFoodPoints(food)

	foodVao = gl.GenVertexArray()
	foodVao.Bind()

	foodVbo = gl.GenBuffer()
	foodVbo.Bind(gl.ARRAY_BUFFER)
	defer foodVbo.Unbind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, binary.Size(foodPointVar)*len(foodPoints), foodPoints, gl.STATIC_DRAW)

	vShader := glh.MakeShader(gl.VERTEX_SHADER, foodV)
	defer vShader.Delete()
	gShader := glh.MakeShader(gl.GEOMETRY_SHADER, foodG)
	defer gShader.Delete()
	fShader := glh.MakeShader(gl.FRAGMENT_SHADER, foodF)
	defer fShader.Delete()

	foodProgram = gl.CreateProgram()
	foodProgram.AttachShader(vShader)
	foodProgram.AttachShader(gShader)
	foodProgram.AttachShader(fShader)
	foodProgram.Link()
	foodProgram.Validate()

	foodProgram.Use()
	defer foodProgram.Unuse()

	positionAttrib := foodProgram.GetAttribLocation("position")
	positionAttrib.AttribPointer(2, gl.FLOAT, false, binary.Size(foodPointVar), unsafe.Offsetof(foodPointVar.position))
	positionAttrib.EnableArray()

	foodVao.Unbind()
}

func renderFood() {
	foodProgram.Use()
	defer foodProgram.Unuse()
	foodVao.Bind()
	defer foodVao.Unbind()
	foodVbo.Bind(gl.ARRAY_BUFFER)
	defer foodVbo.Unbind(gl.ARRAY_BUFFER)

	gl.DrawArrays(gl.POINTS, 0, len(foodPoints))
}

func buildFoodPoints(food *sim.Food) {
	foodPoints = make([]foodPoint, 1)
	pointToScreen(food.Position, &foodPoints[0].position)
}
