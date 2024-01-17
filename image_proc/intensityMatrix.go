package image_proc

// prevent passing raw strings into intensity function. This acts as an enum
type Intensity string

const (
	AVERAGE    Intensity = "average"
	MAX_MIN    Intensity = "max_min"
	LUMINOSITY Intensity = "luminosity"
)

// maps (r, g, b) pixel values into a singular (i) intensity value depending on the algorithm specified by algoName
func GetIntensityMatrix(pixelMatrix [][][]uint8, algoName Intensity) [][]float64 {
	// try float64 intensity matrix
	intensityMatrix := [][]float64{}
	var intensity float64
	for i := 0; i < len(pixelMatrix); i++ {
		row := []float64{}
		for j := 0; j < len(pixelMatrix[i]); j++ {
			a, b, c := pixelMatrix[i][j][0], pixelMatrix[i][j][1], pixelMatrix[i][j][2]
			// convert to 32 bit for temporary math
			switch algoName {
			case "average":
				intensity = (float64(a) + float64(b) + float64(c)) / float64(3)
			case "max_min":
				intensity = (max(float64(a), float64(b), float64(c)) + min(float64(a), float64(b), float64(c))) / 2
			case "luminosity":
				intensity = 0.21*float64(a) + 0.72*float64(b) + 0.72*float64(c)
				// intensity = float64(temp)
			default:
				panic("unknown algorithm name!")
			}

			// intensity matrix is now 2-dimensional
			row = append(row, float64(intensity))
		}
		intensityMatrix = append(intensityMatrix, row)
	}
	return intensityMatrix
}
