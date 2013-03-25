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
  w := world.InitializeWorld(SIZE_X, SIZE_Y, ANTS)
  
  hole := world.NewHole(&world.Point{X: float64(w.SizeX/2), Y: float64(w.SizeY/2)})
  w.Holes.PushBack(hole)
  
  food := world.NewFood(&world.Point{X: float64(w.SizeX*4/5), Y: float64(w.SizeY*4/5)})
  w.Food.PushBack(food)

  go world.GenerateAnts(w)
  go world.EvaporateAll(w.Pheromones)


  // w.Points = make([]([]*Point), w.SizeX)
  // for i:=0;i<w.SizeX;i++ {
  //   w.Points[i] = make([]*Point, w.SizeY)
  //   for j:=0;j<w.SizeY;j++ {
  //     w.Points[i][j] = NewPoint(w, i,j)
  //   }
  // }
  // // w.Points[(sizeX*2)/3][(sizeY*2)/3].HasFood = true
  // // w.Points[sizeX/3][sizeY/3].HasFood = true
  // // w.Points[sizeX/3][(sizeY*2)/3].HasFood = true
  // // w.Points[(sizeX*2)/3][sizeY/3].HasFood = true
  // w.Points[(sizeX*5)/6][(sizeY*5)/6].HasFood = true

  // hole := w.Points[w.SizeX/2][w.SizeY/2]
  // hole.HasHole = true
  
  // go w.evaporatePheromones()

  // hole := w.Points[w.SizeX/2][w.SizeY/2]
  // hole.HasHole = true
  
  // go addAnts(w, hole, ANTS)
  
  ui := ui.NewUI(w)
  ui.Init()
}

// func addAnts(w *world.World, hole *world.Point, max int) {
//   for ; ; {
//     if w.AntsCount < max {
//       ant := world.NewAnt(w)
//       ant.MoveTo(hole)
//       go ant.Move()
//     }
//     time.Sleep(100 * time.Millisecond)
//   }
// }