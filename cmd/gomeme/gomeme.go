package main

import (
	"flag"
	"fmt"
	"image/gif"
	"os"

	"github.com/jpoz/gomeme"
)

func main() {
	var verbose bool
	meme, err := gomeme.NewMeme()
	check(err)

	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] input.gif output.gif\n\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.BoolVar(&verbose, "v", false, "Displays more information.")
	flag.Float64Var(&meme.FontSize, "fs", gomeme.DefaultFontSize, "Font size of the text")
	flag.IntVar(&meme.FontStrokeSize, "ss", gomeme.DefaultStrokeSize, "Stroke size around the text")
	flag.IntVar(&meme.Margin, "m", gomeme.DefaultMargin, "Margin around the text")
	flag.StringVar(&meme.BottomText, "b", "", "Bottom text of the meme.")
	flag.StringVar(&meme.FontPath, "f", "", "TrueType font path. (optional)")
	flag.StringVar(&meme.TopText, "t", "", "Top text of the meme.")

	flag.Parse()

	if verbose {
		fmt.Println(meme)
	}

	if flag.NArg() < 2 {
		fail("Need input and output path\n\n", nil)
	}

	in, err := os.Open(flag.Arg(0))
	if err != nil {
		fail("Could not open input", err)
	}

	meme.GIF, err = gif.DecodeAll(in)
	if err != nil {
		fail("Failed to decode gif", err)
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
