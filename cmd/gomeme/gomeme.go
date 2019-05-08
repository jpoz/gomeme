package main

import (
	"bytes"
	"flag"
	"fmt"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/peterebden/gomeme"
)

// Version of gomeme
var Version = "1.1.0"

func main() {
	var verbose bool
	config := gomeme.NewConfig()

	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] input.gif output.gif\nv%s\n\n",
			os.Args[0],
			Version,
		)
		flag.PrintDefaults()
	}

	flag.BoolVar(&verbose, "v", false, "Displays more information.")
	flag.Float64Var(&config.FontSize, "fs", gomeme.DefaultFontSize, "Font size of the text")
	flag.IntVar(&config.FontStrokeSize, "ss", gomeme.DefaultStrokeSize, "Stroke size around the text")
	flag.IntVar(&config.Margin, "m", gomeme.DefaultMargin, "Margin around the text")
	flag.StringVar(&config.BottomText, "b", "", "Bottom text of the config.")
	flag.StringVar(&config.FontPath, "f", "", "TrueType font path. (optional)")
	flag.StringVar(&config.TopText, "t", "", "Top text of the config.")

	flag.Parse()

	if verbose {
		// TODO make this output better
		fmt.Println(config)
	}

	if flag.NArg() < 2 {
		fail("Need input and output path\n\n", nil)
	}

	in, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		fail("Could not open input", err)
	}

	meme := &gomeme.Meme{
		Config: config,
	}

	contentType := http.DetectContentType(in)
	buff := bytes.NewBuffer(in)

	switch contentType {
	case "image/gif":
		g, err := gif.DecodeAll(buff)
		if err != nil {
			fail("Failed to decode gif", err)
		}
		meme.Memeable = gomeme.GIF{g}
	case "image/jpeg":
		j, err := jpeg.Decode(buff)
		if err != nil {
			fail("Failed to decode jpeg", err)
		}
		meme.Memeable = gomeme.JPEG{j}
	case "image/png":
		p, err := png.Decode(buff)
		if err != nil {
			fail("Failed to decode png", err)
		}
		meme.Memeable = gomeme.PNG{p}
	default:
		fail(fmt.Sprintf("No idea what todo with a %s", contentType), nil)
	}

	out, err := os.Create(flag.Arg(1))
	if err != nil {
		fail("Could not open output file", err)
	}

	check(meme.Write(out))
}

func fail(s string, e error) {
	fmt.Fprintf(os.Stderr, s)
	flag.Usage()
	check(e)
	os.Exit(1)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
