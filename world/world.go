package world

import (
  "container/list"
)

type World struct {
  SizeX int
  SizeY int

  Ants *list.List
  Pheromones *list.List
  Holes *list.List
  Food *list.List

  MaxAnts int
}

func InitializeWorld(sizeX int, sizeY int, maxAnts int) *World {
  return &World{
    SizeX: sizeX,
    SizeY: sizeY,
    Ants: list.New(),
    Pheromones: list.New(),
    Holes: list.New(),
    Food: list.New(),
    MaxAnts: maxAnts,
  }
}

func (w *World) ClosestPheromoneWithin(p *Point, r float64) *Pheromone {
  var best *Pheromone
  var bestDist float64

  for e:= w.Pheromones.Front(); e != nil ; e = e.Next() {
    pheromone := e.Value.(*Pheromone)
    distance, err := DistanceMax(pheromone.Point, p, r)
    if !err && distance < r && (best == nil || bestDist > distance ) {
      best = pheromone
      bestDist = distance
    }
  }
  return best
}
