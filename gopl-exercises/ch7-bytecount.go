// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

// Solving exercise 7.1 - counting words through a Writer interface.
type WordCounter int

func (wc *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*wc += 1
	}
	return int(*wc), nil
}

// Solving exercise 7.2
type ForwardingByteCounter struct {
	count int64
	w     io.Writer
}

func (fbc *ForwardingByteCounter) Write(p []byte) (int, error) {
	fbc.count += int64(len(p))
	return fbc.w.Write(p)
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	fbc := ForwardingByteCounter{0, w}
	return &fbc, &(fbc.count)
}

func main() {
	//!+main
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)

	var wc WordCounter
	wc.Write([]byte("this is the winter of our\n discontent"))
	fmt.Println(wc)

	w, cc := CountingWriter(os.Stdout)
	w.Write([]byte("hello stranger\n"))
	fmt.Println(*cc)
}
