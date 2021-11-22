package bottts

import (
	"fmt"
	"strings"

	"github.com/lafriks/go-avatars/rnd"
)

const (
	sidesChance   = 50
	topChance     = 50
	textureChance = 50
)

func group(r *rnd.Random, content string, chance, x, y int) string {
	if !r.Chance(chance) {
		return ""
	}
	return fmt.Sprintf(`<g transform="translate(%d, %d)">%s</g>`, x, y, content)
}

// Generate generates a bottts avatar.
func Generate(r *rnd.Random) (map[string]string, string) {
	var texture string
	if r.Chance(textureChance) {
		texture = textures[r.Pick(textureNames)]
	}

	primaryColor, secondaryColor := r.PickColors(600, 400)

	var buf strings.Builder
	buf.WriteString(group(r, sides[r.Pick(sideNames)](secondaryColor), sidesChance, 0, 66))
	buf.WriteString(group(r, tops[r.Pick(topNames)](secondaryColor), topChance, 41, 0))
	buf.WriteString(group(r, face[r.Pick(faceNames)](primaryColor, texture), 100, 25, 44))
	buf.WriteString(group(r, mouth[r.Pick(mouthNames)], 100, 52, 124))
	buf.WriteString(group(r, eyes[r.Pick(eyeNames)], 100, 38, 76))

	return map[string]string{
		"viewBox": "0 0 180 180",
	}, buf.String()
}
