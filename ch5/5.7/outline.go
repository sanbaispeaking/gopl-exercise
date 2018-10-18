// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}

}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+call
	forEachNode(doc, startElement, endElement)
	//!-call

	return nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node, hasChild bool)) {
	var hasChild bool
	if n.FirstChild != nil {
		hasChild = true
	}

	if pre != nil {
		pre(n, hasChild)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n, hasChild)
	}
}

//!-forEachNode

//!+startend
var depth int

func startElement(n *html.Node, hasChild bool) {
	if n.Type == html.ElementNode {
		var tagEndding = "/>"
		if hasChild {
			tagEndding = ">"
		}

		if len(n.Attr) == 0 {
			fmt.Printf("%*s<%s%s\n", depth*2, "", n.Data, tagEndding)
			depth++
			return
		}

		attrs := make([]string, 0, len(n.Attr)+1)
		attrs = append(attrs, n.Data)
		for _, a := range n.Attr {
			attrs = append(attrs, fmt.Sprintf("%s=%q", a.Key, a.Val))
		}
		fmt.Printf("%*s<%s%s\n", depth*2, "", strings.Join(attrs, " "), tagEndding)
		depth++
	} else {
		for _, line := range strings.Split(n.Data, "\n") {
			if l := strings.TrimSpace(line); l != "" {
				fmt.Printf("%*s%s\n", depth*2, "", l)
			}
		}
	}
}

func endElement(n *html.Node, hasChild bool) {
	if n.Type == html.ElementNode {
		depth--
		if hasChild {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}

//!-startend
