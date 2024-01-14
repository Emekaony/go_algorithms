package image_proc

// converts RGBA from uint32 to 8-bit pixel value
func RgbaToPixel(r, g, b, _ uint32) (uint8, uint8, uint8) {
	// ignoring the alpha for now!
	red := uint8(r >> 8)
	green := uint8(g >> 8)
	blue := uint8(b >> 8)
	return red, green, blue
}
