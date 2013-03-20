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
  RWMutex *sync.RWMutex // TODO
}

func NewPoint(x int, y int) *Point{
  return &Point{
    X: x, 
    Y: y, 
    RWMutex: &sync.RWMutex{}, 
    Ants: make([]*Ant, 0),
    HasFood: false,
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

  if p.HasFood {
    if char == " " { char = "*" }
    char = "\x1b[32m" + char + "\x1b[0m"
  }

  return char
}
