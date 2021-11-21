package rnd

import (
	"hash/fnv"
	"math/rand"
)

// Random is a random number generator.
type Random struct {
	r *rand.Rand
}

// NewRandom returns a new random number generator using specified seed.
func NewRandom(seed string) *Random {
	// Use hash of string to seed random number generator
	h := fnv.New64a()
	h.Write([]byte(seed))
	i := h.Sum64()
	s := int64(i & 0x7fffffffffffffff)
	if i&0x8000000000000000 != 0 {
		s = -s
	}
	rnd := &Random{r: rand.New(rand.NewSource(s))}
	rnd.r.Seed(s)
	return rnd
}

// Chance returns true if the specified probability is met.
func (r *Random) Chance(n int) bool {
	return r.r.Intn(100) <= n
}

// Pick returns a random element from the specified collection.
func (r *Random) Pick(items []string) string {
	return items[r.r.Intn(len(items))]
}

// PickColors returns a random color with two differnet shades.
func (r *Random) PickColors(primary, secondary int) (string, string) {
	color := r.Pick(colorNames)
	return collection[color][primary], collection[color][secondary]
}
