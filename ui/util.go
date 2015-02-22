package ui

import (
	"errors"
	"image"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/go-gl/gl"
)

func loadDataFile(filePath string) string {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func createTexture(path string) (gl.Texture, error) {
	r, err := os.Open(path)
	if err != nil {
		return gl.Texture(0), err
	}
	defer r.Close()

	img, err := png.Decode(r)
	if err != nil {
		return gl.Texture(0), err
	}

	rgbaImg, ok := img.(*image.NRGBA)
	if !ok {
		return gl.Texture(0), errors.New("texture must be an NRGBA image")
	}

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
