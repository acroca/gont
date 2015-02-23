package ui

import (
	"encoding/binary"
	"unsafe"

	"github.com/acroca/gont/sim"
	"github.com/go-gl/gl/v3.3-core/gl"
)

var (
	antProgram  uint32
	antVao      uint32
	antVbo      uint32
	antPoints   []antPoint
	antPointVar antPoint
)

type antPoint struct {
	position  [2]float32
	direction float32
}

func initAntProgram(ants []*sim.Ant) {
	buildAntPoints(ants)

	gl.GenVertexArrays(1, &antVao)
	gl.BindVertexArray(antVao)

	gl.GenBuffers(1, &antVbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, antVbo)

	gl.BufferData(
		gl.ARRAY_BUFFER,
		binary.Size(antPointVar)*cap(antPoints),
		nil,
		gl.STREAM_DRAW)

	vShader := makeShader(gl.VERTEX_SHADER, antV)
	defer gl.DeleteShader(vShader)
	gShader := makeShader(gl.GEOMETRY_SHADER, antG)
	defer gl.DeleteShader(gShader)
	fShader := makeShader(gl.FRAGMENT_SHADER, antF)
	defer gl.DeleteShader(fShader)

	antProgram = gl.CreateProgram()
	gl.AttachShader(antProgram, vShader)
	gl.AttachShader(antProgram, gShader)
	gl.AttachShader(antProgram, fShader)
	gl.LinkProgram(antProgram)
	gl.ValidateProgram(antProgram)

	gl.UseProgram(antProgram)

	positionAttrib := uint32(gl.GetAttribLocation(antProgram, gl.Str("position\x00")))
	gl.EnableVertexAttribArray(positionAttrib)
	gl.VertexAttribPointer(
		positionAttrib,
		2, gl.FLOAT,
		false,
		int32(binary.Size(antPointVar)),
		gl.PtrOffset(int(unsafe.Offsetof(antPointVar.position))))

	directionAttrib := uint32(gl.GetAttribLocation(antProgram, gl.Str("direction\x00")))
	gl.EnableVertexAttribArray(directionAttrib)
	gl.VertexAttribPointer(
		directionAttrib,
		1, gl.FLOAT,
		false,
		int32(binary.Size(antPointVar)),
		gl.PtrOffset(int(unsafe.Offsetof(antPointVar.direction))))
	// defer gl.DisableVertexAttribArray(directionAttrib)

	tex, err := createTexture(antTex)
	if err != nil {
		panic(err)
	}
	defer gl.DeleteTextures(1, &tex)

	texSampler := gl.GetUniformLocation(antProgram, gl.Str("tex\x00"))
	gl.ActiveTexture(gl.TEXTURE0)
	gl.Uniform1i(texSampler, 0)

}

func renderAnts(ants []*sim.Ant) {
	gl.UseProgram(antProgram)
	gl.BindVertexArray(antVao)
	gl.BindBuffer(antVbo, gl.ARRAY_BUFFER)

	updateAntPoints(ants)

	gl.BufferSubData(gl.ARRAY_BUFFER, 0, binary.Size(antPointVar)*len(antPoints), gl.Ptr(antPoints))
	gl.DrawArrays(gl.POINTS, 0, int32(len(antPoints)))
}

func buildAntPoints(ants []*sim.Ant) {
	antPoints = make([]antPoint, len(ants))
	updateAntPoints(ants)
}

func updateAntPoints(ants []*sim.Ant) {
	for idx := range antPoints {
		pointToScreen(ants[idx].Position, &antPoints[idx].position)
		antPoints[idx].direction = float32(ants[idx].Direction.Angle)
	}
}
