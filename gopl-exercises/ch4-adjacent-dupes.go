package main

import "fmt"

func dedup(s []string) []string {
	writeptr := 0
	for i := range s {
		if i > 0 && s[i] != s[writeptr] {
			writeptr++
			s[writeptr] = s[i]
		}
	}
	return s[:writeptr+1]
}

func main() {
	//ss := []string{"hallo", "hello", "hello", "more", "moar", "moar"}
	ss := []string{"hallo", "hallo", "hallo"}
	fmt.Println(ss)

	ss = dedup(ss)
	fmt.Println(ss)
}
