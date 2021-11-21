package main

import (
	"io/fs"
	"io/ioutil"
	"time"

	"github.com/lafriks/go-avatars"
)

func main() {
	seed := time.Now().Format(time.RFC3339Nano)

	svg, err := avatars.SVG(seed)
	if err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile("test.svg", svg, fs.FileMode(0644)); err != nil {
		panic(err)
	}
	png, err := avatars.PNG(seed)
	if err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile("test.png", png, fs.FileMode(0644)); err != nil {
		panic(err)
	}
}
