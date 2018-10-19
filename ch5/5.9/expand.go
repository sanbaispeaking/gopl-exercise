package main

import (
	"fmt"
	"strings"
	"unicode"
)

// expand replace each substring `$foo` within s by `f(foo)`
func expand(s string, f func(string) string) string {
	l := len(s)
	if f == nil || l == 0 {
		return s
	}

	// consume the first template
	tmplStart := strings.IndexRune(s, '$')
	if tmplStart != -1 {
		tmplEnd := l - 1
		//proceed to template ending
		if tmplStart <= l-1 {
			ending := indexWordEnding(s[tmplStart+1:])
			if ending > 0 {
				tmplEnd = tmplStart + ending + 1
			} else if ending == 0 {
				tmplEnd = tmplStart
			}
		}

		//pass on the rest of string
		if tmplEnd <= l-1 {
			s = s[:tmplEnd+1] + expand(s[tmplEnd+1:], f)
		}

		template := s[tmplStart : tmplEnd+1]
		if len(template) > 2 {
			s = strings.Replace(s, template, f(template[1:]), 1)
		} else {
			s = strings.Replace(s, template, f(""), 1)
		}
	}
	return s
}

func indexWordEnding(s string) int {
	if len(s) == 0 {
		return -1
	}

	var prev rune
	for i, r := range s {
		if !unicode.IsSpace(r) {
			prev = r
		} else {
			if i == 0 {
				return 0
			}
			if unicode.IsGraphic(prev) && !unicode.IsSpace(prev) {
				return i - 1
			}
		}
	}

	return len(s) - 1
}

// beep returns a string of repetitive '*' as long as the original input
func beep(in string) string {
	var b strings.Builder
	for i := len(in); i > 0; i-- {
		b.WriteRune('*')
	}

	return b.String()
}

func main() {
	var in = "\n	 $ fuck $shit\t"
	fmt.Println(expand(in, beep))
}
