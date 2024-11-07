package material

import (
	"jkbkaiser/ray-tracing-go/pkg/color"
	"jkbkaiser/ray-tracing-go/pkg/hitrecord"
	"jkbkaiser/ray-tracing-go/pkg/ray"
	"jkbkaiser/ray-tracing-go/pkg/vec3"
	"math"
	"math/rand/v2"
)

func reflectance(cosine, refractionIndex float64) float64 {
	r0 := (1 - refractionIndex) / (1 + refractionIndex)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}

type Dialectric struct {
	RefractionIndex float64
}

func (m Dialectric) Scatter(r ray.Ray, hitRecord hitrecord.HitRecord) (bool, color.Color, ray.Ray) {
	attenuation := color.Color{R: 1, G: 1, B: 1}

	var ri float64
	if hitRecord.FrontFace {
		ri = 1. / m.RefractionIndex
	} else {
		ri = m.RefractionIndex
	}

	unitDirection := r.Direction.Norm()

	cosTheta := math.Min(unitDirection.Negative().Dot(hitRecord.Normal), 1.)
	sinTheta := math.Sqrt(1. - cosTheta*cosTheta)

	var direction vec3.Vec3

	if ri*sinTheta > 1. || reflectance(cosTheta, ri) > rand.Float64() {
		// Cannot refract
		direction = vec3.Reflect(unitDirection, hitRecord.Normal)
	} else {
		direction = vec3.Refract(unitDirection, hitRecord.Normal, ri)
	}

	scattered := ray.Ray{Origin: hitRecord.Point, Direction: direction}
	return true, attenuation, scattered
}
