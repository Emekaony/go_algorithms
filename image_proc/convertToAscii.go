package image_proc

import "math"

func ConvertToAscii(intensityMatrix [][]float64) [][]string {
	asciiChars := "`^\",:;Il!i~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"
	asciiMatrix := [][]string{}
	for _, row := range intensityMatrix {
		asciiRow := []string{}
		for _, r := range row {
			temp := math.Floor(r * float64(len(asciiChars)-1))
			// fmt.Println("r is ", r, "Length of asciiChars is", len(asciiChars), "and temp is: ", temp)
			// if temp != 0 {
			// 	temp--
			// }

			char := string(asciiChars[int(temp)])
			// fmt.Println("char is ", char)
			asciiRow = append(asciiRow, char)
		}
		asciiMatrix = append(asciiMatrix, asciiRow)
	}
	return asciiMatrix
}
