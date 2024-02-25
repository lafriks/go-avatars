package main

import (
	"image"
	"image/draw"
	"image/png"
	"os"
	"time"

	"github.com/lafriks/go-avatars"
)

func main() {
	size := 80
	img := image.NewRGBA(image.Rect(0, 0, size*10, size*4))

	for i := 0; i < 40; i++ {
		a, err := avatars.Generate(time.Now().Format(time.RFC3339Nano))
		if err != nil {
			panic(err)
		}
		av, err := a.Image(avatars.RenderSize(size))
		if err != nil {
			panic(err)
		}
		pos := image.Rect((i-int(i/10)*10)*size, int(i/10)*size, ((i-int(i/10)*10)+1)*size, (int(i/10)+1)*size)
		draw.Draw(img, pos, av, image.Point{}, draw.Over)
	}

	f, err := os.Create("sample.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		panic(err)
	}
}
