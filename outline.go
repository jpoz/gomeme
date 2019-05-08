package gomeme

import (
	"image"
	"image/color"
	"image/draw"
	"math"
)

// Outline puts an outline of given width around all non-transparent pixels in the image.
// This is not a particularly clever or efficient implementation.
func outline(img *image.RGBA, size int, color color.Color) *image.RGBA {
	width := img.Bounds().Size().X
	height := img.Bounds().Size().Y
	circle := drawCircle(size, color)
	nimg := image.NewRGBA(image.Rect(0, 0, width, height))
	zero := image.Point{}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if img.RGBAAt(x, y).A > 0 {
				rect := image.Rectangle{
					Min: image.Point{X: x - size, Y: y - size},
					Max: image.Point{X: x + size, Y: y + size},
				}
				draw.Over.Draw(nimg, rect, circle, zero)
			}
		}
	}
	draw.Over.Draw(nimg, img.Bounds(), img, zero)
	return nimg
}

// drawCircle returns an image containing a circle of the given size.
func drawCircle(size int, clr color.Color) *image.RGBA {
	sz := size*2 + 1
	img := image.NewRGBA(image.Rect(0, 0, sz, sz))
	r, g, b, _ := clr.RGBA()
	for x := -size + 1; x < size; x++ {
		for y := -size + 1; y < size; y++ {
			xr := math.Abs(float64(x))
			yr := math.Abs(float64(y))
			dist := math.Sqrt(xr*xr + yr*yr)
			alpha := math.Min(math.Max(float64(size)-dist, 0.0), 1.0)
			c := color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(alpha * 255)}
			img.Set(x+size, y+size, c)
		}
	}
	return img
}
