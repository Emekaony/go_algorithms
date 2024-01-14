package main

import (
	"image"
	"log"

	"github.com/disintegration/imaging"
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
	// second task: get the brightness matrix
	// intensityMatrix := getIntensityMatrix(dummyArr)
}

// converts RGBA from uint32 to 8-bit pixel value
func rgbaToPixel(r, g, b, _ uint32) (uint8, uint8, uint8) {
	// ignoring the alpha for now!
	red := uint8(r >> 8)
	green := uint8(g >> 8)
	blue := uint8(b >> 8)
	return red, green, blue
}

func getImageHeightAndWidth(im image.Image) (int, int) {
	return im.Bounds().Dx(), im.Bounds().Dy()
}

// this is the function that will be behind
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

// see how we can pass arrays as references instead of by value
func getIntensityMatrix(pixelMatrix [][][]uint8, algoName string) [][][]uint8 {
	intensityMatrix := make([][][]uint8, len(pixelMatrix))
	var intensity uint32
	for i := 0; i < len(pixelMatrix); i++ {
		for j := 0; j < len(pixelMatrix[i]); j++ {
			a, b, c := pixelMatrix[i][j][0], pixelMatrix[i][j][1], pixelMatrix[i][j][2]
			// convert to 32 bit for temporary math
			switch algoName {
			case "average":
				intensity = (uint32(a) + uint32(b) + uint32(c)) / uint32(3)
			case "max_min":
				intensity = max(uint32(a), uint32(b), uint32(c)) + min(uint32(a), uint32(b), uint32(c))
			case "luminosity":
				temp := 0.21*float32(a) + 0.72*float32(b) + 0.72*float32(c)
				intensity = uint32(temp)
			default:
				panic("unknown algorithm name!")
			}

			// convert back to 8 bit
			intensityMatrix[i][j] = []uint8{uint8(intensity)}
		}
	}
	return pixelMatrix
}
