package sim

import "time"

const (
	antMovementPerSecond    = 0.05
	antMaxRotationPerSecond = 3
	antPheromoneFrequency   = 1000 * time.Millisecond
	pheromoneDuration       = 25 * time.Second
)
