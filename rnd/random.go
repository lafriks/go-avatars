package rnd

import (
	"hash/fnv"
	"math/rand"
)

type Random struct {
	r *rand.Rand
}

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

func (r *Random) Chance(n int) bool {
	return r.r.Intn(100) <= n
}

func (r *Random) Pick(items []string) string {
	return items[r.r.Intn(len(items))]
}
