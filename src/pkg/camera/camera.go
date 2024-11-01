package camera

import (
	"fmt"
	"math"
	"os"

	"jkbkaiser/ray-tracing-go/pkg/color"
	"jkbkaiser/ray-tracing-go/pkg/hittable"
	"jkbkaiser/ray-tracing-go/pkg/progress_bar"
	"jkbkaiser/ray-tracing-go/pkg/ray"
	"jkbkaiser/ray-tracing-go/pkg/util/interval"
	"jkbkaiser/ray-tracing-go/pkg/vec3"
)

type Camera struct {
	maxColor int

	AspectRatio float64

	ImageWidth  int
	imageHeight int

	FocalLength  float64
	CameraCenter vec3.Vec3

	ViewportHeight float64
	viewportWidth  float64

	viewportUpperLeft vec3.Vec3
	pixel00Loc        vec3.Vec3
	deltaU            vec3.Vec3
	deltaV            vec3.Vec3
}

func New() Camera {
	return Camera{
		AspectRatio:    16. / 9.,
		ImageWidth:     400,
		FocalLength:    1.,
		ViewportHeight: 2.,
	}
}

func (c *Camera) Initialize() {
	c.imageHeight = int(float64(c.ImageWidth) / c.AspectRatio)
	c.viewportWidth = c.ViewportHeight * (float64(c.ImageWidth) / float64(c.imageHeight))

	viewportU := vec3.Vec3{X: c.viewportWidth, Y: 0., Z: 0.}
	viewportV := vec3.Vec3{X: 0.0, Y: -c.ViewportHeight, Z: 0.}

	c.deltaU = viewportU.Divide(float64(c.ImageWidth))
	c.deltaV = viewportV.Divide(float64(c.imageHeight))

	c.viewportUpperLeft = c.CameraCenter.
		Subtract(vec3.Vec3{X: 0., Y: 0., Z: c.FocalLength}).
		Subtract(viewportU.Divide(2)).
		Subtract(viewportV.Divide(2))
	c.pixel00Loc = c.viewportUpperLeft.Add(c.deltaU.Add(c.deltaV).Divide(2))
}

func rayColor(r ray.Ray, world hittable.HittableList) color.Color {
	hitRecord := hittable.HitRecord{}

	if world.Hit(r, interval.New(0, math.Inf(1)), &hitRecord) {
		return color.FromVec(
			hitRecord.Normal.Add(vec3.Vec3{X: 1, Y: 1, Z: 1}).Scale(.5),
		)
	}

	a := .5 * (r.Direction.Norm().Y + 1.0)
	startColor := color.Color{R: 1., G: 1., B: 1.}
	endColor := color.Color{R: .5, G: .7, B: 1.}
	return startColor.Scale(1. - a).Add(endColor.Scale(a))
}

func (c Camera) Render(world hittable.HittableList) {
	pb := progress_bar.ProgressBar{Max: c.imageHeight, Length: 30, Writer: os.Stderr}

	fmt.Println("P3")
	fmt.Println(c.ImageWidth, c.imageHeight, c.maxColor)

	for i := range c.imageHeight {
		for j := range c.ImageWidth {
			pixelCenter := c.pixel00Loc.
				Add(c.deltaU.Scale(float64(j))).
				Add(c.deltaV.Scale(float64(i)))

			ray := ray.Ray{
				Origin:    c.CameraCenter,
				Direction: pixelCenter.Subtract(c.CameraCenter),
			}

			c := rayColor(ray, world)
			c.Write()
		}

		pb.Tick()
	}
}
