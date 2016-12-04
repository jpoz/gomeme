package gomeme

import (
	"image"
	"image/color"
	"io/ioutil"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"

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

// Config hold all the configuration needed to make a new meme.
// It Also produces the text image to be overlayed
type Config struct {
	FontPath        string
	FontSize        float64
	FontDPI         float64
	FontColor       image.Image
	FontStrokeSize  int
	FontStrokeColor color.Color

	Margin int

	TopText    string
	BottomText string
}

// NewConfig builds a default configuration.
func NewConfig() *Config {
	meme := &Config{
		FontColor:       image.White,
		FontDPI:         DefaultDPI,
		FontSize:        DefaultFontSize,
		FontStrokeColor: color.Black,
		FontStrokeSize:  DefaultStrokeSize,
		Margin:          DefaultMargin,
	}

	return meme
}

func (c *Config) loadFont() (*truetype.Font, error) {
	var err error
	var fontData []byte

	if c.FontPath == "" {
		fontData, err = Asset("inpact.ttf")
		if err != nil {
			return nil, err
		}
	} else {
		fontData, err = ioutil.ReadFile(c.FontPath)
		if err != nil {
			return nil, err
		}
	}

	return truetype.Parse(fontData)
}

// TextImage produces a image of the bottom and top text
func (c *Config) TextImage(bounds image.Rectangle) (*image.RGBA, error) {
	fnt, err := c.loadFont()
	if err != nil {
		return nil, err
	}

	textImage := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	f := truetype.NewFace(fnt, &truetype.Options{
		Size:    c.FontSize,
		DPI:     c.FontDPI,
		Hinting: font.HintingNone,
	})
	d := &font.Drawer{
		Dst:  textImage,
		Src:  c.FontColor,
		Face: f,
	}

	// Not sure if these are the best metrics for the margin calculations
	metrics := f.Metrics()
	ascent := metrics.Ascent.Ceil()
	descent := metrics.Descent.Ceil()
	// Maybe height should be used?
	//height := metrics.Height.Ceil()

	if c.TopText != "" {
		// Compute the top text position
		y := c.Margin + ascent
		x := (fixed.I(bounds.Dx()) - d.MeasureString(c.TopText)) / 2
		topDot := fixed.Point26_6{
			X: x,
			Y: fixed.I(y),
		}

		// Draw the top text
		d.Dot = topDot
		d.DrawString(c.TopText)
	}

	if c.BottomText != "" {
		// Compute the bottom text position
		y := bounds.Dy() - c.Margin - descent
		x := (fixed.I(bounds.Dx()) - d.MeasureString(c.BottomText)) / 2
		bottomDot := fixed.Point26_6{
			X: x,
			Y: fixed.I(y),
		}

		// Draw the bottom text
		d.Dot = bottomDot
		d.DrawString(c.BottomText)
	}

	// Dialate aka give text a stroke
	dilation.Dialate(textImage, dilation.DialateConfig{
		Stroke:      c.FontStrokeSize,
		StrokeColor: c.FontStrokeColor,
	})

	return textImage, nil
}
