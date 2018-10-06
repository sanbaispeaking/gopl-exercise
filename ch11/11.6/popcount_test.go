package popcount

import "testing"

func BenchmarkShift0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShiftCount(0)
	}
}

func BenchmarkLoopup0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LookupCount(0)
	}
}

func BenchmarkClear0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClearCount(0)
	}
}

func BenchmarkShift64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShiftCount(^uint64(0))
	}
}

func BenchmarkLoopup64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LookupCount(^uint64(0))
	}
}

func BenchmarkClear64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClearCount(^uint64(0))
	}
}
