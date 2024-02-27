package image_proc

const MAX_PIXEL_VALUE float64 = 255.0

func NormalizeIntensityMatrix(intensityMatrix [][]float64) [][]float64 {
	var normalizedIntensityMatrix [][]float64

	// Normalize each pixel value
	for _, row := range intensityMatrix {
		var rescaledRow []float64
		// minPixel := findMinPixel(row)
		// maxPixel := findMaxPixel(row)

		for _, p := range row {
			r := (p / MAX_PIXEL_VALUE)
			rescaledRow = append(rescaledRow, r)
		}
		normalizedIntensityMatrix = append(normalizedIntensityMatrix, rescaledRow)
	}
	// fmt.Println(normalizedIntensityMatrix)

	return normalizedIntensityMatrix
}

// // returns the max value from a vector
// func findMaxPixel(intensityMatrix []float64) float64 {
// 	max := intensityMatrix[0]
// 	for _, item := range intensityMatrix {
// 		if item > max {
// 			max = item
// 		}
// 	}
// 	return max
// }

// // returns the min value from a vector
// func findMinPixel(intensityMatrix []float64) float64 {
// 	min := intensityMatrix[0]
// 	for _, item := range intensityMatrix {
// 		if item < min {
// 			min = item
// 		}
// 	}

// 	return min
// }
