// Code from http://blog.golang.org/pipelines

package main

import (
	"fmt"
	"sync"
)

// This demonstrates how downstream stages can notify upstream senders to stop
// sending / cancel by closing a 'done' channel. A single done channel can work
// for multiple consumers, since we signal we're done by closing the channel -
// this releases the selects in all the consumers. The close is effectively
// a broadcast signal to the consumers.

func gen(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		// This goroutine always closes 'out' before exiting, no matter through
		// which return path.
		defer close(out)
		for _, n := range nums {
			// Try to push a new value to 'out'. Also keep an eye on the 'done'
			// channel - if we could receive from it this means the downstream
			// receivers are done and won't be consuming our values in 'out'
			// any more.
			select {
			case out <- n:
			// A receive on a closed channel can always proceed immediately,
			// yielding the element type's zero value.
			case <-done:
				return
			}
		}
	}()
	return out
}

func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()
	return out
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c or done is closed, then calls
	// wg.Done.
	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	// Set up a done channel that's shared by the whole pipeline,
	// and close that channel when this pipeline exits, as a signal
	// for all the goroutines we started to exit.
	done := make(chan struct{})
	defer close(done)

	in := gen(done, 2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(done, in)
	c2 := sq(done, in)

	// Consume the first value from output.
	out := merge(done, c1, c2)
	fmt.Println(<-out) // 4 or 9

	// done will be closed by the deferred call.
}
