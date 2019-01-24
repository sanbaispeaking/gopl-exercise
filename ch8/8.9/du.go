// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 250.

// The du3 command computes the disk usage of the files in a directory.
package main

// The du3 variant traverses all directories in parallel.
// It uses a concurrency-limiting counting semaphore
// to avoid opening too many files at once.

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

type rootSize struct {
	root     string
	fileSize int64
}
type periodicReport struct {
	nfiles int64
	nbytes int64
}

//!+
func main() {
	// ...determine roots...

	//!-
	flag.Parse()

	// Determine the initial directories.
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	//!+
	// Traverse each root of the file tree in parallel.
	rootSizes := make(chan rootSize)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, root, &n, rootSizes)
	}
	go func() {
		n.Wait()
		close(rootSizes)
	}()
	//!-

	// Print the results periodically.
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}

	dashboard := make(map[string]periodicReport)

loop:
	for {
		select {
		case rs, ok := <-rootSizes:
			if !ok {
				break loop // fileSizes was closed
			}

			pr := dashboard[rs.root]
			dashboard[rs.root] = periodicReport{pr.nfiles + 1, pr.nbytes + rs.fileSize}

		case <-tick:
			printDiskUsageSeparately(dashboard)
		}
	}

	printDiskUsageSeparately(dashboard) // final totals
	//!+
	// ...select loop...
}

//!-

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func printDiskUsageSeparately(dashboard map[string]periodicReport) {
	for root, report := range dashboard {
		fmt.Printf("%s: %d files  %.1f GB\n", root, report.nfiles, float64(report.nbytes)/1e9)
	}
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
//!+walkDir
func walkDir(root, dir string, n *sync.WaitGroup, fileSizes chan<- rootSize) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(root, subdir, n, fileSizes)
		} else {
			fileSizes <- rootSize{root, entry.Size()}
		}
	}
}

//!-walkDir

//!+sema
// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	// ...
	//!-sema

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
