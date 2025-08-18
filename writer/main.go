package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello Playground!")
	fmt.Println(os.Stdout, "Hello Playground!")
	io.WriteString(os.Stdout, "Hello Playground")
}
