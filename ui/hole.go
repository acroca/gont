package ui

import (
	"encoding/binary"
	"unsafe"

	"github.com/acroca/gont/sim"
	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
)

var (
	holeProgram  gl.Program
	holeVao      gl.VertexArray
	holeVbo      gl.Buffer
	holePoints   []holePoint
	holePointVar holePoint
)

type holePoint struct {
	position [2]float32
}

func initHoleProgram(hole *sim.Hole) {
	buildHolePoints(hole)

	holeVao = gl.GenVertexArray()
	holeVao.Bind()

	holeVbo = gl.GenBuffer()
	holeVbo.Bind(gl.ARRAY_BUFFER)
	defer holeVbo.Unbind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, binary.Size(holePointVar)*len(holePoints), holePoints, gl.STATIC_DRAW)

	vShader := glh.MakeShader(gl.VERTEX_SHADER, holeV)
	defer vShader.Delete()
	gShader := glh.MakeShader(gl.GEOMETRY_SHADER, holeG)
	defer gShader.Delete()
	fShader := glh.MakeShader(gl.FRAGMENT_SHADER, holeF)
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
	positionAttrib.AttribPointer(2, gl.FLOAT, false, binary.Size(holePointVar), unsafe.Offsetof(holePointVar.position))
	positionAttrib.EnableArray()

	holeVao.Unbind()
}

func renderHole() {
	holeProgram.Use()
	defer holeProgram.Unuse()
	holeVao.Bind()
	defer holeVao.Unbind()
	holeVbo.Bind(gl.ARRAY_BUFFER)
	defer holeVbo.Unbind(gl.ARRAY_BUFFER)

	gl.DrawArrays(gl.POINTS, 0, len(holePoints))
}

func buildHolePoints(hole *sim.Hole) {
	holePoints = make([]holePoint, 1)
	pointToScreen(hole.Position, &holePoints[0].position)
}
