package gomeme

import (
	"image"
	"image/draw"
	"io"

	"github.com/chai2010/webp"
)

// WebP comprises of all things needed to create a new
// meme from a Image
type WebP struct {
	Image image.Image
}

// Bounds return the bounds of the first frame
func (i WebP) Bounds() image.Rectangle {
	return i.Image.Bounds()
}

// Write Image to writer
func (i WebP) Write(textImage *image.RGBA, w io.Writer) error {
	output := image.NewRGBA(i.Bounds())
	draw.Draw(output, i.Bounds(), i.Image, image.ZP, draw.Src)
	draw.DrawMask(output, i.Bounds(), textImage, image.ZP, textImage, image.ZP, draw.Over)

	return webp.Encode(w, output, &webp.Options{Lossless: true})
}
