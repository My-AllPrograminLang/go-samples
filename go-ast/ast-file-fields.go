// Parsing with parser.ParseFile and exploring different AST fields and
// positions.
package main

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func showIdent(ident *ast.Ident, fset *token.FileSet) string {
	return fmt.Sprintf("Ident('%s') pos=%v", ident.Name, fset.Position(ident.NamePos))
}

func main() {
	fset := token.NewFileSet() // positions are relative to fset

	// Without setting this flag, comments in the code are not parsed.
	flags := parser.ParseComments
	f, err := parser.ParseFile(fset, "<file>", bufio.NewReader(os.Stdin), flags)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Package name:", showIdent(f.Name, fset))
	fmt.Println("Associated doc:")
	if f.Doc == nil {
		fmt.Println("  nil")
	} else {
		fmt.Println(f.Doc.Text())
	}

	fmt.Println("Imports:")
	for _, ispec := range f.Imports {
		fmt.Println("  ", ispec.Path.Value)
	}

	fmt.Println("Functions:")
	for _, decl := range f.Decls {
		if fv, isfunc := decl.(*ast.FuncDecl); isfunc {
			fmt.Println("  Name:", showIdent(fv.Name, fset))
			fmt.Println("  Pos:", fset.Position(fv.Pos()))
			fmt.Println("  End:", fset.Position(fv.End()))
		}
	}

	fmt.Println("Comments:")
	for _, c := range f.Comments {
		fmt.Println(c.Text())
	}
}
