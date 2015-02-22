package ui

import (
	"fmt"
	"runtime"
	"time"

	"github.com/acroca/gont/sim"
	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
)

const (
	title   = "Gont!"
	width   = 800
	height  = 800
	kindAnt = 0
)

var (
	pVar point
)

type point struct {
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

	initHoleProgram(w.world.Hole)
	initAntProgram(w.world.Ants)

	gl.ClearColor(0, 0, 0, 1.0)
	// gl.PointSize(10)

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
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		renderAnts(w.world.Ants)
		renderHole(w.world.Hole)

		frames++
		w.window.SwapBuffers()
		glfw.PollEvents()

		if w.window.GetKey(glfw.KeyEscape) == glfw.Press {
			w.window.SetShouldClose(true)
		}
	}

	return nil
}
