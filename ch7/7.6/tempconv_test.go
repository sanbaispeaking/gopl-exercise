package tempconv

import (
	"testing"
)

func TestfahrenheitFlag(t *testing.T) {
	scenarios := []struct {
		s    string
		want float64
	}{
		{"20.0F", 20.0},
		{"20.0°F", 20.0},
	}
	for _, scen := range scenarios {
		var f *fahrenheitFlag
		if f.Set(scen.s); f.Fahrenheit != Fahrenheit(scen.want) {
			t.Errorf("indexWordEnding(%q) returned %v, want %v", scen.s, f, scen.want)
		}
	}
}
