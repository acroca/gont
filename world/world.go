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
  SIZE_X = 128
  SIZE_Y = 96
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
  world.Points[(SIZE_X*2)/3][(SIZE_Y*2)/3].HasFood = true
  world.Points[SIZE_X/3][SIZE_Y/3].HasFood = true
  world.Points[SIZE_X/3][(SIZE_Y*2)/3].HasFood = true
  world.Points[(SIZE_X*2)/3][SIZE_Y/3].HasFood = true
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
