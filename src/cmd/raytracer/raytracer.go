package main

import (
	"jkbkaiser/ray-tracing-go/pkg/camera"
	"jkbkaiser/ray-tracing-go/pkg/hittable"
	"jkbkaiser/ray-tracing-go/pkg/vec3"
)

func main() {
	var world = hittable.HittableList{}.
		Add(hittable.Sphere{Center: vec3.Vec3{X: 0., Y: 0., Z: -1}, Radius: .5}).
		Add(hittable.Sphere{Center: vec3.Vec3{X: 0., Y: -100.5, Z: -1}, Radius: 100})

	cam := camera.New()
	cam.AspectRatio = 16. / 9.
	cam.ImageWidth = 400

	cam.Initialize()
	cam.Render(world)
}
