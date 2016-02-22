package zhu

import "math"

var Zaiyu = make([]float64, 1)

func Semitone(n float64) float64 {
	return math.Pow(2, n/12) * 1
}

func Guillaume(n float64) float64 {
	return math.Pow(2, n/19) * 1
}

func Microtone(n float64) float64 {
	return math.Pow(2, n/24) * 1
}

func Archicembalo(n float64) float64 {
	return math.Pow(2, n/31) * 1
}

func ChingFang(n float64) float64 {
	return math.Pow(2, n/53) * 1
}
