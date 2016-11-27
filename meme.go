package gomeme

import (
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"io"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"

	"github.com/disintegration/imaging"
	"github.com/golang/freetype/truetype"
	"github.com/jpoz/dilation"
)

// DefaultFontSize is the default size of the font
const DefaultFontSize = 42

// DefaultMargin is the default distance between the text and the top
// and bottem
const DefaultMargin = 10

// DefaultDPI is the fonts DPI
const DefaultDPI = 42

// DefaultStrokeSize default width of the stroke
const DefaultStrokeSize = 4

// Meme comprises of all things needed to create a new
// meme from a gif
type Meme struct {
	Font            *truetype.Font
	FontSize        float64
	FontDPI         float64
	FontColor       image.Image
	FontStrokeSize  int
	FontStrokeColor color.Color

	Margin int

	TopText    string
	BottomText string

	GIF    *gif.GIF
	Bounds image.Rectangle
}

// NewMeme takes a reader and builds a new Meme with
// default configureations.
func NewMeme() (*Meme, error) {
	meme := &Meme{
		FontColor:       image.Black,
		FontDPI:         DefaultDPI,
		FontSize:        DefaultFontSize,
		FontStrokeColor: color.White,
		FontStrokeSize:  DefaultStrokeSize,
		Margin:          DefaultMargin,
	}
	var err error

	var fontData []byte
	fontData, err = Asset("Hack-Bold.ttf")
	if err != nil {
		return meme, err
	}

	meme.Font, err = truetype.Parse(fontData)
	if err != nil {
		return meme, err
	}

	return meme, nil
}

// Write GIF to writer
func (m Meme) Write(w io.Writer) error {
	m.build()

	return gif.EncodeAll(w, m.GIF)
}

// Build will take the current settings of the Meme and updates the GIF
func (m *Meme) build() {
	bounds := m.GIF.Image[0].Bounds()
	textImage := m.textImage()

	// Write on gif
	for _, img := range m.GIF.Image {
		newFrame := imaging.Overlay(img, textImage, image.Pt(0, 0), 1.0)
		bnds := img.Bounds()
		draw.Draw(img, bnds, newFrame, bounds.Min, draw.Src)
	}
}

func (m *Meme) textImage() *image.RGBA {
	bounds := m.GIF.Image[0].Bounds()
	textImage := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	f := truetype.NewFace(m.Font, &truetype.Options{
		Size:    m.FontSize,
		DPI:     m.FontDPI,
		Hinting: font.HintingNone,
	})
	d := &font.Drawer{
		Dst:  textImage,
		Src:  m.FontColor,
		Face: f,
	}

	if m.TopText != "" {
		// Compute the top text position
		metrics := f.Metrics()
		ascent := metrics.Ascent.Floor()
		height := metrics.Height.Floor()
		y := m.Margin + ascent + (ascent - height)
		x := (fixed.I(bounds.Dx()) - d.MeasureString(m.TopText)) / 2
		topDot := fixed.Point26_6{
			X: x,
			Y: fixed.I(y),
		}

		// Draw the top text
		d.Dot = topDot
		d.DrawString(m.TopText)
	}

	if m.BottomText != "" {
		// Compute the bottom text position
		y := bounds.Dy() - m.Margin
		x := (fixed.I(bounds.Dx()) - d.MeasureString(m.BottomText)) / 2
		bottomDot := fixed.Point26_6{
			X: x,
			Y: fixed.I(y),
		}

		// Draw the bottom text
		d.Dot = bottomDot
		d.DrawString(m.BottomText)
	}

	dilation.Dialate(textImage, dilation.DialateConfig{
		Stroke:      m.FontStrokeSize,
		StrokeColor: m.FontStrokeColor,
	})

	return textImage
}
