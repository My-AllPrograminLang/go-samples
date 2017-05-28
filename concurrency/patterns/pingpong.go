package main

import (
	"fmt"
	"time"
)

type Ball struct{ hits int }

func main() {
	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table)

	// Start game by adding a ball; comment this line out to see Go's deadlock
	// detector in action.
	table <- new(Ball)
	time.Sleep(1 * time.Second)

	// Grab the ball back - game ends
	<-table

	// Manually dump stack using panic
	//panic("show me the stacks")
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
