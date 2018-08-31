package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// URL with depth indicator recording how many hoops does it need to reach from index page
type URL struct {
	url string
	gen int
}

func crawl(u URL, cancle <-chan struct{}) []URL {
	fmt.Printf("%v of depth %d\n", u.url, u.gen)
	list, err := Extract(u.url, cancle)
	if err != nil {
		log.Print(err)
	}
	ret := make([]URL, len(list))
	for i, url := range list {
		ret[i].url = url
		ret[i].gen = u.gen + 1
	}
	return ret
}

func parser() (depth int, targets []string) {
	//Isolate `-depth`
	depthPtr := flag.Int("depth", 0, "max hoops allowed from index, valule <=0 will be ignored")

	flag.Parse()
	//Preserve other args as-is
	targets = flag.Args()
	depth = *depthPtr
	return depth, targets
}

//!+
func main() {
	depth, targets := parser()
	log.Println("Start crawling from ", targets, ", depth: ", depth)

	worklist := make(chan []URL)  // lists of URLs, may have duplicates
	unseenLinks := make(chan URL) // de-duplicated URLs
	done := make(chan struct{})   // for http request cancelation

	// Add command-line arguments to worklist.
	go func() {
		urls := make([]URL, len(targets))
		// Wrap url strings
		for i, url := range targets {
			urls[i].url = url
			urls[i].gen = 0
		}
		worklist <- urls
	}()

	// Accept usr input as canclelation sig
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(done)
		for range unseenLinks {
		}
		for range worklist {
		}
	}()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link, done)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link.url] {
				if depth > 0 && link.gen > depth {
					continue
				}
				seen[link.url] = true
				unseenLinks <- link
			}
		}
	}
}

//!-

// Copied from gopl.io/ch5/links.

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func Extract(url string, cancle <-chan struct{}) ([]string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Cancel = cancle

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

//!-Extract

// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
