// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

//!+broadcaster

type namedClient struct {
	ch  chan<- string // an outgoing message channels
	who string
}

var (
	entering = make(chan namedClient)
	leaving  = make(chan namedClient)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[namedClient]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.ch <- msg
			}

		case cli := <-entering:
			if len(clients) > 0 {
				others := "Online:\n"
				for client := range clients {
					others += ("\t" + client.who + "\n")
				}
				cli.ch <- others
			}
			clients[cli] = true
			// Display current active clients to new arrival

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	// setup a client for new commer
	ch := make(chan string) // outgoing client messages
	who := conn.RemoteAddr().String()
	client := namedClient{ch, who}
	go clientWriter(conn, ch)

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- client

	buffer := make(chan string)
	go func() {
		for {
			select {
			case <-time.After(10 * time.Second):
				conn.Close()
				return
			case raw := <-buffer:
				messages <- who + ": " + raw
			}
		}

	}()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		buffer <- input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- client
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main
