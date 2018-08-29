package main

import (
	"flag"
	"fmt"
	"log"

	"gopl.io/ch5/links"
)

// URL with depth indicator recording how many hoops does it need to reach from index page
type URL struct {
	url string
	gen int
}

func crawl(u URL) []URL {
	fmt.Printf("%v of depth %d\n", u.url, u.gen)
	list, err := links.Extract(u.url)
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

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
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
