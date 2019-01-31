package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

type StrinReader string

func (sr StrinReader) Read(p []byte) (n int, err error) {
	nb := copy([]byte(sr), p)
	if nb == len(sr) {
		return nb, io.EOF
	}
	return nb, errors.New("read incomplete")
}

func NewReader(text string) io.Reader {
	sr := StrinReader(text)
	return &sr
}

//!+
func main() {

	page := NewReader(`
	<html>
<head><title>301 Moved Permanently</title></head>
<body bgcolor="white">
<center><h1>301 Moved Permanently</h1></center>
<hr><center>nginx</center>
</body>
</html>`)

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
