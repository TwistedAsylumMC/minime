package main

import (
	_ "embed"
	"fmt"
	"github.com/TwistedAsylumMC/minime"
	"image"
	"image/png"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide input and output files")
		return
	} else if len(os.Args) < 3 {
		fmt.Println("Please provide an output file")
		return
	}

	file, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0666)
	if err != nil {
		fmt.Printf("Failed to open input file: %v\n", err)
		return
	}

	src, err := png.Decode(file)
	if err != nil {
		fmt.Printf("Failed to decode input file: %v\n", err)
		return
	}
	bounds := src.Bounds()
	if !(bounds.Dx() == 64 && bounds.Dy() == 64) && !(bounds.Dx() == 128 && bounds.Dy() == 128) {
		fmt.Printf("Input file must be 64x64 or 128x128 pixels, got %dx%d\n", bounds.Dx(), bounds.Dy())
		return
	}
	res := bounds.Dx() / 64

	var dst image.Image
	if res == 1 {
		dst = minime.Skin64(src)
	} else {
		dst = minime.Skin128(src)
	}

	if err = os.MkdirAll(filepath.Dir(os.Args[2]), 0666); err != nil {
		fmt.Printf("Failed to create output file: %v\n", err)
		return
	}
	out, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Printf("Failed to create output file: %v\n", err)
		return
	}
	defer out.Close()
	if err = png.Encode(out, dst); err != nil {
		panic(err)
	}
}
