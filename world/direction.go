package world

import (
  "math"
  "math/rand"
)

type Direction struct {
  X float64
  Y float64
}

func RandomDirection() *Direction {
  randX := (rand.Float64() * 2) - 1
  randY := (rand.Float64() * 2) - 1
  direction := &Direction{ X: randX, Y: randY, }
  return direction.Normalize()
}

func (d *Direction) Rotate(r float64) *Direction {
  d.X = (d.X * math.Cos(r)) - (d.Y * math.Sin(r))
  d.Y = (d.X * math.Sin(r)) + (d.Y * math.Cos(r))
  return d.Normalize()
}

func (d *Direction) Normalize() *Direction {
  length := math.Sqrt((d.X * d.X) + (d.Y * d.Y))
  d.X /= length
  d.Y /= length
  return d
}
