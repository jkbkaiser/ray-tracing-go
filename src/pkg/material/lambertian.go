package material

import (
	"jkbkaiser/ray-tracing-go/pkg/color"
	"jkbkaiser/ray-tracing-go/pkg/hitrecord"
	"jkbkaiser/ray-tracing-go/pkg/ray"
	"jkbkaiser/ray-tracing-go/pkg/vec3"
)

type Lambertian struct {
	Albedo color.Color
}

func (l Lambertian) Scatter(r ray.Ray, hitRecord hitrecord.HitRecord) (bool, color.Color, ray.Ray) {
	scatterDirection := hitRecord.Normal.Add(vec3.RandomUnit())

	if scatterDirection.NearZero() {
		scatterDirection = hitRecord.Normal
	}

	scattered := ray.Ray{Origin: hitRecord.Point, Direction: scatterDirection}
	return true, l.Albedo, scattered
}
