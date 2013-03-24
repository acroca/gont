package world

import (
  "math"
  "math/rand"
  "time"
)

const (
  BASE_SPEED = 100
)
type Ant struct {
  World *World
  Point *Point
  FacingPoint *Point
  Speed int
  HasFood bool
  HappinessSteps int
  Hunger int
}

func NewAnt(world *World) *Ant {
  world.AntsCount++
  return &Ant{
    World: world,
    Speed: (rand.Int() % 10) + BASE_SPEED,
    HasFood: false,
    HappinessSteps: 0,
    Hunger: 0,
  }
}

func (a *Ant) MoveTo(p *Point) {
  origin_point := a.Point
  if origin_point == p { return }

  if origin_point != nil { origin_point.RWMutex.Lock() }
  p.RWMutex.Lock()

  if origin_point != nil { origin_point.DeleteAnt(a) }

  diffX := 0
  diffY := 0
  if a.Point != nil {
    diffX = p.X - a.Point.X
    diffY = p.Y - a.Point.Y
    a.FacingPoint = a.World.PointAt(p.X + diffX, p.Y + diffY)
  }
  a.Point = p
  p.AddAnt(a)

  if !a.HasFood && p.HasFood {
    a.HasFood = true
    a.FacingPoint = a.World.PointAt(p.X - diffX, p.Y - diffY)
    // a.HappinessSteps = 200
  }
  if a.HasFood && p.HasHole {
    a.HasFood = false
    a.FacingPoint = a.World.PointAt(p.X - diffX, p.Y - diffY)
    a.HappinessSteps = 200
    a.Hunger = 0
  }

  p.RWMutex.Unlock()
  if origin_point != nil { origin_point.RWMutex.Unlock() }

  if p.HasFood { a.HasFood = true }
  a.DropPheromones()
  a.Hunger += 1
}

func (a *Ant) Move() {
  for ; a.Hunger < 1000 ;  {
    time.Sleep(time.Duration(1000/a.Speed) * time.Millisecond)
    p := a.bestCandidatePoint()
    a.MoveTo(p)
  }
  a.Point.DeleteAnt(a)
}

func (a *Ant) candidatePoints() []*Point {
  if a.FacingPoint == nil {
    return a.Point.AdjacentPoints()
  }
  
  candidates := make([]*Point, 3)
  candidates[0] = a.FacingPoint

  diffX := a.FacingPoint.X - a.Point.X
  diffY := a.FacingPoint.Y - a.Point.Y


  if math.Abs(float64(diffX)) == math.Abs(float64(diffY)) { // means corner-to-corner
    candidates[1] = a.World.PointAt(a.FacingPoint.X, a.Point.Y)
    candidates[2] = a.World.PointAt(a.Point.X, a.FacingPoint.Y)
  } else {
    if diffX == 0{
      candidates[1] = a.World.PointAt(a.FacingPoint.X+1, a.FacingPoint.Y)
      candidates[2] = a.World.PointAt(a.FacingPoint.X-1, a.FacingPoint.Y)
    } else {
      candidates[1] = a.World.PointAt(a.FacingPoint.X, a.FacingPoint.Y+1)
      candidates[2] = a.World.PointAt(a.FacingPoint.X, a.FacingPoint.Y-1)
    }
  }

  result := make([]*Point, 3)
  i := 0
  for _, c := range candidates {
    if c != nil {
      result[i] = c
      i++
    }
  }

  if i == 0{
    return a.Point.AdjacentPoints()
  }

  return result[:i]
}

func (a *Ant) bestCandidatePoint() *Point {
  candidates := a.candidatePoints()
  totalPheromones := 0.0
  for _, candidate := range candidates {
    if !a.HasFood && candidate.HasFood { return candidate }
    if a.HasFood && candidate.HasHole { return candidate }
    totalPheromones += candidate.Pheromones
  }
  if totalPheromones > 0{
    rnd := (rand.Float64() * totalPheromones)
    for _, candidate := range candidates {
      rnd -= candidate.Pheromones
      if rnd < 0 {
        return candidate
      }
    }
  }
  return candidates[(rand.Int() % len(candidates))]
}

func (a *Ant) DropPheromones() {
  a.Point.RWMutex.Lock()
  mul := 1.0
  if a.HasFood {
    mul = 20.0
  }
  if a.HappinessSteps > 0 {
    a.HappinessSteps--
    mul = 22.0 
  }
  a.Point.Pheromones += 0.1 * mul
  if a.Point.Pheromones > 1 {
    a.Point.Pheromones = 1.0
  }

  a.Point.RWMutex.Unlock()

  // adjacentPoints := a.Point.AdjacentPoints()
  // for _, p := range adjacentPoints {
  //   if p != nil {
  //     p.RWMutex.Lock()
  //     p.Pheromones += 0.000005 * mul
  //     if p.Pheromones > 1 {
  //       p.Pheromones = 1.0
  //     }
  //     p.RWMutex.Unlock()
  //   }
  // }

}