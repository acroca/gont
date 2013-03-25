package world

import(
  "time"
  "container/list"
)

const (
  MAX_AMOUNT = 100

  EVAPORATION_TIME = 2 * time.Second
  EVAPORATION_DELAY = 100 * time.Millisecond
  EVAPORATION_AMOUNT = int((MAX_AMOUNT * EVAPORATION_TIME) / EVAPORATION_DELAY)
)

type Pheromone struct {
  Point *Point
  Amount int
}

func NewPheromone(p *Point) *Pheromone {
  return &Pheromone{
    Point: p,
    Amount: MAX_AMOUNT,
  }
}

func EvaporateAll(pheromones *list.List) {
  for ; ; {
    time.Sleep(EVAPORATION_DELAY)
    for e := pheromones.Front(); e != nil; e = e.Next() {
      pheromone := e.Value.(*Pheromone)
      pheromone.Amount -= EVAPORATION_AMOUNT
      if pheromone.Amount < 0 { pheromones.Remove(e) }
    }
  }
}