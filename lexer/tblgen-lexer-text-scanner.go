// Lexer for the TableGen language with text.Scanner
// I get slightly different tokens here because the scanner is different, but
// this is just a basic test.

package main

import (
	"fmt"
	"os"
	"text/scanner"
	"time"
)

func readAndMeasure(filename string) {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	toks := make([]string, 0, 200000)
	startTime := time.Now()
	var s scanner.Scanner
	s.Init(f)
	s.Mode = (scanner.ScanChars | scanner.ScanStrings | scanner.ScanIdents |
		scanner.ScanFloats | scanner.ScanChars | scanner.ScanComments)
	var tok rune
	for tok != scanner.EOF {
		tok = s.Scan()
		toks = append(toks, s.TokenText())
	}
	endTime := time.Now()

	fmt.Println("Elapsed:", endTime.Sub(startTime))
	fmt.Println("toks[:10]: ", toks[:10])
}

func main() {
	for i := 0; i < 10; i++ {
		readAndMeasure("/tmp/input.td")
	}
}
