package popcount

import "testing"

func BenchmarkShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShiftCount(^uint64(111))
	}
}

func benchmarkLookup(b *testing.B, size int) {
	b.N = size
	for i := 0; i < b.N; i++ {
		LookupCount(^uint64(111))
	}
}

func BenchmarkLookup1000(b *testing.B) {
	benchmarkLookup(b, 1000)
}
func BenchmarkLookup100000(b *testing.B) {
	benchmarkLookup(b, 100000)
}
func BenchmarkLookup1000000(b *testing.B) {
	benchmarkLookup(b, 1000000)
}

func BenchmarkClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClearCount(^uint64(111))
	}
}
