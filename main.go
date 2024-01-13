package main

import (
	"fmt"
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
	arr := getPixelMatrix(src)
	brightnessMatrix := getBrighnessMatrix(arr)
	fmt.Println(arr[0][0])
	fmt.Println(brightnessMatrix[0][0])
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
func getBrighnessMatrix(pixelMatrix [][][]uint8) [][][]uint8 {
	result := pixelMatrix[:] // make a copy of the pixelMatrix

	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			a, b, c := result[i][j][0], result[i][j][1], result[i][j][2]
			// convert to 32 bit for temporary math
			avg := (uint32(a) + uint32(b) + uint32(c)) / uint32(3)
			// convert back to 8 bit
			result[i][j] = []uint8{uint8(avg)}
		}
	}
	return result
}
