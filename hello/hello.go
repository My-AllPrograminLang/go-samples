package main

import (
	"fmt"
	"go-samples/hello/newmath"

	// dir name for the import
	"go-samples/hello/pkgdir"
)

func main() {
	fmt.Printf("Hello, world.  Sqrt(2) = %v\n", newmath.Sqrt(2))

	// but package name is actually imported
	somepkg.Something("202020")
}
