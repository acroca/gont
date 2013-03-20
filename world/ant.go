package world

import (
  "math/rand"
  "time"
)

const (
  BASE_SPEED = 40
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
  if a.HasFood { a.DropFoodPheromones() }
}

func (a *Ant) Move() {
  for ; ; {
    time.Sleep(time.Duration(1000/a.Speed) * time.Millisecond)
    a.MoveRand()
  }
}

func (a *Ant) MoveHole() {

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

func (a *Ant) DropFoodPheromones() {
  adjacentPoints := a.Point.AdjacentPoints()
  for _, p := range adjacentPoints {
    if p != nil {
      p.RWMutex.Lock()
      if p.FoodPheromones == 0 {
        p.FoodPheromones = 0.1
      } else {
        p.FoodPheromones *= 1.1
        if p.FoodPheromones > 1 { p.FoodPheromones = 1 }
      }
      p.RWMutex.Unlock()
    }
  }
}