package main

import (
	"crypto/rand"
	"image"
	"image/color"
	"image/png"
	"os"

	"go.lukeharris.dev/r30"
)

func main() {
	size := 8

	data := make([]byte, size)
	rand.Read(data)

	w, h := size*8, size*8

	upLeft := image.Point{0, 0}
	lowRight := image.Point{w, h}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if (data[x/8]>>(7-x%8))&0b1 == 1 {
				img.Set(x, y, color.Black)
			} else {
				img.Set(x, y, color.White)
			}
		}
		data = r30.Step(data)
	}

	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
