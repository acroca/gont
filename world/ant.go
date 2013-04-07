package world

import (
  "math"
  "math/rand"
  "time"
)

const (
  BASE_SPEED = 300
  VARIABLE_SPEED = 10
  VISIBLE_RANGE = 70
  ARMS_RANGE = 10
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
    time.Sleep(50 * time.Millisecond)
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

  if !a.HasFood {
    for e:= a.World.Food.Front(); e != nil ; e = e.Next() {
      food := e.Value.(*Food)
      distance, err := DistanceMax(food.Point, a.Vector.Point, VISIBLE_RANGE)
      if !err {
        if distance <= ARMS_RANGE {
          a.HasFood = true
          a.Vector.Rotate(math.Pi)
          return
        } else if distance <= VISIBLE_RANGE {
          a.Vector = VectorFromPoints(a.Vector.Point, food.Point)
          a.Vector.Distance = 1.0
          return
        }
      }
    }
  } else {
    for e:= a.World.Holes.Front(); e != nil ; e = e.Next() {
      hole := e.Value.(*Hole)
      distance := Distance(hole.Point, a.Vector.Point)
      if distance <= ARMS_RANGE {
        a.HasFood = false
        a.Vector.Rotate(math.Pi)
        return
      } else if distance <= VISIBLE_RANGE {
        a.Vector = VectorFromPoints(a.Vector.Point, hole.Point)
        a.Vector.Distance = 1.0
        return
      }
    }
  }

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
    if a.HasFood {
      pheromonesVector.Distance *= 100
    }
    a.Vector = a.Vector.Sum(pheromonesVector)
  } else {
    a.Vector.Rotate((rand.Float64() * (math.Pi/4)) - (math.Pi/8))
  }
  next := a.Vector.TargetPoint()
  if next.X < 0 || next.Y < 0 || next.X > float64(a.World.SizeX) || next.Y > float64(a.World.SizeY) {
    a.Vector.Rotate(math.Pi)
  }
  a.Vector.Distance = 1.0
}

func (a *Ant) DropPheromone() {
  pheromone := NewPheromone( &Point {X: a.Vector.Point.X, Y: a.Vector.Point.Y, })
  a.World.Pheromones.PushBack(pheromone)
}