package ui

import (
	"fmt"

	"github.com/acroca/gont/sim"
	"github.com/acroca/gont/util"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

const (
	title  = "Gont!"
	width  = 800
	height = 800
)

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
	err := glfw.Init()
	if err != nil {
		panic("failed to initialize glfw")
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLDebugContext, glfw.True)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		panic(err)
	}
	w.window = window

	defer window.Destroy()

	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		fmt.Errorf("Glew init failed: %v", err)
	}

	hole := NewHole(w.world)
	food := NewFood(w.world)
	pheromones := NewPheromones(w.world)
	ants := NewAnts(w.world)

	// use vsync
	glfw.SwapInterval(1)
	gl.ClearColor(1, 1, 1, 1.0)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	// gl.PointSize(10)
	width, height := window.GetFramebufferSize()
	gl.Viewport(0, 0, int32(width), int32(height))

	for !w.window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		pheromones.Render()
		hole.Render()
		food.Render()
		ants.Render()

		util.Stats.Frames++

		w.window.SwapBuffers()
		glfw.PollEvents()

		if w.window.GetKey(glfw.KeyEscape) == glfw.Press {
			w.window.SetShouldClose(true)
		}
	}

	return nil
}
