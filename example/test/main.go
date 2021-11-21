package main

import (
	"io/fs"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/lafriks/go-avatars"
)

func main() {
	var err error
	size := 256
	if len(os.Args) > 1 {
		size, err = strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
	}
	seed := time.Now().Format(time.RFC3339Nano)

	svg, err := avatars.SVG(seed)
	if err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile("test.svg", svg, fs.FileMode(0644)); err != nil {
		panic(err)
	}
	png, err := avatars.PNG(seed, avatars.RenderSize(size))
	if err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile("test.png", png, fs.FileMode(0644)); err != nil {
		panic(err)
	}
}
