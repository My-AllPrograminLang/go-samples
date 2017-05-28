// Fan-in from multiple channels using separate go-routines.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	// Return the channel to the caller
	return c
}

// Fans in two channels into a new channel that's returned.
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	// Simultaneosly block-read input1 and input2, pushing values into c, by
	// launching two goroutines that do this.
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func main() {
	rand.Seed(time.Now().UnixNano())
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Bye")
}
