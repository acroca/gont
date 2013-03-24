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
  ANTS = 500
)
func main() {
  rand.Seed( time.Now().UTC().UnixNano())
  w := world.InitializeWorld(SIZE_X, SIZE_Y)

  hole := w.Points[w.SizeX/2][w.SizeY/2]
  hole.HasHole = true
  
  go addAnts(w, hole, ANTS)
  
  ui := ui.NewUI(w)
  ui.Init()
}

func addAnts(w *world.World, hole *world.Point, max int) {
  for ; ; {
    if w.AntsCount < max {
      ant := world.NewAnt(w)
      ant.MoveTo(hole)
      go ant.Move()
    }
    time.Sleep(100 * time.Millisecond)
  }
}