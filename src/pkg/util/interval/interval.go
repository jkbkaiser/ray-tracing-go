package interval

import "math"

type Interval struct {
	Min, Max float64
}

func Default() Interval {
	return Interval{
		Min: math.Inf(-1),
		Max: math.Inf(1),
	}
}

func New(min float64, max float64) Interval {
	return Interval{
		Min: min,
		Max: max,
	}
}

func (i Interval) Size() float64 {
	return i.Max - i.Min
}

func (i Interval) Contains(v float64) bool {
	return i.Min <= v && v <= i.Max
}

func (i Interval) Surrounds(v float64) bool {
	return i.Min < v && v < i.Max
}
