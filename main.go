package main

import (
	"flag"

	"github.com/acroca/gont/sim"
	"github.com/acroca/gont/ui"
)

var (
	numAnts = flag.Int("ants", 10, "Number of ants in the simulator")
)

func main() {
	flag.Parse()
	world := sim.NewWorld(*numAnts)
	window := ui.NewWindow(world)
	go world.Start()
	window.Open()
}
