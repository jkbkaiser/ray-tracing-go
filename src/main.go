package main

import (
	"jkbkaiser/ray-tracing-go/pkg/camera"
	"jkbkaiser/ray-tracing-go/pkg/color"
	"jkbkaiser/ray-tracing-go/pkg/hittable"
	"jkbkaiser/ray-tracing-go/pkg/material"
	"jkbkaiser/ray-tracing-go/pkg/vec3"
)

func main() {
	mat_ground := material.Lambertian{Albedo: color.Color{R: .8, G: .8, B: .0}}
	mat_center := material.Lambertian{Albedo: color.Color{R: .1, G: .2, B: .5}}
	mat_left := material.Metal{Albedo: color.Color{R: .8, G: .8, B: .8}, Fuzz: 0.3}
	mat_right := material.Metal{Albedo: color.Color{R: .8, G: .6, B: .2}, Fuzz: 1.0}

	world := hittable.HittableList{}
	world.Add(hittable.NewSphere(vec3.Vec3{X: 0., Y: -100.5, Z: -1.}, 100., mat_ground))
	world.Add(hittable.NewSphere(vec3.Vec3{X: 0., Y: 0., Z: -1.2}, .5, mat_center))
	world.Add(hittable.NewSphere(vec3.Vec3{X: -1., Y: 0., Z: -1.}, .5, mat_left))
	world.Add(hittable.NewSphere(vec3.Vec3{X: 1., Y: 0., Z: -1.}, .5, mat_right))

	cam := camera.New()
	cam.AspectRatio = 16. / 9.
	cam.ImageWidth = 400
	cam.NumberOfSamples = 100
	cam.MaxDepth = 50

	cam.Initialize()
	cam.Render(world)
}
