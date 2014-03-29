package main

import (
	"path/filepath"
	"fmt"
	"os"
)

func describePath(p string) {
	fmt.Println("path is:", p)
	fmt.Println("dir of path:", filepath.Dir(p))
}

func walker(path string, info os.FileInfo, err error) error {
	fmt.Println("Walk in:", path)
	if info.IsDir() {
		fmt.Println("  a dir")
	}
	return nil
}

func main() {
	path := os.Args[1]
	describePath(path)
	filepath.Walk(path, walker)
}
