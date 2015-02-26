package ui

import (
	"fmt"
	"image"
	"io/ioutil"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

func loadDataFile(filePath string) string {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func createTexture(rgbaImg *image.NRGBA) (uint32, error) {
	var texture uint32
	gl.GenTextures(1, &texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_R, gl.CLAMP_TO_EDGE)

	// generate base level storage
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA,
		int32(rgbaImg.Bounds().Dx()), int32(rgbaImg.Bounds().Dy()),
		0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgbaImg.Pix))
	// generate required number of mipmaps given texture dimensions
	gl.GenerateMipmap(gl.TEXTURE_2D)

	return texture, nil
}

func pointToScreen(point mgl32.Vec2, out *([2]float32)) {
	out[0] = (2 * point.X()) - 1
	out[1] = (2 * point.Y()) - 1
}

func makeShader(shaderType uint32, source string) uint32 {
	shader := gl.CreateShader(shaderType)
	cSource := gl.Str(source + "\x00")
	gl.ShaderSource(shader, 1, &cSource, nil)
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		panic(fmt.Sprintf("failed to compile %v: %v", source, log))
	}
	return shader
}
