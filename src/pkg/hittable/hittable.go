package hittable

import (
	"jkbkaiser/ray-tracing-go/pkg/hitrecord"
	"jkbkaiser/ray-tracing-go/pkg/material"
	"jkbkaiser/ray-tracing-go/pkg/ray"
	"jkbkaiser/ray-tracing-go/pkg/util/interval"
)

type Hittable interface {
	Mat() material.Material
	Hit(ray.Ray, interval.Interval, *hitrecord.HitRecord) bool
}
