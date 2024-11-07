package color

import (
	"fmt"
	"jkbkaiser/ray-tracing-go/pkg/util/interval"
	"jkbkaiser/ray-tracing-go/pkg/vec3"
	"math"
)

type Color struct {
	R, G, B float64
}

func FromVec(v vec3.Vec3) Color {
	return Color{R: v.X, G: v.Y, B: v.Z}
}

func (c Color) Add(other Color) Color {
	return Color{c.R + other.R, c.G + other.G, c.B + other.B}
}

func (c Color) Scale(value float64) Color {
	return Color{value * c.R, value * c.G, value * c.B}
}

func LinearToGamma(component float64) float64 {
	if component > 0 {
		return math.Sqrt(component)
	}

	return component
}

func (c Color) Write() {
	// Convert colors from range [0, 1] to range [0, 255]
	i := interval.New(.000, .999)

	r := i.Clamp(LinearToGamma(c.R))
	g := i.Clamp(LinearToGamma(c.G))
	b := i.Clamp(LinearToGamma(c.B))

	fmt.Println(
		int(r*255.99),
		int(g*255.99),
		int(b*255.99),
	)
}
