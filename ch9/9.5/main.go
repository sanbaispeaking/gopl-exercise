package main

import (
	"fmt"
	"time"
)

func main () {
	in, out := make(chan uint64), make(chan uint64)
	done := make(chan struct{})

	go func () {
		// sending to a unbuffered channel blocks until another goroutine performs
		// receive on the same channel. See page 226
		in <- 0
	}()

	// pingpong read message from in, increase it by 1, send it back to out
	pingpong := func (in, out chan uint64, cancel chan struct{}) {
		for {
			select {
			case ping := <-in:
				ping ++
				out <- ping
			case <- cancel:
				close(out)
				return
			}
		}
	}

	go pingpong(in, out, done)
	go pingpong(out, in, done)

	select {
	case <- time.After(1 * time.Second):
		close(done)
		count := <- in
		countCandidate := <- out
		if count < countCandidate {
			count = countCandidate
		}
		fmt.Println(count)
	}

	// Output:
	// around 2295442 
	// 2.7GHz Intel Core i5 & macOS 10.13.6
}
