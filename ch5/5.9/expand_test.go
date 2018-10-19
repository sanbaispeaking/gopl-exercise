package main

import (
	"testing"
)

func TestIndexWordEnding(t *testing.T) {
	scenarios := []struct {
		s    string
		want int
	}{
		{"", -1},
		{" fuck damn", 0},
		{"abc	", 2},
		{"shit", 3},
	}
	for _, scen := range scenarios {
		if got := indexWordEnding(scen.s); got != scen.want {
			t.Errorf("indexWordEnding(%q) returned %d, want %d", scen.s, got, scen.want)
		}
	}
}

func TestExpand(t *testing.T) {
	scenarios := []struct {
		s    string
		want string
	}{
		{"", ""},
		{"$fuckdamnit", "**********"},
		{"\n	 $ fuck $shit\t", "\n	  fuck ****\t"},
		{"$shit$fuck$damn$it", "*****************"},
	}
	for _, scen := range scenarios {
		if got := expand(scen.s, beep); got != scen.want {
			t.Errorf("expand(%q) returned %q, want %q", scen.s, got, scen.want)
		}
	}
}
