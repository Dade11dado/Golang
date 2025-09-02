package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	ls, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer ls.Close()

	for {
		conn, err := ls.Accept()
		if err != nil {
			log.Fatal(err)

		}
		go handle(conn)
	}
}

func handle(c net.Conn) {
	scanner := bufio.NewScanner(c)

	for scanner.Scan() {
		ls := scanner.Text()
		d := bf13([]byte(ls))
		fmt.Fprintf(c, "%s -  %s \n", ls, d)
	}
}

func bf13(bs []byte) []byte {
	var r13 = make([]byte, len(bs))
	for i, v := range bs {
		if v <= 109 {
			r13[i] = v + 13
		} else {
			r13[i] = v - 13
		}
	}
	return r13
}
