package main

import (
	"bufio"
	"fmt"
)

// WordCounter counts words
type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	count := 0
	for ; len(p) > 0; count++ {
		advance, _, err := bufio.ScanWords(p, true)
		if err != nil {
			return count, err
		}
		p = p[advance:]
	}
	*c += WordCounter(count)
	return count, nil
}

// LineCounter counts input lines
type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	count := 0
	for ; len(p) > 0; count++ {
		advance, _, err := bufio.ScanLines(p, true)
		if err != nil {
			return count, err
		}
		p = p[advance:]
	}
	*c += LineCounter(count)
	return count, nil
}

func main() {
	input := []byte("There must be some way outta here\n said joker to the thief")
	var wc WordCounter
	var lc LineCounter
	words, _ := wc.Write(input)
	lines, _ := lc.Write(input)
	fmt.Printf("%d words, %d lines\n", words, lines)

}
