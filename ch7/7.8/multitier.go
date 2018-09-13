// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 187.

// Sorting sorts a music playlist into a variety of orders.
package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

// Track ...
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

//!+printTracks
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

//!-yearcode

type tierLess func(l, r *Track) bool

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

// makeLess composes multiple cmp funcs
func makeLess(tiers ...func(l, r *Track) int) func(*Track, *Track) bool {
	less := func(x, y *Track) bool {
		for _, tier := range tiers {
			if r := tier(x, y); r == 0 {
				continue
			} else {
				return r == -1
			}
		}
		return false
	}
	return less
}

func main() {
	sort.Sort(customSort{tracks, makeLess(
		// Most recent clicked
		func(x, y *Track) int {
			switch {
			case x.Title < y.Title:
				return -1
			case x.Title > y.Title:
				return 1
			default:
				return 0
			}
		},
		func(x, y *Track) int {
			switch {
			case x.Year < y.Year:
				return -1
			case x.Year > y.Year:
				return 1
			default:
				return 0
			}

		},
		func(x, y *Track) int {
			switch {
			case x.Length < y.Length:
				return -1
			case x.Length > y.Length:
				return 1
			default:
				return 0
			}

		},
	)})

	printTracks(tracks)
}
