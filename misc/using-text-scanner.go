package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

func main() {
	const src = `
	// This is scanned code.
	if a > 10 {
		someParsable = text
	}`
	var s scanner.Scanner
	s.Init(strings.NewReader(src))

	// Set mode to not skip comments
	s.Mode = scanner.ScanIdents | scanner.ScanFloats | scanner.ScanChars | scanner.ScanComments
	var tok rune
	for tok != scanner.EOF {
		tok = s.Scan()
		fmt.Println("At position", s.Pos(), ":", s.TokenText())
	}
}
