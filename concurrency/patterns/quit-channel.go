package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string, quit <-chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
				time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			case <-quit:
				// Note: there's no opportunity for this goroutine to clean up; once
				// select receives the message, the sender will proceed and the program
				// will exit. We could make the quit channel bi-directional and send an
				// ack back to the initiator when cleanup is done. The initiator would
				// send a quit request and then wait for the ack before proceeding.
				return
			}
		}
	}()

	// Return the channel to the caller
	return c
}

func main() {
	rand.Seed(time.Now().UnixNano())
	quit := make(chan bool)
	c := boring("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- true
}
