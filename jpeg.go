package gomeme

import (
	"image"
	"image/draw"
	"image/jpeg"
	"io"
)

// JPEG comprises of all things needed to create a new
// meme from a Image
type JPEG struct {
	Image image.Image
}

// Bounds return the bounds of the first frame
func (i JPEG) Bounds() image.Rectangle {
	return i.Image.Bounds()
}

// Write Image to writer
func (i JPEG) Write(textImage *image.RGBA, w io.Writer) error {
	output := image.NewRGBA(i.Bounds())
	draw.Draw(output, i.Bounds(), i.Image, image.ZP, draw.Src)
	draw.DrawMask(output, i.Bounds(), textImage, image.ZP, textImage, image.ZP, draw.Over)

	return jpeg.Encode(w, output, nil)
}
