package world

import (
  "math"
)

type Point struct {
  X float64
  Y float64
}

func Distance(p1 *Point, p2 *Point) float64{
  return math.Sqrt( math.Pow(p1.X - p2.X, 2.0) + math.Pow(p1.Y - p2.Y, 2.0) )
}