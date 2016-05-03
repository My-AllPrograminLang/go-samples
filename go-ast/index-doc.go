package main

import (
	"fmt"
	"go/build"
	"go/doc"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Type struct {
	Name string
	Syms []string
}

type Package struct {
	Name       string
	ImportPath string
	Syms       []string
	Types      []Type
}

func (pkg *Package) String() string {
	s := fmt.Sprintf("package '%s', import path = '%s'\n", pkg.Name, pkg.ImportPath)
	s += fmt.Sprintf("Syms = [%s]\n", strings.Join(pkg.Syms, ", "))
	for _, typ := range pkg.Types {
		s += fmt.Sprintf("Type '%s':\n", typ.Name)
		s += fmt.Sprintf("  Syms = [%s]\n", strings.Join(typ.Syms, ", "))
	}

	return s
}

func (pkg *Package) Stats() (numSyms int, totalSymLen int) {
	numSyms = len(pkg.Syms)
	numSyms += len(pkg.Types)
	for _, typ := range pkg.Types {
		numSyms += len(typ.Syms)
	}

	for _, s := range pkg.Syms {
		totalSymLen += len(s)
	}
	for _, typ := range pkg.Types {
		totalSymLen += len(typ.Name)
		for _, s := range typ.Syms {
			totalSymLen += len(s)
		}
	}
	return
}

func importDir(dir string) *build.Package {
	// Try to import the directory; if unsuccessful, just return nil as the
	// package.
	pkg, err := build.ImportDir(dir, build.ImportComment)
	if err != nil {
		return nil
	}
	return pkg
}

func parsePackage(buildPkg *build.Package) *Package {
	fs := token.NewFileSet()
	// include tells parser.ParseDir which files to include.
	// That means the file must be in the build package's GoFiles or CgoFiles
	// list only (no tag-ignored files, tests, swig or other non-Go files).
	include := func(info os.FileInfo) bool {
		for _, name := range buildPkg.GoFiles {
			if name == info.Name() {
				return true
			}
		}
		for _, name := range buildPkg.CgoFiles {
			if name == info.Name() {
				return true
			}
		}
		return false
	}
	buildPkgs, err := parser.ParseDir(fs, buildPkg.Dir, include, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure they are all in one package.
	if len(buildPkgs) != 1 {
		log.Fatalf("multiple packages in directory %s", buildPkg.Dir)
	}

	astPkg := buildPkgs[buildPkg.Name]
	docPkg := doc.New(astPkg, buildPkg.ImportPath, doc.AllDecls)
	//for _, typ := range docPkg.Types {
	//docPkg.Consts = append(docPkg.Consts, typ.Consts...)
	//docPkg.Vars = append(docPkg.Vars, typ.Vars...)
	//docPkg.Funcs = append(docPkg.Funcs, typ.Funcs...)
	//}

	pkg := &Package{
		Name:       buildPkg.Name,
		ImportPath: buildPkg.ImportPath}

	// Populate pkg's non-type symbols with exported functions and vars.
	for _, f := range docPkg.Funcs {
		if isExported(f.Name) {
			pkg.Syms = append(pkg.Syms, f.Name)
		}
	}

	// TODO: all vars point to a single package-variables id in the HTML link.
	for _, v := range docPkg.Vars {
		for _, name := range v.Names {
			if isExported(name) {
				pkg.Syms = append(pkg.Syms, name)
			}
		}
	}

	for _, c := range docPkg.Consts {
		for _, name := range c.Names {
			if isExported(name) {
				pkg.Syms = append(pkg.Syms, name)
			}
		}
	}

	for _, docType := range docPkg.Types {
		if isExported(docType.Name) {
			t := Type{Name: docType.Name}

			for _, f := range docType.Funcs {
				if isExported(f.Name) {
					t.Syms = append(t.Syms, f.Name)
				}
			}
			for _, m := range docType.Methods {
				if isExported(m.Name) {
					t.Syms = append(t.Syms, m.Name)
				}
			}
			for _, v := range docType.Vars {
				for _, name := range v.Names {
					if isExported(name) {
						t.Syms = append(t.Syms, name)
					}
				}
			}
			for _, c := range docType.Consts {
				for _, name := range c.Names {
					if isExported(name) {
						t.Syms = append(t.Syms, name)
					}
				}
			}

			pkg.Types = append(pkg.Types, t)
		}
	}

	//fmt.Println("types:")
	//for _, typ := range docPkg.Types {
	//if isExported(typ.Name) {
	//fmt.Printf("%s ", typ.Name)
	//}
	//}
	//fmt.Println()

	return pkg
}

func processPath(dir string) *Package {
	buildPkg := importDir(dir)
	if buildPkg != nil {
		pkg := parsePackage(buildPkg)
		log.Println(pkg)
		numSyms, totalSymLen := pkg.Stats()
		log.Printf("Stats: %d syms, total len = %d\n", numSyms, totalSymLen)
		return pkg
	}
	return nil
}

// startsWithUpper reports whether the name starts with an uppercase letter.
func startsWithUpper(name string) bool {
	ch, _ := utf8.DecodeRuneInString(name)
	return unicode.IsUpper(ch)
}

func isExported(name string) bool {
	return startsWithUpper(name)
}

func main() {
	// Configure logging
	log.SetOutput(ioutil.Discard)

	rootdir := os.Args[1]

	var pkgs []*Package
	walker := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println("ERROR: walking", path)
			return err
		}

		if info.IsDir() {
			if info.Name() == "internal" || info.Name() == "testdata" {
				return filepath.SkipDir
			}
			log.Println("=======>", path)
			pkg := processPath(path)
			if pkg != nil {
				pkgs = append(pkgs, pkg)
			}
		}

		return nil
	}
	filepath.Walk(rootdir, walker)

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("Totals...")
	fmt.Printf("%d packages\n", len(pkgs))

	var nSyms int
	var nSize int
	for _, pkg := range pkgs {
		numSyms, totalSymLen := pkg.Stats()
		nSyms += numSyms
		nSize += totalSymLen
	}

	fmt.Printf("num symbols = %d\n", nSyms)
	fmt.Printf("total size = %d\n", nSize)
}
