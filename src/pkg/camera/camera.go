package camera

import (
	"fmt"
	"math"
	"math/rand/v2"
	"os"

	"jkbkaiser/ray-tracing-go/pkg/color"
	"jkbkaiser/ray-tracing-go/pkg/hittable"
	"jkbkaiser/ray-tracing-go/pkg/progress_bar"
	"jkbkaiser/ray-tracing-go/pkg/ray"
	"jkbkaiser/ray-tracing-go/pkg/util"
	"jkbkaiser/ray-tracing-go/pkg/util/interval"
	"jkbkaiser/ray-tracing-go/pkg/vec3"
)

type Camera struct {
	maxColor int

	AspectRatio float64

	ImageWidth  int
	imageHeight int

	focalLength  float64
	cameraCenter vec3.Vec3
	LookFrom     vec3.Vec3
	LookAt       vec3.Vec3
	VUp          vec3.Vec3
	u            vec3.Vec3
	v            vec3.Vec3
	w            vec3.Vec3

	viewportHeight float64
	viewportWidth  float64

	viewportUpperLeft vec3.Vec3
	pixel00Loc        vec3.Vec3
	deltaU            vec3.Vec3
	deltaV            vec3.Vec3

	NumberOfSamples  int
	pixelSampleScale float64

	MaxDepth int

	FieldOfView float64
}

func New() Camera {
	return Camera{
		AspectRatio:     16. / 9.,
		ImageWidth:      400,
		NumberOfSamples: 100,
		MaxDepth:        50,
		FieldOfView:     90,
		LookFrom:        vec3.Vec3{X: 0, Y: 0, Z: 0},
		LookAt:          vec3.Vec3{X: 0, Y: 0, Z: -1},
		VUp:             vec3.Vec3{X: 0, Y: 1, Z: 0},
	}
}

func (cam *Camera) Initialize() {
	cam.imageHeight = int(float64(cam.ImageWidth) / cam.AspectRatio)
	cam.pixelSampleScale = 1. / float64(cam.NumberOfSamples)
	cam.cameraCenter = cam.LookFrom

	cam.focalLength = cam.LookFrom.Subtract(cam.LookAt).Length()
	theta := util.DegreesToRadians(cam.FieldOfView)
	h := math.Tan(theta / 2)
	cam.viewportHeight = 2 * h * cam.focalLength
	cam.viewportWidth = cam.viewportHeight * (float64(cam.ImageWidth) / float64(cam.imageHeight))

	cam.w = cam.LookFrom.Subtract(cam.LookAt).Norm()
	cam.u = cam.VUp.Cross(cam.w).Norm()
	cam.v = cam.w.Cross(cam.u).Norm()

	viewportU := cam.u.Scale(cam.viewportWidth)
	viewportV := cam.v.Negative().Scale(cam.viewportHeight)

	cam.deltaU = viewportU.Divide(float64(cam.ImageWidth))
	cam.deltaV = viewportV.Divide(float64(cam.imageHeight))

	cam.viewportUpperLeft = cam.cameraCenter.
		Subtract(cam.w.Scale(cam.focalLength)).
		Subtract(viewportU.Divide(2)).
		Subtract(viewportV.Divide(2))
	cam.pixel00Loc = cam.viewportUpperLeft.Add(cam.deltaU.Add(cam.deltaV).Divide(2))
}

func rayColor(r ray.Ray, depth int, world hittable.HittableList) color.Color {
	if depth < 0 {
		return color.Color{}
	}

	if hit, material, hitRecord := world.Hit(r, interval.New(0.001, math.Inf(1))); hit {
		if scattered, attenuation, scatteredRay := material.Scatter(r, hitRecord); scattered {
			return rayColor(scatteredRay, depth-1, world).Mult(attenuation)
		}

		return color.Color{}
	}

	a := .5 * (r.Direction.Norm().Y + 1.0)
	startColor := color.Color{R: 1., G: 1., B: 1.}
	endColor := color.Color{R: .5, G: .7, B: 1.}
	return startColor.Scale(1. - a).Add(endColor.Scale(a))
}

func (cam Camera) getRay(i int, j int) ray.Ray {
	offset := cam.sampleSqure()
	pixelCenter := cam.pixel00Loc.
		Add(cam.deltaU.Scale(float64(j) + offset.X)).
		Add(cam.deltaV.Scale(float64(i) + offset.Y))

	ray := ray.Ray{
		Origin:    cam.cameraCenter,
		Direction: pixelCenter.Subtract(cam.cameraCenter),
	}

	return ray
}

func (cam Camera) sampleSqure() vec3.Vec3 {
	return vec3.Vec3{
		X: rand.Float64() - .5,
		Y: rand.Float64() - .5,
		Z: .0,
	}
}

func (cam Camera) Render(world hittable.HittableList) {
	pb := progress_bar.ProgressBar{Max: cam.imageHeight, Length: 30, Writer: os.Stderr}

	fmt.Println("P3")
	fmt.Println(cam.ImageWidth, cam.imageHeight, cam.maxColor)

	for i := range cam.imageHeight {
		for j := range cam.ImageWidth {
			c := color.Color{}

			for range cam.NumberOfSamples {
				ray := cam.getRay(i, j)
				c = c.Add(rayColor(ray, cam.MaxDepth, world))
			}

			c.Scale(cam.pixelSampleScale).Write()
		}

		pb.Tick()
	}
}
