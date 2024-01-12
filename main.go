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
	pixelArray := getPixelArray(src)

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
func getPixelArray(img image.Image) [][]uint8 {
	width, height := getImageHeightAndWidth(img)
	result := make([][]uint8)
	for i := img.Bounds().Min.X; i < width; i++ {
		for j := img.Bounds().Min.Y; j < height; j++ {

		}
	}
}
