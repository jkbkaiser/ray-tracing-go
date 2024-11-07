package util

import (
	"math"
	"math/rand/v2"
)

const Pi float64 = 3.1415926535897932385

var Inf float64 = math.Inf(1)

func DegreesToRadians(deg float64) float64 {
	return deg * Pi / 180
}

func RandomFloat(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}
