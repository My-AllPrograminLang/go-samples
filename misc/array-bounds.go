package main

import (
	"fmt"
	"reflect"
)

func main() {
	ops := [...]string{
		20: "plus",
		30: "minus",
		40: "equals",
	}
    fmt.Println("Typeof ops: ", reflect.TypeOf(ops))

    //fmt.Println(ops[50])  // out-of-bounds runtime panic

    // Safer way...
    index := 50
    if index >= 0 && index < len(ops) {
        fmt.Println(ops[index])
    } else {
        fmt.Printf("Sorry, %d is out of bounds!\n", index)
    }

    // In-bounds elements that have not been explicitly initialized have the
    // zero value for the respective type.
    fmt.Printf("At index 22: '%s'\n", ops[22])

    floatarr := [...]float32{
        2: 0.1,
        3: 0.11,
    }
    fmt.Printf("floatarr[1]: %f\n", floatarr[1])

	//fmt.Println(ops['a'])
}
