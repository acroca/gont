package world

import (
  "strconv"
  "sync"
)

type Point struct {
  X int
  Y int
  Ants []*Ant
  HasFood bool
  HasHole bool
  FoodPheromones float64
  RWMutex *sync.RWMutex // TODO
}

func NewPoint(x int, y int) *Point{
  return &Point{
    X: x, 
    Y: y, 
    RWMutex: &sync.RWMutex{}, 
    Ants: make([]*Ant, 0),
    HasFood: false,
    HasHole: false,
    FoodPheromones: 0,
  }
}


func (p *Point) AddAnt(ant *Ant) {
  p.Ants = append(p.Ants, ant)  
}

func (p *Point) DeleteAnt(ant *Ant) {
  for idx, v := range p.Ants {
    if v == ant {
      p.Ants = append(p.Ants[:idx], p.Ants[idx+1:]...)
      return
    }
  }
}

func (p *Point) AdjacentPoints() (res []*Point) {
  points := make([]*Point, 8)
  var startX int
  var startY int
  var endX int
  var endY int

  if p.X == 0 { startX = 0 } else { startX = p.X - 1 }
  if p.Y == 0 { startY = 0 } else { startY = p.Y - 1 }
  if p.X == (WORLD.SizeX-1) { endX =(WORLD.SizeX-1) } else { endX = p.X + 1 }
  if p.Y == (WORLD.SizeY-1) { endY =(WORLD.SizeY-1) } else { endY = p.Y + 1 }

  i := 0
  for x := startX; x < endX; x++ {
    for y := startY; y < endY; y++ {
      points[i] = WORLD.Points[x][y]
      i++
    }
  }
  return points
}

func (p *Point) ToString() string {
  length := len(p.Ants)
  var char string

  if length >= 10 {
    char = "#"
  } else if length > 0 {
    char = strconv.Itoa(length)
  } else {
    char = " "
  }

  if p.HasHole {
    if char == " " { char = "O" }
    char = "\x1b[34m" + char + "\x1b[0m"
  } else if p.HasFood {
    if char == " " { char = "*" }
    char = "\x1b[32m" + char + "\x1b[0m"
  } else if p.FoodPheromones > 0 {
    if char == " " { char = "~" }
    char = "\x1b[32m" + char + "\x1b[0m"
  }

  return char
}
