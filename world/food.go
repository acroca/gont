package world

type Food struct {
  Point *Point
}

func NewFood(p *Point) *Food {
  return &Food{Point: p,}
}