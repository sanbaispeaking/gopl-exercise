// Reverb is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

//!+
func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	data := make(chan string)

	//Scan blocks, handle seperately
	go func() {
		defer close(data)
		for input.Scan() {
			data <- input.Text()
		}
	}()

	for {
		select {
		case <-time.After(3 * time.Second):
			c.Close()
			return
		case s := <-data:
			// NOTE: ignoring potential errors from input.Err()
			go echo(c, s, 1*time.Second)
		}
	}
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}

}
