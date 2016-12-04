package gomeme

import "io"

// Meme is the configuration and the "memable" image or gif
type Meme struct {
	Config   *Config
	Memeable Memeable
}

// Write Meme to writer
func (m Meme) Write(w io.Writer) error {
	textImage, err := m.Config.TextImage(m.Memeable.Bounds())
	if err != nil {
		return err
	}

	return m.Memeable.Write(textImage, w)
}
