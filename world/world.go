package world

import (
  "time"
)

type World struct {
  SizeX int
  SizeY int
  Points []([]*Point)
  AntsCount int
}

func InitializeWorld(sizeX int, sizeY int) *World {
  w := &World{
    SizeX: sizeX,
    SizeY: sizeY,
    AntsCount: 0,
  }
  
  w.Points = make([]([]*Point), w.SizeX)
  for i:=0;i<w.SizeX;i++ {
    w.Points[i] = make([]*Point, w.SizeY)
    for j:=0;j<w.SizeY;j++ {
      w.Points[i][j] = NewPoint(w, i,j)
    }
  }
  // w.Points[(sizeX*2)/3][(sizeY*2)/3].HasFood = true
  // w.Points[sizeX/3][sizeY/3].HasFood = true
  // w.Points[sizeX/3][(sizeY*2)/3].HasFood = true
  // w.Points[(sizeX*2)/3][sizeY/3].HasFood = true
  w.Points[(sizeX*5)/6][(sizeY*5)/6].HasFood = true

  hole := w.Points[w.SizeX/2][w.SizeY/2]
  hole.HasHole = true
  
  go w.evaporatePheromones()
  return w
}

func (w *World) PointAt(x int, y int) *Point{
  if x < 0 || y < 0 || x >= w.SizeX || y >= w.SizeY {
    return nil
  }
  return w.Points[x][y]
}

func (w *World) evaporatePheromones() {
  for ; ; {
    time.Sleep(100 * time.Millisecond)
    for _, points := range w.Points {
      for _, point := range points {
        point.RWMutex.Lock() 
        point.Pheromones -= 0.02
        // point.Pheromones *= 0.93
        point.RWMutex.Unlock()
      }
    }
  }
}
