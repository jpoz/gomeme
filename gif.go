package gomeme

import (
	"image"
	"image/draw"
	"image/gif"
	"io"
)

// GIF comprises of all things needed to create a new
// meme from a gif
type GIF struct {
	GIF *gif.GIF
}

// Bounds return the bounds of the first frame
func (g GIF) Bounds() image.Rectangle {
	return g.GIF.Image[0].Bounds()
}

// Write GIF to writer
func (g GIF) Write(textImage *image.RGBA, w io.Writer) error {
	// TODO: Break this out on each CPU
	for _, img := range g.GIF.Image {
		draw.DrawMask(img, textImage.Bounds(), textImage, image.ZP, textImage, image.ZP, draw.Over)
	}

	return gif.EncodeAll(w, g.GIF)
}
