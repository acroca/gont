package world

import(
  "time"
  "container/list"
)

const (
  MAX_AMOUNT = 100
  UNIT_AMOUNT = 20

  EVAPORATION_TIME = 10000 * time.Millisecond
  EVAPORATION_DELAY = 200 * time.Millisecond
  EVAPORATION_AMOUNT = int((MAX_AMOUNT * EVAPORATION_DELAY) / EVAPORATION_TIME)
)

type Pheromone struct {
  Point *Point
  Amount int
}

func NewPheromone(p *Point) *Pheromone {
  return &Pheromone{
    Point: p,
    Amount: UNIT_AMOUNT,
  }
}

func EvaporateAll(pheromones *list.List) {
  for ; ; {
    time.Sleep(EVAPORATION_DELAY)
    var next *list.Element
    for e := pheromones.Front(); e != nil; e = next {
      pheromone := e.Value.(*Pheromone)
      pheromone.Amount -= EVAPORATION_AMOUNT
      next = e.Next()
      if pheromone.Amount < 0 { pheromones.Remove(e) }
    }
  }
}
