package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/net/html"
)

type StrinReader struct {
	s string
	i int64
}

func (sr *StrinReader) Read(p []byte) (n int, err error) {
	if sr.i >= int64(len(sr.s)) {
		return 0, io.EOF
	}

	n = copy(p, sr.s[sr.i:])
	sr.i += int64(n)
	return
}

func NewReader(text string) *StrinReader {
	return &StrinReader{text, 0}
}

//!+
func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	// fmt.Printf(string(input))
	page := NewReader(string(input))

	doc, err := html.Parse(page)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	stack := make(map[string]int)
	outline(stack, doc)
}

func outline(stack map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		stack[n.Data]++
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
