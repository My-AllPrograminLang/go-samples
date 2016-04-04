package main

import (
	"bytes"
	"fmt"
	"strings"
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

	// Invariant: i points to the next place where a comma has to be inserted.
	i := n % 3

	// Emit the part preceding the first comma.
	buf.WriteString(s[:i])

	for i < n {
		if i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(s[i : i+3])
		i += 3
	}
	return buf.String()
}

// exercise 3.11
func commafp(s string) string {
	n := len(s)
	buf := new(bytes.Buffer)

	startofnum := 0

	// Same design as comma_buf, but we ignore an optional leading sign and stop
	// at the floating point.
	// startofnum: index to where the number actually starts (past optional sign).
	if len(s) >= 1 && (s[0] == '+' || s[0] == '-') {
		startofnum = 1
	}

	endofnum := strings.Index(s, ".")
	if endofnum < 0 {
		endofnum = n
	}

	// Write out anything that comes before the number (optional sign).
	buf.WriteString(s[:startofnum])
	numlen := endofnum - startofnum

	// From here on similar to comma_buf, except that comma-tization is applied
	// to the number in range [startofnum:endofnum) instead of [0:n)
	i := startofnum + numlen%3

	buf.WriteString(s[startofnum:i])

	for i < endofnum {
		if i != startofnum {
			buf.WriteByte(',')
		}
		buf.WriteString(s[i : i+3])
		i += 3
	}

	buf.WriteString(s[endofnum:])

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
	fmt.Println(comma_buf("12"))
	fmt.Println(comma_buf(""))
	fmt.Println(comma_buf("123"))
	fmt.Println(comma_buf("123456"))

	fmt.Println(`----`)
	fmt.Println(commafp("1234"))
	fmt.Println(commafp("1234435"))
	fmt.Println(commafp("1"))
	fmt.Println(commafp("12"))
	fmt.Println(commafp(""))
	fmt.Println(commafp("123"))
	fmt.Println(commafp("123456"))
	fmt.Println(commafp("-1"))
	fmt.Println(commafp("-123"))
	fmt.Println(commafp("+1234"))
	fmt.Println(commafp("+1234.12"))
	fmt.Println(commafp("1.12124"))
	fmt.Println(commafp("192.12124"))
}
