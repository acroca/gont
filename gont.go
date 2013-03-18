package main

import (
  "math/rand"
  "time"
  "fmt"
  "gont/world"
)

func main(){
  rand.Seed( time.Now().UTC().UnixNano())
  ants := make([]*world.Ant, 20)
  for i := 0; i<20; i++ {
    ants[i] = &world.Ant{}
    ants[i].MoveTo(world.WORLD.Points[20][10])
  }

  for i := 0; i<100; i++ {
    fmt.Println(world.WORLD.ToString())
    time.Sleep(50 * time.Millisecond)
    for i := 0; i<20; i++ {
      ants[i].Move()
    }
  }
}

