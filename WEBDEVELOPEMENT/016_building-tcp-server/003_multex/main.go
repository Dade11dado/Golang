package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

var m string
var uri string

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

func handle(conn net.Conn) {
	defer conn.Close()

	request(conn)

	respond(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			m = strings.Fields(ln)[0]
			uri = strings.Fields(ln)[1]
			fmt.Println("****METHOD", m)
			fmt.Println("******URI", uri)
		}
		if ln == "" {
			break
		}
		i++
	}
}

func respond(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	</head>
	<title>Hello</title>
	<body> <strong>Ciao</strong>
	</body>
	</html>
	`

	contact := `<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	</head>
	<title>Hello</title>
	<body> <strong>COntact</strong>
	</body>
	</html>
	`

	if uri == "/Home" {
		fmt.Println("Here")
		fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content lenght: %d\r\n", len(body))
		fmt.Fprintf(conn, "Content type: text/html\r\n")
		fmt.Fprintf(conn, "\n")
		fmt.Fprintf(conn, body)
		return
	}

	if uri == "/Contact" {
		fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content lenght: %d\r\n", len(body))
		fmt.Fprintf(conn, "Content type: text/html\r\n")
		fmt.Fprintf(conn, "\n")
		fmt.Fprintf(conn, contact)
		return
	}

	log.Fatal("No route discovered")

}
