package math_alg

func Linspace(start, end float64, numPoints int) []float64 {
	step := (end - start) / float64(numPoints-1)
	result := make([]float64, numPoints)

	for i := 0; i < numPoints; i++ {
		result[i] = start + float64(i)*step
	}

	return result
}
