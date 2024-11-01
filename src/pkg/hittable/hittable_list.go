package hittable

import (
	"jkbkaiser/ray-tracing-go/pkg/ray"
	"jkbkaiser/ray-tracing-go/pkg/util/interval"
)

type HittableList struct {
	objects []Hittable
}

func (h HittableList) Add(object Hittable) HittableList {
	h.objects = append(h.objects, object)
	return h
}

func (h HittableList) Hit(r ray.Ray, rayT interval.Interval, rec *HitRecord) bool {
	tmpRec := HitRecord{}
	hitAnything := false
	closestSoFar := rayT.Max

	for _, object := range h.objects {
		if object.Hit(r, interval.New(rayT.Min, closestSoFar), &tmpRec) {
			hitAnything = true
			closestSoFar = tmpRec.t
			*rec = tmpRec
		}
	}

	return hitAnything
}
