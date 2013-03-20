package main

import (
  "math/rand"
  "time"
  "fmt"
  "gont/world"
)

func main(){
  rand.Seed( time.Now().UTC().UnixNano())
  ants := make([]*world.Ant, world.ANTS)
  hole := world.WORLD.Points[world.SIZE_X/2][world.SIZE_Y/2]

  hole.HasHole = true

  for i := 0; i<world.ANTS; i++ {
    ants[i] = world.NewAnt()
    ants[i].MoveTo(hole)
    fmt.Println("A")
    go ants[i].Move()
  }

  for i := 0; i<100; i++ {
    fmt.Println(world.WORLD.ToString())
    time.Sleep(50 * time.Millisecond)
  }
}