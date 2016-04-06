package main

import "fmt"

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// rotate slice s left by n positions (assuming n < len(s))
func rotaterev(s []int, n int) {
	reverse(s[:n])
	reverse(s[n:])
	reverse(s)
}

func gcd(a, b int) int {
	c := a % b
	if c == 0 {
		return b
	}
	return gcd(b, c)
}

func rotatesinglepass(s []int, n int) {
	numcycles := gcd(len(s), n)

	for cycle := 0; cycle < numcycles; cycle++ {
		i := cycle
		for {
			inext := (i + n) % len(s)
			if inext == cycle {
				break
			}
			s[i], s[inext] = s[inext], s[i]
			i = inext
		}
	}
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	rotatesinglepass(a, 4)
	fmt.Println(a)

	//rotaterev(a, 2)
	//fmt.Println(a)

	//fmt.Println(gcd(4, 6))
	//fmt.Println(gcd(4, 7))
	//fmt.Println(gcd(1, 2))
	//fmt.Println(gcd(2, 2))
	//fmt.Println(gcd(4, 2))
}
