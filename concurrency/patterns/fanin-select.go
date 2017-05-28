package main // Fan-in from multiple channels using select.
import (
	"fmt"
	"math/rand"
	"time"
)

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

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

func main() {
	rand.Seed(time.Now().UnixNano())
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 20; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Bye")
}
