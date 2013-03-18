package world

import (
  "math/rand"
  "time"
)

const (
  BASE_SPEED = 30
)
type Ant struct {
  Point *Point
  Speed int
}

func NewAnt() *Ant {
  return &Ant{Speed: (rand.Int() % 20) + BASE_SPEED}
}

func (a *Ant) ToString() string {
  return "Hi!"
}

func (a *Ant) MoveTo(p *Point) {
  if a.Point == p { return }

  if a.Point != nil {
    a.Point.RWMutex.Lock()
    defer a.Point.RWMutex.Unlock()
  }
  p.RWMutex.Lock()
  defer p.RWMutex.Unlock()

  if a.Point != nil {
    a.Point.DeleteAnt(a)
  }
  a.Point = p
  p.AddAnt(a)
}

func (a *Ant) Move() {
  for ; ; {
    time.Sleep(time.Duration(1000/a.Speed) * time.Millisecond)
    a.MoveRand()
  }
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