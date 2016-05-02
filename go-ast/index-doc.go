package main

import (
	"fmt"
	"go/build"
	"go/doc"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
)

func importDir(dir string) *build.Package {
	// Try to import the directory; if unsuccessful, just return nil as the
	// package.
	pkg, err := build.ImportDir(dir, build.ImportComment)
	if err != nil {
		return nil
	}
	return pkg
}

func parsePackage(pkg *build.Package) {
	fs := token.NewFileSet()
	// include tells parser.ParseDir which files to include.
	// That means the file must be in the build package's GoFiles or CgoFiles
	// list only (no tag-ignored files, tests, swig or other non-Go files).
	include := func(info os.FileInfo) bool {
		for _, name := range pkg.GoFiles {
			if name == info.Name() {
				return true
			}
		}
		for _, name := range pkg.CgoFiles {
			if name == info.Name() {
				return true
			}
		}
		return false
	}
	pkgs, err := parser.ParseDir(fs, pkg.Dir, include, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure they are all in one package.
	if len(pkgs) != 1 {
		log.Fatalf("multiple packages in directory %s", pkg.Dir)
	}
	astPkg := pkgs[pkg.Name]

	docPkg := doc.New(astPkg, pkg.ImportPath, doc.AllDecls)
	for _, typ := range docPkg.Types {
		docPkg.Consts = append(docPkg.Consts, typ.Consts...)
		docPkg.Vars = append(docPkg.Vars, typ.Vars...)
		docPkg.Funcs = append(docPkg.Funcs, typ.Funcs...)
	}

	// TODO: here
	fmt.Println(pkg.ImportPath)
	fmt.Println(astPkg.Name)
	fmt.Println(astPkg)
}

func processPath(dir string) {
	pkg := importDir(dir)
	if pkg != nil {
		parsePackage(pkg)
	}
}

func walker(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println("ERROR: walking", path)
		return err
	}

	if info.IsDir() {
		if info.Name() == "internal" || info.Name() == "testdata" {
			return filepath.SkipDir
		}
		fmt.Println("=======>", path)
		processPath(path)
	}

	return nil
}

func main() {
	rootdir := os.Args[1]
	filepath.Walk(rootdir, walker)
}
