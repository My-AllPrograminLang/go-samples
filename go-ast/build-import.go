// Uses build.ImportDir to import package from a directory.
//
// Walks the given directory tree recursively, tries to import each directory
// found as a Go package and reports some information about it.
package main

import (
	"fmt"
	"go/build"
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

func describePackage(pkg build.Package) {
	fmt.Println("dir:", pkg.Dir)
	fmt.Println("name:", pkg.Name)
	fmt.Println("root:", pkg.Root)
	fmt.Println("import path:", pkg.ImportPath)
	fmt.Println("go files:", pkg.GoFiles)
	fmt.Println("cgo files:", pkg.CgoFiles)
	fmt.Println("imports:", pkg.Imports)
}

func walker(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println("ERROR: walking", path)
		return err
	}

	if info.IsDir() && info.Name() != "internal" && info.Name() != "testdata" {
		fmt.Println("==========================================")
		fmt.Println(path)
		fmt.Println("==========================================")
		pkg := importDir(path)
		if pkg != nil {
			describePackage(*pkg)
		}
	}

	return nil
}

func main() {
	rootdir := os.Args[1]
	filepath.Walk(rootdir, walker)
}
