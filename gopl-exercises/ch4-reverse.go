package main

import "fmt"

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// exercise 4.3
func reversearrptr(aptr *[6]int) {
	for i, j := 0, len(aptr)-1; i < j; i, j = i+1, j-1 {
		aptr[i], aptr[j] = aptr[j], aptr[i]
	}
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(a)

	reverse(a[:])
	fmt.Println(a)

	reversearrptr(&a)
	fmt.Println(a)
}
