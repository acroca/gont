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
}

func NewAnt(w *World, p *Point) *Ant {
  return &Ant{
    World: w,
    Vector: RandomVector(p.X, p.Y),
    Speed: (rand.Int() % VARIABLE_SPEED) + BASE_SPEED,
    HasFood: false,
  }
}

func GenerateAnts(w *World){
  for ; ; {
    time.Sleep(400 * time.Millisecond)
    if w.Ants.Len() < w.MaxAnts {
      hole := w.Holes.Front().Value.(*Hole)
      ant := NewAnt(w, hole.Point)
      w.Ants.PushBack(ant)
      go ant.Move()
    }
  }
}

func (a *Ant) Move() {
  for ; ; {
    for i:=0; i<20; i++ {
      time.Sleep(time.Duration(1000/a.Speed) * time.Millisecond)
      a.Vector.Move()
    }

    a.Reorientate()
    a.DropPheromone()
  }
}

func (a *Ant) Reorientate() {
  var pheromonesVector *Vector

  // For each pheromone ph within VISIBLE_RANGE
  //   Take the vector to the pheromone
  //   Scale the vector based on the angle from the ant's vector (Normal distribution?)
  //   Scale the vector based on the intensity of the pheromone
  // Calculate the sum of all the vectors, including the ant's vector
  // Ant points to that direction

  for e:= a.World.Pheromones.Front(); e != nil ; e = e.Next() {
    pheromone := e.Value.(*Pheromone)
    distance := Distance(pheromone.Point, a.Vector.Point)
    if distance <= VISIBLE_RANGE {
      pheromoneVector := VectorFromPoints(a.Vector.Point, pheromone.Point)

      angle := math.Abs(a.Vector.Angle - pheromoneVector.Angle)
      if angle < math.Pi / 2 {
        pheromoneVector.Distance *= 1 - ((1/math.Pi) * angle)

        pheromoneVector.Distance *= float64(pheromone.Amount) / 100.0

        random := (rand.NormFloat64() * 0.1) + 1
        if random < 0 {
          random = 0
        } else if random > 1 {
          random = 1
        }
        pheromoneVector.Distance *= random

        if pheromonesVector == nil {
          pheromonesVector = pheromoneVector
        } else {
          pheromoneVector = pheromoneVector.Sum(pheromoneVector)
        }
      }
    }
  }
  if pheromonesVector != nil {
    a.Vector = a.Vector.Sum(pheromonesVector)
  } else {
    a.Vector.Rotate((rand.Float64() * (math.Pi/4)) - (math.Pi/8))
  }
  a.Vector.Distance = 1.0
}

func (a *Ant) DropPheromone() {
  pheromone := NewPheromone( &Point {X: a.Vector.Point.X, Y: a.Vector.Point.Y, })
  a.World.Pheromones.PushBack(pheromone)
}