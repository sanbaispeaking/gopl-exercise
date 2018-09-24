package charcount

import "testing"

func TestCharcount(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"", 0},
		{"abc", 3},
	}

	for _, test := range tests {
		if got := Charcount(test.input); got != test.want {
			t.Errorf("Charcount(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}
