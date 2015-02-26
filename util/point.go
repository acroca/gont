package util

import (
	"math/rand"

	"github.com/go-gl/mathgl/mgl32"
)

// RandomPoint returns a random point
func RandomPoint() mgl32.Vec2 {
	return mgl32.Vec2{rand.Float32(), rand.Float32()}
}
