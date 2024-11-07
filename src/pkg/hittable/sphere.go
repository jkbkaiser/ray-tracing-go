package hittable

import (
	"math"

	"jkbkaiser/ray-tracing-go/pkg/hitrecord"
	"jkbkaiser/ray-tracing-go/pkg/material"
	"jkbkaiser/ray-tracing-go/pkg/ray"
	"jkbkaiser/ray-tracing-go/pkg/util/interval"
	"jkbkaiser/ray-tracing-go/pkg/vec3"
)

type Sphere struct {
	Center vec3.Vec3
	Radius float64
	mat    material.Material
}

func NewSphere(center vec3.Vec3, radius float64, mat material.Material) Sphere {
	return Sphere{center, radius, mat}
}

func (s Sphere) Mat() material.Material {
	return s.mat
}

func (s Sphere) Hit(r ray.Ray, rayT interval.Interval, hitRecord *hitrecord.HitRecord) bool {
	oc := s.Center.Subtract(r.Origin)
	a := r.Direction.Dot(r.Direction)
	h := r.Direction.Dot(oc)
	c := oc.Dot(oc) - s.Radius*s.Radius

	discriminant := h*h - a*c

	if discriminant < 0 {
		return false
	}

	sqrtd := math.Sqrt(discriminant)

	root := (h - sqrtd) / a
	if !rayT.Surrounds(root) {
		root = (h + sqrtd) / a

		if !rayT.Surrounds(root) {
			return false
		}
	}

	hitRecord.T = root
	hitRecord.Point = r.At(root)
	outwardNormal := hitRecord.Point.Subtract(s.Center).Divide(s.Radius)
	hitRecord.SetFrontFaceNormal(r, outwardNormal)

	return true
}
