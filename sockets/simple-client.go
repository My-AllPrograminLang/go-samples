package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:2000")
	if err != nil {
		log.Fatal(err)
		// handle error
	}
	fmt.Fprintf(conn, "foobar")
	conn.Close()
}
