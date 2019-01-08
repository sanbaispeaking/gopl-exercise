package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		return
	}
	for {
		c, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(c)
	}

}

func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	cwd, _ := os.Getwd()

	for input.Scan() {
		text := input.Text()
		pieces := strings.Split(text, " ")
		cmd := pieces[0]

		switch cmd {
		case "ls":
			files, _ := ioutil.ReadDir(cwd)
			for _, f := range files {
				fmt.Fprintf(c, "%s\n", f.Name())
			}
		case "cd":
			dir := filepath.Join(cwd, pieces[1])
			os.Chdir(dir)
		case "exit":
			fmt.Fprintf(c, "Bye.\n")
			return
		}
	}

}
