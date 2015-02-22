package ui

import (
	"encoding/binary"
	"unsafe"

	"github.com/acroca/gont/sim"
	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
)

var (
	antProgram gl.Program
	antVao     gl.VertexArray
	antVbo     gl.Buffer
	antPoints  []point
)

func initAntProgram(ants []*sim.Ant) {
	antPoints = buildAntPoints(ants)

	antVao = gl.GenVertexArray()
	antVao.Bind()

	antVbo = gl.GenBuffer()
	antVbo.Bind(gl.ARRAY_BUFFER)
	defer antVbo.Unbind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, binary.Size(pVar)*cap(antPoints), antPoints, gl.STREAM_DRAW)

	vShader := glh.MakeShader(gl.VERTEX_SHADER, loadDataFile("./ui/ant.v.glsl"))
	defer vShader.Delete()
	gShader := glh.MakeShader(gl.GEOMETRY_SHADER, loadDataFile("./ui/ant.g.glsl"))
	defer gShader.Delete()
	fShader := glh.MakeShader(gl.FRAGMENT_SHADER, loadDataFile("./ui/ant.f.glsl"))
	defer fShader.Delete()

	antProgram = gl.CreateProgram()
	antProgram.AttachShader(vShader)
	antProgram.AttachShader(gShader)
	antProgram.AttachShader(fShader)
	antProgram.Link()
	antProgram.Validate()

	antProgram.Use()
	defer antProgram.Unuse()

	positionAttrib := antProgram.GetAttribLocation("position")
	positionAttrib.AttribPointer(2, gl.FLOAT, false, binary.Size(pVar), unsafe.Offsetof(pVar.position))
	positionAttrib.EnableArray()
	defer positionAttrib.DisableArray()
	kindAttrib := antProgram.GetAttribLocation("kind")
	kindAttrib.AttribPointer(1, gl.INT, false, binary.Size(pVar), unsafe.Offsetof(pVar.kind))
	kindAttrib.EnableArray()
	defer kindAttrib.DisableArray()
	directionAttrib := antProgram.GetAttribLocation("direction")
	directionAttrib.AttribPointer(1, gl.FLOAT, false, binary.Size(pVar), unsafe.Offsetof(pVar.direction))
	directionAttrib.EnableArray()
	defer directionAttrib.DisableArray()

	tex, err := createTexture("./ui/ant.png")
	if err != nil {
		panic(err)
	}
	defer tex.Delete()
	texSampler := antProgram.GetUniformLocation("tex")
	gl.ActiveTexture(gl.TEXTURE0)
	texSampler.Uniform1i(0)

	antVao.Unbind()
}

func renderAnts(ants []*sim.Ant) {
	antProgram.Use()
	defer antProgram.Unuse()
	antVao.Bind()
	defer antVao.Unbind()
	antVbo.Bind(gl.ARRAY_BUFFER)
	defer antVbo.Unbind(gl.ARRAY_BUFFER)

	updateAntPoints(ants)
	gl.BufferSubData(gl.ARRAY_BUFFER, 0, binary.Size(pVar)*len(antPoints), antPoints)
	gl.DrawArrays(gl.POINTS, 0, len(antPoints))
}

func buildAntPoints(ants []*sim.Ant) []point {
	res := make([]point, len(ants))
	for idx, ant := range ants {
		res[idx].position[0] = float32(ant.Position.X)
		res[idx].position[1] = float32(ant.Position.Y)
		res[idx].direction = float32(ant.Direction.Angle)
		res[idx].kind = kindAnt
	}
	return res
}

func updateAntPoints(ants []*sim.Ant) {
	for idx := range antPoints {
		antPoints[idx].position[0] = float32((2 * ants[idx].Position.X) - 1)
		antPoints[idx].position[1] = float32((2 * ants[idx].Position.Y) - 1)
		antPoints[idx].direction = float32(ants[idx].Direction.Angle)
	}
}
