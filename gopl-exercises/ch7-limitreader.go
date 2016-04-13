// Implementing io.LimitReader for exercise 7.5

package main

import (
	"fmt"
	"io"
	"strings"
)

type limiter struct {
	r          io.Reader
	nremaining int
}

func (lr *limiter) Read(p []byte) (n int, err error) {
	n, err = lr.r.Read(p)
	lr.nremaining -= n
	if lr.nremaining <= 0 {
		err = io.EOF
	}
	return
}

func LimitReader(r io.Reader, n int) io.Reader {
	return &limiter{r, n}
}

func main() {
	r := strings.NewReader("abracadabra")
	lr := LimitReader(r, 5)

	bb := make([]byte, 3)
	for i := 0; i < 5; i++ {
		nb, err := lr.Read(bb)
		fmt.Println(string(bb), nb, err)
	}
}
