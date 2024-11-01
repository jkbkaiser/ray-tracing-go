package ray

import "jkbkaiser/ray-tracing-go/pkg/vec3"

type Ray struct {
	Origin    vec3.Vec3
	Direction vec3.Vec3
}

func (r *Ray) At(t float64) vec3.Vec3 {
	return r.Origin.Add(r.Direction.Scale(t))
}
