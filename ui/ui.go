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
  
  gl.ClearColor(0, 0, 0, 0.5)
  gl.ClearDepth(1)
  gl.DepthFunc(gl.LEQUAL)

  gl.Viewport(0, 0, Width, Height)
  gl.MatrixMode(gl.PROJECTION)
  gl.LoadIdentity()
  gl.Frustum(-1, 1, -1, 1, 1.0, 10.0)
  gl.MatrixMode(gl.MODELVIEW)
  gl.LoadIdentity()

  gl.PointSize(3.0)

  return
}

func (ui *Ui) destroyScene() {
}

func (ui *Ui) drawScene() {
  gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

  // gl.MatrixMode(gl.MODELVIEW)
  gl.LoadIdentity()
  gl.Translatef(-5, -5, -6)


  for _, points := range ui.World.Points {
    for _, point := range points {
      ui.drawPheromones(point)
      if point.HasFood{
        ui.drawFood(point)
      }
      if len(point.Ants) > 0 {
        ui.drawAnts(point)
      }
    }
  }

}

func (ui *Ui) drawAnts(p *world.Point){
  gl.Begin(gl.POINTS)

  gl.Color3f(0, 1, 0.5)
  baseX := gl.Float((float64(p.X)/float64(ui.World.SizeX))*10) 
  baseY := gl.Float((float64(p.Y)/float64(ui.World.SizeY))*10)

  gl.Vertex2f(baseX, baseY)
  gl.End()
}

func (ui *Ui) drawFood(p *world.Point){
  gl.Begin(gl.POINTS)

  gl.Color3f(0, 1, 0)
  baseX := gl.Float((float64(p.X)/float64(ui.World.SizeX))*10) 
  baseY := gl.Float((float64(p.Y)/float64(ui.World.SizeY))*10)

  gl.Vertex2f(baseX, baseY)
  gl.End()
}

func (ui *Ui) drawPheromones(p *world.Point){

  baseX := gl.Float((float64(p.X)/float64(ui.World.SizeX))*10) 
  baseY := gl.Float((float64(p.Y)/float64(ui.World.SizeY))*10)

  if p.FoodPheromones > 0 {
    gl.Color3f(0, gl.Float(p.FoodPheromones), 0)
  } else if p.PresencePheromones > 0 {
    gl.Color3f(gl.Float(p.PresencePheromones), 0, 0)
  } else {
    return    
  }

  gl.Begin(gl.QUADS)
  // gl.Normal3f(baseX, baseY,1)
  gl.Vertex2f(baseX-0.04, baseY-0.04)
  gl.Vertex2f(baseX+0.04, baseY-0.04)
  gl.Vertex2f(baseX+0.04, baseY+0.04)
  gl.Vertex2f(baseX-0.04, baseY+0.04)
  gl.End()
}