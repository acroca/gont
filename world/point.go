package world

import (
  "strconv"
)

type Point struct {
  X int
  Y int
  Ants []*Ant
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
  if length > 0 {
    return strconv.Itoa(length)
  }
  return " "
}

