package world

import (
  "math/rand"
  "time"
)

const (
  BASE_SPEED = 10
)
type Ant struct {
  Point *Point
  Speed int
  HasFood bool
}

func NewAnt() *Ant {
  return &Ant{
    Speed: (rand.Int() % 20) + BASE_SPEED,
    HasFood: false,
  }
}

func (a *Ant) ToString() string {
  return "Hi!"
}

func (a *Ant) MoveTo(p *Point) {
  origin_point := a.Point
  if origin_point == p { return }

  if origin_point != nil { origin_point.RWMutex.Lock() }
  p.RWMutex.Lock()

  if origin_point != nil { origin_point.DeleteAnt(a) }
  a.Point = p
  p.AddAnt(a)

  p.RWMutex.Unlock()
  if origin_point != nil { origin_point.RWMutex.Unlock() }

  if p.HasFood { a.HasFood = true }
  a.DropPheromones()
}

func (a *Ant) Move() {
  for ; ; {
    time.Sleep(time.Duration(1000/a.Speed) * time.Millisecond)
    if a.HasFood {
      a.MoveHole()
    } else {
      a.MoveFood()
    }
  }
}

func (a *Ant) MoveHole() {
  adjacent := a.Point.AdjacentPoints()

  bestPoint := adjacent[0]

  for _, p := range adjacent {
    if p != nil && p.PresencePheromones > bestPoint.PresencePheromones {
     bestPoint = p 
    }
  }

  if bestPoint.FoodPheromones > 0 {
    a.MoveTo(bestPoint)
  } else {
    a.MoveRand()
  }
}

func (a *Ant) MoveFood() {
  bestPoint := a.bestFoodPoint()

  if bestPoint.HasFood || bestPoint.FoodPheromones > 0 {
    a.MoveTo(bestPoint)
  } else {
    a.MoveRand()
  }
}

func (a *Ant) bestFoodPoint() *Point {
  adjacent := a.Point.AdjacentPoints()
  bestPoint := adjacent[0]

  for _, p := range adjacent {
    if p != nil {
      if p.HasFood { return p }
      if p.FoodPheromones > bestPoint.FoodPheromones { bestPoint = p }
    }
  }
  return bestPoint
}

func (a *Ant) MoveRand() {
  changeX := a.Point.X + (rand.Int() % 3) - 1
  changeY := a.Point.Y + (rand.Int() % 3) - 1
  if changeY < 0 { changeY = 0 }
  if changeX < 0 { changeX = 0 }
  if changeY >= WORLD.SizeY { changeY = WORLD.SizeY - 1 }
  if changeX >= WORLD.SizeX { changeX = WORLD.SizeX - 1 }
  a.MoveTo(WORLD.Points[changeX][changeY])
}

func (a *Ant) DropPheromones() {
  adjacentPoints := a.Point.AdjacentPoints()

  a.Point.RWMutex.Lock()

  if a.HasFood {
    if a.Point.FoodPheromones == 0 {
      a.Point.FoodPheromones = 0.1
    } else {
      a.Point.FoodPheromones *= 1.2
      if a.Point.FoodPheromones > 1 { a.Point.FoodPheromones = 1 }
    }
    for _, p := range adjacentPoints {
      if p != nil {
        if p.FoodPheromones == 0 {
          p.FoodPheromones = 0.1
        } else {
          p.FoodPheromones *= 1.1
          if p.FoodPheromones > 1 { p.FoodPheromones = 1 }
        }
      }
    }
  } else {
    if a.Point.PresencePheromones == 0 {
      a.Point.PresencePheromones = 0.1
    } else {
      a.Point.PresencePheromones *= 1.1
      if a.Point.PresencePheromones > 1 { a.Point.PresencePheromones = 1 }
    }
    for _, p := range adjacentPoints {
      if p != nil {
        if a.Point.PresencePheromones == 0 {
          a.Point.PresencePheromones = 0.1
        } else {
          a.Point.PresencePheromones *= 1.05
          if a.Point.PresencePheromones > 1 { a.Point.PresencePheromones = 1 }
        }
      }
    }
  }

  a.Point.RWMutex.Unlock()
}