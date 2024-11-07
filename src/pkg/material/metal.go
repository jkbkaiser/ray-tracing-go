package material

import (
	"jkbkaiser/ray-tracing-go/pkg/color"
	"jkbkaiser/ray-tracing-go/pkg/hitrecord"
	"jkbkaiser/ray-tracing-go/pkg/ray"
	"jkbkaiser/ray-tracing-go/pkg/vec3"
)

type Metal struct {
	Albedo color.Color
	Fuzz   float64
}

func (m Metal) Scatter(r ray.Ray, hitRecord hitrecord.HitRecord) (bool, color.Color, ray.Ray) {
	reflected := vec3.Reflect(r.Direction, hitRecord.Normal)
	scattered := ray.Ray{Origin: hitRecord.Point, Direction: reflected.Norm().Add(vec3.RandomUnit().Scale(m.Fuzz))}
	return scattered.Direction.Dot(hitRecord.Normal) > 0, m.Albedo, scattered
}
