package ui

import (
	"encoding/binary"
	"unsafe"

	"github.com/acroca/gont/sim"
	"github.com/go-gl/gl/v3.3-core/gl"
)

var (
	holeProgram  uint32
	holeVao      uint32
	holeVbo      uint32
	holePoints   []holePoint
	holePointVar holePoint
)

type holePoint struct {
	position [2]float32
}

func initHoleProgram(hole *sim.Hole) {
	buildHolePoints(hole)

	gl.GenVertexArrays(1, &holeVao)
	gl.BindVertexArray(holeVao)

	gl.GenBuffers(1, &holeVbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, holeVbo)

	gl.BufferData(
		gl.ARRAY_BUFFER,
		binary.Size(holePointVar)*cap(holePoints),
		gl.Ptr(holePoints),
		gl.STATIC_DRAW)

	vShader := makeShader(gl.VERTEX_SHADER, holeV)
	defer gl.DeleteShader(vShader)
	gShader := makeShader(gl.GEOMETRY_SHADER, holeG)
	defer gl.DeleteShader(gShader)
	fShader := makeShader(gl.FRAGMENT_SHADER, holeF)
	defer gl.DeleteShader(fShader)

	holeProgram = gl.CreateProgram()
	gl.AttachShader(holeProgram, vShader)
	gl.AttachShader(holeProgram, gShader)
	gl.AttachShader(holeProgram, fShader)
	gl.LinkProgram(holeProgram)
	gl.ValidateProgram(holeProgram)

	gl.UseProgram(holeProgram)

	positionAttrib := uint32(gl.GetAttribLocation(holeProgram, gl.Str("position\x00")))
	gl.EnableVertexAttribArray(positionAttrib)
	gl.VertexAttribPointer(
		positionAttrib,
		2, gl.FLOAT,
		false,
		int32(binary.Size(holePointVar)),
		gl.PtrOffset(int(unsafe.Offsetof(holePointVar.position))))

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

func renderHole() {
	gl.UseProgram(holeProgram)
	gl.BindVertexArray(holeVao)
	gl.BindBuffer(gl.ARRAY_BUFFER, holeVbo)

	gl.DrawArrays(gl.POINTS, 0, int32(len(holePoints)))
}

func buildHolePoints(hole *sim.Hole) {
	holePoints = make([]holePoint, 1)
	pointToScreen(hole.Position, &holePoints[0].position)
}
