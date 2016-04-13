// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		// Exercise 1.8: prepend http://
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		fmt.Println("Getting URL: ", url)
		fmt.Println("----------------------")
		resp, err := http.Get(url)
		// Exercise 1.9: print status
		fmt.Printf("status: %v\n", resp.Status)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("----------------------")
		// Exercise 1.7: use io.Copy
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Copy: %v\n", err)
		}
		resp.Body.Close()
	}
}
