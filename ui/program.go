package ui

import (
	"image"
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
)

// Program models a OpenGL Program
type Program struct {
	Program uint32
	Vao     uint32
	Vbo     uint32
}

// Shader models all shaders for a program
type Shader struct {
	Vertex   string
	Geometry string
	Fragment string
}

const (
	// AttributeFloat is a float attribute
	AttributeFloat = 0
)

// Attribute models an attribute
type Attribute struct {
	Type          int
	AttributeName string
	Amount        int32
	Stride        int32
	Offset        int
}

// NewProgram builds an returns a program
func NewProgram(bufferSize int, dataPointer unsafe.Pointer, shaders *Shader, attributes []*Attribute) *Program {
	var program Program

	gl.GenVertexArrays(1, &program.Vao)
	gl.BindVertexArray(program.Vao)

	gl.GenBuffers(1, &program.Vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, program.Vbo)

	gl.BufferData(
		gl.ARRAY_BUFFER,
		bufferSize,
		dataPointer,
		gl.STATIC_DRAW)

	program.Program = gl.CreateProgram()

	if shaders.Vertex != "" {
		vShader := makeShader(gl.VERTEX_SHADER, shaders.Vertex)
		gl.AttachShader(program.Program, vShader)
		defer gl.DeleteShader(vShader)
	}

	if shaders.Geometry != "" {
		gShader := makeShader(gl.GEOMETRY_SHADER, shaders.Geometry)
		gl.AttachShader(program.Program, gShader)
		defer gl.DeleteShader(gShader)
	}

	if shaders.Fragment != "" {
		fShader := makeShader(gl.FRAGMENT_SHADER, shaders.Fragment)
		gl.AttachShader(program.Program, fShader)
		defer gl.DeleteShader(fShader)
	}

	gl.LinkProgram(program.Program)
	gl.ValidateProgram(program.Program)
	gl.UseProgram(program.Program)

	for _, attribute := range attributes {
		if attribute.Type == AttributeFloat {
			program.AddFloatAttrib(attribute)
		}
	}
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	return &program
}

// Use uses the program and binds the vao and vbo
func (program *Program) Use() {
	gl.UseProgram(program.Program)
	gl.BindVertexArray(program.Vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, program.Vbo)
}

// AddFloatAttrib registers a float attribute
func (program *Program) AddFloatAttrib(attribute *Attribute) {
	attrib := uint32(gl.GetAttribLocation(program.Program, gl.Str(attribute.AttributeName+"\x00")))
	gl.EnableVertexAttribArray(attrib)
	gl.VertexAttribPointer(attrib, attribute.Amount, gl.FLOAT, false, attribute.Stride, gl.PtrOffset(attribute.Offset))

}

// LoadTexture registers a float attribute
func (program *Program) LoadTexture(attrName string, rgbaImg *image.NRGBA) {
	tex, err := createTexture(rgbaImg)
	if err != nil {
		panic(err)
	}
	defer gl.DeleteTextures(1, &tex)

	texSampler := gl.GetUniformLocation(program.Program, gl.Str(attrName+"\x00"))
	gl.ActiveTexture(gl.TEXTURE0)
	gl.Uniform1i(texSampler, 0)
}
