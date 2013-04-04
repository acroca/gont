package main

import (
  // "math"
  "math/rand"
  "time"
  "gont/world"
  "gont/ui"
)

const (
  SIZE_X = 1280
  SIZE_Y = 960
  ANTS = 20
)
func main() {
  rand.Seed( time.Now().UTC().UnixNano())
  w := world.InitializeWorld(SIZE_X, SIZE_Y, ANTS)

  hole := world.NewHole(&world.Point{X: float64(w.SizeX/2), Y: float64(w.SizeY/2)})
  w.Holes.PushBack(hole)

  food := world.NewFood(&world.Point{X: float64(w.SizeX*4/5), Y: float64(w.SizeY*4/5)})
  w.Food.PushBack(food)

  go world.GenerateAnts(w)
  go world.EvaporateAll(w.Pheromones)

  ui := ui.NewUI(w)
  ui.Init()
}