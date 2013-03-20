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

const (
  ANTS = 100
  SIZE_X = 173
  SIZE_Y = 45
)

func initializeWorld() (w *World) {
  world := &World{
    SizeX: SIZE_X,
    SizeY: SIZE_Y,
  }
  
  world.Points = make([]([]*Point), world.SizeX)
  for i:=0;i<world.SizeX;i++ {
    world.Points[i] = make([]*Point, world.SizeY)
    for j:=0;j<world.SizeY;j++ {
      world.Points[i][j] = NewPoint(i,j)
    }
  }
  world.Points[70][30].HasFood = true
  world.Points[90][30].HasFood = true
  world.Points[120][30].HasFood = true
  return world
}

func (w *World) ToString() string {
  r := ""
  for i:=0;i<w.SizeY+2;i++ {
    r += "\r"
  }
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
