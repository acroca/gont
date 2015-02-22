package ui

import (
	"encoding/binary"
	"fmt"
	"unsafe"

	"github.com/acroca/gont/sim"
	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
)

var (
	holeProgram gl.Program
	holeVao     gl.VertexArray
	holeVbo     gl.Buffer
	holePoints  []point
)

func initHoleProgram(hole *sim.Hole) {
	holePoints = buildHolePoints(hole)

	holeVao = gl.GenVertexArray()
	holeVao.Bind()

	holeVbo = gl.GenBuffer()
	holeVbo.Bind(gl.ARRAY_BUFFER)
	defer holeVbo.Unbind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, binary.Size(pVar)*len(holePoints), holePoints, gl.STATIC_DRAW)

	vShader := glh.MakeShader(gl.VERTEX_SHADER, loadDataFile("./ui/hole.v.glsl"))
	fmt.Println("a")
	defer vShader.Delete()
	gShader := glh.MakeShader(gl.GEOMETRY_SHADER, loadDataFile("./ui/hole.g.glsl"))
	defer gShader.Delete()
	fShader := glh.MakeShader(gl.FRAGMENT_SHADER, loadDataFile("./ui/hole.f.glsl"))
	defer fShader.Delete()

	holeProgram = gl.CreateProgram()
	holeProgram.AttachShader(vShader)
	holeProgram.AttachShader(gShader)
	holeProgram.AttachShader(fShader)
	holeProgram.Link()
	holeProgram.Validate()

	holeProgram.Use()
	defer holeProgram.Unuse()

	positionAttrib := holeProgram.GetAttribLocation("position")
	positionAttrib.AttribPointer(2, gl.FLOAT, false, binary.Size(pVar), unsafe.Offsetof(pVar.position))
	positionAttrib.EnableArray()

	holeVao.Unbind()
}

func renderHole(hole *sim.Hole) {
	holeProgram.Use()
	defer holeProgram.Unuse()
	holeVao.Bind()
	defer holeVao.Unbind()
	holeVbo.Bind(gl.ARRAY_BUFFER)
	defer holeVbo.Unbind(gl.ARRAY_BUFFER)

	gl.DrawArrays(gl.POINTS, 0, len(holePoints))
}

func buildHolePoints(hole *sim.Hole) []point {
	res := make([]point, 1)
	res[0].position[0] = float32(hole.Position.X)
	res[0].position[1] = float32(hole.Position.Y)
	return res
}
