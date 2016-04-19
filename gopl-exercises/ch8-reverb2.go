// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, wg *sync.WaitGroup, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	wg.Done()
}

//!+
func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)

	// Exercise 8.8
	// To do a select-based timeout on the scanner read, launch a goroutine that
	// feeds each line scanned into a channel. Then we can select on this channel.
	inChan := make(chan string)
	go func() {
		for input.Scan() {
			inChan <- input.Text()
		}
		close(inChan)
	}()

	// WaitGroup for solving ex. 8.4
	var wg sync.WaitGroup
LineLoop:
	for {
		select {
		case line, ok := <-inChan:
			if !ok {
				break
			}
			wg.Add(1)
			go echo(c, &wg, line, 1*time.Second)
		case <-time.After(5 * time.Second):
			log.Println("Disconnecting client after timeout")
			break LineLoop
		}
	}
	wg.Wait()
	tcpconn := c.(*net.TCPConn)
	tcpconn.CloseWrite()
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
