package main

import (
  "math/rand"
  "time"
  "gont/world"
  "gont/ui"
)

const (
  SIZE_X = 128
  SIZE_Y = 96
  ANTS = 100
)
func main() {
  rand.Seed( time.Now().UTC().UnixNano())
  w := world.InitializeWorld(128,96)

  hole := w.Points[w.SizeX/2][w.SizeY/2]
  hole.HasHole = true
  
  ants := make([]*world.Ant, ANTS)
  for i := 0; i<ANTS; i++ {
    ants[i] = world.NewAnt(w)
    ants[i].MoveTo(hole)
    go ants[i].Move()
  }

  ui := ui.NewUI(w)
  ui.Init()
}
