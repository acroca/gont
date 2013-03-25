package world

import (
  "math"
  "math/rand"
  "time"
)

const (
  BASE_SPEED = 70
  VARIABLE_SPEED = 10
  VISIBLE_RANGE = 20
)

type Ant struct {
  World *World
  Vector *Vector
  Speed int
  HasFood bool
  Hunger int
}

func NewAnt(w *World, p *Point) *Ant {
  return &Ant{
    World: w,
    Vector: RandomVector(p.X, p.Y),
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
      ant := NewAnt(w, hole.Point)
      w.Ants.PushBack(ant)
      go ant.Move()
    }
  }
}

func (a *Ant) Move() {
  for ; a.Hunger < 1000 ;  {
    for i:=0; i<20; i++ {
      time.Sleep(time.Duration(1000/a.Speed) * time.Millisecond)
      a.Vector.Move()
    }
    
    a.Reorientate()
    a.DropPheromone()
  }
  for e := a.World.Ants.Front() ; e != nil ; e = e.Next() {
    if e.Value.(*Ant) == a {
      a.World.Ants.Remove(e)
      return
    }
  }
}

func (a *Ant) Reorientate() {
  a.Vector.Rotate((rand.Float64() * (math.Pi/2)) - (math.Pi/4))
}

func (a *Ant) DropPheromone() {
  pheromone := NewPheromone( &Point {X: a.Vector.Point.X, Y: a.Vector.Point.Y, })
  a.World.Pheromones.PushBack(pheromone)
}