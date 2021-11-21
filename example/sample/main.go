package main

import (
	"bytes"
	"image"
	"image/draw"
	"image/png"
	"os"
	"time"

	"github.com/lafriks/go-avatars"
)

func main() {
	size := 80
	img := image.NewRGBA(image.Rect(0, 0, size*6, size*2))

	for i := 0; i < 12; i++ {
		seed := time.Now().Format(time.RFC3339Nano)
		buf, err := avatars.PNG(seed, avatars.RenderSize(size))
		if err != nil {
			panic(err)
		}
		av, err := png.Decode(bytes.NewReader(buf))
		if err != nil {
			panic(err)
		}
		pos := image.Rect(i*size, 0, (i+1)*size, size)
		if i >= 6 {
			pos = image.Rect((i-6)*size, size, (i-5)*size, size*2)
		}
		draw.Draw(img, pos, av, image.Point{}, draw.Over)
	}

	f, err := os.Create("sample.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}
