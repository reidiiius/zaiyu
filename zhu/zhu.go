package zhu

import "math"

var Zaiyu = make([]float64, 1)

func Semitone(n float64) float64 {
	return math.Pow(2, n/12) * 1
}
