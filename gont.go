package main

import (
  "math/rand"
  "time"
  "gont/world"
  "gont/ui"
)

func main() {
  game()
  ui := ui.NewUI(world.WORLD)
  ui.Init()
}

func game(){
  rand.Seed( time.Now().UTC().UnixNano())
  ants := make([]*world.Ant, world.ANTS)
  hole := world.WORLD.Points[world.SIZE_X/2][world.SIZE_Y/2]

  hole.HasHole = true

  for i := 0; i<world.ANTS; i++ {
    ants[i] = world.NewAnt()
    ants[i].MoveTo(hole)
    go ants[i].Move()
  }
}