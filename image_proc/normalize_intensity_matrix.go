package image_proc

const MAX_PIXEL_VALUE float64 = 255.0

func NormalizeIntensityMatrix(intensityMatrix [][]float64) [][]float64 {
	var normalizedIntensityMatrix [][]float64

	// Find max and min pixel values
	maxPixel := findMaxPixel(intensityMatrix)
	minPixel := findMinPixel(intensityMatrix)

	// Normalize each pixel value
	for _, row := range intensityMatrix {
		var rescaledRow []float64
		for _, p := range row {
			r := MAX_PIXEL_VALUE * (p - minPixel) / (maxPixel - minPixel)
			rescaledRow = append(rescaledRow, r)
		}
		normalizedIntensityMatrix = append(normalizedIntensityMatrix, rescaledRow)
	}

	return normalizedIntensityMatrix
}

func findMaxPixel(intensityMatrix [][]float64) float64 {
	maxPixel := intensityMatrix[0][0]

	for i := 0; i < len(intensityMatrix); i++ {
		for j := 1; j < len(intensityMatrix[i]); j++ {
			if intensityMatrix[i][j] > maxPixel {
				maxPixel = intensityMatrix[i][j]
			}
		}
	}

	return maxPixel
}

func findMinPixel(intensityMatrix [][]float64) float64 {
	minPixel := intensityMatrix[0][0]

	for i := 0; i < len(intensityMatrix); i++ {
		for j := 1; j < len(intensityMatrix[i]); j++ {
			if intensityMatrix[i][j] < minPixel {
				minPixel = intensityMatrix[i][j]
			}
		}
	}

	return minPixel
}
