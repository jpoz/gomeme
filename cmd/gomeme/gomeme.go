package main

import (
	"flag"
	"fmt"
	"image/gif"
	"os"

	"github.com/jpoz/gomeme"
)

// Version of gomeme set by ldflags
var Version string

// BuildTime is when gomeme was built
var BuildTime string

func main() {
	var verbose bool
	config := gomeme.NewConfig()

	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] input.gif output.gif\nv%s (%s)\n\n",
			os.Args[0],
			Version,
			BuildTime,
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

	in, err := os.Open(flag.Arg(0))
	if err != nil {
		fail("Could not open input", err)
	}

	g, err := gif.DecodeAll(in)
	if err != nil {
		fail("Failed to decode gif", err)
	}

	meme := &gomeme.Meme{
		Config:   config,
		Memeable: gomeme.GIF{g},
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
