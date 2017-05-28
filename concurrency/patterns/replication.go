package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string
type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

// First takes a query string and sends it to all the given replicas; the first
// replica to return a result has its result returned from this function.
func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	res := First("foo",
		fakeSearch("web1"),
		fakeSearch("web2"),
		fakeSearch("web3"))
	elapsed := time.Since(start)
	fmt.Println(res)
	fmt.Println(elapsed)
}
