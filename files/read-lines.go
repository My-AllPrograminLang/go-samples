package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	input := bufio.NewScanner(f)
	input.Split(bufio.ScanLines)
	for input.Scan() {
		fmt.Println("line:", input.Text())
		fmt.Println("asbytes:", input.Bytes())
	}
}
