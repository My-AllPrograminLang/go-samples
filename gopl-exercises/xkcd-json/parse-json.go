package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Episode struct {
	Num        int
	Month      string
	Year       string
	Transcript string
}

func main() {
	filename := os.Args[1]
	fmt.Println(filename)

	var episode Episode
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("%v", err)
	}
	if err := json.Unmarshal(data, &episode); err != nil {
		log.Fatalf("JSON unmarshalling failed: %s", err)
	}

	fmt.Println(episode)
}
