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
	fmt.Fprintln(c, "version 0.1")

	for input.Scan() {
		cwd, _ := os.Getwd()
		args := strings.Split(input.Text(), " ")
		cmd := args[0]
		args = args[1:]

		switch cmd {

		case "ls":
			files, _ := ioutil.ReadDir(cwd)
			for _, f := range files {
				fmt.Fprintf(c, "%s\n", f.Name())
			}

		case "cd":
			if len(args) > 0 {
				dir := filepath.Join(cwd, args[0])
				if err := os.Chdir(dir); err != nil {
					if os.IsPermission(err) {
						fmt.Fprintln(c, err)
					}
				}
			}

		case "get":
			if len(args) < 1 {
				fmt.Fprintln(c, "missing file operand")
				break
			}

		case "exit":
			fmt.Fprintf(c, "Bye.\n")
			return
		}
	}

}
