package main

import "fmt"

func xorBytes(x, y []byte) []byte {
	result := make([]byte, len(x))
	for i := 0; i < len(x); i++ {
		result[i] = x[i] ^ y[i]
	}
	return result
}

func main() {
	s1 := []byte("abracadabr")
	s2 := []byte("0123456789")

	xb := xorBytes(s1, s2)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(xb)
}
