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

const(
  ROUND_PARTS = 1000.0
)

var (
  cachedCos map[float64] float64
  cachedSin map[float64] float64
)

func init() {
  total := int(math.Ceil(math.Pi * 4 * ROUND_PARTS))
  cachedCos = make(map[float64]float64, total)
  cachedSin = make(map[float64]float64, total)

  for i := (0.0) ; i <= (math.Pi * 4) ; i += (1/ROUND_PARTS) {
    v := roundedAngle(i - (2 * math.Pi))
    cachedCos[v] = math.Cos(v)
    cachedSin[v] = math.Sin(v)
  }
}

func RandomVector(x float64, y float64) *Vector {
  return &Vector{
    Point: &Point{
      X: x,
      Y: y,
    },
    Angle: roundedAngle(2 * math.Pi * rand.Float64()),
    Distance: 1.0,
  }
}

func VectorFromPoints(p1, p2 *Point) *Vector{
  v := &Vector{
    Point: &Point{
      X: p1.X,
      Y: p1.Y,
    },
    Angle: math.Atan2(p2.Y-p1.Y, p2.X-p1.X),
    Distance: Distance(p1, p2),
  }
  v.roundAngle()
  return v
}

func (v *Vector) Rotate(r float64) *Vector {
  v.Angle += r
  if v.Angle > (2*math.Pi) { v.Angle -= 2 * math.Pi}
  if v.Angle < (-2 * math.Pi) { v.Angle += 2 * math.Pi}
  v.roundAngle()
  return v
}

func (v *Vector) Move() {
  v.Point.X += cachedCos[v.Angle] * v.Distance
  v.Point.Y += cachedSin[v.Angle] * v.Distance
}

func (v *Vector) TargetPoint() *Point {
  if v.Distance == 0.0 {
    return v.Point
  }
  return &Point{
    X: v.Point.X + (cachedCos[v.Angle] * v.Distance),
    Y: v.Point.Y + (cachedSin[v.Angle] * v.Distance),
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

func (v *Vector) roundAngle(){
  v.Angle = roundedAngle(v.Angle)
}

func roundedAngle(a float64) float64{
  return math.Floor(a * ROUND_PARTS) / ROUND_PARTS
}

