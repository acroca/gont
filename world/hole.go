package world

type Hole struct {
  Point *Point
}

func NewHole(p *Point) *Hole {
  return &Hole{Point: p,}
}