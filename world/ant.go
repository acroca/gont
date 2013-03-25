package world

import (
  "math"
  "math/rand"
  "time"
)

const (
  BASE_SPEED = 30
  VARIABLE_SPEED = 10
)

type Ant struct {
  World *World
  Direction *Direction
  Point *Point
  FacingPoint *Point
  Speed int
  HasFood bool
  Hunger int
}

func NewAnt(w *World, p *Point) *Ant {
  return &Ant{
    World: w,
    Direction: RandomDirection(),
    Point: p,
    Speed: (rand.Int() % VARIABLE_SPEED) + BASE_SPEED,
    HasFood: false,
    Hunger: 0,
  }
}

func GenerateAnts(w *World){
  for ; ; {
    time.Sleep(100 * time.Millisecond)
    if w.Ants.Len() < w.MaxAnts {
      hole := w.Holes.Front().Value.(*Hole)
      ant := NewAnt(w, &Point{X: hole.Point.X, Y: hole.Point.Y})
      w.Ants.PushBack(ant)
      go ant.Move()
    }
  }
}

func (a *Ant) Move() {
  for ; a.Hunger < 1000 ;  {
    time.Sleep(time.Duration(1000/a.Speed) * time.Millisecond)
  
    a.Point.Move(a.Direction)
    a.Direction.Rotate((rand.Float64() * (math.Pi/2)) - (math.Pi/4))

    a.DropPheromone()
  }
  for e := a.World.Ants.Front() ; e != nil ; e = e.Next() {
    if e.Value.(*Ant) == a {
      a.World.Ants.Remove(e)
      return
    }
  }
}

func (a *Ant) DropPheromone() {
  pheromone := NewPheromone( &Point {X: a.Point.X, Y: a.Point.Y, })
  a.World.Pheromones.PushBack(pheromone)
}