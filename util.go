package minime

import (
	"image"
	"image/color"
)

// findSuitablePixel returns the most suitable pixel for the provided coordinates. It first checks if there is a pixel
// at the overlay coordinates, if not it will return the pixel at the normal coordinates.
func findSuitablePixel(img image.Image, overlayX, overlayY, normalX, normalY int) color.Color {
	r, g, b, a := img.At(overlayX, overlayY).RGBA()
	if a == 0 {
		r, g, b, a = img.At(normalX, normalY).RGBA()
	}
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}
}
