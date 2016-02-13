package main

import (
	"strings"
	"testing"
)

func BenchmarkSepAggr(b *testing.B) {
	sar := []string{"jko", "yayaya", "foo", "bar", "baz", "kwa"}
	for i := 0; i < b.N; i++ {
		s, sep := "", ""
		for _, arg := range sar {
			s += sep + arg
			sep = " "
		}
	}
}

var globs string

func BenchmarkJoin(b *testing.B) {
	sar := []string{"jko", "yayaya", "foo", "bar", "baz", "kwa"}
	s := ""
	for i := 0; i < b.N; i++ {
		s = strings.Join(sar, " ")
	}
	globs = s
}
