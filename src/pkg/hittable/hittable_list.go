package hittable

import (
	"jkbkaiser/ray-tracing-go/pkg/hitrecord"
	"jkbkaiser/ray-tracing-go/pkg/material"
	"jkbkaiser/ray-tracing-go/pkg/ray"
	"jkbkaiser/ray-tracing-go/pkg/util/interval"
)

type HittableList struct {
	objects []Hittable
}

func (h *HittableList) Add(object Hittable) {
	h.objects = append(h.objects, object)
}

func (h HittableList) Hit(r ray.Ray, rayT interval.Interval) (bool, material.Material, hitrecord.HitRecord) {
	var rec hitrecord.HitRecord
	var mat material.Material

	tmpRec := hitrecord.HitRecord{}
	hitAnything := false
	closestSoFar := rayT.Max

	for _, object := range h.objects {
		if object.Hit(r, interval.New(rayT.Min, closestSoFar), &tmpRec) {
			hitAnything = true
			closestSoFar = tmpRec.T
			rec = tmpRec
			mat = object.Mat()
		}
	}

	return hitAnything, mat, rec
}
