package avatars

import (
	"bytes"
	"html"
	"image"
	"image/png"

	"github.com/lafriks/go-avatars/bottts"
	"github.com/lafriks/go-avatars/rnd"

	"github.com/fogleman/gg"
	"github.com/lafriks/go-svg"
	"github.com/lafriks/go-svg/renderer"
	rendr_gg "github.com/lafriks/go-svg/renderer/gg"
)

// Avatar that was generated using specified seed.
type Avatar struct {
	r    *rnd.Random
	name string
	svg  []byte
}

// Generate an avatar using the given name as seed.
func Generate(name string, opts ...Option) (*Avatar, error) {
	r := rnd.NewRandom(name)

	attrs, body := bottts.Generate(r)

	res := bytes.Buffer{}
	res.WriteString("<svg")
	for name, val := range attrs {
		res.WriteByte(' ')
		res.WriteString(name)
		res.WriteString(`="`)
		res.WriteString(html.EscapeString(val))
		res.WriteByte('"')
	}
	res.WriteString(">\n")
	res.WriteString(body)
	res.WriteString("\n</svg>")
	return &Avatar{
		name: name,
		r:    r,
		svg:  res.Bytes(),
	}, nil
}

// Name of the avatar.
func (a *Avatar) Name() string {
	return a.name
}

// SVG content of the avatar.
func (a *Avatar) SVG() []byte {
	return a.svg
}

// Image rendered from the SVG.
func (a *Avatar) Image(opts ...RenderOption) (image.Image, error) {
	s, err := svg.Parse(bytes.NewReader(a.svg), svg.IgnoreErrorMode)
	if err != nil {
		return nil, err
	}

	opt := renderOptions(opts...)

	gc := gg.NewContext(opt.Size, opt.Size)
	rendr_gg.Draw(gc, s, renderer.Target(
		float64(opt.Padding),
		float64(opt.Padding),
		float64(opt.Size-opt.Padding*2),
		float64(opt.Size-opt.Padding*2),
	))

	return gc.Image(), nil
}

// PNG encoded image.
func (a *Avatar) PNG(opts ...RenderOption) ([]byte, error) {
	img, err := a.Image(opts...)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(nil)
	if err := png.Encode(buf, img); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
