package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	const two_to_thirtyone = 1 << 31
	const two_to_sixtythree = 1 << 63

	fmt.Println("Sample...")
	fmt.Println("Reflect TypeOf: ", reflect.TypeOf(two_to_thirtyone))

	// From Go 1.1, int is 64-bit on 64-bit machines, so this is OK.
	var try31int int = two_to_thirtyone
	fmt.Printf("%T %v\n", try31int, try31int)

	// This generates an overflow panic because the value is too large for
	// the type.
	//var tryint32 int32 = two_to_thirtyone

	var try31int64 int64 = two_to_thirtyone
	fmt.Printf("%T %v\n", try31int64, try31int64)

	// This overflows int64...
	//var try63int64 int64 = two_to_sixtythree

	var try63uint64 uint64 = two_to_sixtythree
	fmt.Printf("%T %v\n", try63uint64, try63uint64)

	// Constants have arbitrary precision, and the results of operations
	// involving them can be assigned, as long as the final result fits into the
	// type.
	const bigone = (1 << 96) | 6
	var maskedbig uint64 = bigone & 0xFFFF
	fmt.Printf("%T %v\n", maskedbig, maskedbig)

    // unsigned arithmetic wraps around, as in C
	var maxuint32 uint32 = math.MaxUint32 // 1<<32 - 1
	fmt.Printf("%T %x\n", maxuint32, maxuint32)

	var maxuint32plus2 uint32 = maxuint32 + 2
	fmt.Printf("%T %x\n", maxuint32plus2, maxuint32plus2)

    // unsigned arithmetic wraps around, as in C
	var maxint32 int32 = math.MaxInt32
	fmt.Printf("%T %x\n", maxint32, maxint32)

    // signed arithmetic also doesn't raise exceptions when overflowing, but the
    // result is defined by the underlyinig implementation
	var maxint32plus2 int32 = maxint32 + 2
	fmt.Printf("%T %x\n", maxint32plus2, maxint32plus2)
}
