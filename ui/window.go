package ui

import (
	"fmt"
	"runtime"
	"time"

	"github.com/acroca/gont/sim"
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

	initPheromoneProgram(w.world.Pheromones, w.world.MaxPheromones)
	initHoleProgram(w.world.Hole)
	initFoodProgram(w.world.Food)
	initAntProgram(w.world.Ants)

	// use vsync
	glfw.SwapInterval(1)
	gl.ClearColor(0, 0, 0, 1.0)
	// gl.PointSize(10)

	frames := 0
	go func() {
		var m runtime.MemStats
		for {
			runtime.ReadMemStats(&m)
			fmt.Printf("FPS: %d\tMem: %d\n", frames, m.Alloc)
			frames = 0
			time.Sleep(1000 * time.Millisecond)
		}
	}()
	width, height := window.GetFramebufferSize()
	gl.Viewport(0, 0, int32(width), int32(height))

	for !w.window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		renderPheromones(w.world.Pheromones, w.world.MaxPheromones)
		renderHole()
		renderFood()
		renderAnts(w.world.Ants)

		frames++
		w.window.SwapBuffers()
		glfw.PollEvents()

		if w.window.GetKey(glfw.KeyEscape) == glfw.Press {
			w.window.SetShouldClose(true)
		}
	}

	return nil
}
