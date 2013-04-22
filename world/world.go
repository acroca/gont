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

type PheromoneDistance struct {
  Pheromone *Pheromone
  Distance float64
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
  pheromones := make(chan *PheromoneDistance, 10)
  var best *Pheromone
  var bestDist float64
  var pd *PheromoneDistance
  go w.PheromonesWithin(p, r, pheromones)

  for pd = <-pheromones ; pd != nil; pd = <-pheromones {
    if best == nil || bestDist > pd.Distance {
      best = pd.Pheromone
      bestDist = pd.Distance
    }
  }

  return best
}

func (w *World) PheromonesWithin(p *Point, r float64, pheromones chan *PheromoneDistance) {
  for e:= w.Pheromones.Front(); e != nil ; e = e.Next() {
    pheromone := e.Value.(*Pheromone)
    distance, err := DistanceMax(pheromone.Point, p, r)
    if !err && distance < r {
      pheromones <- &PheromoneDistance{Pheromone: pheromone, Distance: distance,}
    }
  }
  pheromones <- nil
}
