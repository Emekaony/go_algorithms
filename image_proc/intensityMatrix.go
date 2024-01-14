package image_proc

// prevent passing raw strings into intensity function. This acts as an enum
type Intensity string

const (
	AVERAGE    Intensity = "average"
	MAX_MIN    Intensity = "max_min"
	LUMINOSITY Intensity = "luminosity"
)

// maps (r, g, b) pixel values into a singular (i) intensity value depending on the algorithm specified by algoName
func GetIntensityMatrix(pixelMatrix [][][]uint8, algoName Intensity) [][][]uint8 {
	intensityMatrix := [][][]uint8{}
	var intensity uint32
	for i := 0; i < len(pixelMatrix); i++ {
		row := [][]uint8{}
		for j := 0; j < len(pixelMatrix[i]); j++ {
			a, b, c := pixelMatrix[i][j][0], pixelMatrix[i][j][1], pixelMatrix[i][j][2]
			// convert to 32 bit for temporary math
			switch algoName {
			case "average":
				intensity = (uint32(a) + uint32(b) + uint32(c)) / uint32(3)
			case "max_min":
				intensity = (max(uint32(a), uint32(b), uint32(c)) + min(uint32(a), uint32(b), uint32(c))) / 2
			case "luminosity":
				temp := 0.21*float32(a) + 0.72*float32(b) + 0.72*float32(c)
				intensity = uint32(temp)
			default:
				panic("unknown algorithm name!")
			}

			// convert back to 8 bit
			row = append(row, []uint8{uint8(intensity)})
		}
		intensityMatrix = append(intensityMatrix, row)
	}
	return intensityMatrix
}
