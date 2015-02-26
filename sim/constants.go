package sim

import "time"

const (
	simSpeed = 4

	antMovementPerSecond    = 0.05 * simSpeed
	antMaxRotationPerSecond = 3 * simSpeed
	antPheromoneFrequency   = (500 / simSpeed) * time.Millisecond
	pheromoneDuration       = (25000 / simSpeed) * time.Millisecond

	pheromoneIndexParts = 100
)
