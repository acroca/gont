package world

type Point struct {
  X float64
  Y float64
}

func (p *Point) Move(d *Direction) {
  // p.X += d.X
  // p.Y += d.Y
  p.X += d.X * 0.1
  p.Y += d.Y * 0.1
}