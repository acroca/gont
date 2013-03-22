package world

import (
  "math"
  "math/rand"
  "time"
  "fmt"
)

const (
  BASE_SPEED = 10
)
type Ant struct {
  World *World
  Point *Point
  LastPoint *Point
  Speed int
  HasFood bool
}

func NewAnt(world *World) *Ant {
  return &Ant{
    World: world,
    Speed: (rand.Int() % 10) + BASE_SPEED,
    HasFood: false,
  }
}

func (a *Ant) MoveTo(p *Point) {
  origin_point := a.Point
  if origin_point == p { return }

  if origin_point != nil { origin_point.RWMutex.Lock() }
  p.RWMutex.Lock()

  if origin_point != nil { origin_point.DeleteAnt(a) }
  a.LastPoint = a.Point
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
    p := a.bestPoint()
    if p == nil {
      a.MoveRand()
    }else{
      a.MoveTo(p)
    }
  }
}

func (a *Ant) bestPoint() *Point {
  return nil
  adjacent := a.Point.AdjacentPoints()
  var bestPoint *Point

  for _, p := range adjacent {
    if p != nil {
      if p.HasFood { return p }
      if p.Pheromones > 0 && (bestPoint == nil || p.Pheromones > bestPoint.Pheromones) { bestPoint = p } 
    }
  }
  return bestPoint
}

func (a *Ant) MoveRand() {
  if a.LastPoint == nil { 
    a.MoveAnywhere()
  } else {
    diffX := a.Point.X - a.LastPoint.X
    diffY := a.Point.Y - a.LastPoint.Y
    changeX := a.Point.X + diffX
    changeY := a.Point.Y + diffY
    fmt.Printf("Diff: %v,%v\n", diffX, diffY)
    if math.Abs(float64(diffX)) == math.Abs(float64(diffY)) { // means corner-to-corner
      rnd := rand.Int() % 3
      switch{
      case rnd == 0:
        // Do nothing
      case rnd == 1:
        changeX = a.Point.X
      case rnd == 2:
        changeY = a.Point.Y
      }
    } else { // Vertical or horizontal
      if diffX == 0{
        changeX += (rand.Int() % 3) - 1
      } else {
        changeY += (rand.Int() % 3) - 1
      }
    }
    if changeY < 0 || changeX < 0 || changeY >= a.World.SizeY || changeX >= a.World.SizeX { 
      a.MoveAnywhere()
    }else{
      if changeX == a.Point.X && changeY == a.Point.Y {
        a.MoveRand()
      } else { 
        a.MoveTo(a.World.Points[changeX][changeY]) 
      }      
    }
  }
}
func (a *Ant) MoveAnywhere() {
  changeX := a.Point.X + (rand.Int() % 3) - 1
  changeY := a.Point.Y + (rand.Int() % 3) - 1
  if changeY < 0 { changeY = 0 }
  if changeX < 0 { changeX = 0 }
  if changeY >= a.World.SizeY { changeY = a.World.SizeY - 1 }
  if changeX >= a.World.SizeX { changeX = a.World.SizeX - 1 }
  a.MoveTo(a.World.Points[changeX][changeY])
}

func (a *Ant) DropPheromones() {
  adjacentPoints := a.Point.AdjacentPoints()

  a.Point.RWMutex.Lock()
  a.Point.Pheromones += 0.02
  a.Point.RWMutex.Unlock()

  for _, p := range adjacentPoints {
    if p != nil {
      p.RWMutex.Lock()
      p.Pheromones += 0.01
      p.RWMutex.Unlock()
    }
  }
}