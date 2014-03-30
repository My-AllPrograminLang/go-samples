package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	path := os.Args[1]

	info, err := os.Stat(path)
	if err != nil {
		log.Fatal("stat error:", err)
	}

	if !info.IsDir() {
		log.Fatal("not a dir")
	}

	infos, _ := ioutil.ReadDir(path)
	for _, info := range infos {
		fmt.Println(info.Name(), "  |  IsDir =", info.IsDir())
	}
}
