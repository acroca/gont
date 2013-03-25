package world

import (
  "math"
  "math/rand"
)

type Vector struct {
  Point *Point
  Angle float64
}

func RandomVector(x float64, y float64) *Vector {
  return &Vector{
    Point: &Point{
      X: x,
      Y: y, 
    },
    Angle: 2 * math.Pi * rand.Float64(),
  }
}

func (d *Vector) Rotate(r float64) *Vector {
  d.Angle += r
  return d
}

func (d *Vector) Move() {
  d.Point.X += math.Cos(d.Angle) * 0.1
  d.Point.Y += math.Sin(d.Angle) * 0.1
}