// This is a small sample to explore how the Go compiler optimized switches.
//
// To compile & disassemble, enter the directory this file is located in and
// type:
//
// $ go build switchfunc.go
// $ objdump -D switchfunc | less
//
// It's immediately obvious that Go does not create a jump table, but rather
// uses a mix of binary search and linear search. More information is available
// in these links:
//
// * https://groups.google.com/forum/#!msg/golang-nuts/IURR4Z2SY7M/R7ORD_yDix4J
// * https://groups.google.com/forum/#!topic/golang-nuts/Sz7i4zQT9os
// * http://stackoverflow.com/questions/9928221/table-of-functions-vs-switch-in-golang
//
package main

import (
	"fmt"
	"os"
	"strconv"
)

func swf(n int) int {
	switch n {
	case 2:
		return 22
	case 4:
		return 204
	case 6:
		return 209
	case 7:
		return 702
	case 8:
		return 809
	case 9:
		return 903
	case 13:
		return 29
	case 14:
		return 30
	case 15:
		return 31
	case 16:
		return 34
	case 17:
		return 56
	case 18:
		return 89
	case 19:
		return 29
	case 20:
		return 44
	case 21:
		return 29
	case 22:
		return 39
	default:
		return 0
	}
}

func main() {
	for i, v := range os.Args {
		fmt.Println(i, v)
		s, _ := strconv.Atoi(v)
		fmt.Println("--", swf(s))
	}
}
