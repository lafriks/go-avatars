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

	a, err := avatars.Generate(seed)
	if err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile("test.svg", a.SVG(), fs.FileMode(0644)); err != nil {
		panic(err)
	}
	img, err := a.PNG(avatars.RenderSize(size))
	if err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile("test.png", img, fs.FileMode(0644)); err != nil {
		panic(err)
	}
}
