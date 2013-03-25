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
  ANTS = 200
)
func main() {
  rand.Seed( time.Now().UTC().UnixNano())
  w := world.InitializeWorld(SIZE_X, SIZE_Y, ANTS)
  hole := world.NewHole(&world.Point{X: float64(w.SizeX/2), Y: float64(w.SizeY/2)})
  w.Holes.PushBack(hole)
  go world.GenerateAnts(w)



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