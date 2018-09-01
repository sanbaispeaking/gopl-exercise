package main

import (
	"fmt"
	"log"
	"net/http"
)

func mirroredQuery() string {
	responses := make(chan string, 3)
	done := make(chan struct{}, 1)
	go func() { responses <- request("http://asia.gopl.io", done) }()
	go func() { responses <- request("http://europe.gopl.io", done) }()
	go func() { responses <- request("http://america.gopl.io", done) }()

	for {
		select {
		case fastest := <-responses:
			close(done)
			for range responses {
			}
			return fastest
		default:
		}
	}
}

func request(hostname string, cancel <-chan struct{}) (response string) {
	// ignore errors
	req, _ := http.NewRequest("GET", hostname, nil)
	req.Cancel = cancel

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("getting %s failed: %s", hostname, err)
	}

	defer rsp.Body.Close()
	return hostname
}

func main() {
	fastest := mirroredQuery()
	fmt.Println("Fastest mirror is: ", fastest)
}
