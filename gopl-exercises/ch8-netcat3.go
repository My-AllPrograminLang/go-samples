// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 227.

// Netcat is a simple read/write client for TCP servers.
package main

import (
	"io"
	"log"
	"net"
	"os"
)

//!+
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	// For solving ex 8.3: use TCPConn type assertion so we can then call
	// CloseWrite on it.
	tcpconn := conn.(*net.TCPConn)
	done := make(chan struct{})
	log.Println("launching listener")
	go func() {
		io.Copy(os.Stdout, tcpconn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	log.Println("copying from stdin")
	mustCopy(tcpconn, os.Stdin)
	log.Println("done copying from stdin - closing conn")
	tcpconn.CloseWrite()
	<-done // wait for background goroutine to finish
}

//!-

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
