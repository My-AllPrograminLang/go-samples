// Exercise 5.9: expand function for replacing $foo vars

package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func expand(s string, f func(string) string) string {
	var parts []string

	i := 0
	for {
		s = s[i:]
		nextrep := strings.IndexRune(s, '$')
		if nextrep == -1 {
			parts = append(parts, s)
			break
		} else {
			parts = append(parts, s[:nextrep])
			s = s[nextrep+1:]
			endofname := strings.IndexFunc(s, func(c rune) bool {
				return !unicode.IsLetter(c)
			})
			if endofname == -1 {
				endofname = len(s)
			}
			parts = append(parts, f(s[:endofname]))
			i = endofname
		}
	}

	return strings.Join(parts, "")
}

func main() {
	fmt.Println(expand(os.Args[1], func(g string) string {
		return ">" + g + "<"
	}))
}
