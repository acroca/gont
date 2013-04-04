package world

import (
  "math"
  "math/rand"
)

type Vector struct {
  Point *Point
  Angle float64
  Distance float64
}

func RandomVector(x float64, y float64) *Vector {
  return &Vector{
    Point: &Point{
      X: x,
      Y: y,
    },
    Angle: 2 * math.Pi * rand.Float64(),
    Distance: 1.0,
  }
}

func VectorFromPoints(p1, p2 *Point) *Vector{
  return &Vector{
    Point: &Point{
      X: p1.X,
      Y: p1.Y,
    },
    Angle: math.Atan2(p2.Y-p1.Y, p2.X-p1.X),
    Distance: Distance(p1, p2),
  }
}

func (v *Vector) Rotate(r float64) *Vector {
  v.Angle += r
  return v
}

func (v *Vector) Move() {
  v.Point.X += math.Cos(v.Angle) * v.Distance
  v.Point.Y += math.Sin(v.Angle) * v.Distance
}

func (v *Vector) TargetPoint() *Point {
  if v.Distance == 0.0 {
    return v.Point
  }
  return &Point{
    X: v.Point.X + (math.Cos(v.Angle) * v.Distance),
    Y: v.Point.Y + (math.Sin(v.Angle) * v.Distance),
  }
}

func (v1 *Vector) SumFromPoint(p *Point) *Vector {
  v2 := VectorFromPoints(v1.Point, p)
  return v1.Sum(v2)
}

func (v1 Vector) Sum(v2 *Vector) *Vector {
  initialPoint := v1.Point
  v1.Point = v1.TargetPoint()
  v1.Angle = v2.Angle
  finalPoint := v1.TargetPoint()
  return VectorFromPoints(initialPoint, finalPoint)
}
