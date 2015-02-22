package main

import (
	"github.com/acroca/gont/sim"
	"github.com/acroca/gont/ui"
)

func main() {
	world := sim.NewWorld(1000)
	ui.NewWindow(world).Open()
}
