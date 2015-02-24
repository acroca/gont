package ui

import (
	"container/list"
	"encoding/binary"
	"unsafe"

	"github.com/acroca/gont/sim"
	"github.com/go-gl/gl/v3.3-core/gl"
)

var (
	pheromoneProgram  uint32
	pheromoneVao      uint32
	pheromoneVbo      uint32
	pheromonePoints   []pheromonePoint
	pheromonePointVar pheromonePoint
)

type pheromonePoint struct {
	position  [2]float32
	intensity float32
}

func initPheromoneProgram(pheromones *list.List, maxPheromones int) {
	buildPheromonePoints(pheromones, maxPheromones)

	gl.GenVertexArrays(1, &pheromoneVao)
	gl.BindVertexArray(pheromoneVao)

	gl.GenBuffers(1, &pheromoneVbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, pheromoneVbo)

	gl.BufferData(
		gl.ARRAY_BUFFER,
		binary.Size(pheromonePointVar)*cap(pheromonePoints),
		nil,
		gl.DYNAMIC_DRAW)

	vShader := makeShader(gl.VERTEX_SHADER, pheromoneV)
	defer gl.DeleteShader(vShader)
	gShader := makeShader(gl.GEOMETRY_SHADER, pheromoneG)
	defer gl.DeleteShader(gShader)
	fShader := makeShader(gl.FRAGMENT_SHADER, pheromoneF)
	defer gl.DeleteShader(fShader)

	pheromoneProgram = gl.CreateProgram()
	gl.AttachShader(pheromoneProgram, vShader)
	gl.AttachShader(pheromoneProgram, gShader)
	gl.AttachShader(pheromoneProgram, fShader)
	gl.LinkProgram(pheromoneProgram)
	gl.ValidateProgram(pheromoneProgram)

	gl.UseProgram(pheromoneProgram)

	positionAttrib := uint32(gl.GetAttribLocation(pheromoneProgram, gl.Str("position\x00")))
	gl.EnableVertexAttribArray(positionAttrib)
	gl.VertexAttribPointer(
		positionAttrib,
		2, gl.FLOAT,
		false,
		int32(binary.Size(pheromonePointVar)),
		gl.PtrOffset(int(unsafe.Offsetof(pheromonePointVar.position))))

	intensityAttrib := uint32(gl.GetAttribLocation(pheromoneProgram, gl.Str("intensity\x00")))
	gl.EnableVertexAttribArray(intensityAttrib)
	gl.VertexAttribPointer(
		intensityAttrib,
		1, gl.FLOAT,
		false,
		int32(binary.Size(pheromonePointVar)),
		gl.PtrOffset(int(unsafe.Offsetof(pheromonePointVar.intensity))))

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

func renderPheromones(pheromones *list.List, maxPheromones int) {
	gl.UseProgram(pheromoneProgram)
	gl.BindVertexArray(pheromoneVao)
	gl.BindBuffer(gl.ARRAY_BUFFER, pheromoneVbo)

	updatePheromonePoints(pheromones, maxPheromones)
	gl.BufferSubData(gl.ARRAY_BUFFER, 0, binary.Size(pheromonePointVar)*pheromones.Len(), gl.Ptr(pheromonePoints))
	gl.DrawArrays(gl.POINTS, 0, int32(pheromones.Len()))
}

func buildPheromonePoints(pheromones *list.List, maxPheromones int) {
	pheromonePoints = make([]pheromonePoint, maxPheromones)
	idx := 0
	var pheromone *sim.Pheromone
	for e := pheromones.Front(); e != nil && idx < maxPheromones; e = e.Next() {
		pheromone = e.Value.(*sim.Pheromone)
		pointToScreen(pheromone.Position, &pheromonePoints[idx].position)
		pheromonePoints[idx].intensity = pheromone.Intensity
		idx++
	}
}

func updatePheromonePoints(pheromones *list.List, maxPheromones int) {
	idx := 0
	var pheromone *sim.Pheromone
	for e := pheromones.Front(); e != nil && idx < maxPheromones; e = e.Next() {
		pheromone = e.Value.(*sim.Pheromone)
		pointToScreen(pheromone.Position, &pheromonePoints[idx].position)
		pheromonePoints[idx].intensity = pheromone.Intensity
		idx++
	}
}
