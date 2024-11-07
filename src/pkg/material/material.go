package material

import (
	"jkbkaiser/ray-tracing-go/pkg/color"
	"jkbkaiser/ray-tracing-go/pkg/hitrecord"
	"jkbkaiser/ray-tracing-go/pkg/ray"
)

type Material interface {
	Scatter(r ray.Ray, hitRecord hitrecord.HitRecord) (bool, color.Color, ray.Ray)
}
