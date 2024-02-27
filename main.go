package main

import (
	"ascii_art/image_proc"
	"log"

	"github.com/disintegration/imaging"
)

func main() {
	// load and open the image
	src, err := imaging.Open("assets/images/ascii_apple.jpg") // this was src
	if err != nil {
		log.Fatalf("Failed to open image: %v", err)
	}
	resized := imaging.Resize(src, 200, 0, imaging.BSpline) // play around withh different sizes until it fits properly on the console.
	// first task: get the pixel matrix
	pixelMatrix := image_proc.GetPixelMatrix(resized) // see if resized works instead of src
	intensityMatrix := image_proc.GetIntensityMatrix(pixelMatrix, image_proc.LUMINOSITY)
	normalixedMatrix := image_proc.NormalizeIntensityMatrix(intensityMatrix)
	asciiMatrix := image_proc.ConvertToAscii(normalixedMatrix)
	image_proc.PrintAsciiMatrix(asciiMatrix) // moment of truth
}
