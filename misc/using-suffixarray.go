package main

import (
	"bytes"
	"fmt"
	"index/suffixarray"
	"regexp"
	"strings"
)

// getStringFromIndex finds the string in data based on given index.
// The index points into data somewhere inside a string that has a \x00 to its
// left and to its right.
// Assumes data is well formed and index is in bounds.
func getStringFromIndex(data []byte, index int) string {
	var start, end int
	for i := index - 1; i >= 0; i-- {
		if data[i] == 0 {
			start = i + 1
			break
		}
	}
	for i := index + 1; i < len(data); i++ {
		if data[i] == 0 {
			end = i
			break
		}
	}
	return string(data[start:end])
}

func main() {
	words := []string{
		"aardvark",
		"happy",
		"hello",
		"hero",
		"he",
		"hotel",
		"hahem",
	}
	// Use \x00 to delimit strings.
	data := []byte("\x00" + strings.Join(words, "\x00") + "\x00")
	sa := suffixarray.New(data)
	buf := &bytes.Buffer{}
	sa.Write(buf)
	fmt.Println("Serialized size:", buf.Len())

	fmt.Println("Using Lookup:")
	indices := sa.Lookup([]byte("he"), -1)

	for _, idx := range indices {
		fmt.Println(getStringFromIndex(data, idx))
	}

	fmt.Println("Using FindAllIndex:")
	r := regexp.MustCompile("he")
	matches := sa.FindAllIndex(r, -1)

	fmt.Println(matches)

	//fmt.Println(idx)

	//match, err := regexp.Compile("\x00he[^\x00]*")
	//if err != nil {
	//panic(err)
	//}
	//ms := sa.FindAllIndex(match, -1)

	//for _, m := range ms {
	//start, end := m[0], m[1]
	//fmt.Printf("match = %q\n", joinedStrings[start+1:end])
	//}
}
