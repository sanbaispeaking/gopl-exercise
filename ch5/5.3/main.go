package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

//!+
func main() {
	tokenizer := html.NewTokenizer(os.Stdin)

	var stack string
	extractText(&stack, tokenizer)
	fmt.Println(stack)
}

// extractText extracts content of text node
func extractText(stack *string, z *html.Tokenizer) error {
	var inScript, inStyle bool
	for {
		t := z.Next()
		switch t {
		case html.ErrorToken:
			return z.Err()
		case html.StartTagToken, html.EndTagToken:
			tn, _ := z.TagName()
			if string(tn) == "script" {
				inScript = !inScript
			} else if string(tn) == "style" {
				inStyle = !inStyle
			}
		case html.TextToken:
			if inScript || inStyle {
				break
			}
			text := string(z.Text())
			if text != "" {
				*stack += text
			}
		}
	}

}
