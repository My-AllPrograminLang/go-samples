package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// exercise 3.10
func comma_buf(s string) string {
	n := len(s)
	buf := new(bytes.Buffer)

	// Emit the part preceding the first comma.
	buf.WriteString(s[:(n % 3)])

	// Invariant: i points to the next place where a comma has to be inserted.
	i := n % 3
	for i < n {
		if i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(s[i : i+3])
		i += 3
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("1234"))
	fmt.Println(comma("1234435"))
	fmt.Println(comma("1"))

	fmt.Println(`----`)
	fmt.Println(comma_buf("1234"))
	fmt.Println(comma_buf("1234435"))
	fmt.Println(comma_buf("1"))
	fmt.Println(comma_buf("123"))
	fmt.Println(comma_buf("123456"))
}
