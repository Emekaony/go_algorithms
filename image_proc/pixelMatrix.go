package image_proc

import "image"

// returns thr (r, g, b) value for each pixel in the image
func GetPixelMatrix(img image.Image) [][][]uint8 {
	width, height := GetImageHeightAndWidth(img)
	result := [][][]uint8{}

	// start from top left corner and go to the right
	for j := img.Bounds().Min.Y; j < height; j++ {
		// we want to have some sort of a row vector?
		row := [][]uint8{}
		for i := img.Bounds().Min.X; i < width; i++ {
			// get the correct red, green, and blue color for each pixel
			r, g, b := RgbaToPixel(img.At(i, j).RGBA())
			tup := []uint8{r, g, b}
			row = append(row, tup)
		}
		result = append(result, row)
	}
	return result
}
