package color

import (
	"fmt"
	"jkbkaiser/ray-tracing-go/pkg/vec3"
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

func (c *Color) Write() {
	// Convert colors from range [0, 1] to range [0, 255]
	fmt.Println(int(c.R*255.99), int(c.G*255.99), int(c.B*255.99))
}
