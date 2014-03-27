//
package main

import (
	"fmt"
)

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -lmylib
// #include "mylib.h"
import "C"


func main() {
	fmt.Println("Running voidfunc")
	C.voidfunc()

	fmt.Println("Running intfunc")
	fmt.Println(C.intfunc(7))
}
