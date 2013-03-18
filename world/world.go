package world

import (
)

type World struct {
  SizeX int
  SizeY int
  Points []([]*Point)
}

var (
  WORLD = initializeWorld()
)

func initializeWorld() (w *World) {
  world := &World{
    SizeX: 40,
    SizeY: 20,
  }
  
  world.Points = make([]([]*Point), world.SizeX)
  for i:=0;i<world.SizeX;i++ {
    world.Points[i] = make([]*Point, world.SizeY)
    for j:=0;j<world.SizeY;j++ {
      world.Points[i][j] = &Point{X: i, Y: j, Ants: make([]*Ant, 0)}
    }
  }
  return world
}

func (w *World) ToString() string {
  r := ""
  for i:=0;i<w.SizeX+2;i++ {
    r += "+"
  }
  r += "\n"
  for j:=0;j<w.SizeY;j++ {
    r += "+"
    for i:=0;i<w.SizeX;i++ {
      r += w.Points[i][j].ToString()
    }
    r += "+\n"
  }
  for i:=0;i<w.SizeX+2;i++ {
    r += "+"
  }
  return r
}
