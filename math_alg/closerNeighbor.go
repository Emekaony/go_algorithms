package math_alg

import (
	"math"
)

func CloserNeighbor(number, neighbor1, neighbor2 float64) float64 {
	distance1 := math.Abs(number - neighbor1)
	distance2 := math.Abs(number - neighbor2)

	if distance1 < distance2 {
		return neighbor1
	}

	return neighbor2
}
