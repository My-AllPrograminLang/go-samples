// Launch a number of concurrent http GET requests.
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var template string = "https://github.com/eliben/pycparser/pull/%d"

func launchConcurrentGets(start, num int) []http.Header {
	c := make(chan http.Header)
	defer close(c)

	for i := start; i < start+num; i++ {
		go func(i int) {
			url := fmt.Sprintf(template, i)
			resp, err := http.Get(url)
			if err != nil {
				log.Fatalf("error in http.Get: %s", err)
			}
			c <- resp.Header
		}(i)
	}

	responses := make([]http.Header, num)
	for i := 0; i < num; i++ {
		responses = append(responses, <-c)
	}
	return responses
}

func main() {
	t1 := time.Now()
	launchConcurrentGets(10, 20)
	log.Printf("elapsed: %s", time.Since(t1))
}
