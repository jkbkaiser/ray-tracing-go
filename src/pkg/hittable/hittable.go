package hittable

import (
	"jkbkaiser/ray-tracing-go/pkg/ray"
	"jkbkaiser/ray-tracing-go/pkg/util/interval"
	"jkbkaiser/ray-tracing-go/pkg/vec3"
)

type HitRecord struct {
	Point     vec3.Vec3
	Normal    vec3.Vec3
	t         float64
	frontFace bool
}

func (h *HitRecord) SetFrontFaceNormal(r ray.Ray, outwardNormal vec3.Vec3) {
	h.frontFace = r.Direction.Dot(outwardNormal) < 0

	if h.frontFace {
		h.Normal = outwardNormal
	} else {
		h.Normal = outwardNormal.Negative()
	}
}

type Hittable interface {
	Hit(ray.Ray, interval.Interval, *HitRecord) bool
}
