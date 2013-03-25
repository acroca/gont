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
