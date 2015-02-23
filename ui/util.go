package ui

import (
	"image"
	"io/ioutil"

	"github.com/acroca/gont/util"
	"github.com/go-gl/gl/v3.3-core/gl"
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

func pointToScreen(point *util.Point, out *([2]float32)) {
	out[0] = float32((2 * point.X) - 1)
	out[1] = float32((2 * point.Y) - 1)
}

func makeShader(shaderType uint32, source string) uint32 {
	shader := gl.CreateShader(shaderType)
	cSource := gl.Str(source + "\x00")
	gl.ShaderSource(shader, 1, &cSource, nil)
	gl.CompileShader(shader)
	return shader
}
