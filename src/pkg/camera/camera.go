package camera

import (
	"sync"
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

	cameraCenter vec3.Vec3
	LookFrom     vec3.Vec3
	LookAt       vec3.Vec3
	VUp          vec3.Vec3
	u            vec3.Vec3
	v            vec3.Vec3
	w            vec3.Vec3

	DefocusAngle float64
	FocusDist    float64
	defocusDiskU vec3.Vec3
	defocusDiskV vec3.Vec3

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
		DefocusAngle:    0,
		FocusDist:       10,
	}
}

func (cam *Camera) Initialize() {
	cam.imageHeight = int(float64(cam.ImageWidth) / cam.AspectRatio)
	cam.pixelSampleScale = 1. / float64(cam.NumberOfSamples)
	cam.cameraCenter = cam.LookFrom

	// cam.focalLength = cam.LookFrom.Subtract(cam.LookAt).Length()
	theta := util.DegreesToRadians(cam.FieldOfView)
	h := math.Tan(theta / 2)
	cam.viewportHeight = 2 * h * cam.FocusDist
	cam.viewportWidth = cam.viewportHeight * (float64(cam.ImageWidth) / float64(cam.imageHeight))

	cam.w = cam.LookFrom.Subtract(cam.LookAt).Norm()
	cam.u = cam.VUp.Cross(cam.w).Norm()
	cam.v = cam.w.Cross(cam.u).Norm()

	viewportU := cam.u.Scale(cam.viewportWidth)
	viewportV := cam.v.Negative().Scale(cam.viewportHeight)

	cam.deltaU = viewportU.Divide(float64(cam.ImageWidth))
	cam.deltaV = viewportV.Divide(float64(cam.imageHeight))

	cam.viewportUpperLeft = cam.cameraCenter.
		Subtract(cam.w.Scale(cam.FocusDist)).
		Subtract(viewportU.Divide(2.)).
		Subtract(viewportV.Divide(2.))
	cam.pixel00Loc = cam.viewportUpperLeft.Add(cam.deltaU.Add(cam.deltaV).Divide(2.))

	defocusRadius := cam.FocusDist * math.Tan(util.DegreesToRadians(cam.DefocusAngle/2.))
	cam.defocusDiskU = cam.u.Scale(defocusRadius)
	cam.defocusDiskV = cam.v.Scale(defocusRadius)
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

func (cam Camera) defocusDiskSample() vec3.Vec3 {
	p := vec3.RandomInUnitDisk()
	return cam.cameraCenter.Add(cam.defocusDiskU.Scale(p.X)).Add(cam.defocusDiskV.Scale(p.Y))
}

func (cam Camera) getRay(i int, j int) ray.Ray {
	offset := cam.sampleSqure()
	pixelSample := cam.pixel00Loc.
		Add(cam.deltaU.Scale(float64(j) + offset.X)).
		Add(cam.deltaV.Scale(float64(i) + offset.Y))

	var rayOrigin vec3.Vec3
	if cam.DefocusAngle <= 0. {
		rayOrigin = cam.cameraCenter
	} else {
		rayOrigin = cam.defocusDiskSample()
	}

	ray := ray.Ray{
		Origin:    rayOrigin,
		Direction: pixelSample.Subtract(rayOrigin),
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
	parrallel := true

	fmt.Println("P3")
	fmt.Println(cam.ImageWidth, cam.imageHeight, cam.maxColor)

	if parrallel {
	  pb := progress_bar.ProgressBar{Max: cam.imageHeight * cam.ImageWidth, Length: 30, Writer: os.Stderr}
		type pixelResult struct {
			color color.Color
			order int
		}
		var wg sync.WaitGroup
		pixChan := make(chan pixelResult, cam.imageHeight * cam.ImageWidth)
		
		order := 0
		for i := range cam.imageHeight {
			for j := range cam.ImageWidth {
				wg.Add(1)
				iCopy, jCopy, orderCopy := i, j, order

				go func() {
					defer wg.Done()
					c := color.Color{}

					for range cam.NumberOfSamples {
						ray := cam.getRay(iCopy, jCopy)
						c = c.Add(rayColor(ray, cam.MaxDepth, world))
					}

					// c.Scale(cam.pixelSampleScale).Write()
					pixChan <- pixelResult{c, orderCopy}
				}()

				order++
			}
		}

		go func() {
			wg.Wait()
			close(pixChan)
		}()

		pixelResults := make([]color.Color, cam.imageHeight * cam.ImageWidth)

		for pixelResult := range pixChan {
			pixelResults[pixelResult.order] = pixelResult.color
			pb.Tick()
		}

		for _, element := range pixelResults {
			element.Scale(cam.pixelSampleScale).Write()
		}
	} else {
	  pb := progress_bar.ProgressBar{Max: cam.imageHeight, Length: 30, Writer: os.Stderr}
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
}
