// Basic parsing of a Go file with parser.ParseFile
package main

import (
	"bufio"
	"fmt"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	fset := token.NewFileSet() // positions are relative to fset

	f, err := parser.ParseFile(fset, "<file>", bufio.NewReader(os.Stdin), 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the imports from the file's AST.
	for _, s := range f.Imports {
		fmt.Println(s.Path.Value)
	}
}
