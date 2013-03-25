package ui

import (
  "fmt"
  gl "github.com/chsc/gogl/gl21"
  "github.com/jteeuwen/glfw"
  "os"
  "gont/world"
)

const (
  Title  = "Spinning Gopher"
  Width  = 1024
  Height = 768
)

type Ui struct{
  World *world.World
}

func NewUI(world *world.World) *Ui {
  return &Ui{
    World: world,
  }
}

func (ui *Ui) Init() {
  if err := glfw.Init(); err != nil {
    fmt.Fprintf(os.Stderr, "glfw: %s\n", err)
    return
  }
  defer glfw.Terminate()

  glfw.OpenWindowHint(glfw.WindowNoResize, 1)
  glfw.OpenWindowHint(glfw.FsaaSamples, 4); // 4x antialiasing

  if err := glfw.OpenWindow(Width, Height, 0, 0, 0, 0, 16, 0, glfw.Windowed); err != nil {
    fmt.Fprintf(os.Stderr, "glfw: %s\n", err)
    return
  }
  defer glfw.CloseWindow()

  glfw.SetSwapInterval(1)
  glfw.SetWindowTitle(Title)

  if err := gl.Init(); err != nil {
    fmt.Fprintf(os.Stderr, "gl: %s\n", err)
  }

  if err := ui.initScene(); err != nil {
    fmt.Fprintf(os.Stderr, "init: %s\n", err)
    return
  }
  defer ui.destroyScene()

  for glfw.WindowParam(glfw.Opened) == 1 {
    ui.drawScene()
    glfw.SwapBuffers()
  }
}

func (ui *Ui) initScene() (err error) {
  gl.Enable(gl.MULTISAMPLE)
  gl.Enable(gl.DEPTH_TEST)
  gl.Enable(gl.TEXTURE_2D)
  gl.Enable(gl.BLEND)

  gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA);
  
  gl.ClearColor(0, 0, 0, 0.5)
  gl.ClearDepth(1)
  gl.DepthFunc(gl.LEQUAL)

  gl.Viewport(0, 0, Width, Height)
  gl.MatrixMode(gl.PROJECTION)
  gl.LoadIdentity()
  gl.Frustum(-1, 1, -1, 1, 1.0, 10.0)
  gl.MatrixMode(gl.MODELVIEW)
  gl.LoadIdentity()

  return
}

func (ui *Ui) destroyScene() {
}

func (ui *Ui) drawScene() {
  gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

  // gl.MatrixMode(gl.MODELVIEW)
  gl.LoadIdentity()
  gl.Translatef(-5, -5, -5.1)

  for e := ui.World.Holes.Front() ; e != nil ; e = e.Next() {
    hole := e.Value.(*world.Hole)
    ui.drawHole(hole)
  }
  for e := ui.World.Pheromones.Front() ; e != nil ; e = e.Next() {
    pheromone := e.Value.(*world.Pheromone)
    ui.drawPheromone(pheromone)
  }
  for e := ui.World.Food.Front() ; e != nil ; e = e.Next() {
    food := e.Value.(*world.Food)
    ui.drawFood(food)
  }
  for e := ui.World.Ants.Front() ; e != nil ; e = e.Next() {
    ant := e.Value.(*world.Ant)
    ui.drawAnts(ant)
  }
}

func (ui *Ui) drawHole(h *world.Hole){
  p := h.Point
  gl.PointSize(gl.Float(3700.0 / float64(ui.World.SizeX)))
  baseX := gl.Float((float64(p.X)/float64(ui.World.SizeX))*10) 
  baseY := gl.Float((float64(p.Y)/float64(ui.World.SizeY))*10)

  gl.Begin(gl.POINTS)

  gl.Color3f(0.7, 0.3, 0.3)
  gl.Vertex2f(baseX, baseY)
  gl.End()

  gl.PointSize(gl.Float(1200.0 / float64(ui.World.SizeX)))

  gl.Begin(gl.POINTS)
  gl.Color3f(0,0,0)
  gl.Vertex2f(baseX, baseY)

  gl.End()
}

func (ui *Ui) drawPheromone(p *world.Pheromone){
  point := p.Point
  gl.PointSize(gl.Float(300.0 / float64(ui.World.SizeX)))
  baseX := gl.Float((float64(point.X)/float64(ui.World.SizeX))*10)
  baseY := gl.Float((float64(point.Y)/float64(ui.World.SizeY))*10)
  intensity := gl.Float((float64(p.Amount) / float64(world.MAX_AMOUNT)) + 0.1)

  gl.Begin(gl.POINTS)

  gl.Color4f(0.9, 0.7, 0.3, intensity)
  gl.Vertex2f(baseX, baseY)
  gl.End()
}

func (ui *Ui) drawFood(f *world.Food){
  p := f.Point
  gl.PointSize(gl.Float(1000.0 / float64(ui.World.SizeX)))
  baseX := gl.Float((float64(p.X)/float64(ui.World.SizeX))*10)
  baseY := gl.Float((float64(p.Y)/float64(ui.World.SizeY))*10)

  gl.Begin(gl.POINTS)

  gl.Color3f(0, 1, 0)
  gl.Vertex2f(baseX, baseY)
  gl.End()
}

func (ui *Ui) drawAnts(a *world.Ant){
  p := a.Vector.Point
  gl.PointSize(gl.Float(700.0 / float64(ui.World.SizeX)))
  gl.Begin(gl.POINTS)

  if a.HasFood {
    gl.Color3f(0.3, 0.7, 0.5)
  } else {
    gl.Color3f(1, 0.2, 0.5)
  }
  baseX := gl.Float((float64(p.X)/float64(ui.World.SizeX))*10) 
  baseY := gl.Float((float64(p.Y)/float64(ui.World.SizeY))*10)

  gl.Vertex2f(baseX, baseY)
  gl.End()
}