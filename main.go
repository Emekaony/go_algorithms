package main

import (
	"fmt"
	"image"
	"log"

	"github.com/disintegration/imaging"
)

// prevent passing raw strings into intensity function. This acts as an enum
type Intensity string

const (
	AVERAGE    Intensity = "average"
	MAX_MIN    Intensity = "max_min"
	LUMINOSITY Intensity = "luminosity"
)

func main() {
	// load and open the image
	src, err := imaging.Open("assets/images/ascii-pineapple.jpg")
	if err != nil {
		log.Fatalf("Failed to open image: %v", err)
	}
	// first task: get the pixel matrix
	arr := getPixelMatrix(src)
	dummyArr := make([][][]uint8, len(arr)) // gotta make a copy so the main one is not altered!
	copy(dummyArr, arr)
	intensityMatrix := getIntensityMatrix(dummyArr, MAX_MIN)
	fmt.Println(intensityMatrix[0])
}

// converts RGBA from uint32 to 8-bit pixel value
func rgbaToPixel(r, g, b, _ uint32) (uint8, uint8, uint8) {
	// ignoring the alpha for now!
	red := uint8(r >> 8)
	green := uint8(g >> 8)
	blue := uint8(b >> 8)
	return red, green, blue
}

// helper function to return the height and width of an image
func getImageHeightAndWidth(im image.Image) (int, int) {
	return im.Bounds().Dx(), im.Bounds().Dy()
}

// returns thr (r, g, b) value for each pixel in the image
func getPixelMatrix(img image.Image) [][][]uint8 {
	width, height := getImageHeightAndWidth(img)
	result := [][][]uint8{}

	// start from top left corner and go to the right
	for j := img.Bounds().Min.Y; j < height; j++ {
		// we want to have some sort of a row vector?
		row := [][]uint8{}
		for i := img.Bounds().Min.X; i < width; i++ {
			// get the correct red, green, and blue color for each pixel
			r, g, b := rgbaToPixel(img.At(i, j).RGBA())
			tup := []uint8{r, g, b}
			row = append(row, tup)
		}
		result = append(result, row)
	}
	return result
}

// maps (r, g, b) pixel values into a singular (i) intensity value depending on the algorithm specified by algoName
func getIntensityMatrix(pixelMatrix [][][]uint8, algoName Intensity) [][][]uint8 {
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
