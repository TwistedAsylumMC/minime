package minime

import (
	"image"
	"image/color"
	"math"
)

// Skin128 returns an avatar generated from a skin that is 128x128 pixels. The skin must also be in the Minecraft format
// otherwise the generated result may not look correct. If the skin is slim (using the Alex model), slim should be set
// to true for the best results.
func Skin128(src image.Image, slim bool) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, 18, 28))
	// Head
	for x := 0; x < 18; x++ {
		for y := 0; y < 18; y++ {
			if x == 0 || x == 17 || y == 0 || y == 17 {
				dst.Set(x, y, color.Black)
				continue
			}
			dst.Set(x, y, findSuitablePixel(src, x+79, y+15, x+15, y+15))
		}
	}

	// Arms
	for x := 0; x < 12; x++ {
		for y := 0; y < 6; y++ {
			if x == 0 || x == 11 || y == 0 || y == 5 {
				dst.Set(x+3, y+17, color.Black)
				continue
			} else if x > 2 && x < 9 {
				continue
			}
			xOff, yOff := (1-(x%2))*4, (y-1)*8
			if slim {
				xOff--
			}
			if x < 3 {
				dst.Set(x+3, y+17, findSuitablePixel(src, 90+xOff, 73+yOff, 90+xOff, 41+yOff))
			} else {
				dst.Set(x+3, y+17, findSuitablePixel(src, 106+xOff, 105+yOff, 74+xOff, 105+yOff))
			}
		}
	}

	// Body
	for x := 0; x < 6; x++ {
		for y := 0; y < 11; y++ {
			if x == 0 || x == 5 || y == 0 || y == 7 || y == 10 {
				dst.Set(x+6, y+17, color.Black)
				continue
			} else if y > 7 {
				xOff, yOff := (int(math.Ceil(float64(x)/2))-1)*4, (y-8)*9
				if x < 3 {
					dst.Set(x+6, y+17, findSuitablePixel(src, 10+xOff, 83+yOff, 10+xOff, 51+yOff))
				} else {
					dst.Set(x+6, y+17, findSuitablePixel(src, 10+xOff, 115+yOff, 42+xOff, 115+yOff))
				}
				continue
			}
			xOff, yOff := (x-1)*4, (y-1)*4
			if x > 2 {
				xOff--
			}
			if y > 3 {
				yOff++
			}
			dst.Set(x+6, y+17, findSuitablePixel(src, 42+xOff, 73+yOff, 42+xOff, 41+yOff))
		}
	}
	return dst
}
