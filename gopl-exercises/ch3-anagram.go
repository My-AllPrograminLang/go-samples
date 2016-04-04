package main

import (
	"fmt"
	"reflect"
	"sort"
)

// This approach uses sorting the strings and comparing them. It has to do some
// prep so that the sort package treats rune slicess as sort.Interface

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool { return s[i] < s[j] }

func (s sortRunes) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s sortRunes) Len() int { return len(s) }

func isanagram_sort(s1, s2 string) bool {
	r1 := []rune(s1)
	sort.Sort(sortRunes(r1))
	r2 := []rune(s2)
	sort.Sort(sortRunes(r2))

	fmt.Println(r1)
	fmt.Println(r2)
	return string(r1) == string(r2)
}

func string2runemap(s string) map[rune]int {
	runemap := make(map[rune]int)
	for _, r := range s {
		runemap[r]++
	}
	return runemap
}

func isanagram_map(s1, s2 string) bool {
	s1runemap := string2runemap(s1)
	s2runemap := string2runemap(s2)
	// Just for fun using DeepEqual ;-)
	return reflect.DeepEqual(s1runemap, s2runemap)
}

func main() {
	s1 := "hello"
	s2 := "eloh"

	for i := range s1 {
		fmt.Println(s1[i])
	}

	fmt.Println(isanagram_sort(s1, s2))
	fmt.Println(isanagram_map(s1, s2))
}
