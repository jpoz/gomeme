package gomeme

import (
	"image"
	"io"
)

// Memeable is takes a image of the text overlays it with its image
// and outputs it to the Writer
type Memeable interface {
	Bounds() image.Rectangle
	Write(textImage *image.RGBA, w io.Writer) error
}
