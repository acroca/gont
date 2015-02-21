package ui

import (
	"encoding/binary"
	"errors"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"os"
	"runtime"
	"time"
	"unsafe"

	"github.com/acroca/gont/sim"
	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	"github.com/go-gl/glh"
)

const (
	title   = "Gont!"
	width   = 800
	height  = 800
	kindAnt = 0
)

type p struct {
	position  [2]float32
	direction float32
	kind      int32
}

// Window represents the simulator window
type Window struct {
	window *glfw.Window
	world  *sim.World
}

// NewWindow builds and returns the window
func NewWindow(world *sim.World) *Window {
	return &Window{
		world: world,
	}
}

// Open opens the window
func (w *Window) Open() error {
	runtime.LockOSThread()
	var pVar p
	if !glfw.Init() {
		panic("glfw init failed")
	}
	defer glfw.Terminate()

	// use OpenGL 4.0 with deprecated functionality removed
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenglForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenglProfile, glfw.OpenglCoreProfile)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	window.MakeContextCurrent()

	// use vsync
	glfw.SwapInterval(1)

	if gl.Init() != 0 {
		panic("glew init failed")
	}
	gl.GetError() // ignore INVALID_ENUM that GLEW raises when using OpenGL 3.2+

	w.window = window

	vao := gl.GenVertexArray()
	vao.Bind()

	vbo := gl.GenBuffer()
	vbo.Bind(gl.ARRAY_BUFFER)
	points := buildPoints(w.world.Ants)
	gl.BufferData(gl.ARRAY_BUFFER, binary.Size(pVar)*cap(w.world.Ants), points, gl.STATIC_DRAW)

	vShader := glh.Shader{Type: gl.VERTEX_SHADER, Program: loadDataFile("./ui/ant.v.glsl")}
	gShader := glh.Shader{Type: gl.GEOMETRY_SHADER, Program: loadDataFile("./ui/ant.g.glsl")}
	fShader := glh.Shader{Type: gl.FRAGMENT_SHADER, Program: loadDataFile("./ui/ant.f.glsl")}

	program := glh.NewProgram(vShader, gShader, fShader)
	program.Use()

	positionAttrib := program.GetAttribLocation("position")
	positionAttrib.AttribPointer(2, gl.FLOAT, false, binary.Size(pVar), unsafe.Offsetof(pVar.position))
	positionAttrib.EnableArray()
	defer positionAttrib.DisableArray()
	kindAttrib := program.GetAttribLocation("kind")
	kindAttrib.AttribPointer(1, gl.INT, false, binary.Size(pVar), unsafe.Offsetof(pVar.kind))
	kindAttrib.EnableArray()
	defer kindAttrib.DisableArray()
	directionAttrib := program.GetAttribLocation("direction")
	directionAttrib.AttribPointer(1, gl.FLOAT, false, binary.Size(pVar), unsafe.Offsetof(pVar.direction))
	directionAttrib.EnableArray()
	defer directionAttrib.DisableArray()
	gl.ClearColor(0, 0, 0, 1.0)
	// gl.PointSize(10)

	_, err = createTexture("./ui/ant.png")
	if err != nil {
		panic(err)
	}
	texSampler := program.GetUniformLocation("tex")
	gl.ActiveTexture(gl.TEXTURE0)
	texSampler.Uniform1i(0)

	frames := 0
	go func() {
		for {
			fmt.Printf("FPS: %d\n", frames)
			frames = 0
			time.Sleep(1000 * time.Millisecond)
		}
	}()
	width, height := window.GetFramebufferSize()
	gl.Viewport(0, 0, width, height)

	for !w.window.ShouldClose() {
		updatePoints(w.world.Ants, points)
		gl.BufferSubData(gl.ARRAY_BUFFER, 0, binary.Size(pVar)*len(w.world.Ants), points)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		gl.DrawArrays(gl.POINTS, 0, len(w.world.Ants))

		frames++
		w.window.SwapBuffers()
		glfw.PollEvents()

		if w.window.GetKey(glfw.KeyEscape) == glfw.Press {
			w.window.SetShouldClose(true)
		}
	}

	return nil
}

func buildPoints(ants []*sim.Ant) []p {
	res := make([]p, len(ants))
	for idx, ant := range ants {
		res[idx].position[0] = float32(ant.Position.X)
		res[idx].position[1] = float32(ant.Position.Y)
		res[idx].direction = float32(ant.Direction.Angle)
		res[idx].kind = kindAnt
	}
	return res
}

func updatePoints(ants []*sim.Ant, points []p) {
	for idx := range points {
		points[idx].position[0] = float32((2 * ants[idx].Position.X) - 1)
		points[idx].position[1] = float32((2 * ants[idx].Position.Y) - 1)
		points[idx].direction = float32(ants[idx].Direction.Angle)
	}
}

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
