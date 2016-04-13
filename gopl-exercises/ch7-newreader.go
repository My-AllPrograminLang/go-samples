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
