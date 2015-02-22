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
	foodPoints = buildFoodPoints(food)

	foodVao = gl.GenVertexArray()
	foodVao.Bind()

	foodVbo = gl.GenBuffer()
	foodVbo.Bind(gl.ARRAY_BUFFER)
	defer foodVbo.Unbind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, binary.Size(foodPointVar)*len(foodPoints), foodPoints, gl.STATIC_DRAW)

	vShader := glh.MakeShader(gl.VERTEX_SHADER, loadDataFile("./ui/food.v.glsl"))
	defer vShader.Delete()
	gShader := glh.MakeShader(gl.GEOMETRY_SHADER, loadDataFile("./ui/food.g.glsl"))
	defer gShader.Delete()
	fShader := glh.MakeShader(gl.FRAGMENT_SHADER, loadDataFile("./ui/food.f.glsl"))
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

func buildFoodPoints(food *sim.Food) []foodPoint {
	res := make([]foodPoint, 1)
	res[0].position[0] = float32(food.Position.X)
	res[0].position[1] = float32(food.Position.Y)
	return res
}
