package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func sepaggr() {
	t1 := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Printf("sepaggr: %d ns\n", time.Since(t1).Nanoseconds())
	fmt.Println(s)
}

func join() {
	t1 := time.Now()
	s := strings.Join(os.Args[0:], " ")
	fmt.Printf("join: %d ns\n", time.Since(t1).Nanoseconds())
	fmt.Println(s)
}

func main() {
	//for i, arg := range os.Args[0:] {
	//fmt.Printf("%d: %s\n", i, arg)
	//}
	sepaggr()
	join()
}
