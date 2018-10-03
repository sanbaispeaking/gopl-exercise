package intset

import (
	"math/rand"
	"testing"
	"time"
)

func randomSet(rng *rand.Rand) *IntSet {
	n := rng.Intn(1000)
	words := make([]uint, n)
	for ; n > 0; n-- {
		words = append(words, uint(rng.Uint64()))
	}
	return &IntSet{words}
}

func BenchmarkUnionWith(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	b.Logf("Rand seed :%d", seed)
	rng := rand.New(rand.NewSource(seed))
	l := randomSet(rng)
	r := randomSet(rng)

	for i := 0; i < b.N; i++ {
		l.UnionWith(r)
	}
}

func BenchmarkIntersect(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	b.Logf("Rand seed :%d", seed)
	rng := rand.New(rand.NewSource(seed))
	l := randomSet(rng)
	r := randomSet(rng)

	for i := 0; i < b.N; i++ {
		l.IntersectWith(r)
	}
}
