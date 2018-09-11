package main

import (
	"bytes"
	"fmt"
	"io"
)

type countingWriter struct {
	counter int64
	w       io.Writer
}

func (c *countingWriter) Write(p []byte) (int, error) {
	n, err := c.w.Write(p)
	c.counter += int64(n)
	return n, err
}

// CountingWriter wraps a writer and makes record of the number of bytes written to
// the new writer
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var cw countingWriter
	cw.w = w
	return &cw, &cw.counter
}

func main() {
	var buf bytes.Buffer
	w, c := CountingWriter(&buf)
	w.Write([]byte("all along"))
	fmt.Printf("%d bytes wrote\n", *c)
	w.Write([]byte("the watch tower"))
	fmt.Printf("%d bytes wrote\n", *c)
}
