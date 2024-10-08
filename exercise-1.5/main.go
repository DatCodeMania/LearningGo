package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.RGBA{0x00, 0xFF, 0x00, 0xFF}, color.Black}

const (
	greenIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	// Create a new file to save the GIF animation
	f, err := os.Create("lissajous.gif")
	if err != nil {
		fmt.Printf("Error creating image file: %v", err)
		os.Exit(1)
	}
	defer f.Close()

	lissajous(f)
}

func lissajous(out *os.File) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	// Encode and write to the output file
	if err := gif.EncodeAll(out, &anim); err != nil {
		fmt.Printf("Error encoding GIF: %v", err)
		os.Exit(1)
	}
}
