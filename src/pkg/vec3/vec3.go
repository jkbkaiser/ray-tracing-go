package vec3

import (
	"jkbkaiser/ray-tracing-go/pkg/util"
	"math"
	"math/rand/v2"
)

type Vec3 struct {
	X, Y, Z float64
}

func (v Vec3) Add(other Vec3) Vec3 {
	return Vec3{v.X + other.X, v.Y + other.Y, v.Z + other.Z}
}

func (v Vec3) Subtract(other Vec3) Vec3 {
	return Vec3{v.X - other.X, v.Y - other.Y, v.Z - other.Z}
}

func (v Vec3) Dot(other Vec3) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

func (v Vec3) Cross(other Vec3) Vec3 {
	return Vec3{
		v.Y*other.Z - v.Z*other.Y,
		v.Z*other.X - v.X*other.Z,
		v.X*other.Y - v.Y*other.X,
	}
}

func (v Vec3) Scale(value float64) Vec3 {
	return Vec3{value * v.X, value * v.Y, value * v.Z}
}

func (v Vec3) Divide(value float64) Vec3 {
	return v.Scale(1 / value)
}

func (v Vec3) LengthSquared() float64 {
	return v.Dot(v)
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.Dot(v))
}

func (v Vec3) Negative() Vec3 {
	return v.Scale(-1)
}

func (v Vec3) Norm() Vec3 {
	return v.Divide(v.Length())
}

func (v Vec3) NearZero() bool {
	s := 1e-8
	return (math.Abs(v.X) < s) && (math.Abs(v.Y) < s) && (math.Abs(v.Z) < s)
}

func Reflect(v, norm Vec3) Vec3 {
	return v.Subtract(norm.Scale(v.Dot(norm)).Scale(2))
}

func Refract(v, norm Vec3, etaiOverEtat float64) Vec3 {
	cosTheta := math.Min(v.Negative().Dot(norm), 1.)
	rOutPerp := v.Add(norm.Scale(cosTheta)).Scale(etaiOverEtat)
	rOutParallel := norm.Scale(-math.Sqrt(math.Abs(1. - rOutPerp.LengthSquared())))
	return rOutPerp.Add(rOutParallel)
}

func Random() Vec3 {
	return Vec3{
		rand.Float64(),
		rand.Float64(),
		rand.Float64(),
	}
}

func RandomRange(min, max float64) Vec3 {
	return Vec3{
		util.RandomFloat(min, max),
		util.RandomFloat(min, max),
		util.RandomFloat(min, max),
	}
}

func RandomUnit() Vec3 {
	for {
		v := RandomRange(-1, 1)
		l := v.LengthSquared()

		if 1e-160 < l && l <= 1 {
			return v.Scale(math.Sqrt(l))
		}
	}
}

func RandomOnHemisphere(norm Vec3) Vec3 {
	v := RandomUnit()

	if v.Dot(norm) > .0 {
		return v
	} else {
		return v.Negative()
	}
}
