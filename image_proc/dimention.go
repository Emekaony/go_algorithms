package image_proc

import "image"

// helper function to return the height and width of an image
func GetImageHeightAndWidth(im image.Image) (int, int) {
	return im.Bounds().Dx(), im.Bounds().Dy()
}
