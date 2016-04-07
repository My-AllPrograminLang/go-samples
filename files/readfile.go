// Read whole file into a byte[]

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	filename := os.Args[1]
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("reading %s: %v", filename, err)
	}

	fmt.Println("Read whole file")
	fmt.Println("---------------")
	fmt.Println(string(data))
}
