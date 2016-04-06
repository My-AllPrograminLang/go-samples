package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func runesInByteslice(s []byte) {
	for i := 0; i < len(s); {
		runeValue, width := utf8.DecodeRune(s[i:])
		fmt.Printf("  %#U starts at byte position %d\n", runeValue, i)
		fmt.Printf("    isspace=%t\n", unicode.IsSpace(runeValue))
		i += width
	}
}

func squashSpaces(s []byte) []byte {
	writeptr := 0
	lastWasSpace := false
	var runeValue rune
	var width int
	for i := 0; i < len(s); {
		runeValue, width = utf8.DecodeRune(s[i:])
		if unicode.IsSpace(runeValue) {
			if !lastWasSpace {
				s[writeptr] = ' '
				writeptr++
				lastWasSpace = true
			}
		} else {
			copy(s[writeptr:], s[i:i+width])
			lastWasSpace = false
			writeptr += width
		}
		i += width
	}
	return s[:writeptr]
}

func main() {
	str := "joyous  help\t   yea  "
	s := []byte(str)
	runesInByteslice(s)

	fmt.Println("---------------------------")
	s = squashSpaces(s)
	runesInByteslice(s)
}
