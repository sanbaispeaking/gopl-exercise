package split

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	scenarios := []struct {
		s    string
		sep  string
		want int
	}{
		{"a:b:c", ":", 3},
		{":b:c", ":", 3},
		{"", ":", 1},
		{"abc", ":", 1},
	}
	for _, scen := range scenarios {
		words := strings.Split(scen.s, scen.sep)
		if got := len(words); got != scen.want {
			t.Errorf("Split(%q, %q) returned %d words, want %d",
				scen.s, scen.sep, got, scen.want)
		}
	}
}
