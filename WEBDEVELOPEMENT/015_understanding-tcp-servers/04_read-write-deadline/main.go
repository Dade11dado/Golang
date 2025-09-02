package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {

	//Listen to the port on second parameters for the specified
	//method
	li, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panic(err)
	}

	defer li.Close()

	//Loop forever
	for {
		//if domebody calls the connection we acept
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heared you say %s\n", ln)
	}
	defer conn.Close()
}
