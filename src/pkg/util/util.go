package util

import "math"

const Pi float64 = 3.1415926535897932385

var Inf float64 = math.Inf(1)

func DegreesToRadians(deg float64) float64 {
	return deg * Pi / 180
}
