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

	// Use ioutil.ReadFile to read the whole file into a byte slice.
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("reading %s: %v", filename, err)
	}

	fmt.Println("Read whole file")
	fmt.Println("---------------")
	fmt.Println(string(data))

	// A more standard method: use os.Open to open the file and then os.File to
	// Read from it into a slice.
	b := make([]byte, 100)
	f, err := os.Open(filename)
	f.Read(b)

	fmt.Println("\n")
	fmt.Println("First 100 bytes from file")
	fmt.Println("-------------------------")
	fmt.Println(string(b))
}
