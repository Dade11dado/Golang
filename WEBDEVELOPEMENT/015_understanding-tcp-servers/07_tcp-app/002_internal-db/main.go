package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handle(conn)
	}
}

func handle(c net.Conn) {
	defer c.Close()

	//istructions
	io.WriteString(c, "\n IN MEMORY DATABASE \n\n"+
		"USE: \n"+
		"SET key value \n"+
		"GET key \n"+
		"DEL key \n")

	//read and write
	data := make(map[string]string)
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)
		switch strings.ToUpper(fs[0]) {
		case "GET":
			k := fs[1]
			v := data[k]
			fmt.Fprintf(c, "%s\n", v)
		case "SET":
			if len(fs) != 3 {
				fmt.Fprintln(c, "EXPECTED VALUE")
				continue
			}
			k := fs[1]
			v := fs[2]
			data[k] = v
		case "DEL":
			k := fs[1]
			delete(data, k)
		default:
			fmt.Fprintln(c, "INVALID INOUT")
		}
	}

}
