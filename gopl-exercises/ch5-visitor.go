// Refactoring ch5-findlinks.go into a more general visitor pattern; as an
// example, solving exercise 5.2 - keeping count of elements by type.

package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

type visitor func(*html.Node)

func visit(n *html.Node, vfunc visitor) {
	if n == nil {
		return
	}
	vfunc(n)
	visit(n.FirstChild, vfunc)
	visit(n.NextSibling, vfunc)
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	typecount := make(map[string]int)
	typecount_visitor := func(n *html.Node) {
		if n.Type == html.ElementNode {
			typecount[n.Data]++
		}
	}

	visit(doc, typecount_visitor)

	for k, v := range typecount {
		fmt.Printf("%s -> %d\n", k, v)
	}
}
