package word

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
	"unicode"
)

func randomNoisyPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x200)) // random rune up to \u99
		runes[i] = r
		r1 := r
		if unicode.IsLetter(r) && unicode.IsLower(r) {
			r = unicode.ToUpper(r)
			if unicode.ToLower(r) != r1 {
				fmt.Printf("cap? %c %c\n", r1, r)
			}
		}
		runes[n-1-i] = r
	}
	return "?" + string(runes) + "!"
}

func TestRandomNonPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomNoisyPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

func TestRandomNoisyPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	n := 0
	for i := 0; i < 1000; i++ {
		p := randomNoisyPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsNoisyPalindrome(%q) = false", p)
			n++
		}
	}
	fmt.Fprintf(os.Stderr, "fail = %d\n", n)
}
