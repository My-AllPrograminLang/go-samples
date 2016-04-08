// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	nfound := ElementByID(doc, "h1")
	fmt.Println(nfound)
	return nil
}

// Modifying forEachNode for ex 5.8 - early exit.
func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) bool {
	if n != nil && n.Type == html.ElementNode {
		fmt.Printf("visiting %s\n", n.Data)
	}

	if pre != nil {
		if cont := pre(n); !cont {
			return false
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if !forEachNode(c, pre, post) {
			return false
		}
	}

	if post != nil {
		if cont := post(n); !cont {
			return false
		}
	}

	return true
}

func ElementByID(doc *html.Node, id string) *html.Node {
	var nfound *html.Node

	pre_finder := func(n *html.Node) bool {
		if n.Type == html.ElementNode && n.Data == id {
			nfound = n
			return false
		} else {
			return true
		}
	}
	forEachNode(doc, pre_finder, nil)
	return nfound
}
