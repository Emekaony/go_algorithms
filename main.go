package main

import (
	"ascii_art/image_proc"
	"fmt"
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
	pixelMatrix := image_proc.GetPixelMatrix(src)
	intensityMatrix := image_proc.GetIntensityMatrix(pixelMatrix, image_proc.MAX_MIN)
	fmt.Println(intensityMatrix[0][0])
	fmt.Println(pixelMatrix[0][0])
}
