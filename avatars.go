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

func SVG(name string) ([]byte, error) {
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

func PNG(name string) ([]byte, error) {
	avatar, err := SVG(name)
	if err != nil {
		return nil, err
	}
	s, err := svg.Parse(bytes.NewReader(avatar), svg.IgnoreErrorMode)
	if err != nil {
		return nil, err
	}

	gc := gg.NewContext(256, 256)
	rendr_gg.Draw(gc, s, renderer.Target(0, 0, 256, 256))

	buf := bytes.NewBuffer(nil)
	if err := png.Encode(buf, gc.Image()); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
