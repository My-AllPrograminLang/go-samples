// A couple of samples from  http://blog.golang.org/c-go-cgo
package main

import (
	"fmt"
)

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "unsafe"

func Print(s string) {
	cs := C.CString(s)
	C.fputs(cs, (*C.FILE)(C.stdout))
	C.free(unsafe.Pointer(cs))
}

func Random() int {
	return int(C.random())
}

func Seed(i int) {
    C.srandom(C.uint(i))
}

func main() {
	Seed(24)
	fmt.Println(Random())
	fmt.Println(Random())
	fmt.Println("now string")
	fmt.Print("joe")
}
