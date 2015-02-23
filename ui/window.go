package ui

import (
	"fmt"
	"time"

	"github.com/acroca/gont/sim"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.0/glfw"
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
	if !glfw.Init() {
		panic("failed to initialize glfw")
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenglForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenglProfile, glfw.OpenglCoreProfile)
	glfw.WindowHint(glfw.OpenglDebugContext, glfw.True)

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

	initHoleProgram(w.world.Hole)
	initFoodProgram(w.world.Food)
	initAntProgram(w.world.Ants)

	// use vsync
	glfw.SwapInterval(1)
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
	gl.Viewport(0, 0, int32(width), int32(height))

	for !w.window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		renderAnts(w.world.Ants)
		renderHole()
		renderFood()

		frames++
		w.window.SwapBuffers()
		glfw.PollEvents()

		if w.window.GetKey(glfw.KeyEscape) == glfw.Press {
			w.window.SetShouldClose(true)
		}
	}

	return nil
}
