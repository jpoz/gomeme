package gomeme

import (
	"image"
	"image/draw"
	"image/gif"
	"io"
	"sync"
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
	var g2 gif.GIF
	g2 = *g.GIF
	var wg sync.WaitGroup
	wg.Add(len(g2.Image))
	frames := make([]*image.Paletted, len(g2.Image))
	for i, img := range g2.Image {
		go func(i int, img *image.Paletted) {
			img2 := image.NewPaletted(img.Bounds(), img.Palette)
			draw.Draw(img2, img2.Bounds(), img, image.ZP, draw.Src)
			draw.DrawMask(img2, textImage.Bounds(), textImage, image.ZP, textImage, image.ZP, draw.Over)
			frames[i] = img2
			wg.Done()
		}(i, img)
	}
	wg.Wait()
	g2.Image = frames
	return gif.EncodeAll(w, &g2)
}
