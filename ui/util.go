package ui

import (
	"image"
	"io/ioutil"

	"github.com/acroca/gont/util"
	"github.com/go-gl/gl"
)

func loadDataFile(filePath string) string {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func createTexture(rgbaImg *image.NRGBA) (gl.Texture, error) {
	texture := gl.GenTexture()
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_R, gl.CLAMP_TO_EDGE)

	// generate base level storage
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA,
		rgbaImg.Bounds().Dx(), rgbaImg.Bounds().Dy(),
		0, gl.RGBA, gl.UNSIGNED_BYTE, rgbaImg.Pix)
	// generate required number of mipmaps given texture dimensions
	gl.GenerateMipmap(gl.TEXTURE_2D)

	return texture, nil
}

func pointToScreen(point *util.Point, out *[2]float32) {
	out[0] = float32((2 * point.X) - 1)
	out[1] = float32((2 * point.Y) - 1)
}
