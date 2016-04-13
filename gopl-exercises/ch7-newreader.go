// Implementing strings.NewReader for exercise 7.4

package main

import (
	"fmt"
	"io"
)

type sreader struct {
	s string
}

func (sr *sreader) Read(p []byte) (int, error) {
	// We keep track of current reading position by reslicing s. Naturally this
	// doesn't support things like "unreading" so it's simplistic. We could
	// instead maintain a read index as an integer separately from s.
	// Also, maybe we shouldn't return EOF when there are more bytes remaining to
	// read (even if the number is smaller than requested)?
	nbytes := len(p)
	eof := false
	if len(sr.s) < nbytes {
		nbytes = len(sr.s)
		eof = true
	}

	for i := 0; i < nbytes; i++ {
		p[i] = sr.s[i]
	}

	sr.s = sr.s[nbytes:]

	if eof {
		return nbytes, io.EOF
	} else {
		return nbytes, nil
	}
}

func NewReader(s string) *sreader {
	return &sreader{s}
}

func main() {
	s := NewReader("foobarba")

	bb := make([]byte, 3)
	nb, err := s.Read(bb)
	fmt.Println(string(bb), nb, err)

	nb, err = s.Read(bb)
	fmt.Println(string(bb), nb, err)

	nb, err = s.Read(bb)
	fmt.Println(string(bb), nb, err)

	nb, err = s.Read(bb)
	fmt.Println(string(bb), nb, err)
}
