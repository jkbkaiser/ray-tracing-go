package main

import (
	"jkbkaiser/ray-tracing-go/pkg/camera"
	"jkbkaiser/ray-tracing-go/pkg/color"
	"jkbkaiser/ray-tracing-go/pkg/hittable"
	"jkbkaiser/ray-tracing-go/pkg/material"
	"jkbkaiser/ray-tracing-go/pkg/util"
	"jkbkaiser/ray-tracing-go/pkg/vec3"
	"math/rand/v2"
)

func main() {
	world := hittable.HittableList{}

	ground_mat := material.Lambertian{Albedo: color.Color{R: .5, G: .5, B: .5}}
	world.Add(hittable.NewSphere(vec3.Vec3{X: 0., Y: -1000., Z: 0.}, 1000., ground_mat))

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := rand.Float64()
			center := vec3.Vec3{
				X: float64(a) + .9*rand.Float64(),
				Y: .2,
				Z: float64(b) + .9*rand.Float64(),
			}

			if center.Subtract(vec3.Vec3{X: 4., Y: .2, Z: 0.}).Length() > .9 {
				if chooseMat < 0.8 {
					albedo := color.FromVec(vec3.Random()).Mult(color.FromVec(vec3.Random()))
					mat := material.Lambertian{Albedo: albedo}
					world.Add(hittable.NewSphere(center, 0.2, mat))
				} else if chooseMat < 0.95 {
					albedo := color.FromVec(vec3.RandomRange(0.5, 1))
					fuzz := util.RandomFloat(0, 0.5)
					mat := material.Metal{Albedo: albedo, Fuzz: fuzz}
					world.Add(hittable.NewSphere(center, 0.2, mat))
				} else {
					mat := material.Dialectric{RefractionIndex: 1.5}
					world.Add(hittable.NewSphere(center, 0.2, mat))
				}
			}
		}
	}

	mat1 := material.Dialectric{RefractionIndex: 1.5}
	world.Add(hittable.NewSphere(vec3.Vec3{X: 0., Y: 1., Z: 0.}, 1., mat1))

	mat2 := material.Lambertian{Albedo: color.Color{R: .4, G: .2, B: .1}}
	world.Add(hittable.NewSphere(vec3.Vec3{X: -4., Y: 1., Z: 0.}, 1., mat2))

	mat3 := material.Metal{Albedo: color.Color{R: .7, G: .6, B: .5}, Fuzz: 0.}
	world.Add(hittable.NewSphere(vec3.Vec3{X: 4., Y: 1., Z: 0.}, 1., mat3))

	cam := camera.New()
	cam.AspectRatio = 16. / 9.
	cam.ImageWidth = 600
	cam.NumberOfSamples = 50
	cam.MaxDepth = 100

	cam.FieldOfView = 20
	cam.LookFrom = vec3.Vec3{X: 13, Y: 2, Z: 3}
	cam.LookAt = vec3.Vec3{X: 0, Y: 0, Z: 0}
	cam.VUp = vec3.Vec3{X: 0, Y: 1, Z: 0}

	cam.DefocusAngle = .6
	cam.FocusDist = 10.

	cam.Initialize()
	cam.Render(world)
}
