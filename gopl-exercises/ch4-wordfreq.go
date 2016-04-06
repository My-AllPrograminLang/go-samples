package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("os.Args: %s\n", os.Args)
	if len(os.Args) < 2 {
		countFreq(os.Stdin)
	} else {
		f, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println(os.Stderr, "wordfreq: %v\n", err)
			os.Exit(1)
		}
		countFreq(f)
	}
}

func countFreq(f *os.File) {
	wordCount := make(map[string]int)

	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		wordCount[input.Text()]++
	}

	for w, count := range wordCount {
		fmt.Printf("%s --> %d\n", w, count)
	}
}
