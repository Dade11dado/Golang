package main

import (
	"fmt"
	"io"
	"log"
	"net"
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
		//write to the connection
		io.WriteString(conn, "\nHello from TCP service")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprint(conn, "%v", "Well, i hope")

		conn.Close()
	}
}
