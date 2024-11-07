package hitrecord

import (
	"jkbkaiser/ray-tracing-go/pkg/ray"
	"jkbkaiser/ray-tracing-go/pkg/vec3"
)

type HitRecord struct {
	Point     vec3.Vec3
	Normal    vec3.Vec3
	T         float64
	FrontFace bool
}

func (h *HitRecord) SetFrontFaceNormal(r ray.Ray, outwardNormal vec3.Vec3) {
	h.FrontFace = r.Direction.Dot(outwardNormal) < 0

	if h.FrontFace {
		h.Normal = outwardNormal
	} else {
		h.Normal = outwardNormal.Negative()
	}
}
