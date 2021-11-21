package avatars

import (
	"bytes"
	"html"
	"image/png"

	"github.com/lafriks/go-avatars/bottts"

	"github.com/fogleman/gg"
	"github.com/lafriks/go-svg"
	"github.com/lafriks/go-svg/renderer"
	rendr_gg "github.com/lafriks/go-svg/renderer/gg"
)

// SVG generates an SVG avatar using the given name as seed.
func SVG(name string, opts ...Option) ([]byte, error) {
	attrs, body := bottts.Generate(name)

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
	return res.Bytes(), nil
}

// PNG generates an PNG avatar using the given name as seed.
func PNG(name string, opts ...Option) ([]byte, error) {
	avatar, err := SVG(name, opts...)
	if err != nil {
		return nil, err
	}

	opt := options(opts...)

	s, err := svg.Parse(bytes.NewReader(avatar), svg.IgnoreErrorMode)
	if err != nil {
		return nil, err
	}

	gc := gg.NewContext(opt.Size, opt.Size)
	rendr_gg.Draw(gc, s, renderer.Target(
		float64(opt.Padding),
		float64(opt.Padding),
		float64(opt.Size-opt.Padding*2),
		float64(opt.Size-opt.Padding*2),
	))

	buf := bytes.NewBuffer(nil)
	if err := png.Encode(buf, gc.Image()); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
